package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ContentCompareResultOverview struct {
	SourceDbName string `json:"source_db_name"`

	TargetDbName string `json:"target_db_name"`

	ContentCompareResult ContentCompareResultOverviewContentCompareResult `json:"content_compare_result"`
}

func (o ContentCompareResultOverview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareResultOverview struct{}"
	}

	return strings.Join([]string{"ContentCompareResultOverview", string(data)}, " ")
}

type ContentCompareResultOverviewContentCompareResult struct {
	value string
}

type ContentCompareResultOverviewContentCompareResultEnum struct {
	CONSISTENT             ContentCompareResultOverviewContentCompareResult
	INCONSISTENT           ContentCompareResultOverviewContentCompareResult
	COMPARING              ContentCompareResultOverviewContentCompareResult
	WAITING_FOR_COMPARISON ContentCompareResultOverviewContentCompareResult
	FAILED_TO_COMPARE      ContentCompareResultOverviewContentCompareResult
	TARGET_DB_NOT_EXIT     ContentCompareResultOverviewContentCompareResult
	CAN_NOT_COMPARE        ContentCompareResultOverviewContentCompareResult
}

func GetContentCompareResultOverviewContentCompareResultEnum() ContentCompareResultOverviewContentCompareResultEnum {
	return ContentCompareResultOverviewContentCompareResultEnum{
		CONSISTENT: ContentCompareResultOverviewContentCompareResult{
			value: "CONSISTENT",
		},
		INCONSISTENT: ContentCompareResultOverviewContentCompareResult{
			value: "INCONSISTENT",
		},
		COMPARING: ContentCompareResultOverviewContentCompareResult{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: ContentCompareResultOverviewContentCompareResult{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: ContentCompareResultOverviewContentCompareResult{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIT: ContentCompareResultOverviewContentCompareResult{
			value: "TARGET_DB_NOT_EXIT",
		},
		CAN_NOT_COMPARE: ContentCompareResultOverviewContentCompareResult{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c ContentCompareResultOverviewContentCompareResult) Value() string {
	return c.value
}

func (c ContentCompareResultOverviewContentCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentCompareResultOverviewContentCompareResult) UnmarshalJSON(b []byte) error {
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
