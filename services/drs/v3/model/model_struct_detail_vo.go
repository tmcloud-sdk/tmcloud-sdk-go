package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StructDetailVo struct {
	Progress *int32 `json:"progress,omitempty"`

	SrcDB *string `json:"src_DB,omitempty"`

	SrcTB *string `json:"src_TB,omitempty"`

	DstDB *string `json:"dst_DB,omitempty"`

	DstTB *string `json:"dst_TB,omitempty"`
}

func (o StructDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructDetailVo struct{}"
	}

	return strings.Join([]string{"StructDetailVo", string(data)}, " ")
}
