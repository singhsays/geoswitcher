package main

import (
	"flag"
	"log"
	"os"
)

var (
	service = flag.String("service", "getflix", "dns service to use. one of getflix, unotelly, dns4me")
	server  = flag.String("server", "", "dns server to use instead of a preset service. overrides service.")
	host    = flag.String("host", "router", "name of the router host.")
	path    = flag.String("path", "/jffs/configs/dnsmasq.conf.add", "path to destination config file.")
	apiKey  = flag.String("apikey", "", "api key for service ip address update.")
	config  = flag.String("config", "~/.geoswitcher.json", "a config json file with service api keys.")
	commit  = flag.Bool("commit", true, "whether to switch config or just write to a temp file.")
	update  = flag.Bool("update", true, "whether to update ip address with the service.")
	// Globals
	keys    map[string]string
	domains = [][]string{
		[]string{"netflix.com", "netflix.net", "nflximg.com", "nflxext.com"},
		[]string{"amazon.com"},
		[]string{"abc.com", "go.com"},
		[]string{"cbs.com"},
		[]string{"hulu.com"},
		[]string{"vudu.com"},
		[]string{"staragvod-vh.akamaihd.net", "staragvod1-vh.akamaihd.net"},
		[]string{"getflix.com.au"},
	}
)

func main() {
	var (
		svc Service
		err error
	)
	// Parse flags
	flag.Parse()
	loadKeys()
	// Select a server to use.
	if *server != "" {
		svc = &GenericService{*server}
	} else {
		svc, err = NewService(*service)
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
	// Write a temporary config file.
	configFile, err := writeTemp(createConfig(svc))
	if err != nil {
		log.Fatal("Error writing temp configuration: %s", err)
	}
	log.Printf("Wrote temp configuration: %s", configFile)

	// If commit is true, scp the config to host and restart dnsmasq
	if *commit {
		err = switchConfig(configFile, *host, *path)
		if err != nil {
			log.Fatalf("Error copying config: %s", err)
		}
		err = os.Remove(configFile)
		if err != nil {
			log.Fatalf("Error removing temp config: %s", err)
		}
	}

	// If update is true, update current ip address with the service.
	if *update {
		if err := svc.UpdateAddress(); err != nil {
			log.Fatalf("Error updating IP address: %s", err)
		}
	}
}
