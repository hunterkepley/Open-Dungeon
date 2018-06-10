/* Hunter Kepley 2018
 * Dungeon crawler w/ full random generation
 * MIT License, check github.com/hunterkepley/Open_Dungeon for more details on that
 */

package main

import (
	"fmt"
	_ "image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	_ "os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	frames       = 0                      // For fps
	second       = time.Tick(time.Second) // For fps
	gameMode     = 1                      // 0 = in main menu, 1 = in game
	camPos       = pixel.ZV
	camZoom      = 0.5
	camZoomSpeed = 1.2
)

const (
	WinWidth  = 800 // Basic starting size.. Not sure if resizing will be added as of yet
	WinHeight = 600
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Open Dungeon",
		Bounds: pixel.R(0, 0, WinWidth, WinHeight),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UTC().UnixNano()) // Seed for random ints

	doneRooms := make(chan bool) // Make sure rooms are finished generating before generating anything else

	go generateRooms(doneRooms, 500, pixel.V(0, 0), pixel.V(500, 500), pixel.V(350, 350))
	if <-doneRooms {
		go generateCorridors()
	}

	player := newPlayer(pixel.V(-12.5, -12.5), pixel.V(25, 25)) // New player, placed in the middle of the camera.

	imd := imdraw.New(nil)

	win.SetSmooth(false) // Smooths out pixels in images/drawn shapes

	last := time.Now()
	for !win.Closed() {
		if gameMode == 1 {
			dt := time.Since(last).Seconds()
			_ = dt
			last = time.Now()

			player.update(playerSpeed, win, dt, &camPos)

			cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
			win.SetMatrix(cam)

			imd.Clear() // Resets shape buffer

			win.Clear(colornames.Black)

			imd.Color = colornames.Lawngreen
			for i := 0; i < len(rooms); i++ {
				rooms[i].render(imd)
			}
			for i := 0; i < len(corridors); i++ {
				corridors[i].render(imd)
			}

			imd.Color = colornames.Ivory
			player.render(imd)

			imd.Draw(win) // Draws shapes
		}

		win.Update()

		frames++ // FPS my guy
		select { // Waits for the block to finish
		case <-second: // A second has passed
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames)) // Appends fps to title for testing
			frames = 0                                                   // Reset it my dude
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
