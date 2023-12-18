package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StructProcessVo struct {
	Type string `json:"type"`

	Status int32 `json:"status"`

	SrcCount int32 `json:"src_count"`

	DstCount int32 `json:"dst_count"`

	StartTime int64 `json:"start_time"`

	EndTime *int64 `json:"end_time,omitempty"`
}

func (o StructProcessVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructProcessVo struct{}"
	}

	return strings.Join([]string{"StructProcessVo", string(data)}, " ")
}
