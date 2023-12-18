package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateFixtedResponseConfig struct {
	StatusCode *string `json:"status_code,omitempty"`

	ContentType *UpdateFixtedResponseConfigContentType `json:"content_type,omitempty"`

	MessageBody *string `json:"message_body,omitempty"`
}

func (o UpdateFixtedResponseConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFixtedResponseConfig struct{}"
	}

	return strings.Join([]string{"UpdateFixtedResponseConfig", string(data)}, " ")
}

type UpdateFixtedResponseConfigContentType struct {
	value string
}

type UpdateFixtedResponseConfigContentTypeEnum struct {
	TEXT_PLAIN             UpdateFixtedResponseConfigContentType
	TEXT_CSS               UpdateFixtedResponseConfigContentType
	TEXT_HTML              UpdateFixtedResponseConfigContentType
	APPLICATION_JAVASCRIPT UpdateFixtedResponseConfigContentType
	APPLICATION_JSON       UpdateFixtedResponseConfigContentType
}

func GetUpdateFixtedResponseConfigContentTypeEnum() UpdateFixtedResponseConfigContentTypeEnum {
	return UpdateFixtedResponseConfigContentTypeEnum{
		TEXT_PLAIN: UpdateFixtedResponseConfigContentType{
			value: "text/plain",
		},
		TEXT_CSS: UpdateFixtedResponseConfigContentType{
			value: "text/css",
		},
		TEXT_HTML: UpdateFixtedResponseConfigContentType{
			value: "text/html",
		},
		APPLICATION_JAVASCRIPT: UpdateFixtedResponseConfigContentType{
			value: "application/javascript",
		},
		APPLICATION_JSON: UpdateFixtedResponseConfigContentType{
			value: "application/json",
		},
	}
}

func (c UpdateFixtedResponseConfigContentType) Value() string {
	return c.value
}

func (c UpdateFixtedResponseConfigContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFixtedResponseConfigContentType) UnmarshalJSON(b []byte) error {
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
