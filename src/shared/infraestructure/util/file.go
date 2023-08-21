package util

import (
	"errors"
	"io"
	"net/http"

	"github.com/getsentry/sentry-go"
)

func ReadFileFromURL(url string) (content []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("404")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return body, nil
}
