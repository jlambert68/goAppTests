package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/sirupsen/logrus"
)

// Generates the GUI objects for the rows in the table
func (mt *MagicTable) CreateTableTypeSelectorObject() (app.HTMLSelect, error) {

	mt.logger.WithFields(logrus.Fields{
		"Id": "80998c75-b018-4a8c-b8e4-d5c8b0adc5da",
	}).Debug("Entering: CreateTableTypeSelectorObject()")

	var err error
	err = nil

	var tableTypeSelectorDisabled bool

	switch mt.tableState {
	case TableState_List:
		tableTypeSelectorDisabled = false

	default:
		tableTypeSelectorDisabled = true
	}

	// Generate list of options for DropDown
	tableTypeSelectorOptions := []app.UI{}
	for _, tableTypeSelectorOption := range mt.tableTypeSelectorOptionsInDB {
		tempOption := app.Option().
			Value(tableTypeSelectorOption.Guid).
			Text(tableTypeSelectorOption.TableName)
		tableTypeSelectorOptions = append(tableTypeSelectorOptions, tempOption)
	}

	tableTypeSelector := app.Select().
		Class("form-select").
		DataSet("aria-label", "Disabled select example").
		Disabled(tableTypeSelectorDisabled).Body(
		app.Option().
			Text("Open this select menu").
			Selected(true),
		app.Range(mt.tableTypeSelectorOptionsInDB).
			Slice(func(arrayCounter int) app.UI {
				return app.Option().
					Value(mt.tableTypeSelectorOptionsInDB[arrayCounter].Guid).
					Text(mt.tableTypeSelectorOptionsInDB[arrayCounter].TableName)
			})).OnChange(mt.MyOnChangeTableEditDropDownWrapper())

	mt.logger.WithFields(logrus.Fields{
		"Id": "a87e32ec-4d87-4f93-9b44-4f1fdcca5dfa",
	}).Debug("Exiting: CreateTableTypeSelectorObject()")

	return tableTypeSelector, err
}

// MyOnChangeTableEditDropDownWrapper is triggered when user change value in tableTypeSelector-dropdown
func (mt *MagicTable) MyOnChangeTableEditDropDownWrapper() app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		guid := ctx.JSSrc.Get("value").String()
		mt.tableTypeGuid = guid

		// Trigger reload correct data

		mt.reloadHeaderMetaData = true
		mt.Update()

	}
}
