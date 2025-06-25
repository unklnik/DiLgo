package main

import "github.com/veandco/go-sdl2/sdl"

var (
	BUTTONCOLON, BUTTONCOLOFF = DARKGREEN(), DARKRED()
	BUTTONSHEET               []BUTTON
	BUTTONSIZE                = float32(32)
)

type BUTTON struct {
	cOn, cOff        sdl.Color
	r                sdl.FRect
	onOff, canToggle bool
	im               IM
}

type WINDOW struct {
	nm                                     string
	r, rBar, rClose                        sdl.FRect
	cBar, cLine, cBG, cText                sdl.Color
	posFix, onoff, closeIcon, shadow, move bool
	outlineW                               int
	txt                                    string
}

// MENUS
func MwindowTextBox() {

}
func MwindowXY(name string, x, y, w, h float32, outlineW int, onOff, positionFixed, closeIcon, shadow bool, colBar, colOutline, colBackground, colText sdl.Color) WINDOW {
	m := WINDOW{}
	m.nm = name
	m.posFix = positionFixed
	m.onoff = onOff
	m.outlineW = outlineW
	m.closeIcon = closeIcon
	m.shadow = shadow
	m.r = sdl.FRect{x, y, w, h}
	m.rBar = sdl.FRect{m.r.X, m.r.Y, m.r.W, FONT1DEFAULT.smlrH + 8}
	wClose := m.rBar.H - (float32(m.outlineW*2) + 4)
	m.rClose = sdl.FRect{m.rBar.X + m.rBar.W - (wClose + (float32(m.outlineW*2) + 2)), m.r.Y + (float32(m.outlineW+2) + 1), wClose, wClose}
	m.cBar = colBar
	m.cLine = colOutline
	m.cBG = colBackground
	m.cText = colText
	return m
}
func MwindowCenter(name string, w, h float32, outlineW int, onOff, positionFixed, closeIcon, shadow bool, colBar, colOutline, colBackground, colText sdl.Color) WINDOW {
	m := WINDOW{}
	m.nm = name
	m.posFix = positionFixed
	m.onoff = onOff
	m.outlineW = outlineW
	m.closeIcon = closeIcon
	m.shadow = shadow
	m.r = sdl.FRect{CNTR.X - w/2, CNTR.Y - h/2, w, h}
	m.rBar = sdl.FRect{m.r.X, m.r.Y, m.r.W, FONT1DEFAULT.smlrH + 8}
	wClose := m.rBar.H - (float32(m.outlineW*2) + 4)
	m.rClose = sdl.FRect{m.rBar.X + m.rBar.W - (wClose + (float32(m.outlineW*2) + 2)), m.r.Y + (float32(m.outlineW+2) + 1), wClose, wClose}
	m.cBar = colBar
	m.cLine = colOutline
	m.cBG = colBackground
	m.cText = colText
	return m
}

// BUTTONS
func UbuttonToggle(b BUTTON) BUTTON {
	b.canToggle = !b.canToggle
	return b
}

func mBUTTONS() {

	b := BUTTON{}
	for i := range len(ICONSSML) {
		b.r = sdl.FRect{0, 0, BUTTONSIZE, BUTTONSIZE}
		b.cOn = BUTTONCOLON
		b.cOff = BUTTONCOLOFF
		b.im = ICONSSML[i]
		BUTTONSHEET = append(BUTTONSHEET, b)
	}

}
