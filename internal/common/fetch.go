package common

import (
	"encoding/base64"
	"hash"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Hashable[T any] interface {
	GetHash(value T) string
}

type (
	Fetch func(string) ([]byte, error)
)

func FetchFrom(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (f Fetch) ReadThrough(hashOf hash.Hash64, ttl time.Duration) Fetch {
	return func(url string) ([]byte, error) {
		fileName := base64.StdEncoding.EncodeToString(hashOf.Sum([]byte(url)))
		file := filepath.Join(os.TempDir(), fileName)

		info, err := os.Stat(file)
		if err == nil && time.Since(info.ModTime()) < ttl {
			return os.ReadFile(file)
		}
		body, err := f(url)
		if err != nil {
			return nil, err
		}
		_ = os.WriteFile(file, body, 0644)
		return body, nil
	}
}
