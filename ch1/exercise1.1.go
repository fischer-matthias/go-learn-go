// exercise1.1 - prints the given arguments in console
package main

import (
	"fmt"
	"os"
)

func main() {
	var str, seperator string

	for i:= 0; i < len(os.Args); i++ {
		str += seperator + os.Args[i]
		seperator = " "
	}

	fmt.Println(str)
}