package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	CNTR                   sdl.FPoint
	FULLSCREEN, OFF        bool
	SCRW, SCRH, WINW, WINH int32
	WINWDEF, WINHDEF       int32 = 1920, 1080
	WIN                    *sdl.Window
	RND                    *sdl.Renderer
	TEX, LOADTEX           *sdl.Texture
	LOADSURF               *sdl.Surface
	DISPLAYMODE            sdl.DisplayMode
	ERR                    error
	ERRLIST                []error
	//SET FPS
	SETFPS     bool
	frameStart time.Time
	frameDelay time.Duration
	TARGETFPS  = int64(60)
)

func EXIT() {
	for i := range ICONSSML {
		ICONSSML[i].tex.Destroy()
	}
	LOADTEX.Destroy()
	RND.Destroy()
	WIN.Destroy()
	OFF = true
}
func B4() {
	if SETFPS {
		frameStart = time.Now()
	}
	RND.SetRenderTarget(TEX)
	RND.SetDrawColor(COLOR2RGBA(BGCOL))
	RND.Clear()

}
func AFTER() {
	if DEBUG {
		DEBUGON()
	}
	dMSGS()

	RND.SetRenderTarget(nil)
	RND.SetDrawColor(COLOR2RGBA(BGCOL))
	RND.Clear()
	RND.Copy(TEX, nil, &sdl.Rect{0, 0, WINW, WINH})
	RND.Present()
	UPD()
	INP()

	if SETFPS {
		elapsed := time.Since(frameStart)
		if remaining := frameDelay - elapsed; remaining > 0 {
			time.Sleep(remaining)
		}
	}
}
func INITIAL(winName string, bgColor sdl.Color) {
	defer WIN.Destroy()
	defer RND.Destroy()
	defer TXSURF.Free()
	defer LOADSURF.Free()
	defer TEX.Destroy()

	STARTTIME = time.Now()
	SECONDTIMER = time.Now()

	if SETFPS {
		frameDelay = time.Second / time.Duration(TARGETFPS)
	}

	BGCOL = bgColor

	//GET SCREEN SIZE
	WIN, ERR = sdl.CreateWindow("", 0, 0, 0, 0, sdl.WINDOW_HIDDEN)
	if ERR != nil {
		mERR(ERR)
	}
	DISPLAYMODE, ERR = sdl.GetCurrentDisplayMode(0)
	if DISPLAYMODE.W > 0 {
		SCRW = DISPLAYMODE.W
	}
	if DISPLAYMODE.H > 0 {
		SCRH = DISPLAYMODE.H
	}
	if FULLSCREEN && SCRW > 0 || SCRW > 0 && SCRW < WINWDEF {
		WINW = SCRW
		WINH = SCRH
	} else {
		WINW = WINWDEF
		WINH = WINHDEF
	}
	WIN.Destroy()
	CNTR = sdl.FPoint{float32(WINW / 2), float32(WINH / 2)}

	//CREATE WINDOW & RENDERER & RENDER TEXTURE
	WIN, ERR = sdl.CreateWindow(winName, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, WINW, WINH, sdl.WINDOW_ALLOW_HIGHDPI)
	if ERR != nil {
		mERR(ERR)
	}
	if SETFPS {
		RND, ERR = sdl.CreateRenderer(WIN, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_TARGETTEXTURE)
	} else {
		RND, ERR = sdl.CreateRenderer(WIN, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_TARGETTEXTURE)
	}
	if ERR != nil {
		mERR(ERR)
	}
	TEX, ERR = RND.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, WINW, WINH)
	if ERR != nil {
		mERR(ERR)
	}

	RND.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	mKEYS()
	mFONTSDEFAULT()
	mIMGS()
	mBUTTONS()

	//EXAMPLES
	mEXAMPLES()

}

func mERR(e error) {
	ERRLIST = append(ERRLIST, e)
}
