package cache

type Cache[K any, V any] interface {
	Put(key K, value V)
	Get(key K) (V, bool)
}
