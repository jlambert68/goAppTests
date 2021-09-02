package server

import (
	"context"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func ListTestInstructionsInDB() ([]api.ListTestInstructionRespons, error) {
	sqlToExecute := "SELECT * FROM testinstructions;"
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var ListTestPluginRespons api.ListTestInstructionRespons
	var returnMessage []api.ListTestInstructionRespons

	for rows.Next() {
		err := rows.Scan(&ListTestPluginRespons.Id, &ListTestPluginRespons.Guid, &ListTestPluginRespons.Name, &ListTestPluginRespons.Description, &ListTestPluginRespons.ReadyForUse, &ListTestPluginRespons.Activated, &ListTestPluginRespons.Deleted, &ListTestPluginRespons.UpdateTimestamp)
		if err != nil {
			return returnMessage, err
		}
		returnMessage = append(returnMessage, ListTestPluginRespons)
		//fmt.Println(ListTestPluginRespons)
	}

	return returnMessage, rows.Err()
}
