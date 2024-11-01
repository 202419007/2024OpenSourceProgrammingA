package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("점수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	var isPrime bool = true

	// 논리적 오류 해결
	if n < 2 {
		isPrime = false
	} else {
		for j := 2; j < n; j++ {
			if n%j == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d", j) // 반복문 횟수 확인용 코드
		}
	}

	if isPrime {
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}
