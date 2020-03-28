// echo2 - prints the given arguments in console
package main

import (
	"fmt"
	"os"
)

func main() {
	str, seperator := "", ""

	for _, arg := range os.Args[1:] {
		str += seperator + arg
		seperator = " "
	}

	fmt.Println(str)
}