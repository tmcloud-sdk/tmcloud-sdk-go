package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ImportImageQuickResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportImageQuickResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageQuickResponse struct{}"
	}

	return strings.Join([]string{"ImportImageQuickResponse", string(data)}, " ")
}
