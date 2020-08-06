package hotkeyd

// Import the libraries we need
import (
	"strings"

	"github.com/MakeNowJust/hotkey"
)

// HotkeySwitch converts hotkey string to uint32
func HotkeySwitch(s string) uint32 {
	// Make a variable for our return value
	var returnval uint32

	// If our hotkey is more than 1 character
	if len(s) == 1 {
		// Convert the string from the config into a rune/char
		r := []rune(strings.ToUpper(s))

		// Return the rune/int32 as an uint32
		returnval = uint32(r[0])
	} else {
		// Make switches for keys that are more than a single character
		// And map them to the correct hotkey.modifier
		switch strings.ToUpper(s) {
		case "SPACE":
			returnval = hotkey.SPACE
		case "LBUTTON":
			returnval = hotkey.LBUTTON
		case "RBUTTON":
			returnval = hotkey.RBUTTON
		case "CANCEL":
			returnval = hotkey.CANCEL
		case "MBUTTON":
			returnval = hotkey.MBUTTON
		case "XBUTTON1":
			returnval = hotkey.XBUTTON1
		case "XBUTTON2":
			returnval = hotkey.XBUTTON2
		case "BACK":
			returnval = hotkey.BACK
		case "TAB":
			returnval = hotkey.TAB
		case "RETURN":
			returnval = hotkey.RETURN
		case "SHIFT":
			returnval = hotkey.SHIFT
		case "CONTROL":
			returnval = hotkey.CONTROL
		case "MENU":
			returnval = hotkey.MENU
		case "PAUSE":
			returnval = hotkey.PAUSE
		case "CAPITAL":
			returnval = hotkey.CAPITAL
		case "KANA":
			returnval = hotkey.KANA
		case "HANGUEL":
			returnval = hotkey.HANGUEL
		case "HANGUL":
			returnval = hotkey.HANGUL
		case "JUNJA":
			returnval = hotkey.JUNJA
		case "FINAL":
			returnval = hotkey.FINAL
		case "HANJA":
			returnval = hotkey.HANJA
		case "KANJI":
			returnval = hotkey.KANJI
		case "ESCAPE":
			returnval = hotkey.ESCAPE
		case "CONVERT":
			returnval = hotkey.CONVERT
		case "NONCONVERT":
			returnval = hotkey.NONCONVERT
		case "ACCEPT":
			returnval = hotkey.ACCEPT
		case "MODECHANGE":
			returnval = hotkey.MODECHANGE
		case "PRIOR":
			returnval = hotkey.PRIOR
		case "NEXT":
			returnval = hotkey.NEXT
		case "END":
			returnval = hotkey.END
		case "HOME":
			returnval = hotkey.HOME
		case "LEFT":
			returnval = hotkey.LEFT
		case "UP":
			returnval = hotkey.UP
		case "RIGHT":
			returnval = hotkey.RIGHT
		case "DOWN":
			returnval = hotkey.DOWN
		case "SELECT":
			returnval = hotkey.SELECT
		case "PRINT":
			returnval = hotkey.PRINT
		case "EXECUTE":
			returnval = hotkey.EXECUTE
		case "SNAPSHOT":
			returnval = hotkey.SNAPSHOT
		case "INSERT":
			returnval = hotkey.INSERT
		case "DELETE":
			returnval = hotkey.DELETE
		case "HELP":
			returnval = hotkey.HELP
		case "LWIN":
			returnval = hotkey.LWIN
		case "RWIN":
			returnval = hotkey.RWIN
		case "APPS":
			returnval = hotkey.APPS
		case "NUMPAD0":
			returnval = hotkey.NUMPAD0
		case "NUMPAD1":
			returnval = hotkey.NUMPAD1
		case "NUMPAD2":
			returnval = hotkey.NUMPAD2
		case "NUMPAD3":
			returnval = hotkey.NUMPAD3
		case "NUMPAD4":
			returnval = hotkey.NUMPAD4
		case "NUMPAD5":
			returnval = hotkey.NUMPAD5
		case "NUMPAD6":
			returnval = hotkey.NUMPAD6
		case "NUMPAD7":
			returnval = hotkey.NUMPAD7
		case "NUMPAD8":
			returnval = hotkey.NUMPAD8
		case "NUMPAD9":
			returnval = hotkey.NUMPAD9
		case "MULTIPLY":
			returnval = hotkey.MULTIPLY
		case "ADD":
			returnval = hotkey.ADD
		case "SEPARATOR":
			returnval = hotkey.SEPARATOR
		case "SUBTRACT":
			returnval = hotkey.SUBTRACT
		case "DECIMAL":
			returnval = hotkey.DECIMAL
		case "DIVIDE":
			returnval = hotkey.DIVIDE
		case "F1":
			returnval = hotkey.F1
		case "F2":
			returnval = hotkey.F2
		case "F3":
			returnval = hotkey.F3
		case "F4":
			returnval = hotkey.F4
		case "F5":
			returnval = hotkey.F5
		case "F6":
			returnval = hotkey.F6
		case "F7":
			returnval = hotkey.F7
		case "F8":
			returnval = hotkey.F8
		case "F9":
			returnval = hotkey.F9
		case "F10":
			returnval = hotkey.F10
		case "F11":
			returnval = hotkey.F11
		case "F12":
			returnval = hotkey.F12
		case "F13":
			returnval = hotkey.F13
		case "F14":
			returnval = hotkey.F14
		case "F15":
			returnval = hotkey.F15
		case "F16":
			returnval = hotkey.F16
		case "F17":
			returnval = hotkey.F17
		case "F18":
			returnval = hotkey.F18
		case "F19":
			returnval = hotkey.F19
		case "F20":
			returnval = hotkey.F20
		case "F21":
			returnval = hotkey.F21
		case "F22":
			returnval = hotkey.F22
		case "F23":
			returnval = hotkey.F23
		case "F24":
			returnval = hotkey.F24
		case "NUMLOCK":
			returnval = hotkey.NUMLOCK
		case "LSHIFT":
			returnval = hotkey.LSHIFT
		case "RSHIFT":
			returnval = hotkey.RSHIFT
		case "SCROLL":
			returnval = hotkey.SCROLL
		case "LCONTROL":
			returnval = hotkey.LCONTROL
		case "RCONTROL":
			returnval = hotkey.RCONTROL
		case "LMENU":
			returnval = hotkey.LMENU
		case "RMENU":
			returnval = hotkey.RMENU
		case "BROWSER_BACK":
			returnval = hotkey.BROWSER_BACK
		case "BROWSER_FORWARD":
			returnval = hotkey.BROWSER_FORWARD
		case "BROWSER_REFRESH":
			returnval = hotkey.BROWSER_REFRESH
		case "BROWSER_STOP":
			returnval = hotkey.BROWSER_STOP
		case "BROWSER_SEARCH":
			returnval = hotkey.BROWSER_SEARCH
		case "BROWSER_FAVORITES":
			returnval = hotkey.BROWSER_FAVORITES
		case "BROWSER_HOME":
			returnval = hotkey.BROWSER_HOME
		case "VOLUME_MUTE":
			returnval = hotkey.VOLUME_MUTE
		case "VOLUME_DOWN":
			returnval = hotkey.VOLUME_DOWN
		case "VOLUME_UP":
			returnval = hotkey.VOLUME_UP
		case "MEDIA_NEXT_TRACK":
			returnval = hotkey.MEDIA_NEXT_TRACK
		case "MEDIA_PREV_TRACK":
			returnval = hotkey.MEDIA_PREV_TRACK
		case "MEDIA_STOP":
			returnval = hotkey.MEDIA_STOP
		case "MEDIA_PLAY_PAUSE":
			returnval = hotkey.MEDIA_PLAY_PAUSE
		case "LAUNCH_MAIL":
			returnval = hotkey.LAUNCH_MAIL
		case "LAUNCH_MEDIA_SELECT":
			returnval = hotkey.LAUNCH_MEDIA_SELECT
		case "LAUNCH_APP1":
			returnval = hotkey.LAUNCH_APP1
		case "LAUNCH_APP2":
			returnval = hotkey.LAUNCH_APP2
		case "OEM_1":
			returnval = hotkey.OEM_1
		case "OEM_PLUS":
			returnval = hotkey.OEM_PLUS
		case "OEM_COMMA":
			returnval = hotkey.OEM_COMMA
		case "OEM_MINUS":
			returnval = hotkey.OEM_MINUS
		case "OEM_PERIOD":
			returnval = hotkey.OEM_PERIOD
		case "OEM_2":
			returnval = hotkey.OEM_2
		case "OEM_3":
			returnval = hotkey.OEM_3
		case "OEM_4":
			returnval = hotkey.OEM_4
		case "OEM_5":
			returnval = hotkey.OEM_5
		case "OEM_6":
			returnval = hotkey.OEM_6
		case "OEM_7":
			returnval = hotkey.OEM_7
		case "OEM_8":
			returnval = hotkey.OEM_8
		case "OEM_102":
			returnval = hotkey.OEM_102
		case "PROCESSKEY":
			returnval = hotkey.PROCESSKEY
		case "PACKET":
			returnval = hotkey.PACKET
		case "ATTN":
			returnval = hotkey.ATTN
		case "CRSEL":
			returnval = hotkey.CRSEL
		case "EXSEL":
			returnval = hotkey.EXSEL
		case "EREOF":
			returnval = hotkey.EREOF
		case "PLAY":
			returnval = hotkey.PLAY
		case "ZOOM":
			returnval = hotkey.ZOOM
		case "NONAME":
			returnval = hotkey.NONAME
		case "PA1":
			returnval = hotkey.PA1
		case "OEM_CLEAR":
			returnval = hotkey.OEM_CLEAR
		}
	}
	return returnval
}
