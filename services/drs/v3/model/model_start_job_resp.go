package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type StartJobResp struct {
	Id string `json:"id"`

	Status *StartJobRespStatus `json:"status,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o StartJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobResp struct{}"
	}

	return strings.Join([]string{"StartJobResp", string(data)}, " ")
}

type StartJobRespStatus struct {
	value string
}

type StartJobRespStatusEnum struct {
	SUCCESS StartJobRespStatus
	FAILED  StartJobRespStatus
}

func GetStartJobRespStatusEnum() StartJobRespStatusEnum {
	return StartJobRespStatusEnum{
		SUCCESS: StartJobRespStatus{
			value: "success",
		},
		FAILED: StartJobRespStatus{
			value: "failed",
		},
	}
}

func (c StartJobRespStatus) Value() string {
	return c.value
}

func (c StartJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartJobRespStatus) UnmarshalJSON(b []byte) error {
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
