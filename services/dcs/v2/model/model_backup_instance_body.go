package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BackupInstanceBody struct {
	Remark *string `json:"remark,omitempty"`

	BackupFormat *BackupInstanceBodyBackupFormat `json:"backup_format,omitempty"`
}

func (o BackupInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInstanceBody struct{}"
	}

	return strings.Join([]string{"BackupInstanceBody", string(data)}, " ")
}

type BackupInstanceBodyBackupFormat struct {
	value string
}

type BackupInstanceBodyBackupFormatEnum struct {
	AOF BackupInstanceBodyBackupFormat
	RDB BackupInstanceBodyBackupFormat
}

func GetBackupInstanceBodyBackupFormatEnum() BackupInstanceBodyBackupFormatEnum {
	return BackupInstanceBodyBackupFormatEnum{
		AOF: BackupInstanceBodyBackupFormat{
			value: "aof",
		},
		RDB: BackupInstanceBodyBackupFormat{
			value: "rdb",
		},
	}
}

func (c BackupInstanceBodyBackupFormat) Value() string {
	return c.value
}

func (c BackupInstanceBodyBackupFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupInstanceBodyBackupFormat) UnmarshalJSON(b []byte) error {
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
