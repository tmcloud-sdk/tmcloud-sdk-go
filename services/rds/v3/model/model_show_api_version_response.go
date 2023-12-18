package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowApiVersionResponse struct {
	Versions *ApiVersion `json:"versions,omitempty"`

	Version        *ApiVersion `json:"version,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowApiVersionResponse", string(data)}, " ")
}
