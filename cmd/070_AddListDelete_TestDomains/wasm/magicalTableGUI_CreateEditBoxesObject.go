package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// CreateEditBoxes Generates the GUI objects for the Edit boxes to be able to do New, Edit, Delete
func (mt *MagicTable) CreateEditBoxes() ([]app.UI, error) {
	var err error
	err = nil

	// Check if they should be enabled or not
	areNewUpdateDeleteTextBoxesEnabled := mt.areNewUpdateDeleteTextBoxesDisabled()

	// Dynamically create column headers for MagicTable
	editRows := []app.UI{}
	for _, columnMetadataResponse := range mt.testDataAndMetaData.magicTableMetaData {

		//rowLabel := app.Label().
		//	Text(columnMetadataResponse.GetColumnHeaderName())
		rowLabel := columnMetadataResponse.GetColumnHeaderName()

		columnDataName := columnMetadataResponse.GetColumnDataName()
		rowTextBoxValue := mt.GetRowTextBoxValueForEdit(columnDataName)

		elem := app.Window().GetElementByID(columnDataName)
		if elem.Type() == 6 {
			elem.Set("value", rowTextBoxValue)
		}
		/*
			rowTextBox := app.Input().
				Value(rowTextBoxValue).
				ID(columnDataName).
				OnDblClick(mt.OnTextboxDblClickappWrapper(columnDataName)).
				Disabled(areNewUpdateDeleteTextBoxesEnabled)
			fmt.Println("elem_xxxx_rowTextBoxValue:", rowTextBoxValue) //elem.Get("value")
		*/
		rowTextBox := app.Div().
			Class("form-floating mb-3").
			Body(app.Input().
				Type("text").
				Class("form-control").
				ID(columnDataName).
				Placeholder(columnDataName).
				Disabled(areNewUpdateDeleteTextBoxesEnabled).
				Value(rowTextBoxValue),

				app.Label().
					For(columnDataName).
					Text(rowLabel))

		editRow := app.Div().Body(
			//rowLabel,
			rowTextBox)

		// Check if column should be shown
		shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
		if shouldBeShown == true {
			editRows = append(editRows, editRow)
		}
	}

	return editRows, err
}
