package main

import (
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (mt *MagicTable) SaveNewOrUpdateTestDomain(newOrUpdateTestDomainData *api.NewOrUpdateTestDomainData, newOrUpdateTestDomainUpdateType api.NewOrUpdateTestDomainUpdateType) {

	var newOrUpdateTestDomainUpdateTypeString string

	mt.logger.WithFields(logrus.Fields{
		"Id":          "87ea65e4-526b-4ed6-8c19-e6f46fea2ffb",
		"NewOrUpdate": newOrUpdateTestDomainUpdateType,
	}).Debug("Entering: SaveNewOrUpdateTestDomain()")

	defer func() {
		mt.logger.WithFields(logrus.Fields{
			"Id":                                    "0fe3ef6b-82aa-43ec-8ff9-ee4c110da78d",
			"newOrUpdateTestDomainUpdateTypeString": newOrUpdateTestDomainUpdateTypeString,
		}).Debug("Exiting: SaveNewOrUpdateTestDomain()")
	}()

	// Unselect rows
	//mt.rowSelected = -1
	newOrUpdateTestDomainUpdateTypeString = ""

	switch newOrUpdateTestDomainUpdateType {
	case api.NewOrUpdateTestDomainUpdateType_NewTestDomain:
		newOrUpdateTestDomainUpdateTypeString = "NewOrUpdateTestDomainUpdateType_NewTestDomain"

	case api.NewOrUpdateTestDomainUpdateType_UpdateTestDomain:
		newOrUpdateTestDomainUpdateTypeString = "NewOrUpdateTestDomainUpdateType_UpdateTestDomain"

	default:
		newOrUpdateTestDomainUpdateTypeString = "Unknown api.NewOrUpdateTestDomainUpdateType"
		mt.logger.WithFields(logrus.Fields{
			"Id":          "a285ca02-1df3-4f37-a012-d50a5ba065a5",
			"NewOrUpdate": newOrUpdateTestDomainUpdateType,
		}).Panic("Unknown api.NewOrUpdateTestDomainUpdateType")

	}
	//newOrUpdateTestDomainResponse
	_, err := api.CallApiSaveNewOrUpdateTestDomain(api.NewOrUpdateTestDomainRequest{
		NewOrUpdate:               api.NewOrUpdateTestDomainUpdateType_NewTestDomain,
		NewOrUpdateTestDomainData: newOrUpdateTestDomainData,
	})

	if err != nil {
		mt.logger.WithFields(logrus.Fields{
			"Id":  "c348f559-b704-4269-a7ac-24a825586bdf",
			"err": err,
		}).Panic("Something got wrong when calling database")
	}

}
