package handler

import (
	"github.com/TimothyYe/godns"
	"github.com/TimothyYe/godns/handler/cloudflare"
)

// IHandler is the interface for all DNS handlers
type IHandler interface {
	SetConfiguration(*godns.Settings)
	DomainLoop(domain *godns.Domain, panicChan chan<- godns.Domain)
}

// CreateHandler creates DNS handler by different providers
func CreateHandler(provider string) IHandler {
	var handler IHandler

	switch provider {
	case godns.CLOUDFLARE:
		handler = IHandler(&cloudflare.Handler{})
	}

	return handler
}
