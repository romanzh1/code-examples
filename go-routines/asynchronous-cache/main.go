package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrNotFound = errors.New("value not found")
var mu = sync.Mutex{}

type Cache map[string]string

func NewCache() Cache {
	cache := make(Cache)

	return cache
}

func (c Cache) Set(key, value string) error {
	if c == nil {
		return fmt.Errorf("cache isn't initialize")
	}

	mu.Lock()
	c[key] = value
	mu.Unlock()

	return nil
}

func (c Cache) Get(key string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("cache isn't initialize")
	}

	mu.Lock()
	val, ok := c[key]
	if !ok {
		return "", ErrNotFound
	}
	mu.Unlock()

	return val, nil
}

func (c Cache) Delete(key string) error {
	if c == nil {
		return fmt.Errorf("cache isn't initialize")
	}

	mu.Lock()
	_, ok := c[key]
	if !ok {
		return ErrNotFound
	}

	delete(c, key)
	mu.Unlock()

	return nil
}

func main() {
	cache := NewCache()

	go func() {
		err := cache.Set("chi", "10")
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		err := cache.Set("gorm", "5")
		if err != nil {
			fmt.Println(err)
		}
	}()
	time.Sleep(time.Millisecond)

	go func() {
		err := cache.Set("fasthttp", "3")
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		gorm, err := cache.Get("gorm")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(gorm)
	}()

	go func() {
		err := cache.Delete("fasthttp")
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		fh, err := cache.Get("fasthttp")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fh)
	}()

	fmt.Println(cache)
	time.Sleep(time.Second)
}
