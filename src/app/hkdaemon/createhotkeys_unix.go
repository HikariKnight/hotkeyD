// +build !windows

package hkdaemon

import (
	"fmt"
)

// CreateHotKeys creates hotkeys based on the configs from hotkeys.ini
func CreateHotKeys() {
	fmt.Println("Unsupported OS, assuming you are just testing. So here is a repsonse from CreateHotKeys()!")
}
