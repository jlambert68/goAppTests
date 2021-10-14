package server

import (
	"fmt"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
	"sort"
)

var currentTableDataPointer interface{}

var orignalData []*api.Instance

var currentTableDataAsSlice []*api.Instance

func forDebuggingSortData(myInstances []*api.Instance) {

	if len(myInstances) > 5 {
		return
	}

	orignalData = myInstances

	// pointer to data
	currentTableDataPointer = to_struct_ptr(orignalData)

	currentTableDataPointerData := reflect.ValueOf(currentTableDataPointer)
	currentTableData := reflect.Indirect(currentTableDataPointerData)
	currentTableDataAsSlice := reflect.Indirect(currentTableDataPointerData).Slice(1, currentTableData.Len())
	//currentTableDataType := reflect.Indirect(currentTableDataPointerData).Type()

	//fmt.Println("mt.testDataAndMetaData.originalTestdataInstances:: ", mt.testDataAndMetaData.originalTestdataInstances)
	//fmt.Println("currentTableDataPointerData:: ", currentTableDataPointerData)
	//fmt.Println("currentTableData:: ", currentTableData, currentTableDataAsSlice)
	//fmt.Println("&mt.currentTableDataPointer:: ", &mt.currentTableDataPointer)
	//fmt.Println("currentTableDataType:: ", currentTableDataType)
	//fmt.Println("mt.testDataAndMetaData.originalTestdataInstancesType:: ", mt.testDataAndMetaData.originalTestdataInstances)

	valuePtr := reflect.ValueOf(currentTableDataPointer)
	value := valuePtr.Elem()
	//fmt.Println("value := valuePtr.Elem() ", value, reflect.TypeOf(value).Kind())

	instanceSlice := make([]*api.Instance, 0)
	sliceType := reflect.TypeOf(instanceSlice)

	//use of MakeSlice() method
	intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)

	v := currentTableDataAsSlice
	rv := reflect.ValueOf(v)
	intSliceReflect = reflect.Append(intSliceReflect, rv)
	//fmt.Println("intSliceReflect", intSliceReflect, reflect.TypeOf(intSliceReflect).Kind())

	//intSlice2 := intSliceReflect.Interface().([]int)
	//fmt.Println(intSlice2)

	//v := ValueOf(slice)

	//if v.Kind() != Slice

	sort.SliceStable(value, func(i, j int) bool {
		var compareResult bool
		// General solution - Nice

		fieldsToExtract := "Name" //mt.testDataAndMetaData.magicTableMetaData[columnNumberThatWasClicked].GetColumnDataName()
		//fmt.Println("fieldsToExtract", fieldsToExtract)
		//value1 := getAttr(mt.testDataAndMetaData.originalTestdataInstances[i], fieldsToExtract)
		//value2 := getAttr(mt.testDataAndMetaData.originalTestdataInstances[j], fieldsToExtract)
		value1 := getAttr(reflect.Indirect(currentTableDataPointerData).Index(i-1), fieldsToExtract)
		value2 := getAttr(reflect.Indirect(currentTableDataPointerData).Index(j-1), fieldsToExtract)

		//switch mt.testDataAndMetaData.magicTableMetaData[columnNumberThatWasClicked].ColumnDataType {
		//case api.MagicTableColumnDataType_String:
		compareResult = value1.String() < value2.String()
		//case api.MagicTableColumnDataType_Float:
		//	compareResult = value1.Float() < value2.Float()
		//}

		return xnor(true, compareResult)

	})
}

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
