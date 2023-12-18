package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type Datastore struct {
	Type DatastoreType `json:"type"`

	Version string `json:"version"`

	CompleteVersion *string `json:"complete_version,omitempty"`
}

func (o Datastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datastore struct{}"
	}

	return strings.Join([]string{"Datastore", string(data)}, " ")
}

type DatastoreType struct {
	value string
}

type DatastoreTypeEnum struct {
	MY_SQL      DatastoreType
	POSTGRE_SQL DatastoreType
	SQL_SERVER  DatastoreType
}

func GetDatastoreTypeEnum() DatastoreTypeEnum {
	return DatastoreTypeEnum{
		MY_SQL: DatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: DatastoreType{
			value: "PostgreSQL",
		},
		SQL_SERVER: DatastoreType{
			value: "SQLServer",
		},
	}
}

func (c DatastoreType) Value() string {
	return c.value
}

func (c DatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatastoreType) UnmarshalJSON(b []byte) error {
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
