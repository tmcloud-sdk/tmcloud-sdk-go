package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type Subnet struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Cidr string `json:"cidr"`

	GatewayIp string `json:"gateway_ip"`

	Ipv6Enable bool `json:"ipv6_enable"`

	CidrV6 string `json:"cidr_v6"`

	GatewayIpV6 string `json:"gateway_ip_v6"`

	DhcpEnable bool `json:"dhcp_enable"`

	PrimaryDns string `json:"primary_dns"`

	SecondaryDns string `json:"secondary_dns"`

	DnsList []string `json:"dnsList"`

	AvailabilityZone string `json:"availability_zone"`

	VpcId string `json:"vpc_id"`

	Status SubnetStatus `json:"status"`

	NeutronNetworkId string `json:"neutron_network_id"`

	NeutronSubnetId string `json:"neutron_subnet_id"`

	NeutronSubnetIdV6 string `json:"neutron_subnet_id_v6"`

	ExtraDhcpOpts []ExtraDhcpOption `json:"extra_dhcp_opts"`

	Scope *string `json:"scope,omitempty"`

	TenantId string `json:"tenant_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o Subnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subnet struct{}"
	}

	return strings.Join([]string{"Subnet", string(data)}, " ")
}

type SubnetStatus struct {
	value string
}

type SubnetStatusEnum struct {
	ACTIVE  SubnetStatus
	UNKNOWN SubnetStatus
	ERROR   SubnetStatus
}

func GetSubnetStatusEnum() SubnetStatusEnum {
	return SubnetStatusEnum{
		ACTIVE: SubnetStatus{
			value: "ACTIVE",
		},
		UNKNOWN: SubnetStatus{
			value: "UNKNOWN",
		},
		ERROR: SubnetStatus{
			value: "ERROR",
		},
	}
}

func (c SubnetStatus) Value() string {
	return c.value
}

func (c SubnetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubnetStatus) UnmarshalJSON(b []byte) error {
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
