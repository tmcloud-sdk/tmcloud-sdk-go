package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StorageSelectors struct {
	Name string `json:"name"`

	StorageType string `json:"storageType"`

	MatchLabels *StorageSelectorsMatchLabels `json:"matchLabels,omitempty"`
}

func (o StorageSelectors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageSelectors struct{}"
	}

	return strings.Join([]string{"StorageSelectors", string(data)}, " ")
}
