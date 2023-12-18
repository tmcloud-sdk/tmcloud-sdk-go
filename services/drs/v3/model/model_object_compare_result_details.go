package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ObjectCompareResultDetails struct {
	SourceDbName string `json:"source_db_name"`

	TargetDbName string `json:"target_db_name"`

	SourceDbValue *string `json:"source_db_value,omitempty"`

	TargetDbValue *string `json:"target_db_value,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ObjectCompareResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectCompareResultDetails struct{}"
	}

	return strings.Join([]string{"ObjectCompareResultDetails", string(data)}, " ")
}
