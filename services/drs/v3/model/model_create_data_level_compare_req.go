package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateDataLevelCompareReq struct {
	ConflictPolicy CreateDataLevelCompareReqConflictPolicy `json:"conflict_policy"`

	CompareType CreateDataLevelCompareReqCompareType `json:"compare_type"`

	CompareMode *CreateDataLevelCompareReqCompareMode `json:"compare_mode,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	CompareObjectInfos *[]CompareObjectInfo `json:"compare_object_infos,omitempty"`

	CompareObjectInfosWithToken *[]CompareObjectInfoWithToken `json:"compare_object_infos_with_token,omitempty"`
}

func (o CreateDataLevelCompareReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataLevelCompareReq struct{}"
	}

	return strings.Join([]string{"CreateDataLevelCompareReq", string(data)}, " ")
}

type CreateDataLevelCompareReqConflictPolicy struct {
	value string
}

type CreateDataLevelCompareReqConflictPolicyEnum struct {
	CANCEL CreateDataLevelCompareReqConflictPolicy
	KEEP   CreateDataLevelCompareReqConflictPolicy
}

func GetCreateDataLevelCompareReqConflictPolicyEnum() CreateDataLevelCompareReqConflictPolicyEnum {
	return CreateDataLevelCompareReqConflictPolicyEnum{
		CANCEL: CreateDataLevelCompareReqConflictPolicy{
			value: "cancel",
		},
		KEEP: CreateDataLevelCompareReqConflictPolicy{
			value: "keep",
		},
	}
}

func (c CreateDataLevelCompareReqConflictPolicy) Value() string {
	return c.value
}

func (c CreateDataLevelCompareReqConflictPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataLevelCompareReqConflictPolicy) UnmarshalJSON(b []byte) error {
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

type CreateDataLevelCompareReqCompareType struct {
	value string
}

type CreateDataLevelCompareReqCompareTypeEnum struct {
	LINES    CreateDataLevelCompareReqCompareType
	CONTENTS CreateDataLevelCompareReqCompareType
}

func GetCreateDataLevelCompareReqCompareTypeEnum() CreateDataLevelCompareReqCompareTypeEnum {
	return CreateDataLevelCompareReqCompareTypeEnum{
		LINES: CreateDataLevelCompareReqCompareType{
			value: "lines",
		},
		CONTENTS: CreateDataLevelCompareReqCompareType{
			value: "contents",
		},
	}
}

func (c CreateDataLevelCompareReqCompareType) Value() string {
	return c.value
}

func (c CreateDataLevelCompareReqCompareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataLevelCompareReqCompareType) UnmarshalJSON(b []byte) error {
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

type CreateDataLevelCompareReqCompareMode struct {
	value string
}

type CreateDataLevelCompareReqCompareModeEnum struct {
	QUICK_COMPARISON CreateDataLevelCompareReqCompareMode
}

func GetCreateDataLevelCompareReqCompareModeEnum() CreateDataLevelCompareReqCompareModeEnum {
	return CreateDataLevelCompareReqCompareModeEnum{
		QUICK_COMPARISON: CreateDataLevelCompareReqCompareMode{
			value: "quick_comparison",
		},
	}
}

func (c CreateDataLevelCompareReqCompareMode) Value() string {
	return c.value
}

func (c CreateDataLevelCompareReqCompareMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataLevelCompareReqCompareMode) UnmarshalJSON(b []byte) error {
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
