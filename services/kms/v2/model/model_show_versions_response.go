package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowVersionsResponse struct {
	Versions       *[]ApiVersionDetail `json:"versions,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionsResponse struct{}"
	}

	return strings.Join([]string{"ShowVersionsResponse", string(data)}, " ")
}
