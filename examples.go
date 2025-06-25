package main

import (
	"fmt"
)

var (
	PLANET, KNIGHTIDLE, KNIGHTATK, KNIGHTWALK, KNIGHTROLL, TILES, SNOWF, SNOWDRAW []IM
)

func mEXAMPLES() {
	//KNIGHT
	KNIGHTIDLE = MimSheetFramesHorizontal("img/knight.png", 32, 32, 4, 0, 0)
	KNIGHTATK = MimSheetFramesHorizontal("img/knight.png", 32, 32, 6, 0, 512)
	KNIGHTATK = MimSheetCollisionRecs(KNIGHTATK, 8, 8, 16, 16)
	KNIGHTWALK = MimSheetFramesHorizontal("img/knight.png", 32, 32, 8, 0, 32)
	KNIGHTROLL = MimSheetFramesHorizontal("img/knight.png", 32, 32, 8, 0, 352)

	//TILES
	TILES = MimSheetMultiRowWidthHeight("img/tiles.png", 0, 0, 16, 16, 256, 192)

	//PLANET
	var filePaths []string
	for i := 1; i < 61; i++ {
		t := "img/planet/" + fmt.Sprint(i) + ".png"
		filePaths = append(filePaths, t)
	}
	PLANET = MimSheetSeperateImageFiles(filePaths)

	//SNOWFLAKES
	SNOWF = MimSheetWidthHorizontal("img/snowflakes.png", 16, 16, 64, 0, 0)
	SNOWDRAW = MdrawSheetRandomPositionsPLUS(SNOWF, 30, 50, 10, 150, 24, 48, 0, 0, float32(WINW), float32(WINH), 0, 45, true, -10, 10)
	SNOWDRAW = UdrawSheetMotionRandom(SNOWDRAW, -1, 1, 4, 8, true, false)

}

func SNOWFLAKES() {
	DdrawSheet(SNOWDRAW)
}
