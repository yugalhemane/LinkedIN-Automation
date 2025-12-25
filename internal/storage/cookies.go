package storage

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod/lib/proto"
)

const CookieFile = "cookies.json"

func SaveCookies(cookies []*proto.NetworkCookie) error {
	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(CookieFile, data, 0600)
}

func LoadCookies() ([]*proto.NetworkCookie, error) {
	data, err := os.ReadFile(CookieFile)
	if err != nil {
		return nil, err
	}

	var cookies []*proto.NetworkCookie
	err = json.Unmarshal(data, &cookies)
	return cookies, err
}
