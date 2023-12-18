package v2

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/evs/v2/model"
)

type EvsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEvsClient(hcClient *http_client.HcHttpClient) *EvsClient {
	return &EvsClient{HcClient: hcClient}
}

func EvsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *EvsClient) BatchCreateVolumeTags(request *model.BatchCreateVolumeTagsRequest) (*model.BatchCreateVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateVolumeTagsResponse), nil
	}
}

func (c *EvsClient) BatchCreateVolumeTagsInvoker(request *model.BatchCreateVolumeTagsRequest) *BatchCreateVolumeTagsInvoker {
	requestDef := GenReqDefForBatchCreateVolumeTags()
	return &BatchCreateVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) BatchDeleteVolumeTags(request *model.BatchDeleteVolumeTagsRequest) (*model.BatchDeleteVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteVolumeTagsResponse), nil
	}
}

func (c *EvsClient) BatchDeleteVolumeTagsInvoker(request *model.BatchDeleteVolumeTagsRequest) *BatchDeleteVolumeTagsInvoker {
	requestDef := GenReqDefForBatchDeleteVolumeTags()
	return &BatchDeleteVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderAcceptVolumeTransfer(request *model.CinderAcceptVolumeTransferRequest) (*model.CinderAcceptVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderAcceptVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderAcceptVolumeTransferResponse), nil
	}
}

func (c *EvsClient) CinderAcceptVolumeTransferInvoker(request *model.CinderAcceptVolumeTransferRequest) *CinderAcceptVolumeTransferInvoker {
	requestDef := GenReqDefForCinderAcceptVolumeTransfer()
	return &CinderAcceptVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderCreateVolumeTransfer(request *model.CinderCreateVolumeTransferRequest) (*model.CinderCreateVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderCreateVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderCreateVolumeTransferResponse), nil
	}
}

func (c *EvsClient) CinderCreateVolumeTransferInvoker(request *model.CinderCreateVolumeTransferRequest) *CinderCreateVolumeTransferInvoker {
	requestDef := GenReqDefForCinderCreateVolumeTransfer()
	return &CinderCreateVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderDeleteVolumeTransfer(request *model.CinderDeleteVolumeTransferRequest) (*model.CinderDeleteVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderDeleteVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderDeleteVolumeTransferResponse), nil
	}
}

func (c *EvsClient) CinderDeleteVolumeTransferInvoker(request *model.CinderDeleteVolumeTransferRequest) *CinderDeleteVolumeTransferInvoker {
	requestDef := GenReqDefForCinderDeleteVolumeTransfer()
	return &CinderDeleteVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderListAvailabilityZones(request *model.CinderListAvailabilityZonesRequest) (*model.CinderListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForCinderListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListAvailabilityZonesResponse), nil
	}
}

func (c *EvsClient) CinderListAvailabilityZonesInvoker(request *model.CinderListAvailabilityZonesRequest) *CinderListAvailabilityZonesInvoker {
	requestDef := GenReqDefForCinderListAvailabilityZones()
	return &CinderListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderListQuotas(request *model.CinderListQuotasRequest) (*model.CinderListQuotasResponse, error) {
	requestDef := GenReqDefForCinderListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListQuotasResponse), nil
	}
}

func (c *EvsClient) CinderListQuotasInvoker(request *model.CinderListQuotasRequest) *CinderListQuotasInvoker {
	requestDef := GenReqDefForCinderListQuotas()
	return &CinderListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderListVolumeTransfers(request *model.CinderListVolumeTransfersRequest) (*model.CinderListVolumeTransfersResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTransfers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTransfersResponse), nil
	}
}

func (c *EvsClient) CinderListVolumeTransfersInvoker(request *model.CinderListVolumeTransfersRequest) *CinderListVolumeTransfersInvoker {
	requestDef := GenReqDefForCinderListVolumeTransfers()
	return &CinderListVolumeTransfersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderListVolumeTypes(request *model.CinderListVolumeTypesRequest) (*model.CinderListVolumeTypesResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTypesResponse), nil
	}
}

func (c *EvsClient) CinderListVolumeTypesInvoker(request *model.CinderListVolumeTypesRequest) *CinderListVolumeTypesInvoker {
	requestDef := GenReqDefForCinderListVolumeTypes()
	return &CinderListVolumeTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderShowVolumeTransfer(request *model.CinderShowVolumeTransferRequest) (*model.CinderShowVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderShowVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderShowVolumeTransferResponse), nil
	}
}

func (c *EvsClient) CinderShowVolumeTransferInvoker(request *model.CinderShowVolumeTransferRequest) *CinderShowVolumeTransferInvoker {
	requestDef := GenReqDefForCinderShowVolumeTransfer()
	return &CinderShowVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListVolumeTags(request *model.ListVolumeTagsRequest) (*model.ListVolumeTagsResponse, error) {
	requestDef := GenReqDefForListVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumeTagsResponse), nil
	}
}

func (c *EvsClient) ListVolumeTagsInvoker(request *model.ListVolumeTagsRequest) *ListVolumeTagsInvoker {
	requestDef := GenReqDefForListVolumeTags()
	return &ListVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListVolumes(request *model.ListVolumesRequest) (*model.ListVolumesResponse, error) {
	requestDef := GenReqDefForListVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesResponse), nil
	}
}

func (c *EvsClient) ListVolumesInvoker(request *model.ListVolumesRequest) *ListVolumesInvoker {
	requestDef := GenReqDefForListVolumes()
	return &ListVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListVolumesByTags(request *model.ListVolumesByTagsRequest) (*model.ListVolumesByTagsResponse, error) {
	requestDef := GenReqDefForListVolumesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesByTagsResponse), nil
	}
}

func (c *EvsClient) ListVolumesByTagsInvoker(request *model.ListVolumesByTagsRequest) *ListVolumesByTagsInvoker {
	requestDef := GenReqDefForListVolumesByTags()
	return &ListVolumesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

func (c *EvsClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ShowVolumeTags(request *model.ShowVolumeTagsRequest) (*model.ShowVolumeTagsResponse, error) {
	requestDef := GenReqDefForShowVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeTagsResponse), nil
	}
}

func (c *EvsClient) ShowVolumeTagsInvoker(request *model.ShowVolumeTagsRequest) *ShowVolumeTagsInvoker {
	requestDef := GenReqDefForShowVolumeTags()
	return &ShowVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListVersions(request *model.ListVersionsRequest) (*model.ListVersionsResponse, error) {
	requestDef := GenReqDefForListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionsResponse), nil
	}
}

func (c *EvsClient) ListVersionsInvoker(request *model.ListVersionsRequest) *ListVersionsInvoker {
	requestDef := GenReqDefForListVersions()
	return &ListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

func (c *EvsClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
