package main

import (
	"basicCLI/processRequest"
	"fmt"
)

func main() {
	dataFiles := make([]string, 0)
	var maxFiles int
	fmt.Print("Enter the maximum number of notes: ")
	fmt.Scan(&maxFiles)
	processRequest.ProcessRequest(dataFiles, maxFiles)
}
