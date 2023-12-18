package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LineCompareResultDetails struct {
	SourceDbName string `json:"source_db_name"`

	LineCompareDetail []LineCompareDetail `json:"line_compare_detail"`

	LineCompareDetailCount int32 `json:"line_compare_detail_count"`
}

func (o LineCompareResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareResultDetails struct{}"
	}

	return strings.Join([]string{"LineCompareResultDetails", string(data)}, " ")
}
