package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateKubernetesClusterCertRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *CertDuration `json:"body,omitempty"`
}

func (o CreateKubernetesClusterCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKubernetesClusterCertRequest struct{}"
	}

	return strings.Join([]string{"CreateKubernetesClusterCertRequest", string(data)}, " ")
}
