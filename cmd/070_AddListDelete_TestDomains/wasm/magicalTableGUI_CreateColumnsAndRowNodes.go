package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/maxence-charriere/go-app/v8/pkg/errors"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
)

// Generates the GUI objects for the columns in the table
func (mt *MagicTable) UpdateColumnsNodes() error {

	mt.logger.WithFields(logrus.Fields{
		"Id": "bc256819-240a-4a37-b1a0-41a298fbfd86",
	}).Debug("Entering: UpdateColumnsNodes()")

	var err error
	err = nil

	mt.tableColumnNodes = []app.UI{}
	for columnCounter, columnMetadataResponse := range mt.testDataAndMetaData.magicTableMetaData {
		// Check if column should be shown
		columnHeader := app.Th().
			Scope("col").
			Body(app.Text(columnMetadataResponse.GetColumnHeaderName() + mt.SetSortIconForTableHeader(columnCounter))).
			OnClick(mt.MyOnColumnClickWrapper(columnCounter))
		shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
		if shouldBeShown == true {
			mt.tableColumnNodes = append(mt.tableColumnNodes, columnHeader)
		}
	}

	mt.logger.WithFields(logrus.Fields{
		"Id": "1393c88b-0135-4a94-ab9b-f61fd5e18bad",
	}).Debug("Exiting: UpdateColumnsNodes()")

	return err
}

// Generates the GUI objects for the rows in the table
func (mt *MagicTable) UpdateRowNodes() error {

	mt.logger.WithFields(logrus.Fields{
		"Id": "f407b197-41ff-43bb-895b-50a033cc70f7",
	}).Debug("Entering: UpdateRowNodes()")

	var err error
	var columnData string
	var rowCounter int64
	var uniqueId int64
	var breakLoop bool
	err = nil

	mt.tableRowsNodes = []app.UI{}

	rowCounter = 0
	breakLoop = false
	for {
		rowNodes := []app.UI{}
		// Loop every column and get value and check if column should be shown

		for _, columnMetadataResponse := range mt.testDataAndMetaData.magicTableMetaData {
			columnName := columnMetadataResponse.GetColumnDataName()

			// Retrieve data from correct table
			switch mt.tableTypeGuid {

			// Original Test table
			case "51253aba-41a9-42ef-b5f1-d8d1d7116b47":
				columnData = mt.formatColumnData_OrignalTestdata(mt.testDataAndMetaData.originalTestdataInstances[rowCounter], columnName)
				uniqueId = mt.testDataAndMetaData.originalTestdataInstances[rowCounter].UniqueId

				// Check for last item
				if rowCounter+1 == int64(len(mt.testDataAndMetaData.originalTestdataInstances)) {
					breakLoop = true
				}

				// TestDomains
			case "8acacaaf-676e-4b36-abe6-c5310822ade1":
				//fmt.Println("mt.testDataAndMetaData.testDomains.len: ", len(mt.testDataAndMetaData.testDomains))
				columnData = mt.formatColumnData_TestDomains(mt.testDataAndMetaData.testDomains[rowCounter], columnName)
				uniqueId = mt.testDataAndMetaData.testDomains[rowCounter].UniqueId

				// Check for last item
				if rowCounter+1 == int64(len(mt.testDataAndMetaData.testDomains)) {
					breakLoop = true
				}

			// TestInstructions
			//case "81c5d008-a38a-4c47-936a-d6c3c258ae13":

			// Unknow Table type
			default:
				fmt.Println("Unknown table type1: ", mt.tableTypeGuid)
				return errors.New("Unknown table type when preparing table data: " + mt.tableTypeGuid)
			}

			shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
			if shouldBeShown == true {
				rowNodes = append(rowNodes, app.Td().Body(app.Text(columnData)))
			}

		}
		// Set background color on selected row
		var rowSelectedColor string
		if uniqueId != mt.uniqueRowSelected {
			rowSelectedColor = "#FFFFFF"
		} else {
			rowSelectedColor = "#E1EAFA"
			mt.rowSelected = rowCounter
		}

		// Append values for each row
		mt.tableRowsNodes = append(mt.tableRowsNodes,
			app.Tr().
				Body(rowNodes...).
				OnClick(mt.MyOnClickOnRowWrapper(rowCounter)).
				OnDblClick(mt.MyOnDblClickOnRowWrapper(rowCounter)).
				Style("background-color", rowSelectedColor))

		rowCounter = rowCounter + 1
		if breakLoop == true {
			break
		}
	}

	/*	switch mt.tableTypeGuid {

		// Original Test table
		case "51253aba-41a9-42ef-b5f1-d8d1d7116b47":
			//Loop every row
			for rowCounter, i := range mt.testDataAndMetaData.originalTestdataInstances {
				rowNodes := []app.UI{}
				// Loop every column and get value and check if column should be shown
				for _, columnMetadataResponse := range mt.testDataAndMetaData.magicTableMetaData {
					columnName := columnMetadataResponse.GetColumnDataName()

					columnData = mt.formatColumnData_OrignalTestdata(i, columnName)

					shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
					if shouldBeShown == true {
						rowNodes = append(rowNodes, app.Td().Body(app.Text(columnData)))
					}
				}

				// Set background color on selected row
				var rowSelectedColor string
				if i.UniqueId != mt.uniqueRowSelected {
					rowSelectedColor = "#FFFFFF"
				} else {
					rowSelectedColor = "#E1EAFA"
					mt.rowSelected = rowCounter
				}

				// Append values for each row
				mt.tableRowsNodes = append(mt.tableRowsNodes,
					app.Tr().
						Body(rowNodes...).
						OnClick(mt.MyOnClickOnRowWrapper(rowCounter)).
						OnDblClick(mt.MyOnDblClickOnRowWrapper(rowCounter)).
						Style("background-color", rowSelectedColor))

			}
			// TestDomains
		case "8acacaaf-676e-4b36-abe6-c5310822ade1":
			columnData = mt.formatColumnData_TestDomains(i, columnName)

		// TestInstructions
		//case "81c5d008-a38a-4c47-936a-d6c3c258ae13":


			// Unknow Table type
		default:
			fmt.Println("Unknown table type: ", mt.tableTypeGuid)
			return errors.New("Unknown table type when preparing table data: " + mt.tableTypeGuid)
		}



	*/

	mt.logger.WithFields(logrus.Fields{
		"Id": "b9342e63-8baf-482c-b110-75c366b5999d",
	}).Debug("Exiting: UpdateRowNodes()")

	return err
}

func (mt *MagicTable) formatColumnData_OrignalTestdata(v *api.Instance, field string) string {

	var returnValue string

	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	returnValue = fmt.Sprintf("%v", f)

	return returnValue
}

func (mt *MagicTable) formatColumnData_TestDomains(v *api.TestDomainForListingMessage, field string) string {

	fmt.Println(v, " - ", field)

	var returnValue string

	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	fmt.Println(r, " - ", f)

	returnValue = fmt.Sprintf("%v", f)

	return returnValue
}
