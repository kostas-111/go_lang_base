package main

import (
	"fmt"
	"job4j.ru/go-lang-base/internal/tracker"
)

func main() {
	fmt.Println("Hello World")
	ui := tracker.UI{
		In:      tracker.ConsoleInput{},
		Out:     tracker.ConsoleOutput{},
		Tracker: tracker.NewTracker(),
	}
	ui.Run()
}
