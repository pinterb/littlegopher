package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("")
	now := time.Now()
	fmt.Println(now.Format(time.UnixDate))
	fmt.Println("")
	fmt.Println("When is Friday?")
	today := time.Now().Weekday()

	switch time.Thursday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
	time.Sleep(3000 * time.Millisecond)
}
