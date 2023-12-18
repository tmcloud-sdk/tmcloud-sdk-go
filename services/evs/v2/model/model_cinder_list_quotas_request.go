package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CinderListQuotasRequest struct {
	TargetProjectId string `json:"target_project_id"`

	Usage CinderListQuotasRequestUsage `json:"usage"`
}

func (o CinderListQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListQuotasRequest struct{}"
	}

	return strings.Join([]string{"CinderListQuotasRequest", string(data)}, " ")
}

type CinderListQuotasRequestUsage struct {
	value string
}

type CinderListQuotasRequestUsageEnum struct {
	TRUE CinderListQuotasRequestUsage
}

func GetCinderListQuotasRequestUsageEnum() CinderListQuotasRequestUsageEnum {
	return CinderListQuotasRequestUsageEnum{
		TRUE: CinderListQuotasRequestUsage{
			value: "True",
		},
	}
}

func (c CinderListQuotasRequestUsage) Value() string {
	return c.value
}

func (c CinderListQuotasRequestUsage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CinderListQuotasRequestUsage) UnmarshalJSON(b []byte) error {
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
