package domain

type CacheService interface {
	Flush()
	Get(key string) (value any, exists bool)
	Set(key string, value any)
}
