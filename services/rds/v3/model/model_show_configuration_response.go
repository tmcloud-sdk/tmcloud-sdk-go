package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ShowConfigurationResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`

	DatastoreName *ShowConfigurationResponseDatastoreName `json:"datastore_name,omitempty"`

	Created *string `json:"created,omitempty"`

	Updated *string `json:"updated,omitempty"`

	ConfigurationParameters *[]ConfigurationParameter `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                       `json:"-"`
}

func (o ShowConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationResponse", string(data)}, " ")
}

type ShowConfigurationResponseDatastoreName struct {
	value string
}

type ShowConfigurationResponseDatastoreNameEnum struct {
	MYSQL      ShowConfigurationResponseDatastoreName
	POSTGRESQL ShowConfigurationResponseDatastoreName
	SQLSERVER  ShowConfigurationResponseDatastoreName
}

func GetShowConfigurationResponseDatastoreNameEnum() ShowConfigurationResponseDatastoreNameEnum {
	return ShowConfigurationResponseDatastoreNameEnum{
		MYSQL: ShowConfigurationResponseDatastoreName{
			value: "mysql",
		},
		POSTGRESQL: ShowConfigurationResponseDatastoreName{
			value: "postgresql",
		},
		SQLSERVER: ShowConfigurationResponseDatastoreName{
			value: "sqlserver",
		},
	}
}

func (c ShowConfigurationResponseDatastoreName) Value() string {
	return c.value
}

func (c ShowConfigurationResponseDatastoreName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigurationResponseDatastoreName) UnmarshalJSON(b []byte) error {
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
