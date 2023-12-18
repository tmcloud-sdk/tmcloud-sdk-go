package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteJobReq struct {
	DeleteType DeleteJobReqDeleteType `json:"delete_type"`

	JobId string `json:"job_id"`
}

func (o DeleteJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobReq struct{}"
	}

	return strings.Join([]string{"DeleteJobReq", string(data)}, " ")
}

type DeleteJobReqDeleteType struct {
	value string
}

type DeleteJobReqDeleteTypeEnum struct {
	TERMINATE       DeleteJobReqDeleteType
	FORCE_TERMINATE DeleteJobReqDeleteType
	DELETE          DeleteJobReqDeleteType
}

func GetDeleteJobReqDeleteTypeEnum() DeleteJobReqDeleteTypeEnum {
	return DeleteJobReqDeleteTypeEnum{
		TERMINATE: DeleteJobReqDeleteType{
			value: "terminate",
		},
		FORCE_TERMINATE: DeleteJobReqDeleteType{
			value: "force_terminate",
		},
		DELETE: DeleteJobReqDeleteType{
			value: "delete",
		},
	}
}

func (c DeleteJobReqDeleteType) Value() string {
	return c.value
}

func (c DeleteJobReqDeleteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteJobReqDeleteType) UnmarshalJSON(b []byte) error {
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
