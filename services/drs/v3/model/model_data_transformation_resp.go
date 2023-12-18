package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DataTransformationResp struct {
	Id *string `json:"id,omitempty"`

	Status *DataTransformationRespStatus `json:"status,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DataTransformationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataTransformationResp struct{}"
	}

	return strings.Join([]string{"DataTransformationResp", string(data)}, " ")
}

type DataTransformationRespStatus struct {
	value string
}

type DataTransformationRespStatusEnum struct {
	SUCCESS DataTransformationRespStatus
	FAILED  DataTransformationRespStatus
}

func GetDataTransformationRespStatusEnum() DataTransformationRespStatusEnum {
	return DataTransformationRespStatusEnum{
		SUCCESS: DataTransformationRespStatus{
			value: "success",
		},
		FAILED: DataTransformationRespStatus{
			value: "failed",
		},
	}
}

func (c DataTransformationRespStatus) Value() string {
	return c.value
}

func (c DataTransformationRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataTransformationRespStatus) UnmarshalJSON(b []byte) error {
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
