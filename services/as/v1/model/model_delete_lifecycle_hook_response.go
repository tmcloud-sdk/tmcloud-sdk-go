package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteLifecycleHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLifecycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLifecycleHookResponse struct{}"
	}

	return strings.Join([]string{"DeleteLifecycleHookResponse", string(data)}, " ")
}
