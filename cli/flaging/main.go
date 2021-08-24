package main

import (
	"flag"
	"fmt"
)

// go run cli/flaging/main.go -count 11
// go run cli/flaging/main.go -count=11
func main() {
	count := flag.Int("count", 10, "use for looping")
	flag.Parse()

	fmt.Println("The count is", *count)
}
