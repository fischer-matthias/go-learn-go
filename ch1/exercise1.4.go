// exercise 1.4 - print name of file in which the line is included
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args) > 1 {
		inFiles := make(map[string]map[string]bool)

		for _, filename := range os.Args[1:] {
			data, err := ioutil.ReadFile(filename)
	
			if err != nil {
				fmt.Fprintf(os.Stderr, "dupl3: %v\n", err)
				continue
			}
	
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
				if inFiles[line] == nil {
					inFiles[line] = make(map[string]bool)
				}
				inFiles[line][filename] = true
			}
		}

		for line, n := range counts {
			for fileName := range inFiles[line] {
				fmt.Printf("-%s:\n", fileName)
			}
	
			fmt.Printf("%dmal \t%s\n\n", n, line)
		}
	} else {
		fmt.Println("Please add a file as argument.")
	}
}