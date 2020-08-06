// +build !windows

package hotkeyd

import (
	"fmt"
)

// Launch a program on unix/unix-like platform (not supported)
func Launch(hidden bool, cmd string, args ...string) {
	fmt.Println("Unsupported OS, assuming you are just testing. So here is a repsonse!")
	fmt.Printf("We received the command: %s\nWe were told to pass these arguments: %s\nShould the window be hidden?: %t\n", cmd, args, hidden)
}
