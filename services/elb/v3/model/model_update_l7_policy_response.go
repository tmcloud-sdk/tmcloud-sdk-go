package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7PolicyResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	L7policy       *L7Policy `json:"l7policy,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateL7PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyResponse", string(data)}, " ")
}
