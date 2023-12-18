package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DetailsBody struct {
	OldCapacity *string `json:"old_capacity,omitempty"`

	NewCapacity *string `json:"new_capacity,omitempty"`

	EnablePublicIp *bool `json:"enable_public_ip,omitempty"`

	PublicIpId *string `json:"public_ip_id,omitempty"`

	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	EnableSsl *bool `json:"enable_ssl,omitempty"`

	OldCacheMode *string `json:"old_cache_mode,omitempty"`

	NewCacheMode *string `json:"new_cache_mode,omitempty"`
}

func (o DetailsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailsBody struct{}"
	}

	return strings.Join([]string{"DetailsBody", string(data)}, " ")
}
