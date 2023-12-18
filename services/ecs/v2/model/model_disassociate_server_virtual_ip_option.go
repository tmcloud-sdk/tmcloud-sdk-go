package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DisassociateServerVirtualIpOption struct {
	SubnetId DisassociateServerVirtualIpOptionSubnetId `json:"subnet_id"`

	IpAddress DisassociateServerVirtualIpOptionIpAddress `json:"ip_address"`

	ReverseBinding *bool `json:"reverse_binding,omitempty"`
}

func (o DisassociateServerVirtualIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateServerVirtualIpOption struct{}"
	}

	return strings.Join([]string{"DisassociateServerVirtualIpOption", string(data)}, " ")
}

type DisassociateServerVirtualIpOptionSubnetId struct {
	value string
}

type DisassociateServerVirtualIpOptionSubnetIdEnum struct {
	EMPTY DisassociateServerVirtualIpOptionSubnetId
}

func GetDisassociateServerVirtualIpOptionSubnetIdEnum() DisassociateServerVirtualIpOptionSubnetIdEnum {
	return DisassociateServerVirtualIpOptionSubnetIdEnum{
		EMPTY: DisassociateServerVirtualIpOptionSubnetId{
			value: "",
		},
	}
}

func (c DisassociateServerVirtualIpOptionSubnetId) Value() string {
	return c.value
}

func (c DisassociateServerVirtualIpOptionSubnetId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisassociateServerVirtualIpOptionSubnetId) UnmarshalJSON(b []byte) error {
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

type DisassociateServerVirtualIpOptionIpAddress struct {
	value string
}

type DisassociateServerVirtualIpOptionIpAddressEnum struct {
	EMPTY DisassociateServerVirtualIpOptionIpAddress
}

func GetDisassociateServerVirtualIpOptionIpAddressEnum() DisassociateServerVirtualIpOptionIpAddressEnum {
	return DisassociateServerVirtualIpOptionIpAddressEnum{
		EMPTY: DisassociateServerVirtualIpOptionIpAddress{
			value: "",
		},
	}
}

func (c DisassociateServerVirtualIpOptionIpAddress) Value() string {
	return c.value
}

func (c DisassociateServerVirtualIpOptionIpAddress) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisassociateServerVirtualIpOptionIpAddress) UnmarshalJSON(b []byte) error {
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
