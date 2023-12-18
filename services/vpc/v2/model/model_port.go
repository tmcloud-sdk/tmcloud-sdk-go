package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type Port struct {
	Id string `json:"id"`

	Name string `json:"name"`

	NetworkId string `json:"network_id"`

	AdminStateUp bool `json:"admin_state_up"`

	MacAddress string `json:"mac_address"`

	FixedIps []FixedIp `json:"fixed_ips"`

	DeviceId string `json:"device_id"`

	DeviceOwner PortDeviceOwner `json:"device_owner"`

	TenantId string `json:"tenant_id"`

	Status PortStatus `json:"status"`

	SecurityGroups []string `json:"security_groups"`

	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs"`

	ExtraDhcpOpts []ExtraDhcpOpt `json:"extra_dhcp_opts"`

	BindingvnicType string `json:"binding:vnic_type"`

	DnsAssignment []DnsAssignMent `json:"dns_assignment"`

	DnsName string `json:"dns_name"`

	BindingvifDetails *BindingVifDetails `json:"binding:vif_details"`

	Bindingprofile *interface{} `json:"binding:profile"`

	InstanceId string `json:"instance_id"`

	InstanceType string `json:"instance_type"`

	PortSecurityEnabled bool `json:"port_security_enabled"`

	ZoneId string `json:"zone_id"`
}

func (o Port) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Port struct{}"
	}

	return strings.Join([]string{"Port", string(data)}, " ")
}

type PortDeviceOwner struct {
	value string
}

type PortDeviceOwnerEnum struct {
	NETWORKDHCP                         PortDeviceOwner
	NETWORKVIP_PORT                     PortDeviceOwner
	NETWORKROUTER_INTERFACE_DISTRIBUTED PortDeviceOwner
	NETWORKROUTER_CENTRALIZED_SNAT      PortDeviceOwner
}

func GetPortDeviceOwnerEnum() PortDeviceOwnerEnum {
	return PortDeviceOwnerEnum{
		NETWORKDHCP: PortDeviceOwner{
			value: "network:dhcp",
		},
		NETWORKVIP_PORT: PortDeviceOwner{
			value: "network:VIP_PORT",
		},
		NETWORKROUTER_INTERFACE_DISTRIBUTED: PortDeviceOwner{
			value: "network:router_interface_distributed",
		},
		NETWORKROUTER_CENTRALIZED_SNAT: PortDeviceOwner{
			value: "network:router_centralized_snat",
		},
	}
}

func (c PortDeviceOwner) Value() string {
	return c.value
}

func (c PortDeviceOwner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PortDeviceOwner) UnmarshalJSON(b []byte) error {
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

type PortStatus struct {
	value string
}

type PortStatusEnum struct {
	ACTIVE PortStatus
	BUILD  PortStatus
	DOWN   PortStatus
}

func GetPortStatusEnum() PortStatusEnum {
	return PortStatusEnum{
		ACTIVE: PortStatus{
			value: "ACTIVE",
		},
		BUILD: PortStatus{
			value: "BUILD",
		},
		DOWN: PortStatus{
			value: "DOWN",
		},
	}
}

func (c PortStatus) Value() string {
	return c.value
}

func (c PortStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PortStatus) UnmarshalJSON(b []byte) error {
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
