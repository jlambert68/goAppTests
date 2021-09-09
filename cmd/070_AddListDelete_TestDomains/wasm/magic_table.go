package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
	"sort"
	"strconv"
)

// Button Pressed
const (
	NewButton    = iota
	EditButton   = iota
	DeleteButton = iota
	SaveButton   = iota
	CancelButton = iota
)

// Statea for Table
const (
	TableState_Search_Load   = iota
	TableState_List          = iota
	TableState_New           = iota
	TableState_New_Save      = iota
	TableState_New_Cancel    = iota
	TableState_Edit          = iota
	TableState_Edit_Save     = iota
	TableState_Edit_Cancel   = iota
	TableState_Delete        = iota
	TableState_Delete_Save   = iota
	TableState_Delete_Cancel = iota
)

const (
	alertType_warning = "alert-warning"
	alertType_danger  = "alert-danger"
	alertType_success = "alert-success"
	alertType_info    = "alert-primary"
)

type magicTableRowStruct struct {
	cellElements []string
}

// Defines where data Table Metadata will be found
type (
	matgicTableDataType    []*api.Instance
	magicTableMetaDataType []*api.MagicTableColumnMetadata
)

type alertMessageStruct struct {
	id           string
	alertType    string
	alertMessage string
	processCount int
	show         bool
}

type MagicTable struct {
	app.Compo
	manager              *MagicManager
	reloadMeaderMetaData bool
	magicTableMetaData   magicTableMetaDataType //[]*api.MagicTableColumnMetadata
	instances            matgicTableDataType    // []*api.Instance
	//instancesSorted    []*api.Instance
	//magicTableRowsData []magicTableRowStruct
	tableType                  api.MagicTableType
	tableTypeGuid              string
	somethingWentWrong         bool
	searchString               string
	searchbarIsVisible         bool
	columnToSortOn             int
	columnSortOrderIsAscending bool
	canBeAnyText               string
	tableState                 int
	rowSelected                int
	uniqueRowSelected          int64
	messagesToAlertToUser      []alertMessageStruct
	alertId                    int
	// Field that reports whether an app update is available. False by default.
	updateAvailable              bool
	tableTypeSelectorOptionsInDB []*api.ListTableToEdit
}

func (p *MagicTable) SetManager(manager *MagicManager) {
	p.manager = manager
}

func (p *MagicTable) SetCanbeAnyText(canBeAnyText string) {
	p.canBeAnyText = canBeAnyText
}

func (p *MagicTable) updateHeaderMetaData() (magicTableMetaDataType, bool, error) {
	columnMeta, searchbarIsVisible, err := p.getColumnsMetadata()
	return columnMeta, searchbarIsVisible, err
}

// OnAppUpdate satisfies the app.Updater interface. It is called when the app is
// updated in background.
func (p *MagicTable) OnAppUpdate(ctx app.Context) {
	p.updateAvailable = ctx.AppUpdateAvailable // Reports that an app update is available.
	p.Update()                                 // Triggers UI update.
}

func (p *MagicTable) Render() app.UI {

	var err error
	var searchInput app.HTMLDiv

	if firstSearch == true {

		// No row is selected
		p.rowSelected = -1
		p.uniqueRowSelected = -1

		// Set Table State
		p.tableState = TableState_List

		p.UpdateInstances("")
		firstSearch = false

		// Try to update manager-ref
		magicManager.magicTable.SetManager(magicManager)

		app.Window().AddEventListener("hide.bs.modal", p.onCloseModalWrapper())
		//app.Window().AddEventListener("closed.bs.alert", p.onCloseAlertWrapper())

		p.alertId = 4

		messagesToAlertToUser := alertMessageStruct{
			id:           "alert0",
			alertType:    alertType_warning,
			alertMessage: "alert0: Detta är en varning!",
			processCount: 0,
			show:         true,
		}
		p.messagesToAlertToUser = append(p.messagesToAlertToUser, messagesToAlertToUser)

		messagesToAlertToUser = alertMessageStruct{
			id:           "alert1",
			alertType:    alertType_info,
			alertMessage: "alert1: Detta är ett infomeddelande.",
			processCount: 0,
			show:         true,
		}
		p.messagesToAlertToUser = append(p.messagesToAlertToUser, messagesToAlertToUser)

		messagesToAlertToUser = alertMessageStruct{
			id:           "alert2",
			alertType:    alertType_success,
			alertMessage: "alert2: Detta är ett meddelande då det gick bra.",
			processCount: 0,
			show:         true,
		}
		p.messagesToAlertToUser = append(p.messagesToAlertToUser, messagesToAlertToUser)

		messagesToAlertToUser = alertMessageStruct{
			id:           "alert3",
			alertType:    alertType_danger,
			alertMessage: "alert3: Detta är ett meddelande då det gick rikigt DÅLIGT!!!",
			processCount: 0,
			show:         true,
		}
		p.messagesToAlertToUser = append(p.messagesToAlertToUser, messagesToAlertToUser)

		// Call DB to get tableTypeSelectorOptions
		err = p.getTableTypeSelectorOptionsFromDB()
		if err != nil {
			fmt.Println(err)
		}

	} //else {

	//app.Window().AddEventListener("hide.bs.modal", p.onCloseModalWrapper())

	//}

	// Retrieve headers metadata, if not have been done
	if p.magicTableMetaData == nil || p.reloadMeaderMetaData == true {
		p.reloadMeaderMetaData = false
		_, p.searchbarIsVisible, err = p.updateHeaderMetaData()
		if err != nil {
			fmt.Println(" p.updateHeaderMetaData() generated Error:", err)

			// If something went wrong, return a message instead
			return app.H1().Text("Couldn't Retrieve Magic table Header Metadata")
		}
	}

	// Generate test-text to see if data can be sent in to this object
	testText := app.H1().Text(p.canBeAnyText)

	// Create DropDown for choosing which table date to show/add/edit
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
	/*
		app.Option().
			Value("2").
			Text("Two"),
		app.Option().
			Value("3").
			Text("Three"))


	*/

	/*
		<select class="form-select" aria-label="Disabled select example" disabled>
		<option selected>Open this select menu</option>
		<option value="1">One</option>
		<option value="2">Two</option>
		<option value="3">Three</option>
		</select>


	*/

	// Create base buttons
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

	// Create Save & Cancel buttons
	SaveCancelButtons := []app.UI{}
	buttonText, buttenDisabled = p.isButtonDisabled(SaveButton)
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

	// Generate New, Edit Delete
	// Check if they should be enabled or not
	areNewUpdateDeleteTextBoxesEnabled := p.areNewUpdateDeleteTextBoxesDisabled()
	// Dynamically create column headers for MagicTable
	editRows := []app.UI{}
	for _, columnMetadataResponse := range p.magicTableMetaData {

		//rowLabel := app.Label().
		//	Text(columnMetadataResponse.GetColumnHeaderName())
		rowLabel := columnMetadataResponse.GetColumnHeaderName()

		columnDataName := columnMetadataResponse.GetColumnDataName()
		rowTextBoxValue := p.GetRowTextBoxValueForEdit(columnDataName)

		elem := app.Window().GetElementByID(columnDataName)
		if elem.Type() == 6 {
			elem.Set("value", rowTextBoxValue)
		}
		/*
			rowTextBox := app.Input().
				Value(rowTextBoxValue).
				ID(columnDataName).
				OnDblClick(p.OnTextboxDblClickappWrapper(columnDataName)).
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

	/*
		<div class="alert alert-warning alert-dismissible fade show" role="alert">
		<strong>Holy guacamole!</strong> You should check in on some of those fields below.
		<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
		</div>
	*/
	/*
		numberOfElementsInAlerts := len(p.messagesToAlertToUser)
		for alertMessageIndex := numberOfElementsInAlerts - 1; alertMessageIndex >= 0; alertMessageIndex-- {
			//fmt.Println("Ska inte bort: ", p.messagesToAlertToUser[alertMessageIndex], app.Window().GetElementByID(p.messagesToAlertToUser[alertMessageIndex].id))
			if p.messagesToAlertToUser[alertMessageIndex].processCount >= 2 {
				//fmt.Println("Vid borttag: ", p.messagesToAlertToUser[alertMessageIndex], app.Window().GetElementByID(p.messagesToAlertToUser[alertMessageIndex].id))
				if app.Window().GetElementByID(p.messagesToAlertToUser[alertMessageIndex].id).IsNull() {
					fmt.Println("Vid borttag2: ", p.messagesToAlertToUser[alertMessageIndex].id)
					p.messagesToAlertToUser = p.removeIndexfromMagicTable(p.messagesToAlertToUser, alertMessageIndex)
					//alertElement := app.Window().GetElementByID(p.messagesToAlertToUser[alertMessageIndex].id)
					//alertElement.Call("dispose")
				}
			}
		}


	*/
	//fmt.Println("Antal Alerts: " + strconv.Itoa(len(p.messagesToAlertToUser)))
	alertMessages := []app.UI{}
	for alertMessageToUserIndex, alertMessageToUser := range p.messagesToAlertToUser {

		//fmt.Println(p.messagesToAlertToUser)

		p.messagesToAlertToUser[alertMessageToUserIndex].processCount = p.messagesToAlertToUser[alertMessageToUserIndex].processCount + 1
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
					//OnClick(p.onCloseAlertWrapper(alertMessageToUser.id)).
					TabIndex(-1))

		if alertMessageToUser.show == true {
			alertMessages = append(alertMessages, alertMessage)
			//fmt.Println("alertMessageToUser.show == true", alertMessageToUser.id, app.Window().GetElementByID(alertMessageToUser.id).Type(), alertMessageToUser.id, app.Window().GetElementByID(alertMessageToUser.id).IsNull())

		} else {
			alertMessages = append(alertMessages, alertMessage)
			fmt.Println("alertMessageToUser.show != true", alertMessageToUser.id, app.Window().GetElementByID(alertMessageToUser.id).Type(), app.Window().GetElementByID(alertMessageToUser.id).IsNull())
			//app.Window().GetElementByID(alertMessageToUser.id).
		}
		//elem := app.Window().
	}

	//p.messagesToAlertToUser = nil
	//fmt.Println("Antal Alerts2: " + strconv.Itoa(len(p.messagesToAlertToUser)))

	buttonText, _ = p.isButtonDisabled(SaveButton)

	openPopUp := app.Button().
		Type("button").
		Class("btn btn-primary").
		DataSet("bs-toggle", "modal").
		DataSet("bs-target", "#staticBackdrop").
		ID("openModalButton").
		Hidden(true)

	popUpp := app.Div().
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

	// Generate Search bar
	if p.searchbarIsVisible == true {
		input := app.Input().
			Class("form-control").
			Value(p.searchString).
			Placeholder("t2.small").
			AutoFocus(true).
			OnKeyup(p.MyOnInputChange)

		searchInput = app.Div().
			Class("input-group").
			Body(
				app.Div().
					Class("input-group-prepend").
					Body(app.Span().
						Class("input-group-text").
						Body(app.Text("🔍"))),
				input,
			)
	}

	// Create table rows
	nodes := []app.UI{}
	//Loop every row
	for rowCounter, i := range p.instances {
		rowNodes := []app.UI{}
		// Loop every column and get value and check if column should be shown
		for _, columnMetadataResponse := range p.magicTableMetaData {
			columnName := columnMetadataResponse.GetColumnDataName()
			columnData := p.formatColumnData(i, columnName)
			shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
			if shouldBeShown == true {
				rowNodes = append(rowNodes, app.Td().Body(app.Text(columnData)))
			}
		}

		// Set background color on selected row
		var rowSelectedColor string
		if i.UniqueId != p.uniqueRowSelected {
			rowSelectedColor = "#FFFFFF"
		} else {
			rowSelectedColor = "#E1EAFA"
			p.rowSelected = rowCounter
		}

		// Append values for each row
		nodes = append(nodes,
			app.Tr().
				Body(rowNodes...).
				OnClick(p.MyOnClickOnRowWrapper(rowCounter)).
				OnDblClick(p.MyOnDblClickOnRowWrapper(rowCounter)).
				Style("background-color", rowSelectedColor))

		/*
			nodes = append(nodes, app.Tr().Body(
				app.Td().Body(app.Text(i.Name)),
				app.Td().Body(app.Text(i.InstanceType)),
				app.Td().Body(app.Text(fmt.Sprintf("%v", i.Ecu))),
				app.Td().Body(app.Text(fmt.Sprintf("%v", i.Memory))),
				app.Td().Body(app.Text(i.Network)),
				app.Td().Body(app.Text(i.Price)),
			).OnDblClick(p.MyOnDblClickOnRowWrapper(rowCounter)))
		*/
	}
	/*
		nodes2 := []app.UI{}
		var rowUIElement app.HTMLTr

		for _, tableRowData := range p.magicTableRowsData{
			tableCells := []app.UI{}
			for _, tableCellValue := range tableRowData.cellElements {
				tableCells = append(tableCells, app.Td().Body(app.Text(tableCellValue)))
			}
			rowUIElement = app.Tr().Body(tableCells...)
			nodes2 = append(nodes2, rowUIElement)
		}
	*/

	// Dynamically create column headers for MagicTable
	columns := []app.UI{}
	for columnCounter, columnMetadataResponse := range p.magicTableMetaData {
		// Check if column should be shown
		columnHeader := app.Th().
			Scope("col").
			Body(app.Text(columnMetadataResponse.GetColumnHeaderName() + p.SetSortIconForTableHeader(columnCounter))).
			OnClick(p.MyOnColumnClickWrapper(columnCounter))
		shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
		if shouldBeShown == true {
			columns = append(columns, columnHeader)
		}
	}

	// Build the table
	magicTableHeaderAndDataRows := app.Table().Class("table").Body(
		app.Tr().Body(columns...),
		app.TBody().Body(nodes...),
		//app.TBody().Body(nodes2...),
	)

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
					Body(baseButtons...),
				//).Else(
				app.Div().
					Class("row").
					Body(editRows...),
				app.Div().
					//Style("color", "#ffff00").
					//Style("border", "solid").
					Style("padding", "10px").
					Body(SaveCancelButtons...)),
		//	),
		searchInput,
		app.Div().
			Class("container-fluid").
			Body(
				magicTableHeaderAndDataRows,
			),
	)

	return magicTableRenderedObject

}

func (p *MagicTable) removeIndexfromMagicTable(s []alertMessageStruct, index int) []alertMessageStruct {
	return append(s[:index], s[index+1:]...)
}

func (p *MagicTable) GetRowTextBoxValueForEdit(columnDataName string) string {

	var returnValue string

	switch p.tableState {

	case TableState_New:
		returnValue = ""

	case TableState_Edit,
		TableState_Delete:
		rowData := p.instances[p.rowSelected]
		//returnValue =rowData.GetName()

		r := reflect.ValueOf(rowData)
		f := reflect.Indirect(r).FieldByName(columnDataName)

		returnValue = fmt.Sprintf("%v", f)

	default:
		returnValue = ""
	}

	return returnValue
}

func (p *MagicTable) StateCheckToShowBaseButtons() bool {
	switch p.tableState {
	case
		TableState_Search_Load,
		TableState_List:
		return true
	}
	return false
}

func (p *MagicTable) SetSortIconForTableHeader(columnCounter int) string {
	var returnMessage string
	//fmt.Println("p.columnToSortOn: ", p.columnToSortOn, p.magicTableMetaData[columnCounter].String(), p.magicTableMetaData[columnCounter].Sortable)

	// Only Set icon on correct Header

	if p.magicTableMetaData[columnCounter].Sortable == true {
		// Column is sortable
		if columnCounter == p.columnToSortOn && firstTimeForSorting == false {
			if p.columnSortOrderIsAscending {
				returnMessage = "↑"
			} else {
				returnMessage = "↓"
			}
		} else {

			returnMessage = "↑↓"
		}
	} else {
		// Column is not sortable
		returnMessage = ""
	}

	return returnMessage
}

func (p *MagicTable) formatColumnData(v *api.Instance, field string) string {

	var returnValue string

	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	returnValue = fmt.Sprintf("%v", f)

	return returnValue
}

//func (p *InstanceTable)  onClick(ctx app.Context, e app.Event) {
//	fmt.Println("onClick is called")
//}
func (p *MagicTable) MyOnDblClickOnRowWrapper(rowThatWasDoubleClickedOn int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		p.MyOnDblClickOnRow(rowThatWasDoubleClickedOn)

		/*
			// Only allow click if Table state is in 'List'
			if  p.tableState != TableState_List {
				return
			}


			fmt.Println("OnDblClick is called::::::" + strconv.Itoa(rowThatWasDoubleClickedOn))
			rowData := p.instances[rowThatWasDoubleClickedOn]
			fmt.Println("RowData: " + rowData.String())
			fmt.Println("")


			// Trigger Edit
			p.onButtonClickWrapper(EditButton)
		*/

	}
}

func (p *MagicTable) MyOnChangeTableEditDropDownWrapper() app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		guid := ctx.JSSrc.Get("value").String()
		p.tableTypeGuid = guid

		p.reloadMeaderMetaData = true
		p.Update()

	}
}

func (p *MagicTable) MyOnDblClickOnRow(rowThatWasDoubleClickedOn int) {
	// Only allow click if Table state is in 'List'
	if p.tableState != TableState_List {
		return
	}

	fmt.Println("OnDblClick is called::::::" + strconv.Itoa(rowThatWasDoubleClickedOn))
	rowData := p.instances[rowThatWasDoubleClickedOn]
	fmt.Println("RowData: " + rowData.String())
	fmt.Println("")

	// Trigger Edit
	p.onButtonClick(EditButton)
}

func (p *MagicTable) MyOnClickOnRowWrapper(rowThatWasClickedOn int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		// Only allow click if Table state is in 'List'
		if p.tableState != TableState_List {
			return
		}

		fmt.Println("OnClick is called::::::" + strconv.Itoa(rowThatWasClickedOn))
		rowData := p.instances[rowThatWasClickedOn]
		fmt.Println("RowData: " + rowData.String())
		fmt.Println("")

		p.rowSelected = rowThatWasClickedOn
		p.uniqueRowSelected = rowData.GetUniqueId()
		p.Update()
	}
}

func (p *MagicTable) OnDblClickapp(src app.Value, e app.Event) {

	// Only allow click if Table state is in 'List'
	if p.tableState != TableState_List {
		return
	}
	fmt.Println("OnDblClick is called - 1025")
	fmt.Println("src - " + src.JSValue().String())
	fmt.Println("e - " + e.Value.String())

}

func (p *MagicTable) OnTextboxDblClickappWrapper(buttonId string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		elem := app.Window().GetElementByID(buttonId)
		fmt.Println("elem - ", buttonId, elem.Get("value").String())
		//elem.Set("value", "APA")

	}
}

func (p *MagicTable) MyOnColumnClickWrapper(columnNumberThatWasClicked int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		// Only allow click if Table state is in 'List'
		if p.tableState != TableState_List {
			return
		}

		fmt.Println("MyOnColumnClickWrapper is called:::::: " + strconv.Itoa(columnNumberThatWasClicked))
		fmt.Println("")

		// Only sort if column is sortable
		if p.magicTableMetaData[columnNumberThatWasClicked].Sortable == true {

			// Check if table for sorted before
			if firstTimeForSorting == true {
				firstTimeForSorting = false

				// First time
				p.columnToSortOn = columnNumberThatWasClicked
				p.columnSortOrderIsAscending = true

				// Not the first time
			} else {
				// If the same column was clicked then invert sort order for column
				if p.columnToSortOn == columnNumberThatWasClicked {
					p.columnSortOrderIsAscending = !p.columnSortOrderIsAscending
				} else {
					// New column to sort on
					p.columnToSortOn = columnNumberThatWasClicked
					p.columnSortOrderIsAscending = true
				}
			}

			var compareResult bool
			// Sort by age, keeping original order or equal elements.
			sort.SliceStable(p.instances, func(i, j int) bool {

				// General solution - Nice
				if true {
					fieldsToExtract := p.magicTableMetaData[columnNumberThatWasClicked].GetColumnDataName()
					//fmt.Println("fieldsToExtract", fieldsToExtract)
					value1 := getAttr(p.instances[i], fieldsToExtract)
					value2 := getAttr(p.instances[j], fieldsToExtract)

					switch p.magicTableMetaData[columnNumberThatWasClicked].ColumnDataType {
					case api.MagicTableColumnDataType_String:
						compareResult = value1.String() < value2.String()
					case api.MagicTableColumnDataType_Float:
						compareResult = value1.Float() < value2.Float()
					}

					return xnor(p.columnSortOrderIsAscending, compareResult)

				} else {
					// Hard defined sorting-types - Should be removed
					switch p.columnToSortOn {
					case 0:
						return xnor(p.columnSortOrderIsAscending, p.instances[i].Name < p.instances[j].Name)

					case 1:
						return xnor(p.columnSortOrderIsAscending, p.instances[i].InstanceType < p.instances[j].InstanceType)

					case 2:
						return xnor(p.columnSortOrderIsAscending, p.instances[i].Ecu < p.instances[j].Ecu)

					case 3:
						return xnor(p.columnSortOrderIsAscending, p.instances[i].Memory < p.instances[j].Memory)

					case 4:
						return xnor(p.columnSortOrderIsAscending, p.instances[i].Network < p.instances[j].Network)

					case 5:
						return xnor(p.columnSortOrderIsAscending, p.instances[i].Price < p.instances[j].Price)

					default:
						fmt.Println("Could not sort on column: " + strconv.Itoa(columnNumberThatWasClicked))
						fmt.Println("Will sort on Name-ascending, instead: ")
						return xnor(true, p.instances[i].Name < p.instances[j].Name)
					}
				}

			})
			p.Update()
			//magicManager.Update()

			// Update selected row
			for rowCounter, i := range p.instances {

				// Check if slected row
				if i.UniqueId == p.uniqueRowSelected {
					p.rowSelected = rowCounter
				}
			}

		}
	}
}

func (p *MagicTable) areNewUpdateDeleteTextBoxesDisabled() (textBoxesAreEnabled bool) {

	switch p.tableState {
	case TableState_New,
		TableState_Edit:
		textBoxesAreEnabled = false

	default:
		textBoxesAreEnabled = true

	}

	return textBoxesAreEnabled

}

func (p *MagicTable) isButtonDisabled(buttonToEvaluate int) (buttonText string, buttonDisbled bool) {

	//fmt.Println("Button isButtonDisabled: " + strconv.Itoa(buttonToEvaluate))

	switch buttonToEvaluate {
	case NewButton:
		//fmt.Println("isButtonDisabled: 'NewButton'")
		buttonText = "New"

		switch p.tableState {
		case TableState_List:
			buttonDisbled = false

		default:
			buttonDisbled = true
		}

	case EditButton:
		//fmt.Println("iisButtonDisabled: 'EditButton'")
		buttonText = "Edit"

		switch p.tableState {
		case TableState_List:
			if p.rowSelected > -1 {
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

		switch p.tableState {
		case TableState_List:
			if p.rowSelected > -1 {
				buttonDisbled = false
			} else {
				buttonDisbled = true
			}

		default:
			buttonDisbled = true
		}

	case SaveButton:
		//fmt.Println("isButtonDisabled: 'SaveButton'")

		switch p.tableState {
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
		switch p.tableState {
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

func (p *MagicTable) onRefreshButtonClickWrapper() app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		/*
			alertIdStr := "alert" + strconv.Itoa(p.alertId)
			p.alertId = p.alertId + 1

			messagesToAlertToUser := alertMessageStruct{
				id:           alertIdStr,
				alertType:    alertType_danger,
				alertMessage: alertIdStr + ": Detta är ett NYTT meddelande då det gick rikigt DÅLIGT!!!",
				processCount: 0,
				show:         true,
			}
			p.messagesToAlertToUser = append(p.messagesToAlertToUser, messagesToAlertToUser)


		*/
		//for _, a := range p.messagesToAlertToUser {
		//	fmt.Println(a.id, app.Window().GetElementByID(a.id), app.Window().GetElementByID(a.id).IsNull())
		//}

		p.Update()
		/*
			ctx.Async(func() {
				time.Sleep(2 * time.Second)
				app.Window().GetElementByID("alert0button").Call("click")
			})


		*/
	}
}
func (p *MagicTable) onCloseAlertWrapper(alertId string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		fmt.Println("***Alert-box (" + alertId + ") stängdes av användaren***")
		//app.Window().Call("hideAlert", alertId)
		for alertCounter, _ := range p.messagesToAlertToUser {
			if p.messagesToAlertToUser[alertCounter].id == alertId {
				fmt.Println("hittade")
				p.messagesToAlertToUser[alertCounter].show = false
				break
			}
		}

		//app.Window().GetElementByID("alert3").Call("focus")
		//p.Update()
	}
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

func (p *MagicTable) onButtonClickWrapper(buttonThatWasClicked int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		p.onButtonClick(buttonThatWasClicked)

	}
}
func (p *MagicTable) onButtonClick(buttonThatWasClicked int) {

	//fmt.Println("Button that was clicked: " + strconv.Itoa(buttonThatWasClicked))

	switch buttonThatWasClicked {
	case NewButton:
		//fmt.Println("onButtonClickWrapper is called: 'NewButton'")
		p.tableState = TableState_New

	case EditButton:
		//fmt.Println("onButtonClickWrapper is called: 'EditButton'")
		p.tableState = TableState_Edit
		for _, columnMetadataResponse := range p.magicTableMetaData {

			columnDataName := columnMetadataResponse.GetColumnDataName()
			//rowTextBoxValue := p.GetRowTextBoxValueForEdit(columnDataName)
			elem := app.Window().
				GetElementByID(columnDataName)
			fmt.Println(columnDataName, elem.IsNull())

		}

	case DeleteButton:
		//fmt.Println("onButtonClickWrapper is called: 'DeleteButton'")
		p.tableState = TableState_Delete

	case SaveButton:
		//fmt.Println("onButtonClickWrapper is called: 'SaveButton'")
		switch p.tableState {
		case TableState_New:
			p.tableState = TableState_New_Save
			//fmt.Println("Current State: 'TableState_New_Save'")

			p.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		case TableState_Edit:
			p.tableState = TableState_Edit_Save
			//fmt.Println("Current State: 'TableState_Edit_Save'")

			p.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		case TableState_Delete:
			p.tableState = TableState_Delete_Save
			//fmt.Println("Current State: 'TableState_Delete_Save'")
			app.Window().
				GetElementByID("openModalButton").
				Call("click")
			//app.Window().("modal")
			//GetElementByID("staticBackdrop").
			//Set("bs-toggle", "modal")
			//Call("modal")
			//p.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		default:
			fmt.Println("Unknown state: " + strconv.Itoa(p.tableState))

		}

	case CancelButton:
		//fmt.Println("onButtonClickWrapper is called: 'CancelButton'")
		switch p.tableState {
		case TableState_New:
			p.tableState = TableState_New_Cancel
			//fmt.Println("Current State: 'TableState_New_Cancel'")

			p.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		case TableState_Edit:
			p.tableState = TableState_Edit_Cancel
			//fmt.Println("Current State: 'TableState_Edit_Cancel'")

			p.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		case TableState_Delete:
			p.tableState = TableState_Delete_Cancel
			//fmt.Println("Current State: 'TableState_Delete_Cancel'")

			p.tableState = TableState_List
			//fmt.Println("Current State: 'TableState_List'")

		default:
			fmt.Println("Unknown state: " + strconv.Itoa(p.tableState))

		}

	default:
		fmt.Println("onButtonClickWrapper is called with unknown value: " + strconv.Itoa(buttonThatWasClicked))
	}

	p.Update()
}

//https://play.golang.org/p/pxj_gCnvF1J
func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}

// XNOR function
func xnor(a, b bool) bool {
	return !((a || b) && (!a || !b))
}

/*
func (h *MagicTable) sortDataOnColumn(columnNumberToSortOn int, sortAscending bool)  () {

	type rowValue struct {
		row int
		value string
	}
	var rowValueToAdd rowValue

	rowValuArray := []rowValue{}

	for rowCounter, i := range h.instances {
		rowValueToAdd = rowValue{
			row:  rowCounter,
			value: i.Price,
		}

		rowValuArray = append(rowValuArray, rowValueToAdd)

	}
	sort.Slice(rowValuArray, func(i, j int) bool {
		return false
	})
}


*/

func (h *MagicTable) getColumnsMetadata() (magicTableMetaDataType, bool, error) {

	magicTableMetadataRequest := api.MagicTableMetadataRequest{
		MagicTableMetadataType: h.tableType,
		TableTypeGuid:          h.tableTypeGuid,
	}
	magicTableMetadataRespons, err := api.CallApiGetMagicTableMetadata(magicTableMetadataRequest)

	if err != nil {
		fmt.Println("api.CallApiGetMagicTableMetadata Error:", err)
		return nil, false, err
	} else {
		h.magicTableMetaData = magicTableMetadataRespons.GetMagicTableColumnsMetadata()
	}

	return magicTableMetadataRespons.GetMagicTableColumnsMetadata(), magicTableMetadataRespons.SearchbarIsVisible, err
}

func (p *MagicTable) MyOnInputChange(ctx app.Context, e app.Event) {
	p.searchString = ctx.JSSrc.Get("value").String()
	//p.searchString = src.Get("value").String()
	//p.Update()
	p.UpdateInstances(p.searchString)
}

func (h *MagicTable) Search(q string) matgicTableDataType {
	instances, err := api.CallApiSearch(api.SearchRequest{
		Query: q,
	})

	if err != nil {
		fmt.Println("Search Error:", err)
		return matgicTableDataType{}
	}

	return instances.Instances
}

func (h *MagicTable) UpdateInstances(q string) {

	// Unselect rows
	h.rowSelected = -1

	instances := h.Search(q)
	h.instances = instances

	//h.convertInstancesIntoStandardMagicTable()
	//fmt.Println(h.magicTableRowsData)

	h.Update()

}

// Get all tables that can be edited
func (h *MagicTable) getTableTypeSelectorOptionsFromDB() error {
	instances, err := api.CallApiListTablesToEdit(api.EmptyParameter{})

	if err != nil {
		fmt.Println("Search Error:", err)
		return err
	}

	h.tableTypeSelectorOptionsInDB = instances.MyListTableToEdit

	return nil
}

/*
func (p *MagicTable) convertInstancesIntoStandardMagicTable() {

	// Create table rows
	magicTableRowsData := []magicTableRowStruct{}

	for _, tableRow := range p.instances {
		var magicTableRow magicTableRowStruct

		magicTableRow.cellElements = append(magicTableRow.cellElements, tableRow.Name)
		magicTableRow.cellElements = append(magicTableRow.cellElements, tableRow.InstanceType)
		magicTableRow.cellElements = append(magicTableRow.cellElements, fmt.Sprintf("%v", tableRow.Ecu))
		magicTableRow.cellElements = append(magicTableRow.cellElements, fmt.Sprintf("%v", tableRow.Memory))
		magicTableRow.cellElements = append(magicTableRow.cellElements, tableRow.Network)
		magicTableRow.cellElements = append(magicTableRow.cellElements, tableRow.Price)

		magicTableRowsData = append(magicTableRowsData, magicTableRow)

	}
	p.magicTableRowsData = magicTableRowsData
}

*/
