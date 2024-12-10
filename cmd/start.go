package cmd

import (
	"acqua-remind-cli/internal/timer"
	"fmt"
)

func Start() {
	fmt.Println("The program is starting. Only for 10 sec!")
	timer.StartTimer()
	fmt.Println("Time is up!The program is closed!")
}
