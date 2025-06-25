/*
Functions starting with a 'D' >> DRAW & RENDER to screen
Functions starting with an 'M' >> MAKE geometries, objects (OBJ), grids
Functions starting with an 'R' >> RENDER to screen only
Functions starting with an 'C' >> CHECK for collisions/intersections/measurements
Functions starting with 'U' >> UPDATE an object/struct
Functions starting with 'UD' >> UPDATES & DRAWS an object/struct
Functions starting with 'F' >> FIND in a slice/list/directory
*/
package main

import "github.com/veandco/go-sdl2/sdl"

var (
	debugNUM int
)

func PLAY() {

	//OUTSIDE RUN LOOP CODE >>
	var siz, spc, x, y float32 = 200, 20, 10, float32(SCRH/2) - 300
	planets := ManimImSheet(PLANET, 20)
	isogrid := MgridIsometricCubesRandomHeights(sdl.FPoint{x + siz/2, y + siz}, siz/10, siz/10, siz/4, 5, 5, DARKGREEN(), MAGENTA(), ORANGE(), BLACK())
	polyro := float64(0)
	star := Mstar(sdl.FPoint{(spc*3 + siz*3) + siz/2, y + siz/2}, 7, siz/2, siz/4, 30, MAGENTA())
	circle := Mcirc(sdl.FPoint{(spc*4 + siz*4) + siz/2, y + siz/2}, siz/2, ORANGE())
	triangle := MtriCenterMultiColor(sdl.FPoint{(spc*5 + siz*5) + siz/2, y + siz/2}, float64(siz/2), 60, MAGENTA(), ORANGE(), GREEN(), WHITE())
	grid := MgridRandomTileIM(spc*6+siz*6, y, siz/5, siz/5, 5, 5, []IM{TILES[0], TILES[2], TILES[4]}, WHITE(), "grid")
	cube := Mcube(sdl.FPoint{(spc*7 + siz*7) + siz/2, y + siz}, siz/2, siz/2, GREEN(), BLACK())

	for !OFF { //INSIDE RUN LOOP >>

		B4() //DO NOT DELETE
		//DRAW/LOOP CODE BELOW >>
		polygon := Mpolygon(sdl.FPoint{(spc*2 + siz*2) + siz/2, y + siz/2}, float64(siz/2), 7, polyro, GREEN())

		DisoGridCubes(isogrid)
		DtxtXY("IsoGrid", x, y+siz+spc, FONT1DEFAULT, 3, WHITE())
		x += siz + spc
		planets = DanimRecLoop(planets, MrecXY(x, y, siz, siz))
		DtxtXY("ANIM", x, y+siz+spc, FONT1DEFAULT, 3, WHITE())
		x += siz + spc
		DpolygonFill(polygon)
		DtxtXY("Polygon", x, y+siz+spc, FONT1DEFAULT, 3, WHITE())
		x += siz + spc
		DstarFill(star)
		DtxtXY("Star", x, y+siz+spc, FONT1DEFAULT, 3, WHITE())
		x += siz + spc
		DcircFillGradientCirc(circle, ORANGE(), MAGENTA())
		DtxtXY("Circle Gradient", x, y+siz+spc, FONT1DEFAULT, 3, WHITE())
		x += siz + spc
		DtriFillLine(triangle)
		DtxtXY("Triangle", x, y+siz+spc, FONT1DEFAULT, 3, WHITE())
		x += siz + spc
		DgridTiles(grid, true, BLACK())
		DtxtXY("Grid Tiles", x, y+siz+spc, FONT1DEFAULT, 3, WHITE())
		x += siz + spc
		DCubeShadow(cube, BLACK(), false, false, 20, 50)
		DtxtXY("Cube", x, y+siz+spc, FONT1DEFAULT, 3, WHITE())
		polyro++

		SNOWFLAKES()
		SCANLINES(1, 2, 120, BLACK())

		Dfps(float32(WINW)-100, 12, FONT1DEFAULT, 1, true, ORANGE())

		x, y = 10, float32(SCRH/2)-300
		//END DRAW/LOOP CODE
		AFTER() //DO NOT DELETE
	}
}

func main() {

	//CHANGE SETTINGS BELOW BEFORE INITIAL()

	//DEBUG
	//DISPLAYERRORS = true //SET TO FALSE TO NOT DISPLAY ONSCREEN ERRORS

	//WINDOW >>
	//FULLSCREEN = true //UNCOMMENT FOR FULLSCREEN
	//WINWDEF, WINHDEF = 1920, 1080 //SET DEFAULT WINDOW SIZE IF NOT USING FULLSCREEN

	//FPS
	//SETFPS = true //SET TO TRUE TO SET A FRAME RATE
	//TARGETFPS = 30 //SET THE DESIRED FRAME RATE

	//BUTTONS >>
	//BUTTONCOLON = GREEN() //BUTTON COLOR ON
	//BUTTONCOLOFF = RED()  //BUTTON COLOR OFF
	//BUTTONSIZE = 64

	//INITIALIZE
	INITIAL("DiLgo >> Go & SDL2", BLACK()) //CREATE WINDOW > (WINDOW TITLE, BACKGROUND COLOR)
	//RUN
	PLAY()
}
