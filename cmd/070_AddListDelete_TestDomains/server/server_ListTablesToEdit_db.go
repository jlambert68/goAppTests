package server

import (
	"context"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func ListTablesToEditInDB() ([]api.ListTableToEdit, error) {
	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT tabletoedit.id, tabletoedit.guid, tabletoedit.table_name "
	sqlToExecute = sqlToExecute + "FROM tabletoedit "
	sqlToExecute = sqlToExecute + "WHERE tabletoedit.valid_for_use = true"

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
