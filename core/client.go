package core

import "net/http"

//small comment

type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}
