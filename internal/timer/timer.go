package timer

import (
	"fmt"
	"os/exec"
	"time"
)

func MainTimer(duration time.Duration) {
	Countdown(duration)
}

func Countdown(duration time.Duration) {
	ticker := time.NewTicker(1 * time.Second)
	stopTime := time.After(duration)
	startTime := time.Now()
	remaining := duration

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			remaining = duration - time.Since(startTime).Round(time.Second)

			minutes := int(remaining.Minutes())
			seconds := int(remaining.Seconds()) % 60

			fmt.Printf("\rTimer: %02d:%02d", minutes, seconds)
		case <-stopTime:
			cmd := exec.Command("afplay", "/System/Library/Sounds/Glass.aiff")
			cmd.Run()
			return
		}
	}
}
