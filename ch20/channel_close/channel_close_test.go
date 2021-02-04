package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(n int, ch chan int, wait *sync.WaitGroup) {
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
		wait.Done()
	}()
}

func dataReceiver(ch chan int, wait *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wait.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wait sync.WaitGroup
	ch := make(chan int)
	wait.Add(1)
	dataProducer(20, ch, &wait)
	wait.Add(1)
	dataReceiver(ch, &wait)
	wait.Add(1)
	dataReceiver(ch, &wait)
	wait.Add(1)
	dataReceiver(ch, &wait)
	wait.Wait()
}
