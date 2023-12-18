package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchListProgressesResponse struct {
	Count *int32 `json:"count,omitempty"`

	Results        *[]QueryProgressResp `json:"results,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchListProgressesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListProgressesResponse struct{}"
	}

	return strings.Join([]string{"BatchListProgressesResponse", string(data)}, " ")
}
