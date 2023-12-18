package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ModifyJobResp struct {
	Id string `json:"id"`

	Status ModifyJobRespStatus `json:"status"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ModifyJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyJobResp struct{}"
	}

	return strings.Join([]string{"ModifyJobResp", string(data)}, " ")
}

type ModifyJobRespStatus struct {
	value string
}

type ModifyJobRespStatusEnum struct {
	SUCCESS ModifyJobRespStatus
	FAILED  ModifyJobRespStatus
}

func GetModifyJobRespStatusEnum() ModifyJobRespStatusEnum {
	return ModifyJobRespStatusEnum{
		SUCCESS: ModifyJobRespStatus{
			value: "success",
		},
		FAILED: ModifyJobRespStatus{
			value: "failed",
		},
	}
}

func (c ModifyJobRespStatus) Value() string {
	return c.value
}

func (c ModifyJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobRespStatus) UnmarshalJSON(b []byte) error {
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
