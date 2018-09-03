/**
 *  author: lim
 *  data  : 18-9-3 下午6:58
 */

package quicklock

type QuickLock struct {
	ch chan bool
}

func (ql *QuickLock) Lock() bool {
	select {
	case ql.ch <- false:
		return true
	default:
		return false
	}
}

func (ql *QuickLock) UnLock() {
	<- ql.ch
}

func NewQL() *QuickLock{
	return &QuickLock{ch: make(chan bool, 1)}
}