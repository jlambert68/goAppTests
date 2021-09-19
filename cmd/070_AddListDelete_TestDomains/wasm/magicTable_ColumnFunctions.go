package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
	"sort"
	"strconv"
)

func (mt *MagicTable) MyOnColumnClickWrapper(columnNumberThatWasClicked int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {

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

			var compareResult bool
			// Sort by age, keeping original order or equal elements.
			sort.SliceStable(mt.testDataAndMetaData.originalTestdataInstances, func(i, j int) bool {

				// General solution - Nice
				if true {
					fieldsToExtract := mt.testDataAndMetaData.magicTableMetaData[columnNumberThatWasClicked].GetColumnDataName()
					//fmt.Println("fieldsToExtract", fieldsToExtract)
					value1 := getAttr(mt.testDataAndMetaData.originalTestdataInstances[i], fieldsToExtract)
					value2 := getAttr(mt.testDataAndMetaData.originalTestdataInstances[j], fieldsToExtract)

					switch mt.testDataAndMetaData.magicTableMetaData[columnNumberThatWasClicked].ColumnDataType {
					case api.MagicTableColumnDataType_String:
						compareResult = value1.String() < value2.String()
					case api.MagicTableColumnDataType_Float:
						compareResult = value1.Float() < value2.Float()
					}

					return xnor(mt.columnSortOrderIsAscending, compareResult)

				} else {
					// Hard defined sorting-types - Should be removed
					switch mt.columnToSortOn {
					case 0:
						return xnor(mt.columnSortOrderIsAscending, mt.testDataAndMetaData.originalTestdataInstances[i].Name < mt.testDataAndMetaData.originalTestdataInstances[j].Name)

					case 1:
						return xnor(mt.columnSortOrderIsAscending, mt.testDataAndMetaData.originalTestdataInstances[i].InstanceType < mt.testDataAndMetaData.originalTestdataInstances[j].InstanceType)

					case 2:
						return xnor(mt.columnSortOrderIsAscending, mt.testDataAndMetaData.originalTestdataInstances[i].Ecu < mt.testDataAndMetaData.originalTestdataInstances[j].Ecu)

					case 3:
						return xnor(mt.columnSortOrderIsAscending, mt.testDataAndMetaData.originalTestdataInstances[i].Memory < mt.testDataAndMetaData.originalTestdataInstances[j].Memory)

					case 4:
						return xnor(mt.columnSortOrderIsAscending, mt.testDataAndMetaData.originalTestdataInstances[i].Network < mt.testDataAndMetaData.originalTestdataInstances[j].Network)

					case 5:
						return xnor(mt.columnSortOrderIsAscending, mt.testDataAndMetaData.originalTestdataInstances[i].Price < mt.testDataAndMetaData.originalTestdataInstances[j].Price)

					default:
						fmt.Println("Could not sort on column: " + strconv.Itoa(columnNumberThatWasClicked))
						fmt.Println("Will sort on Name-ascending, instead: ")
						return xnor(true, mt.testDataAndMetaData.originalTestdataInstances[i].Name < mt.testDataAndMetaData.originalTestdataInstances[j].Name)
					}
				}

			})
			mt.Update()
			//magicManager.Update()
			//TODO Fix for all tables that are supported

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
			/*
				for rowCounter, i := range mt.testDataAndMetaData.originalTestdataInstances {

					// Check if selected row
					if i.UniqueId == mt.uniqueRowSelected {
						mt.rowSelected = int64(rowCounter)
					}
				}

			*/

		}
	}
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
