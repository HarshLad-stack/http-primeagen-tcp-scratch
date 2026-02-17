package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// getLinesChannel stays EXACTLY the same!
// This is the power of using the 'io.ReadCloser' interface.
func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
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
					ch <- currentLine + parts[i]
					currentLine = ""
				}
				currentLine += parts[len(parts)-1]
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
		}
		if currentLine != "" {
			ch <- currentLine
		}
	}()
	return ch
}

func main() {
	// 1. Start the Listener on port 42069
	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 42069...")

	// 2. Infinite loop to keep the server running
	for {
		// 3. Wait for a connection (Program pauses here)
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Connection accepted")

		// 4. Use your existing channel logic to read from the network!
		linesCh := getLinesChannel(conn)

		for line := range linesCh {
			// Print exactly what comes through the pipe
			fmt.Println(line)
		}

		fmt.Println("Connection closed")
	}
}
