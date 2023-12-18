package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListScalingGroupsRequest struct {
	ScalingGroupName *string `json:"scaling_group_name,omitempty"`

	ScalingConfigurationId *string `json:"scaling_configuration_id,omitempty"`

	ScalingGroupStatus *ListScalingGroupsRequestScalingGroupStatus `json:"scaling_group_status,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListScalingGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListScalingGroupsRequest", string(data)}, " ")
}

type ListScalingGroupsRequestScalingGroupStatus struct {
	value string
}

type ListScalingGroupsRequestScalingGroupStatusEnum struct {
	INSERVICE ListScalingGroupsRequestScalingGroupStatus
	PAUSED    ListScalingGroupsRequestScalingGroupStatus
	ERROR     ListScalingGroupsRequestScalingGroupStatus
	DELETING  ListScalingGroupsRequestScalingGroupStatus
}

func GetListScalingGroupsRequestScalingGroupStatusEnum() ListScalingGroupsRequestScalingGroupStatusEnum {
	return ListScalingGroupsRequestScalingGroupStatusEnum{
		INSERVICE: ListScalingGroupsRequestScalingGroupStatus{
			value: "INSERVICE",
		},
		PAUSED: ListScalingGroupsRequestScalingGroupStatus{
			value: "PAUSED",
		},
		ERROR: ListScalingGroupsRequestScalingGroupStatus{
			value: "ERROR",
		},
		DELETING: ListScalingGroupsRequestScalingGroupStatus{
			value: "DELETING",
		},
	}
}

func (c ListScalingGroupsRequestScalingGroupStatus) Value() string {
	return c.value
}

func (c ListScalingGroupsRequestScalingGroupStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScalingGroupsRequestScalingGroupStatus) UnmarshalJSON(b []byte) error {
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
