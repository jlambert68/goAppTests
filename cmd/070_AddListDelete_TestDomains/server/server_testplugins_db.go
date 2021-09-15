package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) ListTestInstructionsInDB() ([]api.ListTestInstructionRespons, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "73768ffb-e069-42cf-945f-009a332a9ac6",
		"Trace": server.trace(false),
	}).Debug("Entering: ListTestInstructionsInDB()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "07e53e86-e45f-4799-a4cf-4d5c73872fbe",
			"Trace": server.trace(false),
		}).Debug("Exiting: ListTestInstructionsInDB()")
	}()

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
