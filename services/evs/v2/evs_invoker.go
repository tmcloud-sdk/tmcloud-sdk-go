package v2

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/evs/v2/model"
)

type BatchCreateVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateVolumeTagsInvoker) Invoke() (*model.BatchCreateVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateVolumeTagsResponse), nil
	}
}

type BatchDeleteVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteVolumeTagsInvoker) Invoke() (*model.BatchDeleteVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteVolumeTagsResponse), nil
	}
}

type CinderAcceptVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderAcceptVolumeTransferInvoker) Invoke() (*model.CinderAcceptVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderAcceptVolumeTransferResponse), nil
	}
}

type CinderCreateVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderCreateVolumeTransferInvoker) Invoke() (*model.CinderCreateVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderCreateVolumeTransferResponse), nil
	}
}

type CinderDeleteVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderDeleteVolumeTransferInvoker) Invoke() (*model.CinderDeleteVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderDeleteVolumeTransferResponse), nil
	}
}

type CinderListAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListAvailabilityZonesInvoker) Invoke() (*model.CinderListAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListAvailabilityZonesResponse), nil
	}
}

type CinderListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListQuotasInvoker) Invoke() (*model.CinderListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListQuotasResponse), nil
	}
}

type CinderListVolumeTransfersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListVolumeTransfersInvoker) Invoke() (*model.CinderListVolumeTransfersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListVolumeTransfersResponse), nil
	}
}

type CinderListVolumeTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListVolumeTypesInvoker) Invoke() (*model.CinderListVolumeTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListVolumeTypesResponse), nil
	}
}

type CinderShowVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderShowVolumeTransferInvoker) Invoke() (*model.CinderShowVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderShowVolumeTransferResponse), nil
	}
}

type ListVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumeTagsInvoker) Invoke() (*model.ListVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumeTagsResponse), nil
	}
}

type ListVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumesInvoker) Invoke() (*model.ListVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumesResponse), nil
	}
}

type ListVolumesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumesByTagsInvoker) Invoke() (*model.ListVolumesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumesByTagsResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowVolumeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVolumeTagsInvoker) Invoke() (*model.ShowVolumeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVolumeTagsResponse), nil
	}
}

type ListVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVersionsInvoker) Invoke() (*model.ListVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionsResponse), nil
	}
}

type ShowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionInvoker) Invoke() (*model.ShowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionResponse), nil
	}
}
