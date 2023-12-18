package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostPaidServerPublicip struct {
	Id *string `json:"id,omitempty"`

	Eip *PostPaidServerEip `json:"eip,omitempty"`

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o PostPaidServerPublicip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerPublicip struct{}"
	}

	return strings.Join([]string{"PostPaidServerPublicip", string(data)}, " ")
}
