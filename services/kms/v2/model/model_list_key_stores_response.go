package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListKeyStoresResponse struct {
	Total *int32 `json:"total,omitempty"`

	Keystores      *[]KeystoreDetails `json:"keystores,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListKeyStoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyStoresResponse struct{}"
	}

	return strings.Join([]string{"ListKeyStoresResponse", string(data)}, " ")
}
