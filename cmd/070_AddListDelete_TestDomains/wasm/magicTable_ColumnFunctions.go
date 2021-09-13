package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
	"sort"
	"strconv"
)

func (p *MagicTable) MyOnColumnClickWrapper(columnNumberThatWasClicked int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {

		// Only allow click if Table state is in 'List'
		if p.tableState != TableState_List {
			return
		}

		fmt.Println("MyOnColumnClickWrapper is called:::::: " + strconv.Itoa(columnNumberThatWasClicked))
		fmt.Println("")

		// Only sort if column is sortable
		if p.testDataAndMetaData.magicTableMetaData[columnNumberThatWasClicked].Sortable == true {

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
			sort.SliceStable(p.testDataAndMetaData.originalTestdataInstances, func(i, j int) bool {

				// General solution - Nice
				if true {
					fieldsToExtract := p.testDataAndMetaData.magicTableMetaData[columnNumberThatWasClicked].GetColumnDataName()
					//fmt.Println("fieldsToExtract", fieldsToExtract)
					value1 := getAttr(p.testDataAndMetaData.originalTestdataInstances[i], fieldsToExtract)
					value2 := getAttr(p.testDataAndMetaData.originalTestdataInstances[j], fieldsToExtract)

					switch p.testDataAndMetaData.magicTableMetaData[columnNumberThatWasClicked].ColumnDataType {
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
						return xnor(p.columnSortOrderIsAscending, p.testDataAndMetaData.originalTestdataInstances[i].Name < p.testDataAndMetaData.originalTestdataInstances[j].Name)

					case 1:
						return xnor(p.columnSortOrderIsAscending, p.testDataAndMetaData.originalTestdataInstances[i].InstanceType < p.testDataAndMetaData.originalTestdataInstances[j].InstanceType)

					case 2:
						return xnor(p.columnSortOrderIsAscending, p.testDataAndMetaData.originalTestdataInstances[i].Ecu < p.testDataAndMetaData.originalTestdataInstances[j].Ecu)

					case 3:
						return xnor(p.columnSortOrderIsAscending, p.testDataAndMetaData.originalTestdataInstances[i].Memory < p.testDataAndMetaData.originalTestdataInstances[j].Memory)

					case 4:
						return xnor(p.columnSortOrderIsAscending, p.testDataAndMetaData.originalTestdataInstances[i].Network < p.testDataAndMetaData.originalTestdataInstances[j].Network)

					case 5:
						return xnor(p.columnSortOrderIsAscending, p.testDataAndMetaData.originalTestdataInstances[i].Price < p.testDataAndMetaData.originalTestdataInstances[j].Price)

					default:
						fmt.Println("Could not sort on column: " + strconv.Itoa(columnNumberThatWasClicked))
						fmt.Println("Will sort on Name-ascending, instead: ")
						return xnor(true, p.testDataAndMetaData.originalTestdataInstances[i].Name < p.testDataAndMetaData.originalTestdataInstances[j].Name)
					}
				}

			})
			p.Update()
			//magicManager.Update()

			// Update selected row
			for rowCounter, i := range p.testDataAndMetaData.originalTestdataInstances {

				// Check if selected row
				if i.UniqueId == p.uniqueRowSelected {
					p.rowSelected = rowCounter
				}
			}

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
