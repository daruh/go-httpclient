package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers http.Header

	maxIdleConnection int
	connectionTimeout time.Duration
	responseTimeout   time.Duration
	disableTimeouts   bool

	client *http.Client
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disableTimeouts bool) ClientBuilder
	SetHttpClient(client *http.Client) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	clientBuilder := &clientBuilder{}
	return clientBuilder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnection = i
	return c
}

func (c *clientBuilder) DisableTimeouts(disableTimeouts bool) ClientBuilder {
	c.disableTimeouts = disableTimeouts
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}
