package auth

import (
	"errors"
	"log"
	"time"

	"linkedin-automation-poc/internal/stealth"

	"github.com/go-rod/rod"
)

type Auth struct {
	Page *rod.Page
}

func New(page *rod.Page) *Auth {
	return &Auth{Page: page}
}

func (a *Auth) Login(email, password string) error {
	// Navigate to LinkedIn login
	a.Page.MustNavigate("https://www.linkedin.com/login")
	time.Sleep(3 * time.Second)

	// Detect login page
	if !a.Page.MustHas("input#username") {
		return errors.New("login page not loaded")
	}

	typer := stealth.NewHumanTyper()

	// Type credentials
	typer.Type(a.Page, "input#username", email)
	time.Sleep(1 * time.Second)
	typer.Type(a.Page, "input#password", password)

	// Submit form
	a.Page.MustElement("button[type=submit]").MustClick()
	time.Sleep(5 * time.Second)

	// Detect security checkpoint (OTP / captcha)
	if a.Page.MustHas("input[name=pin]") ||
		a.Page.MustHas("iframe[src*='captcha']") {

		// Inform user and wait
		log.Println("Security checkpoint detected (OTP / captcha).")
		log.Println("Please complete verification manually in the browser.")
		log.Println("Waiting for login to complete...")

		// Wait until redirected to feed
		a.Page.MustWaitLoad()
		a.Page.MustWaitIdle()

		// Poll until feed is accessible
		a.Page.MustWait(`() => window.location.href.includes("/feed")`)
	}


	// Save cookies after successful login
	session := NewSession(a.Page)
	_ = session.Save()

	return nil
}
