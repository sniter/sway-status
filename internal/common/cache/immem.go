package cache

type InMemCache[K comparable, V any] struct {
	cache map[K]V
}

func (i InMemCache[K, V]) Put(key K, value V) {
	i.cache[key] = value
}

func (i InMemCache[K, V]) Get(key K) (V, bool) {
	value, ok := i.cache[key]
	return value, ok
}

func MakeInMemCache[K comparable, V any]() InMemCache[K, V] {
	return InMemCache[K, V]{
		make(map[K]V),
	}
}
