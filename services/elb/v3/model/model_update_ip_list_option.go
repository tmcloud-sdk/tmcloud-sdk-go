package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateIpListOption struct {
	Name *string `json:"name,omitempty"`

	IpList *[]UpadateIpGroupIpOption `json:"ip_list,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateIpListOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListOption struct{}"
	}

	return strings.Join([]string{"UpdateIpListOption", string(data)}, " ")
}
