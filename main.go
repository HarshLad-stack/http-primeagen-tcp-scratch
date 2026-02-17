package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 8)
	var currentLine string ///<---- This is the waiting room

	for {
		n, err := file.Read(buffer)
		if n > 0 {
			chunk := string(buffer[:n])
			parts := strings.Split(chunk, "\n")

			for i := 0; i < len(parts)-1; i++ {
				fmt.Printf("read: %s\n", currentLine+parts[i])
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]

		}
		if err == io.EOF {
			break
		}

		if currentLine != "" {
			fmt.Printf("read: %s\n", currentLine)
		}
	}

}
