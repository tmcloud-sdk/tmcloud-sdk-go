package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceListImageSchemasResponse struct {
	Name *string `json:"name,omitempty"`

	Properties *interface{} `json:"properties,omitempty"`

	Links          *[]Links `json:"links,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o GlanceListImageSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageSchemasResponse struct{}"
	}

	return strings.Join([]string{"GlanceListImageSchemasResponse", string(data)}, " ")
}
