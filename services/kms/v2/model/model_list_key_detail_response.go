package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListKeyDetailResponse struct {
	KeyInfo        *KeyDetails `json:"key_info,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListKeyDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyDetailResponse struct{}"
	}

	return strings.Join([]string{"ListKeyDetailResponse", string(data)}, " ")
}
