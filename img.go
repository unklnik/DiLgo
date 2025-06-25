package main

import (
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

/*

TO DO: MimSheetSeperateImageFiles add file paths in directory automatically

*/

var (
	ICONSSML, ICONSLRG, CURSORS []IM
	BLANKIM                     = IM{}
)

type ANIM struct {
	startIMnum, endIMnum, frames, dFrame int
	ims                                  []IM
	fps, numType                         int
	timer                                time.Time
	flipH, flipV, off                    bool
}
type ANIMMULTI struct {
	anm                    []ANIM
	anmIdle                ANIM
	keys                   []string
	stateLen, stateCurrent int

	idle, keysOn, timerOn, eventSwitch1, eventSwitch2 bool
}
type IM struct {
	r                              sdl.Rect
	tex                            *sdl.Texture
	rCollision, rGrowShrink, rDraw sdl.FRect
	xShrink, speedX, speedY        float32
	pathPoint                      int
	pathCNTR                       sdl.FPoint
	path                           []sdl.FPoint
	pathRecs, ranPosRecs           []sdl.FRect
	rotation, rotationStep         float64
	alpha                          uint8
	ranFlipH, ranFlipV             []bool
	ranAlpha                       []uint8
	ranAngle                       []float64

	onoffGrowShrink, loopBack, rotates, loopScreen, bounceScreen bool
}

// MARK: TOOLS
func DimResize(im IM, rStart sdl.FRect, showBorder bool, cBorder sdl.Color) IM {
	if im.rGrowShrink == BLANKFREC {
		im.rGrowShrink = rStart
	}

	Dim(im, im.rGrowShrink)
	if showBorder {
		DrecLine(im.rGrowShrink, cBorder)
	}
	if CpointRecBorderCollis(MOUSE, im.rGrowShrink, 4) {
		//DcircleFillCenter(MOUSE, 10, ORANGE())
		sdl.ShowCursor(sdl.DISABLE)
		Dim(CURSORS[1], sdl.FRect{MOUSE.X, MOUSE.Y, 16, 16})
	} else {
		sdl.ShowCursor(sdl.ENABLE)
		DrecLine(sdl.FRect{MOUSE.X, MOUSE.Y, 16, 16}, MAGENTA())
	}
	return im
}

// MARK: ISO
func UisoGridAddCubeIM(ig ISOGRID, im IM) ISOGRID {
	if len(ig.cubes) == 0 {
		Mmsg("ERROR: func UisoGridAddCubeIM: len(ig.cubes) == 0 >> First add cubes to grid >> Will not add image")
	} else {
		for i := range ig.cubes {
			ig.cubes[i].im = im
		}
	}
	return ig
}
func UisoGridAddCubeRandomIM(ig ISOGRID, im []IM) ISOGRID {
	if len(ig.cubes) == 0 {
		Mmsg("ERROR: func UisoGridAddCubeIM: len(ig.cubes) == 0 >> First add cubes to grid >> Will not add images")
	} else if len(im) < 2 {
		Mmsg("ERROR: func UisoGridAddCubeIM: len(im) < 2 >> Must be more than 1 image in the slice >> Will not add images")
	} else {
		for i := range ig.cubes {
			ig.cubes[i].im = im[RINT(0, len(im))]
		}
	}
	return ig
}
func UisoGridAddIsoRecSurfaceIM(ig ISOGRID, im IM) ISOGRID {
	for i := range ig.recs {
		ig.recs[i].imSurface = im
	}
	return ig
}
func UisoGridAddScatterIMrandomSize(ig ISOGRID, im []IM, scatterFrequency1to10 int, maxIMrecW float32) ISOGRID {
	if scatterFrequency1to10 < 1 || scatterFrequency1to10 > 10 {
		Mmsg("ERROR: func UisoGridAddScatterIM: scatterFrequency1to10 must be from 1 to 10 (10 is most frequent) >> scatterFrequency1to10 set to 5")
		scatterFrequency1to10 = 5
	}
	for i := range ig.recs {
		if RINT(0, 20) < scatterFrequency1to10 {
			addIM := im[RINT(0, len(im))]
			addIM.rDraw = RESIZEIMRECSCALEHEIGHT(addIM, RF32(maxIMrecW/2, maxIMrecW))
			addIM.rDraw = MOVERECTOCENTER(addIM.rDraw, ig.recs[i].cnt)
			ig.recs[i].imAdditional = append(ig.recs[i].imAdditional, addIM)
		}
	}
	return ig
}

// MARK: DRAW SHEET
func UdrawSheetMotionRandom(ims []IM, minSpeedX, maxSpeedX, minSpeedY, maxSpeedY float32, loopScreen, bounceScreen bool) []IM {
	if loopScreen && bounceScreen {
		Mmsg("ERROR: func UimSheetDrawRecMotionSet: loopScreen & bounceScreen cannot both be true >> loopScreen set to true & bounceScreen set to false")
		loopScreen = true
		bounceScreen = false
	}
	for i := range ims {
		var speedY, speedX float32
		for {
			speedY = RF32(minSpeedY, maxSpeedY)
			if ABS32(speedY) >= ABS32(maxSpeedY)/10 {
				break
			}
		}
		for {
			speedX = RF32(minSpeedX, maxSpeedX)
			if ABS32(speedX) >= ABS32(maxSpeedX)/10 {
				break
			}
		}
		ims[i].loopScreen = loopScreen
		ims[i].bounceScreen = bounceScreen
		ims[i].speedY = speedY
		ims[i].speedX = speedX
	}
	return ims
}
func UdrawSheetMotionSet(ims []IM, speedX, speedY float32, loopScreen, bounceScreen bool) []IM {
	if loopScreen && bounceScreen {
		Mmsg("ERROR: func UimSheetDrawRecMotionSet: loopScreen & bounceScreen cannot both be true >> loopScreen set to true & bounceScreen set to false")
		loopScreen = true
		bounceScreen = false
	}
	for i := range ims {
		ims[i].loopScreen = loopScreen
		ims[i].bounceScreen = bounceScreen
		ims[i].speedY = speedY
		ims[i].speedX = speedX
	}
	return ims
}

func MdrawSheetTiledRec(x, y float32, numTilesW, numTilesH int, tileSize float32, t IM) []IM {
	ims := []IM{}
	ox := x
	a := numTilesW * numTilesH
	c := 0
	for a > 0 {
		i := IM{}
		i.tex = t.tex
		i.r = t.r
		i.rDraw = sdl.FRect{x, y, tileSize, tileSize}
		ims = append(ims, i)
		x += tileSize
		c++
		a--
		if c == numTilesW {
			c = 0
			x = ox
			y += tileSize
		}
	}
	return ims
}
func MdrawSheetRandomPositionsPLUS(ims []IM, minNumIM, maxNumIM, minAlpha, maxAlpha int, minIMrecW, maxIMrecW, drawAreaX, drawAreaY, drawAreaW, drawAreaH float32, minRotation, maxRotation float64, rotates bool, minRotationStep, maxRotationStep float64) []IM {
	if minNumIM >= maxNumIM {
		Mmsg("ERROR: func MimSheetDrawRecRandomPLUS: minNumIM must be smaller (<) than maxNumIM >> minNumIM set to 10 maxNumIM set to 20")
		minNumIM = 10
		maxNumIM = 20
	}
	if minAlpha >= maxAlpha {
		Mmsg("ERROR: func MimSheetDrawRecRandomPLUS: minAlpha must be smaller (<) than maxAlpha >> minAlpha set to 50 maxAlpha set to 150")
		minAlpha = 50
		maxAlpha = 150
	}
	if minIMrecW >= maxIMrecW {
		Mmsg("ERROR: func MimSheetDrawRecRandomPLUS: minIMrecW must be smaller (<) than maxIMrecW >> minIMrecW set to 32 maxIMrecW set to 64")
		minIMrecW = 32
		maxIMrecW = 64
	}
	if drawAreaW <= 0 || drawAreaH <= 0 {
		Mmsg("ERROR: func MimSheetDrawRecRandomPLUS: drawAreaW & drawAreaH must be larger than zero >> drawAreaW set to 1920 drawAreaH set to 1080")
		drawAreaW = 1920
		drawAreaH = 1080
	}
	if minRotation >= maxRotation {
		Mmsg("ERROR: func MimSheetDrawRecRandomPLUS: minRotation must be smaller (<) than maxRotation >> minRotation set to 0 maxRotation set to 360")
		minRotation = 0
		maxRotation = 360
	}
	if minRotationStep >= maxRotationStep {
		Mmsg("ERROR: func MimSheetDrawRecRandomPLUS: minRotationStep must be smaller (<) than maxRotationStep >> minRotationStep set to 1 maxRotationStep set to 4")
		minRotationStep = 1
		maxRotationStep = 4
	}
	im2 := []IM{}
	num := RINT(minNumIM, maxNumIM)
	for range num {
		w := RF32(minIMrecW, maxIMrecW)
		i := IM{}
		i.rotates = rotates
		for {
			i.rotationStep = RF64(minRotationStep, maxRotationStep)
			if ABS64(i.rotationStep) > ABS64(maxRotationStep)/10 {
				break
			}
		}
		i.rDraw = sdl.FRect{RF32(drawAreaX, drawAreaW-w), RF32(drawAreaY, drawAreaH-w), w, w}
		i.rotation = RF64(minRotation, maxRotation)
		choose := ims[RINT(0, len(ims))]
		i.tex = choose.tex
		i.r = choose.r
		i.alpha = RUINT8(minAlpha, maxAlpha)
		im2 = append(im2, i)
	}

	return im2
}

// MARK: IM SHEET
func MimSheetSeperateImageFilesFromDirectoryJPGandPNG(directoryPath string) []IM {
	paths := FimageFilesJPGandPNG(directoryPath)
	ims := MimSheetSeperateImageFiles(paths)
	return ims
}
func MimSheetMultiRowWidthHeight(filePath string, xStart, yStart, wIM, hIM, wTotal, hTotal int32) []IM {
	var ims []IM
	ox := xStart
	im := IM{}
	im.tex = mTEX(filePath)
	for yStart < hTotal {
		im.r = sdl.Rect{xStart, yStart, wIM, hIM}
		ims = append(ims, im)
		xStart += wIM
		if xStart >= wTotal {
			xStart = ox
			yStart += hIM
		}
	}
	return ims
}

/*
TO DO: MimSheetSeperateImageFiles add file paths in directory automatically
*/
func MimSheetSeperateImageFilesRecSize(filePaths []string, r sdl.Rect) []IM {
	if len(filePaths) <= 1 {
		Mmsg("ERROR: func MimSheetSeperateImageFiles: len(filePaths) must be longer than 1 >> will return empty []IM")
	}
	var ims []IM
	for i := range filePaths {
		im := IM{}
		im.tex = mTEX(filePaths[i])
		im.r = r
		ims = append(ims, im)
	}
	return ims
}
func MimSheetSeperateImageFiles(filePaths []string) []IM {
	if len(filePaths) <= 1 {
		Mmsg("ERROR: func MimSheetSeperateImageFiles: len(filePaths) must be longer than 1 >> will return empty []IM")
	}
	var ims []IM
	for i := range filePaths {
		im := IM{}
		im.tex = mTEX(filePaths[i])
		_, _, w, h, _ := im.tex.Query()
		im.r = sdl.Rect{0, 0, w, h}
		ims = append(ims, im)
	}
	return ims
}
func MimSheetCollisionRecs(ims []IM, xOffset, yOffset, w, h float32) []IM {
	for i := range ims {
		ims[i].rCollision = sdl.FRect{float32(ims[i].r.X) + xOffset, float32(ims[i].r.Y) + yOffset, w, h}
	}
	return ims
}
func MimSheetFramesHorizontal(filePath string, wIM, hIM, framesNum, xStart, y int32) []IM {
	var ims []IM
	im := IM{}
	im.tex = mTEX(filePath)
	for range framesNum {
		im.r = sdl.Rect{xStart, y, wIM, hIM}
		ims = append(ims, im)
		xStart += wIM
	}
	return ims
}
func MimSheetWidthHorizontal(filePath string, wIM, hIM, wTotal, xStart, y int32) []IM {
	var ims []IM
	im := IM{}
	im.tex = mTEX(filePath)
	x := xStart
	for x <= xStart+wTotal-wIM {
		im.r = sdl.Rect{x, y, wIM, hIM}
		ims = append(ims, im)
		x += wIM
	}
	return ims
}

// MARK: ANIMS
func ManimMultiEventSwitch2(a1, a2 ANIM) ANIMMULTI {
	am := ANIMMULTI{}
	am.anm = append(am.anm, a1, a2)
	return am
}
func ManimMultiEventSwitch3(a1, a2, a3 ANIM) ANIMMULTI {
	am := ANIMMULTI{}
	am.anm = append(am.anm, a1, a2, a3)
	return am
}
func ManimMultiKeyStatesIdle(aKeys []ANIM, aIdle ANIM, keys []string) ANIMMULTI {
	am := ANIMMULTI{}
	if len(aKeys) != len(keys) {
		Mmsg("ERROR: func ManimMultiKeyStates: len(a) must equal len(keys) >> key determines which animation plays >> will return empty ANIMMULTI")
	} else {
		am = ManimMulti(aKeys)
		am.stateLen = len(aKeys) - 1
		am.anm = aKeys
		am.keys = keys
		am.keysOn = true
		am.anmIdle = aIdle
		am.idle = true
	}
	return am
}
func ManimMultiKeyStates(a []ANIM, keys []string) ANIMMULTI {
	am := ANIMMULTI{}
	if len(a) != len(keys) {
		Mmsg("ERROR: func ManimMultiKeyStates: len(a) must equal len(keys) >> key determines which animation plays >> will return empty ANIMMULTI")
	} else {
		am = ManimMulti(a)
		am.stateLen = len(a) - 1
		am.anm = a
		am.keys = keys
		am.keysOn = true
	}
	return am
}
func ManimMulti(a []ANIM) ANIMMULTI {
	am := ANIMMULTI{}
	am.stateLen = len(a) - 1
	am.anm = a
	return am
}
func UanimFlip(a ANIM, flipH, flipV bool) ANIM {
	a.flipH = flipH
	a.flipV = flipV
	return a
}
func ManimImSheet(ims []IM, fps1to100 int) ANIM {
	if fps1to100 < 1 || fps1to100 > 100 {
		Mmsg("ERROR: func ManimImSheet: fps1to100 must be from 1 to 100 >> 1 = Slowest / 100 = Fastest >> Set to 50")
		fps1to100 = 50
	}
	a := ANIM{}
	a.numType = 0 //IM SHEET FULL
	a.ims = ims
	a.frames = len(ims) - 1
	a.fps = fps1to100
	return a
}

// MARK: UTILS
func mIMGS() {
	//SMALL ICONS
	ICONSSML = MimSheetWidthHorizontal("img/icons_sml.png", 24, 24, 648, 0, 0)
	//CURSOR
	CURSORS = MimSheetSeperateImageFilesFromDirectoryJPGandPNG("img/cursor/")
}
func mTEX(filePath string) *sdl.Texture {
	LOADSURF, ERR = img.Load(filePath)
	if ERR != nil {
		mERR(ERR)
	}
	defer LOADSURF.Free()
	LOADTEX, ERR = RND.CreateTextureFromSurface(LOADSURF)
	return LOADTEX
}
