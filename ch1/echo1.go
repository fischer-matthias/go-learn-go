// prints the given arguments in console
package main

import (
	"fmt"
	"os"
)

func main() {
	var str, seperator string

	for i:= 1; i < len(os.Args); i++ {
		str += seperator + os.Args[i]

		if i == 1 {
			seperator = " "
		}
	}

	fmt.Println(str)
}