package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ResizeServerResponse struct {
	OrderId *string `json:"order_id,omitempty"`

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeServerResponse struct{}"
	}

	return strings.Join([]string{"ResizeServerResponse", string(data)}, " ")
}
