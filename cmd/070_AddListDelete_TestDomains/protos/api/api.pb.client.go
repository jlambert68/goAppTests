package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CallApiSearch(input SearchRequest) (*Instances, error) {

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api.Api/SearchInDB", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := Instances{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}

func CallApiGetTime(input EmptyParameter) (*TimeMessage, error) {

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api.Api/GetTime", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := TimeMessage{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}

func CallApiGetMagicTableMetadata(input MagicTableMetadataRequest) (*MagicTableMetadataRespons, error) {

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api.Api/GetMagicTableMetadata", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := MagicTableMetadataRespons{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}

func CallApiListTestDomains(input EmptyParameter) (*ListTestDomainsRespons, error) {

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api.Api/ListTestDomains", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := ListTestDomainsRespons{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}

func CallApiSaveNewOrUpdateTestDomain(input NewOrUpdateTestDomainRequest) (*NewOrUpdateTestDomainResponse, error) {

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api.Api/SaveNewOrUpdateTestDomain", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := NewOrUpdateTestDomainResponse{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}

func CallApiListTestInstructions(input EmptyParameter) (*ListTestInstructionsRespons, error) {

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api.Api/ListTestInstructions", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := ListTestInstructionsRespons{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}

func CallApiSaveNewOrUpdateTestInstruction(input NewOrUpdateTestInstructionRequest) (*NewOrUpdateTestInstructionResponse, error) {

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api.Api/SaveNewOrUpdateTestInstruction", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := NewOrUpdateTestInstructionResponse{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}

func CallApiListTablesToEdit(input EmptyParameter) (*ListTablesToEditRespons, error) {

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api.Api/ListTablesToEdit", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := ListTablesToEditRespons{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}
