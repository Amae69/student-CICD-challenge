package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func GetMessage() string {
	return "Hello, DevelopersFoundry fellows!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := color.New(color.FgCyan, color.Bold)
	msg := GetMessage()
	c.Println(msg)
	fmt.Fprintf(w, "%s\n", msg)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
