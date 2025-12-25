package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

type Typer struct {
	minDelay time.Duration
	maxDelay time.Duration
}

func NewHumanTyper() *Typer {
	return &Typer{
		minDelay: 40 * time.Millisecond,
		maxDelay: 140 * time.Millisecond,
	}
}

// randomDelay returns a human-like typing delay
func (t *Typer) randomDelay() time.Duration {
	diff := t.maxDelay - t.minDelay
	return t.minDelay + time.Duration(rand.Int63n(int64(diff)))
}

// Type types text into an element selected by selector
func (t *Typer) Type(page *rod.Page, selector, text string) {
	el := page.MustElement(selector)
	for _, char := range text {
		el.MustInput(string(char))
		time.Sleep(t.randomDelay())
	}
}

// TypeElement types text directly into an element
func (t *Typer) TypeElement(el *rod.Element, text string) {
	for _, char := range text {
		el.MustInput(string(char))
		time.Sleep(t.randomDelay())
	}
}
