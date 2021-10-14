package main

import (
	//"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// GenerateAlertMessages generates the GUI objects for all alert messages
func (mt *MagicTable) GenerateAlertMessages() ([]app.UI, error) {
	var err error
	err = nil

	/*
		<div class="alert alert-warning alert-dismissible fade show" role="alert">
		<strong>Holy guacamole!</strong> You should check in on some of those fields below.
		<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
		</div>
	*/
	/*
		numberOfElementsInAlerts := len(mt.messagesToAlertToUser)
		for alertMessageIndex := numberOfElementsInAlerts - 1; alertMessageIndex >= 0; alertMessageIndex-- {
			//fmt.Println("Ska inte bort: ", mt.messagesToAlertToUser[alertMessageIndex], app.Window().GetElementByID(mt.messagesToAlertToUser[alertMessageIndex].id))
			if mt.messagesToAlertToUser[alertMessageIndex].processCount >= 2 {
				//fmt.Println("Vid borttag: ", mt.messagesToAlertToUser[alertMessageIndex], app.Window().GetElementByID(mt.messagesToAlertToUser[alertMessageIndex].id))
				if app.Window().GetElementByID(mt.messagesToAlertToUser[alertMessageIndex].id).IsNull() {
					fmt.Println("Vid borttag2: ", mt.messagesToAlertToUser[alertMessageIndex].id)
					mt.messagesToAlertToUser = mt.removeIndexFromMagicTable(mt.messagesToAlertToUser, alertMessageIndex)
					//alertElement := app.Window().GetElementByID(mt.messagesToAlertToUser[alertMessageIndex].id)
					//alertElement.Call("dispose")
				}
			}
		}


	*/
	//fmt.Println("Antal Alerts: " + strconv.Itoa(len(mt.messagesToAlertToUser)))
	alertMessages := []app.UI{}
	for alertMessageToUserIndex, alertMessageToUser := range mt.messagesToAlertToUser {

		//fmt.Println(mt.messagesToAlertToUser)

		mt.messagesToAlertToUser[alertMessageToUserIndex].processCount = mt.messagesToAlertToUser[alertMessageToUserIndex].processCount + 1
		showAlertClass := "show"
		//fmt.Println(alertMessageToUser.id, alertMessageToUser.show)
		if alertMessageToUser.show == true {
			showAlertClass = "show"
		} else {
			showAlertClass = "hide"
		}

		showAlertClass = "show"

		alertMessage := app.Div().
			Class("alert "+alertMessageToUser.alertType+" alert-dismissible fade "+showAlertClass).
			Aria("role", "alert").
			Body(
				app.Text(alertMessageToUser.alertMessage),
				app.Button().
					Type("button").
					Class("btn-close").
					DataSet("bs-dismiss", "alert").
					DataSet("aria-label", "Close").
					ID(alertMessageToUser.id+"button").
					//OnClick(mt.onCloseAlertWrapper(alertMessageToUser.id)).
					TabIndex(-1))

		if alertMessageToUser.show == true {
			alertMessages = append(alertMessages, alertMessage)
			//fmt.Println("alertMessageToUser.show == true", alertMessageToUser.id, app.Window().GetElementByID(alertMessageToUser.id).Type(), alertMessageToUser.id, app.Window().GetElementByID(alertMessageToUser.id).IsNull())

		} else {
			alertMessages = append(alertMessages, alertMessage)
			//fmt.Println("alertMessageToUser.show != true", alertMessageToUser.id, app.Window().GetElementByID(alertMessageToUser.id).Type(), app.Window().GetElementByID(alertMessageToUser.id).IsNull())
			//app.Window().GetElementByID(alertMessageToUser.id).
		}
		//elem := app.Window().
	}
	return alertMessages, err
}
