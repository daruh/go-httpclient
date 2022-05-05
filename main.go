package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.github.com"
	client := http.Client{}
	response, err := client.Get(url)
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
