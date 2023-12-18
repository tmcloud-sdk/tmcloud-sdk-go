package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchShowParamsResponse struct {
	ParamsList *[]QueryDbParamsResp `json:"params_list,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchShowParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowParamsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowParamsResponse", string(data)}, " ")
}
