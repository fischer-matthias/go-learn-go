// exercise1.1 - prints the given arguments in console
package main

import (
	"fmt"
	"os"
)

func main() {
	for i:= 1; i < len(os.Args); i++ {
		fmt.Printf("%d\t%s\n", i, os.Args[i])
	}
}