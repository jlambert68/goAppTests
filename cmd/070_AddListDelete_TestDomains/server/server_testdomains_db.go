package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"strconv"
	"time"
)

func (server *Server) ListTestDomainsInDB() ([]api.TestDomainForListingMessage, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "7cf597af-c917-498e-869c-b786a47791b3",
		"Trace": server.trace(false),
	}).Debug("Entering: ListTestDomainsInDB()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "02050f77-6a66-4906-83ab-aab71e4748d4",
			"Trace": server.trace(false),
		}).Debug("Exiting: ListTestDomainsInDB()")
	}()

	sqlToExecute := "SELECT * FROM sp_listtestdomains()"
	/*
		sqlToExecute = sqlToExecute + "SELECT id, guid, name, description, ready_for_use, activated, deleted, update_timestamp "
		sqlToExecute = sqlToExecute + "FROM testdomains "
		sqlToExecute = sqlToExecute + "WHERE deleted = false "
		sqlToExecute = sqlToExecute + "ORDER BY id asc "


	*/
	//rows, _ := DbPool.Query(context.Background(), sqlToExecute)
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var testDomainForListingMessage api.TestDomainForListingMessage
	var returnMessage []api.TestDomainForListingMessage
	var myTimeStamp interface{} //time.Time

	for rows.Next() {

		err := rows.Scan(
			&testDomainForListingMessage.Id,
			&testDomainForListingMessage.Guid,
			&testDomainForListingMessage.Name,
			&testDomainForListingMessage.Description,
			&testDomainForListingMessage.ReadyForUse,
			&testDomainForListingMessage.Activated,
			&testDomainForListingMessage.Deleted,
			myTimeStamp)
		if err != nil {
			return returnMessage, err
		}
		testDomainForListingMessage.UpdateTimestamp = fmt.Sprintf("%v", myTimeStamp) // myTimeStamp.String()
		returnMessage = append(returnMessage, testDomainForListingMessage)

		//fmt.Println("XXXXXXXXX testDomainForListingMessage: ", testDomainForListingMessage)
	}

	return returnMessage, rows.Err()
}

func (server *Server) SaveNewOrUpdateTestDomainDB(newOrUpdateTestDomainRequest *api.NewOrUpdateTestDomainRequest) (api.NewOrUpdateTestDomainData, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "55117ded-c1cf-4a3c-88f6-00660881be93",
		"Trace": server.trace(false),
	}).Debug("Entering: SaveNewOrUpdateTestDomainDB()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "ef6a2a32-d866-4567-b23c-75a7e0bac77b",
			"Trace": server.trace(false),
		}).Debug("Exiting: SaveNewOrUpdateTestDomainDB()")
	}()

	var currentTimeStamp time.Time
	currentTimeStamp = time.Now()

	sqlToExecute := "Select * From sp_insert_new_or_updated_testdomain("
	sqlToExecute = sqlToExecute + "'" + newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Guid + "', "
	sqlToExecute = sqlToExecute + "'" + newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Name + "', "
	sqlToExecute = sqlToExecute + "'" + newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Description + "', "
	sqlToExecute = sqlToExecute + "'" + strconv.FormatBool(newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.ReadyForUse) + "', "
	sqlToExecute = sqlToExecute + "'" + strconv.FormatBool(newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Activated) + "', "
	sqlToExecute = sqlToExecute + "'" + strconv.FormatBool(newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Deleted) + "', "
	sqlToExecute = sqlToExecute + "'" + newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.UpdateTimestamp + "') "

	rows, err := DbPool.Query(context.Background(), sqlToExecute,
		newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Guid,
		newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Name,
		newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Description,
		newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.ReadyForUse,
		newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Activated,
		newOrUpdateTestDomainRequest.NewOrUpdateTestDomainData.Deleted,
		currentTimeStamp,
	)

	if err != nil {
		server.logger.WithFields(logrus.Fields{
			"Id":                           "30d3e05e-8ef5-42b6-bee8-bc0857966901",
			"newOrUpdateTestDomainRequest": newOrUpdateTestDomainRequest,
		}).Error("Something went wrong when creating new or updating TestDomain")
	}

	var newOrUpdateTestDomainData api.NewOrUpdateTestDomainData

	var myTimeStamp interface{} //time.Time

	newOrUpdateTestDomainData = api.NewOrUpdateTestDomainData{}

	for rows.Next() {

		err = rows.Scan(&newOrUpdateTestDomainData.Id, &newOrUpdateTestDomainData.Guid, &newOrUpdateTestDomainData.Name, &newOrUpdateTestDomainData.Description, &newOrUpdateTestDomainData.ReadyForUse, &newOrUpdateTestDomainData.Activated, &newOrUpdateTestDomainData.Deleted, myTimeStamp)
		if err != nil {
			return newOrUpdateTestDomainData, err
		}
		newOrUpdateTestDomainData.UpdateTimestamp = fmt.Sprintf("%v", myTimeStamp) // myTimeStamp.String()

		break
	}

	return newOrUpdateTestDomainData, nil
}
