package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var ()

type PIXL struct {
	xy sdl.FPoint
	c  sdl.Color
}
type TRI struct {
	c, cLine, c1, c2, c3 sdl.Color
	vert                 []sdl.FPoint
	cnt                  sdl.FPoint
}
type POLYGON struct {
	c, c1, c2, c3 sdl.Color
	vert          []sdl.FPoint
	cnt           sdl.FPoint
	tri           []TRI
}
type GRID struct {
	c          sdl.Color
	r          []sdl.FRect
	nm         string
	w, h       float32
	cnt        sdl.FPoint
	tiles      []IM
	numW, numH int
}
type CIRC struct {
	c        sdl.Color
	radius   float32
	cnt      sdl.FPoint
	gradient []PIXL
}
type STAR struct {
	c   sdl.Color
	cnt sdl.FPoint
	p   []POINT64
	tri []TRI
}
type CUBE struct {
	recs                   []ISOREC
	frontCorner            sdl.FPoint
	cntBot, cntMid, cntTop sdl.FPoint
	cLine, cFill           sdl.Color
	zindex                 int
	im                     IM
	wSide, h               float32
}
type ISOGRID struct {
	recs                 []ISOREC
	cubes                []CUBE
	frontCorner, cnt     sdl.FPoint
	name                 string
	w, h, wSide          float32
	numWbloks, numLbloks int
}
type ISOREC struct {
	tri                  []TRI
	vert                 []sdl.FPoint
	zindex               int
	c, cLine, c1, c2, c3 sdl.Color
	cnt                  sdl.FPoint
	solid                bool
	imSurface            IM
	imAdditional         []IM
	wSide                float32
}

// MARK: ISO
func MborderCubesGridIsometric(ig ISOGRID, hCube float32, cFill, cLine sdl.Color) ISOGRID {
	for i := range ig.recs {
		if i < ig.numWbloks {
			ig.cubes = append(ig.cubes, Mcube(ig.recs[i].vert[0], ig.wSide, hCube, cFill, cLine))
		}
		if i > len(ig.recs)-(ig.numWbloks+1) {
			ig.cubes = append(ig.cubes, Mcube(ig.recs[i].vert[0], ig.wSide, hCube, cFill, cLine))
		}
		if i%ig.numWbloks == 0 {
			ig.cubes = append(ig.cubes, Mcube(ig.recs[i].vert[0], ig.wSide, hCube, cFill, cLine))
			if i-ig.numWbloks > 0 && i-ig.numWbloks < len(ig.recs)-(ig.numWbloks+1) {
				ig.cubes = append(ig.cubes, Mcube(ig.recs[i-1].vert[0], ig.wSide, hCube, cFill, cLine))
			}
		}
	}
	return ig
}
func MgridIsometricCubesRandomHeights(frontCorner sdl.FPoint, widthBlockSide, heightCubeMin, heightCubeMax float32, numBlocksWidth, numBlocksLength int, cFillGrid, cLineGrid, cFillCube, cLineCube sdl.Color) ISOGRID {
	ig := MgridIsometric(frontCorner, widthBlockSide, numBlocksWidth, numBlocksLength, cFillGrid, cLineGrid)
	ig.numWbloks = numBlocksWidth
	ig.numLbloks = numBlocksLength
	ig.wSide = widthBlockSide
	for i := range ig.recs {
		heightCube := RF32(heightCubeMin, heightCubeMax)
		c := Mcube(ig.recs[i].vert[0], widthBlockSide, heightCube, cFillCube, cLineCube)
		c.zindex = ig.recs[i].zindex
		ig.cubes = append(ig.cubes, c)
	}
	return ig
}
func MgridIsometricCubesRandomHeightsBelow(frontCorner sdl.FPoint, widthBlockSide, heightCubeMin, heightCubeMax float32, numBlocksWidth, numBlocksLength int, cFillGrid, cLineGrid, cFillCube, cLineCube sdl.Color) ISOGRID {
	ig := MgridIsometric(frontCorner, widthBlockSide, numBlocksWidth, numBlocksLength, cFillGrid, cLineGrid)
	ig.numWbloks = numBlocksWidth
	ig.numLbloks = numBlocksLength
	ig.wSide = widthBlockSide
	for i := range ig.recs {
		heightCube := RF32(heightCubeMin, heightCubeMax)
		v := ig.recs[i].vert[0]
		v.Y += heightCube
		c := Mcube(v, widthBlockSide, heightCube, cFillCube, cLineCube)
		c.zindex = ig.recs[i].zindex
		ig.cubes = append(ig.cubes, c)
	}
	return ig
}
func MgridIsometricCubes(frontCorner sdl.FPoint, widthBlockSide, heightCube float32, numBlocksWidth, numBlocksLength int, cFillGrid, cLineGrid, cFillCube, cLineCube sdl.Color) ISOGRID {
	ig := MgridIsometric(frontCorner, widthBlockSide, numBlocksWidth, numBlocksLength, cFillGrid, cLineGrid)
	ig.numWbloks = numBlocksWidth
	ig.numLbloks = numBlocksLength
	ig.wSide = widthBlockSide
	for i := range ig.recs {
		c := Mcube(ig.recs[i].vert[0], widthBlockSide, heightCube, cFillCube, cLineCube)
		c.zindex = ig.recs[i].zindex
		ig.cubes = append(ig.cubes, c)
	}
	return ig
}
func MgridIsometricCubesBelow(frontCorner sdl.FPoint, widthBlockSide, heightCube float32, numBlocksWidth, numBlocksLength int, cFillGrid, cLineGrid, cFillCube, cLineCube sdl.Color) ISOGRID {
	ig := MgridIsometric(frontCorner, widthBlockSide, numBlocksWidth, numBlocksLength, cFillGrid, cLineGrid)
	ig.numWbloks = numBlocksWidth
	ig.numLbloks = numBlocksLength
	ig.wSide = widthBlockSide
	for i := range ig.recs {
		v := ig.recs[i].vert[0]
		v.Y += heightCube
		c := Mcube(v, widthBlockSide, heightCube, cFillCube, cLineCube)
		c.zindex = ig.recs[i].zindex
		ig.cubes = append(ig.cubes, c)
	}
	return ig
}
func MgridIsometric(frontCorner sdl.FPoint, widthBlockSide float32, numBlocksWidth, numBlocksLength int, cFill, cLine sdl.Color) ISOGRID {
	ig := ISOGRID{}
	ig.numWbloks = numBlocksWidth
	ig.numLbloks = numBlocksLength
	ig.wSide = widthBlockSide
	w := widthBlockSide * float32(numBlocksWidth*2)
	h := widthBlockSide * float32(numBlocksWidth)
	ig.w = w
	ig.h = h
	ig.frontCorner = frontCorner
	ig.cnt = sdl.FPoint{ig.frontCorner.X, ig.frontCorner.Y - h/2}
	x := frontCorner.X
	y := frontCorner.Y
	ox := x
	oy := y
	a := numBlocksLength * numBlocksWidth
	zi := 0
	oz := zi
	c := 0
	for a > 0 {
		ir := MisoRec(sdl.FPoint{x, y}, widthBlockSide, cFill, cLine)
		ir.zindex = zi
		ig.recs = append(ig.recs, ir)
		x -= widthBlockSide
		y -= widthBlockSide / 2
		c++
		zi++
		a--
		if c == numBlocksWidth {
			c = 0
			x = ox
			y = oy
			x += widthBlockSide
			y -= widthBlockSide / 2
			zi = oz
			zi++
			oz = zi
			ox = x
			oy = y
		}
	}
	return ig
}
func MgridIsometricCenter(gridCenter sdl.FPoint, widthBlockSide float32, numBlocksWidth, numBlocksLength int, cFill, cLine sdl.Color) ISOGRID {
	ig := ISOGRID{}
	ig.numWbloks = numBlocksWidth
	ig.numLbloks = numBlocksLength
	ig.wSide = widthBlockSide
	w := widthBlockSide * float32(numBlocksWidth*2)
	h := widthBlockSide * float32(numBlocksWidth)
	ig.w = w
	ig.h = h
	ig.cnt = gridCenter
	ig.frontCorner = sdl.FPoint{gridCenter.X, gridCenter.Y + h/2}
	x := ig.frontCorner.X
	y := ig.frontCorner.Y
	ox := x
	oy := y
	a := numBlocksLength * numBlocksWidth
	zi := 0
	oz := zi
	c := 0
	for a > 0 {
		ir := MisoRec(sdl.FPoint{x, y}, widthBlockSide, cFill, cLine)
		ir.zindex = zi
		ig.recs = append(ig.recs, ir)
		x -= widthBlockSide
		y -= widthBlockSide / 2
		c++
		zi++
		a--
		if c == numBlocksWidth {
			c = 0
			x = ox
			y = oy
			x += widthBlockSide
			y -= widthBlockSide / 2
			zi = oz
			zi++
			oz = zi
			ox = x
			oy = y
		}
	}
	return ig
}
func MisoRec(frontCorner sdl.FPoint, widthSide float32, cFill, cLine sdl.Color) ISOREC {
	x := frontCorner.X
	y := frontCorner.Y
	ir := ISOREC{}
	ir.wSide = widthSide
	ir.cnt = sdl.FPoint{x, y - widthSide/2}
	ir.c = cFill
	ir.cLine = cLine
	ir.vert = append(ir.vert, sdl.FPoint{x, y}, sdl.FPoint{x - widthSide, y - widthSide/2}, sdl.FPoint{x, y - widthSide}, sdl.FPoint{x + widthSide, y - widthSide/2})
	t1 := TRI{}
	t1.vert = append(t1.vert, ir.vert[0], ir.vert[1], ir.vert[2])
	t1.cnt = TRICENTER(t1.vert)
	t1.c = cFill
	t1.cLine = cLine
	ir.tri = append(ir.tri, t1)
	t2 := TRI{}
	t2.vert = append(t2.vert, ir.vert[0], ir.vert[2], ir.vert[3])
	t2.cnt = TRICENTER(t2.vert)
	t2.c = cFill
	t2.cLine = cLine
	ir.tri = append(ir.tri, t2)
	return ir
}
func MisoRecMultiColor(frontCorner sdl.FPoint, widthSide float32, c1, c2, c3, cLine sdl.Color) ISOREC {
	x := frontCorner.X
	y := frontCorner.Y
	ir := ISOREC{}
	ir.cnt = sdl.FPoint{x, y - widthSide/2}
	ir.cLine = cLine
	ir.vert = append(ir.vert, sdl.FPoint{x, y}, sdl.FPoint{x - widthSide, y - widthSide/2}, sdl.FPoint{x, y - widthSide}, sdl.FPoint{x + widthSide, y - widthSide/2})
	ir.tri = append(ir.tri, MtriPointsMultiColor([]sdl.FPoint{ir.vert[0], ir.vert[1], ir.vert[2]}, c1, c2, c3, cLine))
	ir.tri = append(ir.tri, MtriPointsMultiColor([]sdl.FPoint{ir.vert[0], ir.vert[2], ir.vert[3]}, c1, c3, c2, cLine))
	return ir
}

// MARK: CUBES
func Mcube(frontCorner sdl.FPoint, widthSide, height float32, cFill, cLine sdl.Color) CUBE {
	c := CUBE{}
	c.wSide = widthSide
	c.h = height
	c.cFill = cFill
	c.cLine = cLine
	c.frontCorner = frontCorner
	c.cntBot = sdl.FPoint{c.frontCorner.X, frontCorner.Y - widthSide/2}
	c.cntMid = c.cntBot
	c.cntMid.Y -= height / 2
	c.cntTop = c.cntBot
	c.cntTop.Y -= height
	ir := MisoRec(frontCorner, widthSide, cFill, cLine) //BOTTOM
	c.recs = append(c.recs, ir)
	f2 := frontCorner
	f2.Y -= height
	ir = MisoRec(f2, widthSide, cFill, cLine) //TOP
	c.recs = append(c.recs, ir)
	ir = ISOREC{} //LEFT FRONT
	ir.c = cFill
	ir.cLine = cLine
	v1 := frontCorner
	v2 := sdl.FPoint{frontCorner.X - widthSide, frontCorner.Y - widthSide/2}
	v3 := v2
	v3.Y -= height
	v4 := v1
	v4.Y -= height
	ir.vert = append(ir.vert, v1, v2, v3, v4)
	t := TRI{}
	t.c = cFill
	t.cLine = cLine
	t.vert = append(t.vert, ir.vert[0], ir.vert[1], ir.vert[2])
	ir.tri = append(ir.tri, t)
	t = TRI{}
	t.c = cFill
	t.cLine = cLine
	t.vert = append(t.vert, ir.vert[0], ir.vert[2], ir.vert[3])
	ir.tri = append(ir.tri, t)
	c.recs = append(c.recs, ir)
	ir = ISOREC{} //RIGHT FRONT
	ir.c = cFill
	ir.cLine = cLine
	v1 = frontCorner
	v2 = v1
	v2.Y -= height
	v4 = sdl.FPoint{frontCorner.X + widthSide, frontCorner.Y - widthSide/2}
	v3 = v4
	v3.Y -= height
	ir.vert = append(ir.vert, v1, v2, v3, v4)
	t = TRI{}
	t.c = cFill
	t.cLine = cLine
	t.vert = append(t.vert, ir.vert[0], ir.vert[1], ir.vert[2])
	ir.tri = append(ir.tri, t)
	t = TRI{}
	t.c = cFill
	t.cLine = cLine
	t.vert = append(t.vert, ir.vert[0], ir.vert[2], ir.vert[3])
	ir.tri = append(ir.tri, t)
	c.recs = append(c.recs, ir)
	ir = ISOREC{} //LEFT BACK
	ir.c = cFill
	ir.cLine = cLine
	v1 = c.recs[0].vert[1]
	v2 = v1
	v2.Y -= height
	v4 = sdl.FPoint{v1.X + widthSide, v1.Y - widthSide/2}
	v3 = v4
	v3.Y -= height
	ir.vert = append(ir.vert, v1, v2, v3, v4)
	t = TRI{}
	t.c = cFill
	t.cLine = cLine
	t.vert = append(t.vert, ir.vert[0], ir.vert[1], ir.vert[2])
	ir.tri = append(ir.tri, t)
	t = TRI{}
	t.c = cFill
	t.cLine = cLine
	t.vert = append(t.vert, ir.vert[0], ir.vert[2], ir.vert[3])
	ir.tri = append(ir.tri, t)
	c.recs = append(c.recs, ir)
	ir = ISOREC{} //RIGHT BACK
	ir.c = cFill
	ir.cLine = cLine
	v1 = c.recs[0].vert[2]
	v2 = v1
	v2.Y -= height
	v4 = sdl.FPoint{v1.X + widthSide, v1.Y + widthSide/2}
	v3 = v4
	v3.Y -= height
	ir.vert = append(ir.vert, v1, v2, v3, v4)
	t = TRI{}
	t.c = cFill
	t.cLine = cLine
	t.vert = append(t.vert, ir.vert[0], ir.vert[1], ir.vert[2])
	ir.tri = append(ir.tri, t)
	t = TRI{}
	t.c = cFill
	t.cLine = cLine
	t.vert = append(t.vert, ir.vert[0], ir.vert[2], ir.vert[3])
	ir.tri = append(ir.tri, t)
	c.recs = append(c.recs, ir)
	return c

}

// MARK: STARS
func Mstar(cnt sdl.FPoint, numPoints int, outerRadius, innerRadius float32, rotation float64, c sdl.Color) STAR {
	s := STAR{}
	s.c = c
	s.cnt = cnt
	if numPoints < 2 {
		panic("The number of points must be at least 2")
	}
	var points []POINT64
	angleStep := math.Pi / float64(numPoints)
	for i := range numPoints * 2 {
		var radius float64
		if i%2 == 0 {
			radius = float64(outerRadius)
		} else {
			radius = float64(innerRadius)
		}

		// Add the rotation angle to the base angle
		angle := float64(i)*angleStep + rotation
		x := float64(cnt.X) + radius*math.Cos(angle)
		y := float64(cnt.Y) + radius*math.Sin(angle)

		points = append(points, POINT64{x, y})
	}
	s.p = points
	for i := 1; i < len(points)-2; i += 2 {
		t := TRI{}
		t.c = c
		t.vert = append(t.vert, POINT642POINT(points[i]))
		t.vert = append(t.vert, POINT642POINT(points[i+1]))
		t.vert = append(t.vert, POINT642POINT(points[i+2]))
		s.tri = append(s.tri, t)
		t = TRI{}
		t.c = c
		t.vert = append(t.vert, POINT642POINT(points[i]))
		t.vert = append(t.vert, POINT642POINT(points[i+2]))
		t.vert = append(t.vert, cnt)
		s.tri = append(s.tri, t)
	}

	t := TRI{}
	t.c = c
	t.vert = append(t.vert, POINT642POINT(points[len(points)-1]))
	t.vert = append(t.vert, POINT642POINT(points[0]))
	t.vert = append(t.vert, POINT642POINT(points[1]))
	s.tri = append(s.tri, t)
	t = TRI{}
	t.c = c
	t.vert = append(t.vert, POINT642POINT(points[len(points)-1]))
	t.vert = append(t.vert, POINT642POINT(points[1]))
	t.vert = append(t.vert, cnt)
	s.tri = append(s.tri, t)

	return s
}

// MARK: CIRCLES
func Mcirc(cnt sdl.FPoint, radius float32, c sdl.Color) CIRC {
	circ := CIRC{}
	circ.radius = radius
	circ.c = c
	circ.cnt = cnt
	return circ
}

// MARK: RECS
func MrecCenter(cnt sdl.FPoint, w, h float32) sdl.FRect {
	return sdl.FRect{cnt.X - w/2, cnt.Y - h/2, w, h}
}
func MrecXY(x, y, w, h float32) sdl.FRect {
	return sdl.FRect{x, y, w, h}
}
func MsquareCenter(cnt sdl.FPoint, w float32) sdl.FRect {
	return sdl.FRect{cnt.X - w/2, cnt.Y - w/2, w, w}
}

// MARK: POLYGONS
func Mpolygon(cnt sdl.FPoint, radius float64, sides int, rotation float64, c sdl.Color) POLYGON {
	p := POLYGON{}
	p.c = c
	p.cnt = cnt
	verts := polygonVertices(float64(cnt.X), float64(cnt.Y), radius, sides, rotation)
	for i := range verts {
		p.vert = append(p.vert, sdl.FPoint{float32(verts[i].X), float32(verts[i].Y)})
	}
	p.tri = poly2TRI(cnt, p.vert, c)
	return p
}
func MpolygonMultiColor(cnt sdl.FPoint, radius float64, sides int, rotation float64, c1, c2, c3 sdl.Color) POLYGON {
	p := POLYGON{}
	p.c1 = c1
	p.c2 = c2
	p.c3 = c3
	p.cnt = cnt
	verts := polygonVertices(float64(cnt.X), float64(cnt.Y), radius, sides, rotation)
	for i := range verts {
		p.vert = append(p.vert, sdl.FPoint{float32(verts[i].X), float32(verts[i].Y)})
	}
	p.tri = poly2TRImultiColor(cnt, p.vert, c1, c2, c3)
	return p
}
func poly2TRI(cnt sdl.FPoint, p []sdl.FPoint, c sdl.Color) []TRI {
	var t []TRI
	for i := range len(p) - 1 {
		t2 := TRI{}
		t2.c = c
		t2.vert = append(t2.vert, cnt)
		t2.vert = append(t2.vert, p[i])
		t2.vert = append(t2.vert, p[i+1])
		t = append(t, t2)
	}
	t2 := TRI{}
	t2.c = c
	t2.vert = append(t2.vert, cnt)
	t2.vert = append(t2.vert, p[len(p)-1])
	t2.vert = append(t2.vert, p[0])
	t = append(t, t2)
	return t
}
func poly2TRImultiColor(cnt sdl.FPoint, p []sdl.FPoint, c1, c2, c3 sdl.Color) []TRI {
	var t []TRI
	for i := range len(p) - 1 {
		t2 := TRI{}
		t2.c1 = c1
		t2.c2 = c2
		t2.c3 = c3
		t2.vert = append(t2.vert, cnt)
		t2.vert = append(t2.vert, p[i])
		t2.vert = append(t2.vert, p[i+1])
		t = append(t, t2)
	}
	t2 := TRI{}
	t2.c1 = c1
	t2.c2 = c2
	t2.c3 = c3
	t2.vert = append(t2.vert, cnt)
	t2.vert = append(t2.vert, p[len(p)-1])
	t2.vert = append(t2.vert, p[0])
	t = append(t, t2)
	return t
}

func polygonVertices(centerX, centerY, radius float64, sides int, rotationAngle float64) []struct{ X, Y float64 } {
	vertices := make([]struct{ X, Y float64 }, sides)
	angleStep := 2 * math.Pi / float64(sides)

	// Rotation angle in radians
	rotation := rotationAngle * math.Pi / 180

	for i := 0; i < sides; i++ {
		// Compute angle for each vertex
		angle := angleStep*float64(i) + rotation
		x := centerX + radius*math.Cos(angle)
		y := centerY + radius*math.Sin(angle)
		vertices[i] = struct{ X, Y float64 }{X: x, Y: y}
	}
	return vertices
}

// MARK: TRIANGLES
func MtriPointsMultiColor(p []sdl.FPoint, c1, c2, c3, cLine sdl.Color) TRI {
	t := TRI{}
	if len(p) != 3 {
		Mmsg("ERROR: func MtriPoints: len(p) must equal 3 >> Will return empty TRI")
	} else {
		t.c1 = c1
		t.c2 = c2
		t.c3 = c3
		t.cLine = cLine
		t.vert = append(t.vert, p...)
		t.cnt = TRICENTER(p)
	}
	return t
}
func MtriPoints(p []sdl.FPoint, cFill, cLine sdl.Color) TRI {
	t := TRI{}
	if len(p) != 3 {
		Mmsg("ERROR: func MtriPoints: len(p) must equal 3 >> Will return empty TRI")
	} else {
		t.c = cFill
		t.cLine = cLine
		t.vert = append(t.vert, p...)
		t.cnt = TRICENTER(p)
	}
	return t
}
func MtriCenter(cnt sdl.FPoint, radius float64, rotation float64, cFill, cLine sdl.Color) TRI {
	t := TRI{}
	t.cnt = cnt
	t.c = cFill
	t.cLine = cLine
	verts := calculateTriangleVertices(float64(cnt.X), float64(cnt.Y), radius, rotation)
	p := sdl.FPoint{}
	p.X = float32(verts[0][0])
	p.Y = float32(verts[0][1])
	t.vert = append(t.vert, p)
	p.X = float32(verts[1][0])
	p.Y = float32(verts[1][1])
	t.vert = append(t.vert, p)
	p.X = float32(verts[2][0])
	p.Y = float32(verts[2][1])
	t.vert = append(t.vert, p)
	return t
}
func MtriCenterMultiColor(cnt sdl.FPoint, radius float64, rotation float64, c1, c2, c3, cLine sdl.Color) TRI {
	t := TRI{}
	t.cnt = cnt
	t.c1 = c1
	t.c2 = c2
	t.c3 = c3
	t.cLine = cLine
	verts := calculateTriangleVertices(float64(cnt.X), float64(cnt.Y), radius, rotation)
	p := sdl.FPoint{}
	p.X = float32(verts[0][0])
	p.Y = float32(verts[0][1])
	t.vert = append(t.vert, p)
	p.X = float32(verts[1][0])
	p.Y = float32(verts[1][1])
	t.vert = append(t.vert, p)
	p.X = float32(verts[2][0])
	p.Y = float32(verts[2][1])
	t.vert = append(t.vert, p)
	return t
}

func calculateTriangleVertices(centerX, centerY, radius, rotationAngle float64) [3][2]float64 {
	var vertices [3][2]float64
	angles := []float64{0, 2 * math.Pi / 3, 4 * math.Pi / 3}
	rotationRadians := rotationAngle * (math.Pi / 180.0)
	for i, angle := range angles {
		rotatedAngle := angle + rotationRadians
		vertices[i][0] = centerX + radius*math.Cos(rotatedAngle)
		vertices[i][1] = centerY + radius*math.Sin(rotatedAngle)
	}
	return vertices
}

// MARK: GRIDS
func MgridRandomTileIM(x, y, wBlok, hBlok float32, numWbloks, numHbloks int, tileIM []IM, c sdl.Color, name string) GRID {
	g := Mgrid(x, y, wBlok, hBlok, numWbloks, numHbloks, c, name)
	for range g.r {
		g.tiles = append(g.tiles, tileIM[RINT(0, len(tileIM))])
	}
	return g
}
func MgridWinSize(wBlok, hBlok float32, c sdl.Color, randColor bool, name string) GRID {
	g := GRID{}
	g.c = c
	g.cnt = sdl.FPoint{float32(WINW / 2), float32(WINH / 2)}
	var x, y float32 = 0, 0
	countW := 0
	countedW := false
	countH := 0
	for y <= float32(WINH) {
		g.r = append(g.r, sdl.FRect{x, y, wBlok, hBlok})
		x += wBlok
		if !countedW {
			countW++
		}
		if x >= float32(WINW) {
			countedW = true
			x = 0
			y += hBlok
			countH++
		}
	}
	g.numW = countW
	g.numH = countH
	return g
}
func MgridWH(x, y, wGrid, hGrid float32, numWbloks, numHbloks int, c sdl.Color, name string) GRID {
	g := GRID{}
	g.c = c
	g.w = wGrid
	g.h = hGrid
	g.numW = numWbloks
	g.numH = numHbloks
	g.cnt = sdl.FPoint{x + g.w/2, y + g.h/2}
	w := wGrid / float32(numWbloks)
	h := hGrid / float32(numHbloks)
	a := numWbloks * numHbloks
	cc := 0
	ox := x
	for a > 0 {
		g.r = append(g.r, sdl.FRect{x, y, w, h})
		x += w
		cc++
		a--
		if cc == numWbloks {
			cc = 0
			x = ox
			y += h
		}
	}
	return g
}
func Mgrid(x, y, wBlok, hBlok float32, numWbloks, numHbloks int, c sdl.Color, name string) GRID {
	g := GRID{}
	g.c = c
	g.numW = numWbloks
	g.numH = numHbloks
	g.w = float32(numWbloks) * wBlok
	g.h = float32(numHbloks) * hBlok
	g.cnt = sdl.FPoint{x + g.w/2, y + g.h/2}
	a := numWbloks * numHbloks
	cc := 0
	ox := x
	for a > 0 {
		g.r = append(g.r, sdl.FRect{x, y, wBlok, hBlok})
		x += wBlok
		cc++
		a--
		if cc == numWbloks {
			cc = 0
			x = ox
			y += hBlok
		}
	}
	return g
}
