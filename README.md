# hotkeyD
## A hotkey daemon for windows that does not suck.


Let's be honest, the hotkey system for launching programs in windows is just garbage, let us fix that!

For now this is what works<br>
* Assign hotkeys to launch programs
* Launch programs and scripts in a hidden window state (good for automation scripts that you dont want to steal focus)
* Ability to use start command through cmd with the arguments "/c start" launch default applications for things like "start https://github.com/hikariknight" to open this website in your default web browser.

TODO:
* Possible GUI for editing hotkeys.ini? 🤔
<br>

#

### Why use this over AutoHotKey?
This is not meant as a replacement for AutoHotKey, AHK is a fantastic script language and same with AutoIT which it is based on, however this project has no plans to be able to send keystrokes and manipulate windows. This is only meant to be a Launcher for your most used programs or background scripts that you want available with a simple keypress without anti-cheat software getting mad at you for using a tool that *can* be used for cheating.

#

If you plan to use special characters like `.` or `/` then please refer to [Windows Virtual Key Codes documentation](https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes) as some special keys have weird names inside windows like<br>
`OEM_COMMA` for `,`<br>
`OEM_PERIOD` for `.`<br>
and `OEM_2` for `/`

#

Here is an example `hotkeys.ini` config
```ini
# These undefined hotkey options defines the toggle keycombo
# This is mandatory for the daemon to work
# but you can change the keys to what you want
Modkeys=ctrl+alt
Hotkey=q

# Example "open notepad" hotkey
[notepad]
Modkeys=ctrl+alt
Hotkey=n
Launch=C:\windows\system32\notepad.exe

# Example for utilizing cmd.exe, the "start" command and the hide window feature
# To open a url or any known file or protocol in it's default application
[browser]
Modkeys=ctrl+alt
Hotkey=t
Launch=cmd.exe
Args=/c "start https://github.com/HikariKnight/hotkeyd"
Hide=true

# The "name" for the hotkey serves only as a separator for the program
# and you can use any combination of Modkeys (alt, ctrl, win, shift)
# However the win key is often reserved for many hotkey combinations that do
# Not exist yet
[a hotkey can be defined with any name]
Modkeys=alt+shift+win
Hotkey=c
Launch=calc.exe

# There is no built in way to "quit" the application
# This is intentional to avoid reserving more than 1 hotkey combination
# Here is an example on how to make a quit hotkey for hotkeyD, if you ever want one
[kill hotkeyD]
Modkeys=alt+win
Hotkey=q
Launch=cmd.exe
Args=/c "tskill hotkeyd"
Hide=true

# You are able to set a working directory for your program too by using the Workdir
# option, remember to put the path in quotes "like this".
[workdir example]
Modkeys=ctrl+alt
Hotkey=w
Launch=notepad.exe
Workdir="C:\"
```

# 

### Build instructions
There is a [build script](https://github.com/HikariKnight/hotkeyD/blob/master/src/build) written in bash since I developed this on Linux so I needed to cross compile it to use it in my VM, but it should be as simple as running this in the same folder as [hotkeyd.go](https://github.com/HikariKnight/hotkeyD/blob/master/src/app/hotkeyd.go)<br>
`go build -o hotkeyd.exe -v -ldflags -H=windowsgui hotkeyd.go`

If you are using Linux or MacOS as your build machine just run the build script with either `debug` or `release` as the parameter.
The debug build just omits the `-ldflags -H=windowsgui` parameters from the build command which hides the terminal window that go programs tend to have.

Just make sure you `go get` these dependencies before building
```
go get github.com/MakeNowJust/hotkey
go get github.com/kardianos/osext
go get gopkg.in/ini.v1
go get tawesoft.co.uk/go/dialog
```
