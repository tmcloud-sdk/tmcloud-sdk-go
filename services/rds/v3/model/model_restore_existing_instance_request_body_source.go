package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type RestoreExistingInstanceRequestBodySource struct {
	InstanceId string `json:"instance_id"`

	Type *RestoreExistingInstanceRequestBodySourceType `json:"type,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	RestoreTime *int64 `json:"restore_time,omitempty"`

	DatabaseName map[string]string `json:"database_name,omitempty"`

	RestoreAllDatabase *bool `json:"restore_all_database,omitempty"`
}

func (o RestoreExistingInstanceRequestBodySource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistingInstanceRequestBodySource struct{}"
	}

	return strings.Join([]string{"RestoreExistingInstanceRequestBodySource", string(data)}, " ")
}

type RestoreExistingInstanceRequestBodySourceType struct {
	value string
}

type RestoreExistingInstanceRequestBodySourceTypeEnum struct {
	BACKUP    RestoreExistingInstanceRequestBodySourceType
	TIMESTAMP RestoreExistingInstanceRequestBodySourceType
}

func GetRestoreExistingInstanceRequestBodySourceTypeEnum() RestoreExistingInstanceRequestBodySourceTypeEnum {
	return RestoreExistingInstanceRequestBodySourceTypeEnum{
		BACKUP: RestoreExistingInstanceRequestBodySourceType{
			value: "backup",
		},
		TIMESTAMP: RestoreExistingInstanceRequestBodySourceType{
			value: "timestamp",
		},
	}
}

func (c RestoreExistingInstanceRequestBodySourceType) Value() string {
	return c.value
}

func (c RestoreExistingInstanceRequestBodySourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestoreExistingInstanceRequestBodySourceType) UnmarshalJSON(b []byte) error {
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
