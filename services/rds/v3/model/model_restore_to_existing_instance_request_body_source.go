package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type RestoreToExistingInstanceRequestBodySource struct {
	InstanceId string `json:"instance_id"`

	Type *RestoreToExistingInstanceRequestBodySourceType `json:"type,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	RestoreTime *int64 `json:"restore_time,omitempty"`

	DatabaseName map[string]string `json:"database_name,omitempty"`
}

func (o RestoreToExistingInstanceRequestBodySource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreToExistingInstanceRequestBodySource struct{}"
	}

	return strings.Join([]string{"RestoreToExistingInstanceRequestBodySource", string(data)}, " ")
}

type RestoreToExistingInstanceRequestBodySourceType struct {
	value string
}

type RestoreToExistingInstanceRequestBodySourceTypeEnum struct {
	BACKUP    RestoreToExistingInstanceRequestBodySourceType
	TIMESTAMP RestoreToExistingInstanceRequestBodySourceType
}

func GetRestoreToExistingInstanceRequestBodySourceTypeEnum() RestoreToExistingInstanceRequestBodySourceTypeEnum {
	return RestoreToExistingInstanceRequestBodySourceTypeEnum{
		BACKUP: RestoreToExistingInstanceRequestBodySourceType{
			value: "backup",
		},
		TIMESTAMP: RestoreToExistingInstanceRequestBodySourceType{
			value: "timestamp",
		},
	}
}

func (c RestoreToExistingInstanceRequestBodySourceType) Value() string {
	return c.value
}

func (c RestoreToExistingInstanceRequestBodySourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestoreToExistingInstanceRequestBodySourceType) UnmarshalJSON(b []byte) error {
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
