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
