package main

import "github.com/veandco/go-sdl2/sdl"

var (
	BGCOL        sdl.Color
	COLLIST      []COLOBJ
	COLFADE      = uint8(100)
	colFadeONOFF bool
	BLANKCOL     = sdl.Color{}
)

type COLOBJ struct {
	nm string
	c  sdl.Color
}

func COLOR2RGBA(c sdl.Color) (uint8, uint8, uint8, uint8) {
	return c.R, c.G, c.B, c.A
}
func COL(c sdl.Color) {
	RND.SetDrawColor(c.R, c.G, c.B, c.A)
}
func COLORALPHA(c sdl.Color, a uint8) sdl.Color {
	c.A = a
	return c
}

func mCOLLIST() {
	c := COLOBJ{}
	c.nm = "Maroon"
	c.c = MAROON()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Red"
	c.c = DARKRED()
	COLLIST = append(COLLIST, c)
	c.nm = "Brown"
	c.c = BROWN()
	COLLIST = append(COLLIST, c)
	c.nm = "Fire Brick"
	c.c = FIREBRICK()
	COLLIST = append(COLLIST, c)
	c.nm = "Crimson"
	c.c = CRIMSON()
	COLLIST = append(COLLIST, c)
	c.nm = "Red"
	c.c = RED()
	COLLIST = append(COLLIST, c)
	c.nm = "Tomato"
	c.c = TOMATO()
	COLLIST = append(COLLIST, c)
	c.nm = "Coral"
	c.c = CORAL()
	COLLIST = append(COLLIST, c)
	c.nm = "Indian Red"
	c.c = INDIANRED()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Coral"
	c.c = LIGHTCORAL()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Salmon"
	c.c = DARKSALMON()
	COLLIST = append(COLLIST, c)
	c.nm = "Salmon"
	c.c = SALMON()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Salmon"
	c.c = LIGHTSALMON()
	COLLIST = append(COLLIST, c)
	c.nm = "Orange Red"
	c.c = ORANGERED()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Orange"
	c.c = DARKORANGE()
	COLLIST = append(COLLIST, c)
	c.nm = "Orange"
	c.c = ORANGE()
	COLLIST = append(COLLIST, c)
	c.nm = "Gold"
	c.c = GOLD()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Goldenrod"
	c.c = DARKGOLDENROD()
	COLLIST = append(COLLIST, c)
	c.nm = "Goldenrod"
	c.c = GOLDENROD()
	COLLIST = append(COLLIST, c)
	c.nm = "Pale Goldenrod"
	c.c = PALEGOLDENROD()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Khaki"
	c.c = DARKKHAKI()
	COLLIST = append(COLLIST, c)
	c.nm = "Khaki"
	c.c = KHAKI()
	COLLIST = append(COLLIST, c)
	c.nm = "Olive"
	c.c = OLIVE()
	COLLIST = append(COLLIST, c)
	c.nm = "Yellow"
	c.c = YELLOW()
	COLLIST = append(COLLIST, c)
	c.nm = "Yellow Green"
	c.c = YELLOWGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Olive Green"
	c.c = DARKOLIVEGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Olive Drab"
	c.c = OLIVEDRAB()
	COLLIST = append(COLLIST, c)
	c.nm = "Lawn Green"
	c.c = LAWNGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Chartreuse"
	c.c = CHARTREUSE()
	COLLIST = append(COLLIST, c)
	c.nm = "Green Yellow"
	c.c = GREENYELLOW()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Green"
	c.c = DARKGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Green"
	c.c = GREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Forest Green"
	c.c = FORESTGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Lime"
	c.c = LIME()
	COLLIST = append(COLLIST, c)
	c.nm = "Lime Green"
	c.c = LIMEGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Green"
	c.c = LIGHTGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Pale Green"
	c.c = PALEGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Sea Green"
	c.c = DARKSEAGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Spring Green"
	c.c = MEDIUMSPRINGGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Spring Green"
	c.c = SPRINGGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Sea Green"
	c.c = SEAGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Aquamarine"
	c.c = MEDIUMAQUAMARINE()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Sea Green"
	c.c = MEDIUMSEAGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Sea Green"
	c.c = LIGHTSEAGREEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Slate Gray"
	c.c = DARKSLATEGRAY()
	COLLIST = append(COLLIST, c)
	c.nm = "Teal"
	c.c = TEAL()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Cyan"
	c.c = DARKCYAN()
	COLLIST = append(COLLIST, c)
	c.nm = "Aqua"
	c.c = AQUA()
	COLLIST = append(COLLIST, c)
	c.nm = "Cyan"
	c.c = CYAN()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Cyan"
	c.c = LIGHTCYAN()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Turquoise"
	c.c = DARKTURQUOISE()
	COLLIST = append(COLLIST, c)
	c.nm = "Turquoise"
	c.c = TURQUOISE()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Turquoise"
	c.c = MEDIUMTURQUOISE()
	COLLIST = append(COLLIST, c)
	c.nm = "Pale Turquoise"
	c.c = PALETURQUOISE()
	COLLIST = append(COLLIST, c)
	c.nm = "Aquamarine"
	c.c = AQUAMARINE()
	COLLIST = append(COLLIST, c)
	c.nm = "Powder Blue"
	c.c = POWDERBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Cadet Blue"
	c.c = CADETBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Steel Blue"
	c.c = STEELBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Cornflower Blue"
	c.c = CORNFLOWERBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Deep Sky Blue"
	c.c = DEEPSKYBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Dodger Blue"
	c.c = DODGERBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Blue"
	c.c = LIGHTBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Sky Blue"
	c.c = SKYBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Sky Blue"
	c.c = LIGHTSKYBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Midnight Blue"
	c.c = MIDNIGHTBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Navy"
	c.c = NAVY()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Blue"
	c.c = DARKBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Blue"
	c.c = MEDIUMBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Blue"
	c.c = BLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Royal Blue"
	c.c = ROYALBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Blue Violet"
	c.c = BLUEVIOLET()
	COLLIST = append(COLLIST, c)
	c.nm = "Indigo"
	c.c = INDIGO()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Slate Blue"
	c.c = DARKSLATEBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Slate Blue"
	c.c = SLATEBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Slate Blue"
	c.c = MEDIUMSLATEBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Purple"
	c.c = MEDIUMPURPLE()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Magenta"
	c.c = DARKMAGENTA()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Violet"
	c.c = DARKVIOLET()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Orchid"
	c.c = DARKORCHID()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Orchid"
	c.c = MEDIUMORCHID()
	COLLIST = append(COLLIST, c)
	c.nm = "Purple"
	c.c = PURPLE()
	COLLIST = append(COLLIST, c)
	c.nm = "Thistle"
	c.c = THISTLE()
	COLLIST = append(COLLIST, c)
	c.nm = "Plum"
	c.c = PLUM()
	COLLIST = append(COLLIST, c)
	c.nm = "Violet"
	c.c = VIOLET()
	COLLIST = append(COLLIST, c)
	c.nm = "Magenta"
	c.c = MAGENTA()
	COLLIST = append(COLLIST, c)
	c.nm = "Orchid"
	c.c = ORCHID()
	COLLIST = append(COLLIST, c)
	c.nm = "Medium Violet Red"
	c.c = MEDIUMVIOLETRED()
	COLLIST = append(COLLIST, c)
	c.nm = "Pale Violet Red"
	c.c = PALEVIOLETRED()
	COLLIST = append(COLLIST, c)
	c.nm = "Deep Pink"
	c.c = DEEPPINK()
	COLLIST = append(COLLIST, c)
	c.nm = "Hot Pink"
	c.c = HOTPINK()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Pink"
	c.c = LIGHTPINK()
	COLLIST = append(COLLIST, c)
	c.nm = "Pink"
	c.c = PINK()
	COLLIST = append(COLLIST, c)
	c.nm = "Antique White"
	c.c = ANTIQUEWHITE()
	COLLIST = append(COLLIST, c)
	c.nm = "Beige"
	c.c = BEIGE()
	COLLIST = append(COLLIST, c)
	c.nm = "Bisque"
	c.c = BISQUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Blanched Almond"
	c.c = BLANCHEDALMOND()
	COLLIST = append(COLLIST, c)
	c.nm = "Wheat"
	c.c = WHEAT()
	COLLIST = append(COLLIST, c)
	c.nm = "Corn Silk"
	c.c = CORNSILK()
	COLLIST = append(COLLIST, c)
	c.nm = "Lemon Chiffon"
	c.c = LEMONCHIFFON()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Goldenrod Yellow"
	c.c = LIGHTGOLDENRODYELLOW()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Yellow"
	c.c = LIGHTYELLOW()
	COLLIST = append(COLLIST, c)
	c.nm = "Saddle Brown"
	c.c = SADDLEBROWN()
	COLLIST = append(COLLIST, c)
	c.nm = "Sienna"
	c.c = SIENNA()
	COLLIST = append(COLLIST, c)
	c.nm = "Chocolate"
	c.c = CHOCOLATE()
	COLLIST = append(COLLIST, c)
	c.nm = "Peru"
	c.c = PERU()
	COLLIST = append(COLLIST, c)
	c.nm = "Sandy Brown"
	c.c = SANDYBROWN()
	COLLIST = append(COLLIST, c)
	c.nm = "Burlywood"
	c.c = BURLYWOOD()
	COLLIST = append(COLLIST, c)
	c.nm = "Tan"
	c.c = TAN()
	COLLIST = append(COLLIST, c)
	c.nm = "Rosy Brown"
	c.c = ROSYBROWN()
	COLLIST = append(COLLIST, c)
	c.nm = "Moccasin"
	c.c = MOCCASIN()
	COLLIST = append(COLLIST, c)
	c.nm = "Navajo White"
	c.c = NAVAJOWHITE()
	COLLIST = append(COLLIST, c)
	c.nm = "Peach Puff"
	c.c = PEACHPUFF()
	COLLIST = append(COLLIST, c)
	c.nm = "Misty Rose"
	c.c = MISTYROSE()
	COLLIST = append(COLLIST, c)
	c.nm = "Lavender Blush"
	c.c = LAVENDERBLUSH()
	COLLIST = append(COLLIST, c)
	c.nm = "Linen"
	c.c = LINEN()
	COLLIST = append(COLLIST, c)
	c.nm = "Old Lace"
	c.c = OLDLACE()
	COLLIST = append(COLLIST, c)
	c.nm = "Papaya Whip"
	c.c = PAPAYAWHIP()
	COLLIST = append(COLLIST, c)
	c.nm = "Seashell"
	c.c = SEASHELL()
	COLLIST = append(COLLIST, c)
	c.nm = "Mint Cream"
	c.c = MINTCREAM()
	COLLIST = append(COLLIST, c)
	c.nm = "Slate Gray"
	c.c = SLATEGRAY()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Slate Gray"
	c.c = LIGHTSLATEGRAY()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Steel Blue"
	c.c = LIGHTSTEELBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Lavender"
	c.c = LAVENDER()
	COLLIST = append(COLLIST, c)
	c.nm = "Floral White"
	c.c = FLORALWHITE()
	COLLIST = append(COLLIST, c)
	c.nm = "Alice Blue"
	c.c = ALICEBLUE()
	COLLIST = append(COLLIST, c)
	c.nm = "Ghost White"
	c.c = GHOSTWHITE()
	COLLIST = append(COLLIST, c)
	c.nm = "Honeydew"
	c.c = HONEYDEW()
	COLLIST = append(COLLIST, c)
	c.nm = "Ivory"
	c.c = IVORY()
	COLLIST = append(COLLIST, c)
	c.nm = "Azure"
	c.c = AZURE()
	COLLIST = append(COLLIST, c)
	c.nm = "Snow"
	c.c = SNOW()
	COLLIST = append(COLLIST, c)
	c.nm = "Black"
	c.c = BLACK()
	COLLIST = append(COLLIST, c)
	c.nm = "Dim Grey"
	c.c = DIMGREY()
	COLLIST = append(COLLIST, c)
	c.nm = "Grey"
	c.c = GREY()
	COLLIST = append(COLLIST, c)
	c.nm = "Dark Grey"
	c.c = DARKGREY()
	COLLIST = append(COLLIST, c)
	c.nm = "Silver"
	c.c = SILVER()
	COLLIST = append(COLLIST, c)
	c.nm = "Light Grey"
	c.c = LIGHTGREY()
	COLLIST = append(COLLIST, c)
	c.nm = "Gainsboro"
	c.c = GAINSBORO()
	COLLIST = append(COLLIST, c)
	c.nm = "White Smoke"
	c.c = WHITESMOKE()
	COLLIST = append(COLLIST, c)
	c.nm = "White"
	c.c = WHITE()
	COLLIST = append(COLLIST, c)

}

// RANDOM COLORS
func COLORǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255}
}
func DARKGREENǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255}
}
func GREENǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255}
}
func REDǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255}
}
func PINKǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255}
}
func BLUEǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255}
}
func DARKBLUEǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255}
}
func CYANǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255}
}
func YELLOWǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255}
}
func ORANGEǁRANDOM() sdl.Color {
	return sdl.Color{255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255}
}
func BROWNǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255}
}
func GREYǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255}
}
func DARKGREYǁRANDOM() sdl.Color {
	return sdl.Color{uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255}
}

func COLORǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255
}
func DARKGREENǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255
}
func GREENǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255
}
func REDǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255
}
func PINKǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255
}
func BLUEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255
}
func DARKBLUEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255
}
func CYANǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255
}
func YELLOWǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255
}
func ORANGEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return 255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255
}
func BROWNǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255
}
func GREYǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255
}
func DARKGREYǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255
}

func COLORǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255}
}
func DARKGREENǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255}
}
func GREENǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255}
}
func REDǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255}
}
func PINKǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255}
}
func BLUEǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255}
}
func DARKBLUEǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255}
}
func CYANǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255}
}
func YELLOWǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255}
}
func ORANGEǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255}
}
func BROWNǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255}
}
func GREYǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255}
}
func DARKGREYǁRANDOMǁ3() (rgba []uint8) {
	return []uint8{uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255}
}

// SOLID COLORS
func MAROON() sdl.Color               { return sdl.Color{uint8(128), uint8(0), uint8(0), 255} }
func DARKRED() sdl.Color              { return sdl.Color{uint8(139), uint8(0), uint8(0), 255} }
func BROWN() sdl.Color                { return sdl.Color{uint8(165), uint8(42), uint8(42), 255} }
func FIREBRICK() sdl.Color            { return sdl.Color{uint8(178), uint8(34), uint8(34), 255} }
func CRIMSON() sdl.Color              { return sdl.Color{uint8(220), uint8(20), uint8(60), 255} }
func RED() sdl.Color                  { return sdl.Color{uint8(255), uint8(0), uint8(0), 255} }
func TOMATO() sdl.Color               { return sdl.Color{uint8(255), uint8(99), uint8(71), 255} }
func CORAL() sdl.Color                { return sdl.Color{uint8(255), uint8(127), uint8(80), 255} }
func INDIANRED() sdl.Color            { return sdl.Color{uint8(205), uint8(92), uint8(92), 255} }
func LIGHTCORAL() sdl.Color           { return sdl.Color{uint8(240), uint8(128), uint8(128), 255} }
func DARKSALMON() sdl.Color           { return sdl.Color{uint8(233), uint8(150), uint8(122), 255} }
func SALMON() sdl.Color               { return sdl.Color{uint8(250), uint8(128), uint8(114), 255} }
func LIGHTSALMON() sdl.Color          { return sdl.Color{uint8(255), uint8(160), uint8(122), 255} }
func ORANGERED() sdl.Color            { return sdl.Color{uint8(255), uint8(69), uint8(0), 255} }
func DARKORANGE() sdl.Color           { return sdl.Color{uint8(255), uint8(140), uint8(0), 255} }
func ORANGE() sdl.Color               { return sdl.Color{uint8(255), uint8(165), uint8(0), 255} }
func GOLD() sdl.Color                 { return sdl.Color{uint8(255), uint8(215), uint8(0), 255} }
func DARKGOLDENROD() sdl.Color        { return sdl.Color{uint8(184), uint8(134), uint8(11), 255} }
func GOLDENROD() sdl.Color            { return sdl.Color{uint8(218), uint8(165), uint8(32), 255} }
func PALEGOLDENROD() sdl.Color        { return sdl.Color{uint8(238), uint8(232), uint8(170), 255} }
func DARKKHAKI() sdl.Color            { return sdl.Color{uint8(189), uint8(183), uint8(107), 255} }
func KHAKI() sdl.Color                { return sdl.Color{uint8(240), uint8(230), uint8(140), 255} }
func OLIVE() sdl.Color                { return sdl.Color{uint8(128), uint8(128), uint8(0), 255} }
func YELLOW() sdl.Color               { return sdl.Color{uint8(255), uint8(255), uint8(0), 255} }
func YELLOWGREEN() sdl.Color          { return sdl.Color{uint8(154), uint8(205), uint8(50), 255} }
func DARKOLIVEGREEN() sdl.Color       { return sdl.Color{uint8(85), uint8(107), uint8(47), 255} }
func OLIVEDRAB() sdl.Color            { return sdl.Color{uint8(107), uint8(142), uint8(35), 255} }
func LAWNGREEN() sdl.Color            { return sdl.Color{uint8(124), uint8(252), uint8(0), 255} }
func CHARTREUSE() sdl.Color           { return sdl.Color{uint8(127), uint8(255), uint8(0), 255} }
func GREENYELLOW() sdl.Color          { return sdl.Color{uint8(173), uint8(255), uint8(47), 255} }
func DARKGREEN() sdl.Color            { return sdl.Color{uint8(0), uint8(100), uint8(0), 255} }
func GREEN() sdl.Color                { return sdl.Color{uint8(0), uint8(128), uint8(0), 255} }
func FORESTGREEN() sdl.Color          { return sdl.Color{uint8(34), uint8(139), uint8(34), 255} }
func LIME() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(0), 255} }
func LIMEGREEN() sdl.Color            { return sdl.Color{uint8(50), uint8(205), uint8(50), 255} }
func LIGHTGREEN() sdl.Color           { return sdl.Color{uint8(144), uint8(238), uint8(144), 255} }
func PALEGREEN() sdl.Color            { return sdl.Color{uint8(152), uint8(251), uint8(152), 255} }
func DARKSEAGREEN() sdl.Color         { return sdl.Color{uint8(143), uint8(188), uint8(143), 255} }
func MEDIUMSPRINGGREEN() sdl.Color    { return sdl.Color{uint8(0), uint8(250), uint8(154), 255} }
func SPRINGGREEN() sdl.Color          { return sdl.Color{uint8(0), uint8(255), uint8(127), 255} }
func SEAGREEN() sdl.Color             { return sdl.Color{uint8(46), uint8(139), uint8(87), 255} }
func MEDIUMAQUAMARINE() sdl.Color     { return sdl.Color{uint8(102), uint8(205), uint8(170), 255} }
func MEDIUMSEAGREEN() sdl.Color       { return sdl.Color{uint8(60), uint8(179), uint8(113), 255} }
func LIGHTSEAGREEN() sdl.Color        { return sdl.Color{uint8(32), uint8(178), uint8(170), 255} }
func DARKSLATEGRAY() sdl.Color        { return sdl.Color{uint8(47), uint8(79), uint8(79), 255} }
func TEAL() sdl.Color                 { return sdl.Color{uint8(0), uint8(128), uint8(128), 255} }
func DARKCYAN() sdl.Color             { return sdl.Color{uint8(0), uint8(139), uint8(139), 255} }
func AQUA() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(255), 255} }
func CYAN() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(255), 255} }
func LIGHTCYAN() sdl.Color            { return sdl.Color{uint8(224), uint8(255), uint8(255), 255} }
func DARKTURQUOISE() sdl.Color        { return sdl.Color{uint8(0), uint8(206), uint8(209), 255} }
func TURQUOISE() sdl.Color            { return sdl.Color{uint8(64), uint8(224), uint8(208), 255} }
func MEDIUMTURQUOISE() sdl.Color      { return sdl.Color{uint8(72), uint8(209), uint8(204), 255} }
func PALETURQUOISE() sdl.Color        { return sdl.Color{uint8(175), uint8(238), uint8(238), 255} }
func AQUAMARINE() sdl.Color           { return sdl.Color{uint8(127), uint8(255), uint8(212), 255} }
func POWDERBLUE() sdl.Color           { return sdl.Color{uint8(176), uint8(224), uint8(230), 255} }
func CADETBLUE() sdl.Color            { return sdl.Color{uint8(95), uint8(158), uint8(160), 255} }
func STEELBLUE() sdl.Color            { return sdl.Color{uint8(70), uint8(130), uint8(180), 255} }
func CORNFLOWERBLUE() sdl.Color       { return sdl.Color{uint8(100), uint8(149), uint8(237), 255} }
func DEEPSKYBLUE() sdl.Color          { return sdl.Color{uint8(0), uint8(191), uint8(255), 255} }
func DODGERBLUE() sdl.Color           { return sdl.Color{uint8(30), uint8(144), uint8(255), 255} }
func LIGHTBLUE() sdl.Color            { return sdl.Color{uint8(173), uint8(216), uint8(230), 255} }
func SKYBLUE() sdl.Color              { return sdl.Color{uint8(135), uint8(206), uint8(235), 255} }
func LIGHTSKYBLUE() sdl.Color         { return sdl.Color{uint8(135), uint8(206), uint8(250), 255} }
func MIDNIGHTBLUE() sdl.Color         { return sdl.Color{uint8(25), uint8(25), uint8(112), 255} }
func NAVY() sdl.Color                 { return sdl.Color{uint8(0), uint8(0), uint8(128), 255} }
func DARKBLUE() sdl.Color             { return sdl.Color{uint8(0), uint8(0), uint8(139), 255} }
func MEDIUMBLUE() sdl.Color           { return sdl.Color{uint8(0), uint8(0), uint8(205), 255} }
func BLUE() sdl.Color                 { return sdl.Color{uint8(0), uint8(0), uint8(255), 255} }
func ROYALBLUE() sdl.Color            { return sdl.Color{uint8(65), uint8(105), uint8(225), 255} }
func BLUEVIOLET() sdl.Color           { return sdl.Color{uint8(138), uint8(43), uint8(226), 255} }
func INDIGO() sdl.Color               { return sdl.Color{uint8(75), uint8(0), uint8(130), 255} }
func DARKSLATEBLUE() sdl.Color        { return sdl.Color{uint8(72), uint8(61), uint8(139), 255} }
func SLATEBLUE() sdl.Color            { return sdl.Color{uint8(106), uint8(90), uint8(205), 255} }
func MEDIUMSLATEBLUE() sdl.Color      { return sdl.Color{uint8(123), uint8(104), uint8(238), 255} }
func MEDIUMPURPLE() sdl.Color         { return sdl.Color{uint8(147), uint8(112), uint8(219), 255} }
func DARKMAGENTA() sdl.Color          { return sdl.Color{uint8(139), uint8(0), uint8(139), 255} }
func DARKVIOLET() sdl.Color           { return sdl.Color{uint8(148), uint8(0), uint8(211), 255} }
func DARKORCHID() sdl.Color           { return sdl.Color{uint8(153), uint8(50), uint8(204), 255} }
func MEDIUMORCHID() sdl.Color         { return sdl.Color{uint8(186), uint8(85), uint8(211), 255} }
func PURPLE() sdl.Color               { return sdl.Color{uint8(128), uint8(0), uint8(128), 255} }
func THISTLE() sdl.Color              { return sdl.Color{uint8(216), uint8(191), uint8(216), 255} }
func PLUM() sdl.Color                 { return sdl.Color{uint8(221), uint8(160), uint8(221), 255} }
func VIOLET() sdl.Color               { return sdl.Color{uint8(238), uint8(130), uint8(238), 255} }
func MAGENTA() sdl.Color              { return sdl.Color{uint8(255), uint8(0), uint8(255), 255} }
func ORCHID() sdl.Color               { return sdl.Color{uint8(218), uint8(112), uint8(214), 255} }
func MEDIUMVIOLETRED() sdl.Color      { return sdl.Color{uint8(199), uint8(21), uint8(133), 255} }
func PALEVIOLETRED() sdl.Color        { return sdl.Color{uint8(219), uint8(112), uint8(147), 255} }
func DEEPPINK() sdl.Color             { return sdl.Color{uint8(255), uint8(20), uint8(147), 255} }
func HOTPINK() sdl.Color              { return sdl.Color{uint8(255), uint8(105), uint8(180), 255} }
func LIGHTPINK() sdl.Color            { return sdl.Color{uint8(255), uint8(182), uint8(193), 255} }
func PINK() sdl.Color                 { return sdl.Color{uint8(255), uint8(192), uint8(203), 255} }
func ANTIQUEWHITE() sdl.Color         { return sdl.Color{uint8(250), uint8(235), uint8(215), 255} }
func BEIGE() sdl.Color                { return sdl.Color{uint8(245), uint8(245), uint8(220), 255} }
func BISQUE() sdl.Color               { return sdl.Color{uint8(255), uint8(228), uint8(196), 255} }
func BLANCHEDALMOND() sdl.Color       { return sdl.Color{uint8(255), uint8(235), uint8(205), 255} }
func WHEAT() sdl.Color                { return sdl.Color{uint8(245), uint8(222), uint8(179), 255} }
func CORNSILK() sdl.Color             { return sdl.Color{uint8(255), uint8(248), uint8(220), 255} }
func LEMONCHIFFON() sdl.Color         { return sdl.Color{uint8(255), uint8(250), uint8(205), 255} }
func LIGHTGOLDENRODYELLOW() sdl.Color { return sdl.Color{uint8(250), uint8(250), uint8(210), 255} }
func LIGHTYELLOW() sdl.Color          { return sdl.Color{uint8(255), uint8(255), uint8(224), 255} }
func SADDLEBROWN() sdl.Color          { return sdl.Color{uint8(139), uint8(69), uint8(19), 255} }
func SIENNA() sdl.Color               { return sdl.Color{uint8(160), uint8(82), uint8(45), 255} }
func CHOCOLATE() sdl.Color            { return sdl.Color{uint8(210), uint8(105), uint8(30), 255} }
func PERU() sdl.Color                 { return sdl.Color{uint8(205), uint8(133), uint8(63), 255} }
func SANDYBROWN() sdl.Color           { return sdl.Color{uint8(244), uint8(164), uint8(96), 255} }
func BURLYWOOD() sdl.Color            { return sdl.Color{uint8(222), uint8(184), uint8(135), 255} }
func TAN() sdl.Color                  { return sdl.Color{uint8(210), uint8(180), uint8(140), 255} }
func ROSYBROWN() sdl.Color            { return sdl.Color{uint8(188), uint8(143), uint8(143), 255} }
func MOCCASIN() sdl.Color             { return sdl.Color{uint8(255), uint8(228), uint8(181), 255} }
func NAVAJOWHITE() sdl.Color          { return sdl.Color{uint8(255), uint8(222), uint8(173), 255} }
func PEACHPUFF() sdl.Color            { return sdl.Color{uint8(255), uint8(218), uint8(185), 255} }
func MISTYROSE() sdl.Color            { return sdl.Color{uint8(255), uint8(228), uint8(225), 255} }
func LAVENDERBLUSH() sdl.Color        { return sdl.Color{uint8(255), uint8(240), uint8(245), 255} }
func LINEN() sdl.Color                { return sdl.Color{uint8(250), uint8(240), uint8(230), 255} }
func OLDLACE() sdl.Color              { return sdl.Color{uint8(253), uint8(245), uint8(230), 255} }
func PAPAYAWHIP() sdl.Color           { return sdl.Color{uint8(255), uint8(239), uint8(213), 255} }
func SEASHELL() sdl.Color             { return sdl.Color{uint8(255), uint8(245), uint8(238), 255} }
func MINTCREAM() sdl.Color            { return sdl.Color{uint8(245), uint8(255), uint8(250), 255} }
func SLATEGRAY() sdl.Color            { return sdl.Color{uint8(112), uint8(128), uint8(144), 255} }
func LIGHTSLATEGRAY() sdl.Color       { return sdl.Color{uint8(119), uint8(136), uint8(153), 255} }
func LIGHTSTEELBLUE() sdl.Color       { return sdl.Color{uint8(176), uint8(196), uint8(222), 255} }
func LAVENDER() sdl.Color             { return sdl.Color{uint8(230), uint8(230), uint8(250), 255} }
func FLORALWHITE() sdl.Color          { return sdl.Color{uint8(255), uint8(250), uint8(240), 255} }
func ALICEBLUE() sdl.Color            { return sdl.Color{uint8(240), uint8(248), uint8(255), 255} }
func GHOSTWHITE() sdl.Color           { return sdl.Color{uint8(248), uint8(248), uint8(255), 255} }
func HONEYDEW() sdl.Color             { return sdl.Color{uint8(240), uint8(255), uint8(240), 255} }
func IVORY() sdl.Color                { return sdl.Color{uint8(255), uint8(255), uint8(240), 255} }
func AZURE() sdl.Color                { return sdl.Color{uint8(240), uint8(255), uint8(255), 255} }
func SNOW() sdl.Color                 { return sdl.Color{uint8(255), uint8(250), uint8(250), 255} }
func BLACK() sdl.Color                { return sdl.Color{uint8(0), uint8(0), uint8(0), 255} }
func DIMGREY() sdl.Color              { return sdl.Color{uint8(105), uint8(105), uint8(105), 255} }
func GREY() sdl.Color                 { return sdl.Color{uint8(128), uint8(128), uint8(128), 255} }
func DARKGREY() sdl.Color             { return sdl.Color{uint8(169), uint8(169), uint8(169), 255} }
func SILVER() sdl.Color               { return sdl.Color{uint8(192), uint8(192), uint8(192), 255} }
func LIGHTGREY() sdl.Color            { return sdl.Color{uint8(211), uint8(211), uint8(211), 255} }
func GAINSBORO() sdl.Color            { return sdl.Color{uint8(220), uint8(220), uint8(220), 255} }
func WHITESMOKE() sdl.Color           { return sdl.Color{uint8(245), uint8(245), uint8(245), 255} }
func WHITE() sdl.Color                { return sdl.Color{uint8(255), uint8(255), uint8(255), 255} }
