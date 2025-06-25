package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (

	//MENUS
	menuXmouseW, menuYmouseH float32
	//TIMERS
	frameCount, FRAMES, FPSCURRENT uint64
	lastTime                       uint64
	SECONDS, MINUTES, HOURS        int
	TIMESEC, TIMEMIN, TIMEHOUR     int
	STARTTIME, SECONDTIMER         time.Time
)

func UPD() {

	TIMERS()
	INP()

}

// TIMERS
func TIMERS() {

	//RUN TIME
	if time.Since(SECONDTIMER) >= time.Second {
		SECONDTIMER = time.Now()
		SECONDS++
	}
	if SECONDS == 60 {
		MINUTES++
		SECONDS = 0
	}
	if MINUTES == 60 {
		HOURS++
		MINUTES = 0
	}

	//FRAMES
	FRAMES++
	//FPS
	frameCount++
	currentTime := sdl.GetTicks64()
	elapsedTime := currentTime - lastTime
	if elapsedTime >= 1000 {
		FPSCURRENT = frameCount * 1000 / elapsedTime
		frameCount = 0
		lastTime = currentTime
	}

	//KEYS
	if KEYDOWNTIMER > 0 {
		KEYDOWNTIMER--
	}

}
