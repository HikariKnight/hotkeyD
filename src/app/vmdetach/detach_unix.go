// +build !windows

package vmdetach

import (
	"fmt"
)

// Detach from the VM by running a custom putty session using the putty executable in the app directory
func Detach() {
	fmt.Println("Unsupported OS, assuming you are just testing. So here is a repsonse!")
}
