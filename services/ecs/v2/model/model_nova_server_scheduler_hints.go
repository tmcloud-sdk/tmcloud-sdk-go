package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NovaServerSchedulerHints struct {
	Tenancy *[]NovaServerSchedulerHintsTenancy `json:"tenancy,omitempty"`

	DedicatedHostId *[]string `json:"dedicated_host_id,omitempty"`
}

func (o NovaServerSchedulerHints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerSchedulerHints struct{}"
	}

	return strings.Join([]string{"NovaServerSchedulerHints", string(data)}, " ")
}

type NovaServerSchedulerHintsTenancy struct {
	value string
}

type NovaServerSchedulerHintsTenancyEnum struct {
	SHARED    NovaServerSchedulerHintsTenancy
	DEDICATED NovaServerSchedulerHintsTenancy
}

func GetNovaServerSchedulerHintsTenancyEnum() NovaServerSchedulerHintsTenancyEnum {
	return NovaServerSchedulerHintsTenancyEnum{
		SHARED: NovaServerSchedulerHintsTenancy{
			value: "shared",
		},
		DEDICATED: NovaServerSchedulerHintsTenancy{
			value: "dedicated",
		},
	}
}

func (c NovaServerSchedulerHintsTenancy) Value() string {
	return c.value
}

func (c NovaServerSchedulerHintsTenancy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaServerSchedulerHintsTenancy) UnmarshalJSON(b []byte) error {
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
