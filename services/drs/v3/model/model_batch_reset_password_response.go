package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchResetPasswordResponse struct {
	Results *[]ModifyDbPwdResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchResetPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetPasswordResponse struct{}"
	}

	return strings.Join([]string{"BatchResetPasswordResponse", string(data)}, " ")
}
