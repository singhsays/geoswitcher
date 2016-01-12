package main

import (
	"fmt"
	"log"
	"net/http"
)

type DNS4Me struct {
	apiKey string
}

func (s *DNS4Me) Server() string {
	return "103.241.0.199"
}

func (s *DNS4Me) UpdateAddress() error {
	if s.apiKey == "" {
		return fmt.Errorf("No apikey found.")
	}
	resp, err := http.Get("https://dns4me.net/user/update_zone_api/" + s.apiKey)
	log.Printf("Updated Address: %s", resp.Status)
	return err
}
