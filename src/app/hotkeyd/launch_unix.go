// +build !windows

package hotkeyd

import (
	"fmt"
)

// Launch a program on unix/unix-like platform (not supported)
func Launch(hidden bool, cmd string, args ...string) {
	fmt.Println("Unsupported OS, assuming you are just testing. So here is a repsonse!")
}
