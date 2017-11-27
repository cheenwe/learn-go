package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func checkType(name string) {
	switch name {
	case "Album":

		fmt.Println("Album")
	case "Song":
		fmt.Println("Song")
	}
}

// RandInt 生成min~max范围内随机整数
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
func main() {

	coreNum := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("NumCPU: ", coreNum)

	fmt.Println("RandInt(1, 10): ", RandInt(1, 10))

	checkType("Album")
}
