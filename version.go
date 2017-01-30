package main

import (
	"fmt"
)

var (
	Version  string
	Revision string
)

func printVersion() {
	fmt.Println("mr version " + Version + ", build " + Revision)
}
