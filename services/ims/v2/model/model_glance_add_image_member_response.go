package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceAddImageMemberResponse struct {
	Status *string `json:"status,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ImageId *string `json:"image_id,omitempty"`

	MemberId *string `json:"member_id,omitempty"`

	Schema         *string `json:"schema,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceAddImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceAddImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceAddImageMemberResponse", string(data)}, " ")
}
