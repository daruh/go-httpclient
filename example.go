package main

import (
	"fmt"
	"github.com/daruh/go-httpclient/gohttp"
	"io/ioutil"
)

func main() {
	client := gohttp.New()

	resposne, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resposne.StatusCode)
	bytes, _ := ioutil.ReadAll(resposne.Body)
	fmt.Println(string(bytes))
}
