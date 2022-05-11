package gohttp

import (
	"fmt"
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

func TestBodyType(t *testing.T) {
	client := httpClient{}

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		fmt.Println(err)
		fmt.Println(string(body))

		//assert
		if err != nil {
			t.Error("no error expected when marshalled")
		}
		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithXML", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/xml", requestBody)

		fmt.Println(err)
		fmt.Println(string(body))

		//assert
		if err != nil {
			t.Error("no error expected when marshalled")
		}
		if string(body) != "<string>one</string><string>two</string>" {
			t.Error("invalid xml obtained")
		}
	})

	t.Run("BodyWithDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("", requestBody)

		fmt.Println(err)
		fmt.Println(string(body))

		//assert
		if err != nil {
			t.Error("no error expected when marshalled")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})
}
