package server

import (
	"context"
	"encoding/json"
	"goAppTest1/cmd/060_golangWasmMultiProto_goAppv8/protos/api"
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
	instances []*api.Instance
}

func (server *Server) Search(ctx context.Context, in *api.SearchRequest) (*api.Instances, error) {
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

	return &api.Instances{Instances: instances}, nil
}

func (server *Server) parseInstances() {
	fileName := "web/instances.json"
	ec2Instances := []ec2Instance{}
	server.instances = []*api.Instance{}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &ec2Instances)

	for _, e := range ec2Instances {
		server.instances = append(server.instances, &api.Instance{
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

	returnMessage := &api.TimeMessage{
		TimeString: callBackEnd(),
	}
	return returnMessage, nil
}

func callBackEnd() string {

	t := time.Now()
	return "JOnas_v1025  Hi from Server: " + t.String()

}
