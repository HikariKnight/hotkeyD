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

	// Make a map for all the keycodes
	var hotkeymap = map[string]uint32{
		"SPACE":               hotkey.SPACE,
		"LBUTTON":             hotkey.LBUTTON,
		"RBUTTON":             hotkey.RBUTTON,
		"CANCEL":              hotkey.CANCEL,
		"MBUTTON":             hotkey.MBUTTON,
		"XBUTTON1":            hotkey.XBUTTON1,
		"XBUTTON2":            hotkey.XBUTTON2,
		"BACK":                hotkey.BACK,
		"TAB":                 hotkey.TAB,
		"RETURN":              hotkey.RETURN,
		"SHIFT":               hotkey.SHIFT,
		"CONTROL":             hotkey.CONTROL,
		"MENU":                hotkey.MENU,
		"PAUSE":               hotkey.PAUSE,
		"CAPITAL":             hotkey.CAPITAL,
		"KANA":                hotkey.KANA,
		"HANGUEL":             hotkey.HANGUEL,
		"HANGUL":              hotkey.HANGUL,
		"JUNJA":               hotkey.JUNJA,
		"FINAL":               hotkey.FINAL,
		"HANJA":               hotkey.HANJA,
		"KANJI":               hotkey.KANJI,
		"ESCAPE":              hotkey.ESCAPE,
		"CONVERT":             hotkey.CONVERT,
		"NONCONVERT":          hotkey.NONCONVERT,
		"ACCEPT":              hotkey.ACCEPT,
		"MODECHANGE":          hotkey.MODECHANGE,
		"PRIOR":               hotkey.PRIOR,
		"NEXT":                hotkey.NEXT,
		"END":                 hotkey.END,
		"HOME":                hotkey.HOME,
		"LEFT":                hotkey.LEFT,
		"UP":                  hotkey.UP,
		"RIGHT":               hotkey.RIGHT,
		"DOWN":                hotkey.DOWN,
		"SELECT":              hotkey.SELECT,
		"PRINT":               hotkey.PRINT,
		"EXECUTE":             hotkey.EXECUTE,
		"SNAPSHOT":            hotkey.SNAPSHOT,
		"INSERT":              hotkey.INSERT,
		"DELETE":              hotkey.DELETE,
		"HELP":                hotkey.HELP,
		"LWIN":                hotkey.LWIN,
		"RWIN":                hotkey.RWIN,
		"APPS":                hotkey.APPS,
		"NUMPAD0":             hotkey.NUMPAD0,
		"NUMPAD1":             hotkey.NUMPAD1,
		"NUMPAD2":             hotkey.NUMPAD2,
		"NUMPAD3":             hotkey.NUMPAD3,
		"NUMPAD4":             hotkey.NUMPAD4,
		"NUMPAD5":             hotkey.NUMPAD5,
		"NUMPAD6":             hotkey.NUMPAD6,
		"NUMPAD7":             hotkey.NUMPAD7,
		"NUMPAD8":             hotkey.NUMPAD8,
		"NUMPAD9":             hotkey.NUMPAD9,
		"MULTIPLY":            hotkey.MULTIPLY,
		"ADD":                 hotkey.ADD,
		"SEPARATOR":           hotkey.SEPARATOR,
		"SUBTRACT":            hotkey.SUBTRACT,
		"DECIMAL":             hotkey.DECIMAL,
		"DIVIDE":              hotkey.DIVIDE,
		"F1":                  hotkey.F1,
		"F2":                  hotkey.F2,
		"F3":                  hotkey.F3,
		"F4":                  hotkey.F4,
		"F5":                  hotkey.F5,
		"F6":                  hotkey.F6,
		"F7":                  hotkey.F7,
		"F8":                  hotkey.F8,
		"F9":                  hotkey.F9,
		"F10":                 hotkey.F10,
		"F11":                 hotkey.F11,
		"F12":                 hotkey.F12,
		"F13":                 hotkey.F13,
		"F14":                 hotkey.F14,
		"F15":                 hotkey.F15,
		"F16":                 hotkey.F16,
		"F17":                 hotkey.F17,
		"F18":                 hotkey.F18,
		"F19":                 hotkey.F19,
		"F20":                 hotkey.F20,
		"F21":                 hotkey.F21,
		"F22":                 hotkey.F22,
		"F23":                 hotkey.F23,
		"F24":                 hotkey.F24,
		"NUMLOCK":             hotkey.NUMLOCK,
		"LSHIFT":              hotkey.LSHIFT,
		"RSHIFT":              hotkey.RSHIFT,
		"SCROLL":              hotkey.SCROLL,
		"LCONTROL":            hotkey.LCONTROL,
		"RCONTROL":            hotkey.RCONTROL,
		"LMENU":               hotkey.LMENU,
		"RMENU":               hotkey.RMENU,
		"BROWSER_BACK":        hotkey.BROWSER_BACK,
		"BROWSER_FORWARD":     hotkey.BROWSER_FORWARD,
		"BROWSER_REFRESH":     hotkey.BROWSER_REFRESH,
		"BROWSER_STOP":        hotkey.BROWSER_STOP,
		"BROWSER_SEARCH":      hotkey.BROWSER_SEARCH,
		"BROWSER_FAVORITES":   hotkey.BROWSER_FAVORITES,
		"BROWSER_HOME":        hotkey.BROWSER_HOME,
		"VOLUME_MUTE":         hotkey.VOLUME_MUTE,
		"VOLUME_DOWN":         hotkey.VOLUME_DOWN,
		"VOLUME_UP":           hotkey.VOLUME_UP,
		"MEDIA_NEXT_TRACK":    hotkey.MEDIA_NEXT_TRACK,
		"MEDIA_PREV_TRACK":    hotkey.MEDIA_PREV_TRACK,
		"MEDIA_STOP":          hotkey.MEDIA_STOP,
		"MEDIA_PLAY_PAUSE":    hotkey.MEDIA_PLAY_PAUSE,
		"LAUNCH_MAIL":         hotkey.LAUNCH_MAIL,
		"LAUNCH_MEDIA_SELECT": hotkey.LAUNCH_MEDIA_SELECT,
		"LAUNCH_APP1":         hotkey.LAUNCH_APP1,
		"LAUNCH_APP2":         hotkey.LAUNCH_APP2,
		"OEM_1":               hotkey.OEM_1,
		"OEM_PLUS":            hotkey.OEM_PLUS,
		"OEM_COMMA":           hotkey.OEM_COMMA,
		"OEM_MINUS":           hotkey.OEM_MINUS,
		"OEM_PERIOD":          hotkey.OEM_PERIOD,
		"OEM_2":               hotkey.OEM_2,
		"OEM_3":               hotkey.OEM_3,
		"OEM_4":               hotkey.OEM_4,
		"OEM_5":               hotkey.OEM_5,
		"OEM_6":               hotkey.OEM_6,
		"OEM_7":               hotkey.OEM_7,
		"OEM_8":               hotkey.OEM_8,
		"OEM_102":             hotkey.OEM_102,
		"PROCESSKEY":          hotkey.PROCESSKEY,
		"PACKET":              hotkey.PACKET,
		"ATTN":                hotkey.ATTN,
		"CRSEL":               hotkey.CRSEL,
		"EXSEL":               hotkey.EXSEL,
		"EREOF":               hotkey.EREOF,
		"PLAY":                hotkey.PLAY,
		"ZOOM":                hotkey.ZOOM,
		"NONAME":              hotkey.NONAME,
		"PA1":                 hotkey.PA1,
		"OEM_CLEAR":           hotkey.OEM_CLEAR,
	}

	// If our hotkey is more than 1 character
	switch len(s) == 1 {
	case true:
		// Convert the string from the config into a rune/char
		r := []rune(strings.ToUpper(s))

		// Return the rune/int32 as an uint32
		returnval = uint32(r[0])

	case false:
		// Make switches for keys that are more than a single character
		// And map them to the correct hotkey.modifier
		returnval = hotkeymap[strings.ToUpper(s)]
	}

	// Return our unint32 value
	return returnval
}
