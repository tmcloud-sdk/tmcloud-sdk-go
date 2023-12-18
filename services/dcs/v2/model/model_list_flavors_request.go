package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListFlavorsRequest struct {
	InstanceId *string `json:"instance_id,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`

	CacheMode *string `json:"cache_mode,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	CpuType *ListFlavorsRequestCpuType `json:"cpu_type,omitempty"`

	Capacity *string `json:"capacity,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

type ListFlavorsRequestCpuType struct {
	value string
}

type ListFlavorsRequestCpuTypeEnum struct {
	X86_64  ListFlavorsRequestCpuType
	AARCH64 ListFlavorsRequestCpuType
}

func GetListFlavorsRequestCpuTypeEnum() ListFlavorsRequestCpuTypeEnum {
	return ListFlavorsRequestCpuTypeEnum{
		X86_64: ListFlavorsRequestCpuType{
			value: "x86_64",
		},
		AARCH64: ListFlavorsRequestCpuType{
			value: "aarch64",
		},
	}
}

func (c ListFlavorsRequestCpuType) Value() string {
	return c.value
}

func (c ListFlavorsRequestCpuType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestCpuType) UnmarshalJSON(b []byte) error {
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
