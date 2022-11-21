package main

import (
	"fmt"
	"sync"
)

type InMemoryCache struct {
	cache map[int]string
	mx    sync.RWMutex
}

func (c *InMemoryCache) Get(key int) (string, bool) {
	c.mx.RLock()
	data, ok := c.cache[key]
	c.mx.RUnlock()
	return data, ok
}

func (c *InMemoryCache) Put(key int, content string) {

}

func (c *InMemoryCache) Delete(key int) {

}

type SafeCounter struct {
	counter int
	mx      sync.RWMutex
}

func (s *SafeCounter) Inc() {
	s.mx.Lock()
	s.counter++
	s.mx.Unlock()
}

func (s *SafeCounter) String() string {
	s.mx.RLock()
	result := s.counter
	s.mx.RUnlock()
	return fmt.Sprintf("counter is %d", result)
}

func main() {
	println("Lesson-04")

	counter := &SafeCounter{counter: 0}

	for i := 0; i < 1000; i++ {
		go counter.Inc()
	}

	println("for completed")

	var s string
	fmt.Scanln(&s)

	fmt.Println(counter)
}
