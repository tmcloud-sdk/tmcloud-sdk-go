package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceDeleteImageMemberRequest struct {
	ImageId string `json:"image_id"`

	MemberId string `json:"member_id"`
}

func (o GlanceDeleteImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageMemberRequest", string(data)}, " ")
}
