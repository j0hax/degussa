package main

import (
	"fmt"

	"github.com/j0hax/degussa"
)

func main() {
	for _, item := range degussa.All() {
		fmt.Printf("%s\n", item)
	}
}
