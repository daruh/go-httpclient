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
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disableTimeouts bool) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	clientBuilder := &clientBuilder{}
	return clientBuilder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		headers:           c.headers,
		maxIdleConnection: c.maxIdleConnection,
		connectionTimeout: c.connectionTimeout,
		responseTimeout:   c.responseTimeout,
		disableTimeouts:   c.disableTimeouts,
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
