package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
	"time"
)

var blueColor = color.RGBA{0x00, 0x00, 0xff, 0xff}
var greenColor = color.RGBA{0x00, 0xff, 0x00, 0xff}
var redColor = color.RGBA{0xff, 0x00, 0x00, 0xff}

var palette = []color.Color{
	blueColor,
	greenColor,
	redColor,
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var width = 128
	var height = 128

	var images []*image.Paletted
	var delays []int

	var livePixels map[string]*Life
	livePixels = make(map[string]*Life)

	var startCenterCubeEdge = 40
	// var endCenterCubeEdge = width/2 + 4

	steps := 256

	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, width, height), palette)
		images = append(images, img)
		delays = append(delays, 0)

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				img.Set(x, y, greenColor)
			}
		}

		if step == 0 {
			// for x := startCenterCubeEdge; x <= endCenterCubeEdge; x++ {
			// 	for y := startCenterCubeEdge; y <= endCenterCubeEdge; y++ {
			// 		randomNum := rand.Intn(10)

			// 		if randomNum < 4 {
			// 			var life = &Life{x: x, y: y}
			// 			livePixels[life.GetKey()] = life

			// 			img.Set(x, y, blueColor)
			// 		}
			// 	}
			// }

			addLife(img, livePixels, startCenterCubeEdge, startCenterCubeEdge)
			addLife(img, livePixels, startCenterCubeEdge+1, startCenterCubeEdge)
			addLife(img, livePixels, startCenterCubeEdge, startCenterCubeEdge+1)
			addLife(img, livePixels, startCenterCubeEdge+1, startCenterCubeEdge+1)

			addLife(img, livePixels, startCenterCubeEdge+10, startCenterCubeEdge)
			addLife(img, livePixels, startCenterCubeEdge+10, startCenterCubeEdge+1)
			addLife(img, livePixels, startCenterCubeEdge+10, startCenterCubeEdge+2)

			addLife(img, livePixels, startCenterCubeEdge+11, startCenterCubeEdge-1)
			addLife(img, livePixels, startCenterCubeEdge+11, startCenterCubeEdge+3)

			addLife(img, livePixels, startCenterCubeEdge+12, startCenterCubeEdge-2)
			addLife(img, livePixels, startCenterCubeEdge+12, startCenterCubeEdge+4)

			addLife(img, livePixels, startCenterCubeEdge+13, startCenterCubeEdge-2)
			addLife(img, livePixels, startCenterCubeEdge+13, startCenterCubeEdge+4)

			addLife(img, livePixels, startCenterCubeEdge+14, startCenterCubeEdge+1)

			addLife(img, livePixels, startCenterCubeEdge+15, startCenterCubeEdge-1)
			addLife(img, livePixels, startCenterCubeEdge+15, startCenterCubeEdge+3)

			addLife(img, livePixels, startCenterCubeEdge+16, startCenterCubeEdge)
			addLife(img, livePixels, startCenterCubeEdge+16, startCenterCubeEdge+1)
			addLife(img, livePixels, startCenterCubeEdge+16, startCenterCubeEdge+2)
			// Correct

			addLife(img, livePixels, startCenterCubeEdge+17, startCenterCubeEdge+1)

			addLife(img, livePixels, startCenterCubeEdge+20, startCenterCubeEdge-2)
			addLife(img, livePixels, startCenterCubeEdge+20, startCenterCubeEdge-1)
			addLife(img, livePixels, startCenterCubeEdge+20, startCenterCubeEdge)

			addLife(img, livePixels, startCenterCubeEdge+21, startCenterCubeEdge-2)
			addLife(img, livePixels, startCenterCubeEdge+21, startCenterCubeEdge-1)
			addLife(img, livePixels, startCenterCubeEdge+21, startCenterCubeEdge)

			addLife(img, livePixels, startCenterCubeEdge+22, startCenterCubeEdge-3)
			addLife(img, livePixels, startCenterCubeEdge+22, startCenterCubeEdge+1)

			addLife(img, livePixels, startCenterCubeEdge+24, startCenterCubeEdge-4)
			addLife(img, livePixels, startCenterCubeEdge+24, startCenterCubeEdge-3)
			addLife(img, livePixels, startCenterCubeEdge+24, startCenterCubeEdge+1)
			addLife(img, livePixels, startCenterCubeEdge+24, startCenterCubeEdge+2)

			addLife(img, livePixels, startCenterCubeEdge+34, startCenterCubeEdge-1)
			addLife(img, livePixels, startCenterCubeEdge+34, startCenterCubeEdge-2)
			addLife(img, livePixels, startCenterCubeEdge+35, startCenterCubeEdge-1)
			addLife(img, livePixels, startCenterCubeEdge+35, startCenterCubeEdge-2)

			// img.Set(startCenterCubeEdge, startCenterCubeEdge, blueColor)
			// img.Set(startCenterCubeEdge+1, startCenterCubeEdge, blueColor)
			// img.Set(startCenterCubeEdge, startCenterCubeEdge+1, blueColor)
			// img.Set(startCenterCubeEdge+1, startCenterCubeEdge+1, blueColor)

			// img.Set(startCenterCubeEdge+8, startCenterCubeEdge, blueColor)
			// img.Set(startCenterCubeEdge+8, startCenterCubeEdge+1, blueColor)
			// img.Set(startCenterCubeEdge+8, startCenterCubeEdge+2, blueColor)

			// img.Set(startCenterCubeEdge+9, startCenterCubeEdge-1, blueColor)
			// img.Set(startCenterCubeEdge+9, startCenterCubeEdge+3, blueColor)

			// img.Set(startCenterCubeEdge+10, startCenterCubeEdge-2, blueColor)
			// img.Set(startCenterCubeEdge+10, startCenterCubeEdge+4, blueColor)

			// img.Set(startCenterCubeEdge+11, startCenterCubeEdge-2, blueColor)
			// img.Set(startCenterCubeEdge+11, startCenterCubeEdge+4, blueColor)

			// img.Set(startCenterCubeEdge+12, startCenterCubeEdge+1, blueColor)

			// img.Set(startCenterCubeEdge+13, startCenterCubeEdge-1, blueColor)
			// img.Set(startCenterCubeEdge+13, startCenterCubeEdge+3, blueColor)

			// img.Set(startCenterCubeEdge+14, startCenterCubeEdge, blueColor)
			// img.Set(startCenterCubeEdge+14, startCenterCubeEdge+1, blueColor)
			// img.Set(startCenterCubeEdge+14, startCenterCubeEdge+2, blueColor)

			// img.Set(startCenterCubeEdge+17, startCenterCubeEdge, blueColor)
			// img.Set(startCenterCubeEdge+17, startCenterCubeEdge-1, blueColor)
			// img.Set(startCenterCubeEdge+17, startCenterCubeEdge-2, blueColor)

			// img.Set(startCenterCubeEdge+18, startCenterCubeEdge, blueColor)
			// img.Set(startCenterCubeEdge+18, startCenterCubeEdge-1, blueColor)
			// img.Set(startCenterCubeEdge+18, startCenterCubeEdge-2, blueColor)

			// img.Set(startCenterCubeEdge+19, startCenterCubeEdge-3, blueColor)
			// img.Set(startCenterCubeEdge+19, startCenterCubeEdge+1, blueColor)

			// img.Set(startCenterCubeEdge+29, startCenterCubeEdge-1, blueColor)
			// img.Set(startCenterCubeEdge+29, startCenterCubeEdge-2, blueColor)
			// img.Set(startCenterCubeEdge+29, startCenterCubeEdge-1, blueColor)
			// img.Set(startCenterCubeEdge+29, startCenterCubeEdge-2, blueColor)

		} else {
			var pixelsWithLivingNeighbors map[string]*lifeCount
			pixelsWithLivingNeighbors = make(map[string]*lifeCount)

			for _, livePixel := range livePixels {
				var pixelBefore = &Life{x: livePixel.x - 1, y: livePixel.y}
				var pixelAfter = &Life{x: livePixel.x + 1, y: livePixel.y}

				var pixelAbove = &Life{x: livePixel.x, y: livePixel.y + 1}
				var pixelBelow = &Life{x: livePixel.x, y: livePixel.y - 1}

				var pixelTopRight = &Life{x: livePixel.x - 1, y: livePixel.y + 1}
				var pixelTopLeft = &Life{x: livePixel.x - 1, y: livePixel.y - 1}

				var pixelBottomRight = &Life{x: livePixel.x + 1, y: livePixel.y + 1}
				var pixelBottomLeft = &Life{x: livePixel.x + 1, y: livePixel.y - 1}

				addPixelToNeighbors(pixelBefore, pixelsWithLivingNeighbors)
				addPixelToNeighbors(pixelAfter, pixelsWithLivingNeighbors)

				addPixelToNeighbors(pixelAbove, pixelsWithLivingNeighbors)
				addPixelToNeighbors(pixelBelow, pixelsWithLivingNeighbors)

				addPixelToNeighbors(pixelTopRight, pixelsWithLivingNeighbors)
				addPixelToNeighbors(pixelTopLeft, pixelsWithLivingNeighbors)

				addPixelToNeighbors(pixelBottomRight, pixelsWithLivingNeighbors)
				addPixelToNeighbors(pixelBottomLeft, pixelsWithLivingNeighbors)
			}

			var newLivePixels map[string]*Life
			newLivePixels = make(map[string]*Life)

			for _, neighbor := range pixelsWithLivingNeighbors {
				pixel := neighbor.life
				if _, exists := livePixels[pixel.GetKey()]; exists {
					if neighbor.count == 2 || neighbor.count == 3 {
						life := neighbor.life
						newLivePixels[life.GetKey()] = &life
						img.Set(life.x, life.y, blueColor)
					}
				} else {
					if neighbor.count == 3 {
						life := neighbor.life
						newLivePixels[life.GetKey()] = &life
						img.Set(life.x, life.y, blueColor)
					}
				}

			}

			livePixels = newLivePixels

			if len(livePixels) >= 128*128 || len(livePixels) == 0 {
				break
			}
		}
	}

	f, err := os.OpenFile("actual_representation_of_corona.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

type lifeCount struct {
	life  Life
	count int
}

func addPixelToNeighbors(pixel *Life, neighbors map[string]*lifeCount) {
	if _, exists := neighbors[pixel.GetKey()]; exists {
		neighbors[pixel.GetKey()].count++
	} else {
		neighbors[pixel.GetKey()] = &lifeCount{
			life:  *pixel,
			count: 1,
		}
	}
}

func addLife(img *image.Paletted, livePixels map[string]*Life, x int, y int) {
	var life = &Life{x: x, y: y}
	livePixels[life.GetKey()] = life

	img.Set(x, y, blueColor)
}
