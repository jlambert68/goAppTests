package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/060_golangWasmMultiProto_goAppv8/protos/api"
)

// H1

type TempTimeStampAsH1 struct {
	app.Compo
	manager   *MyManager
	timeStamp string
}

func (p *TempTimeStampAsH1) SetManager(manager *MyManager) {
	p.manager = manager
}

func (p *TempTimeStampAsH1) Render() app.UI {
	serverTimeMessage, err := api.CallApiGetTime(api.EmptyParameter{})

	if err != nil {
		fmt.Println("GetTime Error:", err)
	}
	fmt.Println("func (p *TempTimeStampAsH1) Render() app.UI {\n")

	p.timeStamp = serverTimeMessage.String()

	return app.H1().Text("TempTimeStampAsH1: OOO " + p.timeStamp)
}

// H1b

type TempTimeStampAsH1b struct {
	app.Compo
	timeStamp string
}

func (p *TempTimeStampAsH1b) Render() app.UI {
	serverTimeMessage, err := api.CallApiGetTime(api.EmptyParameter{})

	if err != nil {
		fmt.Println("GetTime Error:", err)
	}

	p.timeStamp = serverTimeMessage.String()

	return app.H1().Text("TempTimeStampAsH1b: " + p.timeStamp)
}

// H2

type TempTimeStampAsH2 struct {
	app.Compo
	manager   *MyManager
	timeStamp string
	myText    string
}

func (p *TempTimeStampAsH2) SetManager(manager *MyManager) {
	p.manager = manager
}

func (p *TempTimeStampAsH2) Render() app.UI {
	serverTimeMessage, err := api.CallApiGetTime(api.EmptyParameter{})

	if err != nil {
		fmt.Println("GetTime Error:", err)
	}

	p.timeStamp = serverTimeMessage.String()

	return app.H1().Text("TempTimeStampAsH2: " + p.myText + p.timeStamp)
}
