# hotkeyD
## A hotkey daemon for windows that does not suck.


Let's be honest, the hotkey system for launching programs in windows just suck.
This is ALPHA stage software and I still need to work out a way to toggle the hotkeys on and off.

For now this is what works<br>
* Assign hotkeys to launch programs
* Launch programs and scripts in a hidden window state (good for automation scripts that you dont want to steal focus)
* Ability to use start command through cmd with the arguments "/c start" launch default applications for things like "start https://github.com/hikariknight" to open this website in your default web browser.

TODO:
* Toggle hotkey
* "Recreate" hotkeys when you hit the toggle hotkey again

NOTE: Because there is no way to disable hotkeys in the meantime, you will have to avoid creating hotkeys that conflict with hotkeys from other programs you use, as hotkeyD does not care what is in focus.

Here is an example config
```ini
[notepad]
Modkeys=ctrl+alt
Hotkey=n
Launch=C:\windows\system32\notepad.exe

[browser]
Modkeys=ctrl+alt
Hotkey=t
Launch=cmd.exe
Args=/c "start https://github.com/HikariKnight/hotkeyd"
Hide=true
```