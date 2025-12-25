package search

import (
	"log"
	"strings"
	"time"

	"linkedin-automation-poc/internal/stealth"
	"linkedin-automation-poc/internal/storage"

	"github.com/go-rod/rod"
)

type Searcher struct {
	Page  *rod.Page
	State *storage.State
}

func New(page *rod.Page, state *storage.State) *Searcher {
	return &Searcher{
		Page:  page,
		State: state,
	}
}

func (s *Searcher) SearchByKeyword(keyword string) error {
	searchURL := "https://www.linkedin.com/search/results/people/?keywords=" +
		strings.ReplaceAll(keyword, " ", "%20")

	log.Println("Navigating to search:", searchURL)
	s.Page.MustNavigate(searchURL)
	time.Sleep(5 * time.Second)

	// Scroll to load results
	stealth.RandomScroll(s.Page)

	// Extract profile links
	links := s.Page.MustElements("a[href*='/in/']")

	for _, link := range links {
		href, err := link.Attribute("href")
		if err != nil || href == nil {
			continue
		}

		profileURL := *href
		if !strings.HasPrefix(profileURL, "https://") {
			profileURL = "https://www.linkedin.com" + profileURL
		}

		// Deduplicate
		if s.State.Exists(profileURL) {
			continue
		}

		log.Println("Found profile:", profileURL)
		s.State.AddProfile(profileURL)
	}

	return s.State.Save()
}
