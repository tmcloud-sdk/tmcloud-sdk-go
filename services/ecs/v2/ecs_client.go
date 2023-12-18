package v2

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/ecs/v2/model"
)

type EcsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEcsClient(hcClient *http_client.HcHttpClient) *EcsClient {
	return &EcsClient{HcClient: hcClient}
}

func EcsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *EcsClient) AddServerGroupMember(request *model.AddServerGroupMemberRequest) (*model.AddServerGroupMemberResponse, error) {
	requestDef := GenReqDefForAddServerGroupMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServerGroupMemberResponse), nil
	}
}

func (c *EcsClient) AddServerGroupMemberInvoker(request *model.AddServerGroupMemberRequest) *AddServerGroupMemberInvoker {
	requestDef := GenReqDefForAddServerGroupMember()
	return &AddServerGroupMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) AssociateServerVirtualIp(request *model.AssociateServerVirtualIpRequest) (*model.AssociateServerVirtualIpResponse, error) {
	requestDef := GenReqDefForAssociateServerVirtualIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateServerVirtualIpResponse), nil
	}
}

func (c *EcsClient) AssociateServerVirtualIpInvoker(request *model.AssociateServerVirtualIpRequest) *AssociateServerVirtualIpInvoker {
	requestDef := GenReqDefForAssociateServerVirtualIp()
	return &AssociateServerVirtualIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) AttachServerVolume(request *model.AttachServerVolumeRequest) (*model.AttachServerVolumeResponse, error) {
	requestDef := GenReqDefForAttachServerVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachServerVolumeResponse), nil
	}
}

func (c *EcsClient) AttachServerVolumeInvoker(request *model.AttachServerVolumeRequest) *AttachServerVolumeInvoker {
	requestDef := GenReqDefForAttachServerVolume()
	return &AttachServerVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchAddServerNics(request *model.BatchAddServerNicsRequest) (*model.BatchAddServerNicsResponse, error) {
	requestDef := GenReqDefForBatchAddServerNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddServerNicsResponse), nil
	}
}

func (c *EcsClient) BatchAddServerNicsInvoker(request *model.BatchAddServerNicsRequest) *BatchAddServerNicsInvoker {
	requestDef := GenReqDefForBatchAddServerNics()
	return &BatchAddServerNicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchAttachSharableVolumes(request *model.BatchAttachSharableVolumesRequest) (*model.BatchAttachSharableVolumesResponse, error) {
	requestDef := GenReqDefForBatchAttachSharableVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAttachSharableVolumesResponse), nil
	}
}

func (c *EcsClient) BatchAttachSharableVolumesInvoker(request *model.BatchAttachSharableVolumesRequest) *BatchAttachSharableVolumesInvoker {
	requestDef := GenReqDefForBatchAttachSharableVolumes()
	return &BatchAttachSharableVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchCreateServerTags(request *model.BatchCreateServerTagsRequest) (*model.BatchCreateServerTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateServerTagsResponse), nil
	}
}

func (c *EcsClient) BatchCreateServerTagsInvoker(request *model.BatchCreateServerTagsRequest) *BatchCreateServerTagsInvoker {
	requestDef := GenReqDefForBatchCreateServerTags()
	return &BatchCreateServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchDeleteServerNics(request *model.BatchDeleteServerNicsRequest) (*model.BatchDeleteServerNicsResponse, error) {
	requestDef := GenReqDefForBatchDeleteServerNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServerNicsResponse), nil
	}
}

func (c *EcsClient) BatchDeleteServerNicsInvoker(request *model.BatchDeleteServerNicsRequest) *BatchDeleteServerNicsInvoker {
	requestDef := GenReqDefForBatchDeleteServerNics()
	return &BatchDeleteServerNicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchDeleteServerTags(request *model.BatchDeleteServerTagsRequest) (*model.BatchDeleteServerTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServerTagsResponse), nil
	}
}

func (c *EcsClient) BatchDeleteServerTagsInvoker(request *model.BatchDeleteServerTagsRequest) *BatchDeleteServerTagsInvoker {
	requestDef := GenReqDefForBatchDeleteServerTags()
	return &BatchDeleteServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchRebootServers(request *model.BatchRebootServersRequest) (*model.BatchRebootServersResponse, error) {
	requestDef := GenReqDefForBatchRebootServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebootServersResponse), nil
	}
}

func (c *EcsClient) BatchRebootServersInvoker(request *model.BatchRebootServersRequest) *BatchRebootServersInvoker {
	requestDef := GenReqDefForBatchRebootServers()
	return &BatchRebootServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchResetServersPassword(request *model.BatchResetServersPasswordRequest) (*model.BatchResetServersPasswordResponse, error) {
	requestDef := GenReqDefForBatchResetServersPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResetServersPasswordResponse), nil
	}
}

func (c *EcsClient) BatchResetServersPasswordInvoker(request *model.BatchResetServersPasswordRequest) *BatchResetServersPasswordInvoker {
	requestDef := GenReqDefForBatchResetServersPassword()
	return &BatchResetServersPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchStartServers(request *model.BatchStartServersRequest) (*model.BatchStartServersResponse, error) {
	requestDef := GenReqDefForBatchStartServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartServersResponse), nil
	}
}

func (c *EcsClient) BatchStartServersInvoker(request *model.BatchStartServersRequest) *BatchStartServersInvoker {
	requestDef := GenReqDefForBatchStartServers()
	return &BatchStartServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchStopServers(request *model.BatchStopServersRequest) (*model.BatchStopServersResponse, error) {
	requestDef := GenReqDefForBatchStopServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopServersResponse), nil
	}
}

func (c *EcsClient) BatchStopServersInvoker(request *model.BatchStopServersRequest) *BatchStopServersInvoker {
	requestDef := GenReqDefForBatchStopServers()
	return &BatchStopServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) BatchUpdateServersName(request *model.BatchUpdateServersNameRequest) (*model.BatchUpdateServersNameResponse, error) {
	requestDef := GenReqDefForBatchUpdateServersName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateServersNameResponse), nil
	}
}

func (c *EcsClient) BatchUpdateServersNameInvoker(request *model.BatchUpdateServersNameRequest) *BatchUpdateServersNameInvoker {
	requestDef := GenReqDefForBatchUpdateServersName()
	return &BatchUpdateServersNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ChangeServerOsWithCloudInit(request *model.ChangeServerOsWithCloudInitRequest) (*model.ChangeServerOsWithCloudInitResponse, error) {
	requestDef := GenReqDefForChangeServerOsWithCloudInit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeServerOsWithCloudInitResponse), nil
	}
}

func (c *EcsClient) ChangeServerOsWithCloudInitInvoker(request *model.ChangeServerOsWithCloudInitRequest) *ChangeServerOsWithCloudInitInvoker {
	requestDef := GenReqDefForChangeServerOsWithCloudInit()
	return &ChangeServerOsWithCloudInitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ChangeServerOsWithoutCloudInit(request *model.ChangeServerOsWithoutCloudInitRequest) (*model.ChangeServerOsWithoutCloudInitResponse, error) {
	requestDef := GenReqDefForChangeServerOsWithoutCloudInit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeServerOsWithoutCloudInitResponse), nil
	}
}

func (c *EcsClient) ChangeServerOsWithoutCloudInitInvoker(request *model.ChangeServerOsWithoutCloudInitRequest) *ChangeServerOsWithoutCloudInitInvoker {
	requestDef := GenReqDefForChangeServerOsWithoutCloudInit()
	return &ChangeServerOsWithoutCloudInitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) CreatePostPaidServers(request *model.CreatePostPaidServersRequest) (*model.CreatePostPaidServersResponse, error) {
	requestDef := GenReqDefForCreatePostPaidServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidServersResponse), nil
	}
}

func (c *EcsClient) CreatePostPaidServersInvoker(request *model.CreatePostPaidServersRequest) *CreatePostPaidServersInvoker {
	requestDef := GenReqDefForCreatePostPaidServers()
	return &CreatePostPaidServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) CreateServerGroup(request *model.CreateServerGroupRequest) (*model.CreateServerGroupResponse, error) {
	requestDef := GenReqDefForCreateServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServerGroupResponse), nil
	}
}

func (c *EcsClient) CreateServerGroupInvoker(request *model.CreateServerGroupRequest) *CreateServerGroupInvoker {
	requestDef := GenReqDefForCreateServerGroup()
	return &CreateServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) CreateServers(request *model.CreateServersRequest) (*model.CreateServersResponse, error) {
	requestDef := GenReqDefForCreateServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServersResponse), nil
	}
}

func (c *EcsClient) CreateServersInvoker(request *model.CreateServersRequest) *CreateServersInvoker {
	requestDef := GenReqDefForCreateServers()
	return &CreateServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) DeleteServerGroup(request *model.DeleteServerGroupRequest) (*model.DeleteServerGroupResponse, error) {
	requestDef := GenReqDefForDeleteServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerGroupResponse), nil
	}
}

func (c *EcsClient) DeleteServerGroupInvoker(request *model.DeleteServerGroupRequest) *DeleteServerGroupInvoker {
	requestDef := GenReqDefForDeleteServerGroup()
	return &DeleteServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) DeleteServerGroupMember(request *model.DeleteServerGroupMemberRequest) (*model.DeleteServerGroupMemberResponse, error) {
	requestDef := GenReqDefForDeleteServerGroupMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerGroupMemberResponse), nil
	}
}

func (c *EcsClient) DeleteServerGroupMemberInvoker(request *model.DeleteServerGroupMemberRequest) *DeleteServerGroupMemberInvoker {
	requestDef := GenReqDefForDeleteServerGroupMember()
	return &DeleteServerGroupMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) DeleteServerMetadata(request *model.DeleteServerMetadataRequest) (*model.DeleteServerMetadataResponse, error) {
	requestDef := GenReqDefForDeleteServerMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerMetadataResponse), nil
	}
}

func (c *EcsClient) DeleteServerMetadataInvoker(request *model.DeleteServerMetadataRequest) *DeleteServerMetadataInvoker {
	requestDef := GenReqDefForDeleteServerMetadata()
	return &DeleteServerMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) DeleteServerPassword(request *model.DeleteServerPasswordRequest) (*model.DeleteServerPasswordResponse, error) {
	requestDef := GenReqDefForDeleteServerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerPasswordResponse), nil
	}
}

func (c *EcsClient) DeleteServerPasswordInvoker(request *model.DeleteServerPasswordRequest) *DeleteServerPasswordInvoker {
	requestDef := GenReqDefForDeleteServerPassword()
	return &DeleteServerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) DeleteServers(request *model.DeleteServersRequest) (*model.DeleteServersResponse, error) {
	requestDef := GenReqDefForDeleteServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServersResponse), nil
	}
}

func (c *EcsClient) DeleteServersInvoker(request *model.DeleteServersRequest) *DeleteServersInvoker {
	requestDef := GenReqDefForDeleteServers()
	return &DeleteServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) DetachServerVolume(request *model.DetachServerVolumeRequest) (*model.DetachServerVolumeResponse, error) {
	requestDef := GenReqDefForDetachServerVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachServerVolumeResponse), nil
	}
}

func (c *EcsClient) DetachServerVolumeInvoker(request *model.DetachServerVolumeRequest) *DetachServerVolumeInvoker {
	requestDef := GenReqDefForDetachServerVolume()
	return &DetachServerVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) DisassociateServerVirtualIp(request *model.DisassociateServerVirtualIpRequest) (*model.DisassociateServerVirtualIpResponse, error) {
	requestDef := GenReqDefForDisassociateServerVirtualIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateServerVirtualIpResponse), nil
	}
}

func (c *EcsClient) DisassociateServerVirtualIpInvoker(request *model.DisassociateServerVirtualIpRequest) *DisassociateServerVirtualIpInvoker {
	requestDef := GenReqDefForDisassociateServerVirtualIp()
	return &DisassociateServerVirtualIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

func (c *EcsClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ListResizeFlavors(request *model.ListResizeFlavorsRequest) (*model.ListResizeFlavorsResponse, error) {
	requestDef := GenReqDefForListResizeFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResizeFlavorsResponse), nil
	}
}

func (c *EcsClient) ListResizeFlavorsInvoker(request *model.ListResizeFlavorsRequest) *ListResizeFlavorsInvoker {
	requestDef := GenReqDefForListResizeFlavors()
	return &ListResizeFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ListServerBlockDevices(request *model.ListServerBlockDevicesRequest) (*model.ListServerBlockDevicesResponse, error) {
	requestDef := GenReqDefForListServerBlockDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerBlockDevicesResponse), nil
	}
}

func (c *EcsClient) ListServerBlockDevicesInvoker(request *model.ListServerBlockDevicesRequest) *ListServerBlockDevicesInvoker {
	requestDef := GenReqDefForListServerBlockDevices()
	return &ListServerBlockDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ListServerGroups(request *model.ListServerGroupsRequest) (*model.ListServerGroupsResponse, error) {
	requestDef := GenReqDefForListServerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerGroupsResponse), nil
	}
}

func (c *EcsClient) ListServerGroupsInvoker(request *model.ListServerGroupsRequest) *ListServerGroupsInvoker {
	requestDef := GenReqDefForListServerGroups()
	return &ListServerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ListServerInterfaces(request *model.ListServerInterfacesRequest) (*model.ListServerInterfacesResponse, error) {
	requestDef := GenReqDefForListServerInterfaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerInterfacesResponse), nil
	}
}

func (c *EcsClient) ListServerInterfacesInvoker(request *model.ListServerInterfacesRequest) *ListServerInterfacesInvoker {
	requestDef := GenReqDefForListServerInterfaces()
	return &ListServerInterfacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ListServerTags(request *model.ListServerTagsRequest) (*model.ListServerTagsResponse, error) {
	requestDef := GenReqDefForListServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerTagsResponse), nil
	}
}

func (c *EcsClient) ListServerTagsInvoker(request *model.ListServerTagsRequest) *ListServerTagsInvoker {
	requestDef := GenReqDefForListServerTags()
	return &ListServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ListServersDetails(request *model.ListServersDetailsRequest) (*model.ListServersDetailsResponse, error) {
	requestDef := GenReqDefForListServersDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersDetailsResponse), nil
	}
}

func (c *EcsClient) ListServersDetailsInvoker(request *model.ListServersDetailsRequest) *ListServersDetailsInvoker {
	requestDef := GenReqDefForListServersDetails()
	return &ListServersDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) MigrateServer(request *model.MigrateServerRequest) (*model.MigrateServerResponse, error) {
	requestDef := GenReqDefForMigrateServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateServerResponse), nil
	}
}

func (c *EcsClient) MigrateServerInvoker(request *model.MigrateServerRequest) *MigrateServerInvoker {
	requestDef := GenReqDefForMigrateServer()
	return &MigrateServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaAssociateSecurityGroup(request *model.NovaAssociateSecurityGroupRequest) (*model.NovaAssociateSecurityGroupResponse, error) {
	requestDef := GenReqDefForNovaAssociateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaAssociateSecurityGroupResponse), nil
	}
}

func (c *EcsClient) NovaAssociateSecurityGroupInvoker(request *model.NovaAssociateSecurityGroupRequest) *NovaAssociateSecurityGroupInvoker {
	requestDef := GenReqDefForNovaAssociateSecurityGroup()
	return &NovaAssociateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaAttachInterface(request *model.NovaAttachInterfaceRequest) (*model.NovaAttachInterfaceResponse, error) {
	requestDef := GenReqDefForNovaAttachInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaAttachInterfaceResponse), nil
	}
}

func (c *EcsClient) NovaAttachInterfaceInvoker(request *model.NovaAttachInterfaceRequest) *NovaAttachInterfaceInvoker {
	requestDef := GenReqDefForNovaAttachInterface()
	return &NovaAttachInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaCreateKeypair(request *model.NovaCreateKeypairRequest) (*model.NovaCreateKeypairResponse, error) {
	requestDef := GenReqDefForNovaCreateKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaCreateKeypairResponse), nil
	}
}

func (c *EcsClient) NovaCreateKeypairInvoker(request *model.NovaCreateKeypairRequest) *NovaCreateKeypairInvoker {
	requestDef := GenReqDefForNovaCreateKeypair()
	return &NovaCreateKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaCreateServers(request *model.NovaCreateServersRequest) (*model.NovaCreateServersResponse, error) {
	requestDef := GenReqDefForNovaCreateServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaCreateServersResponse), nil
	}
}

func (c *EcsClient) NovaCreateServersInvoker(request *model.NovaCreateServersRequest) *NovaCreateServersInvoker {
	requestDef := GenReqDefForNovaCreateServers()
	return &NovaCreateServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaDeleteKeypair(request *model.NovaDeleteKeypairRequest) (*model.NovaDeleteKeypairResponse, error) {
	requestDef := GenReqDefForNovaDeleteKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaDeleteKeypairResponse), nil
	}
}

func (c *EcsClient) NovaDeleteKeypairInvoker(request *model.NovaDeleteKeypairRequest) *NovaDeleteKeypairInvoker {
	requestDef := GenReqDefForNovaDeleteKeypair()
	return &NovaDeleteKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaDeleteServer(request *model.NovaDeleteServerRequest) (*model.NovaDeleteServerResponse, error) {
	requestDef := GenReqDefForNovaDeleteServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaDeleteServerResponse), nil
	}
}

func (c *EcsClient) NovaDeleteServerInvoker(request *model.NovaDeleteServerRequest) *NovaDeleteServerInvoker {
	requestDef := GenReqDefForNovaDeleteServer()
	return &NovaDeleteServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaDisassociateSecurityGroup(request *model.NovaDisassociateSecurityGroupRequest) (*model.NovaDisassociateSecurityGroupResponse, error) {
	requestDef := GenReqDefForNovaDisassociateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaDisassociateSecurityGroupResponse), nil
	}
}

func (c *EcsClient) NovaDisassociateSecurityGroupInvoker(request *model.NovaDisassociateSecurityGroupRequest) *NovaDisassociateSecurityGroupInvoker {
	requestDef := GenReqDefForNovaDisassociateSecurityGroup()
	return &NovaDisassociateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaListAvailabilityZones(request *model.NovaListAvailabilityZonesRequest) (*model.NovaListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForNovaListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListAvailabilityZonesResponse), nil
	}
}

func (c *EcsClient) NovaListAvailabilityZonesInvoker(request *model.NovaListAvailabilityZonesRequest) *NovaListAvailabilityZonesInvoker {
	requestDef := GenReqDefForNovaListAvailabilityZones()
	return &NovaListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaListKeypairs(request *model.NovaListKeypairsRequest) (*model.NovaListKeypairsResponse, error) {
	requestDef := GenReqDefForNovaListKeypairs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListKeypairsResponse), nil
	}
}

func (c *EcsClient) NovaListKeypairsInvoker(request *model.NovaListKeypairsRequest) *NovaListKeypairsInvoker {
	requestDef := GenReqDefForNovaListKeypairs()
	return &NovaListKeypairsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaListServerSecurityGroups(request *model.NovaListServerSecurityGroupsRequest) (*model.NovaListServerSecurityGroupsResponse, error) {
	requestDef := GenReqDefForNovaListServerSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListServerSecurityGroupsResponse), nil
	}
}

func (c *EcsClient) NovaListServerSecurityGroupsInvoker(request *model.NovaListServerSecurityGroupsRequest) *NovaListServerSecurityGroupsInvoker {
	requestDef := GenReqDefForNovaListServerSecurityGroups()
	return &NovaListServerSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaListServersDetails(request *model.NovaListServersDetailsRequest) (*model.NovaListServersDetailsResponse, error) {
	requestDef := GenReqDefForNovaListServersDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListServersDetailsResponse), nil
	}
}

func (c *EcsClient) NovaListServersDetailsInvoker(request *model.NovaListServersDetailsRequest) *NovaListServersDetailsInvoker {
	requestDef := GenReqDefForNovaListServersDetails()
	return &NovaListServersDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaShowKeypair(request *model.NovaShowKeypairRequest) (*model.NovaShowKeypairResponse, error) {
	requestDef := GenReqDefForNovaShowKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaShowKeypairResponse), nil
	}
}

func (c *EcsClient) NovaShowKeypairInvoker(request *model.NovaShowKeypairRequest) *NovaShowKeypairInvoker {
	requestDef := GenReqDefForNovaShowKeypair()
	return &NovaShowKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) NovaShowServer(request *model.NovaShowServerRequest) (*model.NovaShowServerResponse, error) {
	requestDef := GenReqDefForNovaShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaShowServerResponse), nil
	}
}

func (c *EcsClient) NovaShowServerInvoker(request *model.NovaShowServerRequest) *NovaShowServerInvoker {
	requestDef := GenReqDefForNovaShowServer()
	return &NovaShowServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) RegisterServerAutoRecovery(request *model.RegisterServerAutoRecoveryRequest) (*model.RegisterServerAutoRecoveryResponse, error) {
	requestDef := GenReqDefForRegisterServerAutoRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterServerAutoRecoveryResponse), nil
	}
}

func (c *EcsClient) RegisterServerAutoRecoveryInvoker(request *model.RegisterServerAutoRecoveryRequest) *RegisterServerAutoRecoveryInvoker {
	requestDef := GenReqDefForRegisterServerAutoRecovery()
	return &RegisterServerAutoRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ReinstallServerWithCloudInit(request *model.ReinstallServerWithCloudInitRequest) (*model.ReinstallServerWithCloudInitResponse, error) {
	requestDef := GenReqDefForReinstallServerWithCloudInit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReinstallServerWithCloudInitResponse), nil
	}
}

func (c *EcsClient) ReinstallServerWithCloudInitInvoker(request *model.ReinstallServerWithCloudInitRequest) *ReinstallServerWithCloudInitInvoker {
	requestDef := GenReqDefForReinstallServerWithCloudInit()
	return &ReinstallServerWithCloudInitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ReinstallServerWithoutCloudInit(request *model.ReinstallServerWithoutCloudInitRequest) (*model.ReinstallServerWithoutCloudInitResponse, error) {
	requestDef := GenReqDefForReinstallServerWithoutCloudInit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReinstallServerWithoutCloudInitResponse), nil
	}
}

func (c *EcsClient) ReinstallServerWithoutCloudInitInvoker(request *model.ReinstallServerWithoutCloudInitRequest) *ReinstallServerWithoutCloudInitInvoker {
	requestDef := GenReqDefForReinstallServerWithoutCloudInit()
	return &ReinstallServerWithoutCloudInitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ResetServerPassword(request *model.ResetServerPasswordRequest) (*model.ResetServerPasswordResponse, error) {
	requestDef := GenReqDefForResetServerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetServerPasswordResponse), nil
	}
}

func (c *EcsClient) ResetServerPasswordInvoker(request *model.ResetServerPasswordRequest) *ResetServerPasswordInvoker {
	requestDef := GenReqDefForResetServerPassword()
	return &ResetServerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ResizePostPaidServer(request *model.ResizePostPaidServerRequest) (*model.ResizePostPaidServerResponse, error) {
	requestDef := GenReqDefForResizePostPaidServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizePostPaidServerResponse), nil
	}
}

func (c *EcsClient) ResizePostPaidServerInvoker(request *model.ResizePostPaidServerRequest) *ResizePostPaidServerInvoker {
	requestDef := GenReqDefForResizePostPaidServer()
	return &ResizePostPaidServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ResizeServer(request *model.ResizeServerRequest) (*model.ResizeServerResponse, error) {
	requestDef := GenReqDefForResizeServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeServerResponse), nil
	}
}

func (c *EcsClient) ResizeServerInvoker(request *model.ResizeServerRequest) *ResizeServerInvoker {
	requestDef := GenReqDefForResizeServer()
	return &ResizeServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowResetPasswordFlag(request *model.ShowResetPasswordFlagRequest) (*model.ShowResetPasswordFlagResponse, error) {
	requestDef := GenReqDefForShowResetPasswordFlag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResetPasswordFlagResponse), nil
	}
}

func (c *EcsClient) ShowResetPasswordFlagInvoker(request *model.ShowResetPasswordFlagRequest) *ShowResetPasswordFlagInvoker {
	requestDef := GenReqDefForShowResetPasswordFlag()
	return &ShowResetPasswordFlagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowServer(request *model.ShowServerRequest) (*model.ShowServerResponse, error) {
	requestDef := GenReqDefForShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerResponse), nil
	}
}

func (c *EcsClient) ShowServerInvoker(request *model.ShowServerRequest) *ShowServerInvoker {
	requestDef := GenReqDefForShowServer()
	return &ShowServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowServerAutoRecovery(request *model.ShowServerAutoRecoveryRequest) (*model.ShowServerAutoRecoveryResponse, error) {
	requestDef := GenReqDefForShowServerAutoRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerAutoRecoveryResponse), nil
	}
}

func (c *EcsClient) ShowServerAutoRecoveryInvoker(request *model.ShowServerAutoRecoveryRequest) *ShowServerAutoRecoveryInvoker {
	requestDef := GenReqDefForShowServerAutoRecovery()
	return &ShowServerAutoRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowServerBlockDevice(request *model.ShowServerBlockDeviceRequest) (*model.ShowServerBlockDeviceResponse, error) {
	requestDef := GenReqDefForShowServerBlockDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerBlockDeviceResponse), nil
	}
}

func (c *EcsClient) ShowServerBlockDeviceInvoker(request *model.ShowServerBlockDeviceRequest) *ShowServerBlockDeviceInvoker {
	requestDef := GenReqDefForShowServerBlockDevice()
	return &ShowServerBlockDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowServerGroup(request *model.ShowServerGroupRequest) (*model.ShowServerGroupResponse, error) {
	requestDef := GenReqDefForShowServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerGroupResponse), nil
	}
}

func (c *EcsClient) ShowServerGroupInvoker(request *model.ShowServerGroupRequest) *ShowServerGroupInvoker {
	requestDef := GenReqDefForShowServerGroup()
	return &ShowServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowServerLimits(request *model.ShowServerLimitsRequest) (*model.ShowServerLimitsResponse, error) {
	requestDef := GenReqDefForShowServerLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerLimitsResponse), nil
	}
}

func (c *EcsClient) ShowServerLimitsInvoker(request *model.ShowServerLimitsRequest) *ShowServerLimitsInvoker {
	requestDef := GenReqDefForShowServerLimits()
	return &ShowServerLimitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowServerPassword(request *model.ShowServerPasswordRequest) (*model.ShowServerPasswordResponse, error) {
	requestDef := GenReqDefForShowServerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerPasswordResponse), nil
	}
}

func (c *EcsClient) ShowServerPasswordInvoker(request *model.ShowServerPasswordRequest) *ShowServerPasswordInvoker {
	requestDef := GenReqDefForShowServerPassword()
	return &ShowServerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowServerRemoteConsole(request *model.ShowServerRemoteConsoleRequest) (*model.ShowServerRemoteConsoleResponse, error) {
	requestDef := GenReqDefForShowServerRemoteConsole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerRemoteConsoleResponse), nil
	}
}

func (c *EcsClient) ShowServerRemoteConsoleInvoker(request *model.ShowServerRemoteConsoleRequest) *ShowServerRemoteConsoleInvoker {
	requestDef := GenReqDefForShowServerRemoteConsole()
	return &ShowServerRemoteConsoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowServerTags(request *model.ShowServerTagsRequest) (*model.ShowServerTagsResponse, error) {
	requestDef := GenReqDefForShowServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerTagsResponse), nil
	}
}

func (c *EcsClient) ShowServerTagsInvoker(request *model.ShowServerTagsRequest) *ShowServerTagsInvoker {
	requestDef := GenReqDefForShowServerTags()
	return &ShowServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) UpdateServer(request *model.UpdateServerRequest) (*model.UpdateServerResponse, error) {
	requestDef := GenReqDefForUpdateServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerResponse), nil
	}
}

func (c *EcsClient) UpdateServerInvoker(request *model.UpdateServerRequest) *UpdateServerInvoker {
	requestDef := GenReqDefForUpdateServer()
	return &UpdateServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) UpdateServerAutoTerminateTime(request *model.UpdateServerAutoTerminateTimeRequest) (*model.UpdateServerAutoTerminateTimeResponse, error) {
	requestDef := GenReqDefForUpdateServerAutoTerminateTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerAutoTerminateTimeResponse), nil
	}
}

func (c *EcsClient) UpdateServerAutoTerminateTimeInvoker(request *model.UpdateServerAutoTerminateTimeRequest) *UpdateServerAutoTerminateTimeInvoker {
	requestDef := GenReqDefForUpdateServerAutoTerminateTime()
	return &UpdateServerAutoTerminateTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) UpdateServerMetadata(request *model.UpdateServerMetadataRequest) (*model.UpdateServerMetadataResponse, error) {
	requestDef := GenReqDefForUpdateServerMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerMetadataResponse), nil
	}
}

func (c *EcsClient) UpdateServerMetadataInvoker(request *model.UpdateServerMetadataRequest) *UpdateServerMetadataInvoker {
	requestDef := GenReqDefForUpdateServerMetadata()
	return &UpdateServerMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EcsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

func (c *EcsClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
