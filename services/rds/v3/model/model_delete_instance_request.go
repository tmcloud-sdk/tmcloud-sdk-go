package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteInstanceRequest struct {
	XLanguage *DeleteInstanceRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o DeleteInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRequest", string(data)}, " ")
}

type DeleteInstanceRequestXLanguage struct {
	value string
}

type DeleteInstanceRequestXLanguageEnum struct {
	ZH_CN DeleteInstanceRequestXLanguage
	EN_US DeleteInstanceRequestXLanguage
}

func GetDeleteInstanceRequestXLanguageEnum() DeleteInstanceRequestXLanguageEnum {
	return DeleteInstanceRequestXLanguageEnum{
		ZH_CN: DeleteInstanceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteInstanceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteInstanceRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteInstanceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteInstanceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
