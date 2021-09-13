package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"goAppTest1/cmd/050_golangWasmMultiProtoServices/protos/api"
)

func main() {
	manager := &Manager{
		searchBar:     &SearchBar{},
		instanceTable: &InstanceTable{},
		timeStampAsH1: &TimeStampAsH1{},
	}

	manager.searchBar.SetManager(manager)
	manager.instanceTable.SetManager(manager)
	manager.timeStampAsH1.SetManager(manager)

	app.Route("/", manager)
	app.Run()
}

// Manager is the main controller of this application, also the root Body
type Manager struct {
	app.Compo
	searchBar     *SearchBar
	instanceTable *InstanceTable
	timeStampAsH1 *TimeStampAsH1
}

func (h *Manager) Render() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.Nav().Class("navbar navbar-expand-lg navbar-light bg-light").Body(
				h.searchBar,
			),
		),
		app.Div().Class("container-fluid").Body(
			h.instanceTable,
		),
		app.Div().Class("container-fluid").Body(
			h.timeStampAsH1,
		),
	)
}

func (h *Manager) Search(q string) []*api.Instance {
	instances, err := api.CallApiSearch(api.SearchRequest{
		Query: q,
	})

	if err != nil {
		fmt.Println("SearchInDB Error:", err)
		return []*api.Instance{}
	}

	return instances.Instances
}

func (h *Manager) GetTime() string {

	serverTimeMessage, err := api.CallApiGetTime(api.EmptyParameter{})

	if err != nil {
		fmt.Println("GetTime Error:", err)
		return serverTimeMessage.String()
	}

	return serverTimeMessage.String()
}

func (h *Manager) UpdateInstances(q string) {
	instances := h.Search(q)
	h.instanceTable.instances = instances
	h.instanceTable.Update()
	h.timeStampAsH1.timeStampAsH1 = h.GetTime() //callBackEnd()
	h.timeStampAsH1.Update()
}
