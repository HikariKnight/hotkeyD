// +build windows
package vmdetach

import (
	"fmt"
	"os/exec"
	"syscall"
)

// Detach from the VM by running a custom putty session using the putty executable in the app directory
func Detach() {
	// Configure the command we will run
	cmnd := exec.Command("putty.exe", "-load", "vm-detach")
	// Hide the putty window
	cmnd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	// Run the command
	cmnd.Start()
	// Tell in prompt what we are doing
	fmt.Println("Putty running to detach usb mouse and keyboard.")
}
