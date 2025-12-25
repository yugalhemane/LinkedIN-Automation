package stealth

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Point struct {
	X float64
	Y float64
}

type MouseMover struct {
	Steps int
}

func NewMouseMover() *MouseMover {
	return &MouseMover{
		Steps: 30 + rand.Intn(20), // variable smoothness
	}
}

// Cubic Bezier curve function
func bezier(t float64, p0, p1, p2, p3 float64) float64 {
	return math.Pow(1-t, 3)*p0 +
		3*math.Pow(1-t, 2)*t*p1 +
		3*(1-t)*math.Pow(t, 2)*p2 +
		math.Pow(t, 3)*p3
}

// Move mouse using Bezier curve
func (m *MouseMover) Move(page *rod.Page, from, to Point) {
	// Random control points to avoid straight lines
	cp1 := Point{
		X: from.X + rand.Float64()*120 - 60,
		Y: from.Y + rand.Float64()*120 - 60,
	}
	cp2 := Point{
		X: to.X + rand.Float64()*120 - 60,
		Y: to.Y + rand.Float64()*120 - 60,
	}

	for i := 0; i <= m.Steps; i++ {
		t := float64(i) / float64(m.Steps)

		x := bezier(t, from.X, cp1.X, cp2.X, to.X)
		y := bezier(t, from.Y, cp1.Y, cp2.Y, to.Y)

		// Correct Rod API
		page.Mouse.MoveTo(proto.Point{
			X: x,
			Y: y,
		})


		// Variable human-like speed
		time.Sleep(time.Duration(5+rand.Intn(15)) * time.Millisecond)
	}
}
