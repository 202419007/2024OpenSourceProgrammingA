package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	//var month int = int(now.Month())
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다\n", now.Year(), int(now.Month()), now.Day())
}
