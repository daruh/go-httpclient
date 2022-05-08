package main

import (
	"fmt"
	"github.com/daruh/go-httpclient/gohttp"
	"io/ioutil"
	"net/http"
)

var client = getGithubClient()

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)
	return client
}

func main() {

	getUrls()
}
func getUrls() {
	resposne, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resposne.StatusCode)
	bytes, _ := ioutil.ReadAll(resposne.Body)
	fmt.Println(string(bytes))
}
