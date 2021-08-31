package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

type InstanceTable struct {
	app.Compo
	manager   *Manager
	instances []*api.Instance
}

func (p *InstanceTable) SetManager(manager *Manager) {
	p.manager = manager
}

func (p *InstanceTable) Render() app.UI {

	nodes := []app.UI{}
	//for _, i := range p.instances {
	for _, i := range tempInstances {
		nodes = append(nodes, app.Tr().Body(
			app.Td().Body(app.Text(i.Name)),
			app.Td().Body(app.Text(i.InstanceType)),
			app.Td().Body(app.Text(fmt.Sprintf("%v", i.Ecu))),
			app.Td().Body(app.Text(fmt.Sprintf("%v", i.Memory))),
			app.Td().Body(app.Text(i.Network)),
			app.Td().Body(app.Text(i.Price)),
		).OnDblClick(p.MyOnDblClickWrapper(i.Name)))

	}

	return app.Table().Class("table").Body(
		app.Tr().Body(
			app.Th().Scope("col").Body(app.Text("Name")),
			app.Th().Scope("col").Body(app.Text("Instance Type")),
			app.Th().Scope("col").Body(app.Text("ECU")),
			app.Th().Scope("col").Body(app.Text("Mem")),
			app.Th().Scope("col").Body(app.Text("Network")),
			app.Th().Scope("col").Body(app.Text("Price")),
		),
		app.TBody().Body(nodes...),
	)

}

//func (p *InstanceTable)  onClick(ctx app.Context, e app.Event) {
//	fmt.Println("onClick is called")
//}
func (p *InstanceTable) MyOnDblClickWrapper(nameThatWasDoubleClicked string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		fmt.Println("OnDblClick is called::::::" + nameThatWasDoubleClicked)
		fmt.Println("")
	}
}

func (p *InstanceTable) OnDblClickapp(src app.Value, e app.Event) {
	fmt.Println("OnDblClick is called - 1025")
	fmt.Println("src - " + src.JSValue().String())
	fmt.Println("e - " + e.Value.String())

}
