package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RouteTableResp struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Default bool `json:"default"`

	Routes []RouteTableRoute `json:"routes"`

	Subnets []SubnetList `json:"subnets"`

	TenantId string `json:"tenant_id"`

	VpcId string `json:"vpc_id"`

	Description string `json:"description"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o RouteTableResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteTableResp struct{}"
	}

	return strings.Join([]string{"RouteTableResp", string(data)}, " ")
}
