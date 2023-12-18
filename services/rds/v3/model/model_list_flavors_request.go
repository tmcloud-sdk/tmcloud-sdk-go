package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListFlavorsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	DatabaseName ListFlavorsRequestDatabaseName `json:"database_name"`

	VersionName *string `json:"version_name,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

type ListFlavorsRequestDatabaseName struct {
	value string
}

type ListFlavorsRequestDatabaseNameEnum struct {
	MY_SQL      ListFlavorsRequestDatabaseName
	POSTGRE_SQL ListFlavorsRequestDatabaseName
	SQL_SERVER  ListFlavorsRequestDatabaseName
}

func GetListFlavorsRequestDatabaseNameEnum() ListFlavorsRequestDatabaseNameEnum {
	return ListFlavorsRequestDatabaseNameEnum{
		MY_SQL: ListFlavorsRequestDatabaseName{
			value: "MySQL",
		},
		POSTGRE_SQL: ListFlavorsRequestDatabaseName{
			value: "PostgreSQL",
		},
		SQL_SERVER: ListFlavorsRequestDatabaseName{
			value: "SQLServer",
		},
	}
}

func (c ListFlavorsRequestDatabaseName) Value() string {
	return c.value
}

func (c ListFlavorsRequestDatabaseName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestDatabaseName) UnmarshalJSON(b []byte) error {
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
