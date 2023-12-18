package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ResourceResult struct {
	Type ResourceResultType `json:"type"`

	Used int32 `json:"used"`

	Quota int32 `json:"quota"`

	Min int32 `json:"min"`
}

func (o ResourceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceResult struct{}"
	}

	return strings.Join([]string{"ResourceResult", string(data)}, " ")
}

type ResourceResultType struct {
	value string
}

type ResourceResultTypeEnum struct {
	VPC                       ResourceResultType
	SUBNET                    ResourceResultType
	SECURITY_GROUP            ResourceResultType
	SECURITY_GROUP_RULE       ResourceResultType
	PUBLIC_IP                 ResourceResultType
	VPN                       ResourceResultType
	VPNGW                     ResourceResultType
	VPC_PEER                  ResourceResultType
	FIREWALL                  ResourceResultType
	SHARE_BANDWIDTH           ResourceResultType
	SHARE_BANDWIDTH_IP        ResourceResultType
	LOADBALANCER              ResourceResultType
	LISTENER                  ResourceResultType
	PHYSICAL_CONNECT          ResourceResultType
	VIRTUAL_INTERFACE         ResourceResultType
	VPC_CONTAIN_ROUTETABLE    ResourceResultType
	ROUTETABLE_CONTAIN_ROUTES ResourceResultType
}

func GetResourceResultTypeEnum() ResourceResultTypeEnum {
	return ResourceResultTypeEnum{
		VPC: ResourceResultType{
			value: "vpc",
		},
		SUBNET: ResourceResultType{
			value: "subnet",
		},
		SECURITY_GROUP: ResourceResultType{
			value: "securityGroup",
		},
		SECURITY_GROUP_RULE: ResourceResultType{
			value: "securityGroupRule",
		},
		PUBLIC_IP: ResourceResultType{
			value: "publicIp",
		},
		VPN: ResourceResultType{
			value: "vpn",
		},
		VPNGW: ResourceResultType{
			value: "vpngw",
		},
		VPC_PEER: ResourceResultType{
			value: "vpcPeer",
		},
		FIREWALL: ResourceResultType{
			value: "firewall",
		},
		SHARE_BANDWIDTH: ResourceResultType{
			value: "shareBandwidth",
		},
		SHARE_BANDWIDTH_IP: ResourceResultType{
			value: "shareBandwidthIP",
		},
		LOADBALANCER: ResourceResultType{
			value: "loadbalancer",
		},
		LISTENER: ResourceResultType{
			value: "listener",
		},
		PHYSICAL_CONNECT: ResourceResultType{
			value: "physicalConnect",
		},
		VIRTUAL_INTERFACE: ResourceResultType{
			value: "virtualInterface",
		},
		VPC_CONTAIN_ROUTETABLE: ResourceResultType{
			value: "vpcContainRoutetable",
		},
		ROUTETABLE_CONTAIN_ROUTES: ResourceResultType{
			value: "routetableContainRoutes",
		},
	}
}

func (c ResourceResultType) Value() string {
	return c.value
}

func (c ResourceResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceResultType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
