package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BackupForList struct {
	Id string `json:"id"`

	InstanceId string `json:"instance_id"`

	Name string `json:"name"`

	Databases *[]BackupDatabase `json:"databases,omitempty"`

	BeginTime string `json:"begin_time"`

	EndTime string `json:"end_time"`

	Status BackupForListStatus `json:"status"`

	Type BackupForListType `json:"type"`

	Size int64 `json:"size"`

	Datastore *BackupDatastore `json:"datastore"`

	AssociatedWithDdm *bool `json:"associated_with_ddm,omitempty"`
}

func (o BackupForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupForList struct{}"
	}

	return strings.Join([]string{"BackupForList", string(data)}, " ")
}

type BackupForListStatus struct {
	value string
}

type BackupForListStatusEnum struct {
	BUILDING  BackupForListStatus
	COMPLETED BackupForListStatus
	FAILED    BackupForListStatus
	DELETING  BackupForListStatus
}

func GetBackupForListStatusEnum() BackupForListStatusEnum {
	return BackupForListStatusEnum{
		BUILDING: BackupForListStatus{
			value: "BUILDING",
		},
		COMPLETED: BackupForListStatus{
			value: "COMPLETED",
		},
		FAILED: BackupForListStatus{
			value: "FAILED",
		},
		DELETING: BackupForListStatus{
			value: "DELETING",
		},
	}
}

func (c BackupForListStatus) Value() string {
	return c.value
}

func (c BackupForListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupForListStatus) UnmarshalJSON(b []byte) error {
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

type BackupForListType struct {
	value string
}

type BackupForListTypeEnum struct {
	AUTO        BackupForListType
	MANUAL      BackupForListType
	FRAGMENT    BackupForListType
	INCREMENTAL BackupForListType
}

func GetBackupForListTypeEnum() BackupForListTypeEnum {
	return BackupForListTypeEnum{
		AUTO: BackupForListType{
			value: "auto",
		},
		MANUAL: BackupForListType{
			value: "manual",
		},
		FRAGMENT: BackupForListType{
			value: "fragment",
		},
		INCREMENTAL: BackupForListType{
			value: "incremental",
		},
	}
}

func (c BackupForListType) Value() string {
	return c.value
}

func (c BackupForListType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupForListType) UnmarshalJSON(b []byte) error {
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
