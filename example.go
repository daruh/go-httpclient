package main

import (
	"fmt"
	"github.com/daruh/go-httpclient/gohttp"
	"time"
)

var client = getGithubClient()

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	return client
}

func main() {

	for i := 0; i < 20; i++ {
		go func() {
			getUrls()
		}()
	}
	time.Sleep(20 * time.Second)
}
func getUrls() {
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status())
}

func createUser(user User) {
	response, err := client.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

}
