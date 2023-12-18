package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DatabaseInfo struct {
	Id *string `json:"id,omitempty"`

	ParentId *string `json:"parent_id,omitempty"`

	ObjectType *DatabaseInfoObjectType `json:"object_type,omitempty"`

	ObjectName *string `json:"object_name,omitempty"`

	ObjectAliasName *string `json:"object_alias_name,omitempty"`

	Select *string `json:"select,omitempty"`
}

func (o DatabaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseInfo struct{}"
	}

	return strings.Join([]string{"DatabaseInfo", string(data)}, " ")
}

type DatabaseInfoObjectType struct {
	value string
}

type DatabaseInfoObjectTypeEnum struct {
	DATABASE DatabaseInfoObjectType
	TABLE    DatabaseInfoObjectType
	SCHEMA   DatabaseInfoObjectType
	VIEW     DatabaseInfoObjectType
}

func GetDatabaseInfoObjectTypeEnum() DatabaseInfoObjectTypeEnum {
	return DatabaseInfoObjectTypeEnum{
		DATABASE: DatabaseInfoObjectType{
			value: "database",
		},
		TABLE: DatabaseInfoObjectType{
			value: "table",
		},
		SCHEMA: DatabaseInfoObjectType{
			value: "schema",
		},
		VIEW: DatabaseInfoObjectType{
			value: "view",
		},
	}
}

func (c DatabaseInfoObjectType) Value() string {
	return c.value
}

func (c DatabaseInfoObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseInfoObjectType) UnmarshalJSON(b []byte) error {
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
