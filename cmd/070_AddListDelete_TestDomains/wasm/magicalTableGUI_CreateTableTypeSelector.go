package main

import "github.com/maxence-charriere/go-app/v8/pkg/app"

// Generates the GUI objects for the rows in the table
func (p *MagicTable) CreateTableTypeSelectorObject() (app.HTMLSelect, error) {
	var err error
	err = nil

	var tableTypeSelectorDisabled bool

	switch p.tableState {
	case TableState_List:
		tableTypeSelectorDisabled = false

	default:
		tableTypeSelectorDisabled = true
	}

	// Generate list of options for DropDown
	tableTypeSelectorOptions := []app.UI{}
	for _, tableTypeSelectorOption := range p.tableTypeSelectorOptionsInDB {
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
		app.Range(p.tableTypeSelectorOptionsInDB).
			Slice(func(arrayCounter int) app.UI {
				return app.Option().
					Value(p.tableTypeSelectorOptionsInDB[arrayCounter].Guid).
					Text(p.tableTypeSelectorOptionsInDB[arrayCounter].TableName)
			})).OnChange(p.MyOnChangeTableEditDropDownWrapper())

	return tableTypeSelector, err
}

// MyOnChangeTableEditDropDownWrapper is triggered when user change value in tableTypeSelector-dropdown
func (p *MagicTable) MyOnChangeTableEditDropDownWrapper() app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		guid := ctx.JSSrc.Get("value").String()
		p.tableTypeGuid = guid

		// Trigger reload correct data

		p.reloadHeaderMetaData = true
		p.Update()

	}
}
