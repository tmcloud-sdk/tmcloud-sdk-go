package v3

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/drs/v3/model"
)

type DrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDrsClient(hcClient *http_client.HcHttpClient) *DrsClient {
	return &DrsClient{HcClient: hcClient}
}

func DrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *DrsClient) BatchChangeData(request *model.BatchChangeDataRequest) (*model.BatchChangeDataResponse, error) {
	requestDef := GenReqDefForBatchChangeData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchChangeDataResponse), nil
	}
}

func (c *DrsClient) BatchChangeDataInvoker(request *model.BatchChangeDataRequest) *BatchChangeDataInvoker {
	requestDef := GenReqDefForBatchChangeData()
	return &BatchChangeDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchCheckJobs(request *model.BatchCheckJobsRequest) (*model.BatchCheckJobsResponse, error) {
	requestDef := GenReqDefForBatchCheckJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckJobsResponse), nil
	}
}

func (c *DrsClient) BatchCheckJobsInvoker(request *model.BatchCheckJobsRequest) *BatchCheckJobsInvoker {
	requestDef := GenReqDefForBatchCheckJobs()
	return &BatchCheckJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchCheckResults(request *model.BatchCheckResultsRequest) (*model.BatchCheckResultsResponse, error) {
	requestDef := GenReqDefForBatchCheckResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckResultsResponse), nil
	}
}

func (c *DrsClient) BatchCheckResultsInvoker(request *model.BatchCheckResultsRequest) *BatchCheckResultsInvoker {
	requestDef := GenReqDefForBatchCheckResults()
	return &BatchCheckResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchCreateJobs(request *model.BatchCreateJobsRequest) (*model.BatchCreateJobsResponse, error) {
	requestDef := GenReqDefForBatchCreateJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateJobsResponse), nil
	}
}

func (c *DrsClient) BatchCreateJobsInvoker(request *model.BatchCreateJobsRequest) *BatchCreateJobsInvoker {
	requestDef := GenReqDefForBatchCreateJobs()
	return &BatchCreateJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchDeleteJobs(request *model.BatchDeleteJobsRequest) (*model.BatchDeleteJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobsResponse), nil
	}
}

func (c *DrsClient) BatchDeleteJobsInvoker(request *model.BatchDeleteJobsRequest) *BatchDeleteJobsInvoker {
	requestDef := GenReqDefForBatchDeleteJobs()
	return &BatchDeleteJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchListJobDetails(request *model.BatchListJobDetailsRequest) (*model.BatchListJobDetailsResponse, error) {
	requestDef := GenReqDefForBatchListJobDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListJobDetailsResponse), nil
	}
}

func (c *DrsClient) BatchListJobDetailsInvoker(request *model.BatchListJobDetailsRequest) *BatchListJobDetailsInvoker {
	requestDef := GenReqDefForBatchListJobDetails()
	return &BatchListJobDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchListJobStatus(request *model.BatchListJobStatusRequest) (*model.BatchListJobStatusResponse, error) {
	requestDef := GenReqDefForBatchListJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListJobStatusResponse), nil
	}
}

func (c *DrsClient) BatchListJobStatusInvoker(request *model.BatchListJobStatusRequest) *BatchListJobStatusInvoker {
	requestDef := GenReqDefForBatchListJobStatus()
	return &BatchListJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchListProgresses(request *model.BatchListProgressesRequest) (*model.BatchListProgressesResponse, error) {
	requestDef := GenReqDefForBatchListProgresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListProgressesResponse), nil
	}
}

func (c *DrsClient) BatchListProgressesInvoker(request *model.BatchListProgressesRequest) *BatchListProgressesInvoker {
	requestDef := GenReqDefForBatchListProgresses()
	return &BatchListProgressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchListRposAndRtos(request *model.BatchListRposAndRtosRequest) (*model.BatchListRposAndRtosResponse, error) {
	requestDef := GenReqDefForBatchListRposAndRtos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListRposAndRtosResponse), nil
	}
}

func (c *DrsClient) BatchListRposAndRtosInvoker(request *model.BatchListRposAndRtosRequest) *BatchListRposAndRtosInvoker {
	requestDef := GenReqDefForBatchListRposAndRtos()
	return &BatchListRposAndRtosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchListStructDetail(request *model.BatchListStructDetailRequest) (*model.BatchListStructDetailResponse, error) {
	requestDef := GenReqDefForBatchListStructDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListStructDetailResponse), nil
	}
}

func (c *DrsClient) BatchListStructDetailInvoker(request *model.BatchListStructDetailRequest) *BatchListStructDetailInvoker {
	requestDef := GenReqDefForBatchListStructDetail()
	return &BatchListStructDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchListStructProcess(request *model.BatchListStructProcessRequest) (*model.BatchListStructProcessResponse, error) {
	requestDef := GenReqDefForBatchListStructProcess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListStructProcessResponse), nil
	}
}

func (c *DrsClient) BatchListStructProcessInvoker(request *model.BatchListStructProcessRequest) *BatchListStructProcessInvoker {
	requestDef := GenReqDefForBatchListStructProcess()
	return &BatchListStructProcessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchResetPassword(request *model.BatchResetPasswordRequest) (*model.BatchResetPasswordResponse, error) {
	requestDef := GenReqDefForBatchResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResetPasswordResponse), nil
	}
}

func (c *DrsClient) BatchResetPasswordInvoker(request *model.BatchResetPasswordRequest) *BatchResetPasswordInvoker {
	requestDef := GenReqDefForBatchResetPassword()
	return &BatchResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchRestoreTask(request *model.BatchRestoreTaskRequest) (*model.BatchRestoreTaskResponse, error) {
	requestDef := GenReqDefForBatchRestoreTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestoreTaskResponse), nil
	}
}

func (c *DrsClient) BatchRestoreTaskInvoker(request *model.BatchRestoreTaskRequest) *BatchRestoreTaskInvoker {
	requestDef := GenReqDefForBatchRestoreTask()
	return &BatchRestoreTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchSetDefiner(request *model.BatchSetDefinerRequest) (*model.BatchSetDefinerResponse, error) {
	requestDef := GenReqDefForBatchSetDefiner()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetDefinerResponse), nil
	}
}

func (c *DrsClient) BatchSetDefinerInvoker(request *model.BatchSetDefinerRequest) *BatchSetDefinerInvoker {
	requestDef := GenReqDefForBatchSetDefiner()
	return &BatchSetDefinerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchSetObjects(request *model.BatchSetObjectsRequest) (*model.BatchSetObjectsResponse, error) {
	requestDef := GenReqDefForBatchSetObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetObjectsResponse), nil
	}
}

func (c *DrsClient) BatchSetObjectsInvoker(request *model.BatchSetObjectsRequest) *BatchSetObjectsInvoker {
	requestDef := GenReqDefForBatchSetObjects()
	return &BatchSetObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchSetPolicy(request *model.BatchSetPolicyRequest) (*model.BatchSetPolicyResponse, error) {
	requestDef := GenReqDefForBatchSetPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetPolicyResponse), nil
	}
}

func (c *DrsClient) BatchSetPolicyInvoker(request *model.BatchSetPolicyRequest) *BatchSetPolicyInvoker {
	requestDef := GenReqDefForBatchSetPolicy()
	return &BatchSetPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchSetSmn(request *model.BatchSetSmnRequest) (*model.BatchSetSmnResponse, error) {
	requestDef := GenReqDefForBatchSetSmn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetSmnResponse), nil
	}
}

func (c *DrsClient) BatchSetSmnInvoker(request *model.BatchSetSmnRequest) *BatchSetSmnInvoker {
	requestDef := GenReqDefForBatchSetSmn()
	return &BatchSetSmnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchSetSpeed(request *model.BatchSetSpeedRequest) (*model.BatchSetSpeedResponse, error) {
	requestDef := GenReqDefForBatchSetSpeed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetSpeedResponse), nil
	}
}

func (c *DrsClient) BatchSetSpeedInvoker(request *model.BatchSetSpeedRequest) *BatchSetSpeedInvoker {
	requestDef := GenReqDefForBatchSetSpeed()
	return &BatchSetSpeedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchShowParams(request *model.BatchShowParamsRequest) (*model.BatchShowParamsResponse, error) {
	requestDef := GenReqDefForBatchShowParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowParamsResponse), nil
	}
}

func (c *DrsClient) BatchShowParamsInvoker(request *model.BatchShowParamsRequest) *BatchShowParamsInvoker {
	requestDef := GenReqDefForBatchShowParams()
	return &BatchShowParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchStartJobs(request *model.BatchStartJobsRequest) (*model.BatchStartJobsResponse, error) {
	requestDef := GenReqDefForBatchStartJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartJobsResponse), nil
	}
}

func (c *DrsClient) BatchStartJobsInvoker(request *model.BatchStartJobsRequest) *BatchStartJobsInvoker {
	requestDef := GenReqDefForBatchStartJobs()
	return &BatchStartJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchStopJobs(request *model.BatchStopJobsRequest) (*model.BatchStopJobsResponse, error) {
	requestDef := GenReqDefForBatchStopJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopJobsResponse), nil
	}
}

func (c *DrsClient) BatchStopJobsInvoker(request *model.BatchStopJobsRequest) *BatchStopJobsInvoker {
	requestDef := GenReqDefForBatchStopJobs()
	return &BatchStopJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchSwitchover(request *model.BatchSwitchoverRequest) (*model.BatchSwitchoverResponse, error) {
	requestDef := GenReqDefForBatchSwitchover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSwitchoverResponse), nil
	}
}

func (c *DrsClient) BatchSwitchoverInvoker(request *model.BatchSwitchoverRequest) *BatchSwitchoverInvoker {
	requestDef := GenReqDefForBatchSwitchover()
	return &BatchSwitchoverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchUpdateJob(request *model.BatchUpdateJobRequest) (*model.BatchUpdateJobResponse, error) {
	requestDef := GenReqDefForBatchUpdateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateJobResponse), nil
	}
}

func (c *DrsClient) BatchUpdateJobInvoker(request *model.BatchUpdateJobRequest) *BatchUpdateJobInvoker {
	requestDef := GenReqDefForBatchUpdateJob()
	return &BatchUpdateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchUpdateUser(request *model.BatchUpdateUserRequest) (*model.BatchUpdateUserResponse, error) {
	requestDef := GenReqDefForBatchUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateUserResponse), nil
	}
}

func (c *DrsClient) BatchUpdateUserInvoker(request *model.BatchUpdateUserRequest) *BatchUpdateUserInvoker {
	requestDef := GenReqDefForBatchUpdateUser()
	return &BatchUpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchValidateClustersConnections(request *model.BatchValidateClustersConnectionsRequest) (*model.BatchValidateClustersConnectionsResponse, error) {
	requestDef := GenReqDefForBatchValidateClustersConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchValidateClustersConnectionsResponse), nil
	}
}

func (c *DrsClient) BatchValidateClustersConnectionsInvoker(request *model.BatchValidateClustersConnectionsRequest) *BatchValidateClustersConnectionsInvoker {
	requestDef := GenReqDefForBatchValidateClustersConnections()
	return &BatchValidateClustersConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) BatchValidateConnections(request *model.BatchValidateConnectionsRequest) (*model.BatchValidateConnectionsResponse, error) {
	requestDef := GenReqDefForBatchValidateConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchValidateConnectionsResponse), nil
	}
}

func (c *DrsClient) BatchValidateConnectionsInvoker(request *model.BatchValidateConnectionsRequest) *BatchValidateConnectionsInvoker {
	requestDef := GenReqDefForBatchValidateConnections()
	return &BatchValidateConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) CreateCompareTask(request *model.CreateCompareTaskRequest) (*model.CreateCompareTaskResponse, error) {
	requestDef := GenReqDefForCreateCompareTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCompareTaskResponse), nil
	}
}

func (c *DrsClient) CreateCompareTaskInvoker(request *model.CreateCompareTaskRequest) *CreateCompareTaskInvoker {
	requestDef := GenReqDefForCreateCompareTask()
	return &CreateCompareTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) ListAvailableZone(request *model.ListAvailableZoneRequest) (*model.ListAvailableZoneResponse, error) {
	requestDef := GenReqDefForListAvailableZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZoneResponse), nil
	}
}

func (c *DrsClient) ListAvailableZoneInvoker(request *model.ListAvailableZoneRequest) *ListAvailableZoneInvoker {
	requestDef := GenReqDefForListAvailableZone()
	return &ListAvailableZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) ListCompareResult(request *model.ListCompareResultRequest) (*model.ListCompareResultResponse, error) {
	requestDef := GenReqDefForListCompareResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCompareResultResponse), nil
	}
}

func (c *DrsClient) ListCompareResultInvoker(request *model.ListCompareResultRequest) *ListCompareResultInvoker {
	requestDef := GenReqDefForListCompareResult()
	return &ListCompareResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

func (c *DrsClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) ShowJobList(request *model.ShowJobListRequest) (*model.ShowJobListResponse, error) {
	requestDef := GenReqDefForShowJobList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobListResponse), nil
	}
}

func (c *DrsClient) ShowJobListInvoker(request *model.ShowJobListRequest) *ShowJobListInvoker {
	requestDef := GenReqDefForShowJobList()
	return &ShowJobListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) ShowMonitoringData(request *model.ShowMonitoringDataRequest) (*model.ShowMonitoringDataResponse, error) {
	requestDef := GenReqDefForShowMonitoringData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonitoringDataResponse), nil
	}
}

func (c *DrsClient) ShowMonitoringDataInvoker(request *model.ShowMonitoringDataRequest) *ShowMonitoringDataInvoker {
	requestDef := GenReqDefForShowMonitoringData()
	return &ShowMonitoringDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

func (c *DrsClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) UpdateParams(request *model.UpdateParamsRequest) (*model.UpdateParamsResponse, error) {
	requestDef := GenReqDefForUpdateParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateParamsResponse), nil
	}
}

func (c *DrsClient) UpdateParamsInvoker(request *model.UpdateParamsRequest) *UpdateParamsInvoker {
	requestDef := GenReqDefForUpdateParams()
	return &UpdateParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DrsClient) UpdateTuningParams(request *model.UpdateTuningParamsRequest) (*model.UpdateTuningParamsResponse, error) {
	requestDef := GenReqDefForUpdateTuningParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTuningParamsResponse), nil
	}
}

func (c *DrsClient) UpdateTuningParamsInvoker(request *model.UpdateTuningParamsRequest) *UpdateTuningParamsInvoker {
	requestDef := GenReqDefForUpdateTuningParams()
	return &UpdateTuningParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
