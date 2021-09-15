package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) ListTestInstructions(ctx context.Context, in *api.EmptyParameter) (*api.ListTestInstructionsRespons, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "830d4bd0-4aef-4cfe-94c6-239de5f1ad12",
		"Trace": server.trace(false),
	}).Debug("Entering: ListTestInstructions()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "b21e5aaa-f5e3-465e-8a65-f3f246d367ce",
			"Trace": server.trace(false),
		}).Debug("Exiting: ListTestInstructions()")
	}()

	var testInstructionsRespons *api.ListTestInstructionsRespons
	var myListTestInstructionsRespons []*api.ListTestInstructionRespons

	testInstructionsFromDB, err := server.ListTestInstructionsInDB()
	if err != nil {
		fmt.Println(err.Error())
		return testInstructionsRespons, err
	}

	//fmt.Println(testInstructionsFromDB)

	for _, testInstructionFromDB := range testInstructionsFromDB {

		testInstruction := &api.ListTestInstructionRespons{
			Id:              testInstructionFromDB.Id,
			Guid:            testInstructionFromDB.Guid,
			Name:            testInstructionFromDB.Name,
			Description:     testInstructionFromDB.Description,
			ReadyForUse:     testInstructionFromDB.ReadyForUse,
			Activated:       testInstructionFromDB.Activated,
			Deleted:         testInstructionFromDB.Deleted,
			UpdateTimestamp: testInstructionFromDB.UpdateTimestamp,
		}

		myListTestInstructionsRespons = append(myListTestInstructionsRespons, testInstruction)

	}

	testInstructionsRespons = &api.ListTestInstructionsRespons{
		MyListTestInstructionsRespons: myListTestInstructionsRespons,
	}

	return testInstructionsRespons, nil

	/*
		var err error
		var returnMessage *api.ListTestInstructionsRespons
		err = nil

		a := &api.ListTestInstructionRespons{
			Id:                   0,
			Guid:                 "12345",
			Name:                 "Custody Cash",
			Description:          "Hanteras alla tester som har med Custody Cash att göra",
			ReadyForUse:          false,
			Activated:            false,
			Deleted:              false,
			UpdateTimestamp:      "2021-08-31",
		}

		returnMessage.MyListTestInstructionsRespons = append(returnMessage.MyListTestInstructionsRespons, a)

		a = &api.ListTestInstructionRespons{
			Id:                   0,
			Guid:                 "986765",
			Name:                 "Custody Arrangement",
			Description:          "Hanteras alla tester som har med Custody Arrangement att göra",
			ReadyForUse:          false,
			Activated:            false,
			Deleted:              false,
			UpdateTimestamp:      "2021-08-30",
		}

		returnMessage.MyListTestInstructionsRespons = append(returnMessage.MyListTestInstructionsRespons, a)



		return returnMessage, err

	*/
}

func (server *Server) SaveNewOrUpdateTestInstruction(ctx context.Context, in *api.NewOrUpdateTestInstructionRequest) (*api.NewOrUpdateTestInstructionResponse, error) {

	var err error
	var returnMessage *api.NewOrUpdateTestInstructionResponse
	err = nil

	switch in.NewOrUpdate {
	case api.NewOrUpdateTestInstructionUpdateType_NewTestInstruction:
		returnMessage = &api.NewOrUpdateTestInstructionResponse{
			Id:              99,
			Guid:            "65555555",
			ResponseStatus:  true,
			ResponseMessage: "Message was saved in database",
		}

	case api.NewOrUpdateTestInstructionUpdateType_UpdateInstruction:
		returnMessage = &api.NewOrUpdateTestInstructionResponse{
			Id:              in.NewOrUpdateTestInstructionData.Id,
			Guid:            in.NewOrUpdateTestInstructionData.Guid,
			ResponseStatus:  true,
			ResponseMessage: "Message was updated in database",
		}

	default:
		returnMessage = &api.NewOrUpdateTestInstructionResponse{
			Id:              in.NewOrUpdateTestInstructionData.Id,
			Guid:            in.NewOrUpdateTestInstructionData.Guid,
			ResponseStatus:  false,
			ResponseMessage: "Unknown api.NewOrUpdateTestInstructionUpdateType",
		}
	}

	return returnMessage, err
}
