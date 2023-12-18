package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ContentCompareDiff struct {
	TargetSelectSql *string `json:"target_select_sql,omitempty"`

	SourceSelectSql *string `json:"source_select_sql,omitempty"`

	SourceKeyValue []string `json:"source_key_value"`

	TargetKeyValue []string `json:"target_key_value"`
}

func (o ContentCompareDiff) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareDiff struct{}"
	}

	return strings.Join([]string{"ContentCompareDiff", string(data)}, " ")
}
