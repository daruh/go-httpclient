package go_httpclient

import "github.com/daruh/go-httpclient/gohttp"

func basicExample() {
	client := gohttp.New()

	client.Get()
}
