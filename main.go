package main

import (
	"fmt"

	"github.com/fatih/color"
)

func GetMessage() string {
	return "Hello, DevelopersFoundry fellows!"
}

func main() {
	c := color.New(color.FgCyan, color.Bold)
	c.Println(GetMessage())
	fmt.Println("My fellow Comrade. My Application started successfully.")
}
