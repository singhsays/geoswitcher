package main

import (
	"fmt"
	"log"
)

type Service interface {
	Server() string
	UpdateAddress() error
}

func NewService(name string) (Service, error) {
	key, ok := keys[name]
	switch *service {
	case "getflix":
		if ok {
			return &GetFlix{key}, nil
		}
		return &GetFlix{}, nil
	case "unotelly":
		if ok {
			return &UnoTelly{key}, nil
		}
		return &UnoTelly{}, nil
	case "dns4me":
		if ok {
			return &DNS4Me{key}, nil
		}
		return &DNS4Me{}, nil
	}
	return nil, fmt.Errorf("Invalid service %s.", name)
}

type GenericService struct {
	server string
}

func (s *GenericService) Server() string {
	return s.server
}

func (s *GenericService) UpdateAddress() error {
	log.Printf("UpdateAddress Not Implemented for Generic Service.")
	return nil
}
