package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// -----------------------------------------------------------
// Function MakePerson creates a new Person instance and returns Person by value

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// -----------------------------------------------------------
// Function MakePersonPtr creates a new Person instance and returns a pointer to Person

func MakePersonPtr(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

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

//-----------------------------------------------------------
// APPROACH 2: Reuse one slice as a buffer
// -----------------------------------------------------------

func readWithReusedBuffer(fileName string) error {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	// Ensure the file is closed when the function returns
	defer file.Close()

	fmt.Println("\n========== APPROACH 2 ==========")
	fmt.Println("Reusing one slice as a buffer")

	// Allocate one slice before entering the loop.
	data := make([]byte, 100) // 100 bytes buffer

	for {

		// resuse the same underlying array for every read
		count, err := file.Read(data)

		// Process only the bytes actually read
		if count > 0 {
			// process function called here
			process(data[:count])
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

	// Second approach
	err = readWithReusedBuffer(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n========== Demonstrating MakePerson and MakePersonPtr ==========")

	// Demonstrate the use of MakePerson and MakePersonPtr
	person1 := MakePerson("John", "Doe", 30)
	fmt.Printf("\nPerson by value: %+v\n", person1)

	person2 := MakePersonPtr("Jane", "Smith", 25)
	fmt.Printf("Person by pointer: %+v\n", *person2)

}
