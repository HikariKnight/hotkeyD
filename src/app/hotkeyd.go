package main

// Import the libraries we need
import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strings"

	"./hotkeyd"
	"github.com/MakeNowJust/hotkey"
	"github.com/kardianos/osext"
	"tawesoft.co.uk/go/dialog"

	"gopkg.in/ini.v1"
)

// Make a main function
func main() {
	// Make a quit method
	quit := make(chan bool)

	if runtime.GOOS != "windows" {
		fmt.Printf("OS not supported!\n")
	}

	// Create hotkeys
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

		// Tell that we cannot find the hotkeys.ini file
		dialog.Alert("No hotkeys.ini found, please make one and put some configs in it. The undefined Modkeys and Hotkey entries at the very top are reserved for the Toggle hotkey!\n\nModkeys=ctrl+alt\nHotkey=q\n[name of hotkey]\nModkeys=ctrl+alt\nHotkey=b\nLaunch=C:\\windows\\system32\\notepad.exe\nArgs=--some random -a -r -g -s=here\nHide=false")

		// Exit with code 1
		os.Exit(1)
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
				// Make an empty intkey
				var intKey = hotkey.None

				// Convert the modkeys to a hotkey.Modifier
				intKey = hotkeyd.String2Mod(modKeys)

				// Get the hotkey from settings and convert to uint32
				var intHotKey uint32 = hotkeyd.HotkeySwitch(hotKey)

				// Make our hotkey
				hkey.Register(intKey, intHotKey, func() {
					fmt.Println("Hotkey Pressed!")

					// Get the absolute path to the executable
					filename, _ := osext.Executable()
					fmt.Printf(filename)

					// Check if --pause is passed to the arguments
					switch strings.Contains(strings.Join(os.Args[1:], " "), "--pause") {
					case false:
						// Start paused instance
						hotkeyd.Launch(false, filename, "--pause")
					case true:
						// Start normal instance
						hotkeyd.Launch(false, filename)
					}
					// Exit this hotkeyd instance
					os.Exit(0)
				})

			default:
				// For any other hotkey definitions
				// Check if --pause is passed as an argument
				switch strings.Contains(strings.Join(os.Args[1:], " "), "--pause") {
				case false:
					// If we are not in paused mode
					// Check if the hotkey options are not undefined
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
}
