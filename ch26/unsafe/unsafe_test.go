package unsafe_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Customer struct {
	Name string
	Age  int
}

func TestUsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

// The cases is suitable fo unsafe
type MyInt int

// 合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

//原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufferPointer unsafe.Pointer
	writeDataFunc := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufferPointer, unsafe.Pointer(&data))
	}

	readDataFunc := func() {
		data := atomic.LoadPointer(&shareBufferPointer)
		fmt.Println(data, *(*[]int)(data))
	}

	var wait sync.WaitGroup
	writeDataFunc()
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFunc()
				time.Sleep(time.Millisecond * 100)
			}
			wait.Done()
		}()
		wait.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFunc()
				time.Sleep(time.Millisecond * 100)
			}
			wait.Done()
		}()
	}
	wait.Wait()
}
