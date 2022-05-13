package examples

import (
	"errors"
	"github.com/daruh/go-httpclient/gohttp"
	"net/http"
	"testing"
)

func TestGetEndpoints(t *testing.T) {

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		//init
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoint"),
		})

		endpoints, err := GetEndpoints()

		//validate
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestUnmarshalResponseBody", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseBody:       `{"current_user_url": 123}`,
			ResponseStatusCode: http.StatusOK,
		})
		endpoints, err := GetEndpoints()

		//validate
		if err != nil {
			t.Error("no error was expected")
		}
		if endpoints == nil {
			t.Error("endpoints were expected and we got nil")
		}
		if err.Error() != "json unmarshal error" {
			t.Error("invalid error message received")
		}

	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
			ResponseStatusCode: http.StatusOK,
		})
		endpoints, err := GetEndpoints()

		//validate
		if err != nil {
			t.Error("no error was expected")
		}
		if endpoints == nil {
			t.Error("endpoints were expected and we got nil")
		}
		if err.Error() != "json unmarshal error" {
			t.Error("invalid error message received")
		}
	})
}
