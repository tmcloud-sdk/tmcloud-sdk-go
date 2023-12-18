package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type HotkeysBody struct {
	Name *string `json:"name,omitempty"`

	Type *HotkeysBodyType `json:"type,omitempty"`

	Shard *string `json:"shard,omitempty"`

	Db *int32 `json:"db,omitempty"`

	Size *int64 `json:"size,omitempty"`

	Unit *string `json:"unit,omitempty"`

	Freq *int32 `json:"freq,omitempty"`
}

func (o HotkeysBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotkeysBody struct{}"
	}

	return strings.Join([]string{"HotkeysBody", string(data)}, " ")
}

type HotkeysBodyType struct {
	value string
}

type HotkeysBodyTypeEnum struct {
	STRING HotkeysBodyType
	LIST   HotkeysBodyType
	SET    HotkeysBodyType
	ZSET   HotkeysBodyType
	HASH   HotkeysBodyType
}

func GetHotkeysBodyTypeEnum() HotkeysBodyTypeEnum {
	return HotkeysBodyTypeEnum{
		STRING: HotkeysBodyType{
			value: "string",
		},
		LIST: HotkeysBodyType{
			value: "list",
		},
		SET: HotkeysBodyType{
			value: "set",
		},
		ZSET: HotkeysBodyType{
			value: "zset",
		},
		HASH: HotkeysBodyType{
			value: "hash",
		},
	}
}

func (c HotkeysBodyType) Value() string {
	return c.value
}

func (c HotkeysBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HotkeysBodyType) UnmarshalJSON(b []byte) error {
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
