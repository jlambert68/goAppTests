package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
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

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT id, guid, name, description, ready_for_use, activated, deleted, update_timestamp "
	sqlToExecute = sqlToExecute + "FROM testdomains "
	sqlToExecute = sqlToExecute + "WHERE deleted = false "
	sqlToExecute = sqlToExecute + "ORDER BY id asc "

	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var testDomainForListingMessage api.TestDomainForListingMessage
	var returnMessage []api.TestDomainForListingMessage
	var myTimeStamp time.Time

	for rows.Next() {
		err := rows.Scan(&testDomainForListingMessage.Id, &testDomainForListingMessage.Guid, &testDomainForListingMessage.Name, &testDomainForListingMessage.Description, &testDomainForListingMessage.ReadyForUse, &testDomainForListingMessage.Activated, &testDomainForListingMessage.Deleted, myTimeStamp)
		if err != nil {
			return returnMessage, err
		}
		testDomainForListingMessage.UpdateTimestamp = myTimeStamp.String()
		returnMessage = append(returnMessage, testDomainForListingMessage)

	}

	return returnMessage, rows.Err()
}
