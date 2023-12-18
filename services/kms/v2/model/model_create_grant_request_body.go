package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateGrantRequestBody struct {
	KeyId string `json:"key_id"`

	GranteePrincipal string `json:"grantee_principal"`

	Operations []string `json:"operations"`

	Name *string `json:"name,omitempty"`

	RetiringPrincipal *string `json:"retiring_principal,omitempty"`

	GranteePrincipalType *CreateGrantRequestBodyGranteePrincipalType `json:"grantee_principal_type,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o CreateGrantRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGrantRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGrantRequestBody", string(data)}, " ")
}

type CreateGrantRequestBodyGranteePrincipalType struct {
	value string
}

type CreateGrantRequestBodyGranteePrincipalTypeEnum struct {
	USER   CreateGrantRequestBodyGranteePrincipalType
	DOMAIN CreateGrantRequestBodyGranteePrincipalType
}

func GetCreateGrantRequestBodyGranteePrincipalTypeEnum() CreateGrantRequestBodyGranteePrincipalTypeEnum {
	return CreateGrantRequestBodyGranteePrincipalTypeEnum{
		USER: CreateGrantRequestBodyGranteePrincipalType{
			value: "user",
		},
		DOMAIN: CreateGrantRequestBodyGranteePrincipalType{
			value: "domain",
		},
	}
}

func (c CreateGrantRequestBodyGranteePrincipalType) Value() string {
	return c.value
}

func (c CreateGrantRequestBodyGranteePrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGrantRequestBodyGranteePrincipalType) UnmarshalJSON(b []byte) error {
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
