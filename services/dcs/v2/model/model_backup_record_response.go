package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BackupRecordResponse struct {
	BackupId *string `json:"backup_id,omitempty"`

	Period *string `json:"period,omitempty"`

	BackupName *string `json:"backup_name,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Size *int64 `json:"size,omitempty"`

	BackupType *BackupRecordResponseBackupType `json:"backup_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Progress *string `json:"progress,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	Remark *string `json:"remark,omitempty"`

	Status *BackupRecordResponseStatus `json:"status,omitempty"`

	IsSupportRestore *string `json:"is_support_restore,omitempty"`
}

func (o BackupRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRecordResponse struct{}"
	}

	return strings.Join([]string{"BackupRecordResponse", string(data)}, " ")
}

type BackupRecordResponseBackupType struct {
	value string
}

type BackupRecordResponseBackupTypeEnum struct {
	MANUAL BackupRecordResponseBackupType
	AUTO   BackupRecordResponseBackupType
}

func GetBackupRecordResponseBackupTypeEnum() BackupRecordResponseBackupTypeEnum {
	return BackupRecordResponseBackupTypeEnum{
		MANUAL: BackupRecordResponseBackupType{
			value: "manual",
		},
		AUTO: BackupRecordResponseBackupType{
			value: "auto",
		},
	}
}

func (c BackupRecordResponseBackupType) Value() string {
	return c.value
}

func (c BackupRecordResponseBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupRecordResponseBackupType) UnmarshalJSON(b []byte) error {
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

type BackupRecordResponseStatus struct {
	value string
}

type BackupRecordResponseStatusEnum struct {
	WAITING   BackupRecordResponseStatus
	BACKUPING BackupRecordResponseStatus
	SUCCEED   BackupRecordResponseStatus
	FAILED    BackupRecordResponseStatus
	EXPIRED   BackupRecordResponseStatus
	DELETED   BackupRecordResponseStatus
}

func GetBackupRecordResponseStatusEnum() BackupRecordResponseStatusEnum {
	return BackupRecordResponseStatusEnum{
		WAITING: BackupRecordResponseStatus{
			value: "waiting",
		},
		BACKUPING: BackupRecordResponseStatus{
			value: "backuping",
		},
		SUCCEED: BackupRecordResponseStatus{
			value: "succeed",
		},
		FAILED: BackupRecordResponseStatus{
			value: "failed",
		},
		EXPIRED: BackupRecordResponseStatus{
			value: "expired",
		},
		DELETED: BackupRecordResponseStatus{
			value: "deleted",
		},
	}
}

func (c BackupRecordResponseStatus) Value() string {
	return c.value
}

func (c BackupRecordResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupRecordResponseStatus) UnmarshalJSON(b []byte) error {
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
