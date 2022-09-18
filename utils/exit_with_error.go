package utils

import (
	"fmt"
	"os"
)

func ExitWithError(err string) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}
