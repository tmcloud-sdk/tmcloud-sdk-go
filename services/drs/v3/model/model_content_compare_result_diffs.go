package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ContentCompareResultDiffs struct {
	SourceDbName string `json:"source_db_name"`

	SourceTableName string `json:"source_table_name"`

	ContentCompareDiff []ContentCompareDiff `json:"content_compare_diff"`

	ContentCompareDiffCount int32 `json:"content_compare_diff_count"`
}

func (o ContentCompareResultDiffs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareResultDiffs struct{}"
	}

	return strings.Join([]string{"ContentCompareResultDiffs", string(data)}, " ")
}
