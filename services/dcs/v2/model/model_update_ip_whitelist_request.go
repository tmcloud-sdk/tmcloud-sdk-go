package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateIpWhitelistRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ModifyIpWhitelistBody `json:"body,omitempty"`
}

func (o UpdateIpWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpWhitelistRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpWhitelistRequest", string(data)}, " ")
}
