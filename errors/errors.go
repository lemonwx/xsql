/**
 *  author: lim
 *  data  : 18-8-8 下午8:49
 */

package errors

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	BTSIZE = 4096
)

type Err struct {
	msg string
	bt  string
}

func (err *Err) Error() string {
	return fmt.Sprintf("%s \n%s", err.msg, err.bt)
}

func New(err error) *Err {
	e := &Err{}
	e.msg = err.Error()

	buf := make([]byte, BTSIZE)
	size := runtime.Stack(buf, false)
	stacks := strings.Split(string(buf[:size]), "\n")

	stacks = append(stacks[:1], stacks[3:]...)
	e.bt = strings.Join(stacks, "\n")

	return e
}
