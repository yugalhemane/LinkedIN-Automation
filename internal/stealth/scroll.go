package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func RandomScroll(page *rod.Page) {
	scrollSteps := 3 + rand.Intn(5)

	for i := 0; i < scrollSteps; i++ {
		scrollBy := 200 + rand.Intn(400)

		page.MustEval(`(y) => window.scrollBy(0, y)`, scrollBy)

		time.Sleep(time.Duration(500+rand.Intn(800)) * time.Millisecond)
	}

	// Occasional scroll back up
	if rand.Intn(2) == 1 {
		page.MustEval(`() => window.scrollBy(0, -300)`)
		time.Sleep(400 * time.Millisecond)
	}
}
