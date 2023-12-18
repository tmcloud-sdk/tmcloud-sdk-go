package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ContentCompareResultDetails struct {
	SourceDbName string `json:"source_db_name"`

	ContentCompareDetail []ContentCompareDetail `json:"content_compare_detail"`

	ContentCompareDetailCount int32 `json:"content_compare_detail_count"`

	ContentUncompareDetail *[]ContentCompareDetail `json:"content_uncompare_detail,omitempty"`

	ContentUncompareDetailCount int32 `json:"content_uncompare_detail_count"`
}

func (o ContentCompareResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareResultDetails struct{}"
	}

	return strings.Join([]string{"ContentCompareResultDetails", string(data)}, " ")
}
