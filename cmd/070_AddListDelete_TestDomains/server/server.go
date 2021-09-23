package server

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
	"io/ioutil"
	"strings"
	"time"
)

////
// ec2Instances file
////

type ec2Instance struct {
	PrettyName   string  `json:"pretty_name,omitempty"`
	InstanceType string  `json:"instance_type,omitempty"`
	ECU          float32 `json:"ECU,omitempty"`
	Memory       float32 `json:"memory,omitempty"`

	NetworkPerformance string `json:"network_performance,omitempty"`

	Pricing map[string]map[string]struct {
		OnDemand string `json:"ondemand,omitempty"`
	} `json:"pricing,omitempty"`
}

////
// Server
////

type Server struct {
	logger            *logrus.Logger
	loggerIsInitiated bool
	instances         []*api.Instance
}

func (server *Server) Search(ctx context.Context, in *api.SearchRequest) (*api.Instances, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "fd765057-6dd8-482b-93f5-e8b115d63e2d",
		"Trace": server.trace(false),
	}).Debug("Entering: Search()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "97628301-4a94-4216-b7fb-6c24cb1950de",
			"Trace": server.trace(false),
		}).Debug("Exiting: Search()")
	}()

	if server.instances == nil {
		server.parseInstances()
	}

	instances := []*api.Instance{}
	for _, instance := range server.instances {
		str, _ := json.Marshal(*instance)
		if strings.Contains(string(str), in.Query) {
			instances = append(instances, instance)
		}
	}

	//forDebuggingSortData(instances)

	return &api.Instances{Instances: instances}, nil
}

func (server *Server) parseInstances() {

	server.logger.WithFields(logrus.Fields{
		"Id":    "95be7b0d-929e-4063-bdcc-df04d65c7cbb",
		"Trace": server.trace(false),
	}).Debug("Entering: parseInstances()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "20549d94-5f4e-4c3e-96fc-f33ec387975a",
			"Trace": server.trace(false),
		}).Debug("Exiting: parseInstances()")
	}()

	fileName := "web/instances.json"
	ec2Instances := []ec2Instance{}
	server.instances = []*api.Instance{}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &ec2Instances)

	for rowNumber, e := range ec2Instances {
		server.instances = append(server.instances, &api.Instance{

			UniqueId:     int64(rowNumber),
			Name:         e.PrettyName,
			InstanceType: e.InstanceType,
			Ecu:          e.ECU,
			Memory:       e.Memory,
			Network:      e.NetworkPerformance,
			Price:        e.Pricing["us-east-1"]["linux"].OnDemand,
		})
	}
}

func (server *Server) GetTime(ctx context.Context, in *api.EmptyParameter) (*api.TimeMessage, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "8c512b55-3e52-468b-8867-0da71ae5a6a2",
		"Trace": server.trace(false),
	}).Debug("Entering: GetTime()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "eab1de7e-183c-4599-8c21-cde9eb76a610",
			"Trace": server.trace(false),
		}).Debug("Exiting: GetTime()")
	}()

	returnMessage := &api.TimeMessage{
		TimeString: callBackEnd(),
	}
	return returnMessage, nil
}

func callBackEnd() string {

	t := time.Now()
	return "JOnas_v1025  Hi from Server: " + t.String()

}
