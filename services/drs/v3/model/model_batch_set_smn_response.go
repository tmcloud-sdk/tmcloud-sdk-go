package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchSetSmnResponse struct {
	Results *[]ImportSmnResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchSetSmnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSmnResponse struct{}"
	}

	return strings.Join([]string{"BatchSetSmnResponse", string(data)}, " ")
}
