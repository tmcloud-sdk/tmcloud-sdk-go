package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListNatGatewayDnatRulesRequest struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ExternalServicePort *int32 `json:"external_service_port,omitempty"`

	FloatingIpAddress *string `json:"floating_ip_address,omitempty"`

	Status *[]ListNatGatewayDnatRulesRequestStatus `json:"status,omitempty"`

	FloatingIpId *string `json:"floating_ip_id,omitempty"`

	InternalServicePort *int32 `json:"internal_service_port,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Id *string `json:"id,omitempty"`

	Description *string `json:"description,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	NatGatewayId *[]string `json:"nat_gateway_id,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	Protocol *[]string `json:"protocol,omitempty"`
}

func (o ListNatGatewayDnatRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewayDnatRulesRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewayDnatRulesRequest", string(data)}, " ")
}

type ListNatGatewayDnatRulesRequestStatus struct {
	value string
}

type ListNatGatewayDnatRulesRequestStatusEnum struct {
	ACTIVE         ListNatGatewayDnatRulesRequestStatus
	PENDING_CREATE ListNatGatewayDnatRulesRequestStatus
	PENDING_UPDATE ListNatGatewayDnatRulesRequestStatus
	PENDING_DELETE ListNatGatewayDnatRulesRequestStatus
	EIP_FREEZED    ListNatGatewayDnatRulesRequestStatus
	INACTIVE       ListNatGatewayDnatRulesRequestStatus
}

func GetListNatGatewayDnatRulesRequestStatusEnum() ListNatGatewayDnatRulesRequestStatusEnum {
	return ListNatGatewayDnatRulesRequestStatusEnum{
		ACTIVE: ListNatGatewayDnatRulesRequestStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: ListNatGatewayDnatRulesRequestStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: ListNatGatewayDnatRulesRequestStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: ListNatGatewayDnatRulesRequestStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: ListNatGatewayDnatRulesRequestStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: ListNatGatewayDnatRulesRequestStatus{
			value: "INACTIVE",
		},
	}
}

func (c ListNatGatewayDnatRulesRequestStatus) Value() string {
	return c.value
}

func (c ListNatGatewayDnatRulesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewayDnatRulesRequestStatus) UnmarshalJSON(b []byte) error {
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
