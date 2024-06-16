package sync_

import (
	"fmt"
	"testing"
	"unsafe"
)

type Aa struct {
	a int
	b float64
	c []int
}

func TestSemaphore(t *testing.T) {
	c := Aa{}
	fmt.Println(unsafe.Sizeof(c))
	c.a = 55
	c.b = 99
	fmt.Println(unsafe.Sizeof(c))

}
