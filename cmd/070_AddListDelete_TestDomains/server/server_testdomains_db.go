package server

import (
	"context"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func ListTestDomainsInDB() ([]api.TestDomainForListingMessage, error) {
	sqlToExecute := "SELECT * FROM testdomains;"
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var testDomainForListingMessage api.TestDomainForListingMessage
	var returnMessage []api.TestDomainForListingMessage

	for rows.Next() {
		err := rows.Scan(&testDomainForListingMessage.Id, &testDomainForListingMessage.Guid, &testDomainForListingMessage.Name, &testDomainForListingMessage.Description, &testDomainForListingMessage.ReadyForUse, &testDomainForListingMessage.Activated, &testDomainForListingMessage.Deleted, &testDomainForListingMessage.UpdateTimestamp)
		if err != nil {
			return returnMessage, err
		}
		returnMessage = append(returnMessage, testDomainForListingMessage)

	}

	return returnMessage, rows.Err()
}
