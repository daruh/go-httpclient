package gohttp_mock

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/daruh/go-httpclient/core"
	"strings"
	"sync"
)

var (
	MockupServer = mockServer{
		mocks:      make(map[string]*Mock),
		httpClient: &(httpClientMock{}),
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	httpClient  core.HttpClient
	mocks       map[string]*Mock
}

func (m *mockServer) Start() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()
	m.enabled = true
}

func (m *mockServer) Stop() {
	m.enabled = false
}

func (m *mockServer) IsMockServerEnabled() bool {
	return m.enabled
}

func (m *mockServer) GetMockedClient() core.HttpClient {
	return m.httpClient
}

func (m *mockServer) FlushMocks() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.mocks = make(map[string]*Mock)
}

func (m *mockServer) AddMock(mock Mock) {
	m.serverMutex.Lock()
	m.serverMutex.Unlock()
	key := m.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	m.mocks[key] = &mock
}

func (m *mockServer) cleanBody(body string) string {

	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) GetMock(method, url, body string) *Mock {
	if !m.enabled {
		return nil
	}
	if mock := m.mocks[m.getMockKey(method, url, body)]; mock != nil {
		return mock
	}
	return &Mock{
		Error: errors.New(fmt.Sprintf("no mocks matching %s from '%s' with given body", method, url)),
	}
}
