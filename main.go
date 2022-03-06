package main

import (
	"flag"
	"fmt"
)

var (
	name = flag.String("name", "world", "name")
)

func main() {
	flag.Parse()
	fmt.Printf("Hello %v\n", *name)
}
