package main

import (
	"fmt"

	"github.com/j0hax/degussa"
)

func main() {
	for i, item := range degussa.All() {
		fmt.Printf("[%d] %s\n", i+1, item)
	}
}
