package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type L7Rule struct {
	AdminStateUp bool `json:"admin_state_up"`

	CompareType string `json:"compare_type"`

	Key string `json:"key"`

	ProjectId string `json:"project_id"`

	Type L7RuleType `json:"type"`

	Value string `json:"value"`

	ProvisioningStatus string `json:"provisioning_status"`

	Invert bool `json:"invert"`

	Id string `json:"id"`

	Conditions []RuleCondition `json:"conditions"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o L7Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7Rule struct{}"
	}

	return strings.Join([]string{"L7Rule", string(data)}, " ")
}

type L7RuleType struct {
	value string
}

type L7RuleTypeEnum struct {
	HOST_NAME    L7RuleType
	PATH         L7RuleType
	METHOD       L7RuleType
	HEADER       L7RuleType
	QUERY_STRING L7RuleType
	SOURCE_IP    L7RuleType
}

func GetL7RuleTypeEnum() L7RuleTypeEnum {
	return L7RuleTypeEnum{
		HOST_NAME: L7RuleType{
			value: "HOST_NAME",
		},
		PATH: L7RuleType{
			value: "PATH",
		},
		METHOD: L7RuleType{
			value: "METHOD",
		},
		HEADER: L7RuleType{
			value: "HEADER",
		},
		QUERY_STRING: L7RuleType{
			value: "QUERY_STRING",
		},
		SOURCE_IP: L7RuleType{
			value: "SOURCE_IP",
		},
	}
}

func (c L7RuleType) Value() string {
	return c.value
}

func (c L7RuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7RuleType) UnmarshalJSON(b []byte) error {
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
