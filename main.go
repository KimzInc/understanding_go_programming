package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// -----------------------------------------------------------
// APPROACH 1: Allocate a new slice for every read
// -----------------------------------------------------------

func readWithNewAllocation(fileName string) error {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	// Ensure the file is closed when the function returns
	defer file.Close()

	fmt.Println("\n========== APPROACH 1 ==========")
	fmt.Println("Allocating a new slice for every read")

	for {
		// Allocate a new slice for every iteration
		dataChunk := make([]byte, 100) // 100 bytes buffer
		count, err := file.Read(dataChunk)

		// Process only the bytes actually read
		if count > 0 {
			// process function called here
			process(dataChunk[:count])

		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil // End of file reached, exit the loop
			}
			return err // Return any other error

		}
	}

}

// -----------------------------------------------------------
// PROCESS FUNCTION
// -----------------------------------------------------------

func process(data []byte) {
	// Placeholder for processing logic
	fmt.Printf("Processed %d bytes: %q\n", len(data), data)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	fileName := os.Args[1]

	// First approach
	err := readWithNewAllocation(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
