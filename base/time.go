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
	//	看上去是不是有点懵？ 2006-01-02 15:04:05这个每组数字都是有独特的含义的，就是相当于拿一段数字来代替了我们其他语言常用的YY:mm:ss HH-MM-SS 这样。但是其实习惯之后发现这个设计还蛮好记的，直接记12345就好了。3那个位置的数这里我使用的15，也就是用24小时格式来显示，如果直接写03则是12小时am pm格式。

	baseTimeFormat = "2006-01-02 15:04:05"
)

func main() {
	t := currentTime()
	currentAt := t.Format(baseTimeFormat)
	fmt.Println("Time is: ", currentAt)

	fmt.Println(time.Now().Unix())

	timestamp := time.Now().Unix()
	fmt.Println(time.Unix(timestamp, 0))

	fmt.Println(time.Date(2017, 11, 27, 20, 20, 20, 20, time.Local))
}
