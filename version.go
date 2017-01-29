package main

import (
	"fmt"
)

var (
	Version  string
	Revision string
)

func printVersion() {
	fmt.Println("mg version " + Version + ", build " + Revision)
}
