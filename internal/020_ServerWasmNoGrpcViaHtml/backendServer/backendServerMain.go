package backendServerMain

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"net/http"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World!")
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func BackendServerMain() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &hello{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// On the server-side, RunWhenOnBrowser() does nothing, which allows the
	// writing of server logic without needing precompiling instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	/*
		http.Handle("/", &app.Handler{
			Name:        "Hello",
			Description: "An Hello World! example",
		})

		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal(err)
		}

	*/
	mux := http.NewServeMux()

	app := &app.Handler{
		Name:        "HelloMuxx",
		Description: "An Hello World! example",
		//		Styles: []string{"bootstrap.css"},
	}

	mux.HandleFunc("/app.wasm", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/app.wasm")
	})

	//mux.HandleFunc("/bootstrap.css", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "external/com_github_bootstrap/file/bootstrap.css")
	//})

	// Handle API
	//api.RegisterApiHTTPMux(mux, &server.Server{})

	// Handle go-app
	mux.Handle("/", app)

	fmt.Println("starting local server on http://localhost:7001")
	log.Fatal(http.ListenAndServe(":7001", mux))

}
