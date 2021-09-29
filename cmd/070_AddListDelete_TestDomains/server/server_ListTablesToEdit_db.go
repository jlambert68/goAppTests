package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) ListTablesToEditInDB() ([]api.ListTableToEdit, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "8555f387-d44d-4d8a-ae4f-0382c20a835a",
		"Trace": server.trace(false),
	}).Debug("Entering: ListTablesToEditInDB()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "1d4a3138-21da-43b4-828b-2a525c75878f",
			"Trace": server.trace(false),
		}).Debug("Exiting: ListTablesToEditInDB()")
	}()

	sqlToExecute := "SELECT * FROM sp_listtablestoedit()"
	/*
		sqlToExecute = sqlToExecute + "SELECT tabletoedit.id, tabletoedit.guid, tabletoedit.table_name "
		sqlToExecute = sqlToExecute + "FROM tabletoedit "
		sqlToExecute = sqlToExecute + "WHERE tabletoedit.valid_for_use = true"
	*/
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var listTableToEdit api.ListTableToEdit
	var returnMessage []api.ListTableToEdit

	for rows.Next() {
		err := rows.Scan(&listTableToEdit.Id, &listTableToEdit.Guid, &listTableToEdit.TableName)
		if err != nil {
			return returnMessage, err
		}
		returnMessage = append(returnMessage, listTableToEdit)
		//fmt.Println(listTableToEdit)
	}

	return returnMessage, rows.Err()
}
