package system

import (
	"sync"
	"testing"
)

func TestRecover(t *testing.T) {
	var err error
	var Panic GSystem
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer Panic.Recover(&err)
		panic("test panic")
	}()
	wg.Wait()
	t.Log(err)
}
