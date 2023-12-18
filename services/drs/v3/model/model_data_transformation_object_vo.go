package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DataTransformationObjectVo struct {
	Id *string `json:"id,omitempty"`

	DataTransformationType *string `json:"data_transformation_type,omitempty"`

	DbName *string `json:"db_name,omitempty"`

	SchemaName *string `json:"schema_name,omitempty"`

	TableName *string `json:"table_name,omitempty"`
}

func (o DataTransformationObjectVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataTransformationObjectVo struct{}"
	}

	return strings.Join([]string{"DataTransformationObjectVo", string(data)}, " ")
}
