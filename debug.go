package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	DEBUG         bool
	DISPLAYERRORS = true
	dispErr       []string
)

func DEBUGON() {
	var x, y float32 = 10, 10
	DrunTime("RUN TIME:", x, y, FONT1DEFAULT, 1, WHITE())
	y += float32(FONT1DEFAULT.charSMLR[0].r.H)
	DtxtXY("START TIME: "+fmt.Sprint(STARTTIME.Format("2006-01-02 15:04:05")), x, y, FONT1DEFAULT, 1, WHITE())
	y += float32(FONT1DEFAULT.charSMLR[0].r.H)
	DtxtXY("SCRW "+fmt.Sprint(SCRW)+" SCRH "+fmt.Sprint(SCRH), x, y, FONT1DEFAULT, 1, WHITE())
	y += float32(FONT1DEFAULT.charSMLR[0].r.H)
	DtxtXY("WINW "+fmt.Sprint(WINW)+" WINH "+fmt.Sprint(WINH), x, y, FONT1DEFAULT, 1, WHITE())
	y += float32(FONT1DEFAULT.charSMLR[0].r.H)
	DtxtXY("FPS "+fmt.Sprint(FPSCURRENT)+" debugNUM "+fmt.Sprint(debugNUM), x, y, FONT1DEFAULT, 1, WHITE())
	y += float32(FONT1DEFAULT.charSMLR[0].r.H)
	DtxtXY("LCLICK "+fmt.Sprint(LCLICK)+" RCLICK "+fmt.Sprint(RCLICK), x, y, FONT1DEFAULT, 1, WHITE())
	y += float32(FONT1DEFAULT.charSMLR[0].r.H)
	DtxtXY("MCLICK "+fmt.Sprint(MCLICK)+" clickHoldT "+fmt.Sprint(clickHoldT), x, y, FONT1DEFAULT, 1, WHITE())
	y += float32(FONT1DEFAULT.charSMLR[0].r.H)
	DtxtXY("len(KEYS) "+fmt.Sprint(len(KEYS))+" clickHoldT "+fmt.Sprint(clickHoldT), x, y, FONT1DEFAULT, 1, WHITE())
	y += float32(FONT1DEFAULT.charSMLR[0].r.H)
}

// DRAW
func DrunTime(leadingText string, x, y float32, fon FONT, siz int, c sdl.Color) {
	s := fmt.Sprint(SECONDS)
	if SECONDS < 10 {
		s = "0" + s
	}
	m := fmt.Sprint(MINUTES)
	if MINUTES < 10 {
		m = "0" + m
	}
	h := fmt.Sprint(HOURS)
	if HOURS < 10 {
		h = "0" + h
	}

	DtxtXY(leadingText+" "+h+":"+m+":"+s, x, y, fon, siz, c)
}

// UTILS
func dMSGS() {
	offset := FONT1DEFAULT.smlrH
	y := float32(2)
	for i := range dispErr {
		DtxtCenterScreen(dispErr[i], 1, y, FONT1DEFAULT, 1, ORANGE())
		y += offset
	}
}
func Mmsg(txt string) {
	if len(dispErr) > 0 {
		canadd := true
		for i := range len(dispErr) {
			if txt == dispErr[i] {
				canadd = false
				break
			}
		}
		if canadd {
			dispErr = append(dispErr, txt)
		}
	} else {
		dispErr = append(dispErr, txt)
	}
}
