package v3

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/cce/v3/model"
)

type CceClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCceClient(hcClient *http_client.HcHttpClient) *CceClient {
	return &CceClient{HcClient: hcClient}
}

func CceClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *CceClient) AwakeCluster(request *model.AwakeClusterRequest) (*model.AwakeClusterResponse, error) {
	requestDef := GenReqDefForAwakeCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AwakeClusterResponse), nil
	}
}

func (c *CceClient) AwakeClusterInvoker(request *model.AwakeClusterRequest) *AwakeClusterInvoker {
	requestDef := GenReqDefForAwakeCluster()
	return &AwakeClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) CreateAddonInstance(request *model.CreateAddonInstanceRequest) (*model.CreateAddonInstanceResponse, error) {
	requestDef := GenReqDefForCreateAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAddonInstanceResponse), nil
	}
}

func (c *CceClient) CreateAddonInstanceInvoker(request *model.CreateAddonInstanceRequest) *CreateAddonInstanceInvoker {
	requestDef := GenReqDefForCreateAddonInstance()
	return &CreateAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) CreateCloudPersistentVolumeClaims(request *model.CreateCloudPersistentVolumeClaimsRequest) (*model.CreateCloudPersistentVolumeClaimsResponse, error) {
	requestDef := GenReqDefForCreateCloudPersistentVolumeClaims()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCloudPersistentVolumeClaimsResponse), nil
	}
}

func (c *CceClient) CreateCloudPersistentVolumeClaimsInvoker(request *model.CreateCloudPersistentVolumeClaimsRequest) *CreateCloudPersistentVolumeClaimsInvoker {
	requestDef := GenReqDefForCreateCloudPersistentVolumeClaims()
	return &CreateCloudPersistentVolumeClaimsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

func (c *CceClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) CreateKubernetesClusterCert(request *model.CreateKubernetesClusterCertRequest) (*model.CreateKubernetesClusterCertResponse, error) {
	requestDef := GenReqDefForCreateKubernetesClusterCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKubernetesClusterCertResponse), nil
	}
}

func (c *CceClient) CreateKubernetesClusterCertInvoker(request *model.CreateKubernetesClusterCertRequest) *CreateKubernetesClusterCertInvoker {
	requestDef := GenReqDefForCreateKubernetesClusterCert()
	return &CreateKubernetesClusterCertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) CreateNode(request *model.CreateNodeRequest) (*model.CreateNodeResponse, error) {
	requestDef := GenReqDefForCreateNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNodeResponse), nil
	}
}

func (c *CceClient) CreateNodeInvoker(request *model.CreateNodeRequest) *CreateNodeInvoker {
	requestDef := GenReqDefForCreateNode()
	return &CreateNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) CreateNodePool(request *model.CreateNodePoolRequest) (*model.CreateNodePoolResponse, error) {
	requestDef := GenReqDefForCreateNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNodePoolResponse), nil
	}
}

func (c *CceClient) CreateNodePoolInvoker(request *model.CreateNodePoolRequest) *CreateNodePoolInvoker {
	requestDef := GenReqDefForCreateNodePool()
	return &CreateNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) DeleteAddonInstance(request *model.DeleteAddonInstanceRequest) (*model.DeleteAddonInstanceResponse, error) {
	requestDef := GenReqDefForDeleteAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddonInstanceResponse), nil
	}
}

func (c *CceClient) DeleteAddonInstanceInvoker(request *model.DeleteAddonInstanceRequest) *DeleteAddonInstanceInvoker {
	requestDef := GenReqDefForDeleteAddonInstance()
	return &DeleteAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) DeleteCloudPersistentVolumeClaims(request *model.DeleteCloudPersistentVolumeClaimsRequest) (*model.DeleteCloudPersistentVolumeClaimsResponse, error) {
	requestDef := GenReqDefForDeleteCloudPersistentVolumeClaims()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCloudPersistentVolumeClaimsResponse), nil
	}
}

func (c *CceClient) DeleteCloudPersistentVolumeClaimsInvoker(request *model.DeleteCloudPersistentVolumeClaimsRequest) *DeleteCloudPersistentVolumeClaimsInvoker {
	requestDef := GenReqDefForDeleteCloudPersistentVolumeClaims()
	return &DeleteCloudPersistentVolumeClaimsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

func (c *CceClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) DeleteNode(request *model.DeleteNodeRequest) (*model.DeleteNodeResponse, error) {
	requestDef := GenReqDefForDeleteNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodeResponse), nil
	}
}

func (c *CceClient) DeleteNodeInvoker(request *model.DeleteNodeRequest) *DeleteNodeInvoker {
	requestDef := GenReqDefForDeleteNode()
	return &DeleteNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) DeleteNodePool(request *model.DeleteNodePoolRequest) (*model.DeleteNodePoolResponse, error) {
	requestDef := GenReqDefForDeleteNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodePoolResponse), nil
	}
}

func (c *CceClient) DeleteNodePoolInvoker(request *model.DeleteNodePoolRequest) *DeleteNodePoolInvoker {
	requestDef := GenReqDefForDeleteNodePool()
	return &DeleteNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) HibernateCluster(request *model.HibernateClusterRequest) (*model.HibernateClusterResponse, error) {
	requestDef := GenReqDefForHibernateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HibernateClusterResponse), nil
	}
}

func (c *CceClient) HibernateClusterInvoker(request *model.HibernateClusterRequest) *HibernateClusterInvoker {
	requestDef := GenReqDefForHibernateCluster()
	return &HibernateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ListAddonInstances(request *model.ListAddonInstancesRequest) (*model.ListAddonInstancesResponse, error) {
	requestDef := GenReqDefForListAddonInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddonInstancesResponse), nil
	}
}

func (c *CceClient) ListAddonInstancesInvoker(request *model.ListAddonInstancesRequest) *ListAddonInstancesInvoker {
	requestDef := GenReqDefForListAddonInstances()
	return &ListAddonInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ListAddonTemplates(request *model.ListAddonTemplatesRequest) (*model.ListAddonTemplatesResponse, error) {
	requestDef := GenReqDefForListAddonTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddonTemplatesResponse), nil
	}
}

func (c *CceClient) ListAddonTemplatesInvoker(request *model.ListAddonTemplatesRequest) *ListAddonTemplatesInvoker {
	requestDef := GenReqDefForListAddonTemplates()
	return &ListAddonTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

func (c *CceClient) ListClustersInvoker(request *model.ListClustersRequest) *ListClustersInvoker {
	requestDef := GenReqDefForListClusters()
	return &ListClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ListNodePools(request *model.ListNodePoolsRequest) (*model.ListNodePoolsResponse, error) {
	requestDef := GenReqDefForListNodePools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodePoolsResponse), nil
	}
}

func (c *CceClient) ListNodePoolsInvoker(request *model.ListNodePoolsRequest) *ListNodePoolsInvoker {
	requestDef := GenReqDefForListNodePools()
	return &ListNodePoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ListNodes(request *model.ListNodesRequest) (*model.ListNodesResponse, error) {
	requestDef := GenReqDefForListNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodesResponse), nil
	}
}

func (c *CceClient) ListNodesInvoker(request *model.ListNodesRequest) *ListNodesInvoker {
	requestDef := GenReqDefForListNodes()
	return &ListNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ShowAddonInstance(request *model.ShowAddonInstanceRequest) (*model.ShowAddonInstanceResponse, error) {
	requestDef := GenReqDefForShowAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAddonInstanceResponse), nil
	}
}

func (c *CceClient) ShowAddonInstanceInvoker(request *model.ShowAddonInstanceRequest) *ShowAddonInstanceInvoker {
	requestDef := GenReqDefForShowAddonInstance()
	return &ShowAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ShowCluster(request *model.ShowClusterRequest) (*model.ShowClusterResponse, error) {
	requestDef := GenReqDefForShowCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterResponse), nil
	}
}

func (c *CceClient) ShowClusterInvoker(request *model.ShowClusterRequest) *ShowClusterInvoker {
	requestDef := GenReqDefForShowCluster()
	return &ShowClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

func (c *CceClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ShowNode(request *model.ShowNodeRequest) (*model.ShowNodeResponse, error) {
	requestDef := GenReqDefForShowNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodeResponse), nil
	}
}

func (c *CceClient) ShowNodeInvoker(request *model.ShowNodeRequest) *ShowNodeInvoker {
	requestDef := GenReqDefForShowNode()
	return &ShowNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) ShowNodePool(request *model.ShowNodePoolRequest) (*model.ShowNodePoolResponse, error) {
	requestDef := GenReqDefForShowNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodePoolResponse), nil
	}
}

func (c *CceClient) ShowNodePoolInvoker(request *model.ShowNodePoolRequest) *ShowNodePoolInvoker {
	requestDef := GenReqDefForShowNodePool()
	return &ShowNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) UpdateAddonInstance(request *model.UpdateAddonInstanceRequest) (*model.UpdateAddonInstanceResponse, error) {
	requestDef := GenReqDefForUpdateAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAddonInstanceResponse), nil
	}
}

func (c *CceClient) UpdateAddonInstanceInvoker(request *model.UpdateAddonInstanceRequest) *UpdateAddonInstanceInvoker {
	requestDef := GenReqDefForUpdateAddonInstance()
	return &UpdateAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) UpdateCluster(request *model.UpdateClusterRequest) (*model.UpdateClusterResponse, error) {
	requestDef := GenReqDefForUpdateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterResponse), nil
	}
}

func (c *CceClient) UpdateClusterInvoker(request *model.UpdateClusterRequest) *UpdateClusterInvoker {
	requestDef := GenReqDefForUpdateCluster()
	return &UpdateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) UpdateClusterEip(request *model.UpdateClusterEipRequest) (*model.UpdateClusterEipResponse, error) {
	requestDef := GenReqDefForUpdateClusterEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterEipResponse), nil
	}
}

func (c *CceClient) UpdateClusterEipInvoker(request *model.UpdateClusterEipRequest) *UpdateClusterEipInvoker {
	requestDef := GenReqDefForUpdateClusterEip()
	return &UpdateClusterEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) UpdateNode(request *model.UpdateNodeRequest) (*model.UpdateNodeResponse, error) {
	requestDef := GenReqDefForUpdateNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodeResponse), nil
	}
}

func (c *CceClient) UpdateNodeInvoker(request *model.UpdateNodeRequest) *UpdateNodeInvoker {
	requestDef := GenReqDefForUpdateNode()
	return &UpdateNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CceClient) UpdateNodePool(request *model.UpdateNodePoolRequest) (*model.UpdateNodePoolResponse, error) {
	requestDef := GenReqDefForUpdateNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodePoolResponse), nil
	}
}

func (c *CceClient) UpdateNodePoolInvoker(request *model.UpdateNodePoolRequest) *UpdateNodePoolInvoker {
	requestDef := GenReqDefForUpdateNodePool()
	return &UpdateNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
