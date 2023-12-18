package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListScalingPoliciesRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingPolicyType *ListScalingPoliciesRequestScalingPolicyType `json:"scaling_policy_type,omitempty"`

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScalingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListScalingPoliciesRequest", string(data)}, " ")
}

type ListScalingPoliciesRequestScalingPolicyType struct {
	value string
}

type ListScalingPoliciesRequestScalingPolicyTypeEnum struct {
	ALARM      ListScalingPoliciesRequestScalingPolicyType
	SCHEDULED  ListScalingPoliciesRequestScalingPolicyType
	RECURRENCE ListScalingPoliciesRequestScalingPolicyType
}

func GetListScalingPoliciesRequestScalingPolicyTypeEnum() ListScalingPoliciesRequestScalingPolicyTypeEnum {
	return ListScalingPoliciesRequestScalingPolicyTypeEnum{
		ALARM: ListScalingPoliciesRequestScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: ListScalingPoliciesRequestScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: ListScalingPoliciesRequestScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c ListScalingPoliciesRequestScalingPolicyType) Value() string {
	return c.value
}

func (c ListScalingPoliciesRequestScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScalingPoliciesRequestScalingPolicyType) UnmarshalJSON(b []byte) error {
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
