package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"
	"strings"
)

var ErrNotFound = errors.New("not found")
var ErrCacheDisable = errors.New("cache is disable")
var CACHE *Memory

const (
	KcachePrefixHome    = "@h"
	KcachePrefixArticle = "@a"
	KcachePrefixCate    = "@c"
	KcachePrefixTag     = "@t"
)

type Memory struct {
	sync.RWMutex
	values map[string][]byte
}

func InitCache(minute int) {
	CACHE = &Memory{values: map[string][]byte{}}
	CACHE.timer(minute)
}

func MakeCacheKey(args... string) string {
	return strings.Join(args, "@")
}

func (memory *Memory) Get(key string) (string, error) {
	if memory == nil {
		return "", ErrCacheDisable
	}

	memory.RLock()
	defer memory.RUnlock()

	if value, ok := memory.values[key]; ok {
		return string(value), nil
	}
	return "", ErrNotFound
}

func (memory *Memory) Unmarshal(key string, object interface{}) error {
	if memory == nil {
		return ErrCacheDisable
	}

	memory.RLock()
	defer memory.RUnlock()

	if value, ok := memory.values[key]; ok {
		return json.Unmarshal(value, object)
	}
	return ErrNotFound
}

func (memory *Memory) Set(key string, value interface{}) {
	if memory == nil {
		return
	}

	memory.Lock()
	memory.values[key] = convertToBytes(value)
	memory.Unlock()
}

func (memory *Memory) Delete(key string) {
	memory.Lock()
	delete(memory.values, key)
	memory.Unlock()
}

// 缓存定时任务
func (memory *Memory) timer(minute int) {
	ticker := time.NewTicker(time.Duration(minute) * time.Minute)
	go func() {
		for _ = range ticker.C {
			log.Println("clear cache....")
			memory.clear()
		}
	}()
}

func (memory *Memory) clear() {
	memory.Lock()
	memory.values = map[string][]byte{}
	memory.Unlock()
}

func convertToBytes(value interface{}) []byte {
	switch result := value.(type) {
	case string:
		return []byte(result)
	case []byte:
		return result
	default:
		bytes, _ := json.Marshal(value)
		return bytes
	}
}
