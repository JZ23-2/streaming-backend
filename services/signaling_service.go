package services

import (
	"fmt"
	"log"
	"main/models"
	"time"

	"github.com/pion/webrtc/v3"
)

func CreatePublisherPC(offer webrtc.SessionDescription) (*webrtc.PeerConnection, *webrtc.SessionDescription, error) {
	pc, err := webrtc.NewPeerConnection(webrtc.Configuration{})

	if err != nil {
		return nil, nil, err
	}

	pc.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		log.Printf("Publisher track : kind=%s, id=%s codec=%s, ", track.Kind(), track.ID(), track.Codec().MimeType)

		models.PublisherLock.Lock()
		models.PublisherRemoteTracks[track.ID()] = track
		models.PublisherLock.Unlock()

		buf := make([]byte, 1500)
		for {
			n, _, readErr := track.Read(buf)
			if readErr != nil {
				log.Printf("Publisher track read error (%s): %v", track.ID(), readErr)
				models.PublisherLock.Lock()
				delete(models.PublisherRemoteTracks, track.ID())
				models.PublisherLock.Unlock()
				return
			}

			models.ViewerLock.Lock()
			for vid, vs := range models.Viewers {
				if lt, ok := vs.LocalTracks[track.ID()]; ok {
					if _, werr := lt.Write(buf[:n]); werr != nil {
						log.Printf("Write to viewer %s failed: %v", vid, werr)
					}
				}
			}
			models.ViewerLock.Unlock()
		}
	})

	if err := pc.SetRemoteDescription(offer); err != nil {
		return nil, nil, err
	}

	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pc.SetLocalDescription(answer); err != nil {
		return nil, nil, err
	}

	<-webrtc.GatheringCompletePromise(pc)
	return pc, pc.LocalDescription(), nil
}

func CreateViewerPC(offer webrtc.SessionDescription) (*webrtc.PeerConnection, *webrtc.SessionDescription, string, error) {
	waitTimeOut := 30 * time.Second

	deadline := time.Now().Add(waitTimeOut)

	for {
		models.PublisherLock.Lock()
		ready := len(models.PublisherRemoteTracks) > 0
		models.PublisherLock.Unlock()
		if ready {
			break
		}

		if time.Now().After(deadline) {
			return nil, nil, "", fmt.Errorf("no publisher available")
		}
		time.Sleep(100 * time.Millisecond)
	}

	pc, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		return nil, nil, "", err
	}

	vs := &models.ViewerState{
		PC:          pc,
		LocalTracks: map[string]*webrtc.TrackLocalStaticRTP{},
	}

	models.PublisherLock.Lock()
	for pid, remote := range models.PublisherRemoteTracks {
		localTrack, err := webrtc.NewTrackLocalStaticRTP(remote.Codec().RTPCodecCapability, remote.ID(), remote.StreamID())

		if err != nil {
			models.PublisherLock.Unlock()
			return nil, nil, "", err
		}

		if _, err := pc.AddTrack(localTrack); err != nil {
			models.PublisherLock.Unlock()
			return nil, nil, "", err
		}

		vs.LocalTracks[pid] = localTrack
	}
	models.PublisherLock.Unlock()

	pc.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		log.Printf("Viewer PC state: %s", state.String())
		if state == webrtc.PeerConnectionStateFailed || state == webrtc.PeerConnectionStateClosed || state == webrtc.PeerConnectionStateDisconnected {
			models.ViewerLock.Lock()

			for k, v := range models.Viewers {
				if v.PC == pc {
					delete(models.Viewers, k)
					break
				}
			}
			models.ViewerLock.Unlock()
			_ = pc.Close()
		}
	})

	vid := fmt.Sprintf("%p", pc)
	models.ViewerLock.Lock()
	models.Viewers[vid] = vs
	models.ViewerLock.Unlock()

	if err := pc.SetRemoteDescription(offer); err != nil {
		models.ViewerLock.Lock()
		delete(models.Viewers, vid)
		models.ViewerLock.Unlock()
		return nil, nil, "", err
	}

	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		models.ViewerLock.Lock()
		delete(models.Viewers, vid)
		models.ViewerLock.Unlock()
		return nil, nil, "", err
	}

	if err := pc.SetLocalDescription(answer); err != nil {
		models.ViewerLock.Lock()
		delete(models.Viewers, vid)
		models.ViewerLock.Unlock()
		return nil, nil, "", err
	}

	<-webrtc.GatheringCompletePromise(pc)
	return pc, pc.LocalDescription(), vid, nil
}
