package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ShowInstanceConfigurationResponse struct {
	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`

	DatastoreName *ShowInstanceConfigurationResponseDatastoreName `json:"datastore_name,omitempty"`

	Created *string `json:"created,omitempty"`

	Updated *string `json:"updated,omitempty"`

	ConfigurationParameters *[]ConfigurationParameter `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                       `json:"-"`
}

func (o ShowInstanceConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationResponse", string(data)}, " ")
}

type ShowInstanceConfigurationResponseDatastoreName struct {
	value string
}

type ShowInstanceConfigurationResponseDatastoreNameEnum struct {
	MYSQL      ShowInstanceConfigurationResponseDatastoreName
	POSTGRESQL ShowInstanceConfigurationResponseDatastoreName
	SQLSERVER  ShowInstanceConfigurationResponseDatastoreName
}

func GetShowInstanceConfigurationResponseDatastoreNameEnum() ShowInstanceConfigurationResponseDatastoreNameEnum {
	return ShowInstanceConfigurationResponseDatastoreNameEnum{
		MYSQL: ShowInstanceConfigurationResponseDatastoreName{
			value: "mysql",
		},
		POSTGRESQL: ShowInstanceConfigurationResponseDatastoreName{
			value: "postgresql",
		},
		SQLSERVER: ShowInstanceConfigurationResponseDatastoreName{
			value: "sqlserver",
		},
	}
}

func (c ShowInstanceConfigurationResponseDatastoreName) Value() string {
	return c.value
}

func (c ShowInstanceConfigurationResponseDatastoreName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceConfigurationResponseDatastoreName) UnmarshalJSON(b []byte) error {
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
