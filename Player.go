package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
)

var (
	playerSpeed = 100.0 // Used for camera speed *and* player speed
)

type Player struct {
	pos pixel.Vec
	size pixel.Vec
}

func newPlayer(pos pixel.Vec, size pixel.Vec) Player {
	return Player{pos, size}
}

func (p *Player) update(playerSpeed float64, win *pixelgl.Window, dt float64, camPos *pixel.Vec) {
	if win.Pressed(pixelgl.KeyA) {
		p.pos.X -= playerSpeed * dt
		camPos.X -= playerSpeed * dt
	}
	if win.Pressed(pixelgl.KeyD) {
		p.pos.X += playerSpeed * dt
		camPos.X += playerSpeed * dt
	}
	if win.Pressed(pixelgl.KeyS) {
		p.pos.Y -= playerSpeed * dt
		camPos.Y -= playerSpeed * dt
	}
	if win.Pressed(pixelgl.KeyW) {
		p.pos.Y += playerSpeed * dt
		camPos.Y += playerSpeed * dt
	}
}

func (p Player) render(imd *imdraw.IMDraw) {
	imd.Push(pixel.V(p.pos.X, p.pos.Y), pixel.V(p.pos.X + p.size.X, p.pos.Y + p.size.Y))
	imd.Rectangle(0)
}

func (p Player) getCenter() pixel.Vec { // Returns center vector of the room
	return pixel.V(p.pos.X+(p.size.X/2), p.pos.Y+(p.size.Y/2))
}