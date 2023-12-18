package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateCompareTaskReq struct {
	JobId string `json:"job_id"`

	ObjectLevelCompareType *CreateCompareTaskReqObjectLevelCompareType `json:"object_level_compare_type,omitempty"`

	DataLevelCompareInfo *CreateDataLevelCompareReq `json:"data_level_compare_info,omitempty"`
}

func (o CreateCompareTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareTaskReq struct{}"
	}

	return strings.Join([]string{"CreateCompareTaskReq", string(data)}, " ")
}

type CreateCompareTaskReqObjectLevelCompareType struct {
	value string
}

type CreateCompareTaskReqObjectLevelCompareTypeEnum struct {
	OBJECTS CreateCompareTaskReqObjectLevelCompareType
}

func GetCreateCompareTaskReqObjectLevelCompareTypeEnum() CreateCompareTaskReqObjectLevelCompareTypeEnum {
	return CreateCompareTaskReqObjectLevelCompareTypeEnum{
		OBJECTS: CreateCompareTaskReqObjectLevelCompareType{
			value: "objects",
		},
	}
}

func (c CreateCompareTaskReqObjectLevelCompareType) Value() string {
	return c.value
}

func (c CreateCompareTaskReqObjectLevelCompareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCompareTaskReqObjectLevelCompareType) UnmarshalJSON(b []byte) error {
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
