package timer

import (
	"testing"
	"time"
)

func TestStartTimer(t *testing.T) {
	// timer := &ConfigurableTimer{3 * time.Second}
	//
	// want := 3 * time.Second
	//
	// Countdown(timer)
	//
	// if <-timer.Run() != want {
	// 	t.Errorf("got %v")
	// }
}

type SpyTimer struct {
	duration time.Duration
	test     time.Time
}

func (st *SpyTimer) Run() <-chan time.Time {
	st.test = <-time.After(st.duration)

	return make(<-chan time.Time)
}
