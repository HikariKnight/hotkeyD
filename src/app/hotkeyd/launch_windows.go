//go:build windows
// +build windows

package hotkeyd

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// Launch a defined program
func Launch(hidden bool, cmd string, workdir string, args ...string) {
	// Configure the command we will run
	cmnd := exec.Command(cmd, args...)
	if hidden {
		// Hide the putty window
		cmnd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}

	// If we have a workdir defined
	if workdir != "" {
		cmnd.Dir = workdir
	}

	// Tell in prompt what we are doing
	fmt.Printf("Executing %s in %s with these arguments:\n%s\n", cmd, workdir, strings.Join(args, " "))

	// Run the command
	cmnd.Start()
}
