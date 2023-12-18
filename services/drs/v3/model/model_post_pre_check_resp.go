package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PostPreCheckResp struct {
	Id *string `json:"id,omitempty"`

	PrecheckId *string `json:"precheck_id,omitempty"`

	Status *PostPreCheckRespStatus `json:"status,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o PostPreCheckResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPreCheckResp struct{}"
	}

	return strings.Join([]string{"PostPreCheckResp", string(data)}, " ")
}

type PostPreCheckRespStatus struct {
	value string
}

type PostPreCheckRespStatusEnum struct {
	SUCCESS PostPreCheckRespStatus
	FAILED  PostPreCheckRespStatus
}

func GetPostPreCheckRespStatusEnum() PostPreCheckRespStatusEnum {
	return PostPreCheckRespStatusEnum{
		SUCCESS: PostPreCheckRespStatus{
			value: "success",
		},
		FAILED: PostPreCheckRespStatus{
			value: "failed",
		},
	}
}

func (c PostPreCheckRespStatus) Value() string {
	return c.value
}

func (c PostPreCheckRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPreCheckRespStatus) UnmarshalJSON(b []byte) error {
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
