package main

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

var (
	corridors []Corridor
	corridorWidth = 10.0
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
		cH := Corridor{pixel.V(rooms[i].getCenter().X, rooms[i].getCenter().Y), pixel.V(rooms[closest].getCenter().X - rooms[i].getCenter().X, corridorWidth)}
		corridors = append(corridors, cH)
		if rooms[i].pos.Y > rooms[closest].pos.Y { // If the closest room is below it, then put the vertical room at the top of the path
			cV := Corridor{pixel.V(cH.pos.X+cH.size.X, cH.pos.Y+cH.size.Y), pixel.V(corridorWidth, rooms[closest].pos.Y - cH.pos.Y)}
			corridors = append(corridors, cV)
		} else { // Else, put it at the top of the path [otherwise, it would only touch corners]
			cV := Corridor{pixel.V(cH.pos.X+cH.size.X, cH.pos.Y), pixel.V(corridorWidth, rooms[closest].pos.Y - cH.pos.Y)}
			corridors = append(corridors, cV)
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