package main

import (
	"fmt"
	"math"

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
		closest := 0
		if i == 0 {
			closest = 1
		}
		for j := 0; j < len(rooms); j++ {
			if j != i { // As long as it's not checking itself, see what the closest room is
				distance := math.Sqrt((rooms[j].getCenter().X-rooms[i].getCenter().X)*(rooms[j].getCenter().X-rooms[i].getCenter().X)+(rooms[j].getCenter().Y-rooms[i].getCenter().Y)*(rooms[j].getCenter().Y-rooms[i].getCenter().Y))
				fmt.Println(distance)
				if distance < math.Sqrt((rooms[closest].getCenter().X-rooms[i].getCenter().X)*(rooms[closest].getCenter().X-rooms[i].getCenter().X)+(rooms[closest].getCenter().Y-rooms[i].getCenter().Y)*(rooms[closest].getCenter().Y-rooms[i].getCenter().Y)) {
					closest = j
					rooms[i].closest = closest
				}
				fmt.Println(fmt.Sprintf("%d is closest to %d", i, closest))
			}
		}
		if i == rooms[rooms[i].closest].closest {
			if len(rooms) > 2 {
				if rooms[i].closest != 0 && i != 0 { // All of this makes sure that the rooms don't connect to eachother.
					rooms[i].closest = 0
				} else if rooms[i].closest != 1 && i != 1 {
					rooms[i].closest = 1
				} else if rooms[i].closest != 2 && i != 2 {
					rooms[i].closest = 2
				}
			}
		}
		cH := Corridor{pixel.V(rooms[i].getCenter().X, rooms[i].getCenter().Y), pixel.V(rooms[rooms[i].closest].getCenter().X - rooms[i].getCenter().X, corridorWidth)}
		corridors = append(corridors, cH)
		cV := Corridor{pixel.V(cH.pos.X+cH.size.X/*+corridorWidth*/, cH.pos.Y+cH.size.Y), pixel.V(corridorWidth, rooms[rooms[i].closest].pos.Y - cH.pos.Y)}
		if cH.size.X > 0 {
			cV = Corridor{pixel.V(cH.pos.X+cH.size.X-corridorWidth, cH.pos.Y+cH.size.Y), pixel.V(corridorWidth, rooms[rooms[i].closest].pos.Y - cH.pos.Y)}
		}
		corridors = append(corridors, cV)
	}
}

func (c Corridor) render(imd *imdraw.IMDraw) {
	imd.Push(pixel.V(c.pos.X, c.pos.Y), pixel.V(c.pos.X + c.size.X, c.pos.Y + c.size.Y))
	imd.Rectangle(0)
}

func (c Corridor) getCenter() pixel.Vec { // Returns center vector of the room
	return pixel.V(c.pos.X+(c.size.X/2), c.pos.Y+(c.size.Y/2))
}