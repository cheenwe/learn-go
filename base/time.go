package main

import (
	"fmt"
	"time"
)

var (
	currentTime = time.Now
)

const (
	baseTimeFormat = "2006-01-02 15:04:05"
)

func main() {
	t := currentTime()
	currentAt := t.Format(baseTimeFormat)
	fmt.Println("Time is: ", currentAt) 
}
