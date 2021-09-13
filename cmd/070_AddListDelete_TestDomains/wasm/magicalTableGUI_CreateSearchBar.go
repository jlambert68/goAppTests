package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// Generates the GUI objects for the columns in the table
func (p *MagicTable) CreateSearchBar() (app.HTMLDiv, error) {
	var err error
	var searchInput app.HTMLDiv
	err = nil
	searchInput = nil

	if p.searchbarIsVisible == true {
		input := app.Input().
			Class("form-control").
			Value(p.searchString).
			Placeholder("t2.small").
			AutoFocus(true).
			OnKeyup(p.MyOnInputChange)

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
