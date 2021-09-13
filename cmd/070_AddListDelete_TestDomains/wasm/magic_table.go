package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
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
	magicTableMetaDataType        []*api.MagicTableColumnMetadata
	originalTestdataInstancesType []*api.Instance
	testDomainsType               []*api.TestDomainForListingMessage
	testInstructionsType          []*api.ListTestInstructionRespons
)

// Defines Column Metadata & data-structires
type metadataAndDataStructuresStruct struct {
	// Metadata for columns
	magicTableMetaData magicTableMetaDataType

	// Original testdata (for testing only)
	originalTestdataInstances originalTestdataInstancesType

	// TestDomains
	testDomains testDomainsType

	// TestInstructions
	testInstructions testInstructionsType
}

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
	reloadHeaderMetaData bool
	//magicTableMetaData   magicTableMetaDataType //[]*api.MagicTableColumnMetadata
	//instances            magicTableOriginalTestDataType    // []*api.Instance
	testDataAndMetaData metadataAndDataStructuresStruct

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

	// Holds the rows in the table as UI-objects
	tableRowsNodes []app.UI

	// Holds the columns in the table as UI-objects
	tableColumnNodes []app.UI
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
	//var searchInput app.HTMLDiv

	if firstSearch == true {

		// No row is selected
		p.rowSelected = -1
		p.uniqueRowSelected = -1

		// Set Table State
		p.tableState = TableState_List

		p.RetrieveTableDataFromDB("")
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
	if p.testDataAndMetaData.magicTableMetaData == nil || p.reloadHeaderMetaData == true {
		p.reloadHeaderMetaData = false
		_, p.searchbarIsVisible, err = p.updateHeaderMetaData()
		if err != nil {
			fmt.Println(" p.updateHeaderMetaData() generated Error:", err)

			// If something went wrong, return a message instead
			return app.H1().Text("Couldn't Retrieve Magic table Header Metadata")
		}
	}

	//p.messagesToAlertToUser = nil
	//fmt.Println("Antal Alerts2: " + strconv.Itoa(len(p.messagesToAlertToUser)))

	magicTableRenderedObject, err := p.AssembleMagicalTableObject()
	if err != nil {
		fmt.Println("Problem assemble the magicTable page object")
	}

	return magicTableRenderedObject

}

func (p *MagicTable) removeIndexFromMagicTable(s []alertMessageStruct, index int) []alertMessageStruct {
	return append(s[:index], s[index+1:]...)
}

func (p *MagicTable) GetRowTextBoxValueForEdit(columnDataName string) string {

	var returnValue string

	switch p.tableState {

	case TableState_New:
		returnValue = ""

	case TableState_Edit,
		TableState_Delete:
		rowData := p.testDataAndMetaData.originalTestdataInstances[p.rowSelected]
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
	//fmt.Println("p.columnToSortOn: ", p.columnToSortOn, p.testDataAndMetaData.magicTableMetaData[columnCounter].String(), p.testDataAndMetaData.magicTableMetaData[columnCounter].Sortable)

	// Only Set icon on correct Header

	if p.testDataAndMetaData.magicTableMetaData[columnCounter].Sortable == true {
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
			rowData := p.testDataAndMetaData.originalTestdataInstances[rowThatWasDoubleClickedOn]
			fmt.Println("RowData: " + rowData.String())
			fmt.Println("")


			// Trigger Edit
			p.onButtonClickWrapper(EditButton)
		*/

	}
}

func (p *MagicTable) MyOnDblClickOnRow(rowThatWasDoubleClickedOn int) {
	// Only allow click if Table state is in 'List'
	if p.tableState != TableState_List {
		return
	}

	fmt.Println("OnDblClick is called::::::" + strconv.Itoa(rowThatWasDoubleClickedOn))
	rowData := p.testDataAndMetaData.originalTestdataInstances[rowThatWasDoubleClickedOn]
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
		rowData := p.testDataAndMetaData.originalTestdataInstances[rowThatWasClickedOn]
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
		h.testDataAndMetaData.magicTableMetaData = magicTableMetadataRespons.GetMagicTableColumnsMetadata()
	}

	return magicTableMetadataRespons.GetMagicTableColumnsMetadata(), magicTableMetadataRespons.SearchbarIsVisible, err
}

func (p *MagicTable) MyOnInputChange(ctx app.Context, e app.Event) {
	p.searchString = ctx.JSSrc.Get("value").String()
	//p.searchString = src.Get("value").String()
	//p.Update()
	p.RetrieveTableDataFromDB(p.searchString)
}

// Get all tables that can be edited
func (h *MagicTable) getTableTypeSelectorOptionsFromDB() error {
	instances, err := api.CallApiListTablesToEdit(api.EmptyParameter{})

	if err != nil {
		fmt.Println("SearchInDB Error:", err)
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
