package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListConfigurationsResponse struct {
	ConfigTime *string `json:"config_time,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	RedisConfig *[]QueryRedisConfig `json:"redis_config,omitempty"`

	ConfigStatus *ListConfigurationsResponseConfigStatus `json:"config_status,omitempty"`

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResponse", string(data)}, " ")
}

type ListConfigurationsResponseConfigStatus struct {
	value string
}

type ListConfigurationsResponseConfigStatusEnum struct {
	UPDATING ListConfigurationsResponseConfigStatus
	FAILURE  ListConfigurationsResponseConfigStatus
	SUCCESS  ListConfigurationsResponseConfigStatus
}

func GetListConfigurationsResponseConfigStatusEnum() ListConfigurationsResponseConfigStatusEnum {
	return ListConfigurationsResponseConfigStatusEnum{
		UPDATING: ListConfigurationsResponseConfigStatus{
			value: "UPDATING",
		},
		FAILURE: ListConfigurationsResponseConfigStatus{
			value: "FAILURE",
		},
		SUCCESS: ListConfigurationsResponseConfigStatus{
			value: "SUCCESS",
		},
	}
}

func (c ListConfigurationsResponseConfigStatus) Value() string {
	return c.value
}

func (c ListConfigurationsResponseConfigStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConfigurationsResponseConfigStatus) UnmarshalJSON(b []byte) error {
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
