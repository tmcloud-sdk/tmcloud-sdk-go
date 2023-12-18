package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type QueryAvailableNodeTypeReq struct {
	EngineType string `json:"engine_type"`

	DbUseType QueryAvailableNodeTypeReqDbUseType `json:"db_use_type"`

	JobDirection QueryAvailableNodeTypeReqJobDirection `json:"job_direction"`

	NodeType string `json:"node_type"`

	MultiWrite *string `json:"multi_write,omitempty"`
}

func (o QueryAvailableNodeTypeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAvailableNodeTypeReq struct{}"
	}

	return strings.Join([]string{"QueryAvailableNodeTypeReq", string(data)}, " ")
}

type QueryAvailableNodeTypeReqDbUseType struct {
	value string
}

type QueryAvailableNodeTypeReqDbUseTypeEnum struct {
	MIGRATION        QueryAvailableNodeTypeReqDbUseType
	SYNC             QueryAvailableNodeTypeReqDbUseType
	CLOUD_DATA_GUARD QueryAvailableNodeTypeReqDbUseType
}

func GetQueryAvailableNodeTypeReqDbUseTypeEnum() QueryAvailableNodeTypeReqDbUseTypeEnum {
	return QueryAvailableNodeTypeReqDbUseTypeEnum{
		MIGRATION: QueryAvailableNodeTypeReqDbUseType{
			value: "migration",
		},
		SYNC: QueryAvailableNodeTypeReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: QueryAvailableNodeTypeReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c QueryAvailableNodeTypeReqDbUseType) Value() string {
	return c.value
}

func (c QueryAvailableNodeTypeReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryAvailableNodeTypeReqDbUseType) UnmarshalJSON(b []byte) error {
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

type QueryAvailableNodeTypeReqJobDirection struct {
	value string
}

type QueryAvailableNodeTypeReqJobDirectionEnum struct {
	UP      QueryAvailableNodeTypeReqJobDirection
	DOWN    QueryAvailableNodeTypeReqJobDirection
	NON_DBS QueryAvailableNodeTypeReqJobDirection
}

func GetQueryAvailableNodeTypeReqJobDirectionEnum() QueryAvailableNodeTypeReqJobDirectionEnum {
	return QueryAvailableNodeTypeReqJobDirectionEnum{
		UP: QueryAvailableNodeTypeReqJobDirection{
			value: "up",
		},
		DOWN: QueryAvailableNodeTypeReqJobDirection{
			value: "down",
		},
		NON_DBS: QueryAvailableNodeTypeReqJobDirection{
			value: "non-dbs",
		},
	}
}

func (c QueryAvailableNodeTypeReqJobDirection) Value() string {
	return c.value
}

func (c QueryAvailableNodeTypeReqJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryAvailableNodeTypeReqJobDirection) UnmarshalJSON(b []byte) error {
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
