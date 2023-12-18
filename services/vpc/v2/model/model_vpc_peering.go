package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type VpcPeering struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Status VpcPeeringStatus `json:"status"`

	RequestVpcInfo *VpcInfo `json:"request_vpc_info"`

	AcceptVpcInfo *VpcInfo `json:"accept_vpc_info"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	Description string `json:"description"`
}

func (o VpcPeering) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcPeering struct{}"
	}

	return strings.Join([]string{"VpcPeering", string(data)}, " ")
}

type VpcPeeringStatus struct {
	value string
}

type VpcPeeringStatusEnum struct {
	PENDING_ACCEPTANCE VpcPeeringStatus
	REJECTED           VpcPeeringStatus
	EXPIRED            VpcPeeringStatus
	DELETED            VpcPeeringStatus
	ACTIVE             VpcPeeringStatus
}

func GetVpcPeeringStatusEnum() VpcPeeringStatusEnum {
	return VpcPeeringStatusEnum{
		PENDING_ACCEPTANCE: VpcPeeringStatus{
			value: "PENDING_ACCEPTANCE",
		},
		REJECTED: VpcPeeringStatus{
			value: "REJECTED",
		},
		EXPIRED: VpcPeeringStatus{
			value: "EXPIRED",
		},
		DELETED: VpcPeeringStatus{
			value: "DELETED",
		},
		ACTIVE: VpcPeeringStatus{
			value: "ACTIVE",
		},
	}
}

func (c VpcPeeringStatus) Value() string {
	return c.value
}

func (c VpcPeeringStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcPeeringStatus) UnmarshalJSON(b []byte) error {
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
