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

func (mt *MagicTable) MyOnColumnClickWrapper(sortOrderThatWasClicked int) app.EventHandler {
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

		fmt.Println("MyOnColumnClickWrapper is called:::::: " + strconv.Itoa(sortOrderThatWasClicked))
		fmt.Println("")
		currentTableDataPointerData := reflect.ValueOf(mt.currentTableDataPointer)
		numberOfRowsInTable := reflect.Indirect(currentTableDataPointerData).Len()
		for rowCounter := 1; rowCounter < numberOfRowsInTable+1; rowCounter++ {
			uniqueId := mt.getUniqueId(rowCounter)
			fmt.Println("uniqueId", uniqueId)
		}
		fmt.Println("mt.columnSortOrderIsAscending before change direction", mt.columnSortOrderIsAscending)

		// Only sort if column is sortable
		for iColumnCounter, magicTableMetaData := range mt.testDataAndMetaData.magicTableMetaData {

			if sortOrderThatWasClicked == int(mt.testDataAndMetaData.magicTableMetaData[iColumnCounter].PresentationOrder) {

				if magicTableMetaData.Sortable == true {

					// Check if table for sorted before
					if firstTimeForSorting == true {
						firstTimeForSorting = false

						// First time
						mt.columnToSortOn = iColumnCounter
						mt.columnSortOrderIsAscending = true

						// Not the first time
					} else {
						// If the same column was clicked then invert sort order for column
						if mt.columnToSortOn == iColumnCounter {
							mt.columnSortOrderIsAscending = !mt.columnSortOrderIsAscending
						} else {
							// New column to sort on
							mt.columnToSortOn = iColumnCounter
							mt.columnSortOrderIsAscending = true
						}
					}
				}
			}
			fmt.Println("mt.columnSortOrderIsAscending after change direction", mt.columnSortOrderIsAscending)

			// Sort the current data
			switch mt.tableTypeGuid {

			// Original Test table
			case "51253aba-41a9-42ef-b5f1-d8d1d7116b47":
				fmt.Println("sort.SliceStable(mt.testDataAndMetaData.originalTestdataInstances, mt.sliceSorter)")
				sort.SliceStable(mt.testDataAndMetaData.originalTestdataInstances, mt.sliceSorter)

				// TestDomains
			case "8acacaaf-676e-4b36-abe6-c5310822ade1":
				fmt.Println("sort.SliceStable(mt.testDataAndMetaData.testDomains, mt.sliceSorter)")
				sort.SliceStable(mt.testDataAndMetaData.testDomains, mt.sliceSorter)

			default:
				mt.logger.WithFields(logrus.Fields{
					"Id":               "534e486c-69a9-4364-b562-7ccf8943e0e7",
					"mt.tableTypeGuid": mt.tableTypeGuid,
				}).Panic("Unknown table reference")

			}

			// Update selected row
			currentTableDataPointerData := reflect.ValueOf(mt.currentTableDataPointer)
			numberOfRowsInTable := reflect.Indirect(currentTableDataPointerData).Len()
			//fmt.Println("numberOfRowsInTable:: ", numberOfRowsInTable)

			// Loop over all data rows
			for rowCounter := 1; rowCounter < numberOfRowsInTable+1; rowCounter++ {
				uniqueId := mt.getUniqueId(rowCounter)

				fmt.Println("uniqueId", uniqueId)

				if uniqueId == mt.uniqueRowSelected {
					mt.rowSelected = int64(rowCounter)

				}
			}
			fmt.Println("mt.columnSortOrderIsAscending after sort", mt.columnSortOrderIsAscending)

			mt.Update()
		}
	}
}

// Function that is passed into "sort.Slicestable" to be able to sort
func (mt *MagicTable) sliceSorter(i, j int) bool {
	// j= i -1

	var compareResult bool = false
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

	fmt.Println("mt.testDataAndMetaData.magicTableMetaData[mt.columnToSortOn].ColumnDataType", mt.testDataAndMetaData.magicTableMetaData[mt.columnToSortOn].ColumnDataType)

	switch mt.testDataAndMetaData.magicTableMetaData[mt.columnToSortOn].ColumnDataType {
	case api.MagicTableColumnDataType_String:
		compareResult = value1.String() < value2.String()
		fmt.Println("case api.MagicTableColumnDataType_String:", value1, "<", value2, compareResult)

	case api.MagicTableColumnDataType_Float:
		compareResult = value1.Float() < value2.Float()
		fmt.Println("case api.MagicTableColumnDataType_Float:", value1, "<", value2, compareResult)

	case api.MagicTableColumnDataType_Int:
		compareResult = value1.Int() < value2.Int()
		fmt.Println("case api.MagicTableColumnDataType_Int:", value1, "<", value2, compareResult)

	case api.MagicTableColumnDataType_Bool:
		compareResult = value1.Bool() == value2.Bool()
		fmt.Println("case api.MagicTableColumnDataType_Bool:", value1, "<", value2, compareResult)

	default:
		mt.logger.WithFields(logrus.Fields{
			"Id": "dbb4e088-f633-4588-8c40-1f7111fd8783",
			"mt.testDataAndMetaData.magicTableMetaData[mt.columnToSortOn].ColumnDataType": mt.testDataAndMetaData.magicTableMetaData[mt.columnToSortOn].ColumnDataType,
		}).Panic("Unknown column type to sort on")
	}

	xnorResults := xor(mt.columnSortOrderIsAscending, compareResult)
	fmt.Println("value1, value2, mt.columnSortOrderIsAscending, compareResult, xnorResults:: ", value1, value2, mt.columnSortOrderIsAscending, compareResult, xnorResults)

	return xnorResults
}

//https://play.golang.org/p/pxj_gCnvF1J
func getAttr(obj interface{}, fieldName string) reflect.Value {
	//fmt.Println("getAttr", obj, fieldName)
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
func xor(a, b bool) bool {
	return ((a && b) || (!a && !b)) //XNOR
	//return ((a || b) && (!a || !b)) //XOR
	//return ((a && b) || (!a && b))

}
