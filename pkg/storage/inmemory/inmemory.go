package inmemory

import (
	"errors"
	"sync"
	"time"
)

type Service struct {
	*sync.RWMutex
	cache map[string]item
}

type item struct {
	value      []byte
	expiration time.Duration
}

func NewService() *Service {
	return &Service{
		RWMutex: &sync.RWMutex{},
		cache:   make(map[string]item),
	}
}

func (s *Service) GetByKey(key string) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()

	value, exists := s.cache[key]
	if !exists {
		return nil, errors.New("no such key")
	}

	return value.value, nil
}

func (s *Service) Set(key string, value []byte, exp time.Duration) error {
	s.Lock()
	defer s.Unlock()

	_, exists := s.cache[key]
	if exists {
		return errors.New("key already exists")
	}

	s.cache[key] = item{
		value:      value,
		expiration: exp,
	}
	return nil
}

func (s *Service) GetKeys() []string {
	s.RLock()
	defer s.RUnlock()

	keys := make([]string, 0, len(s.cache))
	for k := range s.cache {
		keys = append(keys, k)
	}

	return keys
}

func (s *Service) RemoveKey(key string) error {
	s.Lock()
	defer s.Unlock()

	_, exists := s.cache[key]
	if !exists {
		return errors.New("key %s does not exist")
	}
	delete(s.cache, key)
	return nil
}
