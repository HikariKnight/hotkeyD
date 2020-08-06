package main

// Import the libraries we need
import (
	"encoding/csv"
	"fmt"
	"runtime"
	"strings"

	"./hotkeyd"
	"github.com/MakeNowJust/hotkey"
	"tawesoft.co.uk/go/dialog"

	"gopkg.in/ini.v1"
)

var hotkeysEnabled bool = false

// Make a main function
func main() {
	// Make a quit method
	quit := make(chan bool)

	if runtime.GOOS != "windows" {
		fmt.Printf("OS not supported!\n")
	}

	createHotkeys()

	fmt.Println("Started hotkey loop")
	<-quit
}

func createHotkeys() {
	// Make a new hotkey definition
	hkey := hotkey.New()

	// Read our configs
	cfg, cfgerr := ini.Load("hotkeys.ini")
	// If we cannot read the file we use default values
	if cfgerr != nil {
		fmt.Printf("Fail to read file: %v\n", cfgerr)
		fmt.Println("Using defaults instead")

		dialog.Alert("No hotkeys.ini found, please make one and put some configs in it.\n[name of hotkey]\nModkeys=ctrl+alt\nHotkey=b\nLaunch=C:\\windows\\system32\\notepad.exe\nArgs=--some random -a -r -g -s=here\nHide=false")
	} else {
		for _, hotkeyName := range cfg.SectionStrings() {
			// When we have a defined hotkey, get the hotkey config
			hotkeyIni := cfg.Section(hotkeyName)

			// Get all the info for the hotkey
			var hidewindow bool = hotkeyIni.Key("Hide").MustBool(false)
			var launchCMD string = hotkeyIni.Key("Launch").MustString("")
			var cmdArgsStr string = hotkeyIni.Key("Args").MustString("")
			var modKeys string = hotkeyIni.Key("Modkeys").MustString("")
			var hotKey string = hotkeyIni.Key("Hotkey").MustString("")

			switch hotkeyName {
			case "DEFAULT":
				// If the section is named DEFAULT (or empty which returns DEFAULT)
				fmt.Println("TODO: hotkeytoggle()")

				// Make a variable to contain our uint32 value for the key combination
				/*var intKey = hotkey.None

				// Convert the modkeys to a hotkey.Modifier
				intKey = hotkeyd.String2Mod(modKeys)

				// Get the hotkey from settings and convert to uint32
				var intHotKey uint32 = hotkeyd.HotkeySwitch(hotKey)

				// Make our hotkey
				hkey.Register(intKey, intHotKey, func() {
					fmt.Println("Hotkey Pressed!")
					// Execute the defined program
					hkey.Stop()
				})*/

			default:
				// If the hotkeys options are not undefined
				if launchCMD != "" && modKeys != "" && hotKey != "" {
					// Initialize an empty string slice/array
					args := []string{}
					// If the argument string is not empty
					if cmdArgsStr != "" {
						// Make a new csv parser to parse the string
						csv := csv.NewReader(strings.NewReader(cmdArgsStr))
						// Set the "comma" to whitespace
						csv.Comma = ' '

						// Initialize an error variable
						var csvErr error

						// Read the argument string
						args, csvErr = csv.Read()

						// If we get an error
						if csvErr != nil {
							// Print the error and return
							fmt.Println(csvErr)
							return
						}
					}

					// Make a variable to contain our uint32 value for the key combination
					var intKey = hotkey.None

					// Convert the modkeys to a hotkey.Modifier
					intKey = hotkeyd.String2Mod(modKeys)

					// Get the hotkey from settings and convert to uint32
					var intHotKey uint32 = hotkeyd.HotkeySwitch(hotKey)

					// Make our hotkey
					hkey.Register(intKey, intHotKey, func() {
						fmt.Println("Hotkey Pressed!")
						// Execute the defined program
						hotkeyd.Launch(hidewindow, launchCMD, args...)
					})
				}
			}
		}
	}
}
