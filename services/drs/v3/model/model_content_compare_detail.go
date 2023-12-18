package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ContentCompareDetail struct {
	SourceDbName string `json:"source_db_name"`

	TargetDbName string `json:"target_db_name"`

	SourceTableName string `json:"source_table_name"`

	TargetTableName string `json:"target_table_name"`

	SourceRowNum int32 `json:"source_row_num"`

	TargetRowNum int32 `json:"target_row_num"`

	DiffRowNum int32 `json:"diff_row_num"`

	LineCompareResult *ContentCompareDetailLineCompareResult `json:"line_compare_result,omitempty"`

	ContentCompareResult ContentCompareDetailContentCompareResult `json:"content_compare_result"`

	Message *string `json:"message,omitempty"`
}

func (o ContentCompareDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareDetail struct{}"
	}

	return strings.Join([]string{"ContentCompareDetail", string(data)}, " ")
}

type ContentCompareDetailLineCompareResult struct {
	value string
}

type ContentCompareDetailLineCompareResultEnum struct {
	CONSISTENT             ContentCompareDetailLineCompareResult
	INCONSISTENT           ContentCompareDetailLineCompareResult
	COMPARING              ContentCompareDetailLineCompareResult
	WAITING_FOR_COMPARISON ContentCompareDetailLineCompareResult
	FAILED_TO_COMPARE      ContentCompareDetailLineCompareResult
	TARGET_DB_NOT_EXIT     ContentCompareDetailLineCompareResult
	CAN_NOT_COMPARE        ContentCompareDetailLineCompareResult
}

func GetContentCompareDetailLineCompareResultEnum() ContentCompareDetailLineCompareResultEnum {
	return ContentCompareDetailLineCompareResultEnum{
		CONSISTENT: ContentCompareDetailLineCompareResult{
			value: "CONSISTENT",
		},
		INCONSISTENT: ContentCompareDetailLineCompareResult{
			value: "INCONSISTENT",
		},
		COMPARING: ContentCompareDetailLineCompareResult{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: ContentCompareDetailLineCompareResult{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: ContentCompareDetailLineCompareResult{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIT: ContentCompareDetailLineCompareResult{
			value: "TARGET_DB_NOT_EXIT",
		},
		CAN_NOT_COMPARE: ContentCompareDetailLineCompareResult{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c ContentCompareDetailLineCompareResult) Value() string {
	return c.value
}

func (c ContentCompareDetailLineCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentCompareDetailLineCompareResult) UnmarshalJSON(b []byte) error {
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

type ContentCompareDetailContentCompareResult struct {
	value string
}

type ContentCompareDetailContentCompareResultEnum struct {
	CONSISTENT             ContentCompareDetailContentCompareResult
	INCONSISTENT           ContentCompareDetailContentCompareResult
	COMPARING              ContentCompareDetailContentCompareResult
	WAITING_FOR_COMPARISON ContentCompareDetailContentCompareResult
	FAILED_TO_COMPARE      ContentCompareDetailContentCompareResult
	TARGET_DB_NOT_EXIT     ContentCompareDetailContentCompareResult
	CAN_NOT_COMPARE        ContentCompareDetailContentCompareResult
}

func GetContentCompareDetailContentCompareResultEnum() ContentCompareDetailContentCompareResultEnum {
	return ContentCompareDetailContentCompareResultEnum{
		CONSISTENT: ContentCompareDetailContentCompareResult{
			value: "CONSISTENT",
		},
		INCONSISTENT: ContentCompareDetailContentCompareResult{
			value: "INCONSISTENT",
		},
		COMPARING: ContentCompareDetailContentCompareResult{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: ContentCompareDetailContentCompareResult{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: ContentCompareDetailContentCompareResult{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIT: ContentCompareDetailContentCompareResult{
			value: "TARGET_DB_NOT_EXIT",
		},
		CAN_NOT_COMPARE: ContentCompareDetailContentCompareResult{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c ContentCompareDetailContentCompareResult) Value() string {
	return c.value
}

func (c ContentCompareDetailContentCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentCompareDetailContentCompareResult) UnmarshalJSON(b []byte) error {
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
