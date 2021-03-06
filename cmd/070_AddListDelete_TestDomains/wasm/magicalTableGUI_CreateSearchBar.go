package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// Generates the GUI objects for the columns in the table
func (mt *MagicTable) CreateSearchBar() (app.HTMLDiv, error) {
	var err error
	var searchInput app.HTMLDiv
	err = nil
	searchInput = nil

	if mt.searchbarIsVisible == true {
		input := app.Input().
			Class("form-control").
			Value(mt.searchString).
			Placeholder("t2.small").
			AutoFocus(true).
			OnKeyup(mt.MyOnInputChange)

		searchInput = app.Div().
			Class("input-group").
			Body(
				app.Div().
					Class("input-group-prepend").
					Body(app.Span().
						Class("input-group-text").
						Body(app.Text("🔍"))),
				input,
			)
	}

	return searchInput, err
}
