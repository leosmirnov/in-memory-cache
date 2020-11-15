package inmemory

import (
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const defaultExpiration = 5 * time.Minute

type Service struct {
	*sync.RWMutex
	cache           map[string]item
	cleanupInterval *time.Duration
	quitChan        chan interface{}
	logger          logrus.FieldLogger
}

type item struct {
	value      []byte
	expiration int64
}

func NewService(logger logrus.FieldLogger, cleanupInterval *time.Duration) *Service {
	s := &Service{
		RWMutex:         &sync.RWMutex{},
		cache:           make(map[string]item),
		cleanupInterval: cleanupInterval,
		logger:          logger,
		quitChan:        make(chan interface{}),
	}
	go s.startMonitor()
	return s
}

func (s *Service) Close() error {
	s.quitChan <- struct{}{}
	s.logger.Debug("in-memory service has successfully shut down")
	return nil
}

func (s *Service) GetByKey(key string) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()

	value, exists := s.cache[key]
	if !exists || time.Now().UnixNano() > value.expiration {
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

	if exp <= 0 {
		exp = defaultExpiration
	}

	s.cache[key] = item{
		value:      value,
		expiration: time.Now().Add(exp).UnixNano(),
	}
	return nil
}

func (s *Service) GetKeys() []string {
	s.RLock()
	defer s.RUnlock()

	keys := make([]string, 0, len(s.cache))
	for k, v := range s.cache {
		if time.Now().UnixNano() < v.expiration {
			keys = append(keys, k)
		}
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

func (s *Service) startMonitor() {
	for {
		select {
		case <-time.After(*s.cleanupInterval):
		case <-s.quitChan:
			return
		}

		if keys := s.expiredKeys(); len(keys) != 0 {
			s.clearItems(keys)
		}
	}
}

func (s *Service) expiredKeys() (keys []string) {
	s.RLock()
	defer s.RUnlock()

	for k, i := range s.cache {
		if time.Now().UnixNano() > i.expiration {
			keys = append(keys, k)
		}
	}
	return
}

func (s *Service) clearItems(keys []string) {
	s.Lock()
	defer s.Unlock()

	for _, k := range keys {
		delete(s.cache, k)
	}
}
