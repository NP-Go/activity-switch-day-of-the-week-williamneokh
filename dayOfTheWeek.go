package main

import (
	"fmt"
	"time"
)

func main() {
	//Insert code here
	var dayNum int
	today := time.Now().Weekday()
	fmt.Println(today)

	switch time.Now().Weekday() {
	case time.Monday:
		dayNum = 1
	case time.Tuesday:
		dayNum = 2
	case time.Wednesday:
		dayNum = 3
	case time.Thursday:
		dayNum = 4
	case time.Friday:
		dayNum = 5
	case time.Saturday:
		dayNum = 6
	case time.Sunday:
		dayNum = 7

	}
	if dayNum%2 == 0 {
		fmt.Println("It's an even number")
	} else {
		fmt.Println("It's an odd number")
	}

}
