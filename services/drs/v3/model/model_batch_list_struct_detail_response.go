package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchListStructDetailResponse struct {
	Count *int32 `json:"count,omitempty"`

	Results        *[]QueryStructDetailResp `json:"results,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchListStructDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListStructDetailResponse struct{}"
	}

	return strings.Join([]string{"BatchListStructDetailResponse", string(data)}, " ")
}
