package main

import (
	"PostProcessor"
	"flag"
	"fmt"
	"os"
)

func main() {
	formula := flag.String("formula", "", "Example: =SUM[~0,RAND[0,500]]")
	flag.Parse()

	if *formula == "" {
		os.Exit(0)
	}

	value, err := PostProcessor.Calculate(*formula, 5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}
