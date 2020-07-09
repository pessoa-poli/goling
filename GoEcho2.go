// Echo1 prints its command-line arguments.
package main

import (
"fmt"
"os"
)

func main() {
	s, sep := "", "\r\n"
	for i, arg := range os.Args[1:] {
	s = i + " " + arg + sep
	fmt.Println(s)
	}
	
	}