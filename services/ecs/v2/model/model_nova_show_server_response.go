package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaShowServerResponse struct {
	Server         *NovaServer `json:"server,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o NovaShowServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowServerResponse struct{}"
	}

	return strings.Join([]string{"NovaShowServerResponse", string(data)}, " ")
}
