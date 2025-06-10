package cache

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"
	"os"
	"path/filepath"
)

type FileCache struct {
	Dir    string
	Prefix string
	Hasher hash.Hash32
}

func MakeTempFileCache(prefix string) FileCache {
	return FileCache{
		Dir:    os.TempDir(),
		Prefix: prefix,
		Hasher: fnv.New32a(),
	}
}

func encodeKey(hasher hash.Hash32, key string, prefix string) string {
	hasher.Reset()
	hasher.Write([]byte(key))
	hashValue := hasher.Sum32()
	return fmt.Sprintf("%s%d", prefix, hashValue)
}

func (f FileCache) cacheFile(key string) string {
	fileName := encodeKey(f.Hasher, key, f.Prefix)
	return filepath.Join(f.Dir, fileName)
}

func (f FileCache) Put(key string, value []byte) {
	file := f.cacheFile(key)
	err := os.WriteFile(file, value, 0644)
	if err != nil {
		log.Println(err.Error())
	}
}

func (f FileCache) Get(key string) ([]byte, bool) {
	file := f.cacheFile(key)

	if data, err := os.ReadFile(file); err == nil {
		return data, true
	}
	return nil, false
}
