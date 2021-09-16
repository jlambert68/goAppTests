package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (mt *MagicTable) RetrieveTableDataFromDB(q string) {

	// Unselect rows
	//mt.rowSelected = -1

	mt.SearchInDB(q)
	//instances := mt.SearchInDB(q)
	//mt.testDataAndMetaData.originalTestdataInstances = instances

	//mt.convertInstancesIntoStandardMagicTable()
	//fmt.Println(mt.magicTableRowsData)

	mt.Update()

}

func (mt *MagicTable) SearchInDB(q string) {

	// Unselect rows
	mt.rowSelected = -1

	mt.logger.WithFields(logrus.Fields{
		"Id": "843892ce-d69f-407d-8436-af2f6307b9c6",
	}).Debug("Entering: SearchInDB()")

	defer func() {
		mt.logger.WithFields(logrus.Fields{
			"Id": "19b11793-4986-4c5e-825c-0bf0755dcbd9",
		}).Debug("Exiting: SearchInDB()")
	}()

	//var err error
	switch mt.tableTypeGuid {

	// Original Test table
	case "51253aba-41a9-42ef-b5f1-d8d1d7116b47":
		originalInstances, err := api.CallApiSearch(api.SearchRequest{
			Query: q,
		})

		if err != nil {
			fmt.Println("SearchInDB Error:", err)
			mt.testDataAndMetaData.originalTestdataInstances = originalTestdataInstancesType{}
			return
		}

		mt.testDataAndMetaData.originalTestdataInstances = originalInstances.Instances
		fmt.Println("originalInstances.Instances: ", len(originalInstances.Instances))
		return

		// TestDomains
	case "8acacaaf-676e-4b36-abe6-c5310822ade1":
		testDomainInstances, err := api.CallApiListTestDomains(api.EmptyParameter{})

		if err != nil {
			fmt.Println("SearchInDB Error:", err)
			mt.testDataAndMetaData.testDomains = testDomainsType{}
			return
		}

		mt.testDataAndMetaData.testDomains = testDomainInstances.TestDomainForListing
		//fmt.Println("testDomainInstances.TestDomainForListing: ", len(testDomainInstances.TestDomainForListing))
		return

	// TestInstructions
	case "81c5d008-a38a-4c47-936a-d6c3c258ae13":
		testInstructionsInstances, err := api.CallApiListTestInstructions(api.EmptyParameter{})

		if err != nil {
			fmt.Println("SearchInDB Error:", err)
			mt.testDataAndMetaData.testInstructions = testInstructionsType{}
			return
		}

		mt.testDataAndMetaData.testInstructions = testInstructionsInstances.MyListTestInstructionsRespons
		return

		// Unknow Table type
	default:
		fmt.Println("Unknown table type2: ", mt.tableTypeGuid)
		return
	}

}
