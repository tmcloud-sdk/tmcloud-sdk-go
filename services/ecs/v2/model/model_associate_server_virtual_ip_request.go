package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AssociateServerVirtualIpRequest struct {
	NicId string `json:"nic_id"`

	Body *AssociateServerVirtualIpRequestBody `json:"body,omitempty"`
}

func (o AssociateServerVirtualIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateServerVirtualIpRequest struct{}"
	}

	return strings.Join([]string{"AssociateServerVirtualIpRequest", string(data)}, " ")
}
