package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	mouseBlinkA = uint8(255)
	BLANKREC    = sdl.Rect{}
	BLANKFREC   = sdl.FRect{}
)

// MARK: ISO
func DisoGridzIndexPlayerAnimMulti(ig ISOGRID, surfaceIM, additionalIM bool, player ANIMMULTI) ANIMMULTI {

	if surfaceIM {
		DisoGridIsoRecSurfaceIM(ig)
	}
	if additionalIM {
		zrecs := ZINDEXSORTISOREC(ig.recs)
		for i := range zrecs {
			if len(zrecs[i].imAdditional) > 0 {
				for j := range zrecs[i].imAdditional {
					Dim(zrecs[i].imAdditional[j], zrecs[i].imAdditional[j].rDraw)
					//if
				}
			}
		}
	}

	return player
}
func DisoGridAdditionalIM(ig ISOGRID) {
	for i := range ig.recs {
		if len(ig.recs[i].imAdditional) > 0 {
			for j := range ig.recs[i].imAdditional {
				Dim(ig.recs[i].imAdditional[j], ig.recs[i].imAdditional[j].rDraw)
			}
		}
	}
}
func DisoGridIsoRecSurfaceIM(ig ISOGRID) {
	for i := range ig.recs {
		if ig.recs[i].imSurface.r.W == ig.recs[i].imSurface.r.H {
			r := sdl.FRect{ig.recs[i].cnt.X - ig.recs[i].wSide, ig.recs[i].cnt.Y - ig.recs[i].wSide, ig.recs[i].wSide * 2, ig.recs[i].wSide * 2}
			Dim(ig.recs[i].imSurface, r)
		} else if ig.recs[i].imSurface.r.W == ig.recs[i].imSurface.r.H*2 {
			r := sdl.FRect{ig.recs[i].cnt.X - ig.recs[i].wSide, ig.recs[i].cnt.Y - ig.recs[i].wSide/2, ig.recs[i].wSide * 2, ig.recs[i].wSide}
			Dim(ig.recs[i].imSurface, r)
		}
	}
}
func DisoGridCubesIM(ig ISOGRID) {
	c := ig.cubes
	for i := len(c) - 1; i > -1; i-- {
		w := c[i].wSide * 2
		x := c[i].frontCorner.X - ig.cubes[i].wSide
		y := c[i].frontCorner.Y - w
		Dim(c[i].im, sdl.FRect{x, y, w, w})
	}
}
func DisoGridCubes(ig ISOGRID) {
	c := ZINDEXSORTCUBE(ig.cubes)
	for i := range c {
		DcubeLineFill(c[i])
	}
}
func DisoGridCubesShadow(ig ISOGRID, cShadow sdl.Color, topDarker, leftRightDarker bool, minShadow0to255, maxShadow0to255 uint8) {
	c := ZINDEXSORTCUBE(ig.cubes)
	for i := range c {
		DCubeShadow(c[i], cShadow, topDarker, leftRightDarker, minShadow0to255, maxShadow0to255)
	}
}
func DisoGridCubesShadowLine(ig ISOGRID, cShadow sdl.Color, topDarker, leftRightDarker bool, minShadow0to255, maxShadow0to255 uint8) {
	c := ZINDEXSORTCUBE(ig.cubes)
	for i := range c {
		DCubeShadowLine(c[i], cShadow, topDarker, leftRightDarker, minShadow0to255, maxShadow0to255)
	}
}
func DmouseIsoGrid(ig ISOGRID, c sdl.Color, blink bool) {
	if blink {
		mouseBlinkA -= 7
		c.A = mouseBlinkA
	}
	for i := range ig.recs {
		if CtriPoint(MOUSE, ig.recs[i].tri[0]) || CtriPoint(MOUSE, ig.recs[i].tri[1]) {
			DisoRecFillColor(ig.recs[i], c)
			break
		}
	}
}
func DisoGridLineFill(ig ISOGRID) {
	for i := range ig.recs {
		DisoRecFill(ig.recs[i])
		DisoRecLine(ig.recs[i])
		if DEBUG {
			DtxtCenterPoint("zi "+fmt.Sprint(ig.recs[i].zindex), FONT1DEFAULT, 1, WHITE(), ig.recs[i].cnt)
			DcircleFillCenter(ig.cnt, 4, MAGENTA())
		}
	}

}
func DisoRecLineFill(r ISOREC) {
	DtriFill(r.tri[0])
	DtriFill(r.tri[1])
	DisoRecLine(r)
}
func DisoRecFill(r ISOREC) {
	DtriFill(r.tri[0])
	DtriFill(r.tri[1])
}
func DisoRecFillShadow(r ISOREC, c sdl.Color, alpha uint8) {
	DtriFill(r.tri[0])
	DtriFill(r.tri[1])
	DtriShadow(r.tri[0], c, alpha)
	DtriShadow(r.tri[1], c, alpha)
}
func DisoRecFillColor(r ISOREC, c sdl.Color) {
	t1 := r.tri[0]
	t2 := r.tri[1]
	t1.c = c
	t2.c = c
	DtriFill(t1)
	DtriFill(t2)
}
func DisoRecLine(r ISOREC) {
	DlinePointsCloseEnd(r.vert, r.cLine)
}

// MARK: ANIM
func DanimMultiEventSwitch(a ANIMMULTI, r sdl.FRect) ANIMMULTI {
	if len(a.anm) == 2 {
		if a.eventSwitch1 {
			a.anm[1] = DanimRecLoop(a.anm[1], r)
		} else {
			a.anm[0] = DanimRecLoop(a.anm[0], r)
		}
	} else if len(a.anm) == 3 {
		if a.eventSwitch2 {
			a.anm[2] = DanimRecLoop(a.anm[2], r)
		} else if a.eventSwitch1 {
			a.anm[1] = DanimRecLoop(a.anm[1], r)
		} else {
			a.anm[0] = DanimRecLoop(a.anm[0], r)
		}
	}
	return a
}
func DanimMultiRecLoop(a ANIMMULTI, r sdl.FRect) ANIMMULTI {

	for i := range a.keys {
		if KEYPRESS(a.keys[i]) {
			a.stateCurrent = i
		}
	}
	if a.idle {
		if a.anmIdle.flipH {
			DimFlipHorizontal(a.anmIdle.ims[a.anmIdle.dFrame], r)
		} else if a.anm[a.stateCurrent].flipV {
			DimFlipVertical(a.anmIdle.ims[a.anmIdle.dFrame], r)
		} else {
			Dim(a.anmIdle.ims[a.anmIdle.dFrame], r)
		}
		a.anmIdle = dAnimFrame(a.anmIdle)
	} else {
		switch a.anm[a.stateCurrent].numType {
		case 0: //IM SHEET FULL
			if a.anm[a.stateCurrent].flipH {
				DimFlipHorizontal(a.anm[a.stateCurrent].ims[a.anm[a.stateCurrent].dFrame], r)
			} else if a.anm[a.stateCurrent].flipV {
				DimFlipVertical(a.anm[a.stateCurrent].ims[a.anm[a.stateCurrent].dFrame], r)
			} else {
				Dim(a.anm[a.stateCurrent].ims[a.anm[a.stateCurrent].dFrame], r)
			}
			a.anm[a.stateCurrent] = dAnimFrame(a.anm[a.stateCurrent])
		}
	}
	return a
}
func dAnimFrame(a ANIM) ANIM {
	if a.timer.IsZero() {
		a.timer = time.Now()
	}
	if time.Since(a.timer) >= time.Second/time.Duration(a.fps) {
		a.dFrame++
		if a.dFrame > a.frames {
			a.dFrame = 0
		}
		a.timer = time.Now()
	}
	return a
}
func dAnimFrameOnce(a ANIM, clearEnd bool) ANIM {
	if a.timer.IsZero() {
		a.timer = time.Now()
	}
	if time.Since(a.timer) >= time.Second/time.Duration(a.fps) {
		a.dFrame++
		if a.dFrame > a.frames {
			a.dFrame = a.frames
			if clearEnd {
				a.off = true
			}
		}
		a.timer = time.Now()
	}
	return a
}
func DanimRecLoop(a ANIM, r sdl.FRect) ANIM {
	switch a.numType {
	case 0: //IM SHEET FULL
		if a.flipH {
			DimFlipHorizontal(a.ims[a.dFrame], r)
		} else if a.flipV {
			DimFlipVertical(a.ims[a.dFrame], r)
		} else {
			Dim(a.ims[a.dFrame], r)
		}
		a = dAnimFrame(a)
	}
	return a
}
func DanimRecOnce(a ANIM, r sdl.FRect, clearEnd bool) ANIM {
	if !a.off {
		switch a.numType {
		case 0: //IM SHEET FULL
			if a.flipH && a.flipV {
				DimFlipHorizontalVertical(a.ims[a.dFrame], r)
			} else if a.flipH && !a.flipV {
				DimFlipHorizontal(a.ims[a.dFrame], r)
			} else if !a.flipH && a.flipV {
				DimFlipVertical(a.ims[a.dFrame], r)
			} else {
				Dim(a.ims[a.dFrame], r)
			}
			a = dAnimFrameOnce(a, clearEnd)
		}
	}
	return a
}
func DanimXYzoom(a ANIM, x, y, zoom float32) ANIM {
	switch a.numType {
	case 0: //IM SHEET FULL
		r := sdl.FRect{x, y, float32(a.ims[a.dFrame].r.W) * zoom, float32(a.ims[a.dFrame].r.H) * zoom}
		if a.flipH && a.flipV {
			DimFlipHorizontalVertical(a.ims[a.dFrame], r)
		} else if a.flipH && !a.flipV {
			DimFlipHorizontal(a.ims[a.dFrame], r)
		} else if !a.flipH && a.flipV {
			DimFlipVertical(a.ims[a.dFrame], r)
		} else {
			Dim(a.ims[a.dFrame], r)
		}
		a = dAnimFrame(a)
	}
	return a
}
func DanimXYzoomFlipHorizontalVerticle(a ANIM, x, y, zoom float32) ANIM {
	switch a.numType {
	case 0: //IM SHEET FULL
		r := sdl.FRect{x, y, float32(a.ims[a.dFrame].r.W) * zoom, float32(a.ims[a.dFrame].r.H) * zoom}
		DimFlipHorizontalVertical(a.ims[a.dFrame], r)
		a = dAnimFrame(a)
	}
	return a
}
func DanimXYzoomFlipHorizontal(a ANIM, x, y, zoom float32) ANIM {
	switch a.numType {
	case 0: //IM SHEET FULL
		r := sdl.FRect{x, y, float32(a.ims[a.dFrame].r.W) * zoom, float32(a.ims[a.dFrame].r.H) * zoom}
		DimFlipHorizontal(a.ims[a.dFrame], r)
		a = dAnimFrame(a)
	}
	return a
}
func DanimXYzoomFlipVertical(a ANIM, x, y, zoom float32) ANIM {
	switch a.numType {
	case 0: //IM SHEET FULL
		r := sdl.FRect{x, y, float32(a.ims[a.dFrame].r.W) * zoom, float32(a.ims[a.dFrame].r.H) * zoom}
		DimFlipVertical(a.ims[a.dFrame], r)
		a = dAnimFrame(a)
	}
	return a
}

// MARK: IMAGES
func Dim(im IM, r sdl.FRect) {
	RND.CopyF(im.tex, &im.r, &r)
}
func DimDrawRec(im IM) {
	if im.alpha != 0 {
		im.tex = TEXALPHA(im.tex, im.alpha)
	}
	RND.CopyExF(im.tex, &im.r, &im.rDraw, im.rotation, ORIGIN(im.rDraw), sdl.FLIP_NONE)
}
func DdrawSheet(ims []IM) {
	for i := range ims {
		DimDrawRec(ims[i])
		if ims[i].rotates {
			ims[i].rotation += ims[i].rotationStep
		}
		if ims[i].speedX != 0 {
			ims[i].rDraw.X += ims[i].speedX
			if ims[i].loopScreen {
				if ims[i].rDraw.X > float32(WINW) {
					ims[i].rDraw.X = -ims[i].rDraw.W
				} else if ims[i].rDraw.X < -(ims[i].rDraw.W + 1) {
					ims[i].rDraw.X = float32(WINW)
				}
			} else if ims[i].bounceScreen {
				if ims[i].rDraw.X+ims[i].rDraw.W >= float32(WINW) || ims[i].rDraw.X <= 0 {
					ims[i].speedX *= -1
				}
			}
		}
		if ims[i].speedY != 0 {
			ims[i].rDraw.Y += ims[i].speedY
			if ims[i].loopScreen {
				if ims[i].rDraw.Y > float32(WINH) {
					ims[i].rDraw.Y = -ims[i].rDraw.H
				} else if ims[i].rDraw.Y < -(ims[i].rDraw.H + 1) {
					ims[i].rDraw.Y = float32(WINH)
				}
			} else if ims[i].bounceScreen {
				if ims[i].rDraw.Y+ims[i].rDraw.H >= float32(WINH) || ims[i].rDraw.Y <= 0 {
					ims[i].speedY *= -1
				}
			}
		}
	}
}

func DimRandomWindowSizeRandom(im IM, minRecW, maxRecW float32, minNumIM, maxNumIM int, randomFlipHorizontal, randomFlipVertical, randomAngle, randomAlpha bool) IM {
	if minRecW < 0 || maxRecW < 0 {
		if minRecW < 0 {
			Mmsg("ERROR: func DimRandomWindowSizeRandom: minRecW must be larger or equal (>=) to 0 >> Set to 0")
			minRecW = 0
		}
		if maxRecW < 0 {
			Mmsg("ERROR: func DimRandomWindowSizeRandom: maxRecW must be larger (>) than 0 >> Set to 64")
			maxRecW = 64
		}

	}
	if minRecW >= maxRecW {
		Mmsg("ERROR: func DimRandomWindowSizeRandom: minRecW must be smaller (<) than maxRecW >> Set minRecW to 32 maxRecW to 64")
		minRecW = 32
		maxRecW = 64
	}

	if minNumIM < 1 {
		Mmsg("ERROR: func DimRandomWindowSizeRandom: minNumIM must be larger than or equal to (>=) 1 >> Set to 1")
		minNumIM = 1
	}
	if maxNumIM <= minNumIM {
		Mmsg("ERROR: func DimRandomWindowSizeRandom: maxNumIM must be larger than (>) minNumIM >> Set to minNumIM + 10")
		maxNumIM = minNumIM + 10
	}
	if len(im.ranPosRecs) == 0 {
		numIM := RINT(minNumIM, maxNumIM)
		for range numIM {
			recW := RF32(minRecW, maxRecW)
			im.ranPosRecs = append(im.ranPosRecs, sdl.FRect{RF32(0, float32(WINW)-recW), RF32(0, float32(WINH)-recW), recW, recW})
		}
	}
	if randomFlipHorizontal && len(im.ranFlipH) == 0 {
		for range im.ranPosRecs {
			im.ranFlipH = append(im.ranFlipH, FLIPCOIN())
		}
	}
	if randomFlipVertical && len(im.ranFlipV) == 0 {
		for range im.ranPosRecs {
			im.ranFlipV = append(im.ranFlipV, FLIPCOIN())
		}
	}
	if randomAngle && len(im.ranAngle) == 0 {
		for range im.ranPosRecs {
			im.ranAngle = append(im.ranAngle, RF64(0, 360))
		}
	}
	if randomAlpha && len(im.ranAlpha) == 0 {
		for range im.ranPosRecs {
			im.ranAlpha = append(im.ranAlpha, RUINT8(0, 256))
		}
	}
	for i := range im.ranPosRecs {
		if len(im.ranFlipH) > 0 && len(im.ranFlipV) > 0 && len(im.ranAngle) > 0 && len(im.ranAlpha) > 0 {
			im.tex = TEXALPHA(im.tex, im.ranAlpha[i])
			if im.ranFlipH[i] && im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL|sdl.FLIP_VERTICAL)
			} else if im.ranFlipH[i] && !im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL)
			} else if !im.ranFlipH[i] && im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_VERTICAL)
			}
			im.tex = TEXALPHAREVERT(im.tex)
		} else if len(im.ranFlipH) > 0 && len(im.ranFlipV) > 0 && len(im.ranAngle) > 0 && len(im.ranAlpha) == 0 {
			if im.ranFlipH[i] && im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL|sdl.FLIP_VERTICAL)
			} else if im.ranFlipH[i] && !im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL)
			} else if !im.ranFlipH[i] && im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_VERTICAL)
			}
		} else if len(im.ranFlipH) > 0 && len(im.ranFlipV) > 0 && len(im.ranAngle) == 0 && len(im.ranAlpha) == 0 {
			if im.ranFlipH[i] && im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], 0, ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL|sdl.FLIP_VERTICAL)
			} else if im.ranFlipH[i] && !im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], 0, ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL)
			} else if !im.ranFlipH[i] && im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], 0, ORIGIN(im.ranPosRecs[i]), sdl.FLIP_VERTICAL)
			}
		} else if len(im.ranFlipH) > 0 && len(im.ranFlipV) > 0 && len(im.ranAngle) == 0 && len(im.ranAlpha) > 0 {
			im.tex = TEXALPHA(im.tex, im.ranAlpha[i])
			if im.ranFlipH[i] && im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], 0, ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL|sdl.FLIP_VERTICAL)
			} else if im.ranFlipH[i] && !im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], 0, ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL)
			} else if !im.ranFlipH[i] && im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], 0, ORIGIN(im.ranPosRecs[i]), sdl.FLIP_VERTICAL)
			}
			im.tex = TEXALPHAREVERT(im.tex)
		} else if len(im.ranFlipH) > 0 && len(im.ranFlipV) == 0 && len(im.ranAngle) > 0 && len(im.ranAlpha) > 0 {
			im.tex = TEXALPHA(im.tex, im.ranAlpha[i])
			if im.ranFlipH[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL)
			} else if !im.ranFlipH[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_NONE)
			}
			im.tex = TEXALPHAREVERT(im.tex)
		} else if len(im.ranFlipH) == 0 && len(im.ranFlipV) >= 0 && len(im.ranAngle) > 0 && len(im.ranAlpha) > 0 {
			im.tex = TEXALPHA(im.tex, im.ranAlpha[i])
			if im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_VERTICAL)
			} else if !im.ranFlipV[i] {
				RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_NONE)
			}
			im.tex = TEXALPHAREVERT(im.tex)
		} else if len(im.ranFlipH) == 0 && len(im.ranFlipV) == 0 && len(im.ranAngle) > 0 && len(im.ranAlpha) > 0 {
			im.tex = TEXALPHA(im.tex, im.ranAlpha[i])
			RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_NONE)
			im.tex = TEXALPHAREVERT(im.tex)
		} else if len(im.ranFlipH) == 0 && len(im.ranFlipV) == 0 && len(im.ranAngle) == 0 && len(im.ranAlpha) > 0 {
			im.tex = TEXALPHA(im.tex, im.ranAlpha[i])
			RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], 0, ORIGIN(im.ranPosRecs[i]), sdl.FLIP_NONE)
			im.tex = TEXALPHAREVERT(im.tex)
		} else if len(im.ranFlipH) == 0 && len(im.ranFlipV) == 0 && len(im.ranAngle) > 0 && len(im.ranAlpha) == 0 {
			RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_NONE)
		} else if len(im.ranFlipH) > 0 && len(im.ranFlipV) == 0 && len(im.ranAngle) > 0 && len(im.ranAlpha) == 0 {
			RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_HORIZONTAL)
		} else if len(im.ranFlipH) == 0 && len(im.ranFlipV) > 0 && len(im.ranAngle) > 0 && len(im.ranAlpha) == 0 {
			RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], im.ranAngle[i], ORIGIN(im.ranPosRecs[i]), sdl.FLIP_VERTICAL)
		} else if len(im.ranFlipH) == 0 && len(im.ranFlipV) > 0 && len(im.ranAngle) == 0 && len(im.ranAlpha) > 0 {
			im.tex = TEXALPHA(im.tex, im.ranAlpha[i])
			RND.CopyExF(im.tex, &im.r, &im.ranPosRecs[i], 0, ORIGIN(im.ranPosRecs[i]), sdl.FLIP_VERTICAL)
			im.tex = TEXALPHAREVERT(im.tex)
		}

	}

	return im
}
func DimRandomWindowSizeSet(im IM, recW, recH float32, numIM int) IM {
	if numIM < 1 {
		Mmsg("ERROR: func DimRandomWindowSize: numIM ust be larger than 0 >> Set to 10")
		numIM = 10

	}
	if len(im.ranPosRecs) == 0 {
		for range numIM {
			im.ranPosRecs = append(im.ranPosRecs, sdl.FRect{RF32(0, float32(WINW)-recW), RF32(0, float32(WINH)-recH), recW, recH})
		}
	}
	for i := range im.ranPosRecs {
		Dim(im, im.ranPosRecs[i])
	}
	return im
}
func DimColor(im IM, r sdl.FRect, c sdl.Color) {
	im.tex = TEXCOL(im.tex, c)
	RND.CopyF(im.tex, &im.r, &r)
}
func DimColorAlpha(im IM, r sdl.FRect, c sdl.Color, alpha uint8) {
	if alpha < 0 || alpha > 255 {
		Mmsg("ERROR: func DimAlpha: alpha must be from 0 to 255 >> Set to 255")
		alpha = 255
	}
	im.tex = TEXALPHA(im.tex, alpha)
	im.tex = TEXCOL(im.tex, c)
	RND.CopyF(im.tex, &im.r, &r)
}
func DimAlpha(im IM, r sdl.FRect, alpha uint8) {
	if alpha < 0 || alpha > 255 {
		Mmsg("ERROR: func DimAlpha: alpha must be from 0 to 255 >> Set to 255")
		alpha = 255
	}
	im.tex = TEXALPHA(im.tex, alpha)
	RND.CopyF(im.tex, &im.r, &r)
}
func DimPLUS(im IM, r sdl.FRect, flipHorizontal, flipVertical, rotates bool, rotationStep, rotationAngle float64, alpha uint8) IM {
	if alpha < 0 || alpha > 255 {
		Mmsg("ERROR: func DimPLUS: alpha must be from 0 to 255 >> Set to 255")
		alpha = 255
	}
	if rotates && rotationStep == 0 {
		Mmsg("ERROR: func DimPLUS: rotation is true & rotationStep = 0 >> Will not rotate >> rotationStep set to 1")
		rotationStep = 1
	}
	if alpha != 255 {
		im.tex = TEXALPHA(im.tex, alpha)
	}
	if rotates && im.rotation == 0 {
		im.rotation = rotationAngle
	}
	if flipHorizontal && flipVertical {
		RND.CopyExF(im.tex, &im.r, &r, im.rotation, ORIGIN(r), sdl.FLIP_HORIZONTAL|sdl.FLIP_VERTICAL)
	} else if !flipHorizontal && flipVertical {
		RND.CopyExF(im.tex, &im.r, &r, im.rotation, ORIGIN(r), sdl.FLIP_VERTICAL)
	} else if flipHorizontal && !flipVertical {
		RND.CopyExF(im.tex, &im.r, &r, im.rotation, ORIGIN(r), sdl.FLIP_HORIZONTAL)
	} else {
		RND.CopyExF(im.tex, &im.r, &r, im.rotation, ORIGIN(r), sdl.FLIP_NONE)
	}
	if rotates {
		im.rotation += rotationStep
	}
	return im
}
func DimPathPLUS(im IM, recW, recH, speed float32, path []sdl.FPoint, loopBackwards, loopRound bool) IM {
	if len(path) < 2 {
		Mmsg("ERROR: func DimPath: len(path) must be larger than 1 >> Will not draw to screen")
	} else {
		if loopBackwards && loopRound {
			Mmsg("ERROR: func DimPath: loopBackwards either/or loopRound >> Both cannot be true >> Set to loopRound")
			loopBackwards = false
		}
		if len(im.pathRecs) == 0 {
			if loopRound {
				path = append(path, path[0])
			}
			im.path = path
			im.pathCNTR = im.path[0]
			im.pathPoint = 1
			for i := range im.path {
				im.pathRecs = append(im.pathRecs, MrecCenter(im.path[i], 2, 2))
			}
		}
	}

	Dim(im, MrecCenter(im.pathCNTR, recW, recH))

	x, y := SPEEDXY2POINTS(im.pathCNTR, im.path[im.pathPoint], speed)
	im.pathCNTR.X += x
	im.pathCNTR.Y += y
	if CrecPoint(im.pathCNTR, im.pathRecs[im.pathPoint]) {
		if im.loopBack {
			im.pathPoint--
			if im.pathPoint == 0 {
				im.loopBack = false
			}
		} else {
			im.pathPoint++
		}
		if im.pathPoint >= len(im.path)-1 {
			if loopBackwards || loopRound {
				if loopRound {
					im.pathPoint = 0
				} else if loopBackwards {
					im.loopBack = true
				}
			} else {
				im.pathPoint = len(im.path) - 1
			}
		}
	}

	return im
}

func DimPath(im IM, recW, recH, speed float32, path []sdl.FPoint, loopBackwards, loopRound bool) IM {
	if len(path) < 2 {
		Mmsg("ERROR: func DimPath: len(path) must be larger than 1 >> Will not draw to screen")
	} else {
		if loopBackwards && loopRound {
			Mmsg("ERROR: func DimPath: loopBackwards either/or loopRound >> Both cannot be true >> Set to loopRound")
			loopBackwards = false
		}
		if len(im.pathRecs) == 0 {
			if loopRound {
				path = append(path, path[0])
			}
			im.path = path
			im.pathCNTR = im.path[0]
			im.pathPoint = 1
			for i := range im.path {
				im.pathRecs = append(im.pathRecs, MrecCenter(im.path[i], 2, 2))
			}
		}
	}

	Dim(im, MrecCenter(im.pathCNTR, recW, recH))

	x, y := SPEEDXY2POINTS(im.pathCNTR, im.path[im.pathPoint], speed)
	im.pathCNTR.X += x
	im.pathCNTR.Y += y
	if CrecPoint(im.pathCNTR, im.pathRecs[im.pathPoint]) {
		if im.loopBack {
			im.pathPoint--
			if im.pathPoint == 0 {
				im.loopBack = false
			}
		} else {
			im.pathPoint++
		}
		if im.pathPoint >= len(im.path)-1 {
			if loopBackwards || loopRound {
				if loopRound {
					im.pathPoint = 0
				} else if loopBackwards {
					im.loopBack = true
				}
			} else {
				im.pathPoint = len(im.path) - 1
			}
		}
	}

	return im
}

func DimRotate(im IM, r sdl.FRect, angle float64) {
	RND.CopyExF(im.tex, &im.r, &r, angle, ORIGIN(r), sdl.FLIP_NONE)
}
func DimTileRec(im IM, numW, numH int, r sdl.FRect) {
	ox := r.X
	x := ox
	y := r.Y
	w := r.W / float32(numW)
	h := r.W / float32(numH)
	for y < r.Y+r.H {
		Dim(im, sdl.FRect{x, y, w, h})
		x += w
		if x >= r.X+r.W {
			x = ox
			y += h
		}
	}
}
func DimGrowPLUS(im IM, r sdl.FRect, stepIncrease, maxIncrease float32, angle float64, loop, flipH, flipV bool, alpha uint8) IM {
	if stepIncrease > maxIncrease {
		Mmsg("ERROR: func DimGrow: stepIncrease must be smaller (<) than maxIncrease >> Will not draw to screen")
	} else {
		if alpha > 255 || alpha < 0 {
			Mmsg("ERROR: func DimGrow: alpha must be from 0 to 255 >> Set to 255")
			alpha = 255
		}
		if im.rGrowShrink == BLANKFREC {
			im.rGrowShrink = r
		}
		if im.onoffGrowShrink {
			if loop {
				im.rGrowShrink.X += stepIncrease
				im.rGrowShrink.Y += stepIncrease
				im.rGrowShrink.W -= stepIncrease * 2
				im.rGrowShrink.H -= stepIncrease * 2
				if im.rGrowShrink.X >= r.X {
					im.onoffGrowShrink = false
				}
			}
		} else {
			im.rGrowShrink.X -= stepIncrease
			im.rGrowShrink.Y -= stepIncrease
			im.rGrowShrink.W += stepIncrease * 2
			im.rGrowShrink.H += stepIncrease * 2
			if ABSDIFF(r.X, im.rGrowShrink.X) >= maxIncrease {
				im.onoffGrowShrink = true
			}
		}

		im.tex = TEXALPHA(im.tex, alpha)

		if flipH && flipV {
			RND.CopyExF(im.tex, &im.r, &im.rGrowShrink, angle, ORIGIN(im.rGrowShrink), sdl.FLIP_HORIZONTAL|sdl.FLIP_VERTICAL)
		} else if flipH && !flipV {
			RND.CopyExF(im.tex, &im.r, &im.rGrowShrink, angle, ORIGIN(im.rGrowShrink), sdl.FLIP_HORIZONTAL)
		} else if !flipH && flipV {
			RND.CopyExF(im.tex, &im.r, &im.rGrowShrink, angle, ORIGIN(im.rGrowShrink), sdl.FLIP_VERTICAL)
		} else {
			RND.CopyExF(im.tex, &im.r, &im.rGrowShrink, angle, ORIGIN(im.rGrowShrink), sdl.FLIP_NONE)
		}

		im.tex = TEXALPHAREVERT(im.tex)

	}
	return im
}

func DimGrow(im IM, r sdl.FRect, stepIncrease, maxIncrease float32, loop bool) IM {
	if stepIncrease > maxIncrease {
		Mmsg("ERROR: func DimGrow: stepIncrease must be smaller (<) than maxIncrease")
	} else {
		if im.rGrowShrink == BLANKFREC {
			im.rGrowShrink = r
		}
		if im.onoffGrowShrink {
			if loop {
				im.rGrowShrink.X += stepIncrease
				im.rGrowShrink.Y += stepIncrease
				im.rGrowShrink.W -= stepIncrease * 2
				im.rGrowShrink.H -= stepIncrease * 2
				if im.rGrowShrink.X >= r.X {
					im.onoffGrowShrink = false
				}
			}
		} else {
			im.rGrowShrink.X -= stepIncrease
			im.rGrowShrink.Y -= stepIncrease
			im.rGrowShrink.W += stepIncrease * 2
			im.rGrowShrink.H += stepIncrease * 2
			if ABSDIFF(r.X, im.rGrowShrink.X) >= maxIncrease {
				im.onoffGrowShrink = true
			}
		}
		RND.CopyF(im.tex, &im.r, &im.rGrowShrink)
	}
	return im
}

func DimInset(im IM, r sdl.FRect, inset float32) {
	r.X += inset
	r.Y += inset
	r.W -= inset * 2
	r.H -= inset * 2
	RND.CopyF(im.tex, &im.r, &r)
}
func DimFlipHorizontal(im IM, r sdl.FRect) {
	RND.CopyExF(im.tex, &im.r, &r, 0, ORIGIN(r), sdl.FLIP_HORIZONTAL)
}
func DimFlipVertical(im IM, r sdl.FRect) {
	RND.CopyExF(im.tex, &im.r, &r, 0, ORIGIN(r), sdl.FLIP_VERTICAL)
}
func DimFlipHorizontalVertical(im IM, r sdl.FRect) {
	RND.CopyExF(im.tex, &im.r, &r, 0, ORIGIN(r), sdl.FLIP_VERTICAL|sdl.FLIP_HORIZONTAL)
}
func DimSheet(ims []IM, startX, startY, space, zoom float32) {
	//Mmsg("Turn on Debug (F1 key default) to view image & collision rectangles")
	ox := startX
	for i := range ims {
		r := sdl.FRect{startX, startY, float32(ims[i].r.W) * zoom, float32(ims[i].r.H) * zoom}
		Dim(ims[i], r)
		if DEBUG {
			DrecLine(r, ORANGE())
			if ims[i].rCollision != BLANKFREC {
				var xdiff, ydiff float32
				if ims[i].rCollision.X > float32(ims[i].r.X) {
					xdiff = float32(ims[i].rCollision.X - float32(ims[i].r.X))
				} else {
					xdiff = float32(float32(ims[i].r.X) - ims[i].rCollision.X)
				}
				if ims[i].rCollision.Y > float32(ims[i].r.Y) {
					ydiff = float32(ims[i].rCollision.Y - float32(ims[i].r.Y))
				} else {
					ydiff = float32(float32(ims[i].r.Y) - ims[i].rCollision.Y)
				}
				xdiff *= zoom
				ydiff *= zoom
				r2 := sdl.FRect{startX, startY, float32(ims[i].rCollision.W) * zoom, float32(ims[i].rCollision.H) * zoom}
				if ims[i].rCollision.X > float32(ims[i].r.X) {
					r2.X += xdiff
				} else {
					r2.X -= xdiff
				}
				if ims[i].rCollision.Y > float32(ims[i].r.Y) {
					r2.Y += ydiff
				} else {
					r2.Y -= ydiff
				}
				DrecLine(r2, MAGENTA())
			}

		}
		ctxt := WHITE()
		if BGCOL == WHITE() {
			ctxt = BLACK()
		}
		DtxtXY(fmt.Sprint(i), r.X, r.Y+r.H+2, FONT1DEFAULT, 1, ctxt)
		startX += float32(ims[i].r.W) * zoom
		startX += space
		if startX+float32(ims[i].r.W)*zoom > float32(WINW) {
			startX = ox
			startY += float32(ims[i].r.H) * zoom
			startY += space
			startY += FONT1DEFAULT.smlrH
		}
	}
}

// MARK: WINDOWS
func UDwindow(m WINDOW) WINDOW {
	if m.onoff {
		if m.shadow {
			c := COLORALPHA(DARKGREY(), 70)
			r := m.r
			r.X -= 5
			r.Y += 5
			DrecFill(r, c)
		}
		DrecFill(m.r, m.cBG)
		DrecFill(m.rBar, m.cBar)
		DrecLineWidth(m.r, m.outlineW, m.cLine)
		DtxtCenterXrec(m.nm, m.r, 2, FONT1DEFAULT, 1, m.cText)
		//CLOSE
		if m.closeIcon {
			if CrecPoint(MOUSE, m.rClose) && !m.move {
				ICONSSML[0].tex = TEXCOL(ICONSSML[0].tex, ORANGERED())
				Dim(ICONSSML[0], m.rClose)
				ICONSSML[0].tex = TEXCOLREVERT(ICONSSML[0].tex)
				if LCLICK {
					m.onoff = false
				}
			} else {
				ICONSSML[0].tex = TEXCOL(ICONSSML[0].tex, m.cText)
				Dim(ICONSSML[0], m.rClose)
				ICONSSML[0].tex = TEXCOLREVERT(ICONSSML[0].tex)
			}
		}
		//MOVE
		if CrecPoint(MOUSE, m.rBar) && clickHoldT > 0 {
			m.move = true
			if clickHoldT < 5 {
				menuXmouseW = mouseClickPoint.X - m.r.X
				menuYmouseH = mouseClickPoint.Y - m.r.Y
			}
		}
		if m.move {
			m.r.X = MOUSE.X - menuXmouseW
			m.r.Y = MOUSE.Y - menuYmouseH
			if m.r.X < 0 {
				m.r.X = 0
			}
			if m.r.X+m.r.W > float32(WINW) {
				m.r.X = float32(WINW) - m.r.W
			}
			if m.r.Y < 0 {
				m.r.Y = 0
			}
			if m.r.Y+m.r.H > float32(WINH) {
				m.r.Y = float32(WINH) - m.r.H
			}
			m.rBar.X = m.r.X
			m.rBar.Y = m.r.Y
			wClose := m.rBar.H - (float32(m.outlineW*2) + 4)
			m.rClose = sdl.FRect{m.rBar.X + m.rBar.W - (wClose + (float32(m.outlineW*2) + 2)), m.r.Y + (float32(m.outlineW+2) + 1), wClose, wClose}

			if !LCLICK {
				m.move = false
			}
		}

	}
	return m
}

// MARK: BUTTONS
func UDbuttonText(b BUTTON, txt string, upRightDownLeft1234 int, c sdl.Color) BUTTON {
	if upRightDownLeft1234 > 4 || upRightDownLeft1234 < 1 {
		Mmsg("ERROR: func UDbuttonText: upRightDownLeft1234 must be 1,2,3 or 4 / Up, Down, Left, Right")
	}
	l := CtxtLen(txt, FONT1DEFAULT, 1)
	var x, y float32
	cnt := POINTRECCENTER(b.r)
	switch upRightDownLeft1234 {
	case 1:
		x = cnt.X
		x -= l / 2
		y = b.r.Y
		y -= FONT1DEFAULT.smlrH + 4
		DtxtXY(txt, x, y, FONT1DEFAULT, 1, c)
	}
	return b
}
func UDbuttonXY(b BUTTON, x, y, w float32) BUTTON {
	b.r = sdl.FRect{x, y, w, w}
	if b.onOff {
		DrecFill(b.r, b.cOn)
		DimInset(b.im, b.r, 4)
	} else {
		DrecFill(b.r, b.cOff)
		DimInset(b.im, b.r, 4)
	}
	if b.canToggle {
		if CrecPoint(MOUSE, b.r) {
			DrecLineWidth(b.r, 1, COLORǁRANDOM())
			if LCLICK {
				b.onOff = !b.onOff
			}
		}
	}

	return b
}
func UDbutton(b BUTTON) BUTTON {
	if b.onOff {
		DrecFill(b.r, b.cOn)
		DimInset(b.im, b.r, 4)
	} else {
		DrecFill(b.r, b.cOff)
		DimInset(b.im, b.r, 4)
	}
	if b.canToggle {
		if CrecPoint(MOUSE, b.r) {
			DrecLineWidth(b.r, 1, COLORǁRANDOM())
			if LCLICK {
				b.onOff = !b.onOff
			}
		}
	}

	return b
}
func DbuttonSheet(b []BUTTON, startX, startY, space float32) {
	x := startX
	for i := range b {
		b[i].r.X = x
		b[i].r.Y = startY
		b[i] = UDbutton(b[i])
		ctxt := WHITE()
		if BGCOL == WHITE() {
			ctxt = BLACK()
		}
		DtxtXY(fmt.Sprint(i), b[i].r.X, b[i].r.Y+b[i].r.H+2, FONT1DEFAULT, 1, ctxt)
		x += b[i].r.W + space
		if x >= float32(WINW)-b[i].r.W {
			x = startX
			startY += b[i].r.H + space
			startY += FONT1DEFAULT.smlrH
		}
	}
}

// MARK: COLORS
func Dcolors() {
	var x, y float32
	siz := float32(WINH) / 18
	w := siz * 12
	x = CNTR.X - w/2
	xl := x + w
	y = CNTR.Y - w/2
	for i := range COLLIST {
		if CrecPoint(MOUSE, sdl.FRect{x, y, siz, siz}) {
			c := COLORǁRANDOM()
			c.A = mouseBlinkA
			mouseBlinkA -= 7
			DrecFillLineWidth(sdl.FRect{x, y, siz, siz}, 4, COLLIST[i].c, c)
		} else {
			if COLLIST[i].nm == "Black" {
				DrecFillLine(sdl.FRect{x, y, siz, siz}, COLLIST[i].c, WHITE())
			} else {
				DrecFillLine(sdl.FRect{x, y, siz, siz}, COLLIST[i].c, BLACK())
			}
		}
		x += siz
		if x >= xl {
			y += siz
			x = CNTR.X - w/2
		}
	}
}

// MARK: CUBES
func DCubeShadow(c CUBE, cShadow sdl.Color, topDarker, leftRightDarker bool, minShadow0to255, maxShadow0to255 uint8) {
	if minShadow0to255 >= maxShadow0to255 {
		Mmsg("ERROR: func DCubeShadow: minShadow0to255 must be smaller (<) than maxShadow0to255 >> minShadow0to255 set to 50 & maxShadow0to255 set to 150")
		minShadow0to255 = 50
		maxShadow0to255 = 150
	}
	if minShadow0to255 < 0 || maxShadow0to255 < 0 {
		Mmsg("ERROR: func DCubeShadow: minShadow0to255 & maxShadow0to255 must be larger (>) than zero >> minShadow0to255 set to 50 & maxShadow0to255 set to 150")
		minShadow0to255 = 50
		maxShadow0to255 = 150
	}
	if minShadow0to255 > 255 || maxShadow0to255 > 255 {
		Mmsg("ERROR: func DCubeShadow: minShadow0to255 & maxShadow0to255 must be smaller (<)than 256 >> minShadow0to255 set to 50 & maxShadow0to255 set to 150")
		minShadow0to255 = 50
		maxShadow0to255 = 150
	}
	midshadow := minShadow0to255 + ((maxShadow0to255 - minShadow0to255) / 2)
	if topDarker {
		DisoRecFillShadow(c.recs[1], cShadow, maxShadow0to255)
		if leftRightDarker {
			DisoRecFillShadow(c.recs[3], cShadow, midshadow)
			DisoRecFillShadow(c.recs[2], cShadow, minShadow0to255)
		} else {
			DisoRecFillShadow(c.recs[2], cShadow, minShadow0to255)
			DisoRecFillShadow(c.recs[3], cShadow, midshadow)
		}
	} else {
		DisoRecFillShadow(c.recs[1], cShadow, minShadow0to255)
		if leftRightDarker {
			DisoRecFillShadow(c.recs[3], cShadow, maxShadow0to255)
			DisoRecFillShadow(c.recs[2], cShadow, midshadow)
		} else {
			DisoRecFillShadow(c.recs[2], cShadow, midshadow)
			DisoRecFillShadow(c.recs[3], cShadow, maxShadow0to255)
		}
	}
}
func DCubeShadowLine(c CUBE, cShadow sdl.Color, topDarker, leftRightDarker bool, minShadow0to255, maxShadow0to255 uint8) {
	DCubeShadow(c, cShadow, topDarker, leftRightDarker, minShadow0to255, maxShadow0to255)
	DcubeLineFront(c)
}

func DcubeLineAll(c CUBE) {
	for i := range c.recs {
		DisoRecLine(c.recs[i])
	}
}
func DcubeLineBack(c CUBE) {
	DisoRecLine(c.recs[0])
	DisoRecLine(c.recs[4])
	DisoRecLine(c.recs[5])
}
func DcubeLineFront(c CUBE) {
	DisoRecLine(c.recs[1])
	DisoRecLine(c.recs[2])
	DisoRecLine(c.recs[3])
}
func DcubeLineFill(c CUBE) {
	DcubeFill(c)
	DcubeLineFront(c)
	DcubeLineFront(c)
}
func DcubeFill(c CUBE) {
	DisoRecFill(c.recs[1])
	DisoRecFill(c.recs[2])
	DisoRecFill(c.recs[3])
}

// MARK: STARS
func DstarLine(s STAR) {
	DlinePoints64CloseEnd(s.p, s.c)
}
func DstarFill(s STAR) {
	for i := range s.tri {
		DtriFill(s.tri[i])
	}
}

// MARK: CIRCLES ARCS
func DcircFillGradientCirc(circ CIRC, cOuter, cInner sdl.Color) {
	width := int32(circ.radius * 2)
	height := int32(circ.radius * 2)
	circTEX, _ := RND.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, width, height)
	circTEX.SetBlendMode(sdl.BLENDMODE_BLEND)
	pixels := make([]byte, width*height*4)
	centerX := float32(circ.radius)
	centerY := float32(circ.radius)
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			dx := float32(x) - centerX
			dy := float32(y) - centerY
			dist := float32(math.Sqrt(float64(dx*dx + dy*dy)))
			offset := (y*int(width) + x) * 4
			if dist <= circ.radius {
				t := dist / circ.radius
				rVal := uint8(float32(cInner.R)*(1-t) + float32(cOuter.R)*t)
				gVal := uint8(float32(cInner.G)*(1-t) + float32(cOuter.G)*t)
				bVal := uint8(float32(cInner.B)*(1-t) + float32(cOuter.B)*t)
				aVal := uint8(float32(cInner.A)*(1-t) + float32(cOuter.A)*t)
				pixels[offset] = rVal   // Red
				pixels[offset+1] = gVal // Green
				pixels[offset+2] = bVal // Blue
				pixels[offset+3] = aVal // Alpha
			} else {
				pixels[offset] = 0
				pixels[offset+1] = 0
				pixels[offset+2] = 0
				pixels[offset+3] = 0
			}
		}
	}
	pitch := int(width) * 4
	circTEX.Update(nil, unsafe.Pointer(&pixels[0]), pitch)
	RND.CopyF(circTEX, nil, &sdl.FRect{circ.cnt.X - circ.radius, circ.cnt.Y - circ.radius, circ.radius * 2, circ.radius * 2})
}
func DcircFillGradientCenter(cnt sdl.FPoint, radius float32, cOuter, cInner sdl.Color) {
	width := int32(radius * 2)
	height := int32(radius * 2)
	circTEX, _ := RND.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, width, height)
	circTEX.SetBlendMode(sdl.BLENDMODE_BLEND)
	pixels := make([]byte, width*height*4)
	centerX := float32(radius)
	centerY := float32(radius)
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			dx := float32(x) - centerX
			dy := float32(y) - centerY
			dist := float32(math.Sqrt(float64(dx*dx + dy*dy)))
			offset := (y*int(width) + x) * 4
			if dist <= radius {
				t := dist / radius
				rVal := uint8(float32(cInner.R)*(1-t) + float32(cOuter.R)*t)
				gVal := uint8(float32(cInner.G)*(1-t) + float32(cOuter.G)*t)
				bVal := uint8(float32(cInner.B)*(1-t) + float32(cOuter.B)*t)
				aVal := uint8(float32(cInner.A)*(1-t) + float32(cOuter.A)*t)
				pixels[offset] = rVal   // Red
				pixels[offset+1] = gVal // Green
				pixels[offset+2] = bVal // Blue
				pixels[offset+3] = aVal // Alpha
			} else {
				pixels[offset] = 0
				pixels[offset+1] = 0
				pixels[offset+2] = 0
				pixels[offset+3] = 0
			}
		}
	}
	pitch := int(width) * 4
	circTEX.Update(nil, unsafe.Pointer(&pixels[0]), pitch)
	RND.CopyF(circTEX, nil, &sdl.FRect{cnt.X - radius, CNTR.Y - radius, radius * 2, radius * 2})
}

func Darc(x, y, radius float32, startAngle, endAngle float64, c sdl.Color) {
	COL(c)
	step := 0.1
	for angle := startAngle; angle <= endAngle; angle += step {
		rad := angle * math.Pi / 180.0
		x := x + float32(float64(radius)*math.Cos(rad))
		y := y + float32(float64(radius)*math.Sin(rad))
		RND.DrawPointF(x, y)
	}
}
func DarcLineWidth(x, y, radius float32, startAngle, endAngle float64, lineW int, c sdl.Color) {
	COL(c)
	for lineW > 0 {
		step := 0.1
		for angle := startAngle; angle <= endAngle; angle += step {
			rad := angle * math.Pi / 180.0
			x := x + float32(float64(radius)*math.Cos(rad))
			y := y + float32(float64(radius)*math.Sin(rad))
			RND.DrawPointF(x, y)
		}
		radius--
		lineW--
	}
}
func DcircleFill(circ CIRC) {
	if len(circ.gradient) > 0 {
		for i := range circ.gradient {
			COL(circ.gradient[i].c)
			RND.DrawPointF(circ.gradient[i].xy.X, circ.gradient[i].xy.Y)
		}
	} else {
		var x, y float32 = 0, circ.radius
		var d float32 = 3 - 2*circ.radius
		drawCircleLine(circ.cnt.X, circ.cnt.Y, x, y, circ.c)
		for y >= x {
			x++
			if d > 0 {
				y--
				d = d + 4*(x-y) + 10
			} else {
				d = d + 4*x + 6
			}
			drawCircleLine(circ.cnt.X, circ.cnt.Y, x, y, circ.c)
		}
	}
}
func DcircleLine(circ CIRC) {
	var x, y float32 = 0, circ.radius
	var d float32 = 3 - 2*circ.radius
	drawCircle(circ.cnt.X, circ.cnt.Y, x, y, circ.c)
	for y >= x {
		x++
		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}
		drawCircle(circ.cnt.X, circ.cnt.Y, x, y, circ.c)
	}
}
func DcircleLineCenter(cnt sdl.FPoint, radius float32, c sdl.Color) {
	var x, y float32 = 0, radius
	var d float32 = 3 - 2*radius
	drawCircle(cnt.X, cnt.Y, x, y, c)
	for y >= x {
		x++
		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}
		drawCircle(cnt.X, cnt.Y, x, y, c)
	}
}
func DcircleFillCenter(cnt sdl.FPoint, radius float32, c sdl.Color) {
	var x, y float32 = 0, radius
	var d float32 = 3 - 2*radius
	drawCircleLine(cnt.X, cnt.Y, x, y, c)
	for y >= x {
		x++
		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}
		drawCircleLine(cnt.X, cnt.Y, x, y, c)
	}
}
func drawCircle(xc, yc, x, y float32, c sdl.Color) {
	COL(c)
	RND.DrawPointF(xc+x, yc+y)
	RND.DrawPointF(xc-x, yc+y)
	RND.DrawPointF(xc+x, yc-y)
	RND.DrawPointF(xc-x, yc-y)
	RND.DrawPointF(xc+y, yc+x)
	RND.DrawPointF(xc-y, yc+x)
	RND.DrawPointF(xc+y, yc-x)
	RND.DrawPointF(xc-y, yc-x)
}

func drawCircleLine(xc, yc, x, y float32, c sdl.Color) {
	COL(c)
	RND.DrawLineF(xc+x, yc+y, xc-x, yc+y)
	RND.DrawLineF(xc+x, yc-y, xc-x, yc-y)
	RND.DrawLineF(xc+y, yc+x, xc-y, yc+x)
	RND.DrawLineF(xc+y, yc-x, xc-y, yc-x)
}

// MARK: POLYGONS
func DpolygonLine(p POLYGON) {
	for i := range len(p.vert) - 1 {
		Dline(p.vert[i], p.vert[i+1], p.c)
	}
	Dline(p.vert[0], p.vert[len(p.vert)-1], p.c)
}
func DpolygonFill(p POLYGON) {
	for i := range p.tri {
		DtriFill(p.tri[i])
	}
}

// MARK: TRIANGLES
func DtriShadow(t TRI, c sdl.Color, alpha uint8) {
	t2 := TRICOLORFILLALPHA(t, c, alpha)
	DtriFill(t)
	DtriFill(t2)
}
func DtriLine(t TRI) {
	Dline(t.vert[0], t.vert[1], t.c)
	Dline(t.vert[1], t.vert[2], t.c)
	Dline(t.vert[0], t.vert[2], t.c)
}
func DtriFill(t TRI) {
	v := TRIPOINT2VERT(t)
	RND.RenderGeometry(nil, v[:3], nil)
}
func DtriFillLine(t TRI) {
	v := TRIPOINT2VERT(t)
	RND.RenderGeometry(nil, v[:3], nil)
	Dline(t.vert[0], t.vert[1], t.cLine)
	Dline(t.vert[1], t.vert[2], t.cLine)
	Dline(t.vert[0], t.vert[2], t.cLine)
}

// MARK: LINES
func DlinePointsCloseEnd(p []sdl.FPoint, c sdl.Color) {
	for i := range len(p) - 1 {
		Dline(p[i], p[i+1], c)
	}
	Dline(p[0], p[len(p)-1], c)
}
func DlinePoints64CloseEnd(p []POINT64, c sdl.Color) {
	for i := range len(p) - 1 {
		p1 := POINT642POINT(p[i])
		p2 := POINT642POINT(p[i+1])
		Dline(p1, p2, c)
	}
	p1 := POINT642POINT(p[0])
	p2 := POINT642POINT(p[len(p)-1])
	Dline(p1, p2, c)
}
func Dline(p1, p2 sdl.FPoint, c sdl.Color) {
	COL(c)
	RND.DrawLineF(p1.X, p1.Y, p2.X, p2.Y)
}
func DdiagRecLines(r sdl.FRect, c sdl.Color) {
	p := RECPOINTS(r)
	Dline(p[0], p[2], c)
	Dline(p[1], p[3], c)
}

// MARK: GRIDS
func DgridDebug(g GRID, c sdl.Color) {
	for i := range g.r {
		DrecLine(g.r[i], g.c)
		DtxtCenterRec(fmt.Sprint(i), FONT1DEFAULT, 1, c, g.r[i])
		DcircleFillCenter(g.cnt, 4, ORANGE())
	}
	x := g.r[0].X + g.w + 4
	y := g.r[0].Y + 4
	DtxtXY("Num Blocks Width: "+fmt.Sprint(g.numW), x, y, FONT1DEFAULT, 1, c)
	y += FONT1DEFAULT.smlrH
	DtxtXY("Num Blocks Height: "+fmt.Sprint(g.numH), x, y, FONT1DEFAULT, 1, c)

}
func DgridTilesShrink(g GRID, stepDecrease, maxDecrease float32, loop bool, drawLines bool, cLine sdl.Color) {
	if stepDecrease > maxDecrease {
		Mmsg("ERROR: func DimGrow: stepDecrease must be smaller (<) than maxDecrease")
	} else {
		for i := range g.r {
			if g.tiles[i].rGrowShrink == BLANKFREC {
				g.tiles[i].rGrowShrink = g.r[i]
				g.tiles[0].xShrink = g.r[0].X
			}
			if g.tiles[i].onoffGrowShrink {
				if loop {
					g.tiles[i].rGrowShrink.X -= stepDecrease
					g.tiles[i].rGrowShrink.Y -= stepDecrease
					g.tiles[i].rGrowShrink.W += stepDecrease * 2
					g.tiles[i].rGrowShrink.H += stepDecrease * 2
				}
			} else {
				g.tiles[i].rGrowShrink.X += stepDecrease
				g.tiles[i].rGrowShrink.Y += stepDecrease
				g.tiles[i].rGrowShrink.W -= stepDecrease * 2
				g.tiles[i].rGrowShrink.H -= stepDecrease * 2
			}
		}

		w := g.tiles[0].rGrowShrink.W * float32(g.numW)
		h := g.tiles[0].rGrowShrink.H * float32(g.numH)

		c := 0
		x := g.cnt.X - w/2
		ox := x
		y := g.cnt.Y - h/2
		for i := range g.r {
			g.tiles[i].rGrowShrink.X = x
			g.tiles[i].rGrowShrink.Y = y
			x += g.tiles[0].rGrowShrink.W
			c++
			if c == g.numW {
				x = ox
				y += g.tiles[0].rGrowShrink.H
				c = 0
			}
		}

		for i := range g.r {
			Dim(g.tiles[i], g.tiles[i].rGrowShrink)
			if drawLines {
				DrecLine(g.tiles[i].rGrowShrink, cLine)
			}
		}

		if g.tiles[0].rGrowShrink.X >= g.tiles[0].xShrink+maxDecrease {
			for i := range g.r {
				g.tiles[i].onoffGrowShrink = true
			}
		}
		if g.tiles[0].rGrowShrink.X <= g.tiles[0].xShrink {
			for i := range g.r {
				g.tiles[i].onoffGrowShrink = false
			}
		}

	}

}
func DgridTilesGrow(g GRID, stepIncrease, maxIncrease float32, loop bool, drawLines bool, cLine sdl.Color) {
	if stepIncrease > maxIncrease {
		Mmsg("ERROR: func DimGrow: stepIncrease must be smaller (<) than maxIncrease")
	} else {
		for i := range g.r {
			if g.tiles[i].rGrowShrink == BLANKFREC {
				g.tiles[i].rGrowShrink = g.r[i]
			}
			if g.tiles[i].onoffGrowShrink {
				if loop {
					g.tiles[i].rGrowShrink.X += stepIncrease
					g.tiles[i].rGrowShrink.Y += stepIncrease
					g.tiles[i].rGrowShrink.W -= stepIncrease * 2
					g.tiles[i].rGrowShrink.H -= stepIncrease * 2
				}
			} else {
				g.tiles[i].rGrowShrink.X -= stepIncrease
				g.tiles[i].rGrowShrink.Y -= stepIncrease
				g.tiles[i].rGrowShrink.W += stepIncrease * 2
				g.tiles[i].rGrowShrink.H += stepIncrease * 2
			}
		}

		w := g.tiles[0].rGrowShrink.W * float32(g.numW)
		h := g.tiles[0].rGrowShrink.H * float32(g.numH)

		c := 0
		x := g.cnt.X - w/2
		ox := x
		y := g.cnt.Y - h/2
		for i := range g.r {
			g.tiles[i].rGrowShrink.X = x
			g.tiles[i].rGrowShrink.Y = y
			x += g.tiles[0].rGrowShrink.W
			c++
			if c == g.numW {
				x = ox
				y += g.tiles[0].rGrowShrink.H
				c = 0
			}
		}

		for i := range g.r {
			Dim(g.tiles[i], g.tiles[i].rGrowShrink)
			if drawLines {
				DrecLine(g.tiles[i].rGrowShrink, cLine)
			}
		}

		if float32(g.tiles[0].r.X)-ox >= maxIncrease {
			for i := range g.r {
				g.tiles[i].onoffGrowShrink = true
			}
		}
		if g.tiles[0].rGrowShrink.X >= g.r[0].X {
			for i := range g.r {
				g.tiles[i].onoffGrowShrink = false
			}
		}

	}

}
func DgridTilesRotate(g GRID, angle float64, drawLines bool, cLine sdl.Color) {
	for i := range g.r {
		DimRotate(g.tiles[i], g.r[i], angle)
		if drawLines {
			p := Mpolygon(POINTRECCENTER(g.r[i]), float64(RECOUTERCIRCRADIUS(g.r[i].W)), 4, angle-45, cLine)
			DpolygonLine(p)
		}
	}
}
func DgridTiles(g GRID, drawLines bool, cLine sdl.Color) {
	for i := range g.r {
		Dim(g.tiles[i], g.r[i])
		if drawLines {
			DrecLine(g.r[i], cLine)
		}
	}
}
func DgridFillOffset(g GRID, offset float32) {
	for i := range g.r {
		r := g.r[i]
		r.X += offset
		r.Y += offset
		r.W -= offset * 2
		r.H -= offset * 2
		DrecFill(r, g.c)
	}
}

func DgridFill(g GRID, outline bool) {
	for i := range g.r {
		DrecFill(g.r[i], g.c)
	}
	if outline {
		for i := range g.r {
			DrecLine(g.r[i], BLACK())
		}
	}
}
func DgridFillColors(g GRID, cLines, cFill sdl.Color) {
	for i := range g.r {
		DrecFill(g.r[i], cFill)
	}
	for i := range g.r {
		DrecLine(g.r[i], cLines)
	}
}
func DmouseGrid(g GRID, c sdl.Color, blink bool) {
	if blink {
		mouseBlinkA -= 7
		c.A = mouseBlinkA
	}
	for i := range g.r {
		if POINTINREC(MOUSE, g.r[i]) {
			DrecFill(g.r[i], c)
			break
		}
	}
}

func Dgrid(g GRID) {
	for i := range g.r {
		DrecLine(g.r[i], g.c)
	}
}

func DgridXY(x, y, w float32, numW, numH int, c sdl.Color) {
	a := numW * numH
	ox := x
	cc := 0
	for a > 0 {
		DsquareLine(x, y, w, c)
		x += w
		cc++
		a--
		if cc == numW {
			x = ox
			y += w
			cc = 0
		}
	}
}

// MARK: RECS
func DrecSheetNums(r []sdl.FRect, cLine, cFill, cText sdl.Color, fon FONT, siz int) {
	for i := range r {
		DrecFillLine(r[i], cFill, cLine)
		DtxtBottomLeftRec(fmt.Sprint(i), fon, siz, cText, r[i], 4, 4)
	}
}
func DrecFillGradient(rect sdl.FRect, c1, c2 sdl.Color, horizVert bool) {
	// Convert rectangle dimensions to integers.
	startX := rect.X
	startY := rect.Y
	width := rect.W
	height := rect.H

	// Draw gradient either horizontally or vertically.
	if horizVert {
		// Iterate over each x value inside the rectangle.
		// Using width-1 so that factor runs exactly from 0 to 1 over the width.
		for i := 0; i < int(width); i++ {
			factor := float64(i) / float64(width-1)
			c := INTERPOLATECOLOR(c1, c2, factor)
			// Set render draw color to the interpolated color
			RND.SetDrawColor(c.R, c.G, c.B, c.A)
			// Draw a vertical line at x=i across the entire height of the rectangle.
			x := startX + float32(i)
			RND.DrawLineF(x, startY, x, startY+height)
		}
	} else {
		// Iterate over each y value inside the rectangle.
		for i := 0; i < int(height); i++ {
			factor := float64(i) / float64(height-1)
			c := INTERPOLATECOLOR(c1, c2, factor)
			RND.SetDrawColor(c.R, c.G, c.B, c.A)
			// Draw a horizontal line at y=i across the entire width of the rectangle.
			y := startY + float32(i)
			RND.DrawLineF(startX, y, startX+width, y)
		}
	}
}

func DrecXY2ColorDiagonal(x, y, w, h float32, c1, c2, cLine sdl.Color, leftRightDiagonal bool) {
	v1 := sdl.FPoint{x, y}
	v2 := v1
	v2.X += w
	v3 := v2
	v3.Y += h
	v4 := v3
	v4.X -= w
	t := TRI{}
	if leftRightDiagonal {
		t = MtriPointsMultiColor([]sdl.FPoint{v2, v3, v4}, c1, c2, c1, cLine)
		DtriFill(t)
		t = MtriPointsMultiColor([]sdl.FPoint{v4, v1, v2}, c1, c2, c1, cLine)
		DtriFill(t)
	} else {
		t = MtriPointsMultiColor([]sdl.FPoint{v1, v2, v3}, c1, c2, c1, cLine)
		DtriFill(t)
		t = MtriPointsMultiColor([]sdl.FPoint{v3, v4, v1}, c1, c2, c1, cLine)
		DtriFill(t)
	}

	DrecLine(sdl.FRect{x, y, w, h}, cLine)
}
func DrecXY2ColorPyramid(x, y, w, h float32, c1, c2, cLine sdl.Color) {
	v1 := sdl.FPoint{x, y}
	v2 := v1
	v2.X += w
	v3 := v2
	v3.Y += h
	v4 := v3
	v4.X -= w
	v5 := POINTRECCENTERXY(x, y, w, h)
	t := MtriPointsMultiColor([]sdl.FPoint{v1, v2, v5}, c1, c1, c2, cLine)
	DtriFill(t)
	t = MtriPointsMultiColor([]sdl.FPoint{v2, v3, v5}, c1, c1, c2, cLine)
	DtriFill(t)
	t = MtriPointsMultiColor([]sdl.FPoint{v3, v4, v5}, c1, c1, c2, cLine)
	DtriFill(t)
	t = MtriPointsMultiColor([]sdl.FPoint{v4, v1, v5}, c1, c1, c2, cLine)
	DtriFill(t)
	DrecLine(sdl.FRect{x, y, w, h}, cLine)
}
func DrecRoundCornersFillLineWidth(x, y, w, h, radius float32, lineW int, cFill, cLine sdl.Color) {
	for i := range int(h) {
		curY := y + float32(i)
		var left, right float32
		if i < int(radius) {
			dy := float64(radius - float32(i))
			dx := float32(math.Round(math.Sqrt(float64(radius*radius) - dy*dy)))
			left = x + radius - dx
			right = x + w - (radius - dx) - 1
		} else if i >= int(h-radius) {
			dy := float64(float32(i) - (h - radius - 1))
			dx := float32(math.Round(math.Sqrt(float64(radius*radius) - dy*dy)))
			left = x + radius - dx
			right = x + w - (radius - dx) - 1
		} else {
			left = x
			right = x + w - 1
		}
		COL(cFill)
		RND.DrawLineF(left, curY, right, curY)
	}
	DrecRoundCornersLineWidth(x, y, w, h, radius, lineW, cLine)
}
func DrecRoundCornersFill(x, y, w, h, radius float32, c sdl.Color) {
	for i := range int(h) {
		curY := y + float32(i)
		var left, right float32
		if i < int(radius) {
			dy := float64(radius - float32(i))
			dx := float32(math.Round(math.Sqrt(float64(radius*radius) - dy*dy)))
			left = x + radius - dx
			right = x + w - (radius - dx) - 1
		} else if i >= int(h-radius) {
			dy := float64(float32(i) - (h - radius - 1))
			dx := float32(math.Round(math.Sqrt(float64(radius*radius) - dy*dy)))
			left = x + radius - dx
			right = x + w - (radius - dx) - 1
		} else {
			left = x
			right = x + w - 1
		}
		COL(c)
		RND.DrawLineF(left, curY, right, curY)
	}
}
func DrecRoundCornersLineWidth(x, y, w, h, radius float32, lineW int, c sdl.Color) {
	COL(c)
	num := lineW
	y2 := y
	for num > 0 {
		RND.DrawLineF(x+radius, y2, x+w-radius, y2)
		y2++
		num--
	}
	num = lineW
	y2 = y
	for num > 0 {
		RND.DrawLineF(x+radius, y2+h, x+w-radius, y2+h)
		y2--
		num--
	}
	num = lineW
	x2 := x
	for num > 0 {
		RND.DrawLineF(x2, y+radius, x2, y+h-radius)
		x2++
		num--
	}
	num = lineW
	x2 = x
	for num > 0 {
		RND.DrawLineF(x2+w, y+radius, x2+w, y+h-radius)
		x2--
		num--
	}

	DarcLineWidth(x+radius, y+radius, radius, 180, 270, lineW, c)
	DarcLineWidth(x+w-radius, y+radius, radius, 270, 360, lineW, c)
	DarcLineWidth(x+w-radius, y+h-radius, radius, 0, 90, lineW, c)
	DarcLineWidth(x+radius, y+h-radius, radius, 90, 180, lineW, c)
}
func DrecRoundCornersLine(x, y, w, h, radius float32, c sdl.Color) {
	COL(c)
	RND.DrawLineF(x+radius, y, x+w-radius, y)
	RND.DrawLineF(x+radius, y+h, x+w-radius, y+h)
	RND.DrawLineF(x, y+radius, x, y+h-radius)
	RND.DrawLineF(x+w, y+radius, x+w, y+h-radius)
	Darc(x+radius, y+radius, radius, 180, 270, c)
	Darc(x+w-radius, y+radius, radius, 270, 360, c)
	Darc(x+w-radius, y+h-radius, radius, 0, 90, c)
	Darc(x+radius, y+h-radius, radius, 90, 180, c)
}

func DchessBoard(x, y, wBlok float32, numBloksRow int, cLines, cFill sdl.Color) {
	a := numBloksRow * numBloksRow
	var r []sdl.FRect
	ox := x
	c := 0
	for a > 0 {
		r = append(r, sdl.FRect{x, y, wBlok, wBlok})
		x += wBlok
		c++
		if c == numBloksRow {
			c = 0
			x = ox
			y += wBlok
		}
		a--
	}
	onoff := false
	c = 0
	for i := range r {
		if onoff {
			DrecFill(r[i], cFill)
			DrecLine(r[i], cLines)
		} else {
			DrecLine(r[i], cLines)
		}
		c++
		if c == numBloksRow {
			c = 0
		} else {
			onoff = !onoff
		}
	}

}
func DsquareLine(x, y, w float32, c sdl.Color) {
	COL(c)
	RrecLine(sdl.FRect{x, y, w, w})
}
func DsquareFill(x, y, w float32, c sdl.Color) {
	COL(c)
	RrecFill(sdl.FRect{x, y, w, w})
}
func DrecFill(r sdl.FRect, c sdl.Color) {
	COL(c)
	RrecFill(r)
}
func DrecFillLine(r sdl.FRect, c, cLine sdl.Color) {
	COL(c)
	RrecFill(r)
	COL(cLine)
	RrecLine(r)
}
func DrecLine(r sdl.FRect, c sdl.Color) {
	COL(c)
	RrecLine(r)
}
func DrecLineWidth(r sdl.FRect, lineW int, c sdl.Color) {
	COL(c)
	RrecLine(r)
	for lineW > 0 {
		r.X++
		r.Y++
		r.W -= 2
		r.H -= 2
		RrecLine(r)
		lineW--
	}
}
func DrecFillLineWidth(r sdl.FRect, lineW int, c, cLine sdl.Color) {
	COL(c)
	RrecFill(r)
	DrecLineWidth(r, lineW, cLine)
}

// MARK: RENDER
func RrecLine(r sdl.FRect) {
	RND.DrawRectF(&r)
}
func RrecFill(r sdl.FRect) {
	RND.FillRectF(&r)
}
