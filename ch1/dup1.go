// dup1 - prints duplicated sentences with duplication-count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan() {
		if (input.Text() == "q") {
			break
		} else {
			counts[input.Text()]++
		}
	}

	// NOTE: ignoring potential errors
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}