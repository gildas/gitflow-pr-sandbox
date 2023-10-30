package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello, World! version: %s\n", VERSION)
	fmt.Printf("You are running on %s %s architecture\n", runtime.GOOS, runtime.GOARCH)
}
