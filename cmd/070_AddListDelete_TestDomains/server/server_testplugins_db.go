package server

import (
	"context"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func ListTestInstructionsInDB() ([]api.ListTestInstructionRespons, error) {
	sqlToExecute := "SELECT * FROM testinstructions;"
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var ListTestInstructionRespons api.ListTestInstructionRespons
	var returnMessage []api.ListTestInstructionRespons

	for rows.Next() {
		err := rows.Scan(&ListTestInstructionRespons.Id, &ListTestInstructionRespons.Guid, &ListTestInstructionRespons.Name, &ListTestInstructionRespons.Description, &ListTestInstructionRespons.ReadyForUse, &ListTestInstructionRespons.Activated, &ListTestInstructionRespons.Deleted, &ListTestInstructionRespons.UpdateTimestamp)
		if err != nil {
			return returnMessage, err
		}
		returnMessage = append(returnMessage, ListTestInstructionRespons)
		//fmt.Println(ListTestInstructionRespons)
	}

	return returnMessage, rows.Err()
}
