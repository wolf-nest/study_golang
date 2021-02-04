package objectPool

import (
	"errors"
	"time"
)

//
type ReusableObject struct {
}

//
type ObjectPool struct {
	bufferChannle chan *ReusableObject
}

//
func NewObjectPool(number int) *ObjectPool {
	objectPool := ObjectPool{}
	objectPool.bufferChannle = make(chan *ReusableObject, number)
	for i := 0; i < number; i++ {
		objectPool.bufferChannle <- &ReusableObject{}
	}
	return &objectPool
}

//
func (pool *ObjectPool) GetObject(timeout time.Duration) (*ReusableObject, error) {
	select {
	case result := <-pool.bufferChannle:
		return result, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout....")
	}
}

func (pool *ObjectPool) ReleaseObject(object *ReusableObject) error {
	select {
	case pool.bufferChannle <- object:
		return nil
	default:
		return errors.New("overflow")
	}
}
