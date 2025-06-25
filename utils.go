package main

import (
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type POINT64 struct {
	X, Y float64
}

// MARK: IMAGES
func RESIZEIMRECSCALEHEIGHT(im IM, w float32) sdl.FRect {
	h := (w / float32(im.r.W)) * float32(im.r.H)
	return sdl.FRect{float32(im.r.X), float32(im.r.Y), w, h}
}

// MARK: FILES
func FimageFilesJPGandPNG(directoryPath string) []string {
	// Open the directory and list its entries.
	entries, ERR := os.ReadDir(directoryPath)
	if ERR != nil {
		mERR(ERR)
	}

	var imagePaths []string
	for _, entry := range entries {
		// Skip if the entry is a directory.
		if entry.IsDir() {
			continue
		}
		// Convert the filename to lowercase for case-insensitive comparison.
		lowerName := strings.ToLower(entry.Name())
		// Check if the file ends with .jpg or .png.
		if strings.HasSuffix(lowerName, ".jpg") || strings.HasSuffix(lowerName, ".png") {
			fullPath := filepath.Join(directoryPath, entry.Name())
			imagePaths = append(imagePaths, fullPath)
		}
	}
	return imagePaths
}

// MARK: COLORS
func INTERPOLATECOLOR(c1, c2 sdl.Color, factor float64) sdl.Color {
	return sdl.Color{
		R: uint8(float64(c1.R) + (float64(c2.R)-float64(c1.R))*factor),
		G: uint8(float64(c1.G) + (float64(c2.G)-float64(c1.G))*factor),
		B: uint8(float64(c1.B) + (float64(c2.B)-float64(c1.B))*factor),
		A: uint8(float64(c1.A) + (float64(c2.A)-float64(c1.A))*factor),
	}
}

// MARK: ISO
func ZINDEXSORTCUBE(c []CUBE) []CUBE {
	sort.Slice(c, func(i, j int) bool { return c[i].zindex < c[j].zindex })
	return c
}
func ZINDEXSORTISOREC(ir []ISOREC) []ISOREC {
	sort.Slice(ir, func(i, j int) bool { return ir[i].zindex > ir[j].zindex })
	return ir
}

// MARK: MOVEMENT
func SPEEDXY2POINTS(p1, p2 sdl.FPoint, maxspeed float32) (x float32, y float32) {
	x = p2.X - p1.X
	y = p2.Y - p1.Y
	distance := float32(math.Sqrt(float64(x*x + y*y)))
	if distance > maxspeed {
		scale := maxspeed / distance
		x *= scale
		y *= scale
	}
	return x, y
}

// MARK: MATH
func ABS32(n float32) float32 {
	return float32(math.Abs(float64(n)))
}
func ABS64(n float64) float64 {
	return math.Abs(n)
}
func ROUNDFLOAT(val float32, precision int) float32 {
	ratio := math.Pow(10, float64(precision))
	return float32(math.Round(float64(val)*ratio) / ratio)
}
func ABSDIFF(num1, num2 float32) float32 {
	return float32(math.Abs(float64(num1 - num2)))
}
func RECOUTERCIRCRADIUS(wSide float32) float32 {
	return float32(float64(wSide) / math.Sqrt(2))
}

// MARK: TIME
func INT64TODURATION(ms int64) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

// MARK: TEXTURES
func TEXCOL(tex *sdl.Texture, c sdl.Color) *sdl.Texture {
	tex.SetColorMod(c.R, c.G, c.B)
	return tex
}
func TEXCOLREVERT(tex *sdl.Texture) *sdl.Texture {
	tex.SetColorMod(255, 255, 255)
	return tex
}
func TEXALPHA(tex *sdl.Texture, a uint8) *sdl.Texture {
	tex.SetAlphaMod(a)
	return tex
}
func TEXALPHAREVERT(tex *sdl.Texture) *sdl.Texture {
	tex.SetAlphaMod(255)
	return tex
}

// MARK: RANDOM NUMBERS
func RINT(min, max int) int {
	return min + rand.IntN(max-min)
}
func RF32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func FLIPCOIN() bool {
	return rand.IntN(2) == 0
}
func RF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func RUINT8(min, max int) uint8 {
	return uint8(min + rand.IntN(max-min))
}

// MARK: TRIANGLES
func TRICENTER(p []sdl.FPoint) sdl.FPoint {
	return sdl.FPoint{
		(p[0].X + p[1].X + p[2].X) / 3,
		(p[0].Y + p[1].Y + p[2].Y) / 3,
	}
}
func TRIPOINT2VERT(t TRI) []sdl.Vertex {
	var v, v2, v3 sdl.Vertex
	v.Position = t.vert[0]
	v2.Position = t.vert[1]
	v3.Position = t.vert[2]
	if t.c1 != BLANKCOL {
		v.Color = t.c1
		v2.Color = t.c2
		v3.Color = t.c3
	} else {
		v.Color = t.c
		v2.Color = t.c
		v3.Color = t.c
	}
	return []sdl.Vertex{v, v2, v3}
}
func TRIPOINT2VERTMULTICOLOR(t TRI) []sdl.Vertex {
	v := sdl.Vertex{}
	v.Position = t.vert[0]
	v.Color = t.c1
	v2 := sdl.Vertex{}
	v2.Position = t.vert[1]
	v2.Color = t.c2
	v3 := sdl.Vertex{}
	v3.Position = t.vert[2]
	v3.Color = t.c3
	return []sdl.Vertex{v, v2, v3}
}
func TRICOLORFILLALPHA(t TRI, c sdl.Color, alpha uint8) TRI {
	t.c = c
	t.c = COLORALPHA(t.c, alpha)
	return t
}
func TRICOLOR(t TRI, cFill, cLine sdl.Color) TRI {
	t.c = cFill
	t.cLine = cLine
	return t
}

// MARK: POINTS
func ORIGIN(r sdl.FRect) *sdl.FPoint {
	return &sdl.FPoint{r.W / 2, r.H / 2}
}
func POINTRECCENTER(r sdl.FRect) sdl.FPoint {
	return sdl.FPoint{r.X + r.W/2, r.Y + r.H/2}
}
func POINTRECCENTERXY(x, y, w, h float32) sdl.FPoint {
	return sdl.FPoint{x + w/2, y + h/2}
}
func POINT2POINT64(p sdl.FPoint) POINT64 {
	p2 := POINT64{float64(p.X), float64(p.Y)}
	return p2
}
func POINT642POINT(p POINT64) sdl.FPoint {
	p2 := sdl.FPoint{float32(p.X), float32(p.Y)}
	return p2
}

// MARK: RECS

func MOVERECTOCENTER(r sdl.FRect, cnt sdl.FPoint) sdl.FRect {
	r.X = cnt.X - r.W/2
	r.Y = cnt.Y - r.H/2
	return r
}
func RESIZERECSCALEHEIGHT(origWidth, origHeight, newWidth float32) float32 {
	return (newWidth / origWidth) * origHeight
}

func FREC2REC(r sdl.FRect) *sdl.Rect {
	return &sdl.Rect{int32(r.X), int32(r.Y), int32(r.W), int32(r.H)}
}
func RECPOINTS(r sdl.FRect) []sdl.FPoint {
	var p []sdl.FPoint
	p = append(p, sdl.FPoint{r.X, r.Y})
	p = append(p, sdl.FPoint{r.X + r.W, r.Y})
	p = append(p, sdl.FPoint{r.X + r.W, r.Y + r.H})
	p = append(p, sdl.FPoint{r.X, r.Y + r.H})
	return p
}
func POINTINREC(p sdl.FPoint, r sdl.FRect) bool {
	return p.X >= r.X && p.X <= r.X+r.W && p.Y >= r.Y && p.Y <= r.Y+r.H
}

// MARK: COLLISIONS INTERSECTIONS
func CpointRecBorderCollis(p sdl.FPoint, r sdl.FRect, marginTotal float32) bool {
	collides := false
	rL := sdl.FRect{r.X - marginTotal/2, r.Y, marginTotal, r.H}
	rR := sdl.FRect{r.X + r.W - marginTotal/2, r.Y, marginTotal, r.H}
	rT := sdl.FRect{r.X - marginTotal/2, r.Y - marginTotal/2, r.W + marginTotal, marginTotal}
	rB := sdl.FRect{r.X - marginTotal/2, r.Y + r.H - marginTotal/2, r.W + marginTotal, marginTotal}
	if CrecPoint(MOUSE, rL) || CrecPoint(MOUSE, rR) || CrecPoint(MOUSE, rT) || CrecPoint(MOUSE, rB) {
		collides = true
	}
	return collides
}
func CrecRec(r1, r2 sdl.FRect) bool {
	return r1.X < r2.X+r2.W &&
		r1.X+r1.W > r2.X &&
		r1.Y < r2.Y+r2.H &&
		r1.Y+r1.H > r2.Y
}
func CrecPoint(p sdl.FPoint, r sdl.FRect) bool {
	return p.X >= r.X && p.X <= r.X+r.W && p.Y >= r.Y && p.Y <= r.Y+r.H
}
func CstarPoint(p sdl.FPoint, s STAR) bool {
	collides := false
	for i := range s.tri {
		collides = CtriPoint(p, s.tri[i])
		if collides {
			break
		}
	}
	return collides
}
func CcircleCenterPoint(p, cnt sdl.FPoint, radius float32) bool {
	distance := math.Sqrt(math.Pow(float64(p.X-cnt.X), 2) + math.Pow(float64(p.Y-cnt.Y), 2))
	return distance <= float64(radius)
}
func CcirclePoint(p sdl.FPoint, c CIRC) bool {
	distance := math.Sqrt(math.Pow(float64(p.X-c.cnt.X), 2) + math.Pow(float64(p.Y-c.cnt.Y), 2))
	return distance <= float64(c.radius)
}
func CpolygonPoint(p sdl.FPoint, poly POLYGON) bool {
	collides := false
	for i := range poly.tri {
		collides = CtriPoint(p, poly.tri[i])
		if collides {
			break
		}
	}
	return collides
}
func CtriPoint(p sdl.FPoint, t TRI) bool {
	collides := false
	a := POINT2POINT64(t.vert[0])
	b := POINT2POINT64(t.vert[1])
	c := POINT2POINT64(t.vert[2])
	p2 := POINT2POINT64(p)
	collides = pointInTriangle(p2, a, b, c)
	return collides
}
func pointInTriangle(p POINT64, a POINT64, b POINT64, c POINT64) bool {
	areaOriginal := triangleArea(a, b, c)
	area1 := triangleArea(p, b, c)
	area2 := triangleArea(a, p, c)
	area3 := triangleArea(a, b, p)
	return math.Abs((area1+area2+area3)-areaOriginal) < 1e-9
}
func triangleArea(a, b, c POINT64) float64 {
	return math.Abs((a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2.0)
}
