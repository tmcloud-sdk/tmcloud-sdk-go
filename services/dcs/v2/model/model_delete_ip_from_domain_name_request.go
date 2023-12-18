package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteIpFromDomainNameRequest struct {
	InstanceId string `json:"instance_id"`

	GroupId string `json:"group_id"`

	NodeId string `json:"node_id"`
}

func (o DeleteIpFromDomainNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpFromDomainNameRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpFromDomainNameRequest", string(data)}, " ")
}
