package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/karnadattasai/GitTest/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println("Its my new branch")
}
