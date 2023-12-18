package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ModifyTargetParamsReq struct {
	Group ModifyTargetParamsReqGroup `json:"group"`

	Params []ParamsReqBean `json:"params"`
}

func (o ModifyTargetParamsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTargetParamsReq struct{}"
	}

	return strings.Join([]string{"ModifyTargetParamsReq", string(data)}, " ")
}

type ModifyTargetParamsReqGroup struct {
	value string
}

type ModifyTargetParamsReqGroupEnum struct {
	COMMON      ModifyTargetParamsReqGroup
	PERFORMANCE ModifyTargetParamsReqGroup
}

func GetModifyTargetParamsReqGroupEnum() ModifyTargetParamsReqGroupEnum {
	return ModifyTargetParamsReqGroupEnum{
		COMMON: ModifyTargetParamsReqGroup{
			value: "common",
		},
		PERFORMANCE: ModifyTargetParamsReqGroup{
			value: "performance",
		},
	}
}

func (c ModifyTargetParamsReqGroup) Value() string {
	return c.value
}

func (c ModifyTargetParamsReqGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyTargetParamsReqGroup) UnmarshalJSON(b []byte) error {
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
