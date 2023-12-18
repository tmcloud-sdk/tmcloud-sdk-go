package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListSlowlogRequest struct {
	InstanceId string `json:"instance_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	SortKey *ListSlowlogRequestSortKey `json:"sort_key,omitempty"`

	SortDir *ListSlowlogRequestSortDir `json:"sort_dir,omitempty"`

	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`
}

func (o ListSlowlogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogRequest struct{}"
	}

	return strings.Join([]string{"ListSlowlogRequest", string(data)}, " ")
}

type ListSlowlogRequestSortKey struct {
	value string
}

type ListSlowlogRequestSortKeyEnum struct {
	START_TIME ListSlowlogRequestSortKey
	DURATION   ListSlowlogRequestSortKey
}

func GetListSlowlogRequestSortKeyEnum() ListSlowlogRequestSortKeyEnum {
	return ListSlowlogRequestSortKeyEnum{
		START_TIME: ListSlowlogRequestSortKey{
			value: "start_time",
		},
		DURATION: ListSlowlogRequestSortKey{
			value: "duration",
		},
	}
}

func (c ListSlowlogRequestSortKey) Value() string {
	return c.value
}

func (c ListSlowlogRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowlogRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListSlowlogRequestSortDir struct {
	value string
}

type ListSlowlogRequestSortDirEnum struct {
	DESC ListSlowlogRequestSortDir
	ASC  ListSlowlogRequestSortDir
}

func GetListSlowlogRequestSortDirEnum() ListSlowlogRequestSortDirEnum {
	return ListSlowlogRequestSortDirEnum{
		DESC: ListSlowlogRequestSortDir{
			value: "desc",
		},
		ASC: ListSlowlogRequestSortDir{
			value: "asc",
		},
	}
}

func (c ListSlowlogRequestSortDir) Value() string {
	return c.value
}

func (c ListSlowlogRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowlogRequestSortDir) UnmarshalJSON(b []byte) error {
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
