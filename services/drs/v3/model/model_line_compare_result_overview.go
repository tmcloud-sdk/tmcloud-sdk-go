package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type LineCompareResultOverview struct {
	SourceDbName string `json:"source_db_name"`

	TargetDbName string `json:"target_db_name"`

	LineCompareResult LineCompareResultOverviewLineCompareResult `json:"line_compare_result"`
}

func (o LineCompareResultOverview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareResultOverview struct{}"
	}

	return strings.Join([]string{"LineCompareResultOverview", string(data)}, " ")
}

type LineCompareResultOverviewLineCompareResult struct {
	value string
}

type LineCompareResultOverviewLineCompareResultEnum struct {
	CONSISTENT             LineCompareResultOverviewLineCompareResult
	INCONSISTENT           LineCompareResultOverviewLineCompareResult
	COMPARING              LineCompareResultOverviewLineCompareResult
	WAITING_FOR_COMPARISON LineCompareResultOverviewLineCompareResult
	FAILED_TO_COMPARE      LineCompareResultOverviewLineCompareResult
	TARGET_DB_NOT_EXIT     LineCompareResultOverviewLineCompareResult
	CAN_NOT_COMPARE        LineCompareResultOverviewLineCompareResult
}

func GetLineCompareResultOverviewLineCompareResultEnum() LineCompareResultOverviewLineCompareResultEnum {
	return LineCompareResultOverviewLineCompareResultEnum{
		CONSISTENT: LineCompareResultOverviewLineCompareResult{
			value: "CONSISTENT",
		},
		INCONSISTENT: LineCompareResultOverviewLineCompareResult{
			value: "INCONSISTENT",
		},
		COMPARING: LineCompareResultOverviewLineCompareResult{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: LineCompareResultOverviewLineCompareResult{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: LineCompareResultOverviewLineCompareResult{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIT: LineCompareResultOverviewLineCompareResult{
			value: "TARGET_DB_NOT_EXIT",
		},
		CAN_NOT_COMPARE: LineCompareResultOverviewLineCompareResult{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c LineCompareResultOverviewLineCompareResult) Value() string {
	return c.value
}

func (c LineCompareResultOverviewLineCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LineCompareResultOverviewLineCompareResult) UnmarshalJSON(b []byte) error {
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
