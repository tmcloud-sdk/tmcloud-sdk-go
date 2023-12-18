package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BigkeysBody struct {
	Name *string `json:"name,omitempty"`

	Type *BigkeysBodyType `json:"type,omitempty"`

	Shard *string `json:"shard,omitempty"`

	Db *int32 `json:"db,omitempty"`

	Size *int64 `json:"size,omitempty"`

	Unit *string `json:"unit,omitempty"`
}

func (o BigkeysBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BigkeysBody struct{}"
	}

	return strings.Join([]string{"BigkeysBody", string(data)}, " ")
}

type BigkeysBodyType struct {
	value string
}

type BigkeysBodyTypeEnum struct {
	STRING BigkeysBodyType
	LIST   BigkeysBodyType
	SET    BigkeysBodyType
	ZSET   BigkeysBodyType
	HASH   BigkeysBodyType
}

func GetBigkeysBodyTypeEnum() BigkeysBodyTypeEnum {
	return BigkeysBodyTypeEnum{
		STRING: BigkeysBodyType{
			value: "string",
		},
		LIST: BigkeysBodyType{
			value: "list",
		},
		SET: BigkeysBodyType{
			value: "set",
		},
		ZSET: BigkeysBodyType{
			value: "zset",
		},
		HASH: BigkeysBodyType{
			value: "hash",
		},
	}
}

func (c BigkeysBodyType) Value() string {
	return c.value
}

func (c BigkeysBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BigkeysBodyType) UnmarshalJSON(b []byte) error {
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
