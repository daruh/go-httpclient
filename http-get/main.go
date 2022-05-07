package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://localhost:8080"
	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
