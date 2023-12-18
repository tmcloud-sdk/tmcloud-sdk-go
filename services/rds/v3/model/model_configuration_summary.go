package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ConfigurationSummary struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	DatastoreVersionName string `json:"datastore_version_name"`

	DatastoreName ConfigurationSummaryDatastoreName `json:"datastore_name"`

	Created string `json:"created"`

	Updated string `json:"updated"`

	UserDefined bool `json:"user_defined"`
}

func (o ConfigurationSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationSummary struct{}"
	}

	return strings.Join([]string{"ConfigurationSummary", string(data)}, " ")
}

type ConfigurationSummaryDatastoreName struct {
	value string
}

type ConfigurationSummaryDatastoreNameEnum struct {
	MYSQL      ConfigurationSummaryDatastoreName
	POSTGRESQL ConfigurationSummaryDatastoreName
	SQLSERVER  ConfigurationSummaryDatastoreName
}

func GetConfigurationSummaryDatastoreNameEnum() ConfigurationSummaryDatastoreNameEnum {
	return ConfigurationSummaryDatastoreNameEnum{
		MYSQL: ConfigurationSummaryDatastoreName{
			value: "mysql",
		},
		POSTGRESQL: ConfigurationSummaryDatastoreName{
			value: "postgresql",
		},
		SQLSERVER: ConfigurationSummaryDatastoreName{
			value: "sqlserver",
		},
	}
}

func (c ConfigurationSummaryDatastoreName) Value() string {
	return c.value
}

func (c ConfigurationSummaryDatastoreName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationSummaryDatastoreName) UnmarshalJSON(b []byte) error {
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
