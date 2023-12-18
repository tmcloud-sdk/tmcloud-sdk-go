package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrePaidServerSchedulerHints struct {
	Group *string `json:"group,omitempty"`

	Tenancy *PrePaidServerSchedulerHintsTenancy `json:"tenancy,omitempty"`

	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`
}

func (o PrePaidServerSchedulerHints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerSchedulerHints struct{}"
	}

	return strings.Join([]string{"PrePaidServerSchedulerHints", string(data)}, " ")
}

type PrePaidServerSchedulerHintsTenancy struct {
	value string
}

type PrePaidServerSchedulerHintsTenancyEnum struct {
	SHARED    PrePaidServerSchedulerHintsTenancy
	DEDICATED PrePaidServerSchedulerHintsTenancy
}

func GetPrePaidServerSchedulerHintsTenancyEnum() PrePaidServerSchedulerHintsTenancyEnum {
	return PrePaidServerSchedulerHintsTenancyEnum{
		SHARED: PrePaidServerSchedulerHintsTenancy{
			value: "shared",
		},
		DEDICATED: PrePaidServerSchedulerHintsTenancy{
			value: "dedicated",
		},
	}
}

func (c PrePaidServerSchedulerHintsTenancy) Value() string {
	return c.value
}

func (c PrePaidServerSchedulerHintsTenancy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerSchedulerHintsTenancy) UnmarshalJSON(b []byte) error {
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
