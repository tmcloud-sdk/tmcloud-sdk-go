package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type FixtedResponseConfig struct {
	StatusCode string `json:"status_code"`

	ContentType FixtedResponseConfigContentType `json:"content_type"`

	MessageBody string `json:"message_body"`
}

func (o FixtedResponseConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FixtedResponseConfig struct{}"
	}

	return strings.Join([]string{"FixtedResponseConfig", string(data)}, " ")
}

type FixtedResponseConfigContentType struct {
	value string
}

type FixtedResponseConfigContentTypeEnum struct {
	TEXT_PLAIN             FixtedResponseConfigContentType
	TEXT_CSS               FixtedResponseConfigContentType
	TEXT_HTML              FixtedResponseConfigContentType
	APPLICATION_JAVASCRIPT FixtedResponseConfigContentType
	APPLICATION_JSON       FixtedResponseConfigContentType
}

func GetFixtedResponseConfigContentTypeEnum() FixtedResponseConfigContentTypeEnum {
	return FixtedResponseConfigContentTypeEnum{
		TEXT_PLAIN: FixtedResponseConfigContentType{
			value: "text/plain",
		},
		TEXT_CSS: FixtedResponseConfigContentType{
			value: "text/css",
		},
		TEXT_HTML: FixtedResponseConfigContentType{
			value: "text/html",
		},
		APPLICATION_JAVASCRIPT: FixtedResponseConfigContentType{
			value: "application/javascript",
		},
		APPLICATION_JSON: FixtedResponseConfigContentType{
			value: "application/json",
		},
	}
}

func (c FixtedResponseConfigContentType) Value() string {
	return c.value
}

func (c FixtedResponseConfigContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FixtedResponseConfigContentType) UnmarshalJSON(b []byte) error {
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
