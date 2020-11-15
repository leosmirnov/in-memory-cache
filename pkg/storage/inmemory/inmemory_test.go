package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Operations(t *testing.T) {
	assert := assert.New(t)
	s := NewService()
	assert.Equal([]string{}, s.GetKeys())
	err := s.Set("key", []byte("value"), 3)
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
