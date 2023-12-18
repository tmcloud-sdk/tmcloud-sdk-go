package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateCustomTemplateRequest struct {
	Body *CreateCustomTemplateBody `json:"body,omitempty"`
}

func (o CreateCustomTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomTemplateRequest", string(data)}, " ")
}
