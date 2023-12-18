package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateVpcOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Cidr *string `json:"cidr,omitempty"`

	Routes *[]Route `json:"routes,omitempty"`
}

func (o UpdateVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcOption struct{}"
	}

	return strings.Join([]string{"UpdateVpcOption", string(data)}, " ")
}
