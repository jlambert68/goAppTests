package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// Generates the GUI button objects for the New, Edit and Delete buttons
func (p *MagicTable) CreateButtonsForNewEditDelete() ([]app.UI, error) {
	var err error
	err = nil

	// Generate New, Edit Delete buttons
	baseButtons := []app.UI{}
	buttonText, buttenDisabled := p.isButtonDisabled(NewButton)
	newButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((p.onButtonClickWrapper(NewButton)))

	baseButtons = append(baseButtons, newButton)

	buttonText, buttenDisabled = p.isButtonDisabled(EditButton)
	editButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((p.onButtonClickWrapper(EditButton)))
	baseButtons = append(baseButtons, editButton)

	buttonText, buttenDisabled = p.isButtonDisabled(DeleteButton)
	deleteButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((p.onButtonClickWrapper(DeleteButton)))
	baseButtons = append(baseButtons, deleteButton)

	return baseButtons, err
}

// Generates the GUI button objects for the Save and Cancel buttons
func (p *MagicTable) CreateButtonsForSaveCancel() ([]app.UI, error) {
	var err error
	err = nil

	// Create Save & Cancel buttons
	SaveCancelButtons := []app.UI{}
	buttonText, buttenDisabled := p.isButtonDisabled(SaveButton)
	saveButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((p.onButtonClickWrapper(SaveButton)))
	SaveCancelButtons = append(SaveCancelButtons, saveButton)

	buttonText, buttenDisabled = p.isButtonDisabled(CancelButton)
	cancelButton := app.Button().
		Type("button").
		Class("btn btn-secondary").
		Text(buttonText).
		Style("margin", "2px").
		Disabled(buttenDisabled).
		OnClick((p.onButtonClickWrapper(CancelButton)))
	SaveCancelButtons = append(SaveCancelButtons, cancelButton)

	refreshButton := app.Button().
		Text("Refresh").
		Style("margin", "5px").
		Disabled(false).
		OnClick((p.onRefreshButtonClickWrapper()))
	SaveCancelButtons = append(SaveCancelButtons, refreshButton)

	return SaveCancelButtons, err
}
