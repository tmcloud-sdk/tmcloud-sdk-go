package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerSchedulerHints struct {
	Group *[]string `json:"group,omitempty"`

	Tenancy *[]string `json:"tenancy,omitempty"`

	DedicatedHostId *[]string `json:"dedicated_host_id,omitempty"`
}

func (o ServerSchedulerHints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerSchedulerHints struct{}"
	}

	return strings.Join([]string{"ServerSchedulerHints", string(data)}, " ")
}
