package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaAssociateSecurityGroupRequest struct {
	ServerId string `json:"server_id"`

	Body *NovaAssociateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NovaAssociateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAssociateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NovaAssociateSecurityGroupRequest", string(data)}, " ")
}
