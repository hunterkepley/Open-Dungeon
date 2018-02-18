package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

var (
	rooms []Room
)

type Room struct {
	pos pixel.Vec
	size pixel.Vec
	closest int
}

func newRoom(pos pixel.Vec, size pixel.Vec) Room {
	return Room{pos, size}
}

/* Amount of rooms, starting pos of spawning, size of bounds to make new rooms, max size of room, and min size of room*/
func generateRooms(doneRooms chan bool, amount int, startPos pixel.Vec, size pixel.Vec, max pixel.Vec, min pixel.Vec) {
	for i := 0; i < amount; i++ {
		alone := true // Check for intersections with existing rooms, delete if false
		rX := randFloat64(startPos.X, startPos.X+size.X)
		rY := randFloat64(startPos.Y, startPos.Y+size.Y)
		rW := randFloat64(min.X, max.X)
		rH := randFloat64(min.Y, max.Y)
		rPos := pixel.V(rX, rY)
		rSize := pixel.V(rW, rH)
		room := Room{rPos, rSize}
		for j := 0; j < len(rooms); j++ {
			if room.intersectsRoom(rooms[j]) {
				alone = false
				i--
				break
			}
		}
		fmt.Println(len(rooms))
		fmt.Println(fmt.Sprintf("#%d", i))
		if alone {
			rooms = append(0, rooms, room)
		}
	}
	doneRooms <- true
}

func (a Room) intersectsRoom(b Room) bool { // Fix this
	if a.pos.X < b.pos.X + b.size.X &&
	  a.pos.X + a.size.X > b.pos.X &&
	  a.pos.Y < b.pos.Y + b.size.Y &&
	  a.size.Y + a.pos.Y > b.pos.Y {
		return true
	}
	return false
}

func (r Room) render(imd *imdraw.IMDraw) {
	imd.Push(pixel.V(r.pos.X, r.pos.Y), pixel.V(r.pos.X + r.size.X, r.pos.Y + r.size.Y))
	imd.Rectangle(0)
}

func (r Room) getCenter() pixel.Vec { // Returns center vector of the room
	return pixel.V(r.pos.X+(r.size.X/2), r.pos.Y+(r.size.Y/2))
}