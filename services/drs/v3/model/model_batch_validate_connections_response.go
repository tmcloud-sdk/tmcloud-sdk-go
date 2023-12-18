package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchValidateConnectionsResponse struct {
	Results *[]CheckJobResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchValidateConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateConnectionsResponse struct{}"
	}

	return strings.Join([]string{"BatchValidateConnectionsResponse", string(data)}, " ")
}
