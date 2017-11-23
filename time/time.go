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
	//	return filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, timestamp, ext))
	fmt.Println("Time is: ", currentAt) 
}
