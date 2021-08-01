package pipeline

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SearchWord ...
func SearchWord(urls <-chan string, word string) (int, error) {
	var result int

	//TODO
	return result, fmt.Errorf("not implemented")
}

// FetchDocument ...
func FetchDocument(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body == nil {
			return
		}

		_ = resp.Body.Close()
	}()

	return io.ReadAll(resp.Body)
}
