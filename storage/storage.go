package storage

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	SUPABASE_URL = os.Getenv("SUPABASE_URL")
	SUPABASE_KEY = os.Getenv("SUPABASE_KEY")
	BUCKET_NAME  = os.Getenv("SUPABASE_BUCKET")
	client       = &http.Client{}
)

func Init() {
	SUPABASE_URL = os.Getenv("SUPABASE_URL")
	SUPABASE_KEY = os.Getenv("SUPABASE_KEY")
	BUCKET_NAME = os.Getenv("SUPABASE_BUCKET")
}

func UploadFileFromReader(reader io.Reader, destPath string, contentType string) (string, error) {
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", SUPABASE_URL, BUCKET_NAME, destPath)

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+SUPABASE_KEY)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	} else {
		req.Header.Set("Content-Type", "application/octet-stream")
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed: %s", string(b))
	}

	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", SUPABASE_URL, BUCKET_NAME, destPath)
	return publicURL, nil
}
