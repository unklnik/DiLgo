package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	exit bool

	//MOUSE
	MOUSE, mouseClickPoint sdl.FPoint
	LCLICK, RCLICK, MCLICK bool
	clickHoldT             int

	//KEYS
	KEYS         []KEY
	KESC, KF1    bool
	KEYDOWNTIMER int
)

type KEY struct {
	k            string
	code         sdl.Scancode
	codeK        sdl.Keycode
	on, released bool
	timer        time.Duration
}

func INP() {

	sdl.PumpEvents()

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch k := event.(type) {
		case sdl.QuitEvent:
			EXIT()
		case sdl.KeyboardEvent:
			if event.GetType() == sdl.KEYDOWN {
				uKEYSDOWN(k.Keysym.Scancode, k.Keysym.Sym)
			}
			if event.GetType() == sdl.KEYUP {
				uKEYSUP(k.Keysym.Scancode, k.Keysym.Sym)
			}
		case sdl.MouseButtonEvent:
			if event.GetType() == sdl.MOUSEBUTTONDOWN {
				if k.Button == sdl.ButtonLeft {
					if !LCLICK {
						mouseClickPoint = MOUSE
					}
					LCLICK = true
				}
				if k.Button == sdl.ButtonRight {
					RCLICK = true
				}
				if k.Button == sdl.ButtonMiddle {
					MCLICK = true
				}
			}
			if event.GetType() == sdl.MOUSEBUTTONUP {
				if k.Button == sdl.ButtonLeft {
					LCLICK = false
					clickHoldT = 0
				}
				if k.Button == sdl.ButtonRight {
					RCLICK = false
				}
				if k.Button == sdl.ButtonMiddle {
					MCLICK = false
				}
			}
		}

	}

	if LCLICK {
		clickHoldT++
	}

	uKEYS()

	//MOUSE
	x, y, _ := sdl.GetMouseState()
	MOUSE.X = float32(x)
	MOUSE.Y = float32(y)

	uKEYSOFF()
}

// MARK: UPDATE
func uKEYS() {
	if KEYDOWNTIMER == 0 {
		if KEYPRESS("F1") {
			DEBUG = !DEBUG
		}
		if KEYPRESS("F2") {

		}
		if KEYPRESS("ESC") {
			EXIT()
		}
		KEYDOWNTIMER = 7
	}
}
func uKEYSOFF() {
	for i := range KEYS {
		if KEYS[i].on {
			if KEYS[i].released {
				if KEYS[i].timer > 0 {
					KEYS[i].timer -= time.Millisecond
					if KEYS[i].timer < 0 {
						KEYS[i].timer = 0
					}
				} else {
					KEYS[i].on = false
					KEYS[i].released = false
				}
			}
		}
	}
}
func uKEYSDOWN(s sdl.Scancode, k sdl.Keycode) {
	for i := range KEYS {
		if KEYS[i].code == s || KEYS[i].codeK == k {
			KEYS[i].on = true
			KEYS[i].timer = time.Second / 100
			break
		}
	}
}
func uKEYSUP(s sdl.Scancode, k sdl.Keycode) {
	for i := range KEYS {
		if KEYS[i].code == s || KEYS[i].codeK == k {
			KEYS[i].released = true
			break
		}
	}
}
func KEYPRESS(k string) bool {
	on := false
	for i := range KEYS {
		if KEYS[i].k == k {
			on = KEYS[i].on
			break
		}
	}
	return on
}
func KEYDOWN(k string) bool {
	on := false
	for i := range KEYS {
		if KEYS[i].k == k {
			on = KEYS[i].on
			break
		}
	}
	return on
}

// MARK: UTILS
func mStrings2keys(s []string) []KEY {
	var k []KEY
	for i := range s {
		k = append(k, Mstring2key(s[i]))
	}
	return k
}
func Mstring2key(s string) KEY {
	k := KEY{}
	found := false
	for i := range KEYS {
		if KEYS[i].k == s {
			k = KEYS[i]
			found = true
			break
		}
	}
	if !found {
		Mmsg("ERROR: func Mstrings2key: Did not find a matching key >> HINT: Use func Dkeys() to draw available keys")
	}
	return k
}
func Dkeys(x, y float32, c sdl.Color) {
	ox := x
	for i := range KEYS {
		siz := FONT1DEFAULT.smlrH + 8
		if KEYS[i].on {
			DrecFill(sdl.FRect{x - 4, y - 4, CtxtLen(KEYS[i].k, FONT1DEFAULT, 1) + 8, siz}, ORANGE())
		}
		DrecLine(sdl.FRect{x - 4, y - 4, CtxtLen(KEYS[i].k, FONT1DEFAULT, 1) + 8, siz}, c)
		DtxtXY(KEYS[i].k, x, y, FONT1DEFAULT, 1, c)
		x += CtxtLen(KEYS[i].k, FONT1DEFAULT, 1) + 16
		if i < len(KEYS)-2 {
			if x+CtxtLen(KEYS[i+1].k, FONT1DEFAULT, 1)+16 > float32(WINW) {
				x = ox
				y += FONT1DEFAULT.smlrH + 16
			}
		}

	}
}

// MARK: MAKE
func mKEYS() {
	k := KEY{}
	k.k = "0"
	k.code = sdl.SCANCODE_0
	KEYS = append(KEYS, k)
	k.k = "1"
	k.code = sdl.SCANCODE_1
	KEYS = append(KEYS, k)
	k.k = "2"
	k.code = sdl.SCANCODE_2
	KEYS = append(KEYS, k)
	k.k = "3"
	k.code = sdl.SCANCODE_3
	KEYS = append(KEYS, k)
	k.k = "4"
	k.code = sdl.SCANCODE_4
	KEYS = append(KEYS, k)
	k.k = "5"
	k.code = sdl.SCANCODE_5
	KEYS = append(KEYS, k)
	k.k = "6"
	k.code = sdl.SCANCODE_6
	KEYS = append(KEYS, k)
	k.k = "7"
	k.code = sdl.SCANCODE_7
	KEYS = append(KEYS, k)
	k.k = "8"
	k.code = sdl.SCANCODE_8
	KEYS = append(KEYS, k)
	k.k = "9"
	k.code = sdl.SCANCODE_9
	KEYS = append(KEYS, k)
	k.k = "A"
	k.code = sdl.SCANCODE_A
	KEYS = append(KEYS, k)
	k.k = "ACBACK"
	k.code = sdl.SCANCODE_AC_BACK
	KEYS = append(KEYS, k)
	k.k = "ACBOOKMARKS"
	k.code = sdl.SCANCODE_AC_BOOKMARKS
	KEYS = append(KEYS, k)
	k.k = "ACFORWARD"
	k.code = sdl.SCANCODE_AC_FORWARD
	KEYS = append(KEYS, k)
	k.k = "ACHOME"
	k.code = sdl.SCANCODE_AC_HOME
	KEYS = append(KEYS, k)
	k.k = "ACREFRESH"
	k.code = sdl.SCANCODE_AC_REFRESH
	KEYS = append(KEYS, k)
	k.k = "ACSEARCH"
	k.code = sdl.SCANCODE_AC_SEARCH
	KEYS = append(KEYS, k)
	k.k = "ACSTOP"
	k.code = sdl.SCANCODE_AC_STOP
	KEYS = append(KEYS, k)
	k.k = "AGAIN"
	k.code = sdl.SCANCODE_AGAIN
	KEYS = append(KEYS, k)
	k.k = "ALTERASE"
	k.code = sdl.SCANCODE_ALTERASE
	KEYS = append(KEYS, k)
	k.k = "'"
	k.code = sdl.SCANCODE_APOSTROPHE
	KEYS = append(KEYS, k)
	k.k = "APPLICATION"
	k.code = sdl.SCANCODE_APPLICATION
	KEYS = append(KEYS, k)
	k.k = "AUDIOMUTE"
	k.code = sdl.SCANCODE_AUDIOMUTE
	KEYS = append(KEYS, k)
	k.k = "AUDIONEXT"
	k.code = sdl.SCANCODE_AUDIONEXT
	KEYS = append(KEYS, k)
	k.k = "AUDIOPLAY"
	k.code = sdl.SCANCODE_AUDIOPLAY
	KEYS = append(KEYS, k)
	k.k = "AUDIOPREV"
	k.code = sdl.SCANCODE_AUDIOPREV
	KEYS = append(KEYS, k)
	k.k = "AUDIOSTOP"
	k.code = sdl.SCANCODE_AUDIOSTOP
	KEYS = append(KEYS, k)
	k.k = "B"
	k.code = sdl.SCANCODE_B
	KEYS = append(KEYS, k)
	k.k = "\\"
	k.code = sdl.SCANCODE_BACKSLASH
	KEYS = append(KEYS, k)
	k.k = "BACKSPACE"
	k.code = sdl.SCANCODE_BACKSPACE
	KEYS = append(KEYS, k)
	k.k = "BRIGHTNESSDOWN"
	k.code = sdl.SCANCODE_BRIGHTNESSDOWN
	KEYS = append(KEYS, k)
	k.k = "BRIGHTNESSUP"
	k.code = sdl.SCANCODE_BRIGHTNESSUP
	KEYS = append(KEYS, k)
	k.k = "C"
	k.code = sdl.SCANCODE_C
	KEYS = append(KEYS, k)
	k.k = "CALCULATOR"
	k.code = sdl.SCANCODE_CALCULATOR
	KEYS = append(KEYS, k)
	k.k = "CANCEL"
	k.code = sdl.SCANCODE_CANCEL
	KEYS = append(KEYS, k)
	k.k = "CAPSLOCK"
	k.code = sdl.SCANCODE_CAPSLOCK
	KEYS = append(KEYS, k)
	k.k = "CLEAR"
	k.code = sdl.SCANCODE_CLEAR
	KEYS = append(KEYS, k)
	k.k = "CLEAR/AGAIN"
	k.code = sdl.SCANCODE_CLEARAGAIN
	KEYS = append(KEYS, k)
	k.k = ","
	k.code = sdl.SCANCODE_COMMA
	KEYS = append(KEYS, k)
	k.k = "COMPUTER"
	k.code = sdl.SCANCODE_COMPUTER
	KEYS = append(KEYS, k)
	k.k = "COPY"
	k.code = sdl.SCANCODE_COPY
	KEYS = append(KEYS, k)
	k.k = "CRSEL"
	k.code = sdl.SCANCODE_CRSEL
	KEYS = append(KEYS, k)
	k.k = "CURRENCYSUBUNIT"
	k.code = sdl.SCANCODE_CURRENCYSUBUNIT
	KEYS = append(KEYS, k)
	k.k = "CURRENCYUNIT"
	k.code = sdl.SCANCODE_CURRENCYUNIT
	KEYS = append(KEYS, k)
	k.k = "CUT"
	k.code = sdl.SCANCODE_CUT
	KEYS = append(KEYS, k)
	k.k = "D"
	k.code = sdl.SCANCODE_D
	KEYS = append(KEYS, k)
	k.k = "DECIMALSEPARATOR"
	k.code = sdl.SCANCODE_DECIMALSEPARATOR
	KEYS = append(KEYS, k)
	k.k = "DELETE"
	k.code = sdl.SCANCODE_DELETE
	KEYS = append(KEYS, k)
	k.k = "DISPLAYSWITCH"
	k.code = sdl.SCANCODE_DISPLAYSWITCH
	KEYS = append(KEYS, k)
	k.k = "DOWN"
	k.code = sdl.SCANCODE_DOWN
	KEYS = append(KEYS, k)
	k.k = "E"
	k.code = sdl.SCANCODE_E
	KEYS = append(KEYS, k)
	k.k = "EJECT"
	k.code = sdl.SCANCODE_EJECT
	KEYS = append(KEYS, k)
	k.k = "END"
	k.code = sdl.SCANCODE_END
	KEYS = append(KEYS, k)
	k.k = "="
	k.code = sdl.SCANCODE_EQUALS
	KEYS = append(KEYS, k)
	k.k = "ESC"
	k.code = sdl.SCANCODE_ESCAPE
	KEYS = append(KEYS, k)
	k.k = "EXECUTE"
	k.code = sdl.SCANCODE_EXECUTE
	KEYS = append(KEYS, k)
	k.k = "EXSEL"
	k.code = sdl.SCANCODE_EXSEL
	KEYS = append(KEYS, k)
	k.k = "F"
	k.code = sdl.SCANCODE_F
	KEYS = append(KEYS, k)
	k.k = "F1"
	k.code = sdl.SCANCODE_F1
	KEYS = append(KEYS, k)
	k.k = "F10"
	k.code = sdl.SCANCODE_F10
	KEYS = append(KEYS, k)
	k.k = "F11"
	k.code = sdl.SCANCODE_F11
	KEYS = append(KEYS, k)
	k.k = "F12"
	k.code = sdl.SCANCODE_F12
	KEYS = append(KEYS, k)
	k.k = "F13"
	k.code = sdl.SCANCODE_F13
	KEYS = append(KEYS, k)
	k.k = "F14"
	k.code = sdl.SCANCODE_F14
	KEYS = append(KEYS, k)
	k.k = "F15"
	k.code = sdl.SCANCODE_F15
	KEYS = append(KEYS, k)
	k.k = "F16"
	k.code = sdl.SCANCODE_F16
	KEYS = append(KEYS, k)
	k.k = "F17"
	k.code = sdl.SCANCODE_F17
	KEYS = append(KEYS, k)
	k.k = "F18"
	k.code = sdl.SCANCODE_F18
	KEYS = append(KEYS, k)
	k.k = "F19"
	k.code = sdl.SCANCODE_F19
	KEYS = append(KEYS, k)
	k.k = "F2"
	k.code = sdl.SCANCODE_F2
	KEYS = append(KEYS, k)
	k.k = "F20"
	k.code = sdl.SCANCODE_F20
	KEYS = append(KEYS, k)
	k.k = "F21"
	k.code = sdl.SCANCODE_F21
	KEYS = append(KEYS, k)
	k.k = "F22"
	k.code = sdl.SCANCODE_F22
	KEYS = append(KEYS, k)
	k.k = "F23"
	k.code = sdl.SCANCODE_F23
	KEYS = append(KEYS, k)
	k.k = "F24"
	k.code = sdl.SCANCODE_F24
	KEYS = append(KEYS, k)
	k.k = "F3"
	k.code = sdl.SCANCODE_F3
	KEYS = append(KEYS, k)
	k.k = "F4"
	k.code = sdl.SCANCODE_F4
	KEYS = append(KEYS, k)
	k.k = "F5"
	k.code = sdl.SCANCODE_F5
	KEYS = append(KEYS, k)
	k.k = "F6"
	k.code = sdl.SCANCODE_F6
	KEYS = append(KEYS, k)
	k.k = "F7"
	k.code = sdl.SCANCODE_F7
	KEYS = append(KEYS, k)
	k.k = "F8"
	k.code = sdl.SCANCODE_F8
	KEYS = append(KEYS, k)
	k.k = "F9"
	k.code = sdl.SCANCODE_F9
	KEYS = append(KEYS, k)
	k.k = "FIND"
	k.code = sdl.SCANCODE_FIND
	KEYS = append(KEYS, k)
	k.k = "G"
	k.code = sdl.SCANCODE_G
	KEYS = append(KEYS, k)
	k.k = "`"
	k.code = sdl.SCANCODE_GRAVE
	KEYS = append(KEYS, k)
	k.k = "H"
	k.code = sdl.SCANCODE_H
	KEYS = append(KEYS, k)
	k.k = "HELP"
	k.code = sdl.SCANCODE_HELP
	KEYS = append(KEYS, k)
	k.k = "HOME"
	k.code = sdl.SCANCODE_HOME
	KEYS = append(KEYS, k)
	k.k = "I"
	k.code = sdl.SCANCODE_I
	KEYS = append(KEYS, k)
	k.k = "INSERT"
	k.code = sdl.SCANCODE_INSERT
	KEYS = append(KEYS, k)
	k.k = "J"
	k.code = sdl.SCANCODE_J
	KEYS = append(KEYS, k)
	k.k = "K"
	k.code = sdl.SCANCODE_K
	KEYS = append(KEYS, k)
	k.k = "KBDILLUMDOWN"
	k.code = sdl.SCANCODE_KBDILLUMDOWN
	KEYS = append(KEYS, k)
	k.k = "KBDILLUMTOGGLE"
	k.code = sdl.SCANCODE_KBDILLUMTOGGLE
	KEYS = append(KEYS, k)
	k.k = "KBDILLUMUP"
	k.code = sdl.SCANCODE_KBDILLUMUP
	KEYS = append(KEYS, k)
	k.k = "KEYPAD0"
	k.code = sdl.SCANCODE_KP_0
	KEYS = append(KEYS, k)
	k.k = "KEYPAD00"
	k.code = sdl.SCANCODE_KP_00
	KEYS = append(KEYS, k)
	k.k = "KEYPAD000"
	k.code = sdl.SCANCODE_KP_000
	KEYS = append(KEYS, k)
	k.k = "KEYPAD1"
	k.code = sdl.SCANCODE_KP_1
	KEYS = append(KEYS, k)
	k.k = "KEYPAD2"
	k.code = sdl.SCANCODE_KP_2
	KEYS = append(KEYS, k)
	k.k = "KEYPAD3"
	k.code = sdl.SCANCODE_KP_3
	KEYS = append(KEYS, k)
	k.k = "KEYPAD4"
	k.code = sdl.SCANCODE_KP_4
	KEYS = append(KEYS, k)
	k.k = "KEYPAD5"
	k.code = sdl.SCANCODE_KP_5
	KEYS = append(KEYS, k)
	k.k = "KEYPAD6"
	k.code = sdl.SCANCODE_KP_6
	KEYS = append(KEYS, k)
	k.k = "KEYPAD7"
	k.code = sdl.SCANCODE_KP_7
	KEYS = append(KEYS, k)
	k.k = "KEYPAD8"
	k.code = sdl.SCANCODE_KP_8
	KEYS = append(KEYS, k)
	k.k = "KEYPAD9"
	k.code = sdl.SCANCODE_KP_9
	KEYS = append(KEYS, k)
	k.k = "KEYPADA"
	k.code = sdl.SCANCODE_KP_A
	KEYS = append(KEYS, k)
	k.k = "KEYPAD&"
	k.code = sdl.SCANCODE_KP_AMPERSAND
	KEYS = append(KEYS, k)
	k.k = "KEYPAD@"
	k.code = sdl.SCANCODE_KP_AT
	KEYS = append(KEYS, k)
	k.k = "KEYPADB"
	k.code = sdl.SCANCODE_KP_B
	KEYS = append(KEYS, k)
	k.k = "KEYPADBACKSPACE"
	k.code = sdl.SCANCODE_KP_BACKSPACE
	KEYS = append(KEYS, k)
	k.k = "KEYPADBINARY"
	k.code = sdl.SCANCODE_KP_BINARY
	KEYS = append(KEYS, k)
	k.k = "KEYPADC"
	k.code = sdl.SCANCODE_KP_C
	KEYS = append(KEYS, k)
	k.k = "KEYPADCLEAR"
	k.code = sdl.SCANCODE_KP_CLEAR
	KEYS = append(KEYS, k)
	k.k = "KEYPADCLEARENTRY"
	k.code = sdl.SCANCODE_KP_CLEARENTRY
	KEYS = append(KEYS, k)
	k.k = "KEYPAD:"
	k.code = sdl.SCANCODE_KP_COLON
	KEYS = append(KEYS, k)
	k.k = "KEYPAD,"
	k.code = sdl.SCANCODE_KP_COMMA
	KEYS = append(KEYS, k)
	k.k = "KEYPADD"
	k.code = sdl.SCANCODE_KP_D
	KEYS = append(KEYS, k)
	k.k = "KEYPAD&&"
	k.code = sdl.SCANCODE_KP_DBLAMPERSAND
	KEYS = append(KEYS, k)
	k.k = "KEYPAD||"
	k.code = sdl.SCANCODE_KP_DBLVERTICALBAR
	KEYS = append(KEYS, k)
	k.k = "KEYPADDECIMAL"
	k.code = sdl.SCANCODE_KP_DECIMAL
	KEYS = append(KEYS, k)
	k.k = "KEYPAD/"
	k.code = sdl.SCANCODE_KP_DIVIDE
	KEYS = append(KEYS, k)
	k.k = "KEYPADE"
	k.code = sdl.SCANCODE_KP_E
	KEYS = append(KEYS, k)
	k.k = "KEYPADENTER"
	k.code = sdl.SCANCODE_KP_ENTER
	KEYS = append(KEYS, k)
	k.k = "KEYPAD="
	k.code = sdl.SCANCODE_KP_EQUALS
	KEYS = append(KEYS, k)
	k.k = "KEYPAD=(AS400)"
	k.code = sdl.SCANCODE_KP_EQUALSAS400
	KEYS = append(KEYS, k)
	k.k = "KEYPAD!"
	k.code = sdl.SCANCODE_KP_EXCLAM
	KEYS = append(KEYS, k)
	k.k = "KEYPADF"
	k.code = sdl.SCANCODE_KP_F
	KEYS = append(KEYS, k)
	k.k = "KEYPAD>"
	k.code = sdl.SCANCODE_KP_GREATER
	KEYS = append(KEYS, k)
	k.k = "KEYPAD#"
	k.code = sdl.SCANCODE_KP_HASH
	KEYS = append(KEYS, k)
	k.k = "KEYPADHEXADECIMAL"
	k.code = sdl.SCANCODE_KP_HEXADECIMAL
	KEYS = append(KEYS, k)
	k.k = "KEYPAD{"
	k.code = sdl.SCANCODE_KP_LEFTBRACE
	KEYS = append(KEYS, k)
	k.k = "KEYPAD("
	k.code = sdl.SCANCODE_KP_LEFTPAREN
	KEYS = append(KEYS, k)
	k.k = "KEYPAD<"
	k.code = sdl.SCANCODE_KP_LESS
	KEYS = append(KEYS, k)
	k.k = "KEYPADMEMADD"
	k.code = sdl.SCANCODE_KP_MEMADD
	KEYS = append(KEYS, k)
	k.k = "KEYPADMEMCLEAR"
	k.code = sdl.SCANCODE_KP_MEMCLEAR
	KEYS = append(KEYS, k)
	k.k = "KEYPADMEMDIVIDE"
	k.code = sdl.SCANCODE_KP_MEMDIVIDE
	KEYS = append(KEYS, k)
	k.k = "KEYPADMEMMULTIPLY"
	k.code = sdl.SCANCODE_KP_MEMMULTIPLY
	KEYS = append(KEYS, k)
	k.k = "KEYPADMEMRECALL"
	k.code = sdl.SCANCODE_KP_MEMRECALL
	KEYS = append(KEYS, k)
	k.k = "KEYPADMEMSTORE"
	k.code = sdl.SCANCODE_KP_MEMSTORE
	KEYS = append(KEYS, k)
	k.k = "KEYPADMEMSUBTRACT"
	k.code = sdl.SCANCODE_KP_MEMSUBTRACT
	KEYS = append(KEYS, k)
	k.k = "KEYPAD-"
	k.code = sdl.SCANCODE_KP_MINUS
	KEYS = append(KEYS, k)
	k.k = "KEYPAD*"
	k.code = sdl.SCANCODE_KP_MULTIPLY
	KEYS = append(KEYS, k)
	k.k = "KEYPADOCTAL"
	k.code = sdl.SCANCODE_KP_OCTAL
	KEYS = append(KEYS, k)
	k.k = "KEYPAD%"
	k.code = sdl.SCANCODE_KP_PERCENT
	KEYS = append(KEYS, k)
	k.k = "KEYPAD."
	k.code = sdl.SCANCODE_KP_PERIOD
	KEYS = append(KEYS, k)
	k.k = "KEYPAD+"
	k.code = sdl.SCANCODE_KP_PLUS
	KEYS = append(KEYS, k)
	k.k = "KEYPAD+/-"
	k.code = sdl.SCANCODE_KP_PLUSMINUS
	KEYS = append(KEYS, k)
	k.k = "KEYPAD^"
	k.code = sdl.SCANCODE_KP_POWER
	KEYS = append(KEYS, k)
	k.k = "KEYPAD}"
	k.code = sdl.SCANCODE_KP_RIGHTBRACE
	KEYS = append(KEYS, k)
	k.k = "KEYPAD)"
	k.code = sdl.SCANCODE_KP_RIGHTPAREN
	KEYS = append(KEYS, k)
	k.k = "KEYPADSPACE"
	k.code = sdl.SCANCODE_KP_SPACE
	KEYS = append(KEYS, k)
	k.k = "KEYPADTAB"
	k.code = sdl.SCANCODE_KP_TAB
	KEYS = append(KEYS, k)
	k.k = "KEYPAD|"
	k.code = sdl.SCANCODE_KP_VERTICALBAR
	KEYS = append(KEYS, k)
	k.k = "KEYPADXOR"
	k.code = sdl.SCANCODE_KP_XOR
	KEYS = append(KEYS, k)
	k.k = "L"
	k.code = sdl.SCANCODE_L
	KEYS = append(KEYS, k)
	k.k = "LEFTALT"
	k.code = sdl.SCANCODE_LALT
	KEYS = append(KEYS, k)
	k.k = "LEFTCTRL"
	k.code = sdl.SCANCODE_LCTRL
	KEYS = append(KEYS, k)
	k.k = "LEFT"
	k.code = sdl.SCANCODE_LEFT
	KEYS = append(KEYS, k)
	k.k = "["
	k.code = sdl.SCANCODE_LEFTBRACKET
	KEYS = append(KEYS, k)
	k.k = "LEFTGUI"
	k.code = sdl.SCANCODE_LGUI
	KEYS = append(KEYS, k)
	k.k = "LEFTSHIFT"
	k.code = sdl.SCANCODE_LSHIFT
	KEYS = append(KEYS, k)
	k.k = "M"
	k.code = sdl.SCANCODE_M
	KEYS = append(KEYS, k)
	k.k = "MAIL"
	k.code = sdl.SCANCODE_MAIL
	KEYS = append(KEYS, k)
	k.k = "MEDIASELECT"
	k.code = sdl.SCANCODE_MEDIASELECT
	KEYS = append(KEYS, k)
	k.k = "MENU"
	k.code = sdl.SCANCODE_MENU
	KEYS = append(KEYS, k)
	k.k = "-"
	k.code = sdl.SCANCODE_MINUS
	KEYS = append(KEYS, k)
	k.k = "MODESWITCH"
	k.code = sdl.SCANCODE_MODE
	KEYS = append(KEYS, k)
	k.k = "MUTE"
	k.code = sdl.SCANCODE_MUTE
	KEYS = append(KEYS, k)
	k.k = "N"
	k.code = sdl.SCANCODE_N
	KEYS = append(KEYS, k)
	k.k = "NUMLOCK"
	k.code = sdl.SCANCODE_NUMLOCKCLEAR
	KEYS = append(KEYS, k)
	k.k = "O"
	k.code = sdl.SCANCODE_O
	KEYS = append(KEYS, k)
	k.k = "OPER"
	k.code = sdl.SCANCODE_OPER
	KEYS = append(KEYS, k)
	k.k = "OUT"
	k.code = sdl.SCANCODE_OUT
	KEYS = append(KEYS, k)
	k.k = "P"
	k.code = sdl.SCANCODE_P
	KEYS = append(KEYS, k)
	k.k = "PAGEDOWN"
	k.code = sdl.SCANCODE_PAGEDOWN
	KEYS = append(KEYS, k)
	k.k = "PAGEUP"
	k.code = sdl.SCANCODE_PAGEUP
	KEYS = append(KEYS, k)
	k.k = "PASTE"
	k.code = sdl.SCANCODE_PASTE
	KEYS = append(KEYS, k)
	k.k = "PAUSE"
	k.code = sdl.SCANCODE_PAUSE
	KEYS = append(KEYS, k)
	k.k = "."
	k.code = sdl.SCANCODE_PERIOD
	KEYS = append(KEYS, k)
	k.k = "POWER"
	k.code = sdl.SCANCODE_POWER
	KEYS = append(KEYS, k)
	k.k = "PRINTSCREEN"
	k.code = sdl.SCANCODE_PRINTSCREEN
	KEYS = append(KEYS, k)
	k.k = "PRIOR"
	k.code = sdl.SCANCODE_PRIOR
	KEYS = append(KEYS, k)
	k.k = "Q"
	k.code = sdl.SCANCODE_Q
	KEYS = append(KEYS, k)
	k.k = "R"
	k.code = sdl.SCANCODE_R
	KEYS = append(KEYS, k)
	k.k = "RIGHTALT"
	k.code = sdl.SCANCODE_RALT
	KEYS = append(KEYS, k)
	k.k = "RIGHTCTRL"
	k.code = sdl.SCANCODE_RCTRL
	KEYS = append(KEYS, k)
	k.k = "RETURN"
	k.code = sdl.SCANCODE_RETURN
	KEYS = append(KEYS, k)
	k.k = "RETURN"
	k.code = sdl.SCANCODE_RETURN2
	KEYS = append(KEYS, k)
	k.k = "RIGHTGUI"
	k.code = sdl.SCANCODE_RGUI
	KEYS = append(KEYS, k)
	k.k = "RIGHT"
	k.code = sdl.SCANCODE_RIGHT
	KEYS = append(KEYS, k)
	k.k = "]"
	k.code = sdl.SCANCODE_RIGHTBRACKET
	KEYS = append(KEYS, k)
	k.k = "RIGHTSHIFT"
	k.code = sdl.SCANCODE_RSHIFT
	KEYS = append(KEYS, k)
	k.k = "S"
	k.code = sdl.SCANCODE_S
	KEYS = append(KEYS, k)
	k.k = "SCROLLLOCK"
	k.code = sdl.SCANCODE_SCROLLLOCK
	KEYS = append(KEYS, k)
	k.k = "SELECT"
	k.code = sdl.SCANCODE_SELECT
	KEYS = append(KEYS, k)
	k.k = ";"
	k.code = sdl.SCANCODE_SEMICOLON
	KEYS = append(KEYS, k)
	k.k = "SEPARATOR"
	k.code = sdl.SCANCODE_SEPARATOR
	KEYS = append(KEYS, k)
	k.k = "/"
	k.code = sdl.SCANCODE_SLASH
	KEYS = append(KEYS, k)
	k.k = "SLEEP"
	k.code = sdl.SCANCODE_SLEEP
	KEYS = append(KEYS, k)
	k.k = "SPACE"
	k.code = sdl.SCANCODE_SPACE
	KEYS = append(KEYS, k)
	k.k = "STOP"
	k.code = sdl.SCANCODE_STOP
	KEYS = append(KEYS, k)
	k.k = "SYSREQ"
	k.code = sdl.SCANCODE_SYSREQ
	KEYS = append(KEYS, k)
	k.k = "T"
	k.code = sdl.SCANCODE_T
	KEYS = append(KEYS, k)
	k.k = "TAB"
	k.code = sdl.SCANCODE_TAB
	KEYS = append(KEYS, k)
	k.k = "THOUSANDSSEPARATOR"
	k.code = sdl.SCANCODE_THOUSANDSSEPARATOR
	KEYS = append(KEYS, k)
	k.k = "U"
	k.code = sdl.SCANCODE_U
	KEYS = append(KEYS, k)
	k.k = "UNDO"
	k.code = sdl.SCANCODE_UNDO
	KEYS = append(KEYS, k)
	k.k = ""
	k.code = sdl.SCANCODE_UNKNOWN
	KEYS = append(KEYS, k)
	k.k = "UP"
	k.code = sdl.SCANCODE_UP
	KEYS = append(KEYS, k)
	k.k = "V"
	k.code = sdl.SCANCODE_V
	KEYS = append(KEYS, k)
	k.k = "VOLUMEDOWN"
	k.code = sdl.SCANCODE_VOLUMEDOWN
	KEYS = append(KEYS, k)
	k.k = "VOLUMEUP"
	k.code = sdl.SCANCODE_VOLUMEUP
	KEYS = append(KEYS, k)
	k.k = "W"
	k.code = sdl.SCANCODE_W
	KEYS = append(KEYS, k)
	k.k = "WWW"
	k.code = sdl.SCANCODE_WWW
	KEYS = append(KEYS, k)
	k.k = "X"
	k.code = sdl.SCANCODE_X
	KEYS = append(KEYS, k)
	k.k = "Y"
	k.code = sdl.SCANCODE_Y
	KEYS = append(KEYS, k)
	k.k = "Z"
	k.code = sdl.SCANCODE_Z
	KEYS = append(KEYS, k)
	k.k = "&"
	k.codeK = sdl.K_AMPERSAND
	KEYS = append(KEYS, k)
	k.k = "*"
	k.codeK = sdl.K_ASTERISK
	KEYS = append(KEYS, k)
	k.k = "@"
	k.codeK = sdl.K_AT
	KEYS = append(KEYS, k)
	k.k = "^"
	k.codeK = sdl.K_CARET
	KEYS = append(KEYS, k)
	k.k = ":"
	k.codeK = sdl.K_COLON
	KEYS = append(KEYS, k)
	k.k = "$"
	k.codeK = sdl.K_DOLLAR
	KEYS = append(KEYS, k)
	k.k = "!"
	k.codeK = sdl.K_EXCLAIM
	KEYS = append(KEYS, k)
	k.k = ">"
	k.codeK = sdl.K_GREATER
	KEYS = append(KEYS, k)
	k.k = "#"
	k.codeK = sdl.K_HASH
	KEYS = append(KEYS, k)
	k.k = "("
	k.codeK = sdl.K_LEFTPAREN
	KEYS = append(KEYS, k)
	k.k = "<"
	k.codeK = sdl.K_LESS
	KEYS = append(KEYS, k)
	k.k = "%"
	k.codeK = sdl.K_PERCENT
	KEYS = append(KEYS, k)
	k.k = "+"
	k.codeK = sdl.K_PLUS
	KEYS = append(KEYS, k)
	k.k = "?"
	k.codeK = sdl.K_QUESTION
	KEYS = append(KEYS, k)
	k.k = "\""
	k.codeK = sdl.K_QUOTEDBL
	KEYS = append(KEYS, k)
	k.k = "_"
	k.codeK = sdl.K_RIGHTPAREN
	KEYS = append(KEYS, k)
	k.k = "_"
	k.codeK = sdl.K_UNDERSCORE
	KEYS = append(KEYS, k)
}
