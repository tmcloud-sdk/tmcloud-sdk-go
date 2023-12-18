package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListAvailableZoneRequest struct {
	XLanguage *ListAvailableZoneRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryAvailableNodeTypeReq `json:"body,omitempty"`
}

func (o ListAvailableZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZoneRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableZoneRequest", string(data)}, " ")
}

type ListAvailableZoneRequestXLanguage struct {
	value string
}

type ListAvailableZoneRequestXLanguageEnum struct {
	EN_US ListAvailableZoneRequestXLanguage
	ZH_CN ListAvailableZoneRequestXLanguage
}

func GetListAvailableZoneRequestXLanguageEnum() ListAvailableZoneRequestXLanguageEnum {
	return ListAvailableZoneRequestXLanguageEnum{
		EN_US: ListAvailableZoneRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListAvailableZoneRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListAvailableZoneRequestXLanguage) Value() string {
	return c.value
}

func (c ListAvailableZoneRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAvailableZoneRequestXLanguage) UnmarshalJSON(b []byte) error {
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
