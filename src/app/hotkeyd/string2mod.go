package hotkeyd

import (
	"strings"

	"github.com/MakeNowJust/hotkey"
)

// String2Mod converts a modkey string to the correct hotkey.Modifier value
func String2Mod(s string) hotkey.Modifier {
	// Make a variable to contain our modifier key value
	var intkey = hotkey.None

	// If a combination of modkeys are defined (the string will then contain one or more +)
	switch strings.Contains(s, "+") {
	case true:
		// Split the string by +
		keys := strings.Split(s, "+")

		// For every key in the combination
		for v := range keys {
			// Convert the key to a hotkey.modifier
			if strings.EqualFold(keys[v], "ALT") {
				intkey += hotkey.Alt
			}
			if strings.EqualFold(keys[v], "CTRL") {
				intkey += hotkey.Ctrl
			}
			if strings.EqualFold(keys[v], "SHIFT") {
				intkey += hotkey.Shift
			}
			if strings.EqualFold(keys[v], "WIN") {
				intkey += hotkey.Win
			}
		}
	case false:
		// Convert our single modkey to a hotkey.modifier
		if strings.EqualFold(s, "ALT") {
			intkey += hotkey.Alt
		}
		if strings.EqualFold(s, "CTRL") {
			intkey += hotkey.Ctrl
		}
		if strings.EqualFold(s, "SHIFT") {
			intkey += hotkey.Shift
		}
		if strings.EqualFold(s, "WIN") {
			intkey += hotkey.Win
		}
	}

	// Return the hotkey.modifier value
	return intkey
}
