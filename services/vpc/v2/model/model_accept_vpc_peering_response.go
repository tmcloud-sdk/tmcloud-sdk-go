package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type AcceptVpcPeeringResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *AcceptVpcPeeringResponseStatus `json:"status,omitempty"`

	RequestVpcInfo *VpcInfo `json:"request_vpc_info,omitempty"`

	AcceptVpcInfo *VpcInfo `json:"accept_vpc_info,omitempty"`

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AcceptVpcPeeringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptVpcPeeringResponse struct{}"
	}

	return strings.Join([]string{"AcceptVpcPeeringResponse", string(data)}, " ")
}

type AcceptVpcPeeringResponseStatus struct {
	value string
}

type AcceptVpcPeeringResponseStatusEnum struct {
	PENDING_ACCEPTANCE AcceptVpcPeeringResponseStatus
	REJECTED           AcceptVpcPeeringResponseStatus
	EXPIRED            AcceptVpcPeeringResponseStatus
	DELETED            AcceptVpcPeeringResponseStatus
	ACTIVE             AcceptVpcPeeringResponseStatus
}

func GetAcceptVpcPeeringResponseStatusEnum() AcceptVpcPeeringResponseStatusEnum {
	return AcceptVpcPeeringResponseStatusEnum{
		PENDING_ACCEPTANCE: AcceptVpcPeeringResponseStatus{
			value: "PENDING_ACCEPTANCE",
		},
		REJECTED: AcceptVpcPeeringResponseStatus{
			value: "REJECTED",
		},
		EXPIRED: AcceptVpcPeeringResponseStatus{
			value: "EXPIRED",
		},
		DELETED: AcceptVpcPeeringResponseStatus{
			value: "DELETED",
		},
		ACTIVE: AcceptVpcPeeringResponseStatus{
			value: "ACTIVE",
		},
	}
}

func (c AcceptVpcPeeringResponseStatus) Value() string {
	return c.value
}

func (c AcceptVpcPeeringResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AcceptVpcPeeringResponseStatus) UnmarshalJSON(b []byte) error {
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
