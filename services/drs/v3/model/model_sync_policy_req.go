package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SyncPolicyReq struct {
	JobId string `json:"job_id"`

	ConflictPolicy *SyncPolicyReqConflictPolicy `json:"conflict_policy,omitempty"`

	FilterDdlPolicy *SyncPolicyReqFilterDdlPolicy `json:"filter_ddl_policy,omitempty"`

	DdlTrans *bool `json:"ddl_trans,omitempty"`

	IndexTrans *bool `json:"index_trans,omitempty"`

	TopicPolicy *SyncPolicyReqTopicPolicy `json:"topic_policy,omitempty"`

	Topic *string `json:"topic,omitempty"`

	PartitionPolicy *SyncPolicyReqPartitionPolicy `json:"partition_policy,omitempty"`

	KafkaDataFormat *SyncPolicyReqKafkaDataFormat `json:"kafka_data_format,omitempty"`

	TopicNameFormat *string `json:"topic_name_format,omitempty"`

	PartitionsNum *string `json:"partitions_num,omitempty"`

	ReplicationFactor *string `json:"replication_factor,omitempty"`

	IsFillMaterializedView *bool `json:"is_fill_materialized_view,omitempty"`

	ExportSnapshot *bool `json:"export_snapshot,omitempty"`

	SlotName *string `json:"slot_name,omitempty"`
}

func (o SyncPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncPolicyReq struct{}"
	}

	return strings.Join([]string{"SyncPolicyReq", string(data)}, " ")
}

type SyncPolicyReqConflictPolicy struct {
	value string
}

type SyncPolicyReqConflictPolicyEnum struct {
	IGNORE    SyncPolicyReqConflictPolicy
	OVERWRITE SyncPolicyReqConflictPolicy
	STOP      SyncPolicyReqConflictPolicy
}

func GetSyncPolicyReqConflictPolicyEnum() SyncPolicyReqConflictPolicyEnum {
	return SyncPolicyReqConflictPolicyEnum{
		IGNORE: SyncPolicyReqConflictPolicy{
			value: "ignore",
		},
		OVERWRITE: SyncPolicyReqConflictPolicy{
			value: "overwrite",
		},
		STOP: SyncPolicyReqConflictPolicy{
			value: "stop",
		},
	}
}

func (c SyncPolicyReqConflictPolicy) Value() string {
	return c.value
}

func (c SyncPolicyReqConflictPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqConflictPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqFilterDdlPolicy struct {
	value string
}

type SyncPolicyReqFilterDdlPolicyEnum struct {
	DROP_DATABASE SyncPolicyReqFilterDdlPolicy
}

func GetSyncPolicyReqFilterDdlPolicyEnum() SyncPolicyReqFilterDdlPolicyEnum {
	return SyncPolicyReqFilterDdlPolicyEnum{
		DROP_DATABASE: SyncPolicyReqFilterDdlPolicy{
			value: "drop_database",
		},
	}
}

func (c SyncPolicyReqFilterDdlPolicy) Value() string {
	return c.value
}

func (c SyncPolicyReqFilterDdlPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqFilterDdlPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqTopicPolicy struct {
	value string
}

type SyncPolicyReqTopicPolicyEnum struct {
	E_0 SyncPolicyReqTopicPolicy
	E_1 SyncPolicyReqTopicPolicy
	E_2 SyncPolicyReqTopicPolicy
	E_3 SyncPolicyReqTopicPolicy
}

func GetSyncPolicyReqTopicPolicyEnum() SyncPolicyReqTopicPolicyEnum {
	return SyncPolicyReqTopicPolicyEnum{
		E_0: SyncPolicyReqTopicPolicy{
			value: "0",
		},
		E_1: SyncPolicyReqTopicPolicy{
			value: "1",
		},
		E_2: SyncPolicyReqTopicPolicy{
			value: "2",
		},
		E_3: SyncPolicyReqTopicPolicy{
			value: "3",
		},
	}
}

func (c SyncPolicyReqTopicPolicy) Value() string {
	return c.value
}

func (c SyncPolicyReqTopicPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqTopicPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqPartitionPolicy struct {
	value string
}

type SyncPolicyReqPartitionPolicyEnum struct {
	E_0 SyncPolicyReqPartitionPolicy
	E_1 SyncPolicyReqPartitionPolicy
	E_2 SyncPolicyReqPartitionPolicy
	E_3 SyncPolicyReqPartitionPolicy
}

func GetSyncPolicyReqPartitionPolicyEnum() SyncPolicyReqPartitionPolicyEnum {
	return SyncPolicyReqPartitionPolicyEnum{
		E_0: SyncPolicyReqPartitionPolicy{
			value: "0",
		},
		E_1: SyncPolicyReqPartitionPolicy{
			value: "1",
		},
		E_2: SyncPolicyReqPartitionPolicy{
			value: "2",
		},
		E_3: SyncPolicyReqPartitionPolicy{
			value: "3",
		},
	}
}

func (c SyncPolicyReqPartitionPolicy) Value() string {
	return c.value
}

func (c SyncPolicyReqPartitionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqPartitionPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqKafkaDataFormat struct {
	value string
}

type SyncPolicyReqKafkaDataFormatEnum struct {
	JSON SyncPolicyReqKafkaDataFormat
	AVRO SyncPolicyReqKafkaDataFormat
}

func GetSyncPolicyReqKafkaDataFormatEnum() SyncPolicyReqKafkaDataFormatEnum {
	return SyncPolicyReqKafkaDataFormatEnum{
		JSON: SyncPolicyReqKafkaDataFormat{
			value: "json",
		},
		AVRO: SyncPolicyReqKafkaDataFormat{
			value: "avro",
		},
	}
}

func (c SyncPolicyReqKafkaDataFormat) Value() string {
	return c.value
}

func (c SyncPolicyReqKafkaDataFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqKafkaDataFormat) UnmarshalJSON(b []byte) error {
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
