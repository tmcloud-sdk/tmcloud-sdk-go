package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type Privateip struct {
	Status PrivateipStatus `json:"status"`

	Id string `json:"id"`

	SubnetId string `json:"subnet_id"`

	TenantId string `json:"tenant_id"`

	DeviceOwner PrivateipDeviceOwner `json:"device_owner"`

	IpAddress string `json:"ip_address"`
}

func (o Privateip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Privateip struct{}"
	}

	return strings.Join([]string{"Privateip", string(data)}, " ")
}

type PrivateipStatus struct {
	value string
}

type PrivateipStatusEnum struct {
	ACTIVE PrivateipStatus
	DOWN   PrivateipStatus
}

func GetPrivateipStatusEnum() PrivateipStatusEnum {
	return PrivateipStatusEnum{
		ACTIVE: PrivateipStatus{
			value: "ACTIVE",
		},
		DOWN: PrivateipStatus{
			value: "DOWN",
		},
	}
}

func (c PrivateipStatus) Value() string {
	return c.value
}

func (c PrivateipStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateipStatus) UnmarshalJSON(b []byte) error {
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

type PrivateipDeviceOwner struct {
	value string
}

type PrivateipDeviceOwnerEnum struct {
	NETWORKDHCP                         PrivateipDeviceOwner
	NETWORKROUTER_INTERFACE_DISTRIBUTED PrivateipDeviceOwner
	COMPUTEXXX                          PrivateipDeviceOwner
}

func GetPrivateipDeviceOwnerEnum() PrivateipDeviceOwnerEnum {
	return PrivateipDeviceOwnerEnum{
		NETWORKDHCP: PrivateipDeviceOwner{
			value: "network:dhcp",
		},
		NETWORKROUTER_INTERFACE_DISTRIBUTED: PrivateipDeviceOwner{
			value: "network:router_interface_distributed",
		},
		COMPUTEXXX: PrivateipDeviceOwner{
			value: "compute:xxx",
		},
	}
}

func (c PrivateipDeviceOwner) Value() string {
	return c.value
}

func (c PrivateipDeviceOwner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateipDeviceOwner) UnmarshalJSON(b []byte) error {
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
