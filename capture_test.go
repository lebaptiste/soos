package soos_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/lebaptiste/soos"
)

func ExampleCapture() {
	fmt.Println("Hello Gophers!")

	dump := soos.Capture(os.Stdout)
	fmt.Println("How do you do?")
	stdout, _ := dump()

	fmt.Println("Goodbye!")
	fmt.Printf("captured: %v", string(stdout))

	// Output:
	// Hello Gophers!
	// How do you do?
	// Goodbye!
	// captured: How do you do?
}

func TestCaptureError(t *testing.T) {
	// read only file to assess write error
	f2, err := os.Open("capture.go")
	if err != nil {
		panic(err)
	}
	dump := soos.Capture(f2)
	f2.Write([]byte("test"))
	output, err := dump()
	if err == nil {
		t.Error("got <nil>, want an error")
	}
	if output != nil {
		t.Errorf("got %v, want <nil> slice", output)
	}
}
