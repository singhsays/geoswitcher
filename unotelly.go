package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type UnoTelly struct {
	apiKey string
}

func (s *UnoTelly) Server() string {
	return "103.1.187.68"
}

func (s *UnoTelly) UpdateAddress() error {
	if s.apiKey == "" {
		return fmt.Errorf("No apikey found.")
	}
	params := url.Values{
		"user_hash": []string{s.apiKey},
	}
	req, err := http.NewRequest("GET", "http://www.unotelly.com/unodns/auto_auth/hash_update/updateip.php", nil)
	req.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(req)
	log.Printf("Updated Address: %s", resp.Status)
	return err
}
