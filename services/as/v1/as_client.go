package v1

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/as/v1/model"
)

type AsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewAsClient(hcClient *http_client.HcHttpClient) *AsClient {
	return &AsClient{HcClient: hcClient}
}

func AsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *AsClient) AttachCallbackInstanceLifeCycleHook(request *model.AttachCallbackInstanceLifeCycleHookRequest) (*model.AttachCallbackInstanceLifeCycleHookResponse, error) {
	requestDef := GenReqDefForAttachCallbackInstanceLifeCycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachCallbackInstanceLifeCycleHookResponse), nil
	}
}

func (c *AsClient) AttachCallbackInstanceLifeCycleHookInvoker(request *model.AttachCallbackInstanceLifeCycleHookRequest) *AttachCallbackInstanceLifeCycleHookInvoker {
	requestDef := GenReqDefForAttachCallbackInstanceLifeCycleHook()
	return &AttachCallbackInstanceLifeCycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchAddScalingInstances(request *model.BatchAddScalingInstancesRequest) (*model.BatchAddScalingInstancesResponse, error) {
	requestDef := GenReqDefForBatchAddScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddScalingInstancesResponse), nil
	}
}

func (c *AsClient) BatchAddScalingInstancesInvoker(request *model.BatchAddScalingInstancesRequest) *BatchAddScalingInstancesInvoker {
	requestDef := GenReqDefForBatchAddScalingInstances()
	return &BatchAddScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchDeleteScalingConfigs(request *model.BatchDeleteScalingConfigsRequest) (*model.BatchDeleteScalingConfigsResponse, error) {
	requestDef := GenReqDefForBatchDeleteScalingConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteScalingConfigsResponse), nil
	}
}

func (c *AsClient) BatchDeleteScalingConfigsInvoker(request *model.BatchDeleteScalingConfigsRequest) *BatchDeleteScalingConfigsInvoker {
	requestDef := GenReqDefForBatchDeleteScalingConfigs()
	return &BatchDeleteScalingConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchDeleteScalingPolicies(request *model.BatchDeleteScalingPoliciesRequest) (*model.BatchDeleteScalingPoliciesResponse, error) {
	requestDef := GenReqDefForBatchDeleteScalingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteScalingPoliciesResponse), nil
	}
}

func (c *AsClient) BatchDeleteScalingPoliciesInvoker(request *model.BatchDeleteScalingPoliciesRequest) *BatchDeleteScalingPoliciesInvoker {
	requestDef := GenReqDefForBatchDeleteScalingPolicies()
	return &BatchDeleteScalingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchPauseScalingPolicies(request *model.BatchPauseScalingPoliciesRequest) (*model.BatchPauseScalingPoliciesResponse, error) {
	requestDef := GenReqDefForBatchPauseScalingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchPauseScalingPoliciesResponse), nil
	}
}

func (c *AsClient) BatchPauseScalingPoliciesInvoker(request *model.BatchPauseScalingPoliciesRequest) *BatchPauseScalingPoliciesInvoker {
	requestDef := GenReqDefForBatchPauseScalingPolicies()
	return &BatchPauseScalingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchProtectScalingInstances(request *model.BatchProtectScalingInstancesRequest) (*model.BatchProtectScalingInstancesResponse, error) {
	requestDef := GenReqDefForBatchProtectScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchProtectScalingInstancesResponse), nil
	}
}

func (c *AsClient) BatchProtectScalingInstancesInvoker(request *model.BatchProtectScalingInstancesRequest) *BatchProtectScalingInstancesInvoker {
	requestDef := GenReqDefForBatchProtectScalingInstances()
	return &BatchProtectScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchRemoveScalingInstances(request *model.BatchRemoveScalingInstancesRequest) (*model.BatchRemoveScalingInstancesResponse, error) {
	requestDef := GenReqDefForBatchRemoveScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveScalingInstancesResponse), nil
	}
}

func (c *AsClient) BatchRemoveScalingInstancesInvoker(request *model.BatchRemoveScalingInstancesRequest) *BatchRemoveScalingInstancesInvoker {
	requestDef := GenReqDefForBatchRemoveScalingInstances()
	return &BatchRemoveScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchResumeScalingPolicies(request *model.BatchResumeScalingPoliciesRequest) (*model.BatchResumeScalingPoliciesResponse, error) {
	requestDef := GenReqDefForBatchResumeScalingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResumeScalingPoliciesResponse), nil
	}
}

func (c *AsClient) BatchResumeScalingPoliciesInvoker(request *model.BatchResumeScalingPoliciesRequest) *BatchResumeScalingPoliciesInvoker {
	requestDef := GenReqDefForBatchResumeScalingPolicies()
	return &BatchResumeScalingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchSetScalingInstancesStandby(request *model.BatchSetScalingInstancesStandbyRequest) (*model.BatchSetScalingInstancesStandbyResponse, error) {
	requestDef := GenReqDefForBatchSetScalingInstancesStandby()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetScalingInstancesStandbyResponse), nil
	}
}

func (c *AsClient) BatchSetScalingInstancesStandbyInvoker(request *model.BatchSetScalingInstancesStandbyRequest) *BatchSetScalingInstancesStandbyInvoker {
	requestDef := GenReqDefForBatchSetScalingInstancesStandby()
	return &BatchSetScalingInstancesStandbyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchUnprotectScalingInstances(request *model.BatchUnprotectScalingInstancesRequest) (*model.BatchUnprotectScalingInstancesResponse, error) {
	requestDef := GenReqDefForBatchUnprotectScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUnprotectScalingInstancesResponse), nil
	}
}

func (c *AsClient) BatchUnprotectScalingInstancesInvoker(request *model.BatchUnprotectScalingInstancesRequest) *BatchUnprotectScalingInstancesInvoker {
	requestDef := GenReqDefForBatchUnprotectScalingInstances()
	return &BatchUnprotectScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) BatchUnsetScalingInstancesStantby(request *model.BatchUnsetScalingInstancesStantbyRequest) (*model.BatchUnsetScalingInstancesStantbyResponse, error) {
	requestDef := GenReqDefForBatchUnsetScalingInstancesStantby()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUnsetScalingInstancesStantbyResponse), nil
	}
}

func (c *AsClient) BatchUnsetScalingInstancesStantbyInvoker(request *model.BatchUnsetScalingInstancesStantbyRequest) *BatchUnsetScalingInstancesStantbyInvoker {
	requestDef := GenReqDefForBatchUnsetScalingInstancesStantby()
	return &BatchUnsetScalingInstancesStantbyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) CreateLifyCycleHook(request *model.CreateLifyCycleHookRequest) (*model.CreateLifyCycleHookResponse, error) {
	requestDef := GenReqDefForCreateLifyCycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLifyCycleHookResponse), nil
	}
}

func (c *AsClient) CreateLifyCycleHookInvoker(request *model.CreateLifyCycleHookRequest) *CreateLifyCycleHookInvoker {
	requestDef := GenReqDefForCreateLifyCycleHook()
	return &CreateLifyCycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) CreateScalingConfig(request *model.CreateScalingConfigRequest) (*model.CreateScalingConfigResponse, error) {
	requestDef := GenReqDefForCreateScalingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingConfigResponse), nil
	}
}

func (c *AsClient) CreateScalingConfigInvoker(request *model.CreateScalingConfigRequest) *CreateScalingConfigInvoker {
	requestDef := GenReqDefForCreateScalingConfig()
	return &CreateScalingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) CreateScalingGroup(request *model.CreateScalingGroupRequest) (*model.CreateScalingGroupResponse, error) {
	requestDef := GenReqDefForCreateScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingGroupResponse), nil
	}
}

func (c *AsClient) CreateScalingGroupInvoker(request *model.CreateScalingGroupRequest) *CreateScalingGroupInvoker {
	requestDef := GenReqDefForCreateScalingGroup()
	return &CreateScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) CreateScalingNotification(request *model.CreateScalingNotificationRequest) (*model.CreateScalingNotificationResponse, error) {
	requestDef := GenReqDefForCreateScalingNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingNotificationResponse), nil
	}
}

func (c *AsClient) CreateScalingNotificationInvoker(request *model.CreateScalingNotificationRequest) *CreateScalingNotificationInvoker {
	requestDef := GenReqDefForCreateScalingNotification()
	return &CreateScalingNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) CreateScalingPolicy(request *model.CreateScalingPolicyRequest) (*model.CreateScalingPolicyResponse, error) {
	requestDef := GenReqDefForCreateScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingPolicyResponse), nil
	}
}

func (c *AsClient) CreateScalingPolicyInvoker(request *model.CreateScalingPolicyRequest) *CreateScalingPolicyInvoker {
	requestDef := GenReqDefForCreateScalingPolicy()
	return &CreateScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) CreateScalingTagInfo(request *model.CreateScalingTagInfoRequest) (*model.CreateScalingTagInfoResponse, error) {
	requestDef := GenReqDefForCreateScalingTagInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingTagInfoResponse), nil
	}
}

func (c *AsClient) CreateScalingTagInfoInvoker(request *model.CreateScalingTagInfoRequest) *CreateScalingTagInfoInvoker {
	requestDef := GenReqDefForCreateScalingTagInfo()
	return &CreateScalingTagInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) DeleteLifecycleHook(request *model.DeleteLifecycleHookRequest) (*model.DeleteLifecycleHookResponse, error) {
	requestDef := GenReqDefForDeleteLifecycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLifecycleHookResponse), nil
	}
}

func (c *AsClient) DeleteLifecycleHookInvoker(request *model.DeleteLifecycleHookRequest) *DeleteLifecycleHookInvoker {
	requestDef := GenReqDefForDeleteLifecycleHook()
	return &DeleteLifecycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) DeleteScalingConfig(request *model.DeleteScalingConfigRequest) (*model.DeleteScalingConfigResponse, error) {
	requestDef := GenReqDefForDeleteScalingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingConfigResponse), nil
	}
}

func (c *AsClient) DeleteScalingConfigInvoker(request *model.DeleteScalingConfigRequest) *DeleteScalingConfigInvoker {
	requestDef := GenReqDefForDeleteScalingConfig()
	return &DeleteScalingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) DeleteScalingGroup(request *model.DeleteScalingGroupRequest) (*model.DeleteScalingGroupResponse, error) {
	requestDef := GenReqDefForDeleteScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingGroupResponse), nil
	}
}

func (c *AsClient) DeleteScalingGroupInvoker(request *model.DeleteScalingGroupRequest) *DeleteScalingGroupInvoker {
	requestDef := GenReqDefForDeleteScalingGroup()
	return &DeleteScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) DeleteScalingInstance(request *model.DeleteScalingInstanceRequest) (*model.DeleteScalingInstanceResponse, error) {
	requestDef := GenReqDefForDeleteScalingInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingInstanceResponse), nil
	}
}

func (c *AsClient) DeleteScalingInstanceInvoker(request *model.DeleteScalingInstanceRequest) *DeleteScalingInstanceInvoker {
	requestDef := GenReqDefForDeleteScalingInstance()
	return &DeleteScalingInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) DeleteScalingNotification(request *model.DeleteScalingNotificationRequest) (*model.DeleteScalingNotificationResponse, error) {
	requestDef := GenReqDefForDeleteScalingNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingNotificationResponse), nil
	}
}

func (c *AsClient) DeleteScalingNotificationInvoker(request *model.DeleteScalingNotificationRequest) *DeleteScalingNotificationInvoker {
	requestDef := GenReqDefForDeleteScalingNotification()
	return &DeleteScalingNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) DeleteScalingPolicy(request *model.DeleteScalingPolicyRequest) (*model.DeleteScalingPolicyResponse, error) {
	requestDef := GenReqDefForDeleteScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingPolicyResponse), nil
	}
}

func (c *AsClient) DeleteScalingPolicyInvoker(request *model.DeleteScalingPolicyRequest) *DeleteScalingPolicyInvoker {
	requestDef := GenReqDefForDeleteScalingPolicy()
	return &DeleteScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) DeleteScalingTagInfo(request *model.DeleteScalingTagInfoRequest) (*model.DeleteScalingTagInfoResponse, error) {
	requestDef := GenReqDefForDeleteScalingTagInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingTagInfoResponse), nil
	}
}

func (c *AsClient) DeleteScalingTagInfoInvoker(request *model.DeleteScalingTagInfoRequest) *DeleteScalingTagInfoInvoker {
	requestDef := GenReqDefForDeleteScalingTagInfo()
	return &DeleteScalingTagInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ExecuteScalingPolicy(request *model.ExecuteScalingPolicyRequest) (*model.ExecuteScalingPolicyResponse, error) {
	requestDef := GenReqDefForExecuteScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteScalingPolicyResponse), nil
	}
}

func (c *AsClient) ExecuteScalingPolicyInvoker(request *model.ExecuteScalingPolicyRequest) *ExecuteScalingPolicyInvoker {
	requestDef := GenReqDefForExecuteScalingPolicy()
	return &ExecuteScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListHookInstances(request *model.ListHookInstancesRequest) (*model.ListHookInstancesResponse, error) {
	requestDef := GenReqDefForListHookInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHookInstancesResponse), nil
	}
}

func (c *AsClient) ListHookInstancesInvoker(request *model.ListHookInstancesRequest) *ListHookInstancesInvoker {
	requestDef := GenReqDefForListHookInstances()
	return &ListHookInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListLifeCycleHooks(request *model.ListLifeCycleHooksRequest) (*model.ListLifeCycleHooksResponse, error) {
	requestDef := GenReqDefForListLifeCycleHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLifeCycleHooksResponse), nil
	}
}

func (c *AsClient) ListLifeCycleHooksInvoker(request *model.ListLifeCycleHooksRequest) *ListLifeCycleHooksInvoker {
	requestDef := GenReqDefForListLifeCycleHooks()
	return &ListLifeCycleHooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

func (c *AsClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingActivityLogs(request *model.ListScalingActivityLogsRequest) (*model.ListScalingActivityLogsResponse, error) {
	requestDef := GenReqDefForListScalingActivityLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingActivityLogsResponse), nil
	}
}

func (c *AsClient) ListScalingActivityLogsInvoker(request *model.ListScalingActivityLogsRequest) *ListScalingActivityLogsInvoker {
	requestDef := GenReqDefForListScalingActivityLogs()
	return &ListScalingActivityLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingActivityV2Logs(request *model.ListScalingActivityV2LogsRequest) (*model.ListScalingActivityV2LogsResponse, error) {
	requestDef := GenReqDefForListScalingActivityV2Logs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingActivityV2LogsResponse), nil
	}
}

func (c *AsClient) ListScalingActivityV2LogsInvoker(request *model.ListScalingActivityV2LogsRequest) *ListScalingActivityV2LogsInvoker {
	requestDef := GenReqDefForListScalingActivityV2Logs()
	return &ListScalingActivityV2LogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingConfigs(request *model.ListScalingConfigsRequest) (*model.ListScalingConfigsResponse, error) {
	requestDef := GenReqDefForListScalingConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingConfigsResponse), nil
	}
}

func (c *AsClient) ListScalingConfigsInvoker(request *model.ListScalingConfigsRequest) *ListScalingConfigsInvoker {
	requestDef := GenReqDefForListScalingConfigs()
	return &ListScalingConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingGroups(request *model.ListScalingGroupsRequest) (*model.ListScalingGroupsResponse, error) {
	requestDef := GenReqDefForListScalingGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingGroupsResponse), nil
	}
}

func (c *AsClient) ListScalingGroupsInvoker(request *model.ListScalingGroupsRequest) *ListScalingGroupsInvoker {
	requestDef := GenReqDefForListScalingGroups()
	return &ListScalingGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingInstances(request *model.ListScalingInstancesRequest) (*model.ListScalingInstancesResponse, error) {
	requestDef := GenReqDefForListScalingInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingInstancesResponse), nil
	}
}

func (c *AsClient) ListScalingInstancesInvoker(request *model.ListScalingInstancesRequest) *ListScalingInstancesInvoker {
	requestDef := GenReqDefForListScalingInstances()
	return &ListScalingInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingNotifications(request *model.ListScalingNotificationsRequest) (*model.ListScalingNotificationsResponse, error) {
	requestDef := GenReqDefForListScalingNotifications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingNotificationsResponse), nil
	}
}

func (c *AsClient) ListScalingNotificationsInvoker(request *model.ListScalingNotificationsRequest) *ListScalingNotificationsInvoker {
	requestDef := GenReqDefForListScalingNotifications()
	return &ListScalingNotificationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingPolicies(request *model.ListScalingPoliciesRequest) (*model.ListScalingPoliciesResponse, error) {
	requestDef := GenReqDefForListScalingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingPoliciesResponse), nil
	}
}

func (c *AsClient) ListScalingPoliciesInvoker(request *model.ListScalingPoliciesRequest) *ListScalingPoliciesInvoker {
	requestDef := GenReqDefForListScalingPolicies()
	return &ListScalingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingPolicyExecuteLogs(request *model.ListScalingPolicyExecuteLogsRequest) (*model.ListScalingPolicyExecuteLogsResponse, error) {
	requestDef := GenReqDefForListScalingPolicyExecuteLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingPolicyExecuteLogsResponse), nil
	}
}

func (c *AsClient) ListScalingPolicyExecuteLogsInvoker(request *model.ListScalingPolicyExecuteLogsRequest) *ListScalingPolicyExecuteLogsInvoker {
	requestDef := GenReqDefForListScalingPolicyExecuteLogs()
	return &ListScalingPolicyExecuteLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingTagInfosByResourceId(request *model.ListScalingTagInfosByResourceIdRequest) (*model.ListScalingTagInfosByResourceIdResponse, error) {
	requestDef := GenReqDefForListScalingTagInfosByResourceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingTagInfosByResourceIdResponse), nil
	}
}

func (c *AsClient) ListScalingTagInfosByResourceIdInvoker(request *model.ListScalingTagInfosByResourceIdRequest) *ListScalingTagInfosByResourceIdInvoker {
	requestDef := GenReqDefForListScalingTagInfosByResourceId()
	return &ListScalingTagInfosByResourceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingTagInfosByTenantId(request *model.ListScalingTagInfosByTenantIdRequest) (*model.ListScalingTagInfosByTenantIdResponse, error) {
	requestDef := GenReqDefForListScalingTagInfosByTenantId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingTagInfosByTenantIdResponse), nil
	}
}

func (c *AsClient) ListScalingTagInfosByTenantIdInvoker(request *model.ListScalingTagInfosByTenantIdRequest) *ListScalingTagInfosByTenantIdInvoker {
	requestDef := GenReqDefForListScalingTagInfosByTenantId()
	return &ListScalingTagInfosByTenantIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) PauseScalingGroup(request *model.PauseScalingGroupRequest) (*model.PauseScalingGroupResponse, error) {
	requestDef := GenReqDefForPauseScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseScalingGroupResponse), nil
	}
}

func (c *AsClient) PauseScalingGroupInvoker(request *model.PauseScalingGroupRequest) *PauseScalingGroupInvoker {
	requestDef := GenReqDefForPauseScalingGroup()
	return &PauseScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) PauseScalingPolicy(request *model.PauseScalingPolicyRequest) (*model.PauseScalingPolicyResponse, error) {
	requestDef := GenReqDefForPauseScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseScalingPolicyResponse), nil
	}
}

func (c *AsClient) PauseScalingPolicyInvoker(request *model.PauseScalingPolicyRequest) *PauseScalingPolicyInvoker {
	requestDef := GenReqDefForPauseScalingPolicy()
	return &PauseScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ResumeScalingGroup(request *model.ResumeScalingGroupRequest) (*model.ResumeScalingGroupResponse, error) {
	requestDef := GenReqDefForResumeScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeScalingGroupResponse), nil
	}
}

func (c *AsClient) ResumeScalingGroupInvoker(request *model.ResumeScalingGroupRequest) *ResumeScalingGroupInvoker {
	requestDef := GenReqDefForResumeScalingGroup()
	return &ResumeScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ResumeScalingPolicy(request *model.ResumeScalingPolicyRequest) (*model.ResumeScalingPolicyResponse, error) {
	requestDef := GenReqDefForResumeScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeScalingPolicyResponse), nil
	}
}

func (c *AsClient) ResumeScalingPolicyInvoker(request *model.ResumeScalingPolicyRequest) *ResumeScalingPolicyInvoker {
	requestDef := GenReqDefForResumeScalingPolicy()
	return &ResumeScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ShowLifeCycleHook(request *model.ShowLifeCycleHookRequest) (*model.ShowLifeCycleHookResponse, error) {
	requestDef := GenReqDefForShowLifeCycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLifeCycleHookResponse), nil
	}
}

func (c *AsClient) ShowLifeCycleHookInvoker(request *model.ShowLifeCycleHookRequest) *ShowLifeCycleHookInvoker {
	requestDef := GenReqDefForShowLifeCycleHook()
	return &ShowLifeCycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ShowPolicyAndInstanceQuota(request *model.ShowPolicyAndInstanceQuotaRequest) (*model.ShowPolicyAndInstanceQuotaResponse, error) {
	requestDef := GenReqDefForShowPolicyAndInstanceQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyAndInstanceQuotaResponse), nil
	}
}

func (c *AsClient) ShowPolicyAndInstanceQuotaInvoker(request *model.ShowPolicyAndInstanceQuotaRequest) *ShowPolicyAndInstanceQuotaInvoker {
	requestDef := GenReqDefForShowPolicyAndInstanceQuota()
	return &ShowPolicyAndInstanceQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ShowResourceQuota(request *model.ShowResourceQuotaRequest) (*model.ShowResourceQuotaResponse, error) {
	requestDef := GenReqDefForShowResourceQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceQuotaResponse), nil
	}
}

func (c *AsClient) ShowResourceQuotaInvoker(request *model.ShowResourceQuotaRequest) *ShowResourceQuotaInvoker {
	requestDef := GenReqDefForShowResourceQuota()
	return &ShowResourceQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ShowScalingConfig(request *model.ShowScalingConfigRequest) (*model.ShowScalingConfigResponse, error) {
	requestDef := GenReqDefForShowScalingConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingConfigResponse), nil
	}
}

func (c *AsClient) ShowScalingConfigInvoker(request *model.ShowScalingConfigRequest) *ShowScalingConfigInvoker {
	requestDef := GenReqDefForShowScalingConfig()
	return &ShowScalingConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ShowScalingGroup(request *model.ShowScalingGroupRequest) (*model.ShowScalingGroupResponse, error) {
	requestDef := GenReqDefForShowScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingGroupResponse), nil
	}
}

func (c *AsClient) ShowScalingGroupInvoker(request *model.ShowScalingGroupRequest) *ShowScalingGroupInvoker {
	requestDef := GenReqDefForShowScalingGroup()
	return &ShowScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ShowScalingPolicy(request *model.ShowScalingPolicyRequest) (*model.ShowScalingPolicyResponse, error) {
	requestDef := GenReqDefForShowScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingPolicyResponse), nil
	}
}

func (c *AsClient) ShowScalingPolicyInvoker(request *model.ShowScalingPolicyRequest) *ShowScalingPolicyInvoker {
	requestDef := GenReqDefForShowScalingPolicy()
	return &ShowScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) UpdateLifeCycleHook(request *model.UpdateLifeCycleHookRequest) (*model.UpdateLifeCycleHookResponse, error) {
	requestDef := GenReqDefForUpdateLifeCycleHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLifeCycleHookResponse), nil
	}
}

func (c *AsClient) UpdateLifeCycleHookInvoker(request *model.UpdateLifeCycleHookRequest) *UpdateLifeCycleHookInvoker {
	requestDef := GenReqDefForUpdateLifeCycleHook()
	return &UpdateLifeCycleHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) UpdateScalingGroup(request *model.UpdateScalingGroupRequest) (*model.UpdateScalingGroupResponse, error) {
	requestDef := GenReqDefForUpdateScalingGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScalingGroupResponse), nil
	}
}

func (c *AsClient) UpdateScalingGroupInvoker(request *model.UpdateScalingGroupRequest) *UpdateScalingGroupInvoker {
	requestDef := GenReqDefForUpdateScalingGroup()
	return &UpdateScalingGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) UpdateScalingPolicy(request *model.UpdateScalingPolicyRequest) (*model.UpdateScalingPolicyResponse, error) {
	requestDef := GenReqDefForUpdateScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScalingPolicyResponse), nil
	}
}

func (c *AsClient) UpdateScalingPolicyInvoker(request *model.UpdateScalingPolicyRequest) *UpdateScalingPolicyInvoker {
	requestDef := GenReqDefForUpdateScalingPolicy()
	return &UpdateScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

func (c *AsClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

func (c *AsClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) CreateScalingV2Policy(request *model.CreateScalingV2PolicyRequest) (*model.CreateScalingV2PolicyResponse, error) {
	requestDef := GenReqDefForCreateScalingV2Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingV2PolicyResponse), nil
	}
}

func (c *AsClient) CreateScalingV2PolicyInvoker(request *model.CreateScalingV2PolicyRequest) *CreateScalingV2PolicyInvoker {
	requestDef := GenReqDefForCreateScalingV2Policy()
	return &CreateScalingV2PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListAllScalingV2Policies(request *model.ListAllScalingV2PoliciesRequest) (*model.ListAllScalingV2PoliciesResponse, error) {
	requestDef := GenReqDefForListAllScalingV2Policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllScalingV2PoliciesResponse), nil
	}
}

func (c *AsClient) ListAllScalingV2PoliciesInvoker(request *model.ListAllScalingV2PoliciesRequest) *ListAllScalingV2PoliciesInvoker {
	requestDef := GenReqDefForListAllScalingV2Policies()
	return &ListAllScalingV2PoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ListScalingV2Policies(request *model.ListScalingV2PoliciesRequest) (*model.ListScalingV2PoliciesResponse, error) {
	requestDef := GenReqDefForListScalingV2Policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScalingV2PoliciesResponse), nil
	}
}

func (c *AsClient) ListScalingV2PoliciesInvoker(request *model.ListScalingV2PoliciesRequest) *ListScalingV2PoliciesInvoker {
	requestDef := GenReqDefForListScalingV2Policies()
	return &ListScalingV2PoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) ShowScalingV2Policy(request *model.ShowScalingV2PolicyRequest) (*model.ShowScalingV2PolicyResponse, error) {
	requestDef := GenReqDefForShowScalingV2Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingV2PolicyResponse), nil
	}
}

func (c *AsClient) ShowScalingV2PolicyInvoker(request *model.ShowScalingV2PolicyRequest) *ShowScalingV2PolicyInvoker {
	requestDef := GenReqDefForShowScalingV2Policy()
	return &ShowScalingV2PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *AsClient) UpdateScalingV2Policy(request *model.UpdateScalingV2PolicyRequest) (*model.UpdateScalingV2PolicyResponse, error) {
	requestDef := GenReqDefForUpdateScalingV2Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScalingV2PolicyResponse), nil
	}
}

func (c *AsClient) UpdateScalingV2PolicyInvoker(request *model.UpdateScalingV2PolicyRequest) *UpdateScalingV2PolicyInvoker {
	requestDef := GenReqDefForUpdateScalingV2Policy()
	return &UpdateScalingV2PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
