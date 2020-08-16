package main

// Import the libraries we need
import (
	"fmt"
	"runtime"

	"./hkdaemon"
)

// Make a main function
func main() {
	// Make a quit method
	quit := make(chan bool)

	if runtime.GOOS != "windows" {
		fmt.Printf("OS not supported!\n")
	}

	// Create hotkeys
	hkdaemon.CreateHotkeys()

	fmt.Println("Started hotkey loop")
	<-quit
}
