package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateImageRequestBody struct {
	Op UpdateImageRequestBodyOp `json:"op"`

	Path string `json:"path"`

	Value string `json:"value"`
}

func (o UpdateImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateImageRequestBody", string(data)}, " ")
}

type UpdateImageRequestBodyOp struct {
	value string
}

type UpdateImageRequestBodyOpEnum struct {
	ADD     UpdateImageRequestBodyOp
	REPLACE UpdateImageRequestBodyOp
	REMOVE  UpdateImageRequestBodyOp
}

func GetUpdateImageRequestBodyOpEnum() UpdateImageRequestBodyOpEnum {
	return UpdateImageRequestBodyOpEnum{
		ADD: UpdateImageRequestBodyOp{
			value: "add",
		},
		REPLACE: UpdateImageRequestBodyOp{
			value: "replace",
		},
		REMOVE: UpdateImageRequestBodyOp{
			value: "remove",
		},
	}
}

func (c UpdateImageRequestBodyOp) Value() string {
	return c.value
}

func (c UpdateImageRequestBodyOp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageRequestBodyOp) UnmarshalJSON(b []byte) error {
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
