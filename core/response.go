package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
	Status     string
}

func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) String() string {
	return string(r.Body)
}
func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
