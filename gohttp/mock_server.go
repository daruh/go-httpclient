package gohttp

import "sync"

var (
	mockupServer = mockServer{}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	mocks map[string]*Mock
}

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()
	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.enabled = false
}

func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	mockupServer.serverMutex.Unlock()
	key := mock.Method + mock.Url + mock.RequestBody
	mockupServer.mocks[key] = &mock
}
