package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

// H1

type MagicTempTimeStampAsH1 struct {
	app.Compo
	manager   *MagicManager
	timeStamp string
}

func (p *MagicTempTimeStampAsH1) SetManager(manager *MagicManager) {
	p.manager = manager
}

func (p *MagicTempTimeStampAsH1) Render() app.UI {
	serverTimeMessage, err := api.CallApiGetTime(api.EmptyParameter{})

	if err != nil {
		fmt.Println("GetTime Error:", err)
	}
	//fmt.Println("func (p *TempTimeStampAsH1) Render() app.UI {\n")

	p.timeStamp = serverTimeMessage.String()

	return app.H1().Text("TempTimeStampAsH1: OOO " + p.timeStamp)
}
