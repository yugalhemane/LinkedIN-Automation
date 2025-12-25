package auth

import (
	"log"
	"os"

	"linkedin-automation-poc/internal/storage"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Session struct {
	Page *rod.Page
}

func NewSession(page *rod.Page) *Session {
	return &Session{Page: page}
}

func (s *Session) Restore() bool {
	if _, err := os.Stat(storage.CookieFile); err != nil {
		return false
	}

	cookies, err := storage.LoadCookies()
	if err != nil {
		return false
	}

	for _, c := range cookies {
		_, err := proto.NetworkSetCookie{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Secure:   c.Secure,
			HTTPOnly: c.HTTPOnly,
			Expires:  c.Expires,
		}.Call(s.Page)

		if err != nil {
			log.Println("Failed to restore cookie:", c.Name)
		}
	}

	log.Println("Session cookies restored")
	return true
}

func (s *Session) Save() error {
	cookies, err := s.Page.Browser().GetCookies()
	if err != nil {
		return err
	}
	return storage.SaveCookies(cookies)
}
