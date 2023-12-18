package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchListStructProcessResponse struct {
	Results *[]QueryStructProcessResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchListStructProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListStructProcessResponse struct{}"
	}

	return strings.Join([]string{"BatchListStructProcessResponse", string(data)}, " ")
}
