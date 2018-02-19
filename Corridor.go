package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

var (
	corridors []Corridor
	corridorWidth = 5.0
)

type Corridor struct {
	pos pixel.Vec
	size pixel.Vec
}

func newCorridor(pos pixel.Vec, size pixel.Vec) Corridor {
	return Corridor{pos, size}
}

func generateCorridors() {
	for i := 0; i < len(rooms); i++ {
		if !rooms[i].corridored {
			if i > 0 { // Connect to the last made room to make sure they all connect.
				rooms[i].connector = i-1
			}
			cH := Corridor{pixel.V(rooms[i].getCenter().X, rooms[i].getCenter().Y), pixel.V(rooms[rooms[i].connector].getCenter().X - rooms[i].getCenter().X, corridorWidth)}
			corridors = append(corridors, cH)
			cV := Corridor{pixel.V(cH.pos.X+cH.size.X, cH.pos.Y+cH.size.Y), pixel.V(corridorWidth, rooms[rooms[i].connector].pos.Y - cH.pos.Y)}
			if cH.size.X > 0 {
				cV = Corridor{pixel.V(cH.pos.X+cH.size.X-corridorWidth, cH.pos.Y+cH.size.Y), pixel.V(corridorWidth, rooms[rooms[i].connector].pos.Y - cH.pos.Y)}
			}
			corridors = append(corridors, cV)
			rooms[i].corridored = true
		}
	}
}

func (c Corridor) render(imd *imdraw.IMDraw) {
	imd.Push(pixel.V(c.pos.X, c.pos.Y), pixel.V(c.pos.X + c.size.X, c.pos.Y + c.size.Y))
	imd.Rectangle(0)
}

func (c Corridor) getCenter() pixel.Vec { // Returns center vector of the room
	return pixel.V(c.pos.X+(c.size.X/2), c.pos.Y+(c.size.Y/2))
}