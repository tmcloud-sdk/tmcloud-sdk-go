package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CommandTimeTakenList struct {
	TotalNum int32 `json:"total_num"`

	TotalUsecSum float64 `json:"total_usec_sum"`

	Result CommandTimeTakenListResult `json:"result"`

	CommandList []CommandTimeTaken `json:"command_list"`
}

func (o CommandTimeTakenList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommandTimeTakenList struct{}"
	}

	return strings.Join([]string{"CommandTimeTakenList", string(data)}, " ")
}

type CommandTimeTakenListResult struct {
	value string
}

type CommandTimeTakenListResultEnum struct {
	SUCCEED CommandTimeTakenListResult
	FAILED  CommandTimeTakenListResult
}

func GetCommandTimeTakenListResultEnum() CommandTimeTakenListResultEnum {
	return CommandTimeTakenListResultEnum{
		SUCCEED: CommandTimeTakenListResult{
			value: "succeed",
		},
		FAILED: CommandTimeTakenListResult{
			value: "failed",
		},
	}
}

func (c CommandTimeTakenListResult) Value() string {
	return c.value
}

func (c CommandTimeTakenListResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommandTimeTakenListResult) UnmarshalJSON(b []byte) error {
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
