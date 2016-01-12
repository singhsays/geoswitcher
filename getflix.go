package main

import (
	"fmt"
	"log"
	"net/http"
)

type GetFlix struct {
	apiKey string
}

func (s *GetFlix) Server() string {
	return "54.252.183.4"
}

func (s *GetFlix) UpdateAddress() error {
	if s.apiKey == "" {
		return fmt.Errorf("No apikey found.")
	}
	req, err := http.NewRequest("PUT", "https://www.getflix.com.au/api/v1/addresses.json", nil)
	req.SetBasicAuth(s.apiKey, "x")
	resp, err := http.DefaultClient.Do(req)
	log.Printf("Updated Address: %s", resp.Status)
	return err
}
