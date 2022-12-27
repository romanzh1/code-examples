package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var ErrNotFound = errors.New("value not found")

type Cache struct {
	cache map[string]string
	mu    *sync.Mutex
}

func NewCache() Cache {
	cache := Cache{
		cache: make(map[string]string),
		mu:    &sync.Mutex{},
	}

	return cache
}

func (c *Cache) Set(key, value string) error {
	if c == nil {
		return fmt.Errorf("cache isn't initialize")
	}

	c.mu.Lock()
	c.cache[key] = value
	c.mu.Unlock()

	return nil
}

func (c *Cache) Get(key string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("cache isn't initialize")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.cache[key]
	if !ok {
		return "", ErrNotFound
	}
	// mu.Unlock() unlock an unlocked mutex without defer TODO check it out

	return val, nil
}

func (c *Cache) Delete(key string) error {
	if c == nil {
		return fmt.Errorf("cache isn't initialize")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.cache[key]
	if !ok {
		return ErrNotFound
	}

	delete(c.cache, key)
	return nil
}

func CacheWork(cache *Cache) {
	logger := log.New(os.Stderr, "logger: ", log.Ltime|log.Lshortfile)

	go func() {
		err := cache.Set("chi", "10")
		if err != nil {
			logger.Println(err)
		}
	}()

	go func() {
		err := cache.Set("gorm", "5")
		if err != nil {
			logger.Println(err)
		}
	}()
	time.Sleep(time.Millisecond)

	go func() {
		err := cache.Set("fasthttp", "3")
		if err != nil {
			logger.Println(err)
		}
	}()

	go func() {
		gorm, err := cache.Get("gorm")
		if err != nil {
			logger.Println(err)
		}
		fmt.Println(gorm)
	}()

	go func() {
		err := cache.Delete("fasthttp")
		if err != nil {
			logger.Println(err)
		}
	}()

	go func() {
		fh, err := cache.Get("fasthttp")
		if err != nil {
			logger.Println(err)
		}
		fmt.Println(fh)
	}()
}

func main() {
	cache := NewCache()

	CacheWork(&cache)

	time.Sleep(time.Second)
	fmt.Println(cache)
}
