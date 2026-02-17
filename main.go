package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// This function starts a background worker and returns a "pipe" (channel)
func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	// 'go' keyword starts the Goroutine (the background worker)
	go func() {
		// Close the file and the channel when this worker finishes
		defer f.Close()
		defer close(ch)

		buffer := make([]byte, 8)
		var currentLine string

		for {
			n, err := f.Read(buffer)
			if n > 0 {
				chunk := string(buffer[:n])
				parts := strings.Split(chunk, "\n")

				for i := 0; i < len(parts)-1; i++ {
					// Instead of printing, we SEND to the channel
					ch <- currentLine + parts[i]
					currentLine = ""
				}
				currentLine += parts[len(parts)-1]
			}

			if err == io.EOF {
				break
			}
		}

		// Send any leftover text
		if currentLine != "" {
			ch <- currentLine
		}
	}()

	return ch
}

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Get the channel (the conveyor belt)
	linesCh := getLinesChannel(file)

	// Range over the channel. This loop stays open as long as the
	// channel is open. As soon as a line is sent, this loop runs!
	for line := range linesCh {
		fmt.Printf("read: %s\n", line)
	}
}
