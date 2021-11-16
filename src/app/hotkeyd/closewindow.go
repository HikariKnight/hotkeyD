//go:build windows
// +build windows

package hotkeyd

import (
	"github.com/micmonay/keybd_event"
	"time"
)

func CloseWindow()  {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	
	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_F4) 

	// Set ALT to be pressed
	kb.HasALT(true) 

	time.Sleep(10 * time.Millisecond)
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()
}