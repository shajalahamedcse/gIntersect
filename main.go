package main

import (
	"flag"
	"fmt"
	"os"
)

var howToUse = `usage : intersect [File A] [File B]
Prints out the lines that are both in [File A] and [File B]`

//Exit with error message
func exit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {

	flag.Parse()

	if len(flag.Args()) < 2 {
		exit(howToUse)
	}
}
