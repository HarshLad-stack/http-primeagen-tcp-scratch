package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 8)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			fmt.Printf("read: %s\n", string(buffer[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			break
		}
	}

}
