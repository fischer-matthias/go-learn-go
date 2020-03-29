// exercise1.3 - show difference between echo1 and echo3
package main

import (
	"strings"
	"os"
	"fmt"
	"time"
)

func main() {
	startEcho1 := time.Now()
	echo1()
	fmt.Printf("echo1()\t%b elapsed\n", time.Since(startEcho1).Milliseconds())
	
	startEcho3 := time.Now()
	echo3()
	fmt.Printf("echo3()\t%b elapsed\n", time.Since(startEcho3).Milliseconds())
}

func echo1() {
	var str, seperator string

	for i:= 1; i < len(os.Args); i++ {
		str += seperator + os.Args[i]

		if i == 1 {
			seperator = " "
		}
	}

	fmt.Println(str)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}