package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BackupFilesBody struct {
	FileSource *BackupFilesBodyFileSource `json:"file_source,omitempty"`

	BucketName string `json:"bucket_name"`

	Files []Files `json:"files"`

	BackupId *string `json:"backup_id,omitempty"`
}

func (o BackupFilesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupFilesBody struct{}"
	}

	return strings.Join([]string{"BackupFilesBody", string(data)}, " ")
}

type BackupFilesBodyFileSource struct {
	value string
}

type BackupFilesBodyFileSourceEnum struct {
	SELF_BUILD_OBS BackupFilesBodyFileSource
}

func GetBackupFilesBodyFileSourceEnum() BackupFilesBodyFileSourceEnum {
	return BackupFilesBodyFileSourceEnum{
		SELF_BUILD_OBS: BackupFilesBodyFileSource{
			value: "self_build_obs",
		},
	}
}

func (c BackupFilesBodyFileSource) Value() string {
	return c.value
}

func (c BackupFilesBodyFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupFilesBodyFileSource) UnmarshalJSON(b []byte) error {
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
