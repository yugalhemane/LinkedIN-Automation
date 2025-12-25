package message

import (
	"log"
	"strings"
	"time"

	"linkedin-automation-poc/internal/stealth"
	"linkedin-automation-poc/internal/storage"

	"github.com/go-rod/rod"
)

type Messenger struct {
	Page  *rod.Page
	State *storage.State
}

func New(page *rod.Page, state *storage.State) *Messenger {
	return &Messenger{
		Page:  page,
		State: state,
	}
}

// SendMessage sends a follow-up message to an accepted connection
func (m *Messenger) SendMessage(profileURL, template string) {
	if m.State.IsMessaged(profileURL) {
		return
	}

	log.Println("Opening profile for messaging:", profileURL)
	m.Page.MustNavigate(profileURL)
	time.Sleep(4 * time.Second)

	// Check if "Message" button exists (means connected)
	msgBtn, err := m.Page.Timeout(5 * time.Second).
		ElementR("button", "Message")

	if err != nil {
		log.Println("Message button not available (not connected):", profileURL)
		return
	}

	// Click Message
	msgBtn.MustClick()
	time.Sleep(2 * time.Second)

	// Find message box
	input, err := m.Page.Timeout(5 * time.Second).
		Element("div[role='textbox']")

	if err != nil {
		log.Println("Message input not found")
		return
	}

	// Personalize message (basic dynamic variable)
	name := strings.Split(strings.Trim(profileURL, "/"), "/")
	recipient := name[len(name)-1]

	message := strings.ReplaceAll(template, "{{name}}", recipient)

	typer := stealth.NewHumanTyper()
	typer.TypeElement(input, message)

	time.Sleep(1 * time.Second)

	// Send message
	m.Page.MustElementR("button", "Send").MustClick()

	log.Println("Message sent to:", profileURL)

	m.State.MarkMessaged(profileURL)
	_ = m.State.Save()

	time.Sleep(10 * time.Second)
}
