package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type StartInstanceEnlargeVolumeActionRequest struct {
	XLanguage *StartInstanceEnlargeVolumeActionRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *EnlargeVolume `json:"body,omitempty"`
}

func (o StartInstanceEnlargeVolumeActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceEnlargeVolumeActionRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceEnlargeVolumeActionRequest", string(data)}, " ")
}

type StartInstanceEnlargeVolumeActionRequestXLanguage struct {
	value string
}

type StartInstanceEnlargeVolumeActionRequestXLanguageEnum struct {
	ZH_CN StartInstanceEnlargeVolumeActionRequestXLanguage
	EN_US StartInstanceEnlargeVolumeActionRequestXLanguage
}

func GetStartInstanceEnlargeVolumeActionRequestXLanguageEnum() StartInstanceEnlargeVolumeActionRequestXLanguageEnum {
	return StartInstanceEnlargeVolumeActionRequestXLanguageEnum{
		ZH_CN: StartInstanceEnlargeVolumeActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartInstanceEnlargeVolumeActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartInstanceEnlargeVolumeActionRequestXLanguage) Value() string {
	return c.value
}

func (c StartInstanceEnlargeVolumeActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartInstanceEnlargeVolumeActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
