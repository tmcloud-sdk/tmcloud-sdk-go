package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListVpcPeeringsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *ListVpcPeeringsRequestStatus `json:"status,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ListVpcPeeringsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcPeeringsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcPeeringsRequest", string(data)}, " ")
}

type ListVpcPeeringsRequestStatus struct {
	value string
}

type ListVpcPeeringsRequestStatusEnum struct {
	PENDING_ACCEPTANCE ListVpcPeeringsRequestStatus
	REJECTED           ListVpcPeeringsRequestStatus
	EXPIRED            ListVpcPeeringsRequestStatus
	DELETED            ListVpcPeeringsRequestStatus
	ACTIVE             ListVpcPeeringsRequestStatus
}

func GetListVpcPeeringsRequestStatusEnum() ListVpcPeeringsRequestStatusEnum {
	return ListVpcPeeringsRequestStatusEnum{
		PENDING_ACCEPTANCE: ListVpcPeeringsRequestStatus{
			value: "PENDING_ACCEPTANCE",
		},
		REJECTED: ListVpcPeeringsRequestStatus{
			value: "REJECTED",
		},
		EXPIRED: ListVpcPeeringsRequestStatus{
			value: "EXPIRED",
		},
		DELETED: ListVpcPeeringsRequestStatus{
			value: "DELETED",
		},
		ACTIVE: ListVpcPeeringsRequestStatus{
			value: "ACTIVE",
		},
	}
}

func (c ListVpcPeeringsRequestStatus) Value() string {
	return c.value
}

func (c ListVpcPeeringsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVpcPeeringsRequestStatus) UnmarshalJSON(b []byte) error {
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
