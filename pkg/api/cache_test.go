package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/leosmirnov/in-memory-cache/pkg/api/types"
	"github.com/leosmirnov/in-memory-cache/pkg/conf"
	"github.com/leosmirnov/in-memory-cache/pkg/constants"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAPI_AddValue(t *testing.T) {
	assert := assert.New(t)
	tests := []httpTests{
		{name: "Successfully adding a new key",
			expCode: http.StatusCreated,
			mock: func() *mockKVService {
				m := new(mockKVService)
				m.On("Set", "key", []byte("value"), 1*time.Second).Return(nil)
				return m
			}(),
			url:  "/values",
			body: bytes.NewBuffer([]byte(`{"key":"key","value":"value","expiration":"1s"}`)),
		},
		{name: "Key was already added",
			expCode: http.StatusConflict,
			mock: func() *mockKVService {
				m := new(mockKVService)
				m.On("Set", "key", []byte("value"), 1*time.Second).Return(errors.New(""))
				return m
			}(),
			url:  "/values",
			body: bytes.NewBuffer([]byte(`{"key":"key","value":"value","expiration":"1s"}`)),
		},
		{name: "Body is empty",
			expCode: http.StatusBadRequest,
			url:     "/values",
		},
		{name: "Invalid body",
			expCode: http.StatusBadRequest,
			url:     "/values",
			body:    bytes.NewBuffer([]byte(`D":"key","vaue":"value","expsdiration":"1s"}`)),
		},
		{name: "Invalid body request",
			expCode: http.StatusBadRequest,
			url:     "/values",
			body:    bytes.NewBuffer([]byte(`{"INVALID":"key","INVALID":"value","expiration":"2s"}`)),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := New(&conf.API{}, logrus.New(), tt.mock)
			router := httprouter.New()
			router.POST("/values", s.AddValue)

			req, _ := http.NewRequest(http.MethodPost, tt.url, tt.body)
			req.Header.Set(constants.ContentType, constants.ApplicationJSON)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			assert.Equal(tt.expCode, rr.Code, t.Name())
		})
	}
}

func TestAPI_DeleteValue(t *testing.T) {
	assert := assert.New(t)
	tests := []httpTests{
		{name: "Successful deletion",
			expCode: http.StatusOK,
			mock: func() *mockKVService {
				m := new(mockKVService)
				m.On("RemoveKey", "key").Return(nil)
				return m
			}(),
			url: "/values/key",
		},
		{name: "Value was already deleted",
			expCode: http.StatusNotFound,
			mock: func() *mockKVService {
				m := new(mockKVService)
				m.On("RemoveKey", "key").Return(errors.New(""))
				return m
			}(),
			url: "/values/key",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := New(&conf.API{}, logrus.New(), tt.mock)
			router := httprouter.New()
			router.DELETE("/values/:key", s.DeleteValue)

			req, _ := http.NewRequest(http.MethodDelete, tt.url, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			assert.Equal(tt.expCode, rr.Code, t.Name())
		})
	}
}

func TestAPI_GetValue(t *testing.T) {
	assert := assert.New(t)
	tests := []httpTests{
		{name: "Get non created key error",
			expCode: http.StatusNotFound,
			mock: func() *mockKVService {
				m := new(mockKVService)
				m.On("GetByKey", "key").Return(make([]byte, 0), errors.New("error"))
				return m
			}(),
			url: "/values/key",
		},
		{name: "Get values by created key",
			expCode: http.StatusOK,
			mock: func() *mockKVService {
				m := new(mockKVService)
				m.On("GetByKey", "key").Return([]byte{0, 1, 2}, nil)
				return m
			}(),
			url: "/values/key",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := New(&conf.API{}, logrus.New(), tt.mock)
			router := httprouter.New()
			router.GET("/values/:key", s.GetValue)

			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			assert.Equal(tt.expCode, rr.Code, t.Name())
		})
	}
}

func TestAPI_ListKeys(t *testing.T) {
	assert := assert.New(t)
	tests := []httpTests{
		{name: "Get list of keys",
			expCode: http.StatusOK,
			mock: func() *mockKVService {
				m := new(mockKVService)
				m.On("GetKeys").Return([]string{"1", "2"})
				return m
			}(),
			url: "/keys",
			onResponse: func(resp *httptest.ResponseRecorder) {
				body := types.DataBody{}
				assert.Nil(json.NewDecoder(resp.Body).Decode(&body))
				assert.NotEmpty(body.Data)

				list := body.Data.(map[string]interface{})["keys"].([]interface{})
				assert.Len(list, 2)
				assert.Equal("1", list[0].(string))
				assert.Equal("2", list[1].(string))
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := New(&conf.API{}, logrus.New(), tt.mock)
			router := httprouter.New()
			router.GET("/keys", s.ListKeys)

			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			assert.Equal(tt.expCode, rr.Code, t.Name())

			if tt.onResponse != nil {
				tt.onResponse(rr)
			}
		})
	}
}
