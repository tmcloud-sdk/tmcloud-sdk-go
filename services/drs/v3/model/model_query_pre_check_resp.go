package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type QueryPreCheckResp struct {
	PrecheckId *string `json:"precheck_id,omitempty"`

	Result *bool `json:"result,omitempty"`

	Process *string `json:"process,omitempty"`

	TotalPassedRate *string `json:"total_passed_rate,omitempty"`

	RdsInstanceId *string `json:"rds_instance_id,omitempty"`

	JobDirection *QueryPreCheckRespJobDirection `json:"job_direction,omitempty"`

	PrecheckResult *[]PrecheckResult `json:"precheck_result,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`
}

func (o QueryPreCheckResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryPreCheckResp struct{}"
	}

	return strings.Join([]string{"QueryPreCheckResp", string(data)}, " ")
}

type QueryPreCheckRespJobDirection struct {
	value string
}

type QueryPreCheckRespJobDirectionEnum struct {
	UP      QueryPreCheckRespJobDirection
	DOWN    QueryPreCheckRespJobDirection
	NON_DBS QueryPreCheckRespJobDirection
}

func GetQueryPreCheckRespJobDirectionEnum() QueryPreCheckRespJobDirectionEnum {
	return QueryPreCheckRespJobDirectionEnum{
		UP: QueryPreCheckRespJobDirection{
			value: "up",
		},
		DOWN: QueryPreCheckRespJobDirection{
			value: "down",
		},
		NON_DBS: QueryPreCheckRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c QueryPreCheckRespJobDirection) Value() string {
	return c.value
}

func (c QueryPreCheckRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryPreCheckRespJobDirection) UnmarshalJSON(b []byte) error {
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
