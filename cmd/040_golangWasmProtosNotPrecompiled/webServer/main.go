package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"goAppTest1/cmd/040_golangWasmProtosNotPrecompiled/protos/api"
	"goAppTest1/cmd/040_golangWasmProtosNotPrecompiled/server"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	app := &app.Handler{
		Title:  "EC2Instances",
		Author: "Graham Jenson",
		Styles: []string{"bootstrap.css"},
	}

	mux.HandleFunc("/app.wasm", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "wasm/js_wasm_pure_stripped/app.wasm")
		http.ServeFile(w, r, "web/app.wasm")
	})

	mux.HandleFunc("/bootstrap.css", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "external/com_github_bootstrap/file/bootstrap.css")
		http.ServeFile(w, r, "web/bootstrap.css")
	})

	// Handle API
	api.RegisterApiHTTPMux(mux, &server.Server{})

	// Handle go-app
	mux.Handle("/", app)

	fmt.Println("starting local server on http://localhost:7002")
	log.Fatal(http.ListenAndServe(":7002", mux))
}
