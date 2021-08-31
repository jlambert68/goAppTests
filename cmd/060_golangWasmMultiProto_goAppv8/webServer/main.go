package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/060_golangWasmMultiProto_goAppv8/protos/api"
	"goAppTest1/cmd/060_golangWasmMultiProto_goAppv8/server"
	"log"
	"net/http"
)

type TimeStampAsH1 struct {
	app.Compo
	manager       *Manager
	timeStampAsH1 string
}

type TimeStampAsH1_2 struct {
	app.Compo
	timeStampAsH1 string
}

type TempTimeStampAsH1 struct {
	app.Compo
	timeStamp string
}

type TempTimeStampAsH2 struct {
	app.Compo
	timeStamp string
}

type MyManager struct {
	app.Compo
	tempTimeStampAsH1 *TempTimeStampAsH1
	tempTimeStampAsH2 *TempTimeStampAsH2
}

// Manager is the main controller of this application, also the root Body
type Manager struct {
	app.Compo
	searchBar            SearchBar
	instanceTable        InstanceTable
	timeStampAsH1        TimeStampAsH1
	instanceTablePointer *InstanceTable
}

type InstanceTable struct {
	app.Compo
	manager   *Manager
	instances []*api.Instance
}

type SearchBar struct {
	app.Compo
	manager      *Manager
	searchString string
}

func main() {

	manager := &Manager{
		Compo:                app.Compo{},
		searchBar:            SearchBar{},
		instanceTable:        InstanceTable{},
		timeStampAsH1:        TimeStampAsH1{},
		instanceTablePointer: &InstanceTable{},
	}

	myManager := &MyManager{
		tempTimeStampAsH1: &TempTimeStampAsH1{},
		tempTimeStampAsH2: &TempTimeStampAsH2{},
	}

	app.Route("/", manager)
	app.Route("/2", myManager)
	app.Route("/3", &TempTimeStampAsH1{})
	app.Route("/4", &TempTimeStampAsH2{})
	app.Route("/searchBar", &SearchBar{})

	mux := http.NewServeMux()

	app := &app.Handler{
		Title:  "Fenix Inception - Dynamic Administrator ",
		Author: "Jonas",
		Styles: []string{"bootstrap.css"},
	}

	mux.HandleFunc("/app.wasm", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "wasm/js_wasm_pure_stripped/app.wasm")
		fmt.Println("/app.wasm")
		http.ServeFile(w, r, "web/app.wasm")
	})

	mux.HandleFunc("/bootstrap.css", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "external/com_github_bootstrap/file/bootstrap.css"
		fmt.Println("/bootstrap.css")
		http.ServeFile(w, r, "web/bootstrap.css")
	})

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "external/com_github_bootstrap/file/bootstrap.css")
		fmt.Println("/favicon.ico")
		http.ServeFile(w, r, "web/favicon.ico")
	})

	// Handle API
	api.RegisterApiHTTPMux(mux, &server.Server{})

	// Handle go-app
	mux.Handle("/", app)

	fmt.Println("starting local server on http://localhost:7002")
	log.Fatal(http.ListenAndServe(":7002", mux))
}
