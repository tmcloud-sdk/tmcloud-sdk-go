package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreatePoolSessionPersistenceOption struct {
	CookieName *string `json:"cookie_name,omitempty"`

	Type CreatePoolSessionPersistenceOptionType `json:"type"`

	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o CreatePoolSessionPersistenceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolSessionPersistenceOption struct{}"
	}

	return strings.Join([]string{"CreatePoolSessionPersistenceOption", string(data)}, " ")
}

type CreatePoolSessionPersistenceOptionType struct {
	value string
}

type CreatePoolSessionPersistenceOptionTypeEnum struct {
	SOURCE_IP   CreatePoolSessionPersistenceOptionType
	HTTP_COOKIE CreatePoolSessionPersistenceOptionType
	APP_COOKIE  CreatePoolSessionPersistenceOptionType
}

func GetCreatePoolSessionPersistenceOptionTypeEnum() CreatePoolSessionPersistenceOptionTypeEnum {
	return CreatePoolSessionPersistenceOptionTypeEnum{
		SOURCE_IP: CreatePoolSessionPersistenceOptionType{
			value: "SOURCE_IP",
		},
		HTTP_COOKIE: CreatePoolSessionPersistenceOptionType{
			value: "HTTP_COOKIE",
		},
		APP_COOKIE: CreatePoolSessionPersistenceOptionType{
			value: "APP_COOKIE",
		},
	}
}

func (c CreatePoolSessionPersistenceOptionType) Value() string {
	return c.value
}

func (c CreatePoolSessionPersistenceOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePoolSessionPersistenceOptionType) UnmarshalJSON(b []byte) error {
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
