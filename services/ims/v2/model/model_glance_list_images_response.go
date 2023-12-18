package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceListImagesResponse struct {
	First *string `json:"first,omitempty"`

	Images *[]GlanceShowImageResponseBody `json:"images,omitempty"`

	Schema *string `json:"schema,omitempty"`

	Next           *string `json:"next,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceListImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImagesResponse struct{}"
	}

	return strings.Join([]string{"GlanceListImagesResponse", string(data)}, " ")
}
