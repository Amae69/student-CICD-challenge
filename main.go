package main

import (
	"fmt"

	"github.com/fatih/color"
)

func GetMessage() string {
	return "Hello, CI/CD World!"
}

func main() {
	c := color.New(color.FgCyan, color.Bold)
	c.Println(GetMessage())
	fmt.Println("Application started successfully.")
}
