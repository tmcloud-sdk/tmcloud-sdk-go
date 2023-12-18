package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LinksInfoResponse struct {
	Href *string `json:"href,omitempty"`

	Rel *string `json:"rel,omitempty"`
}

func (o LinksInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksInfoResponse struct{}"
	}

	return strings.Join([]string{"LinksInfoResponse", string(data)}, " ")
}
