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
	connector int // Room connected to this room
	corridored bool // Whether it has a corridor or not
}

func newRoom(pos pixel.Vec, size pixel.Vec, i int) Room { // Room constructor
	return Room{pos, size, i, false}
}

/* Amount of rooms, starting pos of spawning, size of bounds to make new rooms, max size of room, and min size of room*/
func generateRooms(doneRooms chan bool, amount int, startPos pixel.Vec, max pixel.Vec, min pixel.Vec) {
	for i := 0; i < amount; i++ {
		alone := true // Check for intersections with existing rooms, delete if false
		rX := startPos.X
		rY := startPos.Y
		if i != 0 {
			rX = randFloat64(rooms[i-1].pos.X - max.X, (rooms[i-1].pos.X + rooms[i-1].size.X) + max.X/2)
			rY = randFloat64(rooms[i-1].pos.Y - max.Y, (rooms[i-1].pos.Y + rooms[i-1].size.Y) + max.Y/2)
		}
		rW := randFloat64(min.X, max.X)
		rH := randFloat64(min.Y, max.Y)
		rPos := pixel.V(rX, rY)
		rSize := pixel.V(rW, rH)
		room := newRoom(rPos, rSize, 0)
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
			rooms = append(rooms, room)
		}
	}
	doneRooms <- true
}

func (a Room) intersectsRoom(b Room) bool { // If 2 rooms intersect, used for generation
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