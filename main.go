package main

import (
	"fmt"
	"os"
	registry "github.com/nicholasdille/registry/registry"
)

func main() {
	fmt.Fprintln(os.Stdout, "Start")

	registry := registry.NewRegistry()
	registry.Connect()

	fmt.Fprintln(os.Stdout, "End")
}