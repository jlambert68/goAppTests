package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/server"
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

type MagicTable struct {
	app.Compo
	magicTableMetaData []*api.MagicTableColumnMetadata
	instances          []*api.Instance
	tableType          api.MagicTableType
}

type MagicManager struct {
	app.Compo
	magicTable             MagicTable
	magicTempTimeStampAsH1 MagicTempTimeStampAsH1
}

type MagicTempTimeStampAsH1 struct {
	app.Compo
	manager   *MagicManager
	timeStamp string
}

func main() {

	// *******************************
	// Set up and connect to DB

	server.ConnectToDB()

	// Close DB-pool when closing program
	defer server.DbPool.Close()

	// *******************************
	// Set up and start Web Server
	var magicTable = &MagicTable{
		tableType: api.MagicTableType_TestModel,
	}

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

	var magicManager = &MagicManager{
		magicTable:             MagicTable{},
		magicTempTimeStampAsH1: MagicTempTimeStampAsH1{},
	}

	app.Route("/", manager)
	app.Route("/2", myManager)
	app.Route("/3", &TempTimeStampAsH1{})
	app.Route("/4", &TempTimeStampAsH2{})
	app.Route("/searchBar", &SearchBar{})
	app.Route("/magicTable", magicTable)
	app.Route("/magicManager", magicManager)

	mux := http.NewServeMux()

	app := &app.Handler{
		Title:  "Fenix Inception - Dynamic Administrator ",
		Author: "Jonas",
		Styles: []string{"bootstrap.css"},
		Scripts: []string{
			"/web/js/bootstrap.bundle.min.js", // Local script
			"/web/js/alertFunctions.js",
			//"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM\" crossorigin=\"anonymous\"", // Remote script
		},
		BackgroundColor: "#151515",

		Description: "List, wake and scan nodes in a network.",
		Icon: app.Icon{
			Default: "/web/FenixIcon.jpeg",
		},
		ThemeColor: "#151515",
	}

	mux.HandleFunc("/app.wasm", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "wasm/js_wasm_pure_stripped/app.wasm")
		fmt.Println("/app.wasm")
		http.ServeFile(w, r, "web/app.wasm")
	})

	mux.HandleFunc("/bootstrap.css", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "external/com_github_bootstrap/file/bootstrap.css"
		fmt.Println("/bootstrap.css")
		http.ServeFile(w, r, "web/css/bootstrap.css")
	})

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "external/com_github_bootstrap/file/bootstrap.css")
		fmt.Println("/favicon.ico")
		http.ServeFile(w, r, "web/favicon.ico")
	})

	mux.HandleFunc("/bootstrap", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "external/com_github_bootstrap/file/bootstrap.css")
		fmt.Println("/bootstrap.html")
		http.ServeFile(w, r, "web/bootstrap.html")
	})

	// Handle API
	api.RegisterApiHTTPMux(mux, &server.Server{})

	// Handle go-app
	mux.Handle("/", app)

	fmt.Println("starting local server on http://localhost:7003")
	log.Fatal(http.ListenAndServe(":7003", mux))
}
