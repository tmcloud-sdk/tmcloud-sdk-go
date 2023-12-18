package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateFixtedResponseConfig struct {
	StatusCode string `json:"status_code"`

	ContentType *CreateFixtedResponseConfigContentType `json:"content_type,omitempty"`

	MessageBody *string `json:"message_body,omitempty"`
}

func (o CreateFixtedResponseConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFixtedResponseConfig struct{}"
	}

	return strings.Join([]string{"CreateFixtedResponseConfig", string(data)}, " ")
}

type CreateFixtedResponseConfigContentType struct {
	value string
}

type CreateFixtedResponseConfigContentTypeEnum struct {
	TEXT_PLAIN             CreateFixtedResponseConfigContentType
	TEXT_CSS               CreateFixtedResponseConfigContentType
	TEXT_HTML              CreateFixtedResponseConfigContentType
	APPLICATION_JAVASCRIPT CreateFixtedResponseConfigContentType
	APPLICATION_JSON       CreateFixtedResponseConfigContentType
}

func GetCreateFixtedResponseConfigContentTypeEnum() CreateFixtedResponseConfigContentTypeEnum {
	return CreateFixtedResponseConfigContentTypeEnum{
		TEXT_PLAIN: CreateFixtedResponseConfigContentType{
			value: "text/plain",
		},
		TEXT_CSS: CreateFixtedResponseConfigContentType{
			value: "text/css",
		},
		TEXT_HTML: CreateFixtedResponseConfigContentType{
			value: "text/html",
		},
		APPLICATION_JAVASCRIPT: CreateFixtedResponseConfigContentType{
			value: "application/javascript",
		},
		APPLICATION_JSON: CreateFixtedResponseConfigContentType{
			value: "application/json",
		},
	}
}

func (c CreateFixtedResponseConfigContentType) Value() string {
	return c.value
}

func (c CreateFixtedResponseConfigContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFixtedResponseConfigContentType) UnmarshalJSON(b []byte) error {
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
