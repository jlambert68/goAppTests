package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// Assemble the GUI object for the complete MagicalTable-page
func (p *MagicTable) AssembleMagicalTableObject() (app.HTMLDiv, error) {
	var err error
	err = nil

	// Generate test-text to see if data can be sent in to this object
	testText := app.H1().Text(p.canBeAnyText)

	// Create DropDown for choosing which table date to show/add/edit
	tableTypeSelector, err := p.CreateTableTypeSelectorObject()
	if err != nil {
		fmt.Println("Problem generating tableTypeSelector for table")
	}

	// Generate Alert messages
	alertMessages, err := p.GenerateAlertMessages()
	if err != nil {
		fmt.Println("Problem generating PopUp")
	}

	// Generate PopUp and Button for opening (no good solution but I don't know how to open Bootstrap-popup in another way
	openPopUp, popUpp, err := p.CreatePopUp()
	if err != nil {
		fmt.Println("Problem generating PopUp")
	}

	// Generate SearchInDB bar
	searchInput, err := p.CreateSearchBar()
	if err != nil {
		fmt.Println("Problem generating SearchInDB Bar")
	}

	// Create table rows
	err = p.UpdateRowNodes()
	if err != nil {
		fmt.Println("Problem generating rows for table")
	}

	// Dynamically create column headers for MagicTable
	err = p.UpdateColumnsNodes()
	if err != nil {
		fmt.Println("Problem generating columns for table")
	}

	// Create the Edit-boxed to be used for doing New, Edit and Delete in table
	editRows, err := p.CreateEditBoxes()
	if err != nil {
		fmt.Println("Problem generating Edit-boxes")
	}

	// Create buttons for New, Edit and Delete
	newEditDeleteButtons, err := p.CreateButtonsForNewEditDelete()
	if err != nil {
		fmt.Println("Problem generating buttons for New, Edit and Delete")
	}

	// Create buttons for Save and Cancel
	saveCancelButtons, err := p.CreateButtonsForSaveCancel()
	if err != nil {
		fmt.Println("Problem generating buttons for Save and Cancel")
	}

	// Build the table
	magicTableHeaderAndDataRows := app.Table().Class("table").Body(
		app.Tr().Body(p.tableColumnNodes...),
		app.TBody().Body(p.tableRowsNodes...),
		//app.TBody().Body(nodes2...),
	)

	// Create the full magicTable GUI object

	magicTableRenderedObject := app.Div().Body(
		app.Div().
			Class("container-fluid").
			Body(
				app.Div().Body(alertMessages...),
				testText,
				openPopUp,
				popUpp,
				app.Div().
					Body(
						tableTypeSelector),
				//app.If(p.StateCheckToShowBaseButtons() == true,
				app.Div().
					//Style("color", "#ff0000").
					//Style("border", "solid").
					Style("padding", "10px").
					Body(newEditDeleteButtons...),
				//).Else(
				app.Div().
					Class("row").
					Body(editRows...),
				app.Div().
					//Style("color", "#ffff00").
					//Style("border", "solid").
					Style("padding", "10px").
					Body(saveCancelButtons...)),
		//	),
		searchInput,
		app.Div().
			Class("container-fluid").
			Body(
				magicTableHeaderAndDataRows,
			),
	)

	return magicTableRenderedObject, err
}
