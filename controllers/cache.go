package controllers

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"
)

var ErrNotFound = errors.New("not found")
var ErrCacheDisable = errors.New("cache is disable")
var CACHE *Memory

const (
	KcachePrefixHome    = "@h"
	KcachePrefixArticle = "@a"
	KcachePrefixCate    = "@c"
	KcachePrefixTag     = "@t"
	KcachePrefixTopic   = "@o"
	KcachePrefixLang    = "@l"
	KcachePrefixDb      = "@d"
)

type Memory struct {
	sync.RWMutex
	values map[string]interface{}
}

func InitCache(minute int) {
	CACHE = &Memory{values: map[string]interface{}{}}
	CACHE.timer(minute)
}

func MakeCacheKey(args ...string) string {
	return strings.Join(args, "@")
}

func (memory *Memory) Get(key string) (interface{}, error) {
	if memory == nil {
		return nil, ErrCacheDisable
	}

	memory.RLock()
	defer memory.RUnlock()

	if value, ok := memory.values[key]; ok {
		return value, nil
	}
	return nil, ErrNotFound
}

func (memory *Memory) Set(key string, value interface{}) {
	if memory == nil {
		return
	}

	memory.Lock()
	memory.values[key] = value
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
	memory.values = map[string]interface{}{}
	memory.Unlock()
}
