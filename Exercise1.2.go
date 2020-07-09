// Echo1 prints its command-line arguments.
package main

import (
"fmt"
"os"
"strconv"
)

func main() {
	s, sep:= "", "\r\n"
	for i, arg := range os.Args[1:] {
	s = strconv.Itoa(i) + " " + arg + sep
	fmt.Println(s)
	}
	
	}