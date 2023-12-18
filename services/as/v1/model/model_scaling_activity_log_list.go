package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type ScalingActivityLogList struct {
	Status *ScalingActivityLogListStatus `json:"status,omitempty"`

	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Id *string `json:"id,omitempty"`

	InstanceRemovedList *string `json:"instance_removed_list,omitempty"`

	InstanceDeletedList *string `json:"instance_deleted_list,omitempty"`

	InstanceAddedList *string `json:"instance_added_list,omitempty"`

	ScalingValue *string `json:"scaling_value,omitempty"`

	Description *string `json:"description,omitempty"`

	InstanceValue *int32 `json:"instance_value,omitempty"`

	DesireValue *int32 `json:"desire_value,omitempty"`
}

func (o ScalingActivityLogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingActivityLogList struct{}"
	}

	return strings.Join([]string{"ScalingActivityLogList", string(data)}, " ")
}

type ScalingActivityLogListStatus struct {
	value string
}

type ScalingActivityLogListStatusEnum struct {
	SUCCESS ScalingActivityLogListStatus
	FAIL    ScalingActivityLogListStatus
	DING    ScalingActivityLogListStatus
}

func GetScalingActivityLogListStatusEnum() ScalingActivityLogListStatusEnum {
	return ScalingActivityLogListStatusEnum{
		SUCCESS: ScalingActivityLogListStatus{
			value: "SUCCESS",
		},
		FAIL: ScalingActivityLogListStatus{
			value: "FAIL",
		},
		DING: ScalingActivityLogListStatus{
			value: "DING",
		},
	}
}

func (c ScalingActivityLogListStatus) Value() string {
	return c.value
}

func (c ScalingActivityLogListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingActivityLogListStatus) UnmarshalJSON(b []byte) error {
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
