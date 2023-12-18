package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchCreateVolumeTagsRequestBody struct {
	Action BatchCreateVolumeTagsRequestBodyAction `json:"action"`

	Tags []Tag `json:"tags"`
}

func (o BatchCreateVolumeTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVolumeTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateVolumeTagsRequestBody", string(data)}, " ")
}

type BatchCreateVolumeTagsRequestBodyAction struct {
	value string
}

type BatchCreateVolumeTagsRequestBodyActionEnum struct {
	CREATE BatchCreateVolumeTagsRequestBodyAction
}

func GetBatchCreateVolumeTagsRequestBodyActionEnum() BatchCreateVolumeTagsRequestBodyActionEnum {
	return BatchCreateVolumeTagsRequestBodyActionEnum{
		CREATE: BatchCreateVolumeTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateVolumeTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateVolumeTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateVolumeTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
