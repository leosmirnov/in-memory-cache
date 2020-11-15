package api

import (
	"io"
	"net/http/httptest"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockKVService struct {
	mock.Mock
}

func (m *mockKVService) GetByKey(key string) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *mockKVService) Set(key string, value []byte, exp time.Duration) error {
	args := m.Called(key, value, exp)
	return args.Error(0)
}

func (m *mockKVService) GetKeys() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

func (m *mockKVService) RemoveKey(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *mockKVService) Close() error {
	args := m.Called()
	return args.Error(0)
}

type httpTests struct {
	name       string
	url        string
	expCode    int
	body       io.Reader
	mock       *mockKVService
	onResponse func(resp *httptest.ResponseRecorder)
}
