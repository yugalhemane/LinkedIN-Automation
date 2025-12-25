package main

import (
	"log"
	"time"

	"linkedin-automation-poc/internal/auth"
	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/config"
	"linkedin-automation-poc/internal/logger"
	"linkedin-automation-poc/internal/stealth"
	"linkedin-automation-poc/internal/storage"
	"linkedin-automation-poc/internal/search"
	"linkedin-automation-poc/internal/connect"
	"linkedin-automation-poc/internal/message"

	"github.com/go-rod/rod"
)

// isLoggedIn checks for an element that ONLY exists when logged in
func isLoggedIn(page *rod.Page) bool {
	return page.MustHas("img.global-nav__me-photo")
}

func main() {
	cfg := config.Load()
	logr := logger.New()

	logr.Println("Starting LinkedIn Automation POC")

	// Launch browser
	br, err := browser.New(cfg.Headless)
	if err != nil {
		log.Fatal(err)
	}
	logr.Println("Browser launched successfully")

	// Apply stealth fingerprint
	if err := stealth.ApplyFingerprint(br.Page); err != nil {
		log.Fatal(err)
	}
	logr.Println("Browser fingerprint applied")

	// --- AUTH SESSION ---
	session := auth.NewSession(br.Page)
	session.Restore()

	// Navigate to feed
	br.Page.MustNavigate("https://www.linkedin.com/feed")
	time.Sleep(5 * time.Second)

	// 🔒 HARD AUTH VERIFICATION
	if !isLoggedIn(br.Page) {
		logr.Println("Session invalid or expired, starting login flow")

		authClient := auth.New(br.Page)
		if err := authClient.Login(cfg.LinkedInEmail, cfg.LinkedInPassword); err != nil {
			log.Fatal(err)
		}

		br.Page.MustNavigate("https://www.linkedin.com/feed")
		time.Sleep(5 * time.Second)

		if !isLoggedIn(br.Page) {
			log.Fatal("Login failed: still not authenticated")
		}

		logr.Println("Login successful and verified")
	} else {
		logr.Println("Authenticated session verified")
	}

	// Human think time
	stealth.NewDefaultTiming().Sleep()

	// --- LOAD STATE ---
	state, err := storage.LoadState()
	if err != nil {
		log.Fatal(err)
	}

	// --- SEARCH ---
	searcher := search.New(br.Page, state)
	if err := searcher.SearchByKeyword("Golang Developer India"); err != nil {
		log.Fatal(err)
	}
	logr.Println("Search completed")

	// --- CONNECT ---
	connector := connect.New(br.Page, state, cfg.DailyConnectLimit)
	note := "Hi, I came across your profile and would like to connect."

	for profileURL := range state.VisitedProfiles {
		if connector.SentToday >= cfg.DailyConnectLimit {
			break
		}
		connector.SendRequest(profileURL, note)
	}
	logr.Println("Connection flow completed")

	// --- MESSAGE ---
	messenger := message.New(br.Page, state)
	template := "Hi {{name}}, thanks for connecting!"

	for profileURL := range state.ConnectedProfiles {
		if state.IsMessaged(profileURL) {
			continue
		}
		messenger.SendMessage(profileURL, template)
		break // safe: 1 message per run
	}
	logr.Println("Messaging flow completed")

	// Idle behavior
	stealth.RandomScroll(br.Page)

	// Keep browser open
	select {}
}
