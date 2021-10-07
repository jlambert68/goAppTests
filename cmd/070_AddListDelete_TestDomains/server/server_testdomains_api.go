package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) ListTestDomains(ctx context.Context, in *api.EmptyParameter) (*api.ListTestDomainsRespons, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "1a0dc1f1-5ac4-4475-b28e-6659885461d0",
		"Trace": server.trace(false),
	}).Debug("Entering: ListTestDomains()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "76b570f8-8ae9-44e8-b17a-550a37c7965d",
			"Trace": server.trace(false),
		}).Debug("Exiting: ListTestDomains()")
	}()

	var testDomainsRespons *api.ListTestDomainsRespons
	var testDomainForListingMessages []*api.TestDomainForListingMessage

	testDomainsFromDB, err := server.ListTestDomainsInDB()
	if err != nil {
		server.logger.WithFields(logrus.Fields{
			"Id": "7c3ec670-105a-4a6b-9d34-4a5694ff194a",
		}).Error(err.Error())
		return testDomainsRespons, err
	}

	//fmt.Println(testDomainsFromDB)

	for rowNumber, testDomainFromDB := range testDomainsFromDB {

		testDomain := &api.TestDomainForListingMessage{
			Id:              testDomainFromDB.Id,
			Guid:            testDomainFromDB.Guid,
			Name:            testDomainFromDB.Name,
			Description:     testDomainFromDB.Description,
			ReadyForUse:     testDomainFromDB.ReadyForUse,
			Activated:       testDomainFromDB.Activated,
			Deleted:         testDomainFromDB.Deleted,
			UpdateTimestamp: testDomainFromDB.UpdateTimestamp,
			UniqueId:        int64(rowNumber),
		}

		testDomainForListingMessages = append(testDomainForListingMessages, testDomain)

	}

	testDomainsRespons = &api.ListTestDomainsRespons{
		TestDomainForListing: testDomainForListingMessages,
	}

	return testDomainsRespons, nil

	/*
		var err error
		var returnMessage *api.ListTestDomainsRespons
		err = nil

		a := &api.ListTestDomainRespons{
			Id:                   0,
			Guid:                 "12345",
			Name:                 "Custody Cash",
			Description:          "Hanteras alla tester som har med Custody Cash att göra",
			ReadyForUse:          false,
			Activated:            false,
			Deleted:              false,
			UpdateTimestamp:      "2021-08-31",
		}

		returnMessage.MyListTestDomainsRespons = append(returnMessage.MyListTestDomainsRespons, a)

		a = &api.ListTestDomainRespons{
			Id:                   0,
			Guid:                 "986765",
			Name:                 "Custody Arrangement",
			Description:          "Hanteras alla tester som har med Custody Arrangement att göra",
			ReadyForUse:          false,
			Activated:            false,
			Deleted:              false,
			UpdateTimestamp:      "2021-08-30",
		}

		returnMessage.MyListTestDomainsRespons = append(returnMessage.MyListTestDomainsRespons, a)



		return returnMessage, err

	*/
}

func (server *Server) SaveNewOrUpdateTestDomain(ctx context.Context, newOrUpdateTestDomainRequest *api.NewOrUpdateTestDomainRequest) (*api.NewOrUpdateTestDomainResponse, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "fd690fdb-69b4-4fae-887f-03fb10d40db7",
		"Trace": server.trace(false),
	}).Debug("Entering: SaveNewOrUpdateTestDomain()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "c987dbef-c5fb-4e55-9de0-0f97ab51366d",
			"Trace": server.trace(false),
		}).Debug("Exiting: SaveNewOrUpdateTestDomain()")
	}()

	var err error
	var returnMessage *api.NewOrUpdateTestDomainResponse
	err = nil

	switch newOrUpdateTestDomainRequest.NewOrUpdate {
	case api.NewOrUpdateTestDomainUpdateType_NewTestDomain:

		returnMessage = &api.NewOrUpdateTestDomainResponse{
			NewOrUpdateTestDomainData: nil,
			ResponseStatus:            true,
			ResponseMessage:           "Message was saved newOrUpdateTestDomainRequest database",
		}

	case api.NewOrUpdateTestDomainUpdateType_UpdateTestDomain:
		returnMessage = &api.NewOrUpdateTestDomainResponse{
			NewOrUpdateTestDomainData: nil,
			ResponseStatus:            true,
			ResponseMessage:           "Message was updated newOrUpdateTestDomainRequest database",
		}

	default:
		returnMessage = &api.NewOrUpdateTestDomainResponse{
			NewOrUpdateTestDomainData: nil,
			ResponseStatus:            false,
			ResponseMessage:           "Unknown api.NewOrUpdateTestDomainUpdateType",
		}
	}

	newOrUpdateTestDomainData, err := server.SaveNewOrUpdateTestDomainDB(newOrUpdateTestDomainRequest)
	if err != nil {
		returnMessage = &api.NewOrUpdateTestDomainResponse{
			NewOrUpdateTestDomainData: nil,
			ResponseStatus:            false,
			ResponseMessage:           err.Error(),
		}

		server.logger.WithFields(logrus.Fields{
			"Id":          "9ff78f70-8237-4b0f-97c0-e197f6ef4149",
			"err.Error()": err.Error(),
		}).Error("Error when calling database")

		return returnMessage, nil
	}

	returnMessage = &api.NewOrUpdateTestDomainResponse{
		NewOrUpdateTestDomainData: &newOrUpdateTestDomainData,
		ResponseStatus:            true,
		ResponseMessage:           "Message was updated newOrUpdateTestDomainRequest database",
	}

	server.logger.WithFields(logrus.Fields{
		"Id":                     "3c4b328f-6923-49eb-80a0-9c83cf917f28",
		"Message to Save/Update": newOrUpdateTestDomainRequest,
		"New Message":            newOrUpdateTestDomainData,
	}).Debug("Message was Saved/Updated in database")

	return returnMessage, err
}
