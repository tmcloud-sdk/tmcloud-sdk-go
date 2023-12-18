package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostPaidServerSchedulerHints struct {
	Group *string `json:"group,omitempty"`

	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	Tenancy *string `json:"tenancy,omitempty"`
}

func (o PostPaidServerSchedulerHints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerSchedulerHints struct{}"
	}

	return strings.Join([]string{"PostPaidServerSchedulerHints", string(data)}, " ")
}
