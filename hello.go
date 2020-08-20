package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/karnadattasai/GitTest/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	// Removed a line here
	fmt.Println("NewTest branch")
	// This is an update on previous commit based on comments to incorporate changes
	// updating remote repo basing on comments 10:01 am Fri 21 Aug 2020
}
