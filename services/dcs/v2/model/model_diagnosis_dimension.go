package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DiagnosisDimension struct {
	Name DiagnosisDimensionName `json:"name"`

	AbnormalNum int32 `json:"abnormal_num"`

	FailedNum int32 `json:"failed_num"`

	DiagnosisItemList []DiagnosisItem `json:"diagnosis_item_list"`
}

func (o DiagnosisDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisDimension struct{}"
	}

	return strings.Join([]string{"DiagnosisDimension", string(data)}, " ")
}

type DiagnosisDimensionName struct {
	value string
}

type DiagnosisDimensionNameEnum struct {
	NETWORK DiagnosisDimensionName
	STORAGE DiagnosisDimensionName
	LOAD    DiagnosisDimensionName
}

func GetDiagnosisDimensionNameEnum() DiagnosisDimensionNameEnum {
	return DiagnosisDimensionNameEnum{
		NETWORK: DiagnosisDimensionName{
			value: "network",
		},
		STORAGE: DiagnosisDimensionName{
			value: "storage",
		},
		LOAD: DiagnosisDimensionName{
			value: "load",
		},
	}
}

func (c DiagnosisDimensionName) Value() string {
	return c.value
}

func (c DiagnosisDimensionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisDimensionName) UnmarshalJSON(b []byte) error {
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
