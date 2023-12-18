package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SlowLogStatisticsForLtsRequest struct {
	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Type *SlowLogStatisticsForLtsRequestType `json:"type,omitempty"`

	Database *string `json:"database,omitempty"`

	Sort *string `json:"sort,omitempty"`

	Order *SlowLogStatisticsForLtsRequestOrder `json:"order,omitempty"`
}

func (o SlowLogStatisticsForLtsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogStatisticsForLtsRequest struct{}"
	}

	return strings.Join([]string{"SlowLogStatisticsForLtsRequest", string(data)}, " ")
}

type SlowLogStatisticsForLtsRequestType struct {
	value string
}

type SlowLogStatisticsForLtsRequestTypeEnum struct {
	INSERT SlowLogStatisticsForLtsRequestType
	UPDATE SlowLogStatisticsForLtsRequestType
	SELECT SlowLogStatisticsForLtsRequestType
	DELETE SlowLogStatisticsForLtsRequestType
	CREATE SlowLogStatisticsForLtsRequestType
	ALL    SlowLogStatisticsForLtsRequestType
}

func GetSlowLogStatisticsForLtsRequestTypeEnum() SlowLogStatisticsForLtsRequestTypeEnum {
	return SlowLogStatisticsForLtsRequestTypeEnum{
		INSERT: SlowLogStatisticsForLtsRequestType{
			value: "INSERT",
		},
		UPDATE: SlowLogStatisticsForLtsRequestType{
			value: "UPDATE",
		},
		SELECT: SlowLogStatisticsForLtsRequestType{
			value: "SELECT",
		},
		DELETE: SlowLogStatisticsForLtsRequestType{
			value: "DELETE",
		},
		CREATE: SlowLogStatisticsForLtsRequestType{
			value: "CREATE",
		},
		ALL: SlowLogStatisticsForLtsRequestType{
			value: "ALL",
		},
	}
}

func (c SlowLogStatisticsForLtsRequestType) Value() string {
	return c.value
}

func (c SlowLogStatisticsForLtsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SlowLogStatisticsForLtsRequestType) UnmarshalJSON(b []byte) error {
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

type SlowLogStatisticsForLtsRequestOrder struct {
	value string
}

type SlowLogStatisticsForLtsRequestOrderEnum struct {
	DESC SlowLogStatisticsForLtsRequestOrder
	ASC  SlowLogStatisticsForLtsRequestOrder
}

func GetSlowLogStatisticsForLtsRequestOrderEnum() SlowLogStatisticsForLtsRequestOrderEnum {
	return SlowLogStatisticsForLtsRequestOrderEnum{
		DESC: SlowLogStatisticsForLtsRequestOrder{
			value: "desc",
		},
		ASC: SlowLogStatisticsForLtsRequestOrder{
			value: "asc",
		},
	}
}

func (c SlowLogStatisticsForLtsRequestOrder) Value() string {
	return c.value
}

func (c SlowLogStatisticsForLtsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SlowLogStatisticsForLtsRequestOrder) UnmarshalJSON(b []byte) error {
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
