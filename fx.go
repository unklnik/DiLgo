package main

import "github.com/veandco/go-sdl2/sdl"

func PIXELNOISEPLUS(numPixels, maxWpixelMAX10 int, c sdl.Color, randomAlpha bool) {
	if maxWpixelMAX10 > 10 || maxWpixelMAX10 < 1 {
		Mmsg("ERROR: func PIXELNOISEPLUS: maxWpixelMAX10 must be from 1 to 10 >> Set to 5")
		maxWpixelMAX10 = 5
	}
	for range numPixels {
		if randomAlpha {
			a := uint8(RINT(1, 256))
			c = COLORALPHA(c, a)
		}
		x, y := RF32(0, float32(WINW)), RF32(0, float32(WINH))
		w := RF32(1, float32(maxWpixelMAX10))
		r := sdl.FRect{x, y, w, w}
		DrecFill(r, c)
	}

}

func PIXELNOISERANDOMCOLOR(numPixels int) {
	for range numPixels {
		COL(COLORǁRANDOM())
		RND.DrawPoint(int32(RINT(0, int(WINW))), int32(RINT(0, int(WINH))))
	}
}

func PIXELNOISE(numPixels int, c sdl.Color) {
	COL(c)
	for range numPixels {
		RND.DrawPoint(int32(RINT(0, int(WINW))), int32(RINT(0, int(WINH))))
	}
}

func SCANLINES(lineThickness, lineSpace int, lineAlpha1to255 uint8, c sdl.Color) {
	var x, y float32 = 0, 0
	c = COLORALPHA(c, lineAlpha1to255)
	for y < float32(WINH) {
		for range lineThickness {
			p1 := sdl.FPoint{x, y}
			p2 := p1
			p2.X += float32(WINW)
			Dline(p1, p2, c)
			y++
		}
		y += float32(lineSpace)
	}
}
