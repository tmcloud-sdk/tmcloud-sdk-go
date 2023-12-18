package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteVolumeTagsRequestBody struct {
	Action BatchDeleteVolumeTagsRequestBodyAction `json:"action"`

	Tags []DeleteTagsOption `json:"tags"`
}

func (o BatchDeleteVolumeTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVolumeTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteVolumeTagsRequestBody", string(data)}, " ")
}

type BatchDeleteVolumeTagsRequestBodyAction struct {
	value string
}

type BatchDeleteVolumeTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteVolumeTagsRequestBodyAction
}

func GetBatchDeleteVolumeTagsRequestBodyActionEnum() BatchDeleteVolumeTagsRequestBodyActionEnum {
	return BatchDeleteVolumeTagsRequestBodyActionEnum{
		DELETE: BatchDeleteVolumeTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteVolumeTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteVolumeTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteVolumeTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
