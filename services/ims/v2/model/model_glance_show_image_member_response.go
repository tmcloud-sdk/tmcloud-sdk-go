package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceShowImageMemberResponse struct {
	Status *string `json:"status,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ImageId *string `json:"image_id,omitempty"`

	MemberId *string `json:"member_id,omitempty"`

	Schema         *string `json:"schema,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceShowImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberResponse", string(data)}, " ")
}
