package server

import (
	"context"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"time"
)

func ListTestDomainsInDB() ([]api.TestDomainForListingMessage, error) {
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
