package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type HistoryInfo struct {
	HistoryId *string `json:"history_id,omitempty"`

	Type *string `json:"type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o HistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryInfo struct{}"
	}

	return strings.Join([]string{"HistoryInfo", string(data)}, " ")
}
