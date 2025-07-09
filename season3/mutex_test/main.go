package main

import (
	"fmt"
	"sync"
)

type PrimaryStruct struct {
	mu              *sync.Mutex
	wg              *sync.WaitGroup
	primaryVariable int
}

type SecondaryStruct struct {
	mu                *sync.Mutex
	wg                *sync.WaitGroup
	secondaryVariable int
}

func (s *PrimaryStruct) Increment() {
	defer s.wg.Done()
	s.mu.Lock()
	s.primaryVariable += 1
	s.mu.Unlock()
}

func (s *SecondaryStruct) Increment() {
	defer s.wg.Done()
	s.secondaryVariable += 1
}

func main() {
	var wg sync.WaitGroup

	ps := PrimaryStruct{
		mu:              &sync.Mutex{},
		wg:              &wg,
		primaryVariable: 0,
	}

	ss := SecondaryStruct{
		mu:                &sync.Mutex{},
		wg:                &wg,
		secondaryVariable: 0,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go ps.Increment()
		go ss.Increment()
	}

	wg.Wait()
	fmt.Println(ps.primaryVariable)
	fmt.Println(ss.secondaryVariable)

}
