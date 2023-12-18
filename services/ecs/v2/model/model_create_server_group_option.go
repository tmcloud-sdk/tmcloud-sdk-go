package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateServerGroupOption struct {
	Name string `json:"name"`

	Policies []CreateServerGroupOptionPolicies `json:"policies"`
}

func (o CreateServerGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupOption struct{}"
	}

	return strings.Join([]string{"CreateServerGroupOption", string(data)}, " ")
}

type CreateServerGroupOptionPolicies struct {
	value string
}

type CreateServerGroupOptionPoliciesEnum struct {
	ANTI_AFFINITY CreateServerGroupOptionPolicies
}

func GetCreateServerGroupOptionPoliciesEnum() CreateServerGroupOptionPoliciesEnum {
	return CreateServerGroupOptionPoliciesEnum{
		ANTI_AFFINITY: CreateServerGroupOptionPolicies{
			value: "anti-affinity",
		},
	}
}

func (c CreateServerGroupOptionPolicies) Value() string {
	return c.value
}

func (c CreateServerGroupOptionPolicies) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServerGroupOptionPolicies) UnmarshalJSON(b []byte) error {
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
