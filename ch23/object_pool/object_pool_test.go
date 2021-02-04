package objectPool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjectPool(t *testing.T) {
	pool := NewObjectPool(10)

	for i := 0; i < 11; i++ {
		if v, err := pool.GetObject(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			// if err := pool.ReleaseObject(v); err != nil {
			// 	t.Error(err)
			// }
		}
	}
	fmt.Println("Done.")
}
