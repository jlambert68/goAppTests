package main

import (
	"fmt"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (h *MagicTable) RetrieveTableDataFromDB(q string) {

	// Unselect rows
	h.rowSelected = -1

	h.SearchInDB(q)
	//instances := h.SearchInDB(q)
	//h.testDataAndMetaData.originalTestdataInstances = instances

	//h.convertInstancesIntoStandardMagicTable()
	//fmt.Println(h.magicTableRowsData)

	h.Update()

}

func (h *MagicTable) SearchInDB(q string) {
	//var err error
	switch h.tableTypeGuid {

	// Original Test table
	case "51253aba-41a9-42ef-b5f1-d8d1d7116b47":
		originalInstances, err := api.CallApiSearch(api.SearchRequest{
			Query: q,
		})

		if err != nil {
			fmt.Println("SearchInDB Error:", err)
			h.testDataAndMetaData.originalTestdataInstances = originalTestdataInstancesType{}
			return
		}

		h.testDataAndMetaData.originalTestdataInstances = originalInstances.Instances
		return

		// TestDomains
	case "8acacaaf-676e-4b36-abe6-c5310822ade1":
		testDomainInstances, err := api.CallApiListTestDomains(api.EmptyParameter{})

		if err != nil {
			fmt.Println("SearchInDB Error:", err)
			h.testDataAndMetaData.testDomains = testDomainsType{}
			return
		}

		h.testDataAndMetaData.testDomains = testDomainInstances.TestDomainForListing
		return

	// TestInstructions
	case "81c5d008-a38a-4c47-936a-d6c3c258ae13":
		testInstructionsInstances, err := api.CallApiListTestInstructions(api.EmptyParameter{})

		if err != nil {
			fmt.Println("SearchInDB Error:", err)
			h.testDataAndMetaData.testInstructions = testInstructionsType{}
			return
		}

		h.testDataAndMetaData.testInstructions = testInstructionsInstances.MyListTestInstructionsRespons
		return

		// Unknow Table type
	default:
		fmt.Println("Unknow table type: ", h.tableTypeGuid)
		return
	}

}
