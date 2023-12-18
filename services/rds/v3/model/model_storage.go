package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Storage struct {
	Name string `json:"name"`

	AzStatus map[string]string `json:"az_status"`

	SupportComputeGroupType *[]string `json:"support_compute_group_type,omitempty"`
}

func (o Storage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Storage struct{}"
	}

	return strings.Join([]string{"Storage", string(data)}, " ")
}
