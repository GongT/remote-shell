package timeout_controller

import (
	"time"
	"errors"
)

type TimeoutController struct {
	mapper map[uint32]*CancellableTimer
	id     uint32
}

func NewTimeoutController() (*TimeoutController) {
	return &TimeoutController{
		mapper: make(map[uint32]*CancellableTimer),
		id:     0,
	}
}

func (tc *TimeoutController) Wait(d time.Duration) (uint32, <-chan error) {
	tc.id++
	id := tc.id

	tmr := NewCancellableTimer()
	tc.mapper[id] = tmr

	timeout := make(chan error)
	go func() {
		if <-tmr.After(d) {
			timeout <- errors.New("timeout.")
		}
		close(timeout)
		delete(tc.mapper, id)
	}()

	return id, timeout
}
func (tc *TimeoutController) Cancel(id uint32) bool {
	tmr, ok := tc.mapper[id]

	if ok {
		tmr.Cancel()
	}

	return ok
}
