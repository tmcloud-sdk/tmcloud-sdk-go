package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceListImageMemberSchemasResponse struct {
	Links *[]Links `json:"links,omitempty"`

	Name *string `json:"name,omitempty"`

	Properties     *interface{} `json:"properties,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o GlanceListImageMemberSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageMemberSchemasResponse struct{}"
	}

	return strings.Join([]string{"GlanceListImageMemberSchemasResponse", string(data)}, " ")
}
