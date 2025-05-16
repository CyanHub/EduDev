package pool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	pool, err := NewPool(10)
	if err != nil {
		t.Fatal(err)	
	}
	for i := 0; i < 100; i++ {
		v := i
		pool.Submit(func() {
			fmt.Println("task", v)
		})
	}
}
