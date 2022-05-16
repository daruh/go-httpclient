package examples

import (
	"errors"
	"fmt"
	"github.com/daruh/go-httpclient/gohttp_mock"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for packages")
	// Tell the HTTP library to mock any further requests from here
	gohttp_mock.MockupServer.Start()
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		//init
		gohttp_mock.MockupServer.FlushMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
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

		if err.Error() != "timeout getting github endpoint" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestUnmarshalResponseBody", func(t *testing.T) {
		gohttp_mock.MockupServer.FlushMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseBody:       `{"current_user_url": "123"}`,
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
	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttp_mock.MockupServer.FlushMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
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
	})
	gohttp_mock.MockupServer.Stop()
}
