package main

import "github.com/veandco/go-sdl2/sdl"

var ()

type OBJ struct {
	cnt                  sdl.FPoint
	isorec               ISOREC
	cube                 CUBE
	name                 string
	isIsometricOBJ       bool
	r, rCollision, rDraw sdl.FRect
	im                   IM
	anm                  ANIM
	anmmulti             ANIMMULTI

	speedX, speedY, speedMax, w, h float32
}
