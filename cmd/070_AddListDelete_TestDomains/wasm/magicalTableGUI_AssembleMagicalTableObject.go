package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/sirupsen/logrus"
)

// Assemble the GUI object for the complete MagicalTable-page
func (mt *MagicTable) AssembleMagicalTableObject() (app.HTMLDiv, error) {

	mt.logger.WithFields(logrus.Fields{
		"Id":    "aa3ccd2d-2e9c-4b51-8241-9c645fe8632c",
		"Trace": mt.trace(false),
	}).Debug("Entering: AssembleMagicalTableObject()")

	defer func() {
		mt.logger.WithFields(logrus.Fields{
			"Id":    "8455b278-cf1a-4e86-b1bd-0c6d0a1bd366",
			"Trace": mt.trace(false),
		}).Debug("Exiting: AssembleMagicalTableObject()")
	}()

	var err error
	err = nil

	// Generate test-text to see if data can be sent in to this object
	testText := app.H1().Text(mt.canBeAnyText)

	// Create DropDown for choosing which table date to show/add/edit
	tableTypeSelector, err := mt.CreateTableTypeSelectorObject()
	if err != nil {
		fmt.Println("Problem generating tableTypeSelector for table")
	}

	// Generate Alert messages
	alertMessages, err := mt.GenerateAlertMessages()
	if err != nil {
		fmt.Println("Problem generating PopUp")
	}

	// Generate PopUp and Button for opening (no good solution but I don't know how to open Bootstrap-popup in another way
	openPopUp, popUpp, err := mt.CreatePopUp()
	if err != nil {
		fmt.Println("Problem generating PopUp")
	}

	// Generate SearchInDB bar
	searchInput, err := mt.CreateSearchBar()
	if err != nil {
		fmt.Println("Problem generating SearchInDB Bar")
	}

	// Create table rows
	err = mt.UpdateRowNodes()
	if err != nil {
		fmt.Println("Problem generating rows for table")
	}

	// Dynamically create column headers for MagicTable
	err = mt.UpdateColumnsNodes()
	if err != nil {
		fmt.Println("Problem generating columns for table")
	}

	// Create the Edit-boxed to be used for doing New, Edit and Delete in table
	editRows, err := mt.CreateEditBoxes()
	if err != nil {
		fmt.Println("Problem generating Edit-boxes")
	}

	// Create buttons for New, Edit and Delete
	newEditDeleteButtons, err := mt.CreateButtonsForNewEditDelete()
	if err != nil {
		fmt.Println("Problem generating buttons for New, Edit and Delete")
	}

	// Create buttons for Save and Cancel
	saveCancelButtons, err := mt.CreateButtonsForSaveCancel()
	if err != nil {
		fmt.Println("Problem generating buttons for Save and Cancel")
	}

	// Build the table
	magicTableHeaderAndDataRows := app.Table().Class("table").Body(
		app.Tr().Body(mt.tableColumnNodes...),
		app.TBody().Body(mt.tableRowsNodes...),
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
				//app.If(mt.StateCheckToShowBaseButtons() == true,
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
