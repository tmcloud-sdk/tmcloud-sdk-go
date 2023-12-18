package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchSetSpeedResponse struct {
	Count *int32 `json:"count,omitempty"`

	Results        *[]ModifyJobResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchSetSpeedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSpeedResponse struct{}"
	}

	return strings.Join([]string{"BatchSetSpeedResponse", string(data)}, " ")
}
