package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ObjectCompareResultOverview struct {
	ObjectType ObjectCompareResultOverviewObjectType `json:"object_type"`

	ObjectCompareResult ObjectCompareResultOverviewObjectCompareResult `json:"object_compare_result"`

	TargetCount int32 `json:"target_count"`

	SourceCount int32 `json:"source_count"`

	DiffCount int32 `json:"diff_count"`
}

func (o ObjectCompareResultOverview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectCompareResultOverview struct{}"
	}

	return strings.Join([]string{"ObjectCompareResultOverview", string(data)}, " ")
}

type ObjectCompareResultOverviewObjectType struct {
	value string
}

type ObjectCompareResultOverviewObjectTypeEnum struct {
	DB                 ObjectCompareResultOverviewObjectType
	TABLE              ObjectCompareResultOverviewObjectType
	VIEW               ObjectCompareResultOverviewObjectType
	EVENT              ObjectCompareResultOverviewObjectType
	ROUTINE            ObjectCompareResultOverviewObjectType
	INDEX              ObjectCompareResultOverviewObjectType
	TRIGGER            ObjectCompareResultOverviewObjectType
	SYNONYM            ObjectCompareResultOverviewObjectType
	FUNCTION           ObjectCompareResultOverviewObjectType
	PROCEDURE          ObjectCompareResultOverviewObjectType
	TYPE               ObjectCompareResultOverviewObjectType
	RULE               ObjectCompareResultOverviewObjectType
	DEFAULT_TYPE       ObjectCompareResultOverviewObjectType
	PLAN_GUIDE         ObjectCompareResultOverviewObjectType
	CONSTRAINT         ObjectCompareResultOverviewObjectType
	FILE_GROUP         ObjectCompareResultOverviewObjectType
	PARTITION_FUNCTION ObjectCompareResultOverviewObjectType
	PARTITION_SCHEME   ObjectCompareResultOverviewObjectType
	TABLE_COLLATION    ObjectCompareResultOverviewObjectType
}

func GetObjectCompareResultOverviewObjectTypeEnum() ObjectCompareResultOverviewObjectTypeEnum {
	return ObjectCompareResultOverviewObjectTypeEnum{
		DB: ObjectCompareResultOverviewObjectType{
			value: "DB",
		},
		TABLE: ObjectCompareResultOverviewObjectType{
			value: "TABLE",
		},
		VIEW: ObjectCompareResultOverviewObjectType{
			value: "VIEW",
		},
		EVENT: ObjectCompareResultOverviewObjectType{
			value: "EVENT",
		},
		ROUTINE: ObjectCompareResultOverviewObjectType{
			value: "ROUTINE",
		},
		INDEX: ObjectCompareResultOverviewObjectType{
			value: "INDEX",
		},
		TRIGGER: ObjectCompareResultOverviewObjectType{
			value: "TRIGGER",
		},
		SYNONYM: ObjectCompareResultOverviewObjectType{
			value: "SYNONYM",
		},
		FUNCTION: ObjectCompareResultOverviewObjectType{
			value: "FUNCTION",
		},
		PROCEDURE: ObjectCompareResultOverviewObjectType{
			value: "PROCEDURE",
		},
		TYPE: ObjectCompareResultOverviewObjectType{
			value: "TYPE",
		},
		RULE: ObjectCompareResultOverviewObjectType{
			value: "RULE",
		},
		DEFAULT_TYPE: ObjectCompareResultOverviewObjectType{
			value: "DEFAULT_TYPE",
		},
		PLAN_GUIDE: ObjectCompareResultOverviewObjectType{
			value: "PLAN_GUIDE",
		},
		CONSTRAINT: ObjectCompareResultOverviewObjectType{
			value: "CONSTRAINT",
		},
		FILE_GROUP: ObjectCompareResultOverviewObjectType{
			value: "FILE_GROUP",
		},
		PARTITION_FUNCTION: ObjectCompareResultOverviewObjectType{
			value: "PARTITION_FUNCTION",
		},
		PARTITION_SCHEME: ObjectCompareResultOverviewObjectType{
			value: "PARTITION_SCHEME",
		},
		TABLE_COLLATION: ObjectCompareResultOverviewObjectType{
			value: "TABLE_COLLATION",
		},
	}
}

func (c ObjectCompareResultOverviewObjectType) Value() string {
	return c.value
}

func (c ObjectCompareResultOverviewObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectCompareResultOverviewObjectType) UnmarshalJSON(b []byte) error {
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

type ObjectCompareResultOverviewObjectCompareResult struct {
	value string
}

type ObjectCompareResultOverviewObjectCompareResultEnum struct {
	CONSISTENT             ObjectCompareResultOverviewObjectCompareResult
	INCONSISTENT           ObjectCompareResultOverviewObjectCompareResult
	COMPARING              ObjectCompareResultOverviewObjectCompareResult
	WAITING_FOR_COMPARISON ObjectCompareResultOverviewObjectCompareResult
	FAILED_TO_COMPARE      ObjectCompareResultOverviewObjectCompareResult
	TARGET_DB_NOT_EXIT     ObjectCompareResultOverviewObjectCompareResult
	CAN_NOT_COMPARE        ObjectCompareResultOverviewObjectCompareResult
}

func GetObjectCompareResultOverviewObjectCompareResultEnum() ObjectCompareResultOverviewObjectCompareResultEnum {
	return ObjectCompareResultOverviewObjectCompareResultEnum{
		CONSISTENT: ObjectCompareResultOverviewObjectCompareResult{
			value: "CONSISTENT",
		},
		INCONSISTENT: ObjectCompareResultOverviewObjectCompareResult{
			value: "INCONSISTENT",
		},
		COMPARING: ObjectCompareResultOverviewObjectCompareResult{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: ObjectCompareResultOverviewObjectCompareResult{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: ObjectCompareResultOverviewObjectCompareResult{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIT: ObjectCompareResultOverviewObjectCompareResult{
			value: "TARGET_DB_NOT_EXIT",
		},
		CAN_NOT_COMPARE: ObjectCompareResultOverviewObjectCompareResult{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c ObjectCompareResultOverviewObjectCompareResult) Value() string {
	return c.value
}

func (c ObjectCompareResultOverviewObjectCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectCompareResultOverviewObjectCompareResult) UnmarshalJSON(b []byte) error {
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
