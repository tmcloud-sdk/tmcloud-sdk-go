package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetSecondLevelMonitorResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSecondLevelMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSecondLevelMonitorResponse struct{}"
	}

	return strings.Join([]string{"SetSecondLevelMonitorResponse", string(data)}, " ")
}
