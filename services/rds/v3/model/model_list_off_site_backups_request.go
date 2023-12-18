package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListOffSiteBackupsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	BackupId *string `json:"backup_id,omitempty"`

	BackupType *ListOffSiteBackupsRequestBackupType `json:"backup_type,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`
}

func (o ListOffSiteBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOffSiteBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListOffSiteBackupsRequest", string(data)}, " ")
}

type ListOffSiteBackupsRequestBackupType struct {
	value string
}

type ListOffSiteBackupsRequestBackupTypeEnum struct {
	AUTO        ListOffSiteBackupsRequestBackupType
	INCREMENTAL ListOffSiteBackupsRequestBackupType
}

func GetListOffSiteBackupsRequestBackupTypeEnum() ListOffSiteBackupsRequestBackupTypeEnum {
	return ListOffSiteBackupsRequestBackupTypeEnum{
		AUTO: ListOffSiteBackupsRequestBackupType{
			value: "auto",
		},
		INCREMENTAL: ListOffSiteBackupsRequestBackupType{
			value: "incremental",
		},
	}
}

func (c ListOffSiteBackupsRequestBackupType) Value() string {
	return c.value
}

func (c ListOffSiteBackupsRequestBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOffSiteBackupsRequestBackupType) UnmarshalJSON(b []byte) error {
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
