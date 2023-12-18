package v2

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/dcs/v2/model"
)

type DcsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDcsClient(hcClient *http_client.HcHttpClient) *DcsClient {
	return &DcsClient{HcClient: hcClient}
}

func DcsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *DcsClient) BatchCreateOrDeleteTags(request *model.BatchCreateOrDeleteTagsRequest) (*model.BatchCreateOrDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteTagsResponse), nil
	}
}

func (c *DcsClient) BatchCreateOrDeleteTagsInvoker(request *model.BatchCreateOrDeleteTagsRequest) *BatchCreateOrDeleteTagsInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteTags()
	return &BatchCreateOrDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) BatchDeleteInstances(request *model.BatchDeleteInstancesRequest) (*model.BatchDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstancesResponse), nil
	}
}

func (c *DcsClient) BatchDeleteInstancesInvoker(request *model.BatchDeleteInstancesRequest) *BatchDeleteInstancesInvoker {
	requestDef := GenReqDefForBatchDeleteInstances()
	return &BatchDeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) BatchShowNodesInformation(request *model.BatchShowNodesInformationRequest) (*model.BatchShowNodesInformationResponse, error) {
	requestDef := GenReqDefForBatchShowNodesInformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowNodesInformationResponse), nil
	}
}

func (c *DcsClient) BatchShowNodesInformationInvoker(request *model.BatchShowNodesInformationRequest) *BatchShowNodesInformationInvoker {
	requestDef := GenReqDefForBatchShowNodesInformation()
	return &BatchShowNodesInformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) BatchStopMigrationTasks(request *model.BatchStopMigrationTasksRequest) (*model.BatchStopMigrationTasksResponse, error) {
	requestDef := GenReqDefForBatchStopMigrationTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopMigrationTasksResponse), nil
	}
}

func (c *DcsClient) BatchStopMigrationTasksInvoker(request *model.BatchStopMigrationTasksRequest) *BatchStopMigrationTasksInvoker {
	requestDef := GenReqDefForBatchStopMigrationTasks()
	return &BatchStopMigrationTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ChangeMasterStandby(request *model.ChangeMasterStandbyRequest) (*model.ChangeMasterStandbyResponse, error) {
	requestDef := GenReqDefForChangeMasterStandby()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeMasterStandbyResponse), nil
	}
}

func (c *DcsClient) ChangeMasterStandbyInvoker(request *model.ChangeMasterStandbyRequest) *ChangeMasterStandbyInvoker {
	requestDef := GenReqDefForChangeMasterStandby()
	return &ChangeMasterStandbyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CopyInstance(request *model.CopyInstanceRequest) (*model.CopyInstanceResponse, error) {
	requestDef := GenReqDefForCopyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyInstanceResponse), nil
	}
}

func (c *DcsClient) CopyInstanceInvoker(request *model.CopyInstanceRequest) *CopyInstanceInvoker {
	requestDef := GenReqDefForCopyInstance()
	return &CopyInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateAutoExpireScanTask(request *model.CreateAutoExpireScanTaskRequest) (*model.CreateAutoExpireScanTaskResponse, error) {
	requestDef := GenReqDefForCreateAutoExpireScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAutoExpireScanTaskResponse), nil
	}
}

func (c *DcsClient) CreateAutoExpireScanTaskInvoker(request *model.CreateAutoExpireScanTaskRequest) *CreateAutoExpireScanTaskInvoker {
	requestDef := GenReqDefForCreateAutoExpireScanTask()
	return &CreateAutoExpireScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateBigkeyScanTask(request *model.CreateBigkeyScanTaskRequest) (*model.CreateBigkeyScanTaskResponse, error) {
	requestDef := GenReqDefForCreateBigkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBigkeyScanTaskResponse), nil
	}
}

func (c *DcsClient) CreateBigkeyScanTaskInvoker(request *model.CreateBigkeyScanTaskRequest) *CreateBigkeyScanTaskInvoker {
	requestDef := GenReqDefForCreateBigkeyScanTask()
	return &CreateBigkeyScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateCustomTemplate(request *model.CreateCustomTemplateRequest) (*model.CreateCustomTemplateResponse, error) {
	requestDef := GenReqDefForCreateCustomTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomTemplateResponse), nil
	}
}

func (c *DcsClient) CreateCustomTemplateInvoker(request *model.CreateCustomTemplateRequest) *CreateCustomTemplateInvoker {
	requestDef := GenReqDefForCreateCustomTemplate()
	return &CreateCustomTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateDiagnosisTask(request *model.CreateDiagnosisTaskRequest) (*model.CreateDiagnosisTaskResponse, error) {
	requestDef := GenReqDefForCreateDiagnosisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDiagnosisTaskResponse), nil
	}
}

func (c *DcsClient) CreateDiagnosisTaskInvoker(request *model.CreateDiagnosisTaskRequest) *CreateDiagnosisTaskInvoker {
	requestDef := GenReqDefForCreateDiagnosisTask()
	return &CreateDiagnosisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateHotkeyScanTask(request *model.CreateHotkeyScanTaskRequest) (*model.CreateHotkeyScanTaskResponse, error) {
	requestDef := GenReqDefForCreateHotkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHotkeyScanTaskResponse), nil
	}
}

func (c *DcsClient) CreateHotkeyScanTaskInvoker(request *model.CreateHotkeyScanTaskRequest) *CreateHotkeyScanTaskInvoker {
	requestDef := GenReqDefForCreateHotkeyScanTask()
	return &CreateHotkeyScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

func (c *DcsClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateMigrationTask(request *model.CreateMigrationTaskRequest) (*model.CreateMigrationTaskResponse, error) {
	requestDef := GenReqDefForCreateMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMigrationTaskResponse), nil
	}
}

func (c *DcsClient) CreateMigrationTaskInvoker(request *model.CreateMigrationTaskRequest) *CreateMigrationTaskInvoker {
	requestDef := GenReqDefForCreateMigrationTask()
	return &CreateMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateOnlineMigrationTask(request *model.CreateOnlineMigrationTaskRequest) (*model.CreateOnlineMigrationTaskResponse, error) {
	requestDef := GenReqDefForCreateOnlineMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOnlineMigrationTaskResponse), nil
	}
}

func (c *DcsClient) CreateOnlineMigrationTaskInvoker(request *model.CreateOnlineMigrationTaskRequest) *CreateOnlineMigrationTaskInvoker {
	requestDef := GenReqDefForCreateOnlineMigrationTask()
	return &CreateOnlineMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateRedislog(request *model.CreateRedislogRequest) (*model.CreateRedislogResponse, error) {
	requestDef := GenReqDefForCreateRedislog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRedislogResponse), nil
	}
}

func (c *DcsClient) CreateRedislogInvoker(request *model.CreateRedislogRequest) *CreateRedislogInvoker {
	requestDef := GenReqDefForCreateRedislog()
	return &CreateRedislogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) CreateRedislogDownloadLink(request *model.CreateRedislogDownloadLinkRequest) (*model.CreateRedislogDownloadLinkResponse, error) {
	requestDef := GenReqDefForCreateRedislogDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRedislogDownloadLinkResponse), nil
	}
}

func (c *DcsClient) CreateRedislogDownloadLinkInvoker(request *model.CreateRedislogDownloadLinkRequest) *CreateRedislogDownloadLinkInvoker {
	requestDef := GenReqDefForCreateRedislogDownloadLink()
	return &CreateRedislogDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) DeleteBackgroundTask(request *model.DeleteBackgroundTaskRequest) (*model.DeleteBackgroundTaskResponse, error) {
	requestDef := GenReqDefForDeleteBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackgroundTaskResponse), nil
	}
}

func (c *DcsClient) DeleteBackgroundTaskInvoker(request *model.DeleteBackgroundTaskRequest) *DeleteBackgroundTaskInvoker {
	requestDef := GenReqDefForDeleteBackgroundTask()
	return &DeleteBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) DeleteBackupFile(request *model.DeleteBackupFileRequest) (*model.DeleteBackupFileResponse, error) {
	requestDef := GenReqDefForDeleteBackupFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupFileResponse), nil
	}
}

func (c *DcsClient) DeleteBackupFileInvoker(request *model.DeleteBackupFileRequest) *DeleteBackupFileInvoker {
	requestDef := GenReqDefForDeleteBackupFile()
	return &DeleteBackupFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) DeleteBigkeyScanTask(request *model.DeleteBigkeyScanTaskRequest) (*model.DeleteBigkeyScanTaskResponse, error) {
	requestDef := GenReqDefForDeleteBigkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBigkeyScanTaskResponse), nil
	}
}

func (c *DcsClient) DeleteBigkeyScanTaskInvoker(request *model.DeleteBigkeyScanTaskRequest) *DeleteBigkeyScanTaskInvoker {
	requestDef := GenReqDefForDeleteBigkeyScanTask()
	return &DeleteBigkeyScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) DeleteHotkeyScanTask(request *model.DeleteHotkeyScanTaskRequest) (*model.DeleteHotkeyScanTaskResponse, error) {
	requestDef := GenReqDefForDeleteHotkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHotkeyScanTaskResponse), nil
	}
}

func (c *DcsClient) DeleteHotkeyScanTaskInvoker(request *model.DeleteHotkeyScanTaskRequest) *DeleteHotkeyScanTaskInvoker {
	requestDef := GenReqDefForDeleteHotkeyScanTask()
	return &DeleteHotkeyScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) DeleteIpFromDomainName(request *model.DeleteIpFromDomainNameRequest) (*model.DeleteIpFromDomainNameResponse, error) {
	requestDef := GenReqDefForDeleteIpFromDomainName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpFromDomainNameResponse), nil
	}
}

func (c *DcsClient) DeleteIpFromDomainNameInvoker(request *model.DeleteIpFromDomainNameRequest) *DeleteIpFromDomainNameInvoker {
	requestDef := GenReqDefForDeleteIpFromDomainName()
	return &DeleteIpFromDomainNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) DeleteMigrationTask(request *model.DeleteMigrationTaskRequest) (*model.DeleteMigrationTaskResponse, error) {
	requestDef := GenReqDefForDeleteMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMigrationTaskResponse), nil
	}
}

func (c *DcsClient) DeleteMigrationTaskInvoker(request *model.DeleteMigrationTaskRequest) *DeleteMigrationTaskInvoker {
	requestDef := GenReqDefForDeleteMigrationTask()
	return &DeleteMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) DeleteSingleInstance(request *model.DeleteSingleInstanceRequest) (*model.DeleteSingleInstanceResponse, error) {
	requestDef := GenReqDefForDeleteSingleInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSingleInstanceResponse), nil
	}
}

func (c *DcsClient) DeleteSingleInstanceInvoker(request *model.DeleteSingleInstanceRequest) *DeleteSingleInstanceInvoker {
	requestDef := GenReqDefForDeleteSingleInstance()
	return &DeleteSingleInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

func (c *DcsClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListBackgroundTask(request *model.ListBackgroundTaskRequest) (*model.ListBackgroundTaskResponse, error) {
	requestDef := GenReqDefForListBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackgroundTaskResponse), nil
	}
}

func (c *DcsClient) ListBackgroundTaskInvoker(request *model.ListBackgroundTaskRequest) *ListBackgroundTaskInvoker {
	requestDef := GenReqDefForListBackgroundTask()
	return &ListBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListBackupFileLinks(request *model.ListBackupFileLinksRequest) (*model.ListBackupFileLinksResponse, error) {
	requestDef := GenReqDefForListBackupFileLinks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupFileLinksResponse), nil
	}
}

func (c *DcsClient) ListBackupFileLinksInvoker(request *model.ListBackupFileLinksRequest) *ListBackupFileLinksInvoker {
	requestDef := GenReqDefForListBackupFileLinks()
	return &ListBackupFileLinksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListBackupRecords(request *model.ListBackupRecordsRequest) (*model.ListBackupRecordsResponse, error) {
	requestDef := GenReqDefForListBackupRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupRecordsResponse), nil
	}
}

func (c *DcsClient) ListBackupRecordsInvoker(request *model.ListBackupRecordsRequest) *ListBackupRecordsInvoker {
	requestDef := GenReqDefForListBackupRecords()
	return &ListBackupRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListBigkeyScanTasks(request *model.ListBigkeyScanTasksRequest) (*model.ListBigkeyScanTasksResponse, error) {
	requestDef := GenReqDefForListBigkeyScanTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBigkeyScanTasksResponse), nil
	}
}

func (c *DcsClient) ListBigkeyScanTasksInvoker(request *model.ListBigkeyScanTasksRequest) *ListBigkeyScanTasksInvoker {
	requestDef := GenReqDefForListBigkeyScanTasks()
	return &ListBigkeyScanTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListConfigHistories(request *model.ListConfigHistoriesRequest) (*model.ListConfigHistoriesResponse, error) {
	requestDef := GenReqDefForListConfigHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigHistoriesResponse), nil
	}
}

func (c *DcsClient) ListConfigHistoriesInvoker(request *model.ListConfigHistoriesRequest) *ListConfigHistoriesInvoker {
	requestDef := GenReqDefForListConfigHistories()
	return &ListConfigHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

func (c *DcsClient) ListConfigurationsInvoker(request *model.ListConfigurationsRequest) *ListConfigurationsInvoker {
	requestDef := GenReqDefForListConfigurations()
	return &ListConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListDiagnosisTasks(request *model.ListDiagnosisTasksRequest) (*model.ListDiagnosisTasksResponse, error) {
	requestDef := GenReqDefForListDiagnosisTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnosisTasksResponse), nil
	}
}

func (c *DcsClient) ListDiagnosisTasksInvoker(request *model.ListDiagnosisTasksRequest) *ListDiagnosisTasksInvoker {
	requestDef := GenReqDefForListDiagnosisTasks()
	return &ListDiagnosisTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

func (c *DcsClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListGroupReplicationInfo(request *model.ListGroupReplicationInfoRequest) (*model.ListGroupReplicationInfoResponse, error) {
	requestDef := GenReqDefForListGroupReplicationInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupReplicationInfoResponse), nil
	}
}

func (c *DcsClient) ListGroupReplicationInfoInvoker(request *model.ListGroupReplicationInfoRequest) *ListGroupReplicationInfoInvoker {
	requestDef := GenReqDefForListGroupReplicationInfo()
	return &ListGroupReplicationInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListHotKeyScanTasks(request *model.ListHotKeyScanTasksRequest) (*model.ListHotKeyScanTasksResponse, error) {
	requestDef := GenReqDefForListHotKeyScanTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHotKeyScanTasksResponse), nil
	}
}

func (c *DcsClient) ListHotKeyScanTasksInvoker(request *model.ListHotKeyScanTasksRequest) *ListHotKeyScanTasksInvoker {
	requestDef := GenReqDefForListHotKeyScanTasks()
	return &ListHotKeyScanTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

func (c *DcsClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListMaintenanceWindows(request *model.ListMaintenanceWindowsRequest) (*model.ListMaintenanceWindowsResponse, error) {
	requestDef := GenReqDefForListMaintenanceWindows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMaintenanceWindowsResponse), nil
	}
}

func (c *DcsClient) ListMaintenanceWindowsInvoker(request *model.ListMaintenanceWindowsRequest) *ListMaintenanceWindowsInvoker {
	requestDef := GenReqDefForListMaintenanceWindows()
	return &ListMaintenanceWindowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListMigrationTask(request *model.ListMigrationTaskRequest) (*model.ListMigrationTaskResponse, error) {
	requestDef := GenReqDefForListMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMigrationTaskResponse), nil
	}
}

func (c *DcsClient) ListMigrationTaskInvoker(request *model.ListMigrationTaskRequest) *ListMigrationTaskInvoker {
	requestDef := GenReqDefForListMigrationTask()
	return &ListMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListMonitoredObjects(request *model.ListMonitoredObjectsRequest) (*model.ListMonitoredObjectsResponse, error) {
	requestDef := GenReqDefForListMonitoredObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitoredObjectsResponse), nil
	}
}

func (c *DcsClient) ListMonitoredObjectsInvoker(request *model.ListMonitoredObjectsRequest) *ListMonitoredObjectsInvoker {
	requestDef := GenReqDefForListMonitoredObjects()
	return &ListMonitoredObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListMonitoredObjectsOfInstance(request *model.ListMonitoredObjectsOfInstanceRequest) (*model.ListMonitoredObjectsOfInstanceResponse, error) {
	requestDef := GenReqDefForListMonitoredObjectsOfInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitoredObjectsOfInstanceResponse), nil
	}
}

func (c *DcsClient) ListMonitoredObjectsOfInstanceInvoker(request *model.ListMonitoredObjectsOfInstanceRequest) *ListMonitoredObjectsOfInstanceInvoker {
	requestDef := GenReqDefForListMonitoredObjectsOfInstance()
	return &ListMonitoredObjectsOfInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListNumberOfInstancesInDifferentStatus(request *model.ListNumberOfInstancesInDifferentStatusRequest) (*model.ListNumberOfInstancesInDifferentStatusResponse, error) {
	requestDef := GenReqDefForListNumberOfInstancesInDifferentStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNumberOfInstancesInDifferentStatusResponse), nil
	}
}

func (c *DcsClient) ListNumberOfInstancesInDifferentStatusInvoker(request *model.ListNumberOfInstancesInDifferentStatusRequest) *ListNumberOfInstancesInDifferentStatusInvoker {
	requestDef := GenReqDefForListNumberOfInstancesInDifferentStatus()
	return &ListNumberOfInstancesInDifferentStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListRedislog(request *model.ListRedislogRequest) (*model.ListRedislogResponse, error) {
	requestDef := GenReqDefForListRedislog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRedislogResponse), nil
	}
}

func (c *DcsClient) ListRedislogInvoker(request *model.ListRedislogRequest) *ListRedislogInvoker {
	requestDef := GenReqDefForListRedislog()
	return &ListRedislogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListRestoreRecords(request *model.ListRestoreRecordsRequest) (*model.ListRestoreRecordsResponse, error) {
	requestDef := GenReqDefForListRestoreRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreRecordsResponse), nil
	}
}

func (c *DcsClient) ListRestoreRecordsInvoker(request *model.ListRestoreRecordsRequest) *ListRestoreRecordsInvoker {
	requestDef := GenReqDefForListRestoreRecords()
	return &ListRestoreRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListSlowlog(request *model.ListSlowlogRequest) (*model.ListSlowlogResponse, error) {
	requestDef := GenReqDefForListSlowlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowlogResponse), nil
	}
}

func (c *DcsClient) ListSlowlogInvoker(request *model.ListSlowlogRequest) *ListSlowlogInvoker {
	requestDef := GenReqDefForListSlowlog()
	return &ListSlowlogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListStatisticsOfRunningInstances(request *model.ListStatisticsOfRunningInstancesRequest) (*model.ListStatisticsOfRunningInstancesResponse, error) {
	requestDef := GenReqDefForListStatisticsOfRunningInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsOfRunningInstancesResponse), nil
	}
}

func (c *DcsClient) ListStatisticsOfRunningInstancesInvoker(request *model.ListStatisticsOfRunningInstancesRequest) *ListStatisticsOfRunningInstancesInvoker {
	requestDef := GenReqDefForListStatisticsOfRunningInstances()
	return &ListStatisticsOfRunningInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ListTagsOfTenant(request *model.ListTagsOfTenantRequest) (*model.ListTagsOfTenantResponse, error) {
	requestDef := GenReqDefForListTagsOfTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsOfTenantResponse), nil
	}
}

func (c *DcsClient) ListTagsOfTenantInvoker(request *model.ListTagsOfTenantRequest) *ListTagsOfTenantInvoker {
	requestDef := GenReqDefForListTagsOfTenant()
	return &ListTagsOfTenantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

func (c *DcsClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) RestartOrFlushInstances(request *model.RestartOrFlushInstancesRequest) (*model.RestartOrFlushInstancesResponse, error) {
	requestDef := GenReqDefForRestartOrFlushInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartOrFlushInstancesResponse), nil
	}
}

func (c *DcsClient) RestartOrFlushInstancesInvoker(request *model.RestartOrFlushInstancesRequest) *RestartOrFlushInstancesInvoker {
	requestDef := GenReqDefForRestartOrFlushInstances()
	return &RestartOrFlushInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) RestoreInstance(request *model.RestoreInstanceRequest) (*model.RestoreInstanceResponse, error) {
	requestDef := GenReqDefForRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceResponse), nil
	}
}

func (c *DcsClient) RestoreInstanceInvoker(request *model.RestoreInstanceRequest) *RestoreInstanceInvoker {
	requestDef := GenReqDefForRestoreInstance()
	return &RestoreInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) SetOnlineMigrationTask(request *model.SetOnlineMigrationTaskRequest) (*model.SetOnlineMigrationTaskResponse, error) {
	requestDef := GenReqDefForSetOnlineMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetOnlineMigrationTaskResponse), nil
	}
}

func (c *DcsClient) SetOnlineMigrationTaskInvoker(request *model.SetOnlineMigrationTaskRequest) *SetOnlineMigrationTaskInvoker {
	requestDef := GenReqDefForSetOnlineMigrationTask()
	return &SetOnlineMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowBigkeyAutoscanConfig(request *model.ShowBigkeyAutoscanConfigRequest) (*model.ShowBigkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForShowBigkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBigkeyAutoscanConfigResponse), nil
	}
}

func (c *DcsClient) ShowBigkeyAutoscanConfigInvoker(request *model.ShowBigkeyAutoscanConfigRequest) *ShowBigkeyAutoscanConfigInvoker {
	requestDef := GenReqDefForShowBigkeyAutoscanConfig()
	return &ShowBigkeyAutoscanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowBigkeyScanTaskDetails(request *model.ShowBigkeyScanTaskDetailsRequest) (*model.ShowBigkeyScanTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowBigkeyScanTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBigkeyScanTaskDetailsResponse), nil
	}
}

func (c *DcsClient) ShowBigkeyScanTaskDetailsInvoker(request *model.ShowBigkeyScanTaskDetailsRequest) *ShowBigkeyScanTaskDetailsInvoker {
	requestDef := GenReqDefForShowBigkeyScanTaskDetails()
	return &ShowBigkeyScanTaskDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowDiagnosisTaskDetails(request *model.ShowDiagnosisTaskDetailsRequest) (*model.ShowDiagnosisTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowDiagnosisTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisTaskDetailsResponse), nil
	}
}

func (c *DcsClient) ShowDiagnosisTaskDetailsInvoker(request *model.ShowDiagnosisTaskDetailsRequest) *ShowDiagnosisTaskDetailsInvoker {
	requestDef := GenReqDefForShowDiagnosisTaskDetails()
	return &ShowDiagnosisTaskDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowHotkeyAutoscanConfig(request *model.ShowHotkeyAutoscanConfigRequest) (*model.ShowHotkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForShowHotkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotkeyAutoscanConfigResponse), nil
	}
}

func (c *DcsClient) ShowHotkeyAutoscanConfigInvoker(request *model.ShowHotkeyAutoscanConfigRequest) *ShowHotkeyAutoscanConfigInvoker {
	requestDef := GenReqDefForShowHotkeyAutoscanConfig()
	return &ShowHotkeyAutoscanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowHotkeyTaskDetails(request *model.ShowHotkeyTaskDetailsRequest) (*model.ShowHotkeyTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowHotkeyTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotkeyTaskDetailsResponse), nil
	}
}

func (c *DcsClient) ShowHotkeyTaskDetailsInvoker(request *model.ShowHotkeyTaskDetailsRequest) *ShowHotkeyTaskDetailsInvoker {
	requestDef := GenReqDefForShowHotkeyTaskDetails()
	return &ShowHotkeyTaskDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

func (c *DcsClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowMigrationTask(request *model.ShowMigrationTaskRequest) (*model.ShowMigrationTaskResponse, error) {
	requestDef := GenReqDefForShowMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationTaskResponse), nil
	}
}

func (c *DcsClient) ShowMigrationTaskInvoker(request *model.ShowMigrationTaskRequest) *ShowMigrationTaskInvoker {
	requestDef := GenReqDefForShowMigrationTask()
	return &ShowMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowMigrationTaskStats(request *model.ShowMigrationTaskStatsRequest) (*model.ShowMigrationTaskStatsResponse, error) {
	requestDef := GenReqDefForShowMigrationTaskStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationTaskStatsResponse), nil
	}
}

func (c *DcsClient) ShowMigrationTaskStatsInvoker(request *model.ShowMigrationTaskStatsRequest) *ShowMigrationTaskStatsInvoker {
	requestDef := GenReqDefForShowMigrationTaskStats()
	return &ShowMigrationTaskStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowQuotaOfTenant(request *model.ShowQuotaOfTenantRequest) (*model.ShowQuotaOfTenantResponse, error) {
	requestDef := GenReqDefForShowQuotaOfTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaOfTenantResponse), nil
	}
}

func (c *DcsClient) ShowQuotaOfTenantInvoker(request *model.ShowQuotaOfTenantRequest) *ShowQuotaOfTenantInvoker {
	requestDef := GenReqDefForShowQuotaOfTenant()
	return &ShowQuotaOfTenantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowTags(request *model.ShowTagsRequest) (*model.ShowTagsResponse, error) {
	requestDef := GenReqDefForShowTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagsResponse), nil
	}
}

func (c *DcsClient) ShowTagsInvoker(request *model.ShowTagsRequest) *ShowTagsInvoker {
	requestDef := GenReqDefForShowTags()
	return &ShowTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) StopMigrationTask(request *model.StopMigrationTaskRequest) (*model.StopMigrationTaskResponse, error) {
	requestDef := GenReqDefForStopMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMigrationTaskResponse), nil
	}
}

func (c *DcsClient) StopMigrationTaskInvoker(request *model.StopMigrationTaskRequest) *StopMigrationTaskInvoker {
	requestDef := GenReqDefForStopMigrationTask()
	return &StopMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) StopMigrationTaskSync(request *model.StopMigrationTaskSyncRequest) (*model.StopMigrationTaskSyncResponse, error) {
	requestDef := GenReqDefForStopMigrationTaskSync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMigrationTaskSyncResponse), nil
	}
}

func (c *DcsClient) StopMigrationTaskSyncInvoker(request *model.StopMigrationTaskSyncRequest) *StopMigrationTaskSyncInvoker {
	requestDef := GenReqDefForStopMigrationTaskSync()
	return &StopMigrationTaskSyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) UpdateBigkeyAutoscanConfig(request *model.UpdateBigkeyAutoscanConfigRequest) (*model.UpdateBigkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForUpdateBigkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBigkeyAutoscanConfigResponse), nil
	}
}

func (c *DcsClient) UpdateBigkeyAutoscanConfigInvoker(request *model.UpdateBigkeyAutoscanConfigRequest) *UpdateBigkeyAutoscanConfigInvoker {
	requestDef := GenReqDefForUpdateBigkeyAutoscanConfig()
	return &UpdateBigkeyAutoscanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) UpdateConfigurations(request *model.UpdateConfigurationsRequest) (*model.UpdateConfigurationsResponse, error) {
	requestDef := GenReqDefForUpdateConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationsResponse), nil
	}
}

func (c *DcsClient) UpdateConfigurationsInvoker(request *model.UpdateConfigurationsRequest) *UpdateConfigurationsInvoker {
	requestDef := GenReqDefForUpdateConfigurations()
	return &UpdateConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) UpdateHotkeyAutoScanConfig(request *model.UpdateHotkeyAutoScanConfigRequest) (*model.UpdateHotkeyAutoScanConfigResponse, error) {
	requestDef := GenReqDefForUpdateHotkeyAutoScanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHotkeyAutoScanConfigResponse), nil
	}
}

func (c *DcsClient) UpdateHotkeyAutoScanConfigInvoker(request *model.UpdateHotkeyAutoScanConfigRequest) *UpdateHotkeyAutoScanConfigInvoker {
	requestDef := GenReqDefForUpdateHotkeyAutoScanConfig()
	return &UpdateHotkeyAutoScanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

func (c *DcsClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) UpdatePassword(request *model.UpdatePasswordRequest) (*model.UpdatePasswordResponse, error) {
	requestDef := GenReqDefForUpdatePassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePasswordResponse), nil
	}
}

func (c *DcsClient) UpdatePasswordInvoker(request *model.UpdatePasswordRequest) *UpdatePasswordInvoker {
	requestDef := GenReqDefForUpdatePassword()
	return &UpdatePasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) UpdateSlavePriority(request *model.UpdateSlavePriorityRequest) (*model.UpdateSlavePriorityResponse, error) {
	requestDef := GenReqDefForUpdateSlavePriority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSlavePriorityResponse), nil
	}
}

func (c *DcsClient) UpdateSlavePriorityInvoker(request *model.UpdateSlavePriorityRequest) *UpdateSlavePriorityInvoker {
	requestDef := GenReqDefForUpdateSlavePriority()
	return &UpdateSlavePriorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) ShowIpWhitelist(request *model.ShowIpWhitelistRequest) (*model.ShowIpWhitelistResponse, error) {
	requestDef := GenReqDefForShowIpWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpWhitelistResponse), nil
	}
}

func (c *DcsClient) ShowIpWhitelistInvoker(request *model.ShowIpWhitelistRequest) *ShowIpWhitelistInvoker {
	requestDef := GenReqDefForShowIpWhitelist()
	return &ShowIpWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcsClient) UpdateIpWhitelist(request *model.UpdateIpWhitelistRequest) (*model.UpdateIpWhitelistResponse, error) {
	requestDef := GenReqDefForUpdateIpWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpWhitelistResponse), nil
	}
}

func (c *DcsClient) UpdateIpWhitelistInvoker(request *model.UpdateIpWhitelistRequest) *UpdateIpWhitelistInvoker {
	requestDef := GenReqDefForUpdateIpWhitelist()
	return &UpdateIpWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
