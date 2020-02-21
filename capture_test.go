package soos_test

import (
	"fmt"
	"os"

	"github.com/lebaptiste/soos"
)

func ExampleCapture() {
	fmt.Println("Hello Gophers!")
	dump := soos.Capture(os.Stdout)
	fmt.Println("How do you do?")
	stdout, err := dump()
	fmt.Println("Goodbye!")

	if err != nil {
		panic(err)
	}

	got := string(stdout)
	want := "How do you do?\n"
	if got != want {
		panic(fmt.Errorf("got %v, want %v", got, want))
	}

	// Output:
	// Hello Gophers!
	// How do you do?
	// Goodbye!
}
