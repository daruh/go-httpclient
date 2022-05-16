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

func StartMockServer() {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()
	MockupServer.enabled = true
}

func StopMockServer() {
	MockupServer.enabled = false
}

func (m *mockServer) IsMockServerEnabled() bool {
	return MockupServer.enabled
}

func (m *mockServer) GetMockedClient() core.HttpClient {
	return m.httpClient
}

func FlushMocks() {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()

	MockupServer.mocks = make(map[string]*Mock)
}

func AddMock(mock Mock) {
	MockupServer.serverMutex.Lock()
	MockupServer.serverMutex.Unlock()
	key := MockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	MockupServer.mocks[key] = &mock
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
