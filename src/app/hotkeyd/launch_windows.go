// +build windows
package hotkeyd

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// Launch a defined program
func Launch(hidden bool, cmd string, args ...string) {
	// Configure the command we will run
	cmnd := exec.Command(cmd, args...)
	if hidden {
		// Hide the putty window
		cmnd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	// Tell in prompt what we are doing
	fmt.Printf("Executing %s with these arguments:\n%s\n", cmd, strings.Join(args, " "))

	// Run the command
	cmnd.Start()
}
