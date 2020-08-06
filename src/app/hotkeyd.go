package main

// Import the libraries we need
import (
	"fmt"
	"runtime"
	"strings"

	"./hotkeyd"

	"github.com/MakeNowJust/hotkey"
	"gopkg.in/ini.v1"
)

// Make a main function
func main() {
	// Make a quit method
	quit := make(chan bool)

	if runtime.GOOS != "windows" {
		fmt.Printf("OS not supported! Exiting...\n")
		quit <- true
	}

	// Make a new hotkey definition
	hkey := hotkey.New()

	// Read our configs
	cfg, cfgerr := ini.Load("hotkeys.ini")
	// If we cannot read the file we use default values
	if cfgerr != nil {
		fmt.Printf("Fail to read file: %v\n", cfgerr)
		fmt.Println("Using defaults instead")

		if runtime.GOOS == "windows" {
			// Make a hotkey to run putty with our profile
			hkey.Register(hotkey.Ctrl+hotkey.Alt, hotkey.SPACE, func() {
				hotkeyd.Launch(false, "cmd.exe")
			})
		}
	} else {
		// Get the modkeys for our vm-detach hotkey
		modKeys := cfg.Section("").Key("Modkeys").String()

		// get the hotkey for our vm-detach hotkey combo
		hotKey := strings.ToUpper(cfg.Section("").Key("Hotkey").String())

		// Make a variable to contain our uint32 value for the key combination
		var intKey = hotkey.None

		// Convert the modkeys to a hotkey.Modifier
		intKey = hotkeyd.String2Mod(modKeys)

		// Get the hotkey from settings and convert to uint32
		var intHotKey uint32 = hotkeyd.HotkeySwitch(hotKey)

		if runtime.GOOS == "windows" {
			// Make our hotkey
			hkey.Register(intKey, intHotKey, func() {
				// Execute putty and detach our mkb from VM
				hotkeyd.Launch(false, "cmd.exe")
			})
		}
	}

	fmt.Println("Started hotkey loop")
	<-quit
}
