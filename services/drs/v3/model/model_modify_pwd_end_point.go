package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ModifyPwdEndPoint struct {
	DbPassword string `json:"db_password"`

	EndPointType ModifyPwdEndPointEndPointType `json:"end_point_type"`

	JobId string `json:"job_id"`

	Kerberos *KerberosVo `json:"kerberos,omitempty"`
}

func (o ModifyPwdEndPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPwdEndPoint struct{}"
	}

	return strings.Join([]string{"ModifyPwdEndPoint", string(data)}, " ")
}

type ModifyPwdEndPointEndPointType struct {
	value string
}

type ModifyPwdEndPointEndPointTypeEnum struct {
	SO ModifyPwdEndPointEndPointType
	TA ModifyPwdEndPointEndPointType
}

func GetModifyPwdEndPointEndPointTypeEnum() ModifyPwdEndPointEndPointTypeEnum {
	return ModifyPwdEndPointEndPointTypeEnum{
		SO: ModifyPwdEndPointEndPointType{
			value: "so",
		},
		TA: ModifyPwdEndPointEndPointType{
			value: "ta",
		},
	}
}

func (c ModifyPwdEndPointEndPointType) Value() string {
	return c.value
}

func (c ModifyPwdEndPointEndPointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyPwdEndPointEndPointType) UnmarshalJSON(b []byte) error {
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
