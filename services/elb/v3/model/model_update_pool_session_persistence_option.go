package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdatePoolSessionPersistenceOption struct {
	CookieName *string `json:"cookie_name,omitempty"`

	Type *UpdatePoolSessionPersistenceOptionType `json:"type,omitempty"`

	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o UpdatePoolSessionPersistenceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolSessionPersistenceOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolSessionPersistenceOption", string(data)}, " ")
}

type UpdatePoolSessionPersistenceOptionType struct {
	value string
}

type UpdatePoolSessionPersistenceOptionTypeEnum struct {
	SOURCE_IP   UpdatePoolSessionPersistenceOptionType
	HTTP_COOKIE UpdatePoolSessionPersistenceOptionType
	APP_COOKIE  UpdatePoolSessionPersistenceOptionType
}

func GetUpdatePoolSessionPersistenceOptionTypeEnum() UpdatePoolSessionPersistenceOptionTypeEnum {
	return UpdatePoolSessionPersistenceOptionTypeEnum{
		SOURCE_IP: UpdatePoolSessionPersistenceOptionType{
			value: "SOURCE_IP",
		},
		HTTP_COOKIE: UpdatePoolSessionPersistenceOptionType{
			value: "HTTP_COOKIE",
		},
		APP_COOKIE: UpdatePoolSessionPersistenceOptionType{
			value: "APP_COOKIE",
		},
	}
}

func (c UpdatePoolSessionPersistenceOptionType) Value() string {
	return c.value
}

func (c UpdatePoolSessionPersistenceOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePoolSessionPersistenceOptionType) UnmarshalJSON(b []byte) error {
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
