package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// Generates the GUI button objects for the New, Edit and Delete buttons
func (mt *MagicTable) CreateButtonsForNewEditDelete() ([]app.UI, error) {
	var err error
	err = nil

	// Generate New, Edit Delete buttons
	baseButtons := []app.UI{}
	buttonText, buttenDisabled := mt.isButtonDisabled(NewButton)
	newButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((mt.onButtonClickWrapper(NewButton)))

	baseButtons = append(baseButtons, newButton)

	buttonText, buttenDisabled = mt.isButtonDisabled(EditButton)
	editButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((mt.onButtonClickWrapper(EditButton)))
	baseButtons = append(baseButtons, editButton)

	buttonText, buttenDisabled = mt.isButtonDisabled(DeleteButton)
	deleteButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((mt.onButtonClickWrapper(DeleteButton)))
	baseButtons = append(baseButtons, deleteButton)

	return baseButtons, err
}

// Generates the GUI button objects for the Save and Cancel buttons
func (mt *MagicTable) CreateButtonsForSaveCancel() ([]app.UI, error) {
	var err error
	err = nil

	// Create Save & Cancel buttons
	SaveCancelButtons := []app.UI{}
	buttonText, buttenDisabled := mt.isButtonDisabled(SaveButton)
	saveButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((mt.onButtonClickWrapper(SaveButton)))
	SaveCancelButtons = append(SaveCancelButtons, saveButton)

	buttonText, buttenDisabled = mt.isButtonDisabled(CancelButton)
	cancelButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((mt.onButtonClickWrapper(CancelButton)))
	SaveCancelButtons = append(SaveCancelButtons, cancelButton)

	refreshButton := app.Button().
		Text("Refresh").
		Style("margin", "5px").
		Disabled(false).
		OnClick((mt.onRefreshButtonClickWrapper()))
	SaveCancelButtons = append(SaveCancelButtons, refreshButton)

	return SaveCancelButtons, err
}
