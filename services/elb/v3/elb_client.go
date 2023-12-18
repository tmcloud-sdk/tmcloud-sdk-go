package v3

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/elb/v3/model"
)

type ElbClient struct {
	HcClient *http_client.HcHttpClient
}

func NewElbClient(hcClient *http_client.HcHttpClient) *ElbClient {
	return &ElbClient{HcClient: hcClient}
}

func ElbClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *ElbClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

func (c *ElbClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

func (c *ElbClient) ListAvailabilityZonesInvoker(request *model.ListAvailabilityZonesRequest) *ListAvailabilityZonesInvoker {
	requestDef := GenReqDefForListAvailabilityZones()
	return &ListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) BatchCreateMembers(request *model.BatchCreateMembersRequest) (*model.BatchCreateMembersResponse, error) {
	requestDef := GenReqDefForBatchCreateMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateMembersResponse), nil
	}
}

func (c *ElbClient) BatchCreateMembersInvoker(request *model.BatchCreateMembersRequest) *BatchCreateMembersInvoker {
	requestDef := GenReqDefForBatchCreateMembers()
	return &BatchCreateMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) BatchDeleteMembers(request *model.BatchDeleteMembersRequest) (*model.BatchDeleteMembersResponse, error) {
	requestDef := GenReqDefForBatchDeleteMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMembersResponse), nil
	}
}

func (c *ElbClient) BatchDeleteMembersInvoker(request *model.BatchDeleteMembersRequest) *BatchDeleteMembersInvoker {
	requestDef := GenReqDefForBatchDeleteMembers()
	return &BatchDeleteMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateMember(request *model.CreateMemberRequest) (*model.CreateMemberResponse, error) {
	requestDef := GenReqDefForCreateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberResponse), nil
	}
}

func (c *ElbClient) CreateMemberInvoker(request *model.CreateMemberRequest) *CreateMemberInvoker {
	requestDef := GenReqDefForCreateMember()
	return &CreateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

func (c *ElbClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListAllMembers(request *model.ListAllMembersRequest) (*model.ListAllMembersResponse, error) {
	requestDef := GenReqDefForListAllMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllMembersResponse), nil
	}
}

func (c *ElbClient) ListAllMembersInvoker(request *model.ListAllMembersRequest) *ListAllMembersInvoker {
	requestDef := GenReqDefForListAllMembers()
	return &ListAllMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

func (c *ElbClient) ListMembersInvoker(request *model.ListMembersRequest) *ListMembersInvoker {
	requestDef := GenReqDefForListMembers()
	return &ListMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowMember(request *model.ShowMemberRequest) (*model.ShowMemberResponse, error) {
	requestDef := GenReqDefForShowMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberResponse), nil
	}
}

func (c *ElbClient) ShowMemberInvoker(request *model.ShowMemberRequest) *ShowMemberInvoker {
	requestDef := GenReqDefForShowMember()
	return &ShowMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberResponse), nil
	}
}

func (c *ElbClient) UpdateMemberInvoker(request *model.UpdateMemberRequest) *UpdateMemberInvoker {
	requestDef := GenReqDefForUpdateMember()
	return &UpdateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreatePool(request *model.CreatePoolRequest) (*model.CreatePoolResponse, error) {
	requestDef := GenReqDefForCreatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePoolResponse), nil
	}
}

func (c *ElbClient) CreatePoolInvoker(request *model.CreatePoolRequest) *CreatePoolInvoker {
	requestDef := GenReqDefForCreatePool()
	return &CreatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
	requestDef := GenReqDefForDeletePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePoolResponse), nil
	}
}

func (c *ElbClient) DeletePoolInvoker(request *model.DeletePoolRequest) *DeletePoolInvoker {
	requestDef := GenReqDefForDeletePool()
	return &DeletePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListPools(request *model.ListPoolsRequest) (*model.ListPoolsResponse, error) {
	requestDef := GenReqDefForListPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolsResponse), nil
	}
}

func (c *ElbClient) ListPoolsInvoker(request *model.ListPoolsRequest) *ListPoolsInvoker {
	requestDef := GenReqDefForListPools()
	return &ListPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowPool(request *model.ShowPoolRequest) (*model.ShowPoolResponse, error) {
	requestDef := GenReqDefForShowPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolResponse), nil
	}
}

func (c *ElbClient) ShowPoolInvoker(request *model.ShowPoolRequest) *ShowPoolInvoker {
	requestDef := GenReqDefForShowPool()
	return &ShowPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdatePool(request *model.UpdatePoolRequest) (*model.UpdatePoolResponse, error) {
	requestDef := GenReqDefForUpdatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePoolResponse), nil
	}
}

func (c *ElbClient) UpdatePoolInvoker(request *model.UpdatePoolRequest) *UpdatePoolInvoker {
	requestDef := GenReqDefForUpdatePool()
	return &UpdatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

func (c *ElbClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

func (c *ElbClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

func (c *ElbClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

func (c *ElbClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

func (c *ElbClient) UpdateCertificateInvoker(request *model.UpdateCertificateRequest) *UpdateCertificateInvoker {
	requestDef := GenReqDefForUpdateCertificate()
	return &UpdateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) BatchUpdatePoliciesPriority(request *model.BatchUpdatePoliciesPriorityRequest) (*model.BatchUpdatePoliciesPriorityResponse, error) {
	requestDef := GenReqDefForBatchUpdatePoliciesPriority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdatePoliciesPriorityResponse), nil
	}
}

func (c *ElbClient) BatchUpdatePoliciesPriorityInvoker(request *model.BatchUpdatePoliciesPriorityRequest) *BatchUpdatePoliciesPriorityInvoker {
	requestDef := GenReqDefForBatchUpdatePoliciesPriority()
	return &BatchUpdatePoliciesPriorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateL7Policy(request *model.CreateL7PolicyRequest) (*model.CreateL7PolicyResponse, error) {
	requestDef := GenReqDefForCreateL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7PolicyResponse), nil
	}
}

func (c *ElbClient) CreateL7PolicyInvoker(request *model.CreateL7PolicyRequest) *CreateL7PolicyInvoker {
	requestDef := GenReqDefForCreateL7Policy()
	return &CreateL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteL7Policy(request *model.DeleteL7PolicyRequest) (*model.DeleteL7PolicyResponse, error) {
	requestDef := GenReqDefForDeleteL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7PolicyResponse), nil
	}
}

func (c *ElbClient) DeleteL7PolicyInvoker(request *model.DeleteL7PolicyRequest) *DeleteL7PolicyInvoker {
	requestDef := GenReqDefForDeleteL7Policy()
	return &DeleteL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListL7Policies(request *model.ListL7PoliciesRequest) (*model.ListL7PoliciesResponse, error) {
	requestDef := GenReqDefForListL7Policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7PoliciesResponse), nil
	}
}

func (c *ElbClient) ListL7PoliciesInvoker(request *model.ListL7PoliciesRequest) *ListL7PoliciesInvoker {
	requestDef := GenReqDefForListL7Policies()
	return &ListL7PoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowL7Policy(request *model.ShowL7PolicyRequest) (*model.ShowL7PolicyResponse, error) {
	requestDef := GenReqDefForShowL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7PolicyResponse), nil
	}
}

func (c *ElbClient) ShowL7PolicyInvoker(request *model.ShowL7PolicyRequest) *ShowL7PolicyInvoker {
	requestDef := GenReqDefForShowL7Policy()
	return &ShowL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateL7Policy(request *model.UpdateL7PolicyRequest) (*model.UpdateL7PolicyResponse, error) {
	requestDef := GenReqDefForUpdateL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7PolicyResponse), nil
	}
}

func (c *ElbClient) UpdateL7PolicyInvoker(request *model.UpdateL7PolicyRequest) *UpdateL7PolicyInvoker {
	requestDef := GenReqDefForUpdateL7Policy()
	return &UpdateL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateL7Rule(request *model.CreateL7RuleRequest) (*model.CreateL7RuleResponse, error) {
	requestDef := GenReqDefForCreateL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7RuleResponse), nil
	}
}

func (c *ElbClient) CreateL7RuleInvoker(request *model.CreateL7RuleRequest) *CreateL7RuleInvoker {
	requestDef := GenReqDefForCreateL7Rule()
	return &CreateL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteL7Rule(request *model.DeleteL7RuleRequest) (*model.DeleteL7RuleResponse, error) {
	requestDef := GenReqDefForDeleteL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7RuleResponse), nil
	}
}

func (c *ElbClient) DeleteL7RuleInvoker(request *model.DeleteL7RuleRequest) *DeleteL7RuleInvoker {
	requestDef := GenReqDefForDeleteL7Rule()
	return &DeleteL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListL7Rules(request *model.ListL7RulesRequest) (*model.ListL7RulesResponse, error) {
	requestDef := GenReqDefForListL7Rules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7RulesResponse), nil
	}
}

func (c *ElbClient) ListL7RulesInvoker(request *model.ListL7RulesRequest) *ListL7RulesInvoker {
	requestDef := GenReqDefForListL7Rules()
	return &ListL7RulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowL7Rule(request *model.ShowL7RuleRequest) (*model.ShowL7RuleResponse, error) {
	requestDef := GenReqDefForShowL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7RuleResponse), nil
	}
}

func (c *ElbClient) ShowL7RuleInvoker(request *model.ShowL7RuleRequest) *ShowL7RuleInvoker {
	requestDef := GenReqDefForShowL7Rule()
	return &ShowL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateL7Rule(request *model.UpdateL7RuleRequest) (*model.UpdateL7RuleResponse, error) {
	requestDef := GenReqDefForUpdateL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7RuleResponse), nil
	}
}

func (c *ElbClient) UpdateL7RuleInvoker(request *model.UpdateL7RuleRequest) *UpdateL7RuleInvoker {
	requestDef := GenReqDefForUpdateL7Rule()
	return &UpdateL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateHealthMonitor(request *model.CreateHealthMonitorRequest) (*model.CreateHealthMonitorResponse, error) {
	requestDef := GenReqDefForCreateHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHealthMonitorResponse), nil
	}
}

func (c *ElbClient) CreateHealthMonitorInvoker(request *model.CreateHealthMonitorRequest) *CreateHealthMonitorInvoker {
	requestDef := GenReqDefForCreateHealthMonitor()
	return &CreateHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteHealthMonitor(request *model.DeleteHealthMonitorRequest) (*model.DeleteHealthMonitorResponse, error) {
	requestDef := GenReqDefForDeleteHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHealthMonitorResponse), nil
	}
}

func (c *ElbClient) DeleteHealthMonitorInvoker(request *model.DeleteHealthMonitorRequest) *DeleteHealthMonitorInvoker {
	requestDef := GenReqDefForDeleteHealthMonitor()
	return &DeleteHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListHealthMonitors(request *model.ListHealthMonitorsRequest) (*model.ListHealthMonitorsResponse, error) {
	requestDef := GenReqDefForListHealthMonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHealthMonitorsResponse), nil
	}
}

func (c *ElbClient) ListHealthMonitorsInvoker(request *model.ListHealthMonitorsRequest) *ListHealthMonitorsInvoker {
	requestDef := GenReqDefForListHealthMonitors()
	return &ListHealthMonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowHealthMonitor(request *model.ShowHealthMonitorRequest) (*model.ShowHealthMonitorResponse, error) {
	requestDef := GenReqDefForShowHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthMonitorResponse), nil
	}
}

func (c *ElbClient) ShowHealthMonitorInvoker(request *model.ShowHealthMonitorRequest) *ShowHealthMonitorInvoker {
	requestDef := GenReqDefForShowHealthMonitor()
	return &ShowHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateHealthMonitor(request *model.UpdateHealthMonitorRequest) (*model.UpdateHealthMonitorResponse, error) {
	requestDef := GenReqDefForUpdateHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthMonitorResponse), nil
	}
}

func (c *ElbClient) UpdateHealthMonitorInvoker(request *model.UpdateHealthMonitorRequest) *UpdateHealthMonitorInvoker {
	requestDef := GenReqDefForUpdateHealthMonitor()
	return &UpdateHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) BatchDeleteIpList(request *model.BatchDeleteIpListRequest) (*model.BatchDeleteIpListResponse, error) {
	requestDef := GenReqDefForBatchDeleteIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteIpListResponse), nil
	}
}

func (c *ElbClient) BatchDeleteIpListInvoker(request *model.BatchDeleteIpListRequest) *BatchDeleteIpListInvoker {
	requestDef := GenReqDefForBatchDeleteIpList()
	return &BatchDeleteIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateIpGroup(request *model.CreateIpGroupRequest) (*model.CreateIpGroupResponse, error) {
	requestDef := GenReqDefForCreateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpGroupResponse), nil
	}
}

func (c *ElbClient) CreateIpGroupInvoker(request *model.CreateIpGroupRequest) *CreateIpGroupInvoker {
	requestDef := GenReqDefForCreateIpGroup()
	return &CreateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteIpGroup(request *model.DeleteIpGroupRequest) (*model.DeleteIpGroupResponse, error) {
	requestDef := GenReqDefForDeleteIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpGroupResponse), nil
	}
}

func (c *ElbClient) DeleteIpGroupInvoker(request *model.DeleteIpGroupRequest) *DeleteIpGroupInvoker {
	requestDef := GenReqDefForDeleteIpGroup()
	return &DeleteIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListIpGroups(request *model.ListIpGroupsRequest) (*model.ListIpGroupsResponse, error) {
	requestDef := GenReqDefForListIpGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpGroupsResponse), nil
	}
}

func (c *ElbClient) ListIpGroupsInvoker(request *model.ListIpGroupsRequest) *ListIpGroupsInvoker {
	requestDef := GenReqDefForListIpGroups()
	return &ListIpGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowIpGroup(request *model.ShowIpGroupRequest) (*model.ShowIpGroupResponse, error) {
	requestDef := GenReqDefForShowIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpGroupResponse), nil
	}
}

func (c *ElbClient) ShowIpGroupInvoker(request *model.ShowIpGroupRequest) *ShowIpGroupInvoker {
	requestDef := GenReqDefForShowIpGroup()
	return &ShowIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateIpGroup(request *model.UpdateIpGroupRequest) (*model.UpdateIpGroupResponse, error) {
	requestDef := GenReqDefForUpdateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpGroupResponse), nil
	}
}

func (c *ElbClient) UpdateIpGroupInvoker(request *model.UpdateIpGroupRequest) *UpdateIpGroupInvoker {
	requestDef := GenReqDefForUpdateIpGroup()
	return &UpdateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateIpList(request *model.UpdateIpListRequest) (*model.UpdateIpListResponse, error) {
	requestDef := GenReqDefForUpdateIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpListResponse), nil
	}
}

func (c *ElbClient) UpdateIpListInvoker(request *model.UpdateIpListRequest) *UpdateIpListInvoker {
	requestDef := GenReqDefForUpdateIpList()
	return &UpdateIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateListener(request *model.CreateListenerRequest) (*model.CreateListenerResponse, error) {
	requestDef := GenReqDefForCreateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateListenerResponse), nil
	}
}

func (c *ElbClient) CreateListenerInvoker(request *model.CreateListenerRequest) *CreateListenerInvoker {
	requestDef := GenReqDefForCreateListener()
	return &CreateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteListener(request *model.DeleteListenerRequest) (*model.DeleteListenerResponse, error) {
	requestDef := GenReqDefForDeleteListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerResponse), nil
	}
}

func (c *ElbClient) DeleteListenerInvoker(request *model.DeleteListenerRequest) *DeleteListenerInvoker {
	requestDef := GenReqDefForDeleteListener()
	return &DeleteListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListListeners(request *model.ListListenersRequest) (*model.ListListenersResponse, error) {
	requestDef := GenReqDefForListListeners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersResponse), nil
	}
}

func (c *ElbClient) ListListenersInvoker(request *model.ListListenersRequest) *ListListenersInvoker {
	requestDef := GenReqDefForListListeners()
	return &ListListenersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowListener(request *model.ShowListenerRequest) (*model.ShowListenerResponse, error) {
	requestDef := GenReqDefForShowListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerResponse), nil
	}
}

func (c *ElbClient) ShowListenerInvoker(request *model.ShowListenerRequest) *ShowListenerInvoker {
	requestDef := GenReqDefForShowListener()
	return &ShowListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
	requestDef := GenReqDefForUpdateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateListenerResponse), nil
	}
}

func (c *ElbClient) UpdateListenerInvoker(request *model.UpdateListenerRequest) *UpdateListenerInvoker {
	requestDef := GenReqDefForUpdateListener()
	return &UpdateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ChangeLoadbalancerChargeMode(request *model.ChangeLoadbalancerChargeModeRequest) (*model.ChangeLoadbalancerChargeModeResponse, error) {
	requestDef := GenReqDefForChangeLoadbalancerChargeMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeLoadbalancerChargeModeResponse), nil
	}
}

func (c *ElbClient) ChangeLoadbalancerChargeModeInvoker(request *model.ChangeLoadbalancerChargeModeRequest) *ChangeLoadbalancerChargeModeInvoker {
	requestDef := GenReqDefForChangeLoadbalancerChargeMode()
	return &ChangeLoadbalancerChargeModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateLoadBalancer(request *model.CreateLoadBalancerRequest) (*model.CreateLoadBalancerResponse, error) {
	requestDef := GenReqDefForCreateLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadBalancerResponse), nil
	}
}

func (c *ElbClient) CreateLoadBalancerInvoker(request *model.CreateLoadBalancerRequest) *CreateLoadBalancerInvoker {
	requestDef := GenReqDefForCreateLoadBalancer()
	return &CreateLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteLoadBalancer(request *model.DeleteLoadBalancerRequest) (*model.DeleteLoadBalancerResponse, error) {
	requestDef := GenReqDefForDeleteLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadBalancerResponse), nil
	}
}

func (c *ElbClient) DeleteLoadBalancerInvoker(request *model.DeleteLoadBalancerRequest) *DeleteLoadBalancerInvoker {
	requestDef := GenReqDefForDeleteLoadBalancer()
	return &DeleteLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListLoadBalancers(request *model.ListLoadBalancersRequest) (*model.ListLoadBalancersResponse, error) {
	requestDef := GenReqDefForListLoadBalancers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadBalancersResponse), nil
	}
}

func (c *ElbClient) ListLoadBalancersInvoker(request *model.ListLoadBalancersRequest) *ListLoadBalancersInvoker {
	requestDef := GenReqDefForListLoadBalancers()
	return &ListLoadBalancersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowLoadBalancer(request *model.ShowLoadBalancerRequest) (*model.ShowLoadBalancerResponse, error) {
	requestDef := GenReqDefForShowLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerResponse), nil
	}
}

func (c *ElbClient) ShowLoadBalancerInvoker(request *model.ShowLoadBalancerRequest) *ShowLoadBalancerInvoker {
	requestDef := GenReqDefForShowLoadBalancer()
	return &ShowLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowLoadBalancerStatus(request *model.ShowLoadBalancerStatusRequest) (*model.ShowLoadBalancerStatusResponse, error) {
	requestDef := GenReqDefForShowLoadBalancerStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerStatusResponse), nil
	}
}

func (c *ElbClient) ShowLoadBalancerStatusInvoker(request *model.ShowLoadBalancerStatusRequest) *ShowLoadBalancerStatusInvoker {
	requestDef := GenReqDefForShowLoadBalancerStatus()
	return &ShowLoadBalancerStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateLoadBalancer(request *model.UpdateLoadBalancerRequest) (*model.UpdateLoadBalancerResponse, error) {
	requestDef := GenReqDefForUpdateLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLoadBalancerResponse), nil
	}
}

func (c *ElbClient) UpdateLoadBalancerInvoker(request *model.UpdateLoadBalancerRequest) *UpdateLoadBalancerInvoker {
	requestDef := GenReqDefForUpdateLoadBalancer()
	return &UpdateLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

func (c *ElbClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowFlavor(request *model.ShowFlavorRequest) (*model.ShowFlavorResponse, error) {
	requestDef := GenReqDefForShowFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlavorResponse), nil
	}
}

func (c *ElbClient) ShowFlavorInvoker(request *model.ShowFlavorRequest) *ShowFlavorInvoker {
	requestDef := GenReqDefForShowFlavor()
	return &ShowFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateLogtank(request *model.CreateLogtankRequest) (*model.CreateLogtankResponse, error) {
	requestDef := GenReqDefForCreateLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogtankResponse), nil
	}
}

func (c *ElbClient) CreateLogtankInvoker(request *model.CreateLogtankRequest) *CreateLogtankInvoker {
	requestDef := GenReqDefForCreateLogtank()
	return &CreateLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteLogtank(request *model.DeleteLogtankRequest) (*model.DeleteLogtankResponse, error) {
	requestDef := GenReqDefForDeleteLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogtankResponse), nil
	}
}

func (c *ElbClient) DeleteLogtankInvoker(request *model.DeleteLogtankRequest) *DeleteLogtankInvoker {
	requestDef := GenReqDefForDeleteLogtank()
	return &DeleteLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListLogtanks(request *model.ListLogtanksRequest) (*model.ListLogtanksResponse, error) {
	requestDef := GenReqDefForListLogtanks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogtanksResponse), nil
	}
}

func (c *ElbClient) ListLogtanksInvoker(request *model.ListLogtanksRequest) *ListLogtanksInvoker {
	requestDef := GenReqDefForListLogtanks()
	return &ListLogtanksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowLogtank(request *model.ShowLogtankRequest) (*model.ShowLogtankResponse, error) {
	requestDef := GenReqDefForShowLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogtankResponse), nil
	}
}

func (c *ElbClient) ShowLogtankInvoker(request *model.ShowLogtankRequest) *ShowLogtankInvoker {
	requestDef := GenReqDefForShowLogtank()
	return &ShowLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateLogtank(request *model.UpdateLogtankRequest) (*model.UpdateLogtankResponse, error) {
	requestDef := GenReqDefForUpdateLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogtankResponse), nil
	}
}

func (c *ElbClient) UpdateLogtankInvoker(request *model.UpdateLogtankRequest) *UpdateLogtankInvoker {
	requestDef := GenReqDefForUpdateLogtank()
	return &UpdateLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListQuotaDetails(request *model.ListQuotaDetailsRequest) (*model.ListQuotaDetailsResponse, error) {
	requestDef := GenReqDefForListQuotaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaDetailsResponse), nil
	}
}

func (c *ElbClient) ListQuotaDetailsInvoker(request *model.ListQuotaDetailsRequest) *ListQuotaDetailsInvoker {
	requestDef := GenReqDefForListQuotaDetails()
	return &ListQuotaDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

func (c *ElbClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CountPreoccupyIpNum(request *model.CountPreoccupyIpNumRequest) (*model.CountPreoccupyIpNumResponse, error) {
	requestDef := GenReqDefForCountPreoccupyIpNum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountPreoccupyIpNumResponse), nil
	}
}

func (c *ElbClient) CountPreoccupyIpNumInvoker(request *model.CountPreoccupyIpNumRequest) *CountPreoccupyIpNumInvoker {
	requestDef := GenReqDefForCountPreoccupyIpNum()
	return &CountPreoccupyIpNumInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateSecurityPolicy(request *model.CreateSecurityPolicyRequest) (*model.CreateSecurityPolicyResponse, error) {
	requestDef := GenReqDefForCreateSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityPolicyResponse), nil
	}
}

func (c *ElbClient) CreateSecurityPolicyInvoker(request *model.CreateSecurityPolicyRequest) *CreateSecurityPolicyInvoker {
	requestDef := GenReqDefForCreateSecurityPolicy()
	return &CreateSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteSecurityPolicy(request *model.DeleteSecurityPolicyRequest) (*model.DeleteSecurityPolicyResponse, error) {
	requestDef := GenReqDefForDeleteSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityPolicyResponse), nil
	}
}

func (c *ElbClient) DeleteSecurityPolicyInvoker(request *model.DeleteSecurityPolicyRequest) *DeleteSecurityPolicyInvoker {
	requestDef := GenReqDefForDeleteSecurityPolicy()
	return &DeleteSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListSecurityPolicies(request *model.ListSecurityPoliciesRequest) (*model.ListSecurityPoliciesResponse, error) {
	requestDef := GenReqDefForListSecurityPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityPoliciesResponse), nil
	}
}

func (c *ElbClient) ListSecurityPoliciesInvoker(request *model.ListSecurityPoliciesRequest) *ListSecurityPoliciesInvoker {
	requestDef := GenReqDefForListSecurityPolicies()
	return &ListSecurityPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListSystemSecurityPolicies(request *model.ListSystemSecurityPoliciesRequest) (*model.ListSystemSecurityPoliciesResponse, error) {
	requestDef := GenReqDefForListSystemSecurityPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSystemSecurityPoliciesResponse), nil
	}
}

func (c *ElbClient) ListSystemSecurityPoliciesInvoker(request *model.ListSystemSecurityPoliciesRequest) *ListSystemSecurityPoliciesInvoker {
	requestDef := GenReqDefForListSystemSecurityPolicies()
	return &ListSystemSecurityPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowSecurityPolicy(request *model.ShowSecurityPolicyRequest) (*model.ShowSecurityPolicyResponse, error) {
	requestDef := GenReqDefForShowSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityPolicyResponse), nil
	}
}

func (c *ElbClient) ShowSecurityPolicyInvoker(request *model.ShowSecurityPolicyRequest) *ShowSecurityPolicyInvoker {
	requestDef := GenReqDefForShowSecurityPolicy()
	return &ShowSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateSecurityPolicy(request *model.UpdateSecurityPolicyRequest) (*model.UpdateSecurityPolicyResponse, error) {
	requestDef := GenReqDefForUpdateSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityPolicyResponse), nil
	}
}

func (c *ElbClient) UpdateSecurityPolicyInvoker(request *model.UpdateSecurityPolicyRequest) *UpdateSecurityPolicyInvoker {
	requestDef := GenReqDefForUpdateSecurityPolicy()
	return &UpdateSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
