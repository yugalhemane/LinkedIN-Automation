package connect

import (
	"log"
	"time"

	"linkedin-automation-poc/internal/stealth"
	"linkedin-automation-poc/internal/storage"

	"github.com/go-rod/rod"
)

type Connector struct {
	Page       *rod.Page
	State      *storage.State
	DailyLimit int
	SentToday  int
}

func New(page *rod.Page, state *storage.State, limit int) *Connector {
	return &Connector{
		Page:       page,
		State:      state,
		DailyLimit: limit,
	}
}

// SendRequest attempts card-level connect first, then profile-level
func (c *Connector) SendRequest(profileURL, note string) {
	if c.SentToday >= c.DailyLimit {
		log.Println("Daily connection limit reached")
		return
	}

	if c.State.IsConnected(profileURL) {
		return
	}

	// 1️⃣ Try CARD-LEVEL Connect (preferred)
	if c.tryCardLevelConnect(note) {
		c.markSuccess(profileURL)
		return
	}

	// 2️⃣ Fallback: PROFILE-LEVEL Connect
	c.tryProfileLevelConnect(profileURL, note)
}

// -------------------- CARD LEVEL --------------------

func (c *Connector) tryCardLevelConnect(note string) bool {
	btn, err := c.Page.Timeout(3 * time.Second).
		ElementR("button", `(?i)connect`)
	if err != nil {
		return false
	}

	log.Println("Card-level Connect found")

	stealth.NewMouseMover().Move(
		c.Page,
		stealth.Point{X: 300, Y: 300},
		stealth.Point{X: 450, Y: 350},
	)

	btn.MustClick()
	time.Sleep(1 * time.Second)

	c.tryAddNote(note)
	return true
}

// -------------------- PROFILE LEVEL --------------------

func (c *Connector) tryProfileLevelConnect(profileURL, note string) {
	log.Println("Visiting profile:", profileURL)
	c.Page.MustNavigate(profileURL)
	time.Sleep(4 * time.Second)

	btn, err := c.Page.Timeout(4 * time.Second).
		ElementR("button", `(?i)connect`)
	if err != nil {
		log.Println("No Connect button on profile:", profileURL)
		return
	}

	stealth.NewMouseMover().Move(
		c.Page,
		stealth.Point{X: 350, Y: 350},
		stealth.Point{X: 520, Y: 420},
	)

	btn.MustClick()
	time.Sleep(1 * time.Second)

	c.tryAddNote(note)
	c.markSuccess(profileURL)
}

// -------------------- NOTE HANDLING --------------------

func (c *Connector) tryAddNote(note string) {
	addNoteBtn, err := c.Page.Timeout(2 * time.Second).
		ElementR("button", `(?i)add a note`)
	if err == nil {
		addNoteBtn.MustClick()
		time.Sleep(500 * time.Millisecond)
	}

	textarea, err := c.Page.Timeout(2 * time.Second).Element("textarea")
	if err == nil {
		textarea.MustInput(note)
		time.Sleep(500 * time.Millisecond)

		c.Page.MustElementR("button", `(?i)send`).MustClick()
	}
}

// -------------------- STATE --------------------

func (c *Connector) markSuccess(profileURL string) {
	log.Println("Connection request sent:", profileURL)
	c.State.MarkConnected(profileURL)
	_ = c.State.Save()
	c.SentToday++

	time.Sleep(time.Duration(15+stealthRand()) * time.Second)
}

func stealthRand() int {
	return 10
}
