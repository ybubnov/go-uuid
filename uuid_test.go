package uuid

import (
	"io/ioutil"
	"testing"
)

// swapDefaultSrc exchanges the default source variable with a given one.
func swapDefaultSrc(new Source) (old Source) {
	old = defaultSrc
	defaultSrc = new
	return old
}

func TestKernelNext(t *testing.T) {
	var src Kernel
	defer src.Stop()

	uuid1, err := src.Next()
	if err != nil {
		t.Fatalf("kernel source failed to generate UUID: %s", err)
	}

	uuid2, err := src.Next()
	if err != nil {
		t.Fatalf("kernel source failed to generate UUID: %s", err)
	}

	if uuid1 == uuid2 {
		t.Fatalf("kernel source produced the same UUIDs")
	}
}

func TestKernelStop(t *testing.T) {
	src := Kernel{MaxProcs: 128}

	src.Next()
	src.Stop()

	// Ensure all routines have been terminated.
	for i := 0; i < src.MaxProcs; i++ {
		<-src.done
	}
}

func TestKernelNextErrStopped(t *testing.T) {
	var src Kernel
	src.Stop()

	_, err := src.Next()
	if err != ErrStopped {
		t.Fatalf("next not allowed on stopped source: %s", err)
	}
}

func TestKernelNextErrorOpen(t *testing.T) {
	src := Kernel{path: "/this/path/should/not/exist"}
	defer src.Stop()

	id, err := src.Next()

	if id != "" {
		t.Fatalf("next should return an empty string UUID: %s", id)
	}

	if err == nil {
		t.Fatalf("next should complain about non-existing file")
	}
}

func TestKernelNextErrIncomplete(t *testing.T) {
	// Create a fake UUID generator with a chunked content in
	// order to simulate incorrect content returned from kernel.
	file, _ := ioutil.TempFile("/tmp", "")
	file.WriteString("307d54bc")
	file.Close()

	src := Kernel{path: file.Name()}
	defer src.Stop()

	_, err := src.Next()
	if err != ErrIncomplete {
		t.Fatalf("next should fail with chunked data error: %s", err)
	}
}

func TestNopSource(t *testing.T) {
	text := "8139da46-d1f9-4993-bea3-d3dc7fc3a7c7"
	src := NopSource(text)
	defer src.Stop()

	id, err := src.Next()
	if err != nil {
		t.Fatalf("next of no-op source should not return errors: %s", err)
	}

	if id != text {
		t.Fatalf("next returned invalid UUID: %s", id)
	}
}

func TestNosupSource(t *testing.T) {
	var src nosupSource
	defer src.Stop()

	_, err := src.Next()
	if err != ErrNotSupported {
		t.Fatalf("next should return not supported error: %s", err)
	}
}

func TestNew(t *testing.T) {
	text := "69ce0151-5df2-4f8d-99e4-8b66c072ffe8"
	src := NopSource(text)

	old := swapDefaultSrc(src)
	defer swapDefaultSrc(old)

	id := New()
	if id != text {
		t.Fatalf("default source returned wrong UUID: %s", id)
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}
