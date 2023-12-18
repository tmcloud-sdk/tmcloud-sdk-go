package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateCustomTemplateBody struct {
	TemplateId string `json:"template_id"`

	Name string `json:"name"`

	Type string `json:"type"`

	Engine *string `json:"engine,omitempty"`

	CacheMode *string `json:"cache_mode,omitempty"`

	Description *string `json:"description,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	Params map[string]string `json:"params"`
}

func (o CreateCustomTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomTemplateBody struct{}"
	}

	return strings.Join([]string{"CreateCustomTemplateBody", string(data)}, " ")
}
