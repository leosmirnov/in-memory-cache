package inmemory

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestService_Operations(t *testing.T) {
	assert := assert.New(t)
	cleanupInterval := 500 * time.Millisecond

	s := NewService(logrus.New(), &cleanupInterval)
	assert.Equal([]string{}, s.GetKeys())
	err := s.Set("key", []byte("value"), 1*time.Minute)
	assert.NoError(err)
	v, err := s.GetByKey("key")
	assert.NoError(err)
	assert.Equal(v, []byte("value"))
	err = s.RemoveKey("INVALID")
	assert.Error(err)
	err = s.RemoveKey("key")
	assert.NoError(err)
	assert.Equal([]string{}, s.GetKeys())
}

func TestService_Cleanup(t *testing.T) {
	assert := assert.New(t)
	cleanupInterval := 10 * time.Microsecond

	s := NewService(logrus.New(), &cleanupInterval)
	assert.Equal([]string{}, s.GetKeys())
	err := s.Set("key", []byte("value"), 50*time.Microsecond)
	assert.NoError(err)
	v, err := s.GetByKey("key")
	assert.NoError(err)
	assert.Equal([]byte("value"), v)
	time.Sleep(70 * time.Microsecond)
	_, err = s.GetByKey("key")
	assert.Error(err)
}
