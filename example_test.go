package uuid_test

import (
	"fmt"

	"github.com/ybubnov/go-uuid"
)

func ExampleKernel() {
	src := uuid.Kernel{MaxIdle: 128, MaxProcs: 16}
	defer src.Stop()

	// Generate the next UUID v4.
	u1, err := src.Next()
	if err != nil {
		fmt.Printf("failed to generate uuid: %s\n", err)
	}

	fmt.Printf("uuid is %cth version\n", u1[14])
	// Output: uuid is 4th version
}
