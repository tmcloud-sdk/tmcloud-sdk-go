package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GlanceUpdateImageRequestBody struct {
	Op GlanceUpdateImageRequestBodyOp `json:"op"`

	Path string `json:"path"`

	Value *string `json:"value,omitempty"`
}

func (o GlanceUpdateImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageRequestBody", string(data)}, " ")
}

type GlanceUpdateImageRequestBodyOp struct {
	value string
}

type GlanceUpdateImageRequestBodyOpEnum struct {
	REPLACE GlanceUpdateImageRequestBodyOp
	ADD     GlanceUpdateImageRequestBodyOp
	REMOVE  GlanceUpdateImageRequestBodyOp
}

func GetGlanceUpdateImageRequestBodyOpEnum() GlanceUpdateImageRequestBodyOpEnum {
	return GlanceUpdateImageRequestBodyOpEnum{
		REPLACE: GlanceUpdateImageRequestBodyOp{
			value: "replace",
		},
		ADD: GlanceUpdateImageRequestBodyOp{
			value: "add",
		},
		REMOVE: GlanceUpdateImageRequestBodyOp{
			value: "remove",
		},
	}
}

func (c GlanceUpdateImageRequestBodyOp) Value() string {
	return c.value
}

func (c GlanceUpdateImageRequestBodyOp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageRequestBodyOp) UnmarshalJSON(b []byte) error {
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
