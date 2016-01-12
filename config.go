package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
	"strings"
)

func loadKeys() {
	path := *config
	if path[:2] == "~/" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = strings.Replace(path, "~", dir, 1)
	}
	c, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error reading config file: %s", err)
		return
	}
	err = json.Unmarshal(c, &keys)
	if err != nil {
		log.Printf("Error parsing config file: %s", err)
		return
	}
}

func createConfig(service Service) []string {
	var lines []string
	server := service.Server()
	for _, domain := range domains {
		pattern := strings.Join(domain, "/")
		entry := strings.Join([]string{"server=", pattern, server}, "/")
		lines = append(lines, entry)
	}
	return lines
}

func writeTemp(lines []string) (string, error) {
	t, err := ioutil.TempFile("/tmp", "geoswitcher")
	if err != nil {
		return "", err
	}
	defer t.Close()
	w := bufio.NewWriter(t)
	for _, line := range lines {
		_, err = w.WriteString(line + "\n")
		if err != nil {
			return "", err
		}
	}
	w.Flush()
	return t.Name(), nil
}
