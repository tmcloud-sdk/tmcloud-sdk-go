package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryFlowCompareDataResp struct {
	TotalRecord *int64 `json:"total_record,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	List *[]StructDetailVo `json:"list,omitempty"`
}

func (o QueryFlowCompareDataResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryFlowCompareDataResp struct{}"
	}

	return strings.Join([]string{"QueryFlowCompareDataResp", string(data)}, " ")
}
