package v2

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/kms/v2/model"
)

type KmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewKmsClient(hcClient *http_client.HcHttpClient) *KmsClient {
	return &KmsClient{HcClient: hcClient}
}

func KmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *KmsClient) BatchCreateKmsTags(request *model.BatchCreateKmsTagsRequest) (*model.BatchCreateKmsTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateKmsTagsResponse), nil
	}
}

func (c *KmsClient) BatchCreateKmsTagsInvoker(request *model.BatchCreateKmsTagsRequest) *BatchCreateKmsTagsInvoker {
	requestDef := GenReqDefForBatchCreateKmsTags()
	return &BatchCreateKmsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CancelGrant(request *model.CancelGrantRequest) (*model.CancelGrantResponse, error) {
	requestDef := GenReqDefForCancelGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelGrantResponse), nil
	}
}

func (c *KmsClient) CancelGrantInvoker(request *model.CancelGrantRequest) *CancelGrantInvoker {
	requestDef := GenReqDefForCancelGrant()
	return &CancelGrantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CancelKeyDeletion(request *model.CancelKeyDeletionRequest) (*model.CancelKeyDeletionResponse, error) {
	requestDef := GenReqDefForCancelKeyDeletion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelKeyDeletionResponse), nil
	}
}

func (c *KmsClient) CancelKeyDeletionInvoker(request *model.CancelKeyDeletionRequest) *CancelKeyDeletionInvoker {
	requestDef := GenReqDefForCancelKeyDeletion()
	return &CancelKeyDeletionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CancelSelfGrant(request *model.CancelSelfGrantRequest) (*model.CancelSelfGrantResponse, error) {
	requestDef := GenReqDefForCancelSelfGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSelfGrantResponse), nil
	}
}

func (c *KmsClient) CancelSelfGrantInvoker(request *model.CancelSelfGrantRequest) *CancelSelfGrantInvoker {
	requestDef := GenReqDefForCancelSelfGrant()
	return &CancelSelfGrantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CreateDatakey(request *model.CreateDatakeyRequest) (*model.CreateDatakeyResponse, error) {
	requestDef := GenReqDefForCreateDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatakeyResponse), nil
	}
}

func (c *KmsClient) CreateDatakeyInvoker(request *model.CreateDatakeyRequest) *CreateDatakeyInvoker {
	requestDef := GenReqDefForCreateDatakey()
	return &CreateDatakeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CreateDatakeyWithoutPlaintext(request *model.CreateDatakeyWithoutPlaintextRequest) (*model.CreateDatakeyWithoutPlaintextResponse, error) {
	requestDef := GenReqDefForCreateDatakeyWithoutPlaintext()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatakeyWithoutPlaintextResponse), nil
	}
}

func (c *KmsClient) CreateDatakeyWithoutPlaintextInvoker(request *model.CreateDatakeyWithoutPlaintextRequest) *CreateDatakeyWithoutPlaintextInvoker {
	requestDef := GenReqDefForCreateDatakeyWithoutPlaintext()
	return &CreateDatakeyWithoutPlaintextInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CreateGrant(request *model.CreateGrantRequest) (*model.CreateGrantResponse, error) {
	requestDef := GenReqDefForCreateGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGrantResponse), nil
	}
}

func (c *KmsClient) CreateGrantInvoker(request *model.CreateGrantRequest) *CreateGrantInvoker {
	requestDef := GenReqDefForCreateGrant()
	return &CreateGrantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CreateKey(request *model.CreateKeyRequest) (*model.CreateKeyResponse, error) {
	requestDef := GenReqDefForCreateKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKeyResponse), nil
	}
}

func (c *KmsClient) CreateKeyInvoker(request *model.CreateKeyRequest) *CreateKeyInvoker {
	requestDef := GenReqDefForCreateKey()
	return &CreateKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CreateKeyStore(request *model.CreateKeyStoreRequest) (*model.CreateKeyStoreResponse, error) {
	requestDef := GenReqDefForCreateKeyStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKeyStoreResponse), nil
	}
}

func (c *KmsClient) CreateKeyStoreInvoker(request *model.CreateKeyStoreRequest) *CreateKeyStoreInvoker {
	requestDef := GenReqDefForCreateKeyStore()
	return &CreateKeyStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CreateKmsTag(request *model.CreateKmsTagRequest) (*model.CreateKmsTagResponse, error) {
	requestDef := GenReqDefForCreateKmsTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKmsTagResponse), nil
	}
}

func (c *KmsClient) CreateKmsTagInvoker(request *model.CreateKmsTagRequest) *CreateKmsTagInvoker {
	requestDef := GenReqDefForCreateKmsTag()
	return &CreateKmsTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CreateParametersForImport(request *model.CreateParametersForImportRequest) (*model.CreateParametersForImportResponse, error) {
	requestDef := GenReqDefForCreateParametersForImport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateParametersForImportResponse), nil
	}
}

func (c *KmsClient) CreateParametersForImportInvoker(request *model.CreateParametersForImportRequest) *CreateParametersForImportInvoker {
	requestDef := GenReqDefForCreateParametersForImport()
	return &CreateParametersForImportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) CreateRandom(request *model.CreateRandomRequest) (*model.CreateRandomResponse, error) {
	requestDef := GenReqDefForCreateRandom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRandomResponse), nil
	}
}

func (c *KmsClient) CreateRandomInvoker(request *model.CreateRandomRequest) *CreateRandomInvoker {
	requestDef := GenReqDefForCreateRandom()
	return &CreateRandomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DecryptData(request *model.DecryptDataRequest) (*model.DecryptDataResponse, error) {
	requestDef := GenReqDefForDecryptData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DecryptDataResponse), nil
	}
}

func (c *KmsClient) DecryptDataInvoker(request *model.DecryptDataRequest) *DecryptDataInvoker {
	requestDef := GenReqDefForDecryptData()
	return &DecryptDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DecryptDatakey(request *model.DecryptDatakeyRequest) (*model.DecryptDatakeyResponse, error) {
	requestDef := GenReqDefForDecryptDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DecryptDatakeyResponse), nil
	}
}

func (c *KmsClient) DecryptDatakeyInvoker(request *model.DecryptDatakeyRequest) *DecryptDatakeyInvoker {
	requestDef := GenReqDefForDecryptDatakey()
	return &DecryptDatakeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DeleteImportedKeyMaterial(request *model.DeleteImportedKeyMaterialRequest) (*model.DeleteImportedKeyMaterialResponse, error) {
	requestDef := GenReqDefForDeleteImportedKeyMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImportedKeyMaterialResponse), nil
	}
}

func (c *KmsClient) DeleteImportedKeyMaterialInvoker(request *model.DeleteImportedKeyMaterialRequest) *DeleteImportedKeyMaterialInvoker {
	requestDef := GenReqDefForDeleteImportedKeyMaterial()
	return &DeleteImportedKeyMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DeleteKey(request *model.DeleteKeyRequest) (*model.DeleteKeyResponse, error) {
	requestDef := GenReqDefForDeleteKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeyResponse), nil
	}
}

func (c *KmsClient) DeleteKeyInvoker(request *model.DeleteKeyRequest) *DeleteKeyInvoker {
	requestDef := GenReqDefForDeleteKey()
	return &DeleteKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DeleteKeyStore(request *model.DeleteKeyStoreRequest) (*model.DeleteKeyStoreResponse, error) {
	requestDef := GenReqDefForDeleteKeyStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeyStoreResponse), nil
	}
}

func (c *KmsClient) DeleteKeyStoreInvoker(request *model.DeleteKeyStoreRequest) *DeleteKeyStoreInvoker {
	requestDef := GenReqDefForDeleteKeyStore()
	return &DeleteKeyStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

func (c *KmsClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DisableKey(request *model.DisableKeyRequest) (*model.DisableKeyResponse, error) {
	requestDef := GenReqDefForDisableKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableKeyResponse), nil
	}
}

func (c *KmsClient) DisableKeyInvoker(request *model.DisableKeyRequest) *DisableKeyInvoker {
	requestDef := GenReqDefForDisableKey()
	return &DisableKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DisableKeyRotation(request *model.DisableKeyRotationRequest) (*model.DisableKeyRotationResponse, error) {
	requestDef := GenReqDefForDisableKeyRotation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableKeyRotationResponse), nil
	}
}

func (c *KmsClient) DisableKeyRotationInvoker(request *model.DisableKeyRotationRequest) *DisableKeyRotationInvoker {
	requestDef := GenReqDefForDisableKeyRotation()
	return &DisableKeyRotationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) DisableKeyStore(request *model.DisableKeyStoreRequest) (*model.DisableKeyStoreResponse, error) {
	requestDef := GenReqDefForDisableKeyStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableKeyStoreResponse), nil
	}
}

func (c *KmsClient) DisableKeyStoreInvoker(request *model.DisableKeyStoreRequest) *DisableKeyStoreInvoker {
	requestDef := GenReqDefForDisableKeyStore()
	return &DisableKeyStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) EnableKey(request *model.EnableKeyRequest) (*model.EnableKeyResponse, error) {
	requestDef := GenReqDefForEnableKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableKeyResponse), nil
	}
}

func (c *KmsClient) EnableKeyInvoker(request *model.EnableKeyRequest) *EnableKeyInvoker {
	requestDef := GenReqDefForEnableKey()
	return &EnableKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) EnableKeyRotation(request *model.EnableKeyRotationRequest) (*model.EnableKeyRotationResponse, error) {
	requestDef := GenReqDefForEnableKeyRotation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableKeyRotationResponse), nil
	}
}

func (c *KmsClient) EnableKeyRotationInvoker(request *model.EnableKeyRotationRequest) *EnableKeyRotationInvoker {
	requestDef := GenReqDefForEnableKeyRotation()
	return &EnableKeyRotationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) EnableKeyStore(request *model.EnableKeyStoreRequest) (*model.EnableKeyStoreResponse, error) {
	requestDef := GenReqDefForEnableKeyStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableKeyStoreResponse), nil
	}
}

func (c *KmsClient) EnableKeyStoreInvoker(request *model.EnableKeyStoreRequest) *EnableKeyStoreInvoker {
	requestDef := GenReqDefForEnableKeyStore()
	return &EnableKeyStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) EncryptData(request *model.EncryptDataRequest) (*model.EncryptDataResponse, error) {
	requestDef := GenReqDefForEncryptData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EncryptDataResponse), nil
	}
}

func (c *KmsClient) EncryptDataInvoker(request *model.EncryptDataRequest) *EncryptDataInvoker {
	requestDef := GenReqDefForEncryptData()
	return &EncryptDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) EncryptDatakey(request *model.EncryptDatakeyRequest) (*model.EncryptDatakeyResponse, error) {
	requestDef := GenReqDefForEncryptDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EncryptDatakeyResponse), nil
	}
}

func (c *KmsClient) EncryptDatakeyInvoker(request *model.EncryptDatakeyRequest) *EncryptDatakeyInvoker {
	requestDef := GenReqDefForEncryptDatakey()
	return &EncryptDatakeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ImportKeyMaterial(request *model.ImportKeyMaterialRequest) (*model.ImportKeyMaterialResponse, error) {
	requestDef := GenReqDefForImportKeyMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportKeyMaterialResponse), nil
	}
}

func (c *KmsClient) ImportKeyMaterialInvoker(request *model.ImportKeyMaterialRequest) *ImportKeyMaterialInvoker {
	requestDef := GenReqDefForImportKeyMaterial()
	return &ImportKeyMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ListGrants(request *model.ListGrantsRequest) (*model.ListGrantsResponse, error) {
	requestDef := GenReqDefForListGrants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGrantsResponse), nil
	}
}

func (c *KmsClient) ListGrantsInvoker(request *model.ListGrantsRequest) *ListGrantsInvoker {
	requestDef := GenReqDefForListGrants()
	return &ListGrantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ListKeyDetail(request *model.ListKeyDetailRequest) (*model.ListKeyDetailResponse, error) {
	requestDef := GenReqDefForListKeyDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeyDetailResponse), nil
	}
}

func (c *KmsClient) ListKeyDetailInvoker(request *model.ListKeyDetailRequest) *ListKeyDetailInvoker {
	requestDef := GenReqDefForListKeyDetail()
	return &ListKeyDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ListKeyStores(request *model.ListKeyStoresRequest) (*model.ListKeyStoresResponse, error) {
	requestDef := GenReqDefForListKeyStores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeyStoresResponse), nil
	}
}

func (c *KmsClient) ListKeyStoresInvoker(request *model.ListKeyStoresRequest) *ListKeyStoresInvoker {
	requestDef := GenReqDefForListKeyStores()
	return &ListKeyStoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ListKeys(request *model.ListKeysRequest) (*model.ListKeysResponse, error) {
	requestDef := GenReqDefForListKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeysResponse), nil
	}
}

func (c *KmsClient) ListKeysInvoker(request *model.ListKeysRequest) *ListKeysInvoker {
	requestDef := GenReqDefForListKeys()
	return &ListKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ListKmsByTags(request *model.ListKmsByTagsRequest) (*model.ListKmsByTagsResponse, error) {
	requestDef := GenReqDefForListKmsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKmsByTagsResponse), nil
	}
}

func (c *KmsClient) ListKmsByTagsInvoker(request *model.ListKmsByTagsRequest) *ListKmsByTagsInvoker {
	requestDef := GenReqDefForListKmsByTags()
	return &ListKmsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ListKmsTags(request *model.ListKmsTagsRequest) (*model.ListKmsTagsResponse, error) {
	requestDef := GenReqDefForListKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKmsTagsResponse), nil
	}
}

func (c *KmsClient) ListKmsTagsInvoker(request *model.ListKmsTagsRequest) *ListKmsTagsInvoker {
	requestDef := GenReqDefForListKmsTags()
	return &ListKmsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ListRetirableGrants(request *model.ListRetirableGrantsRequest) (*model.ListRetirableGrantsResponse, error) {
	requestDef := GenReqDefForListRetirableGrants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRetirableGrantsResponse), nil
	}
}

func (c *KmsClient) ListRetirableGrantsInvoker(request *model.ListRetirableGrantsRequest) *ListRetirableGrantsInvoker {
	requestDef := GenReqDefForListRetirableGrants()
	return &ListRetirableGrantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ShowKeyRotationStatus(request *model.ShowKeyRotationStatusRequest) (*model.ShowKeyRotationStatusResponse, error) {
	requestDef := GenReqDefForShowKeyRotationStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKeyRotationStatusResponse), nil
	}
}

func (c *KmsClient) ShowKeyRotationStatusInvoker(request *model.ShowKeyRotationStatusRequest) *ShowKeyRotationStatusInvoker {
	requestDef := GenReqDefForShowKeyRotationStatus()
	return &ShowKeyRotationStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ShowKeyStore(request *model.ShowKeyStoreRequest) (*model.ShowKeyStoreResponse, error) {
	requestDef := GenReqDefForShowKeyStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKeyStoreResponse), nil
	}
}

func (c *KmsClient) ShowKeyStoreInvoker(request *model.ShowKeyStoreRequest) *ShowKeyStoreInvoker {
	requestDef := GenReqDefForShowKeyStore()
	return &ShowKeyStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ShowKmsTags(request *model.ShowKmsTagsRequest) (*model.ShowKmsTagsResponse, error) {
	requestDef := GenReqDefForShowKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKmsTagsResponse), nil
	}
}

func (c *KmsClient) ShowKmsTagsInvoker(request *model.ShowKmsTagsRequest) *ShowKmsTagsInvoker {
	requestDef := GenReqDefForShowKmsTags()
	return &ShowKmsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ShowPublicKey(request *model.ShowPublicKeyRequest) (*model.ShowPublicKeyResponse, error) {
	requestDef := GenReqDefForShowPublicKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicKeyResponse), nil
	}
}

func (c *KmsClient) ShowPublicKeyInvoker(request *model.ShowPublicKeyRequest) *ShowPublicKeyInvoker {
	requestDef := GenReqDefForShowPublicKey()
	return &ShowPublicKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ShowUserInstances(request *model.ShowUserInstancesRequest) (*model.ShowUserInstancesResponse, error) {
	requestDef := GenReqDefForShowUserInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserInstancesResponse), nil
	}
}

func (c *KmsClient) ShowUserInstancesInvoker(request *model.ShowUserInstancesRequest) *ShowUserInstancesInvoker {
	requestDef := GenReqDefForShowUserInstances()
	return &ShowUserInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ShowUserQuotas(request *model.ShowUserQuotasRequest) (*model.ShowUserQuotasResponse, error) {
	requestDef := GenReqDefForShowUserQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserQuotasResponse), nil
	}
}

func (c *KmsClient) ShowUserQuotasInvoker(request *model.ShowUserQuotasRequest) *ShowUserQuotasInvoker {
	requestDef := GenReqDefForShowUserQuotas()
	return &ShowUserQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) Sign(request *model.SignRequest) (*model.SignResponse, error) {
	requestDef := GenReqDefForSign()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SignResponse), nil
	}
}

func (c *KmsClient) SignInvoker(request *model.SignRequest) *SignInvoker {
	requestDef := GenReqDefForSign()
	return &SignInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) UpdateKeyAlias(request *model.UpdateKeyAliasRequest) (*model.UpdateKeyAliasResponse, error) {
	requestDef := GenReqDefForUpdateKeyAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyAliasResponse), nil
	}
}

func (c *KmsClient) UpdateKeyAliasInvoker(request *model.UpdateKeyAliasRequest) *UpdateKeyAliasInvoker {
	requestDef := GenReqDefForUpdateKeyAlias()
	return &UpdateKeyAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) UpdateKeyDescription(request *model.UpdateKeyDescriptionRequest) (*model.UpdateKeyDescriptionResponse, error) {
	requestDef := GenReqDefForUpdateKeyDescription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyDescriptionResponse), nil
	}
}

func (c *KmsClient) UpdateKeyDescriptionInvoker(request *model.UpdateKeyDescriptionRequest) *UpdateKeyDescriptionInvoker {
	requestDef := GenReqDefForUpdateKeyDescription()
	return &UpdateKeyDescriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) UpdateKeyRotationInterval(request *model.UpdateKeyRotationIntervalRequest) (*model.UpdateKeyRotationIntervalResponse, error) {
	requestDef := GenReqDefForUpdateKeyRotationInterval()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyRotationIntervalResponse), nil
	}
}

func (c *KmsClient) UpdateKeyRotationIntervalInvoker(request *model.UpdateKeyRotationIntervalRequest) *UpdateKeyRotationIntervalInvoker {
	requestDef := GenReqDefForUpdateKeyRotationInterval()
	return &UpdateKeyRotationIntervalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ValidateSignature(request *model.ValidateSignatureRequest) (*model.ValidateSignatureResponse, error) {
	requestDef := GenReqDefForValidateSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateSignatureResponse), nil
	}
}

func (c *KmsClient) ValidateSignatureInvoker(request *model.ValidateSignatureRequest) *ValidateSignatureInvoker {
	requestDef := GenReqDefForValidateSignature()
	return &ValidateSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

func (c *KmsClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *KmsClient) ShowVersions(request *model.ShowVersionsRequest) (*model.ShowVersionsResponse, error) {
	requestDef := GenReqDefForShowVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionsResponse), nil
	}
}

func (c *KmsClient) ShowVersionsInvoker(request *model.ShowVersionsRequest) *ShowVersionsInvoker {
	requestDef := GenReqDefForShowVersions()
	return &ShowVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
