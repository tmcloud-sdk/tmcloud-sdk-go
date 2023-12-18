package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateKubernetesClusterCertResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Preferences *interface{} `json:"preferences,omitempty"`

	Clusters *[]Clusters `json:"clusters,omitempty"`

	Users *[]Users `json:"users,omitempty"`

	Contexts *[]Contexts `json:"contexts,omitempty"`

	CurrentContext *string `json:"current-context,omitempty"`

	PortID         *string `json:"Port-ID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKubernetesClusterCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKubernetesClusterCertResponse struct{}"
	}

	return strings.Join([]string{"CreateKubernetesClusterCertResponse", string(data)}, " ")
}
