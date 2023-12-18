package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchResetServersPasswordResponse struct {
	Response       *[]ServerId `json:"response,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o BatchResetServersPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetServersPasswordResponse struct{}"
	}

	return strings.Join([]string{"BatchResetServersPasswordResponse", string(data)}, " ")
}
