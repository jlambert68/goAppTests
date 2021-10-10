package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"reflect"
	"strconv"
)

type keyValueMapType map[string]string

func (mt *MagicTable) SaveNewOrUpdateTestDomain(newOrUpdateTestDomainDataKeyValueMap keyValueMapType, newOrUpdateTestDomainUpdateType api.NewOrUpdateTestDomainUpdateType) {

	var newOrUpdateTestDomainUpdateTypeString string
	var newOrUpdateTestDomainData api.NewOrUpdateTestDomainData

	mt.logger.WithFields(logrus.Fields{
		"Id":          "87ea65e4-526b-4ed6-8c19-e6f46fea2ffb",
		"NewOrUpdate": newOrUpdateTestDomainUpdateType,
	}).Debug("Entering: DeleteTestDomain()")

	defer func() {
		mt.logger.WithFields(logrus.Fields{
			"Id":                                    "0fe3ef6b-82aa-43ec-8ff9-ee4c110da78d",
			"newOrUpdateTestDomainUpdateTypeString": newOrUpdateTestDomainUpdateTypeString,
		}).Debug("Exiting: DeleteTestDomain()")
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
						fmt.Println(reflect.ValueOf(value).Kind())
						//intValue := value.(int) //reflect.ValueOf(value).Int()

						// Only try to reflect if there are a value
						if len(value) != 0 {
							intValue, err := strconv.Atoi(value)
							if err != nil {
								mt.logger.WithFields(logrus.Fields{
									"Id":              "2e811f13-bd0f-4dc1-9d17-40b396a5e76a",
									"err.Error()":     err.Error(),
									"key to parse":    key,
									"Value to  parse": value,
								}).Panic("Error when converting into Integer")
							}

							if !structField.OverflowInt(int64(intValue)) {
								structField.SetInt(int64(intValue))
							}
						} else {
							structField.SetInt(int64(0))
						}

					case reflect.String:
						stringValue := reflect.ValueOf(value)
						structField.SetString(stringValue.String())

					case reflect.Bool:

						//boolValue := value.(bool)  //reflect.ValueOf(value).String()
						//fmt.Println(" reflect.Bool", value, reflect.ValueOf(value).Kind(), boolValue)

						// Only try to reflect if there are a value
						if len(value) != 0 {
							boolValue, err := strconv.ParseBool(value)
							if err != nil {
								mt.logger.WithFields(logrus.Fields{
									"Id":              "400b4d21-2b85-47e5-af6f-68201c252814",
									"err.Error()":     err.Error(),
									"key to parse":    key,
									"Value to  parse": value,
									"lenght of bool":  len(value),
								}).Panic("Error when converting into Boolean")
							}

							structField.SetBool(boolValue)

						} else {
							structField.SetBool(false)
						}

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

	// Update list
	mt.SearchInDB(mt.searchString)

}

func (mt *MagicTable) DeleteTestDomain(deleteTestDomainDataKeyValueMap keyValueMapType) {

	var deleteTestDomainRequest api.DeleteTestDomainRequest

	mt.logger.WithFields(logrus.Fields{
		"Id": "72f1e015-0d1d-4eb1-bad9-e4d0bb6c050d",
	}).Debug("Entering: DeleteTestDomain()")

	defer func() {
		mt.logger.WithFields(logrus.Fields{
			"Id": "b3d05bc7-bbff-4e38-b727-9c097d1fa7a4",
		}).Debug("Exiting: DeleteTestDomain()")
	}()

	for key, value := range deleteTestDomainDataKeyValueMap {

		// Only keep value for "guid"
		if key != "Guid" {
			continue
		}

		pointToStructDeleteTestDomainData := reflect.ValueOf(&deleteTestDomainRequest)
		// struct
		myStruct := pointToStructDeleteTestDomainData.Elem()
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

					case reflect.String:
						stringValue := reflect.ValueOf(value)
						structField.SetString(stringValue.String())

					default:
						mt.logger.WithFields(logrus.Fields{
							"Id":                 "357306d9-e0ae-4357-b6b8-0bc73d1d8eb7",
							"structField.Kind()": structField.Kind(),
						}).Panic("Unknown 'structField.Kind()'")
					}

				}
			}
		}
	}

	// Unselect rows
	//mt.rowSelected = -1

	//Call database for deleteing
	_, err := api.CallApiDeleteTestDomain(deleteTestDomainRequest)

	if err != nil {
		mt.logger.WithFields(logrus.Fields{
			"Id":  "a7b1911b-754e-4186-85ac-48469037cc5b",
			"err": err,
		}).Panic("Something got wrong when calling database")
	}

	// Update list
	mt.SearchInDB(mt.searchString)

}
