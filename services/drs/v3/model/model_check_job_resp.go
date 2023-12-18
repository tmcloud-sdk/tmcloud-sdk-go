package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CheckJobResp struct {
	Id string `json:"id"`

	Status CheckJobRespStatus `json:"status"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func (o CheckJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobResp struct{}"
	}

	return strings.Join([]string{"CheckJobResp", string(data)}, " ")
}

type CheckJobRespStatus struct {
	value string
}

type CheckJobRespStatusEnum struct {
	TRUE  CheckJobRespStatus
	FALSE CheckJobRespStatus
}

func GetCheckJobRespStatusEnum() CheckJobRespStatusEnum {
	return CheckJobRespStatusEnum{
		TRUE: CheckJobRespStatus{
			value: "true",
		},
		FALSE: CheckJobRespStatus{
			value: "false",
		},
	}
}

func (c CheckJobRespStatus) Value() string {
	return c.value
}

func (c CheckJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckJobRespStatus) UnmarshalJSON(b []byte) error {
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
