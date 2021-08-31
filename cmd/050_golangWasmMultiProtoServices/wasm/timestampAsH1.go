package main

import (
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type TimeStampAsH1 struct {
	app.Compo
	manager       *Manager
	timeStampAsH1 string
}

func (p *TimeStampAsH1) SetManager(manager *Manager) {
	p.manager = manager
}

func (p *TimeStampAsH1) Render() app.UI {
	return app.H1().Text(p.timeStampAsH1)
}
