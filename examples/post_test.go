package examples

import (
	"fmt"
	"github.com/daruh/go-httpclient/gohttp_mock"
	"net/http"
	"testing"
)

type Repository struct {
	Name string `json:"name"`
}

func TestPost(t *testing.T) {
	repo := Repository{
		Name: "testing_repo",
	}

	gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
		Method:             http.MethodPost,
		Url:                "https://api.github.com",
		ResponseBody:       `{"current_user_url": "123"}`,
		ResponseStatusCode: http.StatusOK,
	})

	response, err := httpClient.Post("https://api.github.com", repo, nil)

	fmt.Println(err)
	fmt.Println(response)
}
