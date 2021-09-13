package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/060_golangWasmMultiProto_goAppv8/protos/api"
)

var tempTimeStampAsH2__b = &TempTimeStampAsH2{myText: "Startvärde-1"}
var tempTimeStampAsH2__c = &TempTimeStampAsH2{myText: "Startvärde-2"}
var resultTable = &InstanceTable{}

var manager = &Manager{
	searchBar: SearchBar{},
	//instanceTable: InstanceTable{},
	timeStampAsH1:        TimeStampAsH1{},
	instanceTablePointer: &InstanceTable{},
}

func main() {

	manager.searchBar.SetManager(manager)
	//manager.instanceTable.SetManager(manager)
	manager.timeStampAsH1.SetManager(manager)
	manager.instanceTablePointer.SetManager(manager)

	app.Route("/", manager)
	app.Route("/2", mymanager)
	app.Route("/3", &TempTimeStampAsH1{})
	app.Route("/4", &TempTimeStampAsH2{})
	app.Route("/searchBar", &SearchBar{})
	app.RunWhenOnBrowser()
}

// Manager is the main controller of this application, also the root Body

type Manager struct {
	app.Compo
	searchBar            SearchBar
	instanceTable        InstanceTable
	timeStampAsH1        TimeStampAsH1
	instanceTablePointer *InstanceTable
}

type searchAndResultTableType struct {
	searchBar     *SearchBar
	instanceTable *InstanceTable
}

var searchAndResultTable = &searchAndResultTableType{searchBar: &SearchBar{}, instanceTable: &InstanceTable{}}

var mymanager = &MyManager{}

type MyManager struct {
	app.Compo
}

func (h *MyManager) Render() app.UI {
	serverTimeMessage, err := api.CallApiGetTime(api.EmptyParameter{})

	if err != nil {
		fmt.Println("GetTime Error:", err)
	}

	serverTime := serverTimeMessage.String()
	fmt.Println("clientCompiledTimeStamp: " + clientCompiledTimeStamp)
	fmt.Println("TimeFromServer. " + serverTime)

	return app.Div().Body(
		app.Div().Class("container-fluid").Body(
			tempTimeStampAsH2__b,
		),
		app.Div().Class("container-fluid").Body(
			tempTimeStampAsH2__c,
		),

		app.Header().Body(
			app.Nav().Class("navbar navbar-expand-lg navbar-light bg-light").Body(
				searchAndResultTable.searchBar,
			),
		),
		app.Header().Body(
			app.Nav().Class("navbar navbar-expand-lg navbar-light bg-light").Body(
				searchAndResultTable.instanceTable,
			),
		),
	)

}

func (h *Manager) Render() app.UI {

	fmt.Println("clientCompiledTimeStamp: " + clientCompiledTimeStamp)
	fmt.Println("TimeFromServer. " + h.GetTime())
	//return app.H1().Text("Hello World!")
	//tempTimeStampAsH2__b.myText ="MinNyaTest"

	return app.Div().Body(
		app.Div().Class("container-fluid").Body(
			tempTimeStampAsH2__b,
		),
		app.Div().Class("container-fluid").Body(
			tempTimeStampAsH2__c,
		),
		app.Div().Class("container-fluid").Body(
			&h.timeStampAsH1,
		),

		app.Header().Body(
			app.Nav().Class("navbar navbar-expand-lg navbar-light bg-light").Body(
				&h.searchBar,
			),
		),
		app.Header().Body(
			app.Nav().Class("navbar navbar-expand-lg navbar-light bg-light").Body(
				resultTable,
			),
		),

		app.Div().Class("container-fluid").Body(
			&h.instanceTable,
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

func (h *searchAndResultTableType) Search(q string) []*api.Instance {
	instances, err := api.CallApiSearch(api.SearchRequest{
		Query: q,
	})

	if err != nil {
		fmt.Println("SearchInDB Error:", err)
		return []*api.Instance{}
	}

	return instances.Instances
}

func (h *searchAndResultTableType) GetTime() string {

	serverTimeMessage, err := api.CallApiGetTime(api.EmptyParameter{})

	if err != nil {
		fmt.Println("GetTime Error:", err)
		return serverTimeMessage.String()
	}

	return serverTimeMessage.String()
}

var tempInstances []*api.Instance

func (h *Manager) UpdateInstances(q string) {
	fmt.Println("func (h *Manager) RetrieveTableDataFromDB(q string) { " + q)
	instances := h.Search(q)
	tempInstances = instances
	fmt.Println("qqqqqqqqqqqqqqqqqqqqqq")
	fmt.Println(instances)
	fmt.Println("sssssssssssssssssss")
	resultTable.instances = instances
	//myI := h.instanceTable
	//myI.instances = instances
	//h.instances = &instances
	fmt.Println("2222222222222")
	resultTable.Update()
	//h.instanceTable.instances = instances
	fmt.Println("wwwwwwwwwww")
	//h.instanceTable.Update()
	//h.instanceTablePointer.Update()
	fmt.Println("cccccccccccccccccccccc")
	//h.timeStampAsH1.timeStampAsH1 = h.GetTime() //callBackEnd()
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	//h.timeStampAsH1.Update()
	fmt.Println("aaaaaaaaaaaaa")
	tempTimeStampAsH2__b.myText = "Hej Hej-AAAA"
	tempTimeStampAsH2__c.myText = "Hej Hej-BBBB"
	fmt.Println("cccccccccccccccccccccc")
	tempTimeStampAsH2__b.Update()
	tempTimeStampAsH2__c.Update()
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	//manager.Update()
}

func (h *MyManager) UpdateInstances(q string) {
	fmt.Println("func (h *Manager) RetrieveTableDataFromDB(q string) { " + q)
	instances := searchAndResultTable.Search(q)
	tempInstances = instances
	fmt.Println("qqqqqqqqqqqqqqqqqqqqqq")
	fmt.Println(instances)
	fmt.Println("sssssssssssssssssss")
	searchAndResultTable.instanceTable.instances = instances
	//	resultTable.instances = instances
	//myI := h.instanceTable
	//myI.instances = instances
	//h.instances = &instances
	fmt.Println("2222222222222")
	searchAndResultTable.instanceTable.Update()
	//resultTable.Update()

	//h.instanceTable.instances = instances
	fmt.Println("wwwwwwwwwww")
	//h.instanceTable.Update()
	//h.instanceTablePointer.Update()
	fmt.Println("cccccccccccccccccccccc")
	//h.timeStampAsH1.timeStampAsH1 = h.GetTime() //callBackEnd()
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	//h.timeStampAsH1.Update()
	fmt.Println("aaaaaaaaaaaaa")
	tempTimeStampAsH2__b.myText = "Hej Hej-AAAA"
	tempTimeStampAsH2__c.myText = "Hej Hej-BBBB"
	fmt.Println("cccccccccccccccccccccc")
	tempTimeStampAsH2__b.Update()
	tempTimeStampAsH2__c.Update()
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	//manager.Update()
}
