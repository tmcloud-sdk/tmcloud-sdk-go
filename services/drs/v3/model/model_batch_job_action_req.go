package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchJobActionReq struct {
	Action BatchJobActionReqAction `json:"action"`

	JobId string `json:"job_id"`

	Property string `json:"property"`
}

func (o BatchJobActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchJobActionReq struct{}"
	}

	return strings.Join([]string{"BatchJobActionReq", string(data)}, " ")
}

type BatchJobActionReqAction struct {
	value string
}

type BatchJobActionReqActionEnum struct {
	TEST_CONNECTION BatchJobActionReqAction
}

func GetBatchJobActionReqActionEnum() BatchJobActionReqActionEnum {
	return BatchJobActionReqActionEnum{
		TEST_CONNECTION: BatchJobActionReqAction{
			value: "testConnection",
		},
	}
}

func (c BatchJobActionReqAction) Value() string {
	return c.value
}

func (c BatchJobActionReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchJobActionReqAction) UnmarshalJSON(b []byte) error {
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
