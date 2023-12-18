package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrecheckResult struct {
	Item *string `json:"item,omitempty"`

	Result *PrecheckResultResult `json:"result,omitempty"`

	FailedReason *string `json:"failed_reason,omitempty"`

	Data *string `json:"data,omitempty"`

	RawErrorMsg *string `json:"raw_error_msg,omitempty"`

	Group *string `json:"group,omitempty"`

	FailedSubJobs *[]PrecheckFailSubJobVo `json:"failed_sub_jobs,omitempty"`
}

func (o PrecheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckResult struct{}"
	}

	return strings.Join([]string{"PrecheckResult", string(data)}, " ")
}

type PrecheckResultResult struct {
	value string
}

type PrecheckResultResultEnum struct {
	PASSED PrecheckResultResult
	ALARM  PrecheckResultResult
	FAILED PrecheckResultResult
}

func GetPrecheckResultResultEnum() PrecheckResultResultEnum {
	return PrecheckResultResultEnum{
		PASSED: PrecheckResultResult{
			value: "PASSED",
		},
		ALARM: PrecheckResultResult{
			value: "ALARM",
		},
		FAILED: PrecheckResultResult{
			value: "FAILED",
		},
	}
}

func (c PrecheckResultResult) Value() string {
	return c.value
}

func (c PrecheckResultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrecheckResultResult) UnmarshalJSON(b []byte) error {
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
