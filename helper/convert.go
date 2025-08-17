package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type ffprobeOutput struct {
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

func GetVideoDurationFromURL(videoURL string) (float64, error) {
	resp, err := http.Get(videoURL)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch video: %w", err)
	}
	defer resp.Body.Close()

	tmpFile, err := os.CreateTemp("", "video-*.mp4")
	if err != nil {
		return 0, err
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to save temp video: %w", err)
	}
	tmpFile.Close()

	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", tmpFile.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("ffprobe failed: %w", err)
	}

	var probe ffprobeOutput
	if err := json.Unmarshal(out.Bytes(), &probe); err != nil {
		return 0, fmt.Errorf("json parse failed: %w", err)
	}

	return strconv.ParseFloat(probe.Format.Duration, 64)
}
