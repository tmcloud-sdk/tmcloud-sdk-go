package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type LineCompareDetail struct {
	SourceTableName string `json:"source_table_name"`

	TargetTableName string `json:"target_table_name"`

	SourceRowNum int32 `json:"source_row_num"`

	TargetRowNum int32 `json:"target_row_num"`

	DiffRowNum int32 `json:"diff_row_num"`

	LineCompareResult LineCompareDetailLineCompareResult `json:"line_compare_result"`

	Message *string `json:"message,omitempty"`
}

func (o LineCompareDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareDetail struct{}"
	}

	return strings.Join([]string{"LineCompareDetail", string(data)}, " ")
}

type LineCompareDetailLineCompareResult struct {
	value string
}

type LineCompareDetailLineCompareResultEnum struct {
	CONSISTENT             LineCompareDetailLineCompareResult
	INCONSISTENT           LineCompareDetailLineCompareResult
	COMPARING              LineCompareDetailLineCompareResult
	WAITING_FOR_COMPARISON LineCompareDetailLineCompareResult
	FAILED_TO_COMPARE      LineCompareDetailLineCompareResult
	TARGET_DB_NOT_EXIT     LineCompareDetailLineCompareResult
	CAN_NOT_COMPARE        LineCompareDetailLineCompareResult
}

func GetLineCompareDetailLineCompareResultEnum() LineCompareDetailLineCompareResultEnum {
	return LineCompareDetailLineCompareResultEnum{
		CONSISTENT: LineCompareDetailLineCompareResult{
			value: "CONSISTENT",
		},
		INCONSISTENT: LineCompareDetailLineCompareResult{
			value: "INCONSISTENT",
		},
		COMPARING: LineCompareDetailLineCompareResult{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: LineCompareDetailLineCompareResult{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: LineCompareDetailLineCompareResult{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIT: LineCompareDetailLineCompareResult{
			value: "TARGET_DB_NOT_EXIT",
		},
		CAN_NOT_COMPARE: LineCompareDetailLineCompareResult{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c LineCompareDetailLineCompareResult) Value() string {
	return c.value
}

func (c LineCompareDetailLineCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LineCompareDetailLineCompareResult) UnmarshalJSON(b []byte) error {
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
