package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PauseInfo struct {
	JobId string `json:"job_id"`

	PauseMode PauseInfoPauseMode `json:"pause_mode"`
}

func (o PauseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseInfo struct{}"
	}

	return strings.Join([]string{"PauseInfo", string(data)}, " ")
}

type PauseInfoPauseMode struct {
	value string
}

type PauseInfoPauseModeEnum struct {
	TARGET PauseInfoPauseMode
	ALL    PauseInfoPauseMode
}

func GetPauseInfoPauseModeEnum() PauseInfoPauseModeEnum {
	return PauseInfoPauseModeEnum{
		TARGET: PauseInfoPauseMode{
			value: "target",
		},
		ALL: PauseInfoPauseMode{
			value: "all",
		},
	}
}

func (c PauseInfoPauseMode) Value() string {
	return c.value
}

func (c PauseInfoPauseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PauseInfoPauseMode) UnmarshalJSON(b []byte) error {
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
