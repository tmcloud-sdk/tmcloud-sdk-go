package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchSwitchoverResponse struct {
	Results *[]SwitchoverResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchSwitchoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSwitchoverResponse struct{}"
	}

	return strings.Join([]string{"BatchSwitchoverResponse", string(data)}, " ")
}
