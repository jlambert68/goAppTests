package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
	"sort"
	"strconv"
)

func (mt *MagicTable) MyOnColumnClickWrapper(columnNumberThatWasClicked int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		mt.logger.WithFields(logrus.Fields{
			"Id": "fa5a1cf2-06ed-4710-86ef-dbd316ad855c",
		}).Debug("Entering: MyOnColumnClickWrapper()")

		defer func() {
			mt.logger.WithFields(logrus.Fields{
				"Id": "d069adc2-8944-48b6-a9b2-1a6a7c1cca86",
			}).Debug("Exiting: MyOnColumnClickWrapper()")
		}()

		// Only allow click if Table state is in 'List'
		if mt.tableState != TableState_List {
			return
		}

		fmt.Println("MyOnColumnClickWrapper is called:::::: " + strconv.Itoa(columnNumberThatWasClicked))
		fmt.Println("")

		// Only sort if column is sortable
		if mt.testDataAndMetaData.magicTableMetaData[columnNumberThatWasClicked].Sortable == true {

			// Check if table for sorted before
			if firstTimeForSorting == true {
				firstTimeForSorting = false

				// First time
				mt.columnToSortOn = columnNumberThatWasClicked
				mt.columnSortOrderIsAscending = true

				// Not the first time
			} else {
				// If the same column was clicked then invert sort order for column
				if mt.columnToSortOn == columnNumberThatWasClicked {
					mt.columnSortOrderIsAscending = !mt.columnSortOrderIsAscending
				} else {
					// New column to sort on
					mt.columnToSortOn = columnNumberThatWasClicked
					mt.columnSortOrderIsAscending = true
				}
			}

			// Sort the current data
			switch mt.tableTypeGuid {

			// Original Test table
			case "51253aba-41a9-42ef-b5f1-d8d1d7116b47":
				sort.SliceStable(mt.testDataAndMetaData.originalTestdataInstances, mt.sliceSorter)

				// TestDomains
			case "8acacaaf-676e-4b36-abe6-c5310822ade1":
				sort.SliceStable(mt.testDataAndMetaData.testDomains, mt.sliceSorter)

			default:
				mt.logger.WithFields(logrus.Fields{
					"Id":               "534e486c-69a9-4364-b562-7ccf8943e0e7",
					"mt.tableTypeGuid": mt.tableTypeGuid,
				}).Panic("Unknown table reference")

			}

			mt.Update()

			// Update selected row
			currentTableDataPointerData := reflect.ValueOf(mt.currentTableDataPointer)
			numberOfRowsInTable := reflect.Indirect(currentTableDataPointerData).Len()
			//fmt.Println("numberOfRowsInTable:: ", numberOfRowsInTable)

			// Loop over all data rows
			for rowCounter := 1; rowCounter < numberOfRowsInTable+1; rowCounter++ {
				uniqueId := mt.getUniqueId(rowCounter)

				if uniqueId == mt.uniqueRowSelected {
					mt.rowSelected = int64(rowCounter)

				}
			}
		}
	}
}

// Function that is passed into "sort.Slicestable" to be able to sort
func (mt *MagicTable) sliceSorter(i, j int) bool {

	var compareResult bool
	var value1 reflect.Value
	var value2 reflect.Value

	fieldsToExtract := mt.testDataAndMetaData.magicTableMetaData[mt.columnToSortOn].GetColumnDataName()
	fmt.Println("fieldsToExtract", fieldsToExtract)

	switch mt.tableTypeGuid {

	// Original Test table
	case "51253aba-41a9-42ef-b5f1-d8d1d7116b47":
		value1 = getAttr(mt.testDataAndMetaData.originalTestdataInstances[i], fieldsToExtract)
		value2 = getAttr(mt.testDataAndMetaData.originalTestdataInstances[j], fieldsToExtract)

		// TestDomains
	case "8acacaaf-676e-4b36-abe6-c5310822ade1":
		value1 = getAttr(mt.testDataAndMetaData.testDomains[i], fieldsToExtract)
		value2 = getAttr(mt.testDataAndMetaData.testDomains[j], fieldsToExtract)

	default:
		mt.logger.WithFields(logrus.Fields{
			"Id":               "14f9d72b-d768-4933-bfb1-4d86b00ff972",
			"mt.tableTypeGuid": mt.tableTypeGuid,
		}).Panic("Unknown table reference")

	}
	//value1 := getAttr(reflect.Indirect(currentTableData).Index(i-1), fieldsToExtract)
	//value2 := getAttr(reflect.Indirect(currentTableData).Index(j-1), fieldsToExtract)

	switch mt.testDataAndMetaData.magicTableMetaData[mt.columnToSortOn].ColumnDataType {
	case api.MagicTableColumnDataType_String:
		compareResult = value1.String() < value2.String()
	case api.MagicTableColumnDataType_Float:
		compareResult = value1.Float() < value2.Float()
	}

	return xnor(mt.columnSortOrderIsAscending, compareResult)

}

//https://play.golang.org/p/pxj_gCnvF1J
func getAttr(obj interface{}, fieldName string) reflect.Value {
	fmt.Println("getAttr", obj, fieldName)
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
