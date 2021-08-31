package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
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
	return app.H1().Text("JOnas_v0915 " + p.timeStampAsH1)
}

type TimeStampAsH1_2 struct {
	app.Compo
	timeStampAsH1 string
}

func (p *TimeStampAsH1_2) Render() app.UI {

	serverTimeMessage, err := api.CallApiGetTime(api.EmptyParameter{})

	if err != nil {
		fmt.Println("GetTime Error:", err)
	}

	p.timeStampAsH1 = serverTimeMessage.String()
	return app.Div().Body(
		app.Div().Class("container-fluid").Body(app.H2().
			Text("JOnas_v666 " + p.timeStampAsH1)))
}
