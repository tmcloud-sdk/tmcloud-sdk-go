package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type Grants struct {
	KeyId *string `json:"key_id,omitempty"`

	GrantId *string `json:"grant_id,omitempty"`

	GranteePrincipal *string `json:"grantee_principal,omitempty"`

	GranteePrincipalType *GrantsGranteePrincipalType `json:"grantee_principal_type,omitempty"`

	Operations *[]string `json:"operations,omitempty"`

	IssuingPrincipal *string `json:"issuing_principal,omitempty"`

	CreationDate *string `json:"creation_date,omitempty"`

	Name *string `json:"name,omitempty"`

	RetiringPrincipal *string `json:"retiring_principal,omitempty"`
}

func (o Grants) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Grants struct{}"
	}

	return strings.Join([]string{"Grants", string(data)}, " ")
}

type GrantsGranteePrincipalType struct {
	value string
}

type GrantsGranteePrincipalTypeEnum struct {
	USER   GrantsGranteePrincipalType
	DOMAIN GrantsGranteePrincipalType
}

func GetGrantsGranteePrincipalTypeEnum() GrantsGranteePrincipalTypeEnum {
	return GrantsGranteePrincipalTypeEnum{
		USER: GrantsGranteePrincipalType{
			value: "user",
		},
		DOMAIN: GrantsGranteePrincipalType{
			value: "domain",
		},
	}
}

func (c GrantsGranteePrincipalType) Value() string {
	return c.value
}

func (c GrantsGranteePrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GrantsGranteePrincipalType) UnmarshalJSON(b []byte) error {
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
