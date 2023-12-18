package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListVpcRoutesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *string `json:"id,omitempty"`

	Type *ListVpcRoutesRequestType `json:"type,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	Destination *string `json:"destination,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (o ListVpcRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcRoutesRequest struct{}"
	}

	return strings.Join([]string{"ListVpcRoutesRequest", string(data)}, " ")
}

type ListVpcRoutesRequestType struct {
	value string
}

type ListVpcRoutesRequestTypeEnum struct {
	PEERING ListVpcRoutesRequestType
}

func GetListVpcRoutesRequestTypeEnum() ListVpcRoutesRequestTypeEnum {
	return ListVpcRoutesRequestTypeEnum{
		PEERING: ListVpcRoutesRequestType{
			value: "peering",
		},
	}
}

func (c ListVpcRoutesRequestType) Value() string {
	return c.value
}

func (c ListVpcRoutesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVpcRoutesRequestType) UnmarshalJSON(b []byte) error {
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
