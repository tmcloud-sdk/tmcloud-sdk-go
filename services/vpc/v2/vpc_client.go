package v2

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/vpc/v2/model"
)

type VpcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVpcClient(hcClient *http_client.HcHttpClient) *VpcClient {
	return &VpcClient{HcClient: hcClient}
}

func VpcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *VpcClient) AcceptVpcPeering(request *model.AcceptVpcPeeringRequest) (*model.AcceptVpcPeeringResponse, error) {
	requestDef := GenReqDefForAcceptVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptVpcPeeringResponse), nil
	}
}

func (c *VpcClient) AcceptVpcPeeringInvoker(request *model.AcceptVpcPeeringRequest) *AcceptVpcPeeringInvoker {
	requestDef := GenReqDefForAcceptVpcPeering()
	return &AcceptVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) AssociateRouteTable(request *model.AssociateRouteTableRequest) (*model.AssociateRouteTableResponse, error) {
	requestDef := GenReqDefForAssociateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRouteTableResponse), nil
	}
}

func (c *VpcClient) AssociateRouteTableInvoker(request *model.AssociateRouteTableRequest) *AssociateRouteTableInvoker {
	requestDef := GenReqDefForAssociateRouteTable()
	return &AssociateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) BatchCreateSubnetTags(request *model.BatchCreateSubnetTagsRequest) (*model.BatchCreateSubnetTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateSubnetTagsResponse), nil
	}
}

func (c *VpcClient) BatchCreateSubnetTagsInvoker(request *model.BatchCreateSubnetTagsRequest) *BatchCreateSubnetTagsInvoker {
	requestDef := GenReqDefForBatchCreateSubnetTags()
	return &BatchCreateSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) BatchDeleteSubnetTags(request *model.BatchDeleteSubnetTagsRequest) (*model.BatchDeleteSubnetTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSubnetTagsResponse), nil
	}
}

func (c *VpcClient) BatchDeleteSubnetTagsInvoker(request *model.BatchDeleteSubnetTagsRequest) *BatchDeleteSubnetTagsInvoker {
	requestDef := GenReqDefForBatchDeleteSubnetTags()
	return &BatchDeleteSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreatePort(request *model.CreatePortRequest) (*model.CreatePortResponse, error) {
	requestDef := GenReqDefForCreatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePortResponse), nil
	}
}

func (c *VpcClient) CreatePortInvoker(request *model.CreatePortRequest) *CreatePortInvoker {
	requestDef := GenReqDefForCreatePort()
	return &CreatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateRouteTable(request *model.CreateRouteTableRequest) (*model.CreateRouteTableResponse, error) {
	requestDef := GenReqDefForCreateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRouteTableResponse), nil
	}
}

func (c *VpcClient) CreateRouteTableInvoker(request *model.CreateRouteTableRequest) *CreateRouteTableInvoker {
	requestDef := GenReqDefForCreateRouteTable()
	return &CreateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateSecurityGroup(request *model.CreateSecurityGroupRequest) (*model.CreateSecurityGroupResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupResponse), nil
	}
}

func (c *VpcClient) CreateSecurityGroupInvoker(request *model.CreateSecurityGroupRequest) *CreateSecurityGroupInvoker {
	requestDef := GenReqDefForCreateSecurityGroup()
	return &CreateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateSecurityGroupRule(request *model.CreateSecurityGroupRuleRequest) (*model.CreateSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) CreateSecurityGroupRuleInvoker(request *model.CreateSecurityGroupRuleRequest) *CreateSecurityGroupRuleInvoker {
	requestDef := GenReqDefForCreateSecurityGroupRule()
	return &CreateSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateSubnet(request *model.CreateSubnetRequest) (*model.CreateSubnetResponse, error) {
	requestDef := GenReqDefForCreateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubnetResponse), nil
	}
}

func (c *VpcClient) CreateSubnetInvoker(request *model.CreateSubnetRequest) *CreateSubnetInvoker {
	requestDef := GenReqDefForCreateSubnet()
	return &CreateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateSubnetTag(request *model.CreateSubnetTagRequest) (*model.CreateSubnetTagResponse, error) {
	requestDef := GenReqDefForCreateSubnetTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubnetTagResponse), nil
	}
}

func (c *VpcClient) CreateSubnetTagInvoker(request *model.CreateSubnetTagRequest) *CreateSubnetTagInvoker {
	requestDef := GenReqDefForCreateSubnetTag()
	return &CreateSubnetTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateVpcPeering(request *model.CreateVpcPeeringRequest) (*model.CreateVpcPeeringResponse, error) {
	requestDef := GenReqDefForCreateVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcPeeringResponse), nil
	}
}

func (c *VpcClient) CreateVpcPeeringInvoker(request *model.CreateVpcPeeringRequest) *CreateVpcPeeringInvoker {
	requestDef := GenReqDefForCreateVpcPeering()
	return &CreateVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeletePort(request *model.DeletePortRequest) (*model.DeletePortResponse, error) {
	requestDef := GenReqDefForDeletePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePortResponse), nil
	}
}

func (c *VpcClient) DeletePortInvoker(request *model.DeletePortRequest) *DeletePortInvoker {
	requestDef := GenReqDefForDeletePort()
	return &DeletePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteRouteTable(request *model.DeleteRouteTableRequest) (*model.DeleteRouteTableResponse, error) {
	requestDef := GenReqDefForDeleteRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRouteTableResponse), nil
	}
}

func (c *VpcClient) DeleteRouteTableInvoker(request *model.DeleteRouteTableRequest) *DeleteRouteTableInvoker {
	requestDef := GenReqDefForDeleteRouteTable()
	return &DeleteRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteSecurityGroup(request *model.DeleteSecurityGroupRequest) (*model.DeleteSecurityGroupResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupResponse), nil
	}
}

func (c *VpcClient) DeleteSecurityGroupInvoker(request *model.DeleteSecurityGroupRequest) *DeleteSecurityGroupInvoker {
	requestDef := GenReqDefForDeleteSecurityGroup()
	return &DeleteSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteSecurityGroupRule(request *model.DeleteSecurityGroupRuleRequest) (*model.DeleteSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) DeleteSecurityGroupRuleInvoker(request *model.DeleteSecurityGroupRuleRequest) *DeleteSecurityGroupRuleInvoker {
	requestDef := GenReqDefForDeleteSecurityGroupRule()
	return &DeleteSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteSubnet(request *model.DeleteSubnetRequest) (*model.DeleteSubnetResponse, error) {
	requestDef := GenReqDefForDeleteSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubnetResponse), nil
	}
}

func (c *VpcClient) DeleteSubnetInvoker(request *model.DeleteSubnetRequest) *DeleteSubnetInvoker {
	requestDef := GenReqDefForDeleteSubnet()
	return &DeleteSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteSubnetTag(request *model.DeleteSubnetTagRequest) (*model.DeleteSubnetTagResponse, error) {
	requestDef := GenReqDefForDeleteSubnetTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubnetTagResponse), nil
	}
}

func (c *VpcClient) DeleteSubnetTagInvoker(request *model.DeleteSubnetTagRequest) *DeleteSubnetTagInvoker {
	requestDef := GenReqDefForDeleteSubnetTag()
	return &DeleteSubnetTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteVpcPeering(request *model.DeleteVpcPeeringRequest) (*model.DeleteVpcPeeringResponse, error) {
	requestDef := GenReqDefForDeleteVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcPeeringResponse), nil
	}
}

func (c *VpcClient) DeleteVpcPeeringInvoker(request *model.DeleteVpcPeeringRequest) *DeleteVpcPeeringInvoker {
	requestDef := GenReqDefForDeleteVpcPeering()
	return &DeleteVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DisassociateRouteTable(request *model.DisassociateRouteTableRequest) (*model.DisassociateRouteTableResponse, error) {
	requestDef := GenReqDefForDisassociateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRouteTableResponse), nil
	}
}

func (c *VpcClient) DisassociateRouteTableInvoker(request *model.DisassociateRouteTableRequest) *DisassociateRouteTableInvoker {
	requestDef := GenReqDefForDisassociateRouteTable()
	return &DisassociateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListPorts(request *model.ListPortsRequest) (*model.ListPortsResponse, error) {
	requestDef := GenReqDefForListPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortsResponse), nil
	}
}

func (c *VpcClient) ListPortsInvoker(request *model.ListPortsRequest) *ListPortsInvoker {
	requestDef := GenReqDefForListPorts()
	return &ListPortsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListRouteTables(request *model.ListRouteTablesRequest) (*model.ListRouteTablesResponse, error) {
	requestDef := GenReqDefForListRouteTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRouteTablesResponse), nil
	}
}

func (c *VpcClient) ListRouteTablesInvoker(request *model.ListRouteTablesRequest) *ListRouteTablesInvoker {
	requestDef := GenReqDefForListRouteTables()
	return &ListRouteTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListSecurityGroupRules(request *model.ListSecurityGroupRulesRequest) (*model.ListSecurityGroupRulesResponse, error) {
	requestDef := GenReqDefForListSecurityGroupRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupRulesResponse), nil
	}
}

func (c *VpcClient) ListSecurityGroupRulesInvoker(request *model.ListSecurityGroupRulesRequest) *ListSecurityGroupRulesInvoker {
	requestDef := GenReqDefForListSecurityGroupRules()
	return &ListSecurityGroupRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListSecurityGroups(request *model.ListSecurityGroupsRequest) (*model.ListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForListSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupsResponse), nil
	}
}

func (c *VpcClient) ListSecurityGroupsInvoker(request *model.ListSecurityGroupsRequest) *ListSecurityGroupsInvoker {
	requestDef := GenReqDefForListSecurityGroups()
	return &ListSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListSubnetTags(request *model.ListSubnetTagsRequest) (*model.ListSubnetTagsResponse, error) {
	requestDef := GenReqDefForListSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetTagsResponse), nil
	}
}

func (c *VpcClient) ListSubnetTagsInvoker(request *model.ListSubnetTagsRequest) *ListSubnetTagsInvoker {
	requestDef := GenReqDefForListSubnetTags()
	return &ListSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListSubnets(request *model.ListSubnetsRequest) (*model.ListSubnetsResponse, error) {
	requestDef := GenReqDefForListSubnets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetsResponse), nil
	}
}

func (c *VpcClient) ListSubnetsInvoker(request *model.ListSubnetsRequest) *ListSubnetsInvoker {
	requestDef := GenReqDefForListSubnets()
	return &ListSubnetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListSubnetsByTags(request *model.ListSubnetsByTagsRequest) (*model.ListSubnetsByTagsResponse, error) {
	requestDef := GenReqDefForListSubnetsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetsByTagsResponse), nil
	}
}

func (c *VpcClient) ListSubnetsByTagsInvoker(request *model.ListSubnetsByTagsRequest) *ListSubnetsByTagsInvoker {
	requestDef := GenReqDefForListSubnetsByTags()
	return &ListSubnetsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListVpcPeerings(request *model.ListVpcPeeringsRequest) (*model.ListVpcPeeringsResponse, error) {
	requestDef := GenReqDefForListVpcPeerings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcPeeringsResponse), nil
	}
}

func (c *VpcClient) ListVpcPeeringsInvoker(request *model.ListVpcPeeringsRequest) *ListVpcPeeringsInvoker {
	requestDef := GenReqDefForListVpcPeerings()
	return &ListVpcPeeringsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) RejectVpcPeering(request *model.RejectVpcPeeringRequest) (*model.RejectVpcPeeringResponse, error) {
	requestDef := GenReqDefForRejectVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RejectVpcPeeringResponse), nil
	}
}

func (c *VpcClient) RejectVpcPeeringInvoker(request *model.RejectVpcPeeringRequest) *RejectVpcPeeringInvoker {
	requestDef := GenReqDefForRejectVpcPeering()
	return &RejectVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowPort(request *model.ShowPortRequest) (*model.ShowPortResponse, error) {
	requestDef := GenReqDefForShowPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPortResponse), nil
	}
}

func (c *VpcClient) ShowPortInvoker(request *model.ShowPortRequest) *ShowPortInvoker {
	requestDef := GenReqDefForShowPort()
	return &ShowPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

func (c *VpcClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowRouteTable(request *model.ShowRouteTableRequest) (*model.ShowRouteTableResponse, error) {
	requestDef := GenReqDefForShowRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRouteTableResponse), nil
	}
}

func (c *VpcClient) ShowRouteTableInvoker(request *model.ShowRouteTableRequest) *ShowRouteTableInvoker {
	requestDef := GenReqDefForShowRouteTable()
	return &ShowRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowSecurityGroup(request *model.ShowSecurityGroupRequest) (*model.ShowSecurityGroupResponse, error) {
	requestDef := GenReqDefForShowSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupResponse), nil
	}
}

func (c *VpcClient) ShowSecurityGroupInvoker(request *model.ShowSecurityGroupRequest) *ShowSecurityGroupInvoker {
	requestDef := GenReqDefForShowSecurityGroup()
	return &ShowSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowSecurityGroupRule(request *model.ShowSecurityGroupRuleRequest) (*model.ShowSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForShowSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) ShowSecurityGroupRuleInvoker(request *model.ShowSecurityGroupRuleRequest) *ShowSecurityGroupRuleInvoker {
	requestDef := GenReqDefForShowSecurityGroupRule()
	return &ShowSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowSubnet(request *model.ShowSubnetRequest) (*model.ShowSubnetResponse, error) {
	requestDef := GenReqDefForShowSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubnetResponse), nil
	}
}

func (c *VpcClient) ShowSubnetInvoker(request *model.ShowSubnetRequest) *ShowSubnetInvoker {
	requestDef := GenReqDefForShowSubnet()
	return &ShowSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowSubnetTags(request *model.ShowSubnetTagsRequest) (*model.ShowSubnetTagsResponse, error) {
	requestDef := GenReqDefForShowSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubnetTagsResponse), nil
	}
}

func (c *VpcClient) ShowSubnetTagsInvoker(request *model.ShowSubnetTagsRequest) *ShowSubnetTagsInvoker {
	requestDef := GenReqDefForShowSubnetTags()
	return &ShowSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowVpcPeering(request *model.ShowVpcPeeringRequest) (*model.ShowVpcPeeringResponse, error) {
	requestDef := GenReqDefForShowVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcPeeringResponse), nil
	}
}

func (c *VpcClient) ShowVpcPeeringInvoker(request *model.ShowVpcPeeringRequest) *ShowVpcPeeringInvoker {
	requestDef := GenReqDefForShowVpcPeering()
	return &ShowVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdatePort(request *model.UpdatePortRequest) (*model.UpdatePortResponse, error) {
	requestDef := GenReqDefForUpdatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePortResponse), nil
	}
}

func (c *VpcClient) UpdatePortInvoker(request *model.UpdatePortRequest) *UpdatePortInvoker {
	requestDef := GenReqDefForUpdatePort()
	return &UpdatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateRouteTable(request *model.UpdateRouteTableRequest) (*model.UpdateRouteTableResponse, error) {
	requestDef := GenReqDefForUpdateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRouteTableResponse), nil
	}
}

func (c *VpcClient) UpdateRouteTableInvoker(request *model.UpdateRouteTableRequest) *UpdateRouteTableInvoker {
	requestDef := GenReqDefForUpdateRouteTable()
	return &UpdateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateSubnet(request *model.UpdateSubnetRequest) (*model.UpdateSubnetResponse, error) {
	requestDef := GenReqDefForUpdateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubnetResponse), nil
	}
}

func (c *VpcClient) UpdateSubnetInvoker(request *model.UpdateSubnetRequest) *UpdateSubnetInvoker {
	requestDef := GenReqDefForUpdateSubnet()
	return &UpdateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateVpcPeering(request *model.UpdateVpcPeeringRequest) (*model.UpdateVpcPeeringResponse, error) {
	requestDef := GenReqDefForUpdateVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcPeeringResponse), nil
	}
}

func (c *VpcClient) UpdateVpcPeeringInvoker(request *model.UpdateVpcPeeringRequest) *UpdateVpcPeeringInvoker {
	requestDef := GenReqDefForUpdateVpcPeering()
	return &UpdateVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreatePrivateip(request *model.CreatePrivateipRequest) (*model.CreatePrivateipResponse, error) {
	requestDef := GenReqDefForCreatePrivateip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateipResponse), nil
	}
}

func (c *VpcClient) CreatePrivateipInvoker(request *model.CreatePrivateipRequest) *CreatePrivateipInvoker {
	requestDef := GenReqDefForCreatePrivateip()
	return &CreatePrivateipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeletePrivateip(request *model.DeletePrivateipRequest) (*model.DeletePrivateipResponse, error) {
	requestDef := GenReqDefForDeletePrivateip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateipResponse), nil
	}
}

func (c *VpcClient) DeletePrivateipInvoker(request *model.DeletePrivateipRequest) *DeletePrivateipInvoker {
	requestDef := GenReqDefForDeletePrivateip()
	return &DeletePrivateipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListPrivateips(request *model.ListPrivateipsRequest) (*model.ListPrivateipsResponse, error) {
	requestDef := GenReqDefForListPrivateips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateipsResponse), nil
	}
}

func (c *VpcClient) ListPrivateipsInvoker(request *model.ListPrivateipsRequest) *ListPrivateipsInvoker {
	requestDef := GenReqDefForListPrivateips()
	return &ListPrivateipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowNetworkIpAvailabilities(request *model.ShowNetworkIpAvailabilitiesRequest) (*model.ShowNetworkIpAvailabilitiesResponse, error) {
	requestDef := GenReqDefForShowNetworkIpAvailabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkIpAvailabilitiesResponse), nil
	}
}

func (c *VpcClient) ShowNetworkIpAvailabilitiesInvoker(request *model.ShowNetworkIpAvailabilitiesRequest) *ShowNetworkIpAvailabilitiesInvoker {
	requestDef := GenReqDefForShowNetworkIpAvailabilities()
	return &ShowNetworkIpAvailabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowPrivateip(request *model.ShowPrivateipRequest) (*model.ShowPrivateipResponse, error) {
	requestDef := GenReqDefForShowPrivateip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateipResponse), nil
	}
}

func (c *VpcClient) ShowPrivateipInvoker(request *model.ShowPrivateipRequest) *ShowPrivateipInvoker {
	requestDef := GenReqDefForShowPrivateip()
	return &ShowPrivateipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronCreateSecurityGroup(request *model.NeutronCreateSecurityGroupRequest) (*model.NeutronCreateSecurityGroupResponse, error) {
	requestDef := GenReqDefForNeutronCreateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateSecurityGroupResponse), nil
	}
}

func (c *VpcClient) NeutronCreateSecurityGroupInvoker(request *model.NeutronCreateSecurityGroupRequest) *NeutronCreateSecurityGroupInvoker {
	requestDef := GenReqDefForNeutronCreateSecurityGroup()
	return &NeutronCreateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronCreateSecurityGroupRule(request *model.NeutronCreateSecurityGroupRuleRequest) (*model.NeutronCreateSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForNeutronCreateSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) NeutronCreateSecurityGroupRuleInvoker(request *model.NeutronCreateSecurityGroupRuleRequest) *NeutronCreateSecurityGroupRuleInvoker {
	requestDef := GenReqDefForNeutronCreateSecurityGroupRule()
	return &NeutronCreateSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronDeleteSecurityGroup(request *model.NeutronDeleteSecurityGroupRequest) (*model.NeutronDeleteSecurityGroupResponse, error) {
	requestDef := GenReqDefForNeutronDeleteSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteSecurityGroupResponse), nil
	}
}

func (c *VpcClient) NeutronDeleteSecurityGroupInvoker(request *model.NeutronDeleteSecurityGroupRequest) *NeutronDeleteSecurityGroupInvoker {
	requestDef := GenReqDefForNeutronDeleteSecurityGroup()
	return &NeutronDeleteSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronDeleteSecurityGroupRule(request *model.NeutronDeleteSecurityGroupRuleRequest) (*model.NeutronDeleteSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForNeutronDeleteSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) NeutronDeleteSecurityGroupRuleInvoker(request *model.NeutronDeleteSecurityGroupRuleRequest) *NeutronDeleteSecurityGroupRuleInvoker {
	requestDef := GenReqDefForNeutronDeleteSecurityGroupRule()
	return &NeutronDeleteSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronListSecurityGroupRules(request *model.NeutronListSecurityGroupRulesRequest) (*model.NeutronListSecurityGroupRulesResponse, error) {
	requestDef := GenReqDefForNeutronListSecurityGroupRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListSecurityGroupRulesResponse), nil
	}
}

func (c *VpcClient) NeutronListSecurityGroupRulesInvoker(request *model.NeutronListSecurityGroupRulesRequest) *NeutronListSecurityGroupRulesInvoker {
	requestDef := GenReqDefForNeutronListSecurityGroupRules()
	return &NeutronListSecurityGroupRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronListSecurityGroups(request *model.NeutronListSecurityGroupsRequest) (*model.NeutronListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForNeutronListSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListSecurityGroupsResponse), nil
	}
}

func (c *VpcClient) NeutronListSecurityGroupsInvoker(request *model.NeutronListSecurityGroupsRequest) *NeutronListSecurityGroupsInvoker {
	requestDef := GenReqDefForNeutronListSecurityGroups()
	return &NeutronListSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronShowSecurityGroup(request *model.NeutronShowSecurityGroupRequest) (*model.NeutronShowSecurityGroupResponse, error) {
	requestDef := GenReqDefForNeutronShowSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowSecurityGroupResponse), nil
	}
}

func (c *VpcClient) NeutronShowSecurityGroupInvoker(request *model.NeutronShowSecurityGroupRequest) *NeutronShowSecurityGroupInvoker {
	requestDef := GenReqDefForNeutronShowSecurityGroup()
	return &NeutronShowSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronShowSecurityGroupRule(request *model.NeutronShowSecurityGroupRuleRequest) (*model.NeutronShowSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForNeutronShowSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) NeutronShowSecurityGroupRuleInvoker(request *model.NeutronShowSecurityGroupRuleRequest) *NeutronShowSecurityGroupRuleInvoker {
	requestDef := GenReqDefForNeutronShowSecurityGroupRule()
	return &NeutronShowSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronUpdateSecurityGroup(request *model.NeutronUpdateSecurityGroupRequest) (*model.NeutronUpdateSecurityGroupResponse, error) {
	requestDef := GenReqDefForNeutronUpdateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateSecurityGroupResponse), nil
	}
}

func (c *VpcClient) NeutronUpdateSecurityGroupInvoker(request *model.NeutronUpdateSecurityGroupRequest) *NeutronUpdateSecurityGroupInvoker {
	requestDef := GenReqDefForNeutronUpdateSecurityGroup()
	return &NeutronUpdateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronAddFirewallRule(request *model.NeutronAddFirewallRuleRequest) (*model.NeutronAddFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronAddFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronAddFirewallRuleResponse), nil
	}
}

func (c *VpcClient) NeutronAddFirewallRuleInvoker(request *model.NeutronAddFirewallRuleRequest) *NeutronAddFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronAddFirewallRule()
	return &NeutronAddFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronCreateFirewallGroup(request *model.NeutronCreateFirewallGroupRequest) (*model.NeutronCreateFirewallGroupResponse, error) {
	requestDef := GenReqDefForNeutronCreateFirewallGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateFirewallGroupResponse), nil
	}
}

func (c *VpcClient) NeutronCreateFirewallGroupInvoker(request *model.NeutronCreateFirewallGroupRequest) *NeutronCreateFirewallGroupInvoker {
	requestDef := GenReqDefForNeutronCreateFirewallGroup()
	return &NeutronCreateFirewallGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronCreateFirewallPolicy(request *model.NeutronCreateFirewallPolicyRequest) (*model.NeutronCreateFirewallPolicyResponse, error) {
	requestDef := GenReqDefForNeutronCreateFirewallPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateFirewallPolicyResponse), nil
	}
}

func (c *VpcClient) NeutronCreateFirewallPolicyInvoker(request *model.NeutronCreateFirewallPolicyRequest) *NeutronCreateFirewallPolicyInvoker {
	requestDef := GenReqDefForNeutronCreateFirewallPolicy()
	return &NeutronCreateFirewallPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronCreateFirewallRule(request *model.NeutronCreateFirewallRuleRequest) (*model.NeutronCreateFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronCreateFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateFirewallRuleResponse), nil
	}
}

func (c *VpcClient) NeutronCreateFirewallRuleInvoker(request *model.NeutronCreateFirewallRuleRequest) *NeutronCreateFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronCreateFirewallRule()
	return &NeutronCreateFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronDeleteFirewallGroup(request *model.NeutronDeleteFirewallGroupRequest) (*model.NeutronDeleteFirewallGroupResponse, error) {
	requestDef := GenReqDefForNeutronDeleteFirewallGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteFirewallGroupResponse), nil
	}
}

func (c *VpcClient) NeutronDeleteFirewallGroupInvoker(request *model.NeutronDeleteFirewallGroupRequest) *NeutronDeleteFirewallGroupInvoker {
	requestDef := GenReqDefForNeutronDeleteFirewallGroup()
	return &NeutronDeleteFirewallGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronDeleteFirewallPolicy(request *model.NeutronDeleteFirewallPolicyRequest) (*model.NeutronDeleteFirewallPolicyResponse, error) {
	requestDef := GenReqDefForNeutronDeleteFirewallPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteFirewallPolicyResponse), nil
	}
}

func (c *VpcClient) NeutronDeleteFirewallPolicyInvoker(request *model.NeutronDeleteFirewallPolicyRequest) *NeutronDeleteFirewallPolicyInvoker {
	requestDef := GenReqDefForNeutronDeleteFirewallPolicy()
	return &NeutronDeleteFirewallPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronDeleteFirewallRule(request *model.NeutronDeleteFirewallRuleRequest) (*model.NeutronDeleteFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronDeleteFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteFirewallRuleResponse), nil
	}
}

func (c *VpcClient) NeutronDeleteFirewallRuleInvoker(request *model.NeutronDeleteFirewallRuleRequest) *NeutronDeleteFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronDeleteFirewallRule()
	return &NeutronDeleteFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronListFirewallGroups(request *model.NeutronListFirewallGroupsRequest) (*model.NeutronListFirewallGroupsResponse, error) {
	requestDef := GenReqDefForNeutronListFirewallGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListFirewallGroupsResponse), nil
	}
}

func (c *VpcClient) NeutronListFirewallGroupsInvoker(request *model.NeutronListFirewallGroupsRequest) *NeutronListFirewallGroupsInvoker {
	requestDef := GenReqDefForNeutronListFirewallGroups()
	return &NeutronListFirewallGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronListFirewallPolicies(request *model.NeutronListFirewallPoliciesRequest) (*model.NeutronListFirewallPoliciesResponse, error) {
	requestDef := GenReqDefForNeutronListFirewallPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListFirewallPoliciesResponse), nil
	}
}

func (c *VpcClient) NeutronListFirewallPoliciesInvoker(request *model.NeutronListFirewallPoliciesRequest) *NeutronListFirewallPoliciesInvoker {
	requestDef := GenReqDefForNeutronListFirewallPolicies()
	return &NeutronListFirewallPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronListFirewallRules(request *model.NeutronListFirewallRulesRequest) (*model.NeutronListFirewallRulesResponse, error) {
	requestDef := GenReqDefForNeutronListFirewallRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListFirewallRulesResponse), nil
	}
}

func (c *VpcClient) NeutronListFirewallRulesInvoker(request *model.NeutronListFirewallRulesRequest) *NeutronListFirewallRulesInvoker {
	requestDef := GenReqDefForNeutronListFirewallRules()
	return &NeutronListFirewallRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronRemoveFirewallRule(request *model.NeutronRemoveFirewallRuleRequest) (*model.NeutronRemoveFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronRemoveFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronRemoveFirewallRuleResponse), nil
	}
}

func (c *VpcClient) NeutronRemoveFirewallRuleInvoker(request *model.NeutronRemoveFirewallRuleRequest) *NeutronRemoveFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronRemoveFirewallRule()
	return &NeutronRemoveFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronShowFirewallGroup(request *model.NeutronShowFirewallGroupRequest) (*model.NeutronShowFirewallGroupResponse, error) {
	requestDef := GenReqDefForNeutronShowFirewallGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowFirewallGroupResponse), nil
	}
}

func (c *VpcClient) NeutronShowFirewallGroupInvoker(request *model.NeutronShowFirewallGroupRequest) *NeutronShowFirewallGroupInvoker {
	requestDef := GenReqDefForNeutronShowFirewallGroup()
	return &NeutronShowFirewallGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronShowFirewallPolicy(request *model.NeutronShowFirewallPolicyRequest) (*model.NeutronShowFirewallPolicyResponse, error) {
	requestDef := GenReqDefForNeutronShowFirewallPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowFirewallPolicyResponse), nil
	}
}

func (c *VpcClient) NeutronShowFirewallPolicyInvoker(request *model.NeutronShowFirewallPolicyRequest) *NeutronShowFirewallPolicyInvoker {
	requestDef := GenReqDefForNeutronShowFirewallPolicy()
	return &NeutronShowFirewallPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronShowFirewallRule(request *model.NeutronShowFirewallRuleRequest) (*model.NeutronShowFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronShowFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowFirewallRuleResponse), nil
	}
}

func (c *VpcClient) NeutronShowFirewallRuleInvoker(request *model.NeutronShowFirewallRuleRequest) *NeutronShowFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronShowFirewallRule()
	return &NeutronShowFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronUpdateFirewallGroup(request *model.NeutronUpdateFirewallGroupRequest) (*model.NeutronUpdateFirewallGroupResponse, error) {
	requestDef := GenReqDefForNeutronUpdateFirewallGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateFirewallGroupResponse), nil
	}
}

func (c *VpcClient) NeutronUpdateFirewallGroupInvoker(request *model.NeutronUpdateFirewallGroupRequest) *NeutronUpdateFirewallGroupInvoker {
	requestDef := GenReqDefForNeutronUpdateFirewallGroup()
	return &NeutronUpdateFirewallGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronUpdateFirewallPolicy(request *model.NeutronUpdateFirewallPolicyRequest) (*model.NeutronUpdateFirewallPolicyResponse, error) {
	requestDef := GenReqDefForNeutronUpdateFirewallPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateFirewallPolicyResponse), nil
	}
}

func (c *VpcClient) NeutronUpdateFirewallPolicyInvoker(request *model.NeutronUpdateFirewallPolicyRequest) *NeutronUpdateFirewallPolicyInvoker {
	requestDef := GenReqDefForNeutronUpdateFirewallPolicy()
	return &NeutronUpdateFirewallPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) NeutronUpdateFirewallRule(request *model.NeutronUpdateFirewallRuleRequest) (*model.NeutronUpdateFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronUpdateFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateFirewallRuleResponse), nil
	}
}

func (c *VpcClient) NeutronUpdateFirewallRuleInvoker(request *model.NeutronUpdateFirewallRuleRequest) *NeutronUpdateFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronUpdateFirewallRule()
	return &NeutronUpdateFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) BatchCreateVpcTags(request *model.BatchCreateVpcTagsRequest) (*model.BatchCreateVpcTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateVpcTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateVpcTagsResponse), nil
	}
}

func (c *VpcClient) BatchCreateVpcTagsInvoker(request *model.BatchCreateVpcTagsRequest) *BatchCreateVpcTagsInvoker {
	requestDef := GenReqDefForBatchCreateVpcTags()
	return &BatchCreateVpcTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) BatchDeleteVpcTags(request *model.BatchDeleteVpcTagsRequest) (*model.BatchDeleteVpcTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteVpcTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteVpcTagsResponse), nil
	}
}

func (c *VpcClient) BatchDeleteVpcTagsInvoker(request *model.BatchDeleteVpcTagsRequest) *BatchDeleteVpcTagsInvoker {
	requestDef := GenReqDefForBatchDeleteVpcTags()
	return &BatchDeleteVpcTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateVpc(request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	requestDef := GenReqDefForCreateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcResponse), nil
	}
}

func (c *VpcClient) CreateVpcInvoker(request *model.CreateVpcRequest) *CreateVpcInvoker {
	requestDef := GenReqDefForCreateVpc()
	return &CreateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateVpcResourceTag(request *model.CreateVpcResourceTagRequest) (*model.CreateVpcResourceTagResponse, error) {
	requestDef := GenReqDefForCreateVpcResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcResourceTagResponse), nil
	}
}

func (c *VpcClient) CreateVpcResourceTagInvoker(request *model.CreateVpcResourceTagRequest) *CreateVpcResourceTagInvoker {
	requestDef := GenReqDefForCreateVpcResourceTag()
	return &CreateVpcResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateVpcRoute(request *model.CreateVpcRouteRequest) (*model.CreateVpcRouteResponse, error) {
	requestDef := GenReqDefForCreateVpcRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcRouteResponse), nil
	}
}

func (c *VpcClient) CreateVpcRouteInvoker(request *model.CreateVpcRouteRequest) *CreateVpcRouteInvoker {
	requestDef := GenReqDefForCreateVpcRoute()
	return &CreateVpcRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteVpc(request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	requestDef := GenReqDefForDeleteVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcResponse), nil
	}
}

func (c *VpcClient) DeleteVpcInvoker(request *model.DeleteVpcRequest) *DeleteVpcInvoker {
	requestDef := GenReqDefForDeleteVpc()
	return &DeleteVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteVpcRoute(request *model.DeleteVpcRouteRequest) (*model.DeleteVpcRouteResponse, error) {
	requestDef := GenReqDefForDeleteVpcRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcRouteResponse), nil
	}
}

func (c *VpcClient) DeleteVpcRouteInvoker(request *model.DeleteVpcRouteRequest) *DeleteVpcRouteInvoker {
	requestDef := GenReqDefForDeleteVpcRoute()
	return &DeleteVpcRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteVpcTag(request *model.DeleteVpcTagRequest) (*model.DeleteVpcTagResponse, error) {
	requestDef := GenReqDefForDeleteVpcTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcTagResponse), nil
	}
}

func (c *VpcClient) DeleteVpcTagInvoker(request *model.DeleteVpcTagRequest) *DeleteVpcTagInvoker {
	requestDef := GenReqDefForDeleteVpcTag()
	return &DeleteVpcTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListVpcRoutes(request *model.ListVpcRoutesRequest) (*model.ListVpcRoutesResponse, error) {
	requestDef := GenReqDefForListVpcRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcRoutesResponse), nil
	}
}

func (c *VpcClient) ListVpcRoutesInvoker(request *model.ListVpcRoutesRequest) *ListVpcRoutesInvoker {
	requestDef := GenReqDefForListVpcRoutes()
	return &ListVpcRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListVpcTags(request *model.ListVpcTagsRequest) (*model.ListVpcTagsResponse, error) {
	requestDef := GenReqDefForListVpcTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcTagsResponse), nil
	}
}

func (c *VpcClient) ListVpcTagsInvoker(request *model.ListVpcTagsRequest) *ListVpcTagsInvoker {
	requestDef := GenReqDefForListVpcTags()
	return &ListVpcTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListVpcs(request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	requestDef := GenReqDefForListVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcsResponse), nil
	}
}

func (c *VpcClient) ListVpcsInvoker(request *model.ListVpcsRequest) *ListVpcsInvoker {
	requestDef := GenReqDefForListVpcs()
	return &ListVpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListVpcsByTags(request *model.ListVpcsByTagsRequest) (*model.ListVpcsByTagsResponse, error) {
	requestDef := GenReqDefForListVpcsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcsByTagsResponse), nil
	}
}

func (c *VpcClient) ListVpcsByTagsInvoker(request *model.ListVpcsByTagsRequest) *ListVpcsByTagsInvoker {
	requestDef := GenReqDefForListVpcsByTags()
	return &ListVpcsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowVpc(request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	requestDef := GenReqDefForShowVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcResponse), nil
	}
}

func (c *VpcClient) ShowVpcInvoker(request *model.ShowVpcRequest) *ShowVpcInvoker {
	requestDef := GenReqDefForShowVpc()
	return &ShowVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowVpcRoute(request *model.ShowVpcRouteRequest) (*model.ShowVpcRouteResponse, error) {
	requestDef := GenReqDefForShowVpcRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcRouteResponse), nil
	}
}

func (c *VpcClient) ShowVpcRouteInvoker(request *model.ShowVpcRouteRequest) *ShowVpcRouteInvoker {
	requestDef := GenReqDefForShowVpcRoute()
	return &ShowVpcRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowVpcTags(request *model.ShowVpcTagsRequest) (*model.ShowVpcTagsResponse, error) {
	requestDef := GenReqDefForShowVpcTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcTagsResponse), nil
	}
}

func (c *VpcClient) ShowVpcTagsInvoker(request *model.ShowVpcTagsRequest) *ShowVpcTagsInvoker {
	requestDef := GenReqDefForShowVpcTags()
	return &ShowVpcTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateVpc(request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	requestDef := GenReqDefForUpdateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcResponse), nil
	}
}

func (c *VpcClient) UpdateVpcInvoker(request *model.UpdateVpcRequest) *UpdateVpcInvoker {
	requestDef := GenReqDefForUpdateVpc()
	return &UpdateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
