package storage

import (
	"encoding/json"
	"os"
	"sync"
)

const StateFile = "state.json"

type State struct {
	VisitedProfiles   map[string]bool `json:"visited_profiles"`
	ConnectedProfiles map[string]bool `json:"connected_profiles"`
	mu                sync.Mutex
	MessagedProfiles map[string]bool `json:"messaged_profiles"`

}
func NewState() *State {
	return &State{
		VisitedProfiles:   make(map[string]bool),
		ConnectedProfiles: make(map[string]bool),
		MessagedProfiles: make(map[string]bool),
	}
}



func LoadState() (*State, error) {
	state := &State{
		VisitedProfiles: make(map[string]bool),
	}

	data, err := os.ReadFile(StateFile)
	if err != nil {
		return state, nil // first run
	}

	_ = json.Unmarshal(data, state)
	return state, nil
}

func (s *State) Save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(StateFile, data, 0600)
}

func (s *State) AddProfile(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.VisitedProfiles[url] = true
}

func (s *State) Exists(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.VisitedProfiles[url]
}

func (s *State) MarkConnected(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ConnectedProfiles[url] = true
}

func (s *State) IsConnected(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ConnectedProfiles[url]
}

func (s *State) MarkMessaged(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.MessagedProfiles[url] = true
}

func (s *State) IsMessaged(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.MessagedProfiles[url]
}
