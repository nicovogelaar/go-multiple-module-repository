package main

import (
	"log"

	"github.com/nicovogelaar/go-multiple-module-repository/blue/red"
)

func main() {
	Hello()
	red.Hello()
}

func Hello() {
	log.Print("Hello, blue!")
}
