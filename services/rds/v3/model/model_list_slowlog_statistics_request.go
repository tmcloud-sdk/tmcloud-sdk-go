package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListSlowlogStatisticsRequest struct {
	XLanguage *ListSlowlogStatisticsRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	CurPage int32 `json:"cur_page"`

	PerPage int32 `json:"per_page"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	Type ListSlowlogStatisticsRequestType `json:"type"`

	Sort *string `json:"sort,omitempty"`
}

func (o ListSlowlogStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListSlowlogStatisticsRequest", string(data)}, " ")
}

type ListSlowlogStatisticsRequestXLanguage struct {
	value string
}

type ListSlowlogStatisticsRequestXLanguageEnum struct {
	ZH_CN ListSlowlogStatisticsRequestXLanguage
	EN_US ListSlowlogStatisticsRequestXLanguage
}

func GetListSlowlogStatisticsRequestXLanguageEnum() ListSlowlogStatisticsRequestXLanguageEnum {
	return ListSlowlogStatisticsRequestXLanguageEnum{
		ZH_CN: ListSlowlogStatisticsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSlowlogStatisticsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSlowlogStatisticsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSlowlogStatisticsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowlogStatisticsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListSlowlogStatisticsRequestType struct {
	value string
}

type ListSlowlogStatisticsRequestTypeEnum struct {
	INSERT ListSlowlogStatisticsRequestType
	UPDATE ListSlowlogStatisticsRequestType
	SELECT ListSlowlogStatisticsRequestType
	DELETE ListSlowlogStatisticsRequestType
	CREATE ListSlowlogStatisticsRequestType
	ALL    ListSlowlogStatisticsRequestType
}

func GetListSlowlogStatisticsRequestTypeEnum() ListSlowlogStatisticsRequestTypeEnum {
	return ListSlowlogStatisticsRequestTypeEnum{
		INSERT: ListSlowlogStatisticsRequestType{
			value: "INSERT",
		},
		UPDATE: ListSlowlogStatisticsRequestType{
			value: "UPDATE",
		},
		SELECT: ListSlowlogStatisticsRequestType{
			value: "SELECT",
		},
		DELETE: ListSlowlogStatisticsRequestType{
			value: "DELETE",
		},
		CREATE: ListSlowlogStatisticsRequestType{
			value: "CREATE",
		},
		ALL: ListSlowlogStatisticsRequestType{
			value: "ALL",
		},
	}
}

func (c ListSlowlogStatisticsRequestType) Value() string {
	return c.value
}

func (c ListSlowlogStatisticsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowlogStatisticsRequestType) UnmarshalJSON(b []byte) error {
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
