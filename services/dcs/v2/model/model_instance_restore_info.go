package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type InstanceRestoreInfo struct {
	BackupId *string `json:"backup_id,omitempty"`

	RestoreId *string `json:"restore_id,omitempty"`

	BackupName *string `json:"backup_name,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	RestoreRemark *string `json:"restore_remark,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	Progress *string `json:"progress,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	RestoreName *string `json:"restore_name,omitempty"`

	BackupRemark *string `json:"backup_remark,omitempty"`

	Status *InstanceRestoreInfoStatus `json:"status,omitempty"`
}

func (o InstanceRestoreInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceRestoreInfo struct{}"
	}

	return strings.Join([]string{"InstanceRestoreInfo", string(data)}, " ")
}

type InstanceRestoreInfoStatus struct {
	value string
}

type InstanceRestoreInfoStatusEnum struct {
	WAITING   InstanceRestoreInfoStatus
	RESTORING InstanceRestoreInfoStatus
	SUCCEED   InstanceRestoreInfoStatus
	FAILED    InstanceRestoreInfoStatus
}

func GetInstanceRestoreInfoStatusEnum() InstanceRestoreInfoStatusEnum {
	return InstanceRestoreInfoStatusEnum{
		WAITING: InstanceRestoreInfoStatus{
			value: "waiting",
		},
		RESTORING: InstanceRestoreInfoStatus{
			value: "restoring",
		},
		SUCCEED: InstanceRestoreInfoStatus{
			value: "succeed",
		},
		FAILED: InstanceRestoreInfoStatus{
			value: "failed",
		},
	}
}

func (c InstanceRestoreInfoStatus) Value() string {
	return c.value
}

func (c InstanceRestoreInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceRestoreInfoStatus) UnmarshalJSON(b []byte) error {
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
