package main

import (
	"flag"
	"fmt"
)

func main() {
	flagName := flag.String("name", "xcccc", "name flag for test")
	flagCount := flag.Int("count", 12, "count flag for test")
	flag.Parse()

	fmt.Println("flag name:", *flagName, " --- ", flagName)
	fmt.Println("flag count:", *flagCount, " --- ", flagCount)
	fmt.Println("flag args:", flag.Args())
}
