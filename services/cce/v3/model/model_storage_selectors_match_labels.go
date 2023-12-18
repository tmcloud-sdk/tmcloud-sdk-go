package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StorageSelectorsMatchLabels struct {
	Size *string `json:"size,omitempty"`

	VolumeType *string `json:"volumeType,omitempty"`

	MetadataEncrypted *string `json:"metadataEncrypted,omitempty"`

	MetadataCmkid *string `json:"metadataCmkid,omitempty"`

	Count *string `json:"count,omitempty"`
}

func (o StorageSelectorsMatchLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageSelectorsMatchLabels struct{}"
	}

	return strings.Join([]string{"StorageSelectorsMatchLabels", string(data)}, " ")
}
