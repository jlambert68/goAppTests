package server

import (
	"context"
	"fmt"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) ListTestDomains(ctx context.Context, in *api.EmptyParameter) (*api.ListTestDomainsRespons, error) {

	var testDomainsRespons *api.ListTestDomainsRespons
	var testDomainForListingMessages []*api.TestDomainForListingMessage

	testDomainsFromDB, err := ListTestDomainsInDB()
	if err != nil {
		fmt.Println(err.Error())
		return testDomainsRespons, err
	}

	//fmt.Println(testDomainsFromDB)

	for _, testDomainFromDB := range testDomainsFromDB {

		testDomain := &api.TestDomainForListingMessage{
			Id:              testDomainFromDB.Id,
			Guid:            testDomainFromDB.Guid,
			Name:            testDomainFromDB.Name,
			Description:     testDomainFromDB.Description,
			ReadyForUse:     testDomainFromDB.ReadyForUse,
			Activated:       testDomainFromDB.Activated,
			Deleted:         testDomainFromDB.Deleted,
			UpdateTimestamp: testDomainFromDB.UpdateTimestamp,
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

func (server *Server) SaveNewOrUpdateTestDomain(ctx context.Context, in *api.NewOrUpdateTestDomainRequest) (*api.NewOrUpdateTestDomainResponse, error) {

	var err error
	var returnMessage *api.NewOrUpdateTestDomainResponse
	err = nil

	switch in.NewOrUpdate {
	case api.NewOrUpdateTestDomainUpdateType_NewTestDomain:
		returnMessage = &api.NewOrUpdateTestDomainResponse{
			Id:              99,
			Guid:            "65555555",
			ResponseStatus:  true,
			ResponseMessage: "Message was saved in database",
		}

	case api.NewOrUpdateTestDomainUpdateType_UpdateTestDomain:
		returnMessage = &api.NewOrUpdateTestDomainResponse{
			Id:              in.NewOrUpdateTestDomainData.Id,
			Guid:            in.NewOrUpdateTestDomainData.Guid,
			ResponseStatus:  true,
			ResponseMessage: "Message was updated in database",
		}

	default:
		returnMessage = &api.NewOrUpdateTestDomainResponse{
			Id:              in.NewOrUpdateTestDomainData.Id,
			Guid:            in.NewOrUpdateTestDomainData.Guid,
			ResponseStatus:  false,
			ResponseMessage: "Unknown api.NewOrUpdateTestDomainUpdateType",
		}
	}

	return returnMessage, err
}
