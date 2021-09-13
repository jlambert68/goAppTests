package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// Generates the GUI objects for the columns in the table
func (p *MagicTable) CreatePopUp() (app.HTMLButton, app.HTMLDiv, error) {
	var err error
	err = nil

	buttonText, _ := p.isButtonDisabled(SaveButton)

	openPopUp := app.Button().
		Type("button").
		Class("btn btn-primary").
		DataSet("bs-toggle", "modal").
		DataSet("bs-target", "#staticBackdrop").
		ID("openModalButton").
		Hidden(true)

	popUp := app.Div().
		Class("modal fade").
		ID("staticBackdrop").
		DataSet("bs-backdrop", "static").
		DataSet("bs-keyboard", "false").
		DataSet("tabindex", "-1").
		DataSet("aria-labelledby", "staticBackdropLabel").
		DataSet("aria-hidden", "false").
		Body(
			app.Div().
				Class("modal-dialog").
				Body(
					app.Div().
						Class("modal-content").
						Body(
							app.Div().
								Class("modal-header").
								Body(
									app.H5().
										Class("modal-title").
										ID("staticBackdropLabel").
										Text("Modal title"),
									app.Button().
										Type("button").
										Class("btn-close").
										DataSet("bs-dismiss", "modal").
										DataSet("aria-label", "Close")),
							app.Div().
								Class("modal-body").
								Text("Detta är texten som kommer att visas!!!"),
							app.Div().
								Class("modal-footer").
								Body(
									app.Button().
										Type("button").
										Class("btn btn-secondary").
										DataSet("bs-dismiss", "modal").
										Text("Close").
										ID("ModalCancel"),

									app.Button().Type("button").
										Class("btn btn-primary").
										Text(buttonText).
										OnClick(p.onModalOKClicked())))))

	return openPopUp, popUp, err
}

func (p *MagicTable) onCloseModalWrapper() app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		p.changeStateToList()
	}
}

func (p *MagicTable) changeStateToList() {
	p.tableState = TableState_List
	fmt.Println("Close Modal triggered this")
	p.Update()
}

func (p *MagicTable) onModalOKClicked() app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		switch p.tableState {

		case TableState_Delete_Save:
			fmt.Println("Send 'Delete' to DB")
			p.rowSelected = -1
			p.uniqueRowSelected = -1

		default:
			fmt.Println("Wrong TableState, shouldn't be here:", p.tableState)

		}
		app.Window().
			GetElementByID("ModalCancel").
			Call("click")
		p.tableState = TableState_List
		p.Update()

	}
}

func (p *MagicTable) onModalCancelClicked() app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		switch p.tableState {

		case TableState_Delete_Save:
			fmt.Println("Cancelling Send to DB")

		default:
			fmt.Println("Wrong TableState, shouldn't be here:", p.tableState)

		}
		app.Window().
			GetElementByID("ModalCancel").
			Call("click")
		p.tableState = TableState_List
		p.Update()

	}
}
