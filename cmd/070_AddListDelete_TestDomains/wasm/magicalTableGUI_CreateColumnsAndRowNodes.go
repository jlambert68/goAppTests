package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
	"strconv"
)

// Generates the GUI objects for the columns in the table
func (mt *MagicTable) UpdateColumnsNodes() error {

	mt.logger.WithFields(logrus.Fields{
		"Id": "bc256819-240a-4a37-b1a0-41a298fbfd86",
	}).Debug("Entering: UpdateColumnsNodes()")

	var err error
	err = nil

	mt.tableColumnNodes = []app.UI{}
	for _, columnMetadataResponse := range mt.testDataAndMetaData.magicTableMetaData {
		// Check if column should be shown
		//fmt.Println("columnMetadataResponse ", columnMetadataResponse)
		columnHeader := app.Th().
			Scope("col").
			Body(app.Text(columnMetadataResponse.GetColumnHeaderName() + mt.SetSortIconForTableHeader(int(columnMetadataResponse.PresentationOrder)))).
			OnClick(mt.MyOnColumnClickWrapper(int(columnMetadataResponse.PresentationOrder)))
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
		"Id": "699e1ec1-135a-458b-a571-621155184857",
	}).Debug("Entering: UpdateRowNodes()")

	defer func() {
		mt.logger.WithFields(logrus.Fields{
			"Id": "7dd8326b-c503-4b18-9632-a0069f3be08e",
		}).Debug("Exiting: UpdateRowNodes()")
	}()

	var err error
	var columnData string
	var uniqueId int64
	err = nil

	mt.tableRowsNodes = []app.UI{}

	currentTableDataPointerData := reflect.ValueOf(mt.currentTableDataPointer)
	numberOfRowsInTable := reflect.Indirect(currentTableDataPointerData).Len()
	//fmt.Println("numberOfRowsInTable:: ", numberOfRowsInTable)

	// Loop over all data rows
	for rowCounter := 1; rowCounter < numberOfRowsInTable+1; rowCounter++ {
		rowNodes := []app.UI{}
		uniqueId = -1
		//fmt.Println("rowCounter::", rowCounter)

		// Loop every column and get header value and check if column should be shown
		for _, columnMetadataResponse := range mt.testDataAndMetaData.magicTableMetaData {
			columnName := columnMetadataResponse.GetColumnDataName()

			// Get column data as string
			columnData = mt.formatColumnData(rowCounter, columnName)

			//fmt.Println("columnName - columnData:: ", columnName, columnData)

			//fmt.Println("mt.getUniqueId(rowCounter)", mt.getUniqueId(rowCounter))

			// Only retrieve UniqueId once per row
			if uniqueId == -1 {
				uniqueId = mt.getUniqueId(rowCounter)
			}
			//fmt.Println("uniqueId:: ", uniqueId)

			// Check if column data should be visible or not
			shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
			if shouldBeShown == true {
				rowNodes = append(rowNodes, app.Td().Body(app.Text(columnData)))
			}

		}

		// Set background color if the row was selected by user
		var rowSelectedColor string
		if uniqueId != mt.uniqueRowSelected {
			rowSelectedColor = "#FFFFFF"
		} else {
			rowSelectedColor = "#E1EAFA"
			mt.rowSelected = int64(rowCounter)
		}

		// Append values for each row
		mt.tableRowsNodes = append(mt.tableRowsNodes,
			app.Tr().
				Body(rowNodes...).
				OnClick(mt.MyOnClickOnRowWrapper(int64(rowCounter))).
				OnDblClick(mt.MyOnDblClickOnRowWrapper(int64(rowCounter))).
				Style("background-color", rowSelectedColor))

	}
	//mt.tableRowsNodes = []app.UI{}

	/*
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

					r := reflect.ValueOf(mt.currentTableDataPointer)
					f := reflect.Indirect(r).Len()
					fmt.Println("mt.currentTableDataPointer.LEN::::", f)
					fmt.Println("XXXXXXXDetta ska jag kolla efter")


					columnData = mt.formatColumnData_OrignalTestdata(mt.testDataAndMetaData.originalTestdataInstances[rowCounter], columnName)
					uniqueId = mt.testDataAndMetaData.originalTestdataInstances[rowCounter].UniqueId

					// Check for last item
					if rowCounter+1 == int64(len(mt.testDataAndMetaData.originalTestdataInstances)) {
						breakLoop = true
					}

					// TestDomains
				case "8acacaaf-676e-4b36-abe6-c5310822ade1":
					fmt.Println("mt.testDataAndMetaData.testDomains.len: ", len(mt.testDataAndMetaData.testDomains))
					sp := to_struct_ptr(mt.testDataAndMetaData.testDomains)
					test_ptr(sp)
					fmt.Println("sp is",sp)
					r := reflect.ValueOf(sp)
					f := reflect.Indirect(r).Slice(0,1)
					flen := reflect.Indirect(r).Len()
					fmt.Println("f:: ", f)
					fmt.Println("Len(): ", flen)


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
	*/

	return err
}

func (mt *MagicTable) formatColumnData(rowNumber int, field string) string {

	var returnValue string

	currentTableDataPointerData := reflect.ValueOf(mt.currentTableDataPointer)
	dataSlice := reflect.Indirect(currentTableDataPointerData).Index(rowNumber - 1)
	//fmt.Println("XXXX dataSlice::field:", dataSlice, field)

	keyValue := reflect.Indirect(dataSlice).FieldByName(field)
	//fmt.Println("keyValue", keyValue)
	/*
		fmt.Println("XXXX dataSlice::field:", dataSlice, field)

		fieldByName := reflect.Indirect(currentTableDataPointerData).Index(rowNumber).FieldByName(field)

		fmt.Println("XXXX fieldByName, field:", fieldByName,field)

		f := reflect.Indirect(dataSlice).FieldByName(field)

	*/

	returnValue = fmt.Sprintf("%v", keyValue)

	return returnValue
}

func (mt *MagicTable) getUniqueId(rowNumber int) int64 {

	var returnValue int64

	currentTableDataPointerData := reflect.ValueOf(mt.currentTableDataPointer)
	dataSlice := reflect.Indirect(currentTableDataPointerData).Index(rowNumber - 1)

	keyValue := reflect.Indirect(dataSlice).FieldByName("UniqueId")

	returnValueString := fmt.Sprintf("%v", keyValue)

	returnValue, err := strconv.ParseInt(returnValueString, 10, 64)
	if err != nil {
		fmt.Printf("%d of type %T\n", returnValue, returnValue)
		//fmt.Println("")
	}

	//fmt.Println("getUniqueId", rowNumber, dataSlice, keyValue, returnValueString, returnValue, err )

	return returnValue
}

func (mt *MagicTable) formatColumnData_OrignalTestdata(v *api.Instance, field string) string {

	var returnValue string

	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	returnValue = fmt.Sprintf("%v", f)

	return returnValue
}

func (mt *MagicTable) formatColumnData_TestDomains(v *api.TestDomainForListingMessage, field string) string {

	var returnValue string

	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	returnValue = fmt.Sprintf("%v", f)

	return returnValue
}

// From
// https://play.golang.org/p/hYnsAijyCE

//
// Dump out a pointer ...
//
func test_ptr(ptr interface{}) {
	v := reflect.ValueOf(ptr)
	fmt.Println("ptr is a", v.Type(), "to a", reflect.Indirect(v).Type())
}

//
// Return a pointer to the supplied struct via interface{}
//
func to_struct_ptr(obj interface{}) interface{} {

	//fmt.Println("obj is a", reflect.TypeOf(obj).Name())

	// Create a new instance of the underlying type
	vp := reflect.New(reflect.TypeOf(obj))

	// Should be a *Cat and Cat respectively
	//fmt.Println("vp is", vp.Type(), " to a ", vp.Elem().Type())

	vp.Elem().Set(reflect.ValueOf(obj))

	// NOTE: `vp.Elem().Set(reflect.ValueOf(&obj).Elem())` does not work

	// Return a `Cat` pointer to obj -- i.e. &obj.(*Cat)
	return vp.Interface()
}

/*

func main() {
	cat := Cat{name: "Fuzzy Wuzzy"}

	// Reports "*main.Cat"
	test_ptr(&cat)

	// Get a "*Cat" generically via interface{}
	sp := to_struct_ptr(cat)

	// *should* report "*main.Cat" also
	test_ptr(sp)

	fmt.Println("sp is",sp)
}
*/
