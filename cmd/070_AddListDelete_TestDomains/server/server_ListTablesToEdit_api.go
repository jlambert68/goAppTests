package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) ListTablesToEdit(ctx context.Context, in *api.EmptyParameter) (*api.ListTablesToEditRespons, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "e0de54a8-a031-49ba-aa0f-a9b4f7d7f9c9",
		"Trace": server.trace(false),
	}).Debug("Entering: ListTablesToEdit()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "0f719af4-309f-4fba-bc5b-f6efaecbd282",
			"Trace": server.trace(false),
		}).Debug("Exiting: ListTablesToEdit()")
	}()

	var listTablesToEditRespons *api.ListTablesToEditRespons
	var myListTablesToEdit []*api.ListTableToEdit

	listTablesToEditInDB, err := server.ListTablesToEditInDB()
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
