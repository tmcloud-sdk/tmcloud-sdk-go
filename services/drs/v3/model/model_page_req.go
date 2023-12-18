package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PageReq struct {
	CurPage *int32 `json:"cur_page,omitempty"`

	PerPage *int32 `json:"per_page,omitempty"`
}

func (o PageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageReq struct{}"
	}

	return strings.Join([]string{"PageReq", string(data)}, " ")
}
