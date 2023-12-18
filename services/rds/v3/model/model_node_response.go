package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NodeResponse struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Role string `json:"role"`

	Status string `json:"status"`

	AvailabilityZone string `json:"availability_zone"`
}

func (o NodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeResponse struct{}"
	}

	return strings.Join([]string{"NodeResponse", string(data)}, " ")
}
