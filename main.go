package main

import (
	"fmt"
	"log"
	"flag"
	registry "github.com/nicholasdille/go-containers/registry"
)

var hostname string
var port int
var insecure bool
var username string
var password string

func main() {
	flag.StringVar(&hostname, "hostname", "localhost", "Hostname of the registry")
	flag.IntVar(   &port,     "port",     443,         "Port of the registry")
	flag.BoolVar(  &insecure, "insecure", false,       "Connection scheme")
	flag.StringVar(&username, "username", "",          "Username for registry")
	flag.StringVar(&password, "password", "",          "Password for registry")
	flag.Parse()

	registry := registry.NewRegistry(hostname, port, insecure, username, password)
	err := registry.Check()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(registry.String())
}