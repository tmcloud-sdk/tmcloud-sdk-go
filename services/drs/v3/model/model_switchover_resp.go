package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SwitchoverResp struct {
	JobId *string `json:"job_id,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	SourceDb *EndpointVo `json:"source_db,omitempty"`

	TargetDb *EndpointVo `json:"target_db,omitempty"`

	JobDirection *SwitchoverRespJobDirection `json:"job_direction,omitempty"`

	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`
}

func (o SwitchoverResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverResp struct{}"
	}

	return strings.Join([]string{"SwitchoverResp", string(data)}, " ")
}

type SwitchoverRespJobDirection struct {
	value string
}

type SwitchoverRespJobDirectionEnum struct {
	UP      SwitchoverRespJobDirection
	DOWN    SwitchoverRespJobDirection
	NON_DBS SwitchoverRespJobDirection
}

func GetSwitchoverRespJobDirectionEnum() SwitchoverRespJobDirectionEnum {
	return SwitchoverRespJobDirectionEnum{
		UP: SwitchoverRespJobDirection{
			value: "up",
		},
		DOWN: SwitchoverRespJobDirection{
			value: "down",
		},
		NON_DBS: SwitchoverRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c SwitchoverRespJobDirection) Value() string {
	return c.value
}

func (c SwitchoverRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchoverRespJobDirection) UnmarshalJSON(b []byte) error {
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
