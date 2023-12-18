package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceShowImageMemberSchemasResponse struct {
	Name *string `json:"name,omitempty"`

	Properties     *interface{} `json:"properties,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o GlanceShowImageMemberSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberSchemasResponse struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberSchemasResponse", string(data)}, " ")
}
