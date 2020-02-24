package soos

import (
	"bytes"
	"io"
	"os"
	"sync"
)

// Capture reads from the *os.File until the dump func is called.
// The dump func returns the bytes captured since.
// The function is not thread safe.
func Capture(f *os.File) (dump func() ([]byte, error)) {
	var err error
	source := *f
	var buf bytes.Buffer
	w := io.MultiWriter(&source, &buf)

	r, out, err := os.Pipe()
	if err != nil {
		return func() ([]byte, error) { return nil, err }
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_, err = io.Copy(w, r)
		wg.Done()
	}()

	*f = *out

	return func() ([]byte, error) {
		out.Close()
		*f = source
		wg.Wait()
		if err != nil {
			return nil, err
		}
		txt := buf.Bytes()
		return txt, nil
	}
}
