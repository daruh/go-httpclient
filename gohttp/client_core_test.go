package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//arrange
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	//act
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-ID", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	//assert
	if len(finalHeaders) != 3 {
		t.Error("we expect 3 header")
	}
}
