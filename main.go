package main

import (
	"github.com/fogleman/gg"
	"math"
)

type Point struct {
	X, Y int
}

func main() {
	dc := gg.NewContext(100, 100)

	setColor(dc, "white")
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.Fill()

	drawAxis(dc, "black")
	brezCirc(dc, Point{-10, 10}, 20, "teal")
	brez(dc, Point{30, 80}, Point{12, -5}, "pink")
	brezArc(dc, Point{0, 0}, 50, 30, 60, "")
	polygon(dc, []Point{{0, 0}, {10, 10}, {35, 10}, {-10, 40}}, "")

	err := dc.SavePNG("out.png")
	if err != nil {
		panic(err)
	}
}

func polygon(dc *gg.Context, pointArray []Point, color string) {
	for i := 0; i < len(pointArray)-1; i++ {
		brez(dc, pointArray[i], pointArray[i+1], color)
	}
	brez(dc, pointArray[len(pointArray)-1], pointArray[0], color)
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func trans(dc *gg.Context, a Point) (x int, y int) {
	return a.X + dc.Width()/2, -a.Y + dc.Height()/2
}

func setColor(dc *gg.Context, color string) {
	switch {
	case color == "white":
		dc.SetRGB255(255, 255, 255)
	case color == "gray":
		dc.SetRGB255(123, 123, 123)
	case color == "pink":
		dc.SetRGB255(255, 192, 203)
	case color == "teal":
		dc.SetRGB255(0, 128, 128)
	case color == "black":
		dc.SetRGB255(0, 0, 0)
	default:
		dc.SetRGB255(255, 165, 0)
	}
}

func drawAxis(dc *gg.Context, color string) {
	setColor(dc, color)
	for i := 0; i < dc.Width(); i++ {
		dc.SetPixel(dc.Height()/2, i)
	}
	for i := 0; i < dc.Height(); i++ {
		dc.SetPixel(i, dc.Width()/2)
	}
}

func brezArc(dc *gg.Context, zerop Point, r float64, startAngle float64, endAngle float64, color string) {
	setColor(dc, color)
	if startAngle > endAngle {
		startAngle, endAngle = endAngle, startAngle
	}
	startAngle *= math.Pi / 180
	endAngle *= math.Pi / 180
	x := int(math.Cos(endAngle) * r)
	y := int(math.Sin(endAngle) * r)
	delta := 2 * (1 - int(r))
	for y >= int(math.Sin(startAngle)*r) {
		dc.SetPixel(trans(dc, Point{x + zerop.X, y + zerop.Y}))
		if delta < 0 && 2*delta+2*y-1 <= 0 {
			x++
			delta += 2*x - 1
		} else if delta > 0 && 2*delta-2*x-1 > 0 {
			y--
			delta += 1 - 2*y
		} else {
			x++
			y--
			delta += 2*x - 2*y + 2
		}
	}
}

func brezCirc(dc *gg.Context, c Point, r int, color string) {
	setColor(dc, color)
	var x int
	y := r
	delta := 2 - 2 - y
	for y >= 0 {
		dc.SetPixel(trans(dc, Point{c.X + x, c.Y + y}))
		dc.SetPixel(trans(dc, Point{c.X + x, c.Y - y}))
		dc.SetPixel(trans(dc, Point{c.X - x, c.Y + y}))
		dc.SetPixel(trans(dc, Point{c.X - x, c.Y - y}))

		if delta < 0 {
			buffer := 2*delta + 2*y - 1
			x++
			if buffer <= 0 {
				delta += 2*x + 1
			} else {
				y--
				delta += 2*x - 2*y + 2
			}
			continue
		}

		if delta > 0 {
			buffer := 2*delta - 2*x - 1
			y--
			if buffer > 0 {
				delta += -2*y + 1
			} else {
				x++
				delta += 2*x - 2*y + 2
			}
			continue
		}

		if delta == 0 {
			x++
			y--
			delta += 2*x - 2*y + 2
		}
	}
}

func brez(dc *gg.Context, a Point, b Point, color string) {
	setColor(dc, color)
	x := a.X
	y := a.Y
	dx := math.Abs(float64(b.X - a.X))
	dy := math.Abs(float64(b.Y - a.Y))
	s1 := sign(b.X - a.X)
	s2 := sign(b.Y - a.Y)
	var strg int
	if dy > dx {
		dx, dy = dy, dx
		strg = 1
	}
	e := 2*dy - dx
	for i := 1; i < int(dx); i++ {
		dc.SetPixel(trans(dc, Point{x, y}))
		for e >= 0 {
			if strg == 1 {
				x += s1
			} else {
				y += s2
			}
			e -= 2 * dx
		}
		if strg == 1 {
			y += s2
		} else {
			x += s1
		}
		e += 2 * dy
	}
}
