package stealth

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func ApplyFingerprint(page *rod.Page) error {
	// Set realistic User-Agent via CDP
	err := proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/120.0.0.0 Safari/537.36",
		AcceptLanguage: "en-US,en;q=0.9",
	}.Call(page)
	if err != nil {
		return err
	}

	// Set common laptop viewport
	err = proto.EmulationSetDeviceMetricsOverride{
		Width:             1366,
		Height:            768,
		DeviceScaleFactor: 1,
		Mobile:            false,
	}.Call(page)
	if err != nil {
		return err
	}

	// Remove navigator.webdriver
	page.MustEval(`() => {
	Object.defineProperty(navigator, 'webdriver', {
		get: () => undefined
	});
}`)


	return nil
}
