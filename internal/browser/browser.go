package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Browser struct {
	Browser *rod.Browser
	Page    *rod.Page
}

func New(headless bool) (*Browser, error) {
	l := launcher.New().
		Headless(headless).
		Leakless(false). // 🔑 DISABLE leakless
		Set("start-maximized").
		Set("disable-blink-features", "AutomationControlled")

	url := l.MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage("")

	return &Browser{
		Browser: browser,
		Page:    page,
	}, nil
}
