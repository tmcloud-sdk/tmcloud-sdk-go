package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchChangeDataResponse struct {
	Results *[]DataTransformationResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchChangeDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeDataResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeDataResponse", string(data)}, " ")
}
