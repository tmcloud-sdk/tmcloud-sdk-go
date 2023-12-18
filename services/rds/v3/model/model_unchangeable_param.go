package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UnchangeableParam struct {
	LowerCaseTableNames *string `json:"lower_case_table_names,omitempty"`
}

func (o UnchangeableParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnchangeableParam struct{}"
	}

	return strings.Join([]string{"UnchangeableParam", string(data)}, " ")
}
