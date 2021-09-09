package server

import (
	"context"
	"fmt"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) ListTablesToEdit(ctx context.Context, in *api.EmptyParameter) (*api.ListTablesToEditRespons, error) {

	var listTablesToEditRespons *api.ListTablesToEditRespons
	var myListTablesToEdit []*api.ListTableToEdit

	listTablesToEditInDB, err := ListTablesToEditInDB()
	if err != nil {
		fmt.Println(err.Error())
		return listTablesToEditRespons, err
	}

	//fmt.Println(testInstructionsFromDB)

	for _, listTableToEditInDB := range listTablesToEditInDB {

		listTableToEdit := &api.ListTableToEdit{
			Id:        listTableToEditInDB.Id,
			Guid:      listTableToEditInDB.Guid,
			TableName: listTableToEditInDB.TableName,
		}

		myListTablesToEdit = append(myListTablesToEdit, listTableToEdit)

	}

	listTablesToEditRespons = &api.ListTablesToEditRespons{
		MyListTableToEdit: myListTablesToEdit,
	}

	return listTablesToEditRespons, nil

}
