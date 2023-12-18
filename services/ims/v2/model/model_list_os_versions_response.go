package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListOsVersionsResponse struct {
	Body           *[]ListOsVersionsResponseBody `json:"body,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListOsVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOsVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListOsVersionsResponse", string(data)}, " ")
}
