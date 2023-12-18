package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ConfigurationSummaryForCreate struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	DatastoreVersionName string `json:"datastore_version_name"`

	DatastoreName ConfigurationSummaryForCreateDatastoreName `json:"datastore_name"`

	Created string `json:"created"`

	Updated string `json:"updated"`
}

func (o ConfigurationSummaryForCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationSummaryForCreate struct{}"
	}

	return strings.Join([]string{"ConfigurationSummaryForCreate", string(data)}, " ")
}

type ConfigurationSummaryForCreateDatastoreName struct {
	value string
}

type ConfigurationSummaryForCreateDatastoreNameEnum struct {
	MYSQL      ConfigurationSummaryForCreateDatastoreName
	POSTGRESQL ConfigurationSummaryForCreateDatastoreName
	SQLSERVER  ConfigurationSummaryForCreateDatastoreName
}

func GetConfigurationSummaryForCreateDatastoreNameEnum() ConfigurationSummaryForCreateDatastoreNameEnum {
	return ConfigurationSummaryForCreateDatastoreNameEnum{
		MYSQL: ConfigurationSummaryForCreateDatastoreName{
			value: "mysql",
		},
		POSTGRESQL: ConfigurationSummaryForCreateDatastoreName{
			value: "postgresql",
		},
		SQLSERVER: ConfigurationSummaryForCreateDatastoreName{
			value: "sqlserver",
		},
	}
}

func (c ConfigurationSummaryForCreateDatastoreName) Value() string {
	return c.value
}

func (c ConfigurationSummaryForCreateDatastoreName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationSummaryForCreateDatastoreName) UnmarshalJSON(b []byte) error {
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
