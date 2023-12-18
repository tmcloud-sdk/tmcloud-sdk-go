package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ContainerNetwork struct {
	Mode ContainerNetworkMode `json:"mode"`

	Cidr *string `json:"cidr,omitempty"`

	Cidrs *[]ContainerCidr `json:"cidrs,omitempty"`
}

func (o ContainerNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerNetwork struct{}"
	}

	return strings.Join([]string{"ContainerNetwork", string(data)}, " ")
}

type ContainerNetworkMode struct {
	value string
}

type ContainerNetworkModeEnum struct {
	OVERLAY_L2 ContainerNetworkMode
	VPC_ROUTER ContainerNetworkMode
	ENI        ContainerNetworkMode
}

func GetContainerNetworkModeEnum() ContainerNetworkModeEnum {
	return ContainerNetworkModeEnum{
		OVERLAY_L2: ContainerNetworkMode{
			value: "overlay_l2",
		},
		VPC_ROUTER: ContainerNetworkMode{
			value: "vpc-router",
		},
		ENI: ContainerNetworkMode{
			value: "eni",
		},
	}
}

func (c ContainerNetworkMode) Value() string {
	return c.value
}

func (c ContainerNetworkMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContainerNetworkMode) UnmarshalJSON(b []byte) error {
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
