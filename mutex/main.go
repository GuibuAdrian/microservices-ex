package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

var (
	counter       = 0
	aux           atomic.Value
	lock          sync.Mutex
	atomicCounter = AtomicInt{}
)

func (i *AtomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *AtomicInt) Decrease() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value--
}

func (i *AtomicInt) Value() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("final counter :%d", counter))
	fmt.Println(fmt.Sprintf("final atomic counter value: %d", atomicCounter.Value()))
}

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()

	atomicCounter.Decrease()
	wg.Done()
}
