package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"strconv"
)

func (mt *MagicTable) isButtonDisabled(buttonToEvaluate int) (buttonText string, buttonDisbled bool) {

	//fmt.Println("Button isButtonDisabled: " + strconv.Itoa(buttonToEvaluate))

	switch buttonToEvaluate {
	case NewButton:
		//fmt.Println("isButtonDisabled: 'NewButton'")
		buttonText = "New"

		switch mt.tableState {
		case TableState_List:
			buttonDisbled = false

		default:
			buttonDisbled = true
		}

	case EditButton:
		//fmt.Println("iisButtonDisabled: 'EditButton'")
		buttonText = "Edit"

		switch mt.tableState {
		case TableState_List:
			if mt.rowSelected > -1 {
				buttonDisbled = false
			} else {
				buttonDisbled = true
			}

		default:
			buttonDisbled = true
		}

	case DeleteButton:
		//fmt.Println("isButtonDisabled: 'DeleteButton'")
		buttonText = "Delete"

		switch mt.tableState {
		case TableState_List:
			if mt.rowSelected > -1 {
				buttonDisbled = false
			} else {
				buttonDisbled = true
			}

		default:
			buttonDisbled = true
		}

	case SaveButton:
		//fmt.Println("isButtonDisabled: 'SaveButton'")

		switch mt.tableState {
		case TableState_New:
			buttonText = "Save"
			buttonDisbled = false

		case TableState_Edit:
			buttonText = "Update"
			buttonDisbled = false

		case TableState_Delete:
			buttonText = "Delete"
			buttonDisbled = false

		default:
			buttonText = "Save/Update/Delete"
			buttonDisbled = true

		}

	case CancelButton:
		//fmt.Println("onButtonClickWrapper is called: 'CancelButton'")
		switch mt.tableState {
		case TableState_New,
			TableState_Edit,
			TableState_Delete:
			buttonText = "Cancel"
			buttonDisbled = false

		default:
			buttonText = "Cancel"
			buttonDisbled = true

		}

	default:
		fmt.Println("isButtonDisabled is called with unknown value: " + strconv.Itoa(buttonToEvaluate))
	}

	return buttonText, buttonDisbled
}

func (mt *MagicTable) onButtonClickWrapper(buttonThatWasClicked int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		mt.onButtonClick(buttonThatWasClicked)

	}
}
func (mt *MagicTable) onButtonClick(buttonThatWasClicked int) {

	//fmt.Println("Button that was clicked: " + strconv.Itoa(buttonThatWasClicked))

	switch buttonThatWasClicked {
	case NewButton:
		//fmt.Println("onButtonClickWrapper is called: 'NewButton'")
		mt.tableState = TableState_New

	case EditButton:
		//fmt.Println("onButtonClickWrapper is called: 'EditButton'")
		mt.tableState = TableState_Edit
		for _, columnMetadataResponse := range mt.testDataAndMetaData.magicTableMetaData {

			columnDataName := columnMetadataResponse.GetColumnDataName()
			//rowTextBoxValue := mt.GetRowTextBoxValueForEdit(columnDataName)
			elem := app.Window().
				GetElementByID(columnDataName)
			fmt.Println(columnDataName, elem.IsNull())

		}

	case DeleteButton:
		//fmt.Println("onButtonClickWrapper is called: 'DeleteButton'")
		mt.tableState = TableState_Delete

	case SaveButton:
		//fmt.Println("onButtonClickWrapper is called: 'SaveButton'")
		switch mt.tableState {
		case TableState_New:
			mt.tableState = TableState_New_Save
			//fmt.Println("Current State: 'TableState_New_Save'")

			mt.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		case TableState_Edit:
			mt.tableState = TableState_Edit_Save
			//fmt.Println("Current State: 'TableState_Edit_Save'")

			mt.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		case TableState_Delete:
			mt.tableState = TableState_Delete_Save
			//fmt.Println("Current State: 'TableState_Delete_Save'")
			app.Window().
				GetElementByID("openModalButton").
				Call("click")
			//app.Window().("modal")
			//GetElementByID("staticBackdrop").
			//Set("bs-toggle", "modal")
			//Call("modal")
			//mt.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		default:
			fmt.Println("Unknown state: " + strconv.Itoa(mt.tableState))

		}

	case CancelButton:
		//fmt.Println("onButtonClickWrapper is called: 'CancelButton'")
		switch mt.tableState {
		case TableState_New:
			mt.tableState = TableState_New_Cancel
			//fmt.Println("Current State: 'TableState_New_Cancel'")

			mt.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		case TableState_Edit:
			mt.tableState = TableState_Edit_Cancel
			//fmt.Println("Current State: 'TableState_Edit_Cancel'")

			mt.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		case TableState_Delete:
			mt.tableState = TableState_Delete_Cancel
			//fmt.Println("Current State: 'TableState_Delete_Cancel'")

			mt.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		default:
			fmt.Println("Unknown state: " + strconv.Itoa(mt.tableState))

		}

	default:
		fmt.Println("onButtonClickWrapper is called with unknown value: " + strconv.Itoa(buttonThatWasClicked))
	}

	mt.Update()
}
