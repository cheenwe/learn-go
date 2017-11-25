package main

import (
	"fmt"
	"time"
)

var (
	currentTime = time.Now
)

const (
	//	为什么是这个时间一直搞不懂, 换成其他时间会出错
	baseTimeFormat = "2006-01-02 15:04:05"
)

func main() {
	t := currentTime()
	currentAt := t.Format(baseTimeFormat)
	fmt.Println("Time is: ", currentAt)
}
