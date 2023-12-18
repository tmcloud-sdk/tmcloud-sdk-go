package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CancelGrantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelGrantResponse struct{}"
	}

	return strings.Join([]string{"CancelGrantResponse", string(data)}, " ")
}
