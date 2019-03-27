package timeout_controller

import "time"

var tc *TimeoutController

func init() {
	tc = NewTimeoutController()
}

func Wait(d time.Duration) (uint32, <-chan error) {
	return tc.Wait(d)
}
func Cancel(id uint32) bool {
	return tc.Cancel(id)
}
