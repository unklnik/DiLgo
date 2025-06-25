package main

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	FONSPC, FONLINEH           float32 = 2, 2
	TXSURF                     *sdl.Surface
	FONT1DEFAULTPATH           = "fonts/Rubik-Medium.ttf"
	FONT2DEFAULTPATH           = "fonts/RubikDoodleShadow-Regular.ttf"
	FONT1DEFAULT, FONT2DEFAULT FONT

	standardCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:;<=>?!#$%&'()*+,-./@[]^_`{|}~'\"' "
)

type CHAR struct {
	ch    string
	tex   *sdl.Texture
	r     sdl.Rect
	c1    sdl.Color
	state int
}
type WORD struct {
	t string
	l float32
}

type FONT struct {
	name                                          string
	sizeSMLR, sizeSML, sizeMED, sizeLRG, sizeLRGR int
	charSMLR, charSML, charMED, charLRG, charLRGR []CHAR
	smlrH, smlH, medH, lrgH, lrgrH                float32
}

// DRAW
func DtxtCenterPoint(txt string, fon FONT, siz int, c sdl.Color, p sdl.FPoint) {
	x := p.X - CtxtLen(txt, fon, siz)/2
	y := p.Y - CfontHeight(fon, siz)/2
	x = ROUNDFLOAT(x, 0)
	DtxtXY(txt, x, y, fon, siz, c)
}
func DtxtTopLeftRec(txt string, fon FONT, siz int, c sdl.Color, r sdl.FRect, xOffset, yOffset float32) {
	DtxtXY(txt, r.X+xOffset, r.Y+yOffset, fon, siz, c)
}
func DtxtTopRightRec(txt string, fon FONT, siz int, c sdl.Color, r sdl.FRect, xOffset, yOffset float32) {
	l := CtxtLen(txt, fon, siz)
	DtxtXY(txt, (r.X+r.W)-(l+xOffset), r.Y+yOffset, fon, siz, c)
}
func DtxtBottomLeftRec(txt string, fon FONT, siz int, c sdl.Color, r sdl.FRect, xOffset, yOffset float32) {
	h := CfontHeight(fon, siz)
	DtxtXY(txt, r.X+xOffset, (r.Y+r.H)-(h+yOffset), fon, siz, c)
}
func DtxtBottomRightRec(txt string, fon FONT, siz int, c sdl.Color, r sdl.FRect, xOffset, yOffset float32) {
	l := CtxtLen(txt, fon, siz)
	h := CfontHeight(fon, siz)
	DtxtXY(txt, (r.X+r.W)-(l+xOffset), (r.Y+r.H)-(h+yOffset), fon, siz, c)
}
func DtxtCenterRec(txt string, fon FONT, siz int, c sdl.Color, r sdl.FRect) {
	l := CtxtLen(txt, fon, siz)
	h := CfontHeight(fon, siz)
	x := (r.X + r.W/2) - l/2
	y := (r.Y + r.H/2) - h/2
	DtxtXY(txt, x, y, fon, siz, c)
}
func DtxtCenterScreenletterAngle(txt string, topMiddleBottom123 int, offset float32, angle float64, fon FONT, siz int, c sdl.Color) {
	if siz > 5 || siz < 1 {
		Mmsg("ERROR: func DtxtCenterScreen: Size must be from 1 to 5 >> 1 = Smallest / 5 = Largest >> Set to 1")
		siz = 1
	}
	if topMiddleBottom123 > 3 || topMiddleBottom123 < 1 {
		Mmsg("ERROR: func DtxtCenterScreen: topMiddleBottom123 must be from 1 to 3 >> 1 = Top / 2 = Middle / 3 = Bottom >> Set to 2")
		topMiddleBottom123 = 2
	}
	l := CtxtLen(txt, fon, siz)
	x := ROUNDFLOAT(CNTR.X-l/2, 0)
	var y float32
	h := CfontHeight(fon, siz)
	switch topMiddleBottom123 {
	case 1:
		y = offset
	case 2:
		y = float32(int(CNTR.Y-h/2)) + offset
	case 3:
		y = float32(WINH) - (h + offset)
	}
	DtxtXYletterAngle(txt, x, y, angle, fon, siz, c)
}
func DtxtCenterScreenRandomColors(txt string, topMiddleBottom123 int, offset float32, fon FONT, siz int) {
	if siz > 5 || siz < 1 {
		Mmsg("ERROR: func DtxtCenterScreen: Size must be from 1 to 5 >> 1 = Smallest / 5 = Largest >> Set to 1")
		siz = 1
	}
	if topMiddleBottom123 > 3 || topMiddleBottom123 < 1 {
		Mmsg("ERROR: func DtxtCenterScreen: topMiddleBottom123 must be from 1 to 3 >> 1 = Top / 2 = Middle / 3 = Bottom >> Set to 2")
		topMiddleBottom123 = 2
	}
	l := CtxtLen(txt, fon, siz)
	x := ROUNDFLOAT(CNTR.X-l/2, 0)
	var y float32
	h := CfontHeight(fon, siz)
	switch topMiddleBottom123 {
	case 1:
		y = offset
	case 2:
		y = float32(int(CNTR.Y-h/2)) + offset
	case 3:
		y = float32(WINH) - (h + offset)
	}
	DtxtXYrandomColors(txt, x, y, fon, siz)
}
func DtxtXYletterAngleShadow(txt string, x, y, xShadowOffset, yShadowOffset float32, angle float64, fon FONT, siz int, cText, cShadow sdl.Color) {
	if siz > 5 || siz < 1 {
		Mmsg("ERROR: func DtxtXYletterAngleShadow: Size must be from 1 to 5 >> 1 = Smallest / 5 = Largest >> Set to 1")
		siz = 1
	}
	t := strings.Split(txt, "")
	for i := range t {
		for j := range fon.charSMLR {
			if t[i] == fon.charSMLR[j].ch {
				switch siz {
				case 1:
					//SHADOW
					fon.charSMLR[j].tex = TEXCOL(fon.charSMLR[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charSMLR[j].r.W), float32(fon.charSMLR[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyExF(fon.charSMLR[j].tex, &fon.charSMLR[j].r, &r, angle, ORIGIN(r), sdl.FLIP_NONE)
					x += float32(fon.charSMLR[j].r.W) + FONSPC
					fon.charSMLR[j].tex = TEXCOLREVERT(fon.charSMLR[j].tex)
					//TEXT
					fon.charSMLR[j].tex = TEXCOL(fon.charSMLR[j].tex, cText)
					RND.CopyExF(fon.charSMLR[j].tex, &fon.charSMLR[j].r, &sdl.FRect{x, y, float32(fon.charSMLR[j].r.W), float32(fon.charSMLR[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charSMLR[j].r.W), float32(fon.charSMLR[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charSMLR[j].r.W) + FONSPC
					fon.charSMLR[j].tex = TEXCOLREVERT(fon.charSMLR[j].tex)
					break
				case 2:
					//SHADOW
					fon.charSML[j].tex = TEXCOL(fon.charSML[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charSML[j].r.W), float32(fon.charSML[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyExF(fon.charSML[j].tex, &fon.charSML[j].r, &r, angle, ORIGIN(r), sdl.FLIP_NONE)
					x += float32(fon.charSML[j].r.W) + FONSPC
					fon.charSML[j].tex = TEXCOLREVERT(fon.charSML[j].tex)
					//TEXT
					fon.charSML[j].tex = TEXCOL(fon.charSML[j].tex, cText)
					RND.CopyExF(fon.charSML[j].tex, &fon.charSML[j].r, &sdl.FRect{x, y, float32(fon.charSML[j].r.W), float32(fon.charSML[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charSML[j].r.W), float32(fon.charSML[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charSML[j].r.W) + FONSPC
					fon.charSML[j].tex = TEXCOLREVERT(fon.charSML[j].tex)
					break
				case 3:
					//SHADOW
					fon.charMED[j].tex = TEXCOL(fon.charMED[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charMED[j].r.W), float32(fon.charMED[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyExF(fon.charMED[j].tex, &fon.charMED[j].r, &r, angle, ORIGIN(r), sdl.FLIP_NONE)
					x += float32(fon.charMED[j].r.W) + FONSPC
					fon.charMED[j].tex = TEXCOLREVERT(fon.charMED[j].tex)
					//TEXT
					fon.charMED[j].tex = TEXCOL(fon.charMED[j].tex, cText)
					RND.CopyExF(fon.charMED[j].tex, &fon.charMED[j].r, &sdl.FRect{x, y, float32(fon.charMED[j].r.W), float32(fon.charMED[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charMED[j].r.W), float32(fon.charMED[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charMED[j].r.W) + FONSPC
					fon.charMED[j].tex = TEXCOLREVERT(fon.charMED[j].tex)
					break
				case 4:
					//SHADOW
					fon.charLRG[j].tex = TEXCOL(fon.charLRG[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charLRG[j].r.W), float32(fon.charLRG[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyExF(fon.charLRG[j].tex, &fon.charLRG[j].r, &r, angle, ORIGIN(r), sdl.FLIP_NONE)
					x += float32(fon.charLRG[j].r.W) + FONSPC
					fon.charLRG[j].tex = TEXCOLREVERT(fon.charLRG[j].tex)
					//TEXT
					fon.charLRG[j].tex = TEXCOL(fon.charLRG[j].tex, cText)
					RND.CopyExF(fon.charLRG[j].tex, &fon.charLRG[j].r, &sdl.FRect{x, y, float32(fon.charLRG[j].r.W), float32(fon.charLRG[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charLRG[j].r.W), float32(fon.charLRG[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charLRG[j].r.W) + FONSPC
					fon.charLRG[j].tex = TEXCOLREVERT(fon.charLRG[j].tex)
					break
				case 5:
					//SHADOW
					fon.charLRGR[j].tex = TEXCOL(fon.charLRGR[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charLRGR[j].r.W), float32(fon.charSML[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyExF(fon.charLRGR[j].tex, &fon.charLRGR[j].r, &r, angle, ORIGIN(r), sdl.FLIP_NONE)
					x += float32(fon.charSML[j].r.W) + FONSPC
					fon.charSML[j].tex = TEXCOLREVERT(fon.charSML[j].tex)
					//TEXT
					fon.charLRGR[j].tex = TEXCOL(fon.charLRGR[j].tex, cText)
					RND.CopyExF(fon.charLRGR[j].tex, &fon.charLRGR[j].r, &sdl.FRect{x, y, float32(fon.charLRGR[j].r.W), float32(fon.charLRGR[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charLRGR[j].r.W), float32(fon.charLRGR[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charLRGR[j].r.W) + FONSPC
					fon.charLRGR[j].tex = TEXCOLREVERT(fon.charLRGR[j].tex)
					break
				}
			}
		}
	}
}
func DtxtXYletterAngle(txt string, x, y float32, angle float64, fon FONT, siz int, c sdl.Color) {
	if siz > 5 || siz < 1 {
		Mmsg("ERROR: func DtxtXYletterAngle: Size must be from 1 to 5 >> 1 = Smallest / 5 = Largest >> Set to 1")
		siz = 1
	}
	t := strings.Split(txt, "")
	for i := range t {
		for j := range fon.charSMLR {
			if t[i] == fon.charSMLR[j].ch {
				switch siz {
				case 1:
					fon.charSMLR[j].tex = TEXCOL(fon.charSMLR[j].tex, c)
					RND.CopyExF(fon.charSMLR[j].tex, &fon.charSMLR[j].r, &sdl.FRect{x, y, float32(fon.charSMLR[j].r.W), float32(fon.charSMLR[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charSMLR[j].r.W), float32(fon.charSMLR[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charSMLR[j].r.W) + FONSPC
					fon.charSMLR[j].tex = TEXCOLREVERT(fon.charSMLR[j].tex)
					break
				case 2:
					fon.charSML[j].tex = TEXCOL(fon.charSML[j].tex, c)
					RND.CopyExF(fon.charSML[j].tex, &fon.charSML[j].r, &sdl.FRect{x, y, float32(fon.charSML[j].r.W), float32(fon.charSML[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charSML[j].r.W), float32(fon.charSML[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charSML[j].r.W) + FONSPC
					fon.charSML[j].tex = TEXCOLREVERT(fon.charSML[j].tex)
					break
				case 3:
					fon.charMED[j].tex = TEXCOL(fon.charMED[j].tex, c)
					RND.CopyExF(fon.charMED[j].tex, &fon.charMED[j].r, &sdl.FRect{x, y, float32(fon.charMED[j].r.W), float32(fon.charMED[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charMED[j].r.W), float32(fon.charMED[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charMED[j].r.W) + FONSPC
					fon.charMED[j].tex = TEXCOLREVERT(fon.charMED[j].tex)
					break
				case 4:
					fon.charLRG[j].tex = TEXCOL(fon.charLRG[j].tex, c)
					RND.CopyExF(fon.charLRG[j].tex, &fon.charLRG[j].r, &sdl.FRect{x, y, float32(fon.charLRG[j].r.W), float32(fon.charLRG[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charLRG[j].r.W), float32(fon.charLRG[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charLRG[j].r.W) + FONSPC
					fon.charLRG[j].tex = TEXCOLREVERT(fon.charLRG[j].tex)
					break
				case 5:
					fon.charLRGR[j].tex = TEXCOL(fon.charLRGR[j].tex, c)
					RND.CopyExF(fon.charLRGR[j].tex, &fon.charLRGR[j].r, &sdl.FRect{x, y, float32(fon.charLRGR[j].r.W), float32(fon.charLRGR[j].r.H)}, angle, ORIGIN(sdl.FRect{x, y, float32(fon.charLRGR[j].r.W), float32(fon.charLRGR[j].r.H)}), sdl.FLIP_NONE)
					x += float32(fon.charLRGR[j].r.W) + FONSPC
					fon.charLRGR[j].tex = TEXCOLREVERT(fon.charLRGR[j].tex)
					break
				}
			}
		}
	}
}
func DtxtXYshadowMulti(txt string, x, y, xShadowOffset, yShadowOffset float32, fon FONT, siz int, cText sdl.Color, cShadow []sdl.Color) {
	for i := len(cShadow) - 1; i > 0; i-- {
		xShadow := xShadowOffset * float32(i)
		yShadow := yShadowOffset * float32(i)
		DtxtShadowOnly(txt, x, y, xShadow, yShadow, fon, siz, cShadow[i])
	}

	DtxtXY(txt, x, y, fon, siz, cText)
}
func DtxtXYshadow(txt string, x, y, xShadowOffset, yShadowOffset float32, fon FONT, siz int, cText, cShadow sdl.Color) {
	DtxtShadowOnly(txt, x, y, xShadowOffset, yShadowOffset, fon, siz, cShadow)
	DtxtXY(txt, x, y, fon, siz, cText)
}
func DtxtShadowOnly(txt string, x, y, xShadowOffset, yShadowOffset float32, fon FONT, siz int, cShadow sdl.Color) {
	t := strings.Split(txt, "")
	for i := range t {
		for j := range fon.charSMLR {
			if t[i] == fon.charSMLR[j].ch {
				switch siz {
				case 1:
					fon.charSMLR[j].tex = TEXCOL(fon.charSMLR[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charSMLR[j].r.W), float32(fon.charSMLR[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyF(fon.charSMLR[j].tex, &fon.charSMLR[j].r, &r)
					x += float32(fon.charSMLR[j].r.W) + FONSPC
					fon.charSMLR[j].tex = TEXCOLREVERT(fon.charSMLR[j].tex)
					break
				case 2:
					fon.charSML[j].tex = TEXCOL(fon.charSML[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charSML[j].r.W), float32(fon.charSML[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyF(fon.charSML[j].tex, &fon.charSML[j].r, &r)
					x += float32(fon.charSML[j].r.W) + FONSPC
					fon.charSML[j].tex = TEXCOLREVERT(fon.charSML[j].tex)
					break
				case 3:
					fon.charMED[j].tex = TEXCOL(fon.charMED[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charMED[j].r.W), float32(fon.charMED[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyF(fon.charMED[j].tex, &fon.charMED[j].r, &r)
					x += float32(fon.charMED[j].r.W) + FONSPC
					fon.charMED[j].tex = TEXCOLREVERT(fon.charMED[j].tex)
					break
				case 4:
					fon.charLRG[j].tex = TEXCOL(fon.charLRG[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charLRG[j].r.W), float32(fon.charLRG[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyF(fon.charLRG[j].tex, &fon.charLRG[j].r, &r)
					x += float32(fon.charLRG[j].r.W) + FONSPC
					fon.charLRG[j].tex = TEXCOLREVERT(fon.charLRG[j].tex)
					break
				case 5:
					fon.charLRGR[j].tex = TEXCOL(fon.charLRGR[j].tex, cShadow)
					r := sdl.FRect{x, y, float32(fon.charLRGR[j].r.W), float32(fon.charLRGR[j].r.H)}
					r.X += xShadowOffset
					r.Y += yShadowOffset
					RND.CopyF(fon.charLRGR[j].tex, &fon.charLRGR[j].r, &r)
					x += float32(fon.charLRGR[j].r.W) + FONSPC
					fon.charLRGR[j].tex = TEXCOLREVERT(fon.charLRGR[j].tex)
					break
				}
			}
		}
	}
}
func DtxtXYrandomColors(txt string, x, y float32, fon FONT, siz int) {
	if siz > 5 || siz < 1 {
		Mmsg("ERROR: func DtxtXY: Size must be from 1 to 5 >> 1 = Smallest / 5 = Largest >> Set to 1")
		siz = 1
	}
	t := strings.Split(txt, "")
	for i := range t {
		for j := range fon.charSMLR {
			if t[i] == fon.charSMLR[j].ch {
				switch siz {
				case 1:
					if fon.charSMLR[j].c1 == BLANKCOL {
						fon.charSMLR[j].c1 = COLORǁRANDOM()
					}
					fon.charSMLR[j].tex = TEXCOL(fon.charSMLR[j].tex, fon.charSMLR[j].c1)
					RND.CopyF(fon.charSMLR[j].tex, &fon.charSMLR[j].r, &sdl.FRect{x, y, float32(fon.charSMLR[j].r.W), float32(fon.charSMLR[j].r.H)})
					x += float32(fon.charSMLR[j].r.W) + FONSPC
					fon.charSMLR[j].tex = TEXCOLREVERT(fon.charSMLR[j].tex)
					break
				case 2:
					if fon.charSML[j].c1 == BLANKCOL {
						fon.charSML[j].c1 = COLORǁRANDOM()
					}
					fon.charSML[j].tex = TEXCOL(fon.charSML[j].tex, fon.charSML[j].c1)
					RND.CopyF(fon.charSML[j].tex, &fon.charSML[j].r, &sdl.FRect{x, y, float32(fon.charSML[j].r.W), float32(fon.charSML[j].r.H)})
					x += float32(fon.charSML[j].r.W) + FONSPC
					fon.charSML[j].tex = TEXCOLREVERT(fon.charSML[j].tex)
					break
				case 3:
					if fon.charMED[j].c1 == BLANKCOL {
						fon.charMED[j].c1 = COLORǁRANDOM()
					}
					fon.charMED[j].tex = TEXCOL(fon.charMED[j].tex, fon.charMED[j].c1)
					RND.CopyF(fon.charMED[j].tex, &fon.charMED[j].r, &sdl.FRect{x, y, float32(fon.charMED[j].r.W), float32(fon.charMED[j].r.H)})
					x += float32(fon.charMED[j].r.W) + FONSPC
					fon.charMED[j].tex = TEXCOLREVERT(fon.charMED[j].tex)
					break
				case 4:
					if fon.charLRG[j].c1 == BLANKCOL {
						fon.charLRG[j].c1 = COLORǁRANDOM()
					}
					fon.charLRG[j].tex = TEXCOL(fon.charLRG[j].tex, fon.charLRG[j].c1)
					RND.CopyF(fon.charLRG[j].tex, &fon.charLRG[j].r, &sdl.FRect{x, y, float32(fon.charLRG[j].r.W), float32(fon.charLRG[j].r.H)})
					x += float32(fon.charLRG[j].r.W) + FONSPC
					fon.charLRG[j].tex = TEXCOLREVERT(fon.charLRG[j].tex)
					break
				case 5:
					if fon.charLRGR[j].c1 == BLANKCOL {
						fon.charLRGR[j].c1 = COLORǁRANDOM()
					}
					fon.charLRGR[j].tex = TEXCOL(fon.charLRGR[j].tex, fon.charLRGR[j].c1)
					RND.CopyF(fon.charLRGR[j].tex, &fon.charLRGR[j].r, &sdl.FRect{x, y, float32(fon.charLRGR[j].r.W), float32(fon.charLRGR[j].r.H)})
					x += float32(fon.charLRGR[j].r.W) + FONSPC
					fon.charLRGR[j].tex = TEXCOLREVERT(fon.charLRGR[j].tex)
					break
				}
			}
		}
	}
}
func DtxtCenterScreen(txt string, topMiddleBottom123 int, offset float32, fon FONT, siz int, c sdl.Color) {
	if siz > 5 || siz < 1 {
		Mmsg("ERROR: func DtxtCenterScreen: Size must be from 1 to 5 >> 1 = Smallest / 5 = Largest >> Set to 1")
		siz = 1
	}
	if topMiddleBottom123 > 3 || topMiddleBottom123 < 1 {
		Mmsg("ERROR: func DtxtCenterScreen: topMiddleBottom123 must be from 1 to 3 >> 1 = Top / 2 = Middle / 3 = Bottom >> Set to 2")
		topMiddleBottom123 = 2
	}
	l := CtxtLen(txt, fon, siz)
	x := ROUNDFLOAT(CNTR.X-l/2, 0)
	var y float32
	h := CfontHeight(fon, siz)
	switch topMiddleBottom123 {
	case 1:
		y = offset
	case 2:
		y = float32(int(CNTR.Y-h/2)) + offset
	case 3:
		y = float32(WINH) - (h + offset)
	}
	DtxtXY(txt, x, y, fon, siz, c)
}
func DtxtXY(txt string, x, y float32, fon FONT, siz int, c sdl.Color) {
	if siz > 5 || siz < 1 {
		Mmsg("ERROR: func DtxtXY: Size must be from 1 to 5 >> 1 = Smallest / 5 = Largest >> Set to 1")
		siz = 1
	}
	x = ROUNDFLOAT(x, 0)
	y = ROUNDFLOAT(y, 0)
	t := strings.Split(txt, "")
	for i := range t {
		for j := range fon.charSMLR {
			if t[i] == fon.charSMLR[j].ch {
				switch siz {
				case 1:
					fon.charSMLR[j].tex = TEXCOL(fon.charSMLR[j].tex, c)
					RND.CopyF(fon.charSMLR[j].tex, &fon.charSMLR[j].r, &sdl.FRect{x, y, float32(fon.charSMLR[j].r.W), float32(fon.charSMLR[j].r.H)})
					x += float32(fon.charSMLR[j].r.W) + FONSPC
					fon.charSMLR[j].tex = TEXCOLREVERT(fon.charSMLR[j].tex)
					break
				case 2:
					fon.charSML[j].tex = TEXCOL(fon.charSML[j].tex, c)
					RND.CopyF(fon.charSML[j].tex, &fon.charSML[j].r, &sdl.FRect{x, y, float32(fon.charSML[j].r.W), float32(fon.charSML[j].r.H)})
					x += float32(fon.charSML[j].r.W) + FONSPC
					fon.charSML[j].tex = TEXCOLREVERT(fon.charSML[j].tex)
					break
				case 3:
					fon.charMED[j].tex = TEXCOL(fon.charMED[j].tex, c)
					RND.CopyF(fon.charMED[j].tex, &fon.charMED[j].r, &sdl.FRect{x, y, float32(fon.charMED[j].r.W), float32(fon.charMED[j].r.H)})
					x += float32(fon.charMED[j].r.W) + FONSPC
					fon.charMED[j].tex = TEXCOLREVERT(fon.charMED[j].tex)
					break
				case 4:
					fon.charLRG[j].tex = TEXCOL(fon.charLRG[j].tex, c)
					RND.CopyF(fon.charLRG[j].tex, &fon.charLRG[j].r, &sdl.FRect{x, y, float32(fon.charLRG[j].r.W), float32(fon.charLRG[j].r.H)})
					x += float32(fon.charLRG[j].r.W) + FONSPC
					fon.charLRG[j].tex = TEXCOLREVERT(fon.charLRG[j].tex)
					break
				case 5:
					fon.charLRGR[j].tex = TEXCOL(fon.charLRGR[j].tex, c)
					RND.CopyF(fon.charLRGR[j].tex, &fon.charLRGR[j].r, &sdl.FRect{x, y, float32(fon.charLRGR[j].r.W), float32(fon.charLRGR[j].r.H)})
					x += float32(fon.charLRGR[j].r.W) + FONSPC
					fon.charLRGR[j].tex = TEXCOLREVERT(fon.charLRGR[j].tex)
					break
				}
			}
		}
	}
}
func DtxtWrap(txt string, r sdl.FRect, inset, lineSpace float32, fon FONT, siz int, c sdl.Color) {
	x := r.X + inset
	xr := x + (r.W - inset*2)
	y := r.Y + inset
	var h float32
	switch siz {
	case 1:
		h = fon.smlrH
	case 2:
		h = fon.smlH
	case 3:
		h = fon.medH
	case 4:
		h = fon.lrgH
	case 5:
		h = fon.lrgrH
	}

	words := Mwords(txt, fon, siz)
	for i := range words {
		if x+words[i].l > xr {
			x = r.X + inset
			y += h + lineSpace
		}
		t := strings.Split(words[i].t, "")
		for j := range t {
			DtxtXY(t[j], x, y, fon, siz, c)
			x += CtxtLen(t[j], fon, siz)
		}

	}
}
func DtxtCenterXrec(txt string, r sdl.FRect, yOffset float32, fon FONT, siz int, c sdl.Color) {
	l := CtxtLen(txt, fon, siz)
	x := r.X + r.W/2
	x -= l / 2
	y := r.Y
	y += yOffset
	DtxtXY(txt, x, y, fon, siz, c)
}
func Dfps(x, y float32, fon FONT, siz int, fpsLeadingTxt bool, c sdl.Color) {
	t := fmt.Sprint(FPSCURRENT)
	if fpsLeadingTxt {
		t = "FPS: " + t
	}
	DtxtXY(t, x, y, fon, siz, c)
}
func DwordLen(w []WORD, x, y float32, c sdl.Color) {
	DtxtXY("Pixel length includes single character (5px) space at end of word", x, y, FONT1DEFAULT, 1, ORANGE())
	y += FONT1DEFAULT.smlrH
	for i := range w {
		DtxtXY("len: "+fmt.Sprint(w[i].l)+" word: "+w[i].t, x, y, FONT1DEFAULT, 1, c)
		y += FONT1DEFAULT.smlrH
	}
}
func Dwords(w []WORD, x, y float32, fon FONT, siz int, c sdl.Color) {
	ox := x
	for i := range w {
		if x+w[i].l > float32(WINW) {
			y += FONT1DEFAULT.smlrH
			x = ox
		}
		DrecLine(sdl.FRect{x, y - 2, w[i].l, FONT1DEFAULT.smlrH + 4}, ORANGE())
		t := strings.Split(w[i].t, "")
		for j := range t {
			DtxtXY(t[j], x, y, fon, siz, c)
			x += CtxtLen(t[j], fon, siz)
		}

	}

}

// MARK: UTILS
func CfontHeight(fon FONT, siz int) float32 {
	var h float32
	switch siz {
	case 1:
		h = fon.smlrH
	case 2:
		h = fon.smlH
	case 3:
		h = fon.medH
	case 4:
		h = fon.lrgH
	case 5:
		h = fon.lrgrH
	}
	return h
}
func CtxtLineLen(txt string, r sdl.FRect, inset, lineSpace float32, fon FONT, siz int) []float32 {
	x := r.X + inset
	xr := x + (r.W - inset*2)
	y := r.Y + inset
	var h float32
	switch siz {
	case 1:
		h = fon.smlrH
	case 2:
		h = fon.smlH
	case 3:
		h = fon.medH
	case 4:
		h = fon.lrgH
	case 5:
		h = fon.lrgrH
	}

	words := Mwords(txt, fon, siz)
	var l float32
	var linel []float32
	for i := range words {
		if x+words[i].l > xr {
			linel = append(linel, l)
			l = 0
			x = r.X + inset
			y += h + lineSpace
		}
		t := strings.Split(words[i].t, "")
		for j := range t {
			x += CtxtLen(t[j], fon, siz)
		}
		l += words[i].l
	}
	linel = append(linel, l)
	return linel
}

func CtxtLen(txt string, fon FONT, siz int) float32 {
	var txlen float32
	t := strings.Split(txt, "")
	for i := range t {
		for j := range FONT1DEFAULT.charSMLR {
			if t[i] == FONT1DEFAULT.charSMLR[j].ch {
				switch siz {
				case 1: //SMLR
					txlen += float32(fon.charSMLR[j].r.W) + FONSPC
					break
				case 2: //SML
					txlen += float32(fon.charSML[j].r.W) + FONSPC
					break
				case 3: //MED
					txlen += float32(fon.charMED[j].r.W) + FONSPC
					break
				case 4: //LRG
					txlen += float32(fon.charLRG[j].r.W) + FONSPC
					break
				case 5: //LRGR
					txlen += float32(fon.charLRGR[j].r.W) + FONSPC
					break
				}
			}
		}
	}
	return txlen
}

// MARK: MAKE

func Mwords(txt string, fon FONT, siz int) []WORD {
	t := strings.Split(txt, "")
	w := WORD{}
	var words []WORD
	letters := ""
	var l float32

	for i := range t {
		if t[i] == " " || i == len(t)-1 {
			letters = letters + t[i]
			l += CtxtLen(t[i], fon, siz)
			w.t = letters
			w.l = l
			words = append(words, w)
			l = 0
			letters = ""
			w = WORD{}
		} else {
			letters = letters + t[i]
			l += CtxtLen(t[i], fon, siz)
		}
	}
	return words
}

func Mfont(filePath string, name string, sizeSMLR, sizeSML, sizeMED, sizeLRG, sizeLRGR int) FONT {
	ttf.Init()
	fon := FONT{}
	sizes := []int{sizeSMLR, sizeSML, sizeMED, sizeLRG, sizeLRGR}
	t := strings.Split(standardCharacters, "")
	for i := range sizes {
		f, ERR := ttf.OpenFont(filePath, sizes[i])
		if ERR != nil {
			mERR(ERR)
		}
		for j := range t {
			switch i {
			case 0: //sizeSMLR
				fon.charSMLR = append(fon.charSMLR, mCHAR(t[j], f))
				fon.sizeSMLR = sizes[i]
				fon.smlrH = float32(fon.charSMLR[0].r.H)
			case 1: //sizeSML
				fon.charSML = append(fon.charSML, mCHAR(t[j], f))
				fon.sizeSML = sizes[i]
				fon.smlH = float32(fon.charSML[0].r.H)
			case 2: //sizeMED
				fon.charMED = append(fon.charMED, mCHAR(t[j], f))
				fon.sizeMED = sizes[i]
				fon.medH = float32(fon.charMED[0].r.H)
			case 3: //sizeLRG
				fon.charLRG = append(fon.charLRG, mCHAR(t[j], f))
				fon.sizeLRG = sizes[i]
				fon.lrgH = float32(fon.charLRG[0].r.H)
			case 4: //sizeLRGR
				fon.charLRGR = append(fon.charLRGR, mCHAR(t[j], f))
				fon.sizeLRGR = sizes[i]
				fon.lrgrH = float32(fon.charLRGR[0].r.H)
			}
		}
	}
	return fon
}

func mFONTSDEFAULT() {
	ttf.Init()
	fon := FONT{}
	sizes := []int{16, 18, 24, 36, 48}
	t := strings.Split(standardCharacters, "")
	for i := range sizes {
		f, ERR := ttf.OpenFont(FONT1DEFAULTPATH, sizes[i])
		if ERR != nil {
			mERR(ERR)
		}
		for j := range t {
			switch i {
			case 0: //sizeSMLR
				fon.charSMLR = append(fon.charSMLR, mCHAR(t[j], f))
				fon.sizeSMLR = sizes[i]
				fon.smlrH = float32(fon.charSMLR[0].r.H)
			case 1: //sizeSML
				fon.charSML = append(fon.charSML, mCHAR(t[j], f))
				fon.sizeSML = sizes[i]
				fon.smlH = float32(fon.charSML[0].r.H)
			case 2: //sizeMED
				fon.charMED = append(fon.charMED, mCHAR(t[j], f))
				fon.sizeMED = sizes[i]
				fon.medH = float32(fon.charMED[0].r.H)
			case 3: //sizeLRG
				fon.charLRG = append(fon.charLRG, mCHAR(t[j], f))
				fon.sizeLRG = sizes[i]
				fon.lrgH = float32(fon.charLRG[0].r.H)
			case 4: //sizeLRGR
				fon.charLRGR = append(fon.charLRGR, mCHAR(t[j], f))
				fon.sizeLRGR = sizes[i]
				fon.lrgrH = float32(fon.charLRGR[0].r.H)
			}
		}
	}
	FONT1DEFAULT = fon
	fon = FONT{}

	for i := range sizes {
		f, ERR := ttf.OpenFont(FONT2DEFAULTPATH, sizes[i])
		if ERR != nil {
			mERR(ERR)
		}
		for j := range t {
			switch i {
			case 0: //sizeSMLR
				fon.charSMLR = append(fon.charSMLR, mCHAR(t[j], f))
				fon.sizeSMLR = sizes[i]
				fon.smlrH = float32(fon.charSMLR[0].r.H)
			case 1: //sizeSML
				fon.charSML = append(fon.charSML, mCHAR(t[j], f))
				fon.sizeSML = sizes[i]
				fon.smlH = float32(fon.charSML[0].r.H)
			case 2: //sizeMED
				fon.charMED = append(fon.charMED, mCHAR(t[j], f))
				fon.sizeMED = sizes[i]
				fon.medH = float32(fon.charMED[0].r.H)
			case 3: //sizeLRG
				fon.charLRG = append(fon.charLRG, mCHAR(t[j], f))
				fon.sizeLRG = sizes[i]
				fon.lrgH = float32(fon.charLRG[0].r.H)
			case 4: //sizeLRGR
				fon.charLRGR = append(fon.charLRGR, mCHAR(t[j], f))
				fon.sizeLRGR = sizes[i]
				fon.lrgrH = float32(fon.charLRGR[0].r.H)
			}
		}
	}
	FONT2DEFAULT = fon
}

func mCHAR(c string, f *ttf.Font) CHAR {
	var w, h int
	TXSURF, _ = f.RenderUTF8Blended(c, WHITE())
	defer TXSURF.Free()
	w, h, _ = f.SizeUTF8(c)
	t := CHAR{}
	t.tex, _ = RND.CreateTextureFromSurface(TXSURF)
	t.tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	t.r = sdl.Rect{0, 0, int32(w), int32(h)}
	t.ch = c
	return t
}
