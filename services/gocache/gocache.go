package gocache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Service struct {
	c *cache.Cache
}

func NewService() *Service {
	return &Service{cache.New(cache.NoExpiration, 60*time.Minute)}
}

func (s *Service) Flush() {
	s.c.Flush()
}

func (s *Service) Get(key string) (value any, exists bool) {
	return s.c.Get(key)
}

func (s *Service) Set(key string, value any) {
	s.c.Set(key, value, cache.DefaultExpiration)
}
