package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ConfigTransformationVo struct {
	DbTableName string `json:"db_table_name"`

	DbName string `json:"db_name"`

	TableName string `json:"table_name"`

	ColNames string `json:"col_names"`

	PrimKeyOrIndex string `json:"prim_key_or_index"`

	Indexs string `json:"indexs"`

	Values string `json:"values"`
}

func (o ConfigTransformationVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigTransformationVo struct{}"
	}

	return strings.Join([]string{"ConfigTransformationVo", string(data)}, " ")
}
