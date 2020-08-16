// +build !windows

package hkdaemon

import (
	"fmt"
)

// CreateHotkeys creates hotkeys based on the configs from hotkeys.ini
func CreateHotkeys() {
	fmt.Println("Unsupported OS, assuming you are just testing. So here is a repsonse from CreateHotKeys()!")
}
