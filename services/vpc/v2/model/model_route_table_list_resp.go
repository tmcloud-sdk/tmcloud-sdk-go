package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RouteTableListResp struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Default bool `json:"default"`

	Subnets []SubnetList `json:"subnets"`

	TenantId string `json:"tenant_id"`

	VpcId string `json:"vpc_id"`

	Description string `json:"description"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o RouteTableListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteTableListResp struct{}"
	}

	return strings.Join([]string{"RouteTableListResp", string(data)}, " ")
}
