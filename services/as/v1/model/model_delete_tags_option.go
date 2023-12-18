package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteTagsOption struct {
	Tags []TagsSingleValue `json:"tags"`

	Action DeleteTagsOptionAction `json:"action"`
}

func (o DeleteTagsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagsOption struct{}"
	}

	return strings.Join([]string{"DeleteTagsOption", string(data)}, " ")
}

type DeleteTagsOptionAction struct {
	value string
}

type DeleteTagsOptionActionEnum struct {
	DELETE DeleteTagsOptionAction
}

func GetDeleteTagsOptionActionEnum() DeleteTagsOptionActionEnum {
	return DeleteTagsOptionActionEnum{
		DELETE: DeleteTagsOptionAction{
			value: "delete",
		},
	}
}

func (c DeleteTagsOptionAction) Value() string {
	return c.value
}

func (c DeleteTagsOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTagsOptionAction) UnmarshalJSON(b []byte) error {
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
