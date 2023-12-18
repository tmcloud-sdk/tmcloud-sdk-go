package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchListJobDetailsResponse struct {
	Count *int32 `json:"count,omitempty"`

	Results        *[]QueryJobResp `json:"results,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o BatchListJobDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobDetailsResponse struct{}"
	}

	return strings.Join([]string{"BatchListJobDetailsResponse", string(data)}, " ")
}
