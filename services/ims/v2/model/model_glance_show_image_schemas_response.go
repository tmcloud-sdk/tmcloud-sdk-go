package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceShowImageSchemasResponse struct {
	AdditionalProperties *AdditionalProperties `json:"additionalProperties,omitempty"`

	Name *string `json:"name,omitempty"`

	Properties *interface{} `json:"properties,omitempty"`

	Links          *[]Links `json:"links,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o GlanceShowImageSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageSchemasResponse struct{}"
	}

	return strings.Join([]string{"GlanceShowImageSchemasResponse", string(data)}, " ")
}
