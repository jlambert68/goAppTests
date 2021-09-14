package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/sirupsen/logrus"
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
	logger *logrus.Logger
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
	rowSelected                int64
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

func (mt *MagicTable) SetManager(manager *MagicManager) {
	mt.manager = manager
}

func (mt *MagicTable) SetCanbeAnyText(canBeAnyText string) {
	mt.canBeAnyText = canBeAnyText
}

func (mt *MagicTable) updateHeaderMetaData() (magicTableMetaDataType, bool, error) {
	columnMeta, searchbarIsVisible, err := mt.getColumnsMetadata()

	return columnMeta, searchbarIsVisible, err
}

// OnAppUpdate satisfies the app.Updater interface. It is called when the app is
// updated in background.
func (mt *MagicTable) OnAppUpdate(ctx app.Context) {
	mt.updateAvailable = ctx.AppUpdateAvailable // Reports that an app update is available.
	mt.Update()                                 // Triggers UI update.
}

func (mt *MagicTable) Render() app.UI {

	var err error
	//var searchInput app.HTMLDiv

	if firstSearch == true {

		// Initiate logrus logger
		mt.InitLogger("")

		// Only for testing
		mt.tableTypeGuid = "51253aba-41a9-42ef-b5f1-d8d1d7116b47"

		// No row is selected
		mt.rowSelected = -1
		mt.uniqueRowSelected = -1

		// Set Table State
		mt.tableState = TableState_List

		mt.RetrieveTableDataFromDB("")
		firstSearch = false

		// Try to update manager-ref
		magicManager.magicTable.SetManager(magicManager)

		app.Window().AddEventListener("hide.bs.modal", mt.onCloseModalWrapper())
		//app.Window().AddEventListener("closed.bs.alert", mt.onCloseAlertWrapper())

		mt.alertId = 4

		messagesToAlertToUser := alertMessageStruct{
			id:           "alert0",
			alertType:    alertType_warning,
			alertMessage: "alert0: Detta är en varning!",
			processCount: 0,
			show:         true,
		}
		mt.messagesToAlertToUser = append(mt.messagesToAlertToUser, messagesToAlertToUser)

		messagesToAlertToUser = alertMessageStruct{
			id:           "alert1",
			alertType:    alertType_info,
			alertMessage: "alert1: Detta är ett infomeddelande.",
			processCount: 0,
			show:         true,
		}
		mt.messagesToAlertToUser = append(mt.messagesToAlertToUser, messagesToAlertToUser)

		messagesToAlertToUser = alertMessageStruct{
			id:           "alert2",
			alertType:    alertType_success,
			alertMessage: "alert2: Detta är ett meddelande då det gick bra.",
			processCount: 0,
			show:         true,
		}
		mt.messagesToAlertToUser = append(mt.messagesToAlertToUser, messagesToAlertToUser)

		messagesToAlertToUser = alertMessageStruct{
			id:           "alert3",
			alertType:    alertType_danger,
			alertMessage: "alert3: Detta är ett meddelande då det gick rikigt DÅLIGT!!!",
			processCount: 0,
			show:         true,
		}
		mt.messagesToAlertToUser = append(mt.messagesToAlertToUser, messagesToAlertToUser)

		// Call DB to get tableTypeSelectorOptions
		err = mt.getTableTypeSelectorOptionsFromDB()
		if err != nil {
			fmt.Println(err)
		}

	} //else {

	//app.Window().AddEventListener("hide.bs.modal", mt.onCloseModalWrapper())

	//}

	// Retrieve headers metadata, if not have been done
	if (mt.testDataAndMetaData.magicTableMetaData == nil || mt.reloadHeaderMetaData == true) && len(mt.tableTypeGuid) > 0 {
		mt.reloadHeaderMetaData = false
		_, mt.searchbarIsVisible, err = mt.updateHeaderMetaData()
		if err != nil {
			fmt.Println(" mt.updateHeaderMetaData() generated Error:", err)

			// If something went wrong, return a message instead
			return app.H1().Text("Couldn't Retrieve Magic table Header Metadata")
		}

	}

	//mt.messagesToAlertToUser = nil
	//fmt.Println("Antal Alerts2: " + strconv.Itoa(len(mt.messagesToAlertToUser)))

	magicTableRenderedObject, err := mt.AssembleMagicalTableObject()
	if err != nil {
		fmt.Println("Problem assemble the magicTable page object")
	}

	return magicTableRenderedObject

}

func (mt *MagicTable) removeIndexFromMagicTable(s []alertMessageStruct, index int) []alertMessageStruct {
	return append(s[:index], s[index+1:]...)
}

func (mt *MagicTable) GetRowTextBoxValueForEdit(columnDataName string) string {

	var returnValue string

	switch mt.tableState {

	case TableState_New:
		returnValue = ""

	case TableState_Edit,
		TableState_Delete:
		rowData := mt.testDataAndMetaData.originalTestdataInstances[mt.rowSelected]
		//returnValue =rowData.GetName()

		r := reflect.ValueOf(rowData)
		f := reflect.Indirect(r).FieldByName(columnDataName)

		returnValue = fmt.Sprintf("%v", f)

	default:
		returnValue = ""
	}

	return returnValue
}

func (mt *MagicTable) StateCheckToShowBaseButtons() bool {
	switch mt.tableState {
	case
		TableState_Search_Load,
		TableState_List:
		return true
	}
	return false
}

func (mt *MagicTable) SetSortIconForTableHeader(columnCounter int) string {
	var returnMessage string
	//fmt.Println("mt.columnToSortOn: ", mt.columnToSortOn, mt.testDataAndMetaData.magicTableMetaData[columnCounter].String(), mt.testDataAndMetaData.magicTableMetaData[columnCounter].Sortable)

	// Only Set icon on correct Header

	if mt.testDataAndMetaData.magicTableMetaData[columnCounter].Sortable == true {
		// Column is sortable
		if columnCounter == mt.columnToSortOn && firstTimeForSorting == false {
			if mt.columnSortOrderIsAscending {
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

//func (p *InstanceTable)  onClick(ctx app.Context, e app.Event) {
//	fmt.Println("onClick is called")
//}
func (mt *MagicTable) MyOnDblClickOnRowWrapper(rowThatWasDoubleClickedOn int64) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		mt.MyOnDblClickOnRow(rowThatWasDoubleClickedOn)

		/*
			// Only allow click if Table state is in 'List'
			if  mt.tableState != TableState_List {
				return
			}


			fmt.Println("OnDblClick is called::::::" + strconv.Itoa(rowThatWasDoubleClickedOn))
			rowData := mt.testDataAndMetaData.originalTestdataInstances[rowThatWasDoubleClickedOn]
			fmt.Println("RowData: " + rowData.String())
			fmt.Println("")


			// Trigger Edit
			mt.onButtonClickWrapper(EditButton)
		*/

	}
}

func (mt *MagicTable) MyOnDblClickOnRow(rowThatWasDoubleClickedOn int64) {
	// Only allow click if Table state is in 'List'
	if mt.tableState != TableState_List {
		return
	}

	fmt.Println("OnDblClick is called::::::" + strconv.FormatInt(rowThatWasDoubleClickedOn, 10))
	rowData := mt.testDataAndMetaData.originalTestdataInstances[rowThatWasDoubleClickedOn]
	fmt.Println("RowData: " + rowData.String())
	fmt.Println("")

	// Trigger Edit
	mt.onButtonClick(EditButton)
}

func (mt *MagicTable) MyOnClickOnRowWrapper(rowThatWasClickedOn int64) app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		// Only allow click if Table state is in 'List'
		if mt.tableState != TableState_List {
			return
		}

		fmt.Println("OnClick is called::::::" + strconv.FormatInt(rowThatWasClickedOn, 10))
		rowData := mt.testDataAndMetaData.originalTestdataInstances[rowThatWasClickedOn]
		fmt.Println("RowData: " + rowData.String())
		fmt.Println("")

		mt.rowSelected = rowThatWasClickedOn
		mt.uniqueRowSelected = rowData.GetUniqueId()
		mt.Update()
	}
}

func (mt *MagicTable) OnDblClickapp(src app.Value, e app.Event) {

	// Only allow click if Table state is in 'List'
	if mt.tableState != TableState_List {
		return
	}
	fmt.Println("OnDblClick is called - 1025")
	fmt.Println("src - " + src.JSValue().String())
	fmt.Println("e - " + e.Value.String())

}

func (mt *MagicTable) OnTextboxDblClickappWrapper(buttonId string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		elem := app.Window().GetElementByID(buttonId)
		fmt.Println("elem - ", buttonId, elem.Get("value").String())
		//elem.Set("value", "APA")

	}
}

func (mt *MagicTable) areNewUpdateDeleteTextBoxesDisabled() (textBoxesAreEnabled bool) {

	switch mt.tableState {
	case TableState_New,
		TableState_Edit:
		textBoxesAreEnabled = false

	default:
		textBoxesAreEnabled = true

	}

	return textBoxesAreEnabled

}

func (mt *MagicTable) onRefreshButtonClickWrapper() app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		/*
			alertIdStr := "alert" + strconv.Itoa(mt.alertId)
			mt.alertId = mt.alertId + 1

			messagesToAlertToUser := alertMessageStruct{
				id:           alertIdStr,
				alertType:    alertType_danger,
				alertMessage: alertIdStr + ": Detta är ett NYTT meddelande då det gick rikigt DÅLIGT!!!",
				processCount: 0,
				show:         true,
			}
			mt.messagesToAlertToUser = append(mt.messagesToAlertToUser, messagesToAlertToUser)


		*/
		//for _, a := range mt.messagesToAlertToUser {
		//	fmt.Println(a.id, app.Window().GetElementByID(a.id), app.Window().GetElementByID(a.id).IsNull())
		//}

		mt.Update()
		/*
			ctx.Async(func() {
				time.Sleep(2 * time.Second)
				app.Window().GetElementByID("alert0button").Call("click")
			})


		*/
	}
}
func (mt *MagicTable) onCloseAlertWrapper(alertId string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		fmt.Println("***Alert-box (" + alertId + ") stängdes av användaren***")
		//app.Window().Call("hideAlert", alertId)
		for alertCounter, _ := range mt.messagesToAlertToUser {
			if mt.messagesToAlertToUser[alertCounter].id == alertId {
				fmt.Println("hittade")
				mt.messagesToAlertToUser[alertCounter].show = false
				break
			}
		}

		//app.Window().GetElementByID("alert3").Call("focus")
		//mt.Update()
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

func (mt *MagicTable) getColumnsMetadata() (magicTableMetaDataType, bool, error) {

	mt.logger.WithFields(logrus.Fields{
		"Id": "ca5dd6c2-b071-4d15-b406-ced54c3e124e",
	}).Debug("Entering: getColumnsMetadata()")

	defer func() {
		mt.logger.WithFields(logrus.Fields{
			"Id": "08e24c3b-6872-402e-aac0-067d8f043561",
		}).Debug("Exiting: getColumnsMetadata()")
	}()

	magicTableMetadataRequest := api.MagicTableMetadataRequest{
		MagicTableMetadataType: mt.tableType,
		TableTypeGuid:          mt.tableTypeGuid,
	}
	magicTableMetadataRespons, err := api.CallApiGetMagicTableMetadata(magicTableMetadataRequest)

	if err != nil {
		fmt.Println("api.CallApiGetMagicTableMetadata Error:", err)
		return nil, false, err
	} else {
		mt.testDataAndMetaData.magicTableMetaData = magicTableMetadataRespons.GetMagicTableColumnsMetadata()
	}

	return magicTableMetadataRespons.GetMagicTableColumnsMetadata(), magicTableMetadataRespons.SearchbarIsVisible, err
}

func (mt *MagicTable) MyOnInputChange(ctx app.Context, e app.Event) {
	mt.searchString = ctx.JSSrc.Get("value").String()
	//mt.searchString = src.Get("value").String()
	//mt.Update()
	mt.RetrieveTableDataFromDB(mt.searchString)
}

// Get all tables that can be edited
func (mt *MagicTable) getTableTypeSelectorOptionsFromDB() error {
	instances, err := api.CallApiListTablesToEdit(api.EmptyParameter{})

	if err != nil {
		fmt.Println("SearchInDB Error:", err)
		return err
	}

	mt.tableTypeSelectorOptionsInDB = instances.MyListTableToEdit

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
