package models

import (
	"sync"

	"github.com/pion/webrtc/v3"
)

var (
	PublisherRemoteTracks = map[string]*webrtc.TrackRemote{}
	PublisherLock         sync.Mutex
)

type ViewerState struct {
	PC          *webrtc.PeerConnection
	LocalTracks map[string]*webrtc.TrackLocalStaticRTP
}

var (
	Viewers    = map[string]*ViewerState{}
	ViewerLock sync.Mutex
)
