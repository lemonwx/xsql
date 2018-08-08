/**
 *  author: lim
 *  data  : 18-8-8 下午8:53
 */

package errors

import (
	"fmt"
	"testing"
)

func foo3() error {
	return New(fmt.Errorf("this is foo3's return error"))
}

func foo2() error {
	return foo3()
}

func foo1() error {
	return foo2()
}

func TestMainFunc(t *testing.T) {

	err := foo1()
	t.Log(err)
}
