package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateCustomTemplateResponse struct {
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCustomTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomTemplateResponse", string(data)}, " ")
}
