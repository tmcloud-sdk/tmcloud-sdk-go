package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListApiVersionNewResponse struct {
	Versions       *[]ApiVersion `json:"versions,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListApiVersionNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionNewResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionNewResponse", string(data)}, " ")
}
