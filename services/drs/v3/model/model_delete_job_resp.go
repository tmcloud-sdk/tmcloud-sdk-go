package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteJobResp struct {
	Id *string `json:"id,omitempty"`

	Status *DeleteJobRespStatus `json:"status,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DeleteJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobResp struct{}"
	}

	return strings.Join([]string{"DeleteJobResp", string(data)}, " ")
}

type DeleteJobRespStatus struct {
	value string
}

type DeleteJobRespStatusEnum struct {
	SUCCESS DeleteJobRespStatus
	FAILED  DeleteJobRespStatus
}

func GetDeleteJobRespStatusEnum() DeleteJobRespStatusEnum {
	return DeleteJobRespStatusEnum{
		SUCCESS: DeleteJobRespStatus{
			value: "success",
		},
		FAILED: DeleteJobRespStatus{
			value: "failed",
		},
	}
}

func (c DeleteJobRespStatus) Value() string {
	return c.value
}

func (c DeleteJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteJobRespStatus) UnmarshalJSON(b []byte) error {
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
