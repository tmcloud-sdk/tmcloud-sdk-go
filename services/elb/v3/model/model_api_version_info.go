package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ApiVersionInfo struct {
	Id string `json:"id"`

	Status ApiVersionInfoStatus `json:"status"`
}

func (o ApiVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionInfo struct{}"
	}

	return strings.Join([]string{"ApiVersionInfo", string(data)}, " ")
}

type ApiVersionInfoStatus struct {
	value string
}

type ApiVersionInfoStatusEnum struct {
	CURRENT    ApiVersionInfoStatus
	STABLE     ApiVersionInfoStatus
	DEPRECATED ApiVersionInfoStatus
}

func GetApiVersionInfoStatusEnum() ApiVersionInfoStatusEnum {
	return ApiVersionInfoStatusEnum{
		CURRENT: ApiVersionInfoStatus{
			value: "CURRENT",
		},
		STABLE: ApiVersionInfoStatus{
			value: "STABLE",
		},
		DEPRECATED: ApiVersionInfoStatus{
			value: "DEPRECATED",
		},
	}
}

func (c ApiVersionInfoStatus) Value() string {
	return c.value
}

func (c ApiVersionInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiVersionInfoStatus) UnmarshalJSON(b []byte) error {
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
