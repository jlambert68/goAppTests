package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
)

func (mt *MagicTable) SaveNewOrUpdateTestDomain(newOrUpdateTestDomainDataKeyValueMap map[string]interface{}, newOrUpdateTestDomainUpdateType api.NewOrUpdateTestDomainUpdateType) {

	var newOrUpdateTestDomainUpdateTypeString string
	var newOrUpdateTestDomainData api.NewOrUpdateTestDomainData

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

	for key, value := range newOrUpdateTestDomainDataKeyValueMap {
		fmt.Println("key, value:", key, value)
		pointToStructNewOrUpdateTestDomainData := reflect.ValueOf(&newOrUpdateTestDomainData)
		// struct
		myStruct := pointToStructNewOrUpdateTestDomainData.Elem()
		if myStruct.Kind() == reflect.Struct {
			// exported field
			structField := myStruct.FieldByName(key)
			if structField.IsValid() {
				// A Value can be changed only if it is
				// addressable and was not obtained by
				// the use of unexported struct fields.
				if structField.CanSet() {
					// change value of N
					switch structField.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						intValue := reflect.ValueOf(value)
						/*
							intValue, err := strconv.Atoi(value)
							if err != nil {
								mt.logger.WithFields(logrus.Fields{
									"Id":          "2e811f13-bd0f-4dc1-9d17-40b396a5e76a",
									"err.Error()": err.Error(),
								}).Panic("Error when converting into Integer")
							}
						*/

						if !structField.OverflowInt(int64(intValue)) {
							structField.SetInt(int64(intValue))
						}

					case reflect.String:
						stringValue := reflect.ValueOf(value)
						structField.SetString(stringValue.String())

					case reflect.Bool:
						boolValue := reflect.ValueOf(value)
						/*
							boolValue, err := strconv.ParseBool(value)
							if err != nil {
								mt.logger.WithFields(logrus.Fields{
									"Id":          "400b4d21-2b85-47e5-af6f-68201c252814",
									"err.Error()": err.Error(),
								}).Panic("Error when converting into Boolean")
							}

						*/
						structField.SetBool(boolValue)

					//TODO case reflect.Float64:

					default:
						mt.logger.WithFields(logrus.Fields{
							"Id":                 "ccff9b3d-d733-465c-b026-a485bc261377",
							"structField.Kind()": structField.Kind(),
						}).Panic("Unknown 'structField.Kind()'")
					}

				}
			}
		}
	}

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
		NewOrUpdate:               newOrUpdateTestDomainUpdateType,
		NewOrUpdateTestDomainData: &newOrUpdateTestDomainData,
	})

	if err != nil {
		mt.logger.WithFields(logrus.Fields{
			"Id":  "c348f559-b704-4269-a7ac-24a825586bdf",
			"err": err,
		}).Panic("Something got wrong when calling database")
	}

}
