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
	if strings.Contains(s, "+") {
		// Split the string by +
		keys := strings.Split(s, "+")

		// For every key in the combination
		for v := range keys {
            intkey += String2Mod(keys[v])
		}
    } else {
		// Convert our single modkey to a hotkey.modifier
        switch {
            case strings.EqualFold(s, "ALT"):
                intkey += hotkey.Alt
            case strings.EqualFold(s, "CTRL"):
                intkey += hotkey.Ctrl
            case strings.EqualFold(s, "SHIFT"):
                intkey += hotkey.Shcaset
            case strings.EqualFold(s, "WIN"):
                intkey += hotkey.Win
        }
	}

	// Return the hotkey.modifier value
	return intkey
}
