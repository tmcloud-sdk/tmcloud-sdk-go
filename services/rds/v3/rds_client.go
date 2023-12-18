package v3

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/rds/v3/model"
)

type RdsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewRdsClient(hcClient *http_client.HcHttpClient) *RdsClient {
	return &RdsClient{HcClient: hcClient}
}

func RdsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *RdsClient) ApplyConfigurationAsync(request *model.ApplyConfigurationAsyncRequest) (*model.ApplyConfigurationAsyncResponse, error) {
	requestDef := GenReqDefForApplyConfigurationAsync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyConfigurationAsyncResponse), nil
	}
}

func (c *RdsClient) ApplyConfigurationAsyncInvoker(request *model.ApplyConfigurationAsyncRequest) *ApplyConfigurationAsyncInvoker {
	requestDef := GenReqDefForApplyConfigurationAsync()
	return &ApplyConfigurationAsyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) AttachEip(request *model.AttachEipRequest) (*model.AttachEipResponse, error) {
	requestDef := GenReqDefForAttachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachEipResponse), nil
	}
}

func (c *RdsClient) AttachEipInvoker(request *model.AttachEipRequest) *AttachEipInvoker {
	requestDef := GenReqDefForAttachEip()
	return &AttachEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) BatchTagAddAction(request *model.BatchTagAddActionRequest) (*model.BatchTagAddActionResponse, error) {
	requestDef := GenReqDefForBatchTagAddAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagAddActionResponse), nil
	}
}

func (c *RdsClient) BatchTagAddActionInvoker(request *model.BatchTagAddActionRequest) *BatchTagAddActionInvoker {
	requestDef := GenReqDefForBatchTagAddAction()
	return &BatchTagAddActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) BatchTagDelAction(request *model.BatchTagDelActionRequest) (*model.BatchTagDelActionResponse, error) {
	requestDef := GenReqDefForBatchTagDelAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagDelActionResponse), nil
	}
}

func (c *RdsClient) BatchTagDelActionInvoker(request *model.BatchTagDelActionRequest) *BatchTagDelActionInvoker {
	requestDef := GenReqDefForBatchTagDelAction()
	return &BatchTagDelActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ChangeFailoverMode(request *model.ChangeFailoverModeRequest) (*model.ChangeFailoverModeResponse, error) {
	requestDef := GenReqDefForChangeFailoverMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeFailoverModeResponse), nil
	}
}

func (c *RdsClient) ChangeFailoverModeInvoker(request *model.ChangeFailoverModeRequest) *ChangeFailoverModeInvoker {
	requestDef := GenReqDefForChangeFailoverMode()
	return &ChangeFailoverModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ChangeFailoverStrategy(request *model.ChangeFailoverStrategyRequest) (*model.ChangeFailoverStrategyResponse, error) {
	requestDef := GenReqDefForChangeFailoverStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeFailoverStrategyResponse), nil
	}
}

func (c *RdsClient) ChangeFailoverStrategyInvoker(request *model.ChangeFailoverStrategyRequest) *ChangeFailoverStrategyInvoker {
	requestDef := GenReqDefForChangeFailoverStrategy()
	return &ChangeFailoverStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ChangeOpsWindow(request *model.ChangeOpsWindowRequest) (*model.ChangeOpsWindowResponse, error) {
	requestDef := GenReqDefForChangeOpsWindow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeOpsWindowResponse), nil
	}
}

func (c *RdsClient) ChangeOpsWindowInvoker(request *model.ChangeOpsWindowRequest) *ChangeOpsWindowInvoker {
	requestDef := GenReqDefForChangeOpsWindow()
	return &ChangeOpsWindowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CopyConfiguration(request *model.CopyConfigurationRequest) (*model.CopyConfigurationResponse, error) {
	requestDef := GenReqDefForCopyConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyConfigurationResponse), nil
	}
}

func (c *RdsClient) CopyConfigurationInvoker(request *model.CopyConfigurationRequest) *CopyConfigurationInvoker {
	requestDef := GenReqDefForCopyConfiguration()
	return &CopyConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateConfiguration(request *model.CreateConfigurationRequest) (*model.CreateConfigurationResponse, error) {
	requestDef := GenReqDefForCreateConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationResponse), nil
	}
}

func (c *RdsClient) CreateConfigurationInvoker(request *model.CreateConfigurationRequest) *CreateConfigurationInvoker {
	requestDef := GenReqDefForCreateConfiguration()
	return &CreateConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateDnsName(request *model.CreateDnsNameRequest) (*model.CreateDnsNameResponse, error) {
	requestDef := GenReqDefForCreateDnsName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDnsNameResponse), nil
	}
}

func (c *RdsClient) CreateDnsNameInvoker(request *model.CreateDnsNameRequest) *CreateDnsNameInvoker {
	requestDef := GenReqDefForCreateDnsName()
	return &CreateDnsNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

func (c *RdsClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateManualBackup(request *model.CreateManualBackupRequest) (*model.CreateManualBackupResponse, error) {
	requestDef := GenReqDefForCreateManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManualBackupResponse), nil
	}
}

func (c *RdsClient) CreateManualBackupInvoker(request *model.CreateManualBackupRequest) *CreateManualBackupInvoker {
	requestDef := GenReqDefForCreateManualBackup()
	return &CreateManualBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateRestoreInstance(request *model.CreateRestoreInstanceRequest) (*model.CreateRestoreInstanceResponse, error) {
	requestDef := GenReqDefForCreateRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRestoreInstanceResponse), nil
	}
}

func (c *RdsClient) CreateRestoreInstanceInvoker(request *model.CreateRestoreInstanceRequest) *CreateRestoreInstanceInvoker {
	requestDef := GenReqDefForCreateRestoreInstance()
	return &CreateRestoreInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeleteConfiguration(request *model.DeleteConfigurationRequest) (*model.DeleteConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationResponse), nil
	}
}

func (c *RdsClient) DeleteConfigurationInvoker(request *model.DeleteConfigurationRequest) *DeleteConfigurationInvoker {
	requestDef := GenReqDefForDeleteConfiguration()
	return &DeleteConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

func (c *RdsClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeleteManualBackup(request *model.DeleteManualBackupRequest) (*model.DeleteManualBackupResponse, error) {
	requestDef := GenReqDefForDeleteManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteManualBackupResponse), nil
	}
}

func (c *RdsClient) DeleteManualBackupInvoker(request *model.DeleteManualBackupRequest) *DeleteManualBackupInvoker {
	requestDef := GenReqDefForDeleteManualBackup()
	return &DeleteManualBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DownloadSlowlog(request *model.DownloadSlowlogRequest) (*model.DownloadSlowlogResponse, error) {
	requestDef := GenReqDefForDownloadSlowlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadSlowlogResponse), nil
	}
}

func (c *RdsClient) DownloadSlowlogInvoker(request *model.DownloadSlowlogRequest) *DownloadSlowlogInvoker {
	requestDef := GenReqDefForDownloadSlowlog()
	return &DownloadSlowlogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) EnableConfiguration(request *model.EnableConfigurationRequest) (*model.EnableConfigurationResponse, error) {
	requestDef := GenReqDefForEnableConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableConfigurationResponse), nil
	}
}

func (c *RdsClient) EnableConfigurationInvoker(request *model.EnableConfigurationRequest) *EnableConfigurationInvoker {
	requestDef := GenReqDefForEnableConfiguration()
	return &EnableConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListAuditlogs(request *model.ListAuditlogsRequest) (*model.ListAuditlogsResponse, error) {
	requestDef := GenReqDefForListAuditlogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditlogsResponse), nil
	}
}

func (c *RdsClient) ListAuditlogsInvoker(request *model.ListAuditlogsRequest) *ListAuditlogsInvoker {
	requestDef := GenReqDefForListAuditlogs()
	return &ListAuditlogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

func (c *RdsClient) ListBackupsInvoker(request *model.ListBackupsRequest) *ListBackupsInvoker {
	requestDef := GenReqDefForListBackups()
	return &ListBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListCollations(request *model.ListCollationsRequest) (*model.ListCollationsResponse, error) {
	requestDef := GenReqDefForListCollations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollationsResponse), nil
	}
}

func (c *RdsClient) ListCollationsInvoker(request *model.ListCollationsRequest) *ListCollationsInvoker {
	requestDef := GenReqDefForListCollations()
	return &ListCollationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

func (c *RdsClient) ListConfigurationsInvoker(request *model.ListConfigurationsRequest) *ListConfigurationsInvoker {
	requestDef := GenReqDefForListConfigurations()
	return &ListConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListDatastores(request *model.ListDatastoresRequest) (*model.ListDatastoresResponse, error) {
	requestDef := GenReqDefForListDatastores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatastoresResponse), nil
	}
}

func (c *RdsClient) ListDatastoresInvoker(request *model.ListDatastoresRequest) *ListDatastoresInvoker {
	requestDef := GenReqDefForListDatastores()
	return &ListDatastoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListDrRelations(request *model.ListDrRelationsRequest) (*model.ListDrRelationsResponse, error) {
	requestDef := GenReqDefForListDrRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDrRelationsResponse), nil
	}
}

func (c *RdsClient) ListDrRelationsInvoker(request *model.ListDrRelationsRequest) *ListDrRelationsInvoker {
	requestDef := GenReqDefForListDrRelations()
	return &ListDrRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListErrorLogs(request *model.ListErrorLogsRequest) (*model.ListErrorLogsResponse, error) {
	requestDef := GenReqDefForListErrorLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListErrorLogsResponse), nil
	}
}

func (c *RdsClient) ListErrorLogsInvoker(request *model.ListErrorLogsRequest) *ListErrorLogsInvoker {
	requestDef := GenReqDefForListErrorLogs()
	return &ListErrorLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListErrorLogsNew(request *model.ListErrorLogsNewRequest) (*model.ListErrorLogsNewResponse, error) {
	requestDef := GenReqDefForListErrorLogsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListErrorLogsNewResponse), nil
	}
}

func (c *RdsClient) ListErrorLogsNewInvoker(request *model.ListErrorLogsNewRequest) *ListErrorLogsNewInvoker {
	requestDef := GenReqDefForListErrorLogsNew()
	return &ListErrorLogsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListErrorlogForLts(request *model.ListErrorlogForLtsRequest) (*model.ListErrorlogForLtsResponse, error) {
	requestDef := GenReqDefForListErrorlogForLts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListErrorlogForLtsResponse), nil
	}
}

func (c *RdsClient) ListErrorlogForLtsInvoker(request *model.ListErrorlogForLtsRequest) *ListErrorlogForLtsInvoker {
	requestDef := GenReqDefForListErrorlogForLts()
	return &ListErrorlogForLtsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

func (c *RdsClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListInstanceParamHistories(request *model.ListInstanceParamHistoriesRequest) (*model.ListInstanceParamHistoriesResponse, error) {
	requestDef := GenReqDefForListInstanceParamHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceParamHistoriesResponse), nil
	}
}

func (c *RdsClient) ListInstanceParamHistoriesInvoker(request *model.ListInstanceParamHistoriesRequest) *ListInstanceParamHistoriesInvoker {
	requestDef := GenReqDefForListInstanceParamHistories()
	return &ListInstanceParamHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

func (c *RdsClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListJobInfo(request *model.ListJobInfoRequest) (*model.ListJobInfoResponse, error) {
	requestDef := GenReqDefForListJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobInfoResponse), nil
	}
}

func (c *RdsClient) ListJobInfoInvoker(request *model.ListJobInfoRequest) *ListJobInfoInvoker {
	requestDef := GenReqDefForListJobInfo()
	return &ListJobInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListJobInfoDetail(request *model.ListJobInfoDetailRequest) (*model.ListJobInfoDetailResponse, error) {
	requestDef := GenReqDefForListJobInfoDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobInfoDetailResponse), nil
	}
}

func (c *RdsClient) ListJobInfoDetailInvoker(request *model.ListJobInfoDetailRequest) *ListJobInfoDetailInvoker {
	requestDef := GenReqDefForListJobInfoDetail()
	return &ListJobInfoDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListOffSiteBackups(request *model.ListOffSiteBackupsRequest) (*model.ListOffSiteBackupsResponse, error) {
	requestDef := GenReqDefForListOffSiteBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOffSiteBackupsResponse), nil
	}
}

func (c *RdsClient) ListOffSiteBackupsInvoker(request *model.ListOffSiteBackupsRequest) *ListOffSiteBackupsInvoker {
	requestDef := GenReqDefForListOffSiteBackups()
	return &ListOffSiteBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListOffSiteInstances(request *model.ListOffSiteInstancesRequest) (*model.ListOffSiteInstancesResponse, error) {
	requestDef := GenReqDefForListOffSiteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOffSiteInstancesResponse), nil
	}
}

func (c *RdsClient) ListOffSiteInstancesInvoker(request *model.ListOffSiteInstancesRequest) *ListOffSiteInstancesInvoker {
	requestDef := GenReqDefForListOffSiteInstances()
	return &ListOffSiteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListOffSiteRestoreTimes(request *model.ListOffSiteRestoreTimesRequest) (*model.ListOffSiteRestoreTimesResponse, error) {
	requestDef := GenReqDefForListOffSiteRestoreTimes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOffSiteRestoreTimesResponse), nil
	}
}

func (c *RdsClient) ListOffSiteRestoreTimesInvoker(request *model.ListOffSiteRestoreTimesRequest) *ListOffSiteRestoreTimesInvoker {
	requestDef := GenReqDefForListOffSiteRestoreTimes()
	return &ListOffSiteRestoreTimesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

func (c *RdsClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListRecycleInstances(request *model.ListRecycleInstancesRequest) (*model.ListRecycleInstancesResponse, error) {
	requestDef := GenReqDefForListRecycleInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecycleInstancesResponse), nil
	}
}

func (c *RdsClient) ListRecycleInstancesInvoker(request *model.ListRecycleInstancesRequest) *ListRecycleInstancesInvoker {
	requestDef := GenReqDefForListRecycleInstances()
	return &ListRecycleInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListRestoreTimes(request *model.ListRestoreTimesRequest) (*model.ListRestoreTimesResponse, error) {
	requestDef := GenReqDefForListRestoreTimes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreTimesResponse), nil
	}
}

func (c *RdsClient) ListRestoreTimesInvoker(request *model.ListRestoreTimesRequest) *ListRestoreTimesInvoker {
	requestDef := GenReqDefForListRestoreTimes()
	return &ListRestoreTimesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListSlowLogFile(request *model.ListSlowLogFileRequest) (*model.ListSlowLogFileResponse, error) {
	requestDef := GenReqDefForListSlowLogFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogFileResponse), nil
	}
}

func (c *RdsClient) ListSlowLogFileInvoker(request *model.ListSlowLogFileRequest) *ListSlowLogFileInvoker {
	requestDef := GenReqDefForListSlowLogFile()
	return &ListSlowLogFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListSlowLogStatisticsForLts(request *model.ListSlowLogStatisticsForLtsRequest) (*model.ListSlowLogStatisticsForLtsResponse, error) {
	requestDef := GenReqDefForListSlowLogStatisticsForLts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogStatisticsForLtsResponse), nil
	}
}

func (c *RdsClient) ListSlowLogStatisticsForLtsInvoker(request *model.ListSlowLogStatisticsForLtsRequest) *ListSlowLogStatisticsForLtsInvoker {
	requestDef := GenReqDefForListSlowLogStatisticsForLts()
	return &ListSlowLogStatisticsForLtsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListSlowLogs(request *model.ListSlowLogsRequest) (*model.ListSlowLogsResponse, error) {
	requestDef := GenReqDefForListSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogsResponse), nil
	}
}

func (c *RdsClient) ListSlowLogsInvoker(request *model.ListSlowLogsRequest) *ListSlowLogsInvoker {
	requestDef := GenReqDefForListSlowLogs()
	return &ListSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListSlowLogsNew(request *model.ListSlowLogsNewRequest) (*model.ListSlowLogsNewResponse, error) {
	requestDef := GenReqDefForListSlowLogsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogsNewResponse), nil
	}
}

func (c *RdsClient) ListSlowLogsNewInvoker(request *model.ListSlowLogsNewRequest) *ListSlowLogsNewInvoker {
	requestDef := GenReqDefForListSlowLogsNew()
	return &ListSlowLogsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListSlowlogForLts(request *model.ListSlowlogForLtsRequest) (*model.ListSlowlogForLtsResponse, error) {
	requestDef := GenReqDefForListSlowlogForLts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowlogForLtsResponse), nil
	}
}

func (c *RdsClient) ListSlowlogForLtsInvoker(request *model.ListSlowlogForLtsRequest) *ListSlowlogForLtsInvoker {
	requestDef := GenReqDefForListSlowlogForLts()
	return &ListSlowlogForLtsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListSlowlogStatistics(request *model.ListSlowlogStatisticsRequest) (*model.ListSlowlogStatisticsResponse, error) {
	requestDef := GenReqDefForListSlowlogStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowlogStatisticsResponse), nil
	}
}

func (c *RdsClient) ListSlowlogStatisticsInvoker(request *model.ListSlowlogStatisticsRequest) *ListSlowlogStatisticsInvoker {
	requestDef := GenReqDefForListSlowlogStatistics()
	return &ListSlowlogStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListStorageTypes(request *model.ListStorageTypesRequest) (*model.ListStorageTypesResponse, error) {
	requestDef := GenReqDefForListStorageTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageTypesResponse), nil
	}
}

func (c *RdsClient) ListStorageTypesInvoker(request *model.ListStorageTypesRequest) *ListStorageTypesInvoker {
	requestDef := GenReqDefForListStorageTypes()
	return &ListStorageTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) MigrateFollower(request *model.MigrateFollowerRequest) (*model.MigrateFollowerResponse, error) {
	requestDef := GenReqDefForMigrateFollower()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateFollowerResponse), nil
	}
}

func (c *RdsClient) MigrateFollowerInvoker(request *model.MigrateFollowerRequest) *MigrateFollowerInvoker {
	requestDef := GenReqDefForMigrateFollower()
	return &MigrateFollowerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) RestoreExistInstance(request *model.RestoreExistInstanceRequest) (*model.RestoreExistInstanceResponse, error) {
	requestDef := GenReqDefForRestoreExistInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreExistInstanceResponse), nil
	}
}

func (c *RdsClient) RestoreExistInstanceInvoker(request *model.RestoreExistInstanceRequest) *RestoreExistInstanceInvoker {
	requestDef := GenReqDefForRestoreExistInstance()
	return &RestoreExistInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) RestoreTables(request *model.RestoreTablesRequest) (*model.RestoreTablesResponse, error) {
	requestDef := GenReqDefForRestoreTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreTablesResponse), nil
	}
}

func (c *RdsClient) RestoreTablesInvoker(request *model.RestoreTablesRequest) *RestoreTablesInvoker {
	requestDef := GenReqDefForRestoreTables()
	return &RestoreTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) RestoreToExistingInstance(request *model.RestoreToExistingInstanceRequest) (*model.RestoreToExistingInstanceResponse, error) {
	requestDef := GenReqDefForRestoreToExistingInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreToExistingInstanceResponse), nil
	}
}

func (c *RdsClient) RestoreToExistingInstanceInvoker(request *model.RestoreToExistingInstanceRequest) *RestoreToExistingInstanceInvoker {
	requestDef := GenReqDefForRestoreToExistingInstance()
	return &RestoreToExistingInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetAuditlogPolicy(request *model.SetAuditlogPolicyRequest) (*model.SetAuditlogPolicyResponse, error) {
	requestDef := GenReqDefForSetAuditlogPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditlogPolicyResponse), nil
	}
}

func (c *RdsClient) SetAuditlogPolicyInvoker(request *model.SetAuditlogPolicyRequest) *SetAuditlogPolicyInvoker {
	requestDef := GenReqDefForSetAuditlogPolicy()
	return &SetAuditlogPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetAutoEnlargePolicy(request *model.SetAutoEnlargePolicyRequest) (*model.SetAutoEnlargePolicyResponse, error) {
	requestDef := GenReqDefForSetAutoEnlargePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAutoEnlargePolicyResponse), nil
	}
}

func (c *RdsClient) SetAutoEnlargePolicyInvoker(request *model.SetAutoEnlargePolicyRequest) *SetAutoEnlargePolicyInvoker {
	requestDef := GenReqDefForSetAutoEnlargePolicy()
	return &SetAutoEnlargePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetBackupPolicy(request *model.SetBackupPolicyRequest) (*model.SetBackupPolicyResponse, error) {
	requestDef := GenReqDefForSetBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBackupPolicyResponse), nil
	}
}

func (c *RdsClient) SetBackupPolicyInvoker(request *model.SetBackupPolicyRequest) *SetBackupPolicyInvoker {
	requestDef := GenReqDefForSetBackupPolicy()
	return &SetBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetBinlogClearPolicy(request *model.SetBinlogClearPolicyRequest) (*model.SetBinlogClearPolicyResponse, error) {
	requestDef := GenReqDefForSetBinlogClearPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBinlogClearPolicyResponse), nil
	}
}

func (c *RdsClient) SetBinlogClearPolicyInvoker(request *model.SetBinlogClearPolicyRequest) *SetBinlogClearPolicyInvoker {
	requestDef := GenReqDefForSetBinlogClearPolicy()
	return &SetBinlogClearPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetOffSiteBackupPolicy(request *model.SetOffSiteBackupPolicyRequest) (*model.SetOffSiteBackupPolicyResponse, error) {
	requestDef := GenReqDefForSetOffSiteBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetOffSiteBackupPolicyResponse), nil
	}
}

func (c *RdsClient) SetOffSiteBackupPolicyInvoker(request *model.SetOffSiteBackupPolicyRequest) *SetOffSiteBackupPolicyInvoker {
	requestDef := GenReqDefForSetOffSiteBackupPolicy()
	return &SetOffSiteBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetSecondLevelMonitor(request *model.SetSecondLevelMonitorRequest) (*model.SetSecondLevelMonitorResponse, error) {
	requestDef := GenReqDefForSetSecondLevelMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSecondLevelMonitorResponse), nil
	}
}

func (c *RdsClient) SetSecondLevelMonitorInvoker(request *model.SetSecondLevelMonitorRequest) *SetSecondLevelMonitorInvoker {
	requestDef := GenReqDefForSetSecondLevelMonitor()
	return &SetSecondLevelMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetSecurityGroup(request *model.SetSecurityGroupRequest) (*model.SetSecurityGroupResponse, error) {
	requestDef := GenReqDefForSetSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSecurityGroupResponse), nil
	}
}

func (c *RdsClient) SetSecurityGroupInvoker(request *model.SetSecurityGroupRequest) *SetSecurityGroupInvoker {
	requestDef := GenReqDefForSetSecurityGroup()
	return &SetSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetSensitiveSlowLog(request *model.SetSensitiveSlowLogRequest) (*model.SetSensitiveSlowLogResponse, error) {
	requestDef := GenReqDefForSetSensitiveSlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSensitiveSlowLogResponse), nil
	}
}

func (c *RdsClient) SetSensitiveSlowLogInvoker(request *model.SetSensitiveSlowLogRequest) *SetSensitiveSlowLogInvoker {
	requestDef := GenReqDefForSetSensitiveSlowLog()
	return &SetSensitiveSlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowAuditlogDownloadLink(request *model.ShowAuditlogDownloadLinkRequest) (*model.ShowAuditlogDownloadLinkResponse, error) {
	requestDef := GenReqDefForShowAuditlogDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditlogDownloadLinkResponse), nil
	}
}

func (c *RdsClient) ShowAuditlogDownloadLinkInvoker(request *model.ShowAuditlogDownloadLinkRequest) *ShowAuditlogDownloadLinkInvoker {
	requestDef := GenReqDefForShowAuditlogDownloadLink()
	return &ShowAuditlogDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowAuditlogPolicy(request *model.ShowAuditlogPolicyRequest) (*model.ShowAuditlogPolicyResponse, error) {
	requestDef := GenReqDefForShowAuditlogPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditlogPolicyResponse), nil
	}
}

func (c *RdsClient) ShowAuditlogPolicyInvoker(request *model.ShowAuditlogPolicyRequest) *ShowAuditlogPolicyInvoker {
	requestDef := GenReqDefForShowAuditlogPolicy()
	return &ShowAuditlogPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowAutoEnlargePolicy(request *model.ShowAutoEnlargePolicyRequest) (*model.ShowAutoEnlargePolicyResponse, error) {
	requestDef := GenReqDefForShowAutoEnlargePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoEnlargePolicyResponse), nil
	}
}

func (c *RdsClient) ShowAutoEnlargePolicyInvoker(request *model.ShowAutoEnlargePolicyRequest) *ShowAutoEnlargePolicyInvoker {
	requestDef := GenReqDefForShowAutoEnlargePolicy()
	return &ShowAutoEnlargePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowBackupDownloadLink(request *model.ShowBackupDownloadLinkRequest) (*model.ShowBackupDownloadLinkResponse, error) {
	requestDef := GenReqDefForShowBackupDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupDownloadLinkResponse), nil
	}
}

func (c *RdsClient) ShowBackupDownloadLinkInvoker(request *model.ShowBackupDownloadLinkRequest) *ShowBackupDownloadLinkInvoker {
	requestDef := GenReqDefForShowBackupDownloadLink()
	return &ShowBackupDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowBackupPolicy(request *model.ShowBackupPolicyRequest) (*model.ShowBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPolicyResponse), nil
	}
}

func (c *RdsClient) ShowBackupPolicyInvoker(request *model.ShowBackupPolicyRequest) *ShowBackupPolicyInvoker {
	requestDef := GenReqDefForShowBackupPolicy()
	return &ShowBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowBinlogClearPolicy(request *model.ShowBinlogClearPolicyRequest) (*model.ShowBinlogClearPolicyResponse, error) {
	requestDef := GenReqDefForShowBinlogClearPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBinlogClearPolicyResponse), nil
	}
}

func (c *RdsClient) ShowBinlogClearPolicyInvoker(request *model.ShowBinlogClearPolicyRequest) *ShowBinlogClearPolicyInvoker {
	requestDef := GenReqDefForShowBinlogClearPolicy()
	return &ShowBinlogClearPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowConfiguration(request *model.ShowConfigurationRequest) (*model.ShowConfigurationResponse, error) {
	requestDef := GenReqDefForShowConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationResponse), nil
	}
}

func (c *RdsClient) ShowConfigurationInvoker(request *model.ShowConfigurationRequest) *ShowConfigurationInvoker {
	requestDef := GenReqDefForShowConfiguration()
	return &ShowConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowDnsName(request *model.ShowDnsNameRequest) (*model.ShowDnsNameResponse, error) {
	requestDef := GenReqDefForShowDnsName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDnsNameResponse), nil
	}
}

func (c *RdsClient) ShowDnsNameInvoker(request *model.ShowDnsNameRequest) *ShowDnsNameInvoker {
	requestDef := GenReqDefForShowDnsName()
	return &ShowDnsNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowDomainName(request *model.ShowDomainNameRequest) (*model.ShowDomainNameResponse, error) {
	requestDef := GenReqDefForShowDomainName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainNameResponse), nil
	}
}

func (c *RdsClient) ShowDomainNameInvoker(request *model.ShowDomainNameRequest) *ShowDomainNameInvoker {
	requestDef := GenReqDefForShowDomainName()
	return &ShowDomainNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowDrReplicaStatus(request *model.ShowDrReplicaStatusRequest) (*model.ShowDrReplicaStatusResponse, error) {
	requestDef := GenReqDefForShowDrReplicaStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDrReplicaStatusResponse), nil
	}
}

func (c *RdsClient) ShowDrReplicaStatusInvoker(request *model.ShowDrReplicaStatusRequest) *ShowDrReplicaStatusInvoker {
	requestDef := GenReqDefForShowDrReplicaStatus()
	return &ShowDrReplicaStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowInstanceConfiguration(request *model.ShowInstanceConfigurationRequest) (*model.ShowInstanceConfigurationResponse, error) {
	requestDef := GenReqDefForShowInstanceConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceConfigurationResponse), nil
	}
}

func (c *RdsClient) ShowInstanceConfigurationInvoker(request *model.ShowInstanceConfigurationRequest) *ShowInstanceConfigurationInvoker {
	requestDef := GenReqDefForShowInstanceConfiguration()
	return &ShowInstanceConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowOffSiteBackupPolicy(request *model.ShowOffSiteBackupPolicyRequest) (*model.ShowOffSiteBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowOffSiteBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOffSiteBackupPolicyResponse), nil
	}
}

func (c *RdsClient) ShowOffSiteBackupPolicyInvoker(request *model.ShowOffSiteBackupPolicyRequest) *ShowOffSiteBackupPolicyInvoker {
	requestDef := GenReqDefForShowOffSiteBackupPolicy()
	return &ShowOffSiteBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

func (c *RdsClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowRecyclePolicy(request *model.ShowRecyclePolicyRequest) (*model.ShowRecyclePolicyResponse, error) {
	requestDef := GenReqDefForShowRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecyclePolicyResponse), nil
	}
}

func (c *RdsClient) ShowRecyclePolicyInvoker(request *model.ShowRecyclePolicyRequest) *ShowRecyclePolicyInvoker {
	requestDef := GenReqDefForShowRecyclePolicy()
	return &ShowRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowReplicationStatus(request *model.ShowReplicationStatusRequest) (*model.ShowReplicationStatusResponse, error) {
	requestDef := GenReqDefForShowReplicationStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReplicationStatusResponse), nil
	}
}

func (c *RdsClient) ShowReplicationStatusInvoker(request *model.ShowReplicationStatusRequest) *ShowReplicationStatusInvoker {
	requestDef := GenReqDefForShowReplicationStatus()
	return &ShowReplicationStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowSecondLevelMonitoring(request *model.ShowSecondLevelMonitoringRequest) (*model.ShowSecondLevelMonitoringResponse, error) {
	requestDef := GenReqDefForShowSecondLevelMonitoring()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecondLevelMonitoringResponse), nil
	}
}

func (c *RdsClient) ShowSecondLevelMonitoringInvoker(request *model.ShowSecondLevelMonitoringRequest) *ShowSecondLevelMonitoringInvoker {
	requestDef := GenReqDefForShowSecondLevelMonitoring()
	return &ShowSecondLevelMonitoringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StartFailover(request *model.StartFailoverRequest) (*model.StartFailoverResponse, error) {
	requestDef := GenReqDefForStartFailover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartFailoverResponse), nil
	}
}

func (c *RdsClient) StartFailoverInvoker(request *model.StartFailoverRequest) *StartFailoverInvoker {
	requestDef := GenReqDefForStartFailover()
	return &StartFailoverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StartInstanceEnlargeVolumeAction(request *model.StartInstanceEnlargeVolumeActionRequest) (*model.StartInstanceEnlargeVolumeActionResponse, error) {
	requestDef := GenReqDefForStartInstanceEnlargeVolumeAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceEnlargeVolumeActionResponse), nil
	}
}

func (c *RdsClient) StartInstanceEnlargeVolumeActionInvoker(request *model.StartInstanceEnlargeVolumeActionRequest) *StartInstanceEnlargeVolumeActionInvoker {
	requestDef := GenReqDefForStartInstanceEnlargeVolumeAction()
	return &StartInstanceEnlargeVolumeActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StartInstanceRestartAction(request *model.StartInstanceRestartActionRequest) (*model.StartInstanceRestartActionResponse, error) {
	requestDef := GenReqDefForStartInstanceRestartAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceRestartActionResponse), nil
	}
}

func (c *RdsClient) StartInstanceRestartActionInvoker(request *model.StartInstanceRestartActionRequest) *StartInstanceRestartActionInvoker {
	requestDef := GenReqDefForStartInstanceRestartAction()
	return &StartInstanceRestartActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StartInstanceSingleToHaAction(request *model.StartInstanceSingleToHaActionRequest) (*model.StartInstanceSingleToHaActionResponse, error) {
	requestDef := GenReqDefForStartInstanceSingleToHaAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceSingleToHaActionResponse), nil
	}
}

func (c *RdsClient) StartInstanceSingleToHaActionInvoker(request *model.StartInstanceSingleToHaActionRequest) *StartInstanceSingleToHaActionInvoker {
	requestDef := GenReqDefForStartInstanceSingleToHaAction()
	return &StartInstanceSingleToHaActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StartRecyclePolicy(request *model.StartRecyclePolicyRequest) (*model.StartRecyclePolicyResponse, error) {
	requestDef := GenReqDefForStartRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartRecyclePolicyResponse), nil
	}
}

func (c *RdsClient) StartRecyclePolicyInvoker(request *model.StartRecyclePolicyRequest) *StartRecyclePolicyInvoker {
	requestDef := GenReqDefForStartRecyclePolicy()
	return &StartRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StartResizeFlavorAction(request *model.StartResizeFlavorActionRequest) (*model.StartResizeFlavorActionResponse, error) {
	requestDef := GenReqDefForStartResizeFlavorAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartResizeFlavorActionResponse), nil
	}
}

func (c *RdsClient) StartResizeFlavorActionInvoker(request *model.StartResizeFlavorActionRequest) *StartResizeFlavorActionInvoker {
	requestDef := GenReqDefForStartResizeFlavorAction()
	return &StartResizeFlavorActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StartupInstance(request *model.StartupInstanceRequest) (*model.StartupInstanceResponse, error) {
	requestDef := GenReqDefForStartupInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartupInstanceResponse), nil
	}
}

func (c *RdsClient) StartupInstanceInvoker(request *model.StartupInstanceRequest) *StartupInstanceInvoker {
	requestDef := GenReqDefForStartupInstance()
	return &StartupInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StopInstance(request *model.StopInstanceRequest) (*model.StopInstanceResponse, error) {
	requestDef := GenReqDefForStopInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInstanceResponse), nil
	}
}

func (c *RdsClient) StopInstanceInvoker(request *model.StopInstanceRequest) *StopInstanceInvoker {
	requestDef := GenReqDefForStopInstance()
	return &StopInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SwitchSsl(request *model.SwitchSslRequest) (*model.SwitchSslResponse, error) {
	requestDef := GenReqDefForSwitchSsl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSslResponse), nil
	}
}

func (c *RdsClient) SwitchSslInvoker(request *model.SwitchSslRequest) *SwitchSslInvoker {
	requestDef := GenReqDefForSwitchSsl()
	return &SwitchSslInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateConfiguration(request *model.UpdateConfigurationRequest) (*model.UpdateConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationResponse), nil
	}
}

func (c *RdsClient) UpdateConfigurationInvoker(request *model.UpdateConfigurationRequest) *UpdateConfigurationInvoker {
	requestDef := GenReqDefForUpdateConfiguration()
	return &UpdateConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateDataIp(request *model.UpdateDataIpRequest) (*model.UpdateDataIpResponse, error) {
	requestDef := GenReqDefForUpdateDataIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataIpResponse), nil
	}
}

func (c *RdsClient) UpdateDataIpInvoker(request *model.UpdateDataIpRequest) *UpdateDataIpInvoker {
	requestDef := GenReqDefForUpdateDataIp()
	return &UpdateDataIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateDnsName(request *model.UpdateDnsNameRequest) (*model.UpdateDnsNameResponse, error) {
	requestDef := GenReqDefForUpdateDnsName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDnsNameResponse), nil
	}
}

func (c *RdsClient) UpdateDnsNameInvoker(request *model.UpdateDnsNameRequest) *UpdateDnsNameInvoker {
	requestDef := GenReqDefForUpdateDnsName()
	return &UpdateDnsNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateInstanceConfiguration(request *model.UpdateInstanceConfigurationRequest) (*model.UpdateInstanceConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateInstanceConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceConfigurationResponse), nil
	}
}

func (c *RdsClient) UpdateInstanceConfigurationInvoker(request *model.UpdateInstanceConfigurationRequest) *UpdateInstanceConfigurationInvoker {
	requestDef := GenReqDefForUpdateInstanceConfiguration()
	return &UpdateInstanceConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateInstanceConfigurationAsync(request *model.UpdateInstanceConfigurationAsyncRequest) (*model.UpdateInstanceConfigurationAsyncResponse, error) {
	requestDef := GenReqDefForUpdateInstanceConfigurationAsync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceConfigurationAsyncResponse), nil
	}
}

func (c *RdsClient) UpdateInstanceConfigurationAsyncInvoker(request *model.UpdateInstanceConfigurationAsyncRequest) *UpdateInstanceConfigurationAsyncInvoker {
	requestDef := GenReqDefForUpdateInstanceConfigurationAsync()
	return &UpdateInstanceConfigurationAsyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateInstanceName(request *model.UpdateInstanceNameRequest) (*model.UpdateInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceNameResponse), nil
	}
}

func (c *RdsClient) UpdateInstanceNameInvoker(request *model.UpdateInstanceNameRequest) *UpdateInstanceNameInvoker {
	requestDef := GenReqDefForUpdateInstanceName()
	return &UpdateInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdatePort(request *model.UpdatePortRequest) (*model.UpdatePortResponse, error) {
	requestDef := GenReqDefForUpdatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePortResponse), nil
	}
}

func (c *RdsClient) UpdatePortInvoker(request *model.UpdatePortRequest) *UpdatePortInvoker {
	requestDef := GenReqDefForUpdatePort()
	return &UpdatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdatePostgresqlInstanceAlias(request *model.UpdatePostgresqlInstanceAliasRequest) (*model.UpdatePostgresqlInstanceAliasResponse, error) {
	requestDef := GenReqDefForUpdatePostgresqlInstanceAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePostgresqlInstanceAliasResponse), nil
	}
}

func (c *RdsClient) UpdatePostgresqlInstanceAliasInvoker(request *model.UpdatePostgresqlInstanceAliasRequest) *UpdatePostgresqlInstanceAliasInvoker {
	requestDef := GenReqDefForUpdatePostgresqlInstanceAlias()
	return &UpdatePostgresqlInstanceAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpgradeDbVersion(request *model.UpgradeDbVersionRequest) (*model.UpgradeDbVersionResponse, error) {
	requestDef := GenReqDefForUpgradeDbVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeDbVersionResponse), nil
	}
}

func (c *RdsClient) UpgradeDbVersionInvoker(request *model.UpgradeDbVersionRequest) *UpgradeDbVersionInvoker {
	requestDef := GenReqDefForUpgradeDbVersion()
	return &UpgradeDbVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListApiVersion(request *model.ListApiVersionRequest) (*model.ListApiVersionResponse, error) {
	requestDef := GenReqDefForListApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionResponse), nil
	}
}

func (c *RdsClient) ListApiVersionInvoker(request *model.ListApiVersionRequest) *ListApiVersionInvoker {
	requestDef := GenReqDefForListApiVersion()
	return &ListApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListApiVersionNew(request *model.ListApiVersionNewRequest) (*model.ListApiVersionNewResponse, error) {
	requestDef := GenReqDefForListApiVersionNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionNewResponse), nil
	}
}

func (c *RdsClient) ListApiVersionNewInvoker(request *model.ListApiVersionNewRequest) *ListApiVersionNewInvoker {
	requestDef := GenReqDefForListApiVersionNew()
	return &ListApiVersionNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

func (c *RdsClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) AllowDbUserPrivilege(request *model.AllowDbUserPrivilegeRequest) (*model.AllowDbUserPrivilegeResponse, error) {
	requestDef := GenReqDefForAllowDbUserPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowDbUserPrivilegeResponse), nil
	}
}

func (c *RdsClient) AllowDbUserPrivilegeInvoker(request *model.AllowDbUserPrivilegeRequest) *AllowDbUserPrivilegeInvoker {
	requestDef := GenReqDefForAllowDbUserPrivilege()
	return &AllowDbUserPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateDatabase(request *model.CreateDatabaseRequest) (*model.CreateDatabaseResponse, error) {
	requestDef := GenReqDefForCreateDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseResponse), nil
	}
}

func (c *RdsClient) CreateDatabaseInvoker(request *model.CreateDatabaseRequest) *CreateDatabaseInvoker {
	requestDef := GenReqDefForCreateDatabase()
	return &CreateDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateDbUser(request *model.CreateDbUserRequest) (*model.CreateDbUserResponse, error) {
	requestDef := GenReqDefForCreateDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDbUserResponse), nil
	}
}

func (c *RdsClient) CreateDbUserInvoker(request *model.CreateDbUserRequest) *CreateDbUserInvoker {
	requestDef := GenReqDefForCreateDbUser()
	return &CreateDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeleteDatabase(request *model.DeleteDatabaseRequest) (*model.DeleteDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseResponse), nil
	}
}

func (c *RdsClient) DeleteDatabaseInvoker(request *model.DeleteDatabaseRequest) *DeleteDatabaseInvoker {
	requestDef := GenReqDefForDeleteDatabase()
	return &DeleteDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeleteDbUser(request *model.DeleteDbUserRequest) (*model.DeleteDbUserResponse, error) {
	requestDef := GenReqDefForDeleteDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDbUserResponse), nil
	}
}

func (c *RdsClient) DeleteDbUserInvoker(request *model.DeleteDbUserRequest) *DeleteDbUserInvoker {
	requestDef := GenReqDefForDeleteDbUser()
	return &DeleteDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListAuthorizedDatabases(request *model.ListAuthorizedDatabasesRequest) (*model.ListAuthorizedDatabasesResponse, error) {
	requestDef := GenReqDefForListAuthorizedDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizedDatabasesResponse), nil
	}
}

func (c *RdsClient) ListAuthorizedDatabasesInvoker(request *model.ListAuthorizedDatabasesRequest) *ListAuthorizedDatabasesInvoker {
	requestDef := GenReqDefForListAuthorizedDatabases()
	return &ListAuthorizedDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListAuthorizedDbUsers(request *model.ListAuthorizedDbUsersRequest) (*model.ListAuthorizedDbUsersResponse, error) {
	requestDef := GenReqDefForListAuthorizedDbUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizedDbUsersResponse), nil
	}
}

func (c *RdsClient) ListAuthorizedDbUsersInvoker(request *model.ListAuthorizedDbUsersRequest) *ListAuthorizedDbUsersInvoker {
	requestDef := GenReqDefForListAuthorizedDbUsers()
	return &ListAuthorizedDbUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListDatabases(request *model.ListDatabasesRequest) (*model.ListDatabasesResponse, error) {
	requestDef := GenReqDefForListDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabasesResponse), nil
	}
}

func (c *RdsClient) ListDatabasesInvoker(request *model.ListDatabasesRequest) *ListDatabasesInvoker {
	requestDef := GenReqDefForListDatabases()
	return &ListDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListDbUsers(request *model.ListDbUsersRequest) (*model.ListDbUsersResponse, error) {
	requestDef := GenReqDefForListDbUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbUsersResponse), nil
	}
}

func (c *RdsClient) ListDbUsersInvoker(request *model.ListDbUsersRequest) *ListDbUsersInvoker {
	requestDef := GenReqDefForListDbUsers()
	return &ListDbUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ResetPwd(request *model.ResetPwdRequest) (*model.ResetPwdResponse, error) {
	requestDef := GenReqDefForResetPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPwdResponse), nil
	}
}

func (c *RdsClient) ResetPwdInvoker(request *model.ResetPwdRequest) *ResetPwdInvoker {
	requestDef := GenReqDefForResetPwd()
	return &ResetPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) Revoke(request *model.RevokeRequest) (*model.RevokeResponse, error) {
	requestDef := GenReqDefForRevoke()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevokeResponse), nil
	}
}

func (c *RdsClient) RevokeInvoker(request *model.RevokeRequest) *RevokeInvoker {
	requestDef := GenReqDefForRevoke()
	return &RevokeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetDbUserPwd(request *model.SetDbUserPwdRequest) (*model.SetDbUserPwdResponse, error) {
	requestDef := GenReqDefForSetDbUserPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetDbUserPwdResponse), nil
	}
}

func (c *RdsClient) SetDbUserPwdInvoker(request *model.SetDbUserPwdRequest) *SetDbUserPwdInvoker {
	requestDef := GenReqDefForSetDbUserPwd()
	return &SetDbUserPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetReadOnlySwitch(request *model.SetReadOnlySwitchRequest) (*model.SetReadOnlySwitchResponse, error) {
	requestDef := GenReqDefForSetReadOnlySwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetReadOnlySwitchResponse), nil
	}
}

func (c *RdsClient) SetReadOnlySwitchInvoker(request *model.SetReadOnlySwitchRequest) *SetReadOnlySwitchInvoker {
	requestDef := GenReqDefForSetReadOnlySwitch()
	return &SetReadOnlySwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateDatabase(request *model.UpdateDatabaseRequest) (*model.UpdateDatabaseResponse, error) {
	requestDef := GenReqDefForUpdateDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatabaseResponse), nil
	}
}

func (c *RdsClient) UpdateDatabaseInvoker(request *model.UpdateDatabaseRequest) *UpdateDatabaseInvoker {
	requestDef := GenReqDefForUpdateDatabase()
	return &UpdateDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateDbUserComment(request *model.UpdateDbUserCommentRequest) (*model.UpdateDbUserCommentResponse, error) {
	requestDef := GenReqDefForUpdateDbUserComment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDbUserCommentResponse), nil
	}
}

func (c *RdsClient) UpdateDbUserCommentInvoker(request *model.UpdateDbUserCommentRequest) *UpdateDbUserCommentInvoker {
	requestDef := GenReqDefForUpdateDbUserComment()
	return &UpdateDbUserCommentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) AllowDbPrivilege(request *model.AllowDbPrivilegeRequest) (*model.AllowDbPrivilegeResponse, error) {
	requestDef := GenReqDefForAllowDbPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowDbPrivilegeResponse), nil
	}
}

func (c *RdsClient) AllowDbPrivilegeInvoker(request *model.AllowDbPrivilegeRequest) *AllowDbPrivilegeInvoker {
	requestDef := GenReqDefForAllowDbPrivilege()
	return &AllowDbPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ChangeProxyScale(request *model.ChangeProxyScaleRequest) (*model.ChangeProxyScaleResponse, error) {
	requestDef := GenReqDefForChangeProxyScale()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeProxyScaleResponse), nil
	}
}

func (c *RdsClient) ChangeProxyScaleInvoker(request *model.ChangeProxyScaleRequest) *ChangeProxyScaleInvoker {
	requestDef := GenReqDefForChangeProxyScale()
	return &ChangeProxyScaleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ChangeTheDelayThreshold(request *model.ChangeTheDelayThresholdRequest) (*model.ChangeTheDelayThresholdResponse, error) {
	requestDef := GenReqDefForChangeTheDelayThreshold()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeTheDelayThresholdResponse), nil
	}
}

func (c *RdsClient) ChangeTheDelayThresholdInvoker(request *model.ChangeTheDelayThresholdRequest) *ChangeTheDelayThresholdInvoker {
	requestDef := GenReqDefForChangeTheDelayThreshold()
	return &ChangeTheDelayThresholdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreatePostgresqlDatabase(request *model.CreatePostgresqlDatabaseRequest) (*model.CreatePostgresqlDatabaseResponse, error) {
	requestDef := GenReqDefForCreatePostgresqlDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostgresqlDatabaseResponse), nil
	}
}

func (c *RdsClient) CreatePostgresqlDatabaseInvoker(request *model.CreatePostgresqlDatabaseRequest) *CreatePostgresqlDatabaseInvoker {
	requestDef := GenReqDefForCreatePostgresqlDatabase()
	return &CreatePostgresqlDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreatePostgresqlDatabaseSchema(request *model.CreatePostgresqlDatabaseSchemaRequest) (*model.CreatePostgresqlDatabaseSchemaResponse, error) {
	requestDef := GenReqDefForCreatePostgresqlDatabaseSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostgresqlDatabaseSchemaResponse), nil
	}
}

func (c *RdsClient) CreatePostgresqlDatabaseSchemaInvoker(request *model.CreatePostgresqlDatabaseSchemaRequest) *CreatePostgresqlDatabaseSchemaInvoker {
	requestDef := GenReqDefForCreatePostgresqlDatabaseSchema()
	return &CreatePostgresqlDatabaseSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreatePostgresqlDbUser(request *model.CreatePostgresqlDbUserRequest) (*model.CreatePostgresqlDbUserResponse, error) {
	requestDef := GenReqDefForCreatePostgresqlDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostgresqlDbUserResponse), nil
	}
}

func (c *RdsClient) CreatePostgresqlDbUserInvoker(request *model.CreatePostgresqlDbUserRequest) *CreatePostgresqlDbUserInvoker {
	requestDef := GenReqDefForCreatePostgresqlDbUser()
	return &CreatePostgresqlDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreatePostgresqlExtension(request *model.CreatePostgresqlExtensionRequest) (*model.CreatePostgresqlExtensionResponse, error) {
	requestDef := GenReqDefForCreatePostgresqlExtension()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostgresqlExtensionResponse), nil
	}
}

func (c *RdsClient) CreatePostgresqlExtensionInvoker(request *model.CreatePostgresqlExtensionRequest) *CreatePostgresqlExtensionInvoker {
	requestDef := GenReqDefForCreatePostgresqlExtension()
	return &CreatePostgresqlExtensionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeletePostgresqlExtension(request *model.DeletePostgresqlExtensionRequest) (*model.DeletePostgresqlExtensionResponse, error) {
	requestDef := GenReqDefForDeletePostgresqlExtension()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePostgresqlExtensionResponse), nil
	}
}

func (c *RdsClient) DeletePostgresqlExtensionInvoker(request *model.DeletePostgresqlExtensionRequest) *DeletePostgresqlExtensionInvoker {
	requestDef := GenReqDefForDeletePostgresqlExtension()
	return &DeletePostgresqlExtensionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListPostgresqlDatabaseSchemas(request *model.ListPostgresqlDatabaseSchemasRequest) (*model.ListPostgresqlDatabaseSchemasResponse, error) {
	requestDef := GenReqDefForListPostgresqlDatabaseSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPostgresqlDatabaseSchemasResponse), nil
	}
}

func (c *RdsClient) ListPostgresqlDatabaseSchemasInvoker(request *model.ListPostgresqlDatabaseSchemasRequest) *ListPostgresqlDatabaseSchemasInvoker {
	requestDef := GenReqDefForListPostgresqlDatabaseSchemas()
	return &ListPostgresqlDatabaseSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListPostgresqlDatabases(request *model.ListPostgresqlDatabasesRequest) (*model.ListPostgresqlDatabasesResponse, error) {
	requestDef := GenReqDefForListPostgresqlDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPostgresqlDatabasesResponse), nil
	}
}

func (c *RdsClient) ListPostgresqlDatabasesInvoker(request *model.ListPostgresqlDatabasesRequest) *ListPostgresqlDatabasesInvoker {
	requestDef := GenReqDefForListPostgresqlDatabases()
	return &ListPostgresqlDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListPostgresqlDbUserPaginated(request *model.ListPostgresqlDbUserPaginatedRequest) (*model.ListPostgresqlDbUserPaginatedResponse, error) {
	requestDef := GenReqDefForListPostgresqlDbUserPaginated()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPostgresqlDbUserPaginatedResponse), nil
	}
}

func (c *RdsClient) ListPostgresqlDbUserPaginatedInvoker(request *model.ListPostgresqlDbUserPaginatedRequest) *ListPostgresqlDbUserPaginatedInvoker {
	requestDef := GenReqDefForListPostgresqlDbUserPaginated()
	return &ListPostgresqlDbUserPaginatedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListPostgresqlExtension(request *model.ListPostgresqlExtensionRequest) (*model.ListPostgresqlExtensionResponse, error) {
	requestDef := GenReqDefForListPostgresqlExtension()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPostgresqlExtensionResponse), nil
	}
}

func (c *RdsClient) ListPostgresqlExtensionInvoker(request *model.ListPostgresqlExtensionRequest) *ListPostgresqlExtensionInvoker {
	requestDef := GenReqDefForListPostgresqlExtension()
	return &ListPostgresqlExtensionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SearchQueryScaleComputeFlavors(request *model.SearchQueryScaleComputeFlavorsRequest) (*model.SearchQueryScaleComputeFlavorsResponse, error) {
	requestDef := GenReqDefForSearchQueryScaleComputeFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQueryScaleComputeFlavorsResponse), nil
	}
}

func (c *RdsClient) SearchQueryScaleComputeFlavorsInvoker(request *model.SearchQueryScaleComputeFlavorsRequest) *SearchQueryScaleComputeFlavorsInvoker {
	requestDef := GenReqDefForSearchQueryScaleComputeFlavors()
	return &SearchQueryScaleComputeFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SearchQueryScaleFlavors(request *model.SearchQueryScaleFlavorsRequest) (*model.SearchQueryScaleFlavorsResponse, error) {
	requestDef := GenReqDefForSearchQueryScaleFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQueryScaleFlavorsResponse), nil
	}
}

func (c *RdsClient) SearchQueryScaleFlavorsInvoker(request *model.SearchQueryScaleFlavorsRequest) *SearchQueryScaleFlavorsInvoker {
	requestDef := GenReqDefForSearchQueryScaleFlavors()
	return &SearchQueryScaleFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetDatabaseUserPrivilege(request *model.SetDatabaseUserPrivilegeRequest) (*model.SetDatabaseUserPrivilegeResponse, error) {
	requestDef := GenReqDefForSetDatabaseUserPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetDatabaseUserPrivilegeResponse), nil
	}
}

func (c *RdsClient) SetDatabaseUserPrivilegeInvoker(request *model.SetDatabaseUserPrivilegeRequest) *SetDatabaseUserPrivilegeInvoker {
	requestDef := GenReqDefForSetDatabaseUserPrivilege()
	return &SetDatabaseUserPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) SetPostgresqlDbUserPwd(request *model.SetPostgresqlDbUserPwdRequest) (*model.SetPostgresqlDbUserPwdResponse, error) {
	requestDef := GenReqDefForSetPostgresqlDbUserPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetPostgresqlDbUserPwdResponse), nil
	}
}

func (c *RdsClient) SetPostgresqlDbUserPwdInvoker(request *model.SetPostgresqlDbUserPwdRequest) *SetPostgresqlDbUserPwdInvoker {
	requestDef := GenReqDefForSetPostgresqlDbUserPwd()
	return &SetPostgresqlDbUserPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowInformationAboutDatabaseProxy(request *model.ShowInformationAboutDatabaseProxyRequest) (*model.ShowInformationAboutDatabaseProxyResponse, error) {
	requestDef := GenReqDefForShowInformationAboutDatabaseProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInformationAboutDatabaseProxyResponse), nil
	}
}

func (c *RdsClient) ShowInformationAboutDatabaseProxyInvoker(request *model.ShowInformationAboutDatabaseProxyRequest) *ShowInformationAboutDatabaseProxyInvoker {
	requestDef := GenReqDefForShowInformationAboutDatabaseProxy()
	return &ShowInformationAboutDatabaseProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ShowPostgresqlParamValue(request *model.ShowPostgresqlParamValueRequest) (*model.ShowPostgresqlParamValueResponse, error) {
	requestDef := GenReqDefForShowPostgresqlParamValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPostgresqlParamValueResponse), nil
	}
}

func (c *RdsClient) ShowPostgresqlParamValueInvoker(request *model.ShowPostgresqlParamValueRequest) *ShowPostgresqlParamValueInvoker {
	requestDef := GenReqDefForShowPostgresqlParamValue()
	return &ShowPostgresqlParamValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StartDatabaseProxy(request *model.StartDatabaseProxyRequest) (*model.StartDatabaseProxyResponse, error) {
	requestDef := GenReqDefForStartDatabaseProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDatabaseProxyResponse), nil
	}
}

func (c *RdsClient) StartDatabaseProxyInvoker(request *model.StartDatabaseProxyRequest) *StartDatabaseProxyInvoker {
	requestDef := GenReqDefForStartDatabaseProxy()
	return &StartDatabaseProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) StopDatabaseProxy(request *model.StopDatabaseProxyRequest) (*model.StopDatabaseProxyResponse, error) {
	requestDef := GenReqDefForStopDatabaseProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopDatabaseProxyResponse), nil
	}
}

func (c *RdsClient) StopDatabaseProxyInvoker(request *model.StopDatabaseProxyRequest) *StopDatabaseProxyInvoker {
	requestDef := GenReqDefForStopDatabaseProxy()
	return &StopDatabaseProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateDbUserPrivilege(request *model.UpdateDbUserPrivilegeRequest) (*model.UpdateDbUserPrivilegeResponse, error) {
	requestDef := GenReqDefForUpdateDbUserPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDbUserPrivilegeResponse), nil
	}
}

func (c *RdsClient) UpdateDbUserPrivilegeInvoker(request *model.UpdateDbUserPrivilegeRequest) *UpdateDbUserPrivilegeInvoker {
	requestDef := GenReqDefForUpdateDbUserPrivilege()
	return &UpdateDbUserPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdatePostgresqlParameterValue(request *model.UpdatePostgresqlParameterValueRequest) (*model.UpdatePostgresqlParameterValueResponse, error) {
	requestDef := GenReqDefForUpdatePostgresqlParameterValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePostgresqlParameterValueResponse), nil
	}
}

func (c *RdsClient) UpdatePostgresqlParameterValueInvoker(request *model.UpdatePostgresqlParameterValueRequest) *UpdatePostgresqlParameterValueInvoker {
	requestDef := GenReqDefForUpdatePostgresqlParameterValue()
	return &UpdatePostgresqlParameterValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) UpdateReadWeight(request *model.UpdateReadWeightRequest) (*model.UpdateReadWeightResponse, error) {
	requestDef := GenReqDefForUpdateReadWeight()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReadWeightResponse), nil
	}
}

func (c *RdsClient) UpdateReadWeightInvoker(request *model.UpdateReadWeightRequest) *UpdateReadWeightInvoker {
	requestDef := GenReqDefForUpdateReadWeight()
	return &UpdateReadWeightInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) AllowSqlserverDbUserPrivilege(request *model.AllowSqlserverDbUserPrivilegeRequest) (*model.AllowSqlserverDbUserPrivilegeResponse, error) {
	requestDef := GenReqDefForAllowSqlserverDbUserPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowSqlserverDbUserPrivilegeResponse), nil
	}
}

func (c *RdsClient) AllowSqlserverDbUserPrivilegeInvoker(request *model.AllowSqlserverDbUserPrivilegeRequest) *AllowSqlserverDbUserPrivilegeInvoker {
	requestDef := GenReqDefForAllowSqlserverDbUserPrivilege()
	return &AllowSqlserverDbUserPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) BatchAddMsdtcs(request *model.BatchAddMsdtcsRequest) (*model.BatchAddMsdtcsResponse, error) {
	requestDef := GenReqDefForBatchAddMsdtcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddMsdtcsResponse), nil
	}
}

func (c *RdsClient) BatchAddMsdtcsInvoker(request *model.BatchAddMsdtcsRequest) *BatchAddMsdtcsInvoker {
	requestDef := GenReqDefForBatchAddMsdtcs()
	return &BatchAddMsdtcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateSqlserverDatabase(request *model.CreateSqlserverDatabaseRequest) (*model.CreateSqlserverDatabaseResponse, error) {
	requestDef := GenReqDefForCreateSqlserverDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlserverDatabaseResponse), nil
	}
}

func (c *RdsClient) CreateSqlserverDatabaseInvoker(request *model.CreateSqlserverDatabaseRequest) *CreateSqlserverDatabaseInvoker {
	requestDef := GenReqDefForCreateSqlserverDatabase()
	return &CreateSqlserverDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) CreateSqlserverDbUser(request *model.CreateSqlserverDbUserRequest) (*model.CreateSqlserverDbUserResponse, error) {
	requestDef := GenReqDefForCreateSqlserverDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlserverDbUserResponse), nil
	}
}

func (c *RdsClient) CreateSqlserverDbUserInvoker(request *model.CreateSqlserverDbUserRequest) *CreateSqlserverDbUserInvoker {
	requestDef := GenReqDefForCreateSqlserverDbUser()
	return &CreateSqlserverDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeleteSqlserverDatabase(request *model.DeleteSqlserverDatabaseRequest) (*model.DeleteSqlserverDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteSqlserverDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlserverDatabaseResponse), nil
	}
}

func (c *RdsClient) DeleteSqlserverDatabaseInvoker(request *model.DeleteSqlserverDatabaseRequest) *DeleteSqlserverDatabaseInvoker {
	requestDef := GenReqDefForDeleteSqlserverDatabase()
	return &DeleteSqlserverDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeleteSqlserverDatabaseEx(request *model.DeleteSqlserverDatabaseExRequest) (*model.DeleteSqlserverDatabaseExResponse, error) {
	requestDef := GenReqDefForDeleteSqlserverDatabaseEx()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlserverDatabaseExResponse), nil
	}
}

func (c *RdsClient) DeleteSqlserverDatabaseExInvoker(request *model.DeleteSqlserverDatabaseExRequest) *DeleteSqlserverDatabaseExInvoker {
	requestDef := GenReqDefForDeleteSqlserverDatabaseEx()
	return &DeleteSqlserverDatabaseExInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) DeleteSqlserverDbUser(request *model.DeleteSqlserverDbUserRequest) (*model.DeleteSqlserverDbUserResponse, error) {
	requestDef := GenReqDefForDeleteSqlserverDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlserverDbUserResponse), nil
	}
}

func (c *RdsClient) DeleteSqlserverDbUserInvoker(request *model.DeleteSqlserverDbUserRequest) *DeleteSqlserverDbUserInvoker {
	requestDef := GenReqDefForDeleteSqlserverDbUser()
	return &DeleteSqlserverDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListAuthorizedSqlserverDbUsers(request *model.ListAuthorizedSqlserverDbUsersRequest) (*model.ListAuthorizedSqlserverDbUsersResponse, error) {
	requestDef := GenReqDefForListAuthorizedSqlserverDbUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizedSqlserverDbUsersResponse), nil
	}
}

func (c *RdsClient) ListAuthorizedSqlserverDbUsersInvoker(request *model.ListAuthorizedSqlserverDbUsersRequest) *ListAuthorizedSqlserverDbUsersInvoker {
	requestDef := GenReqDefForListAuthorizedSqlserverDbUsers()
	return &ListAuthorizedSqlserverDbUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListMsdtcHosts(request *model.ListMsdtcHostsRequest) (*model.ListMsdtcHostsResponse, error) {
	requestDef := GenReqDefForListMsdtcHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMsdtcHostsResponse), nil
	}
}

func (c *RdsClient) ListMsdtcHostsInvoker(request *model.ListMsdtcHostsRequest) *ListMsdtcHostsInvoker {
	requestDef := GenReqDefForListMsdtcHosts()
	return &ListMsdtcHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListSqlserverDatabases(request *model.ListSqlserverDatabasesRequest) (*model.ListSqlserverDatabasesResponse, error) {
	requestDef := GenReqDefForListSqlserverDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlserverDatabasesResponse), nil
	}
}

func (c *RdsClient) ListSqlserverDatabasesInvoker(request *model.ListSqlserverDatabasesRequest) *ListSqlserverDatabasesInvoker {
	requestDef := GenReqDefForListSqlserverDatabases()
	return &ListSqlserverDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) ListSqlserverDbUsers(request *model.ListSqlserverDbUsersRequest) (*model.ListSqlserverDbUsersResponse, error) {
	requestDef := GenReqDefForListSqlserverDbUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlserverDbUsersResponse), nil
	}
}

func (c *RdsClient) ListSqlserverDbUsersInvoker(request *model.ListSqlserverDbUsersRequest) *ListSqlserverDbUsersInvoker {
	requestDef := GenReqDefForListSqlserverDbUsers()
	return &ListSqlserverDbUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *RdsClient) RevokeSqlserverDbUserPrivilege(request *model.RevokeSqlserverDbUserPrivilegeRequest) (*model.RevokeSqlserverDbUserPrivilegeResponse, error) {
	requestDef := GenReqDefForRevokeSqlserverDbUserPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevokeSqlserverDbUserPrivilegeResponse), nil
	}
}

func (c *RdsClient) RevokeSqlserverDbUserPrivilegeInvoker(request *model.RevokeSqlserverDbUserPrivilegeRequest) *RevokeSqlserverDbUserPrivilegeInvoker {
	requestDef := GenReqDefForRevokeSqlserverDbUserPrivilege()
	return &RevokeSqlserverDbUserPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
