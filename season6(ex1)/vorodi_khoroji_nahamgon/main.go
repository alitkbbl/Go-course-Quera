package main

import (
	"sync/atomic"
	"time"
)

type FutureResult struct {
	Done       atomic.Bool
	ResultChan chan string
	// TODO
}

type Task func() string

func Async(t Task) *FutureResult {
}
func AsyncWithTimeout(t Task, timeout time.Duration) *FutureResult {
}

func (fResult *FutureResult) Await() string {
}

func CombineFutureResults(fResults ...*FutureResult) *FutureResult {
}
