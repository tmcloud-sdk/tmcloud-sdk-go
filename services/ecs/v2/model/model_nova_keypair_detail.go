package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaKeypairDetail struct {
	PublicKey string `json:"public_key"`

	Name string `json:"name"`

	Fingerprint string `json:"fingerprint"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	Deleted bool `json:"deleted"`

	DeletedAt *sdktime.SdkTime `json:"deleted_at"`

	Id int32 `json:"id"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	UserId string `json:"user_id"`

	Type *string `json:"type,omitempty"`
}

func (o NovaKeypairDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaKeypairDetail struct{}"
	}

	return strings.Join([]string{"NovaKeypairDetail", string(data)}, " ")
}
