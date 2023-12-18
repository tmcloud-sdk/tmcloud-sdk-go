package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListNatGatewaySnatRulesRequest struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Cidr *string `json:"cidr,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	FloatingIpAddress *string `json:"floating_ip_address,omitempty"`

	FloatingIpId *string `json:"floating_ip_id,omitempty"`

	Id *string `json:"id,omitempty"`

	Description *string `json:"description,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	NatGatewayId *[]string `json:"nat_gateway_id,omitempty"`

	NetworkId *string `json:"network_id,omitempty"`

	SourceType *int32 `json:"source_type,omitempty"`

	Status *ListNatGatewaySnatRulesRequestStatus `json:"status,omitempty"`
}

func (o ListNatGatewaySnatRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaySnatRulesRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewaySnatRulesRequest", string(data)}, " ")
}

type ListNatGatewaySnatRulesRequestStatus struct {
	value string
}

type ListNatGatewaySnatRulesRequestStatusEnum struct {
	ACTIVE         ListNatGatewaySnatRulesRequestStatus
	PENDING_CREATE ListNatGatewaySnatRulesRequestStatus
	PENDING_UPDATE ListNatGatewaySnatRulesRequestStatus
	PENDING_DELETE ListNatGatewaySnatRulesRequestStatus
	EIP_FREEZED    ListNatGatewaySnatRulesRequestStatus
	INACTIVE       ListNatGatewaySnatRulesRequestStatus
}

func GetListNatGatewaySnatRulesRequestStatusEnum() ListNatGatewaySnatRulesRequestStatusEnum {
	return ListNatGatewaySnatRulesRequestStatusEnum{
		ACTIVE: ListNatGatewaySnatRulesRequestStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: ListNatGatewaySnatRulesRequestStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: ListNatGatewaySnatRulesRequestStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: ListNatGatewaySnatRulesRequestStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: ListNatGatewaySnatRulesRequestStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: ListNatGatewaySnatRulesRequestStatus{
			value: "INACTIVE",
		},
	}
}

func (c ListNatGatewaySnatRulesRequestStatus) Value() string {
	return c.value
}

func (c ListNatGatewaySnatRulesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewaySnatRulesRequestStatus) UnmarshalJSON(b []byte) error {
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
