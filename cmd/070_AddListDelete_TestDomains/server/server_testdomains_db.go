package server

import (
	"context"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func ListTestDomainsInDB() ([]api.ListTestDomainRespons, error) {
	sqlToExecute := "SELECT * FROM testdomains;"
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var ListTestDomainRespons api.ListTestDomainRespons
	var returnMessage []api.ListTestDomainRespons

	for rows.Next() {
		err := rows.Scan(&ListTestDomainRespons.Id, &ListTestDomainRespons.Guid, &ListTestDomainRespons.Name, &ListTestDomainRespons.Description, &ListTestDomainRespons.ReadyForUse, &ListTestDomainRespons.Activated, &ListTestDomainRespons.Deleted, &ListTestDomainRespons.UpdateTimestamp)
		if err != nil {
			return returnMessage, err
		}
		returnMessage = append(returnMessage, ListTestDomainRespons)
		//fmt.Println(ListTestDomainRespons)
	}

	return returnMessage, rows.Err()
}
