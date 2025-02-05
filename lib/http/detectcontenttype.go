package libhttp

import (
	"bytes"
	"net/http"
)

func DetectContentType(data []byte) string {
	const sniffLen int = 2048

	if sniffLen < len(data) {
		data = data[:sniffLen]
	}

	switch {
	case bytes.Contains(data, magicActivityStreams):
		return "application/activity+json"
	default:
		return http.DetectContentType(data)
	}
}
