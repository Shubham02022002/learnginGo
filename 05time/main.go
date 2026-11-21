package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In this module we'll be looking into time")
	currentTime := time.Now()
	fmt.Println(currentTime)

	fmt.Println(currentTime.Format("01-02-2006 "))
	fmt.Println(currentTime.Format("01-02-2006 Monday "))
	fmt.Println(currentTime.Format("01-02-2006 Monday 15:04:04"))

	newDate := time.Date(2047, time.August, 15, 00, 00, 0, 0, time.UTC)
	fmt.Println("Celebrating 100 years of freedom", newDate.Format("01-02-2006 Monday"))
}
