package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type FlavorsItems struct {
	SpecCode *string `json:"spec_code,omitempty"`

	CloudServiceTypeCode *string `json:"cloud_service_type_code,omitempty"`

	CloudResourceTypeCode *string `json:"cloud_resource_type_code,omitempty"`

	CacheMode *string `json:"cache_mode,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	ProductType *string `json:"product_type,omitempty"`

	CpuType *string `json:"cpu_type,omitempty"`

	StorageType *string `json:"storage_type,omitempty"`

	Capacity *[]string `json:"capacity,omitempty"`

	BillingMode *[]string `json:"billing_mode,omitempty"`

	TenantIpCount *int32 `json:"tenant_ip_count,omitempty"`

	PricingType *string `json:"pricing_type,omitempty"`

	IsDec *bool `json:"is_dec,omitempty"`

	Attrs *[]AttrsObject `json:"attrs,omitempty"`

	FlavorsAvailableZones *[]FlavorAzObject `json:"flavors_available_zones,omitempty"`
}

func (o FlavorsItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorsItems struct{}"
	}

	return strings.Join([]string{"FlavorsItems", string(data)}, " ")
}
