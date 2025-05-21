package cmd

import (
	"acqua-remind-cli/internal/timer"
	"fmt"
	"time"
)

const timeDuration = 3 * time.Second

func Start() {
	fmt.Printf("The program is starting! Only for %v!\n", timeDuration)
	timer.MainTimer(timeDuration)
	fmt.Println("\nTime is up!The program is closed!")
}
