package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ExportImageResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageResponse struct{}"
	}

	return strings.Join([]string{"ExportImageResponse", string(data)}, " ")
}
