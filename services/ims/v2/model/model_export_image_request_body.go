package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ExportImageRequestBody struct {
	BucketUrl string `json:"bucket_url"`

	FileFormat ExportImageRequestBodyFileFormat `json:"file_format"`

	IsQuickExport *bool `json:"is_quick_export,omitempty"`
}

func (o ExportImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageRequestBody struct{}"
	}

	return strings.Join([]string{"ExportImageRequestBody", string(data)}, " ")
}

type ExportImageRequestBodyFileFormat struct {
	value string
}

type ExportImageRequestBodyFileFormatEnum struct {
	QCOW2 ExportImageRequestBodyFileFormat
	VHD   ExportImageRequestBodyFileFormat
	ZVHD  ExportImageRequestBodyFileFormat
	VMDK  ExportImageRequestBodyFileFormat
}

func GetExportImageRequestBodyFileFormatEnum() ExportImageRequestBodyFileFormatEnum {
	return ExportImageRequestBodyFileFormatEnum{
		QCOW2: ExportImageRequestBodyFileFormat{
			value: "qcow2",
		},
		VHD: ExportImageRequestBodyFileFormat{
			value: "vhd",
		},
		ZVHD: ExportImageRequestBodyFileFormat{
			value: "zvhd",
		},
		VMDK: ExportImageRequestBodyFileFormat{
			value: "vmdk",
		},
	}
}

func (c ExportImageRequestBodyFileFormat) Value() string {
	return c.value
}

func (c ExportImageRequestBodyFileFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportImageRequestBodyFileFormat) UnmarshalJSON(b []byte) error {
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
