package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ShowMonitoringDataRequest struct {
	XLanguage *ShowMonitoringDataRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryJobReq `json:"body,omitempty"`
}

func (o ShowMonitoringDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitoringDataRequest struct{}"
	}

	return strings.Join([]string{"ShowMonitoringDataRequest", string(data)}, " ")
}

type ShowMonitoringDataRequestXLanguage struct {
	value string
}

type ShowMonitoringDataRequestXLanguageEnum struct {
	EN_US ShowMonitoringDataRequestXLanguage
	ZH_CN ShowMonitoringDataRequestXLanguage
}

func GetShowMonitoringDataRequestXLanguageEnum() ShowMonitoringDataRequestXLanguageEnum {
	return ShowMonitoringDataRequestXLanguageEnum{
		EN_US: ShowMonitoringDataRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowMonitoringDataRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowMonitoringDataRequestXLanguage) Value() string {
	return c.value
}

func (c ShowMonitoringDataRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMonitoringDataRequestXLanguage) UnmarshalJSON(b []byte) error {
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
