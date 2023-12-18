package v2

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/ims/v2/model"
)

type ImsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewImsClient(hcClient *http_client.HcHttpClient) *ImsClient {
	return &ImsClient{HcClient: hcClient}
}

func ImsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *ImsClient) AddImageTag(request *model.AddImageTagRequest) (*model.AddImageTagResponse, error) {
	requestDef := GenReqDefForAddImageTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddImageTagResponse), nil
	}
}

func (c *ImsClient) AddImageTagInvoker(request *model.AddImageTagRequest) *AddImageTagInvoker {
	requestDef := GenReqDefForAddImageTag()
	return &AddImageTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) BatchAddMembers(request *model.BatchAddMembersRequest) (*model.BatchAddMembersResponse, error) {
	requestDef := GenReqDefForBatchAddMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddMembersResponse), nil
	}
}

func (c *ImsClient) BatchAddMembersInvoker(request *model.BatchAddMembersRequest) *BatchAddMembersInvoker {
	requestDef := GenReqDefForBatchAddMembers()
	return &BatchAddMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) BatchAddOrDeleteTags(request *model.BatchAddOrDeleteTagsRequest) (*model.BatchAddOrDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchAddOrDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddOrDeleteTagsResponse), nil
	}
}

func (c *ImsClient) BatchAddOrDeleteTagsInvoker(request *model.BatchAddOrDeleteTagsRequest) *BatchAddOrDeleteTagsInvoker {
	requestDef := GenReqDefForBatchAddOrDeleteTags()
	return &BatchAddOrDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) BatchDeleteMembers(request *model.BatchDeleteMembersRequest) (*model.BatchDeleteMembersResponse, error) {
	requestDef := GenReqDefForBatchDeleteMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMembersResponse), nil
	}
}

func (c *ImsClient) BatchDeleteMembersInvoker(request *model.BatchDeleteMembersRequest) *BatchDeleteMembersInvoker {
	requestDef := GenReqDefForBatchDeleteMembers()
	return &BatchDeleteMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) BatchUpdateMembers(request *model.BatchUpdateMembersRequest) (*model.BatchUpdateMembersResponse, error) {
	requestDef := GenReqDefForBatchUpdateMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateMembersResponse), nil
	}
}

func (c *ImsClient) BatchUpdateMembersInvoker(request *model.BatchUpdateMembersRequest) *BatchUpdateMembersInvoker {
	requestDef := GenReqDefForBatchUpdateMembers()
	return &BatchUpdateMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) CopyImageCrossRegion(request *model.CopyImageCrossRegionRequest) (*model.CopyImageCrossRegionResponse, error) {
	requestDef := GenReqDefForCopyImageCrossRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyImageCrossRegionResponse), nil
	}
}

func (c *ImsClient) CopyImageCrossRegionInvoker(request *model.CopyImageCrossRegionRequest) *CopyImageCrossRegionInvoker {
	requestDef := GenReqDefForCopyImageCrossRegion()
	return &CopyImageCrossRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) CopyImageInRegion(request *model.CopyImageInRegionRequest) (*model.CopyImageInRegionResponse, error) {
	requestDef := GenReqDefForCopyImageInRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyImageInRegionResponse), nil
	}
}

func (c *ImsClient) CopyImageInRegionInvoker(request *model.CopyImageInRegionRequest) *CopyImageInRegionInvoker {
	requestDef := GenReqDefForCopyImageInRegion()
	return &CopyImageInRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) CreateDataImage(request *model.CreateDataImageRequest) (*model.CreateDataImageResponse, error) {
	requestDef := GenReqDefForCreateDataImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataImageResponse), nil
	}
}

func (c *ImsClient) CreateDataImageInvoker(request *model.CreateDataImageRequest) *CreateDataImageInvoker {
	requestDef := GenReqDefForCreateDataImage()
	return &CreateDataImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) CreateImage(request *model.CreateImageRequest) (*model.CreateImageResponse, error) {
	requestDef := GenReqDefForCreateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageResponse), nil
	}
}

func (c *ImsClient) CreateImageInvoker(request *model.CreateImageRequest) *CreateImageInvoker {
	requestDef := GenReqDefForCreateImage()
	return &CreateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) CreateOrUpdateTags(request *model.CreateOrUpdateTagsRequest) (*model.CreateOrUpdateTagsResponse, error) {
	requestDef := GenReqDefForCreateOrUpdateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrUpdateTagsResponse), nil
	}
}

func (c *ImsClient) CreateOrUpdateTagsInvoker(request *model.CreateOrUpdateTagsRequest) *CreateOrUpdateTagsInvoker {
	requestDef := GenReqDefForCreateOrUpdateTags()
	return &CreateOrUpdateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) CreateWholeImage(request *model.CreateWholeImageRequest) (*model.CreateWholeImageResponse, error) {
	requestDef := GenReqDefForCreateWholeImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWholeImageResponse), nil
	}
}

func (c *ImsClient) CreateWholeImageInvoker(request *model.CreateWholeImageRequest) *CreateWholeImageInvoker {
	requestDef := GenReqDefForCreateWholeImage()
	return &CreateWholeImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) DeleteImageTag(request *model.DeleteImageTagRequest) (*model.DeleteImageTagResponse, error) {
	requestDef := GenReqDefForDeleteImageTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageTagResponse), nil
	}
}

func (c *ImsClient) DeleteImageTagInvoker(request *model.DeleteImageTagRequest) *DeleteImageTagInvoker {
	requestDef := GenReqDefForDeleteImageTag()
	return &DeleteImageTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ExportImage(request *model.ExportImageRequest) (*model.ExportImageResponse, error) {
	requestDef := GenReqDefForExportImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportImageResponse), nil
	}
}

func (c *ImsClient) ExportImageInvoker(request *model.ExportImageRequest) *ExportImageInvoker {
	requestDef := GenReqDefForExportImage()
	return &ExportImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ImportImageQuick(request *model.ImportImageQuickRequest) (*model.ImportImageQuickResponse, error) {
	requestDef := GenReqDefForImportImageQuick()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportImageQuickResponse), nil
	}
}

func (c *ImsClient) ImportImageQuickInvoker(request *model.ImportImageQuickRequest) *ImportImageQuickInvoker {
	requestDef := GenReqDefForImportImageQuick()
	return &ImportImageQuickInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ListImageByTags(request *model.ListImageByTagsRequest) (*model.ListImageByTagsResponse, error) {
	requestDef := GenReqDefForListImageByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageByTagsResponse), nil
	}
}

func (c *ImsClient) ListImageByTagsInvoker(request *model.ListImageByTagsRequest) *ListImageByTagsInvoker {
	requestDef := GenReqDefForListImageByTags()
	return &ListImageByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ListImageTags(request *model.ListImageTagsRequest) (*model.ListImageTagsResponse, error) {
	requestDef := GenReqDefForListImageTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageTagsResponse), nil
	}
}

func (c *ImsClient) ListImageTagsInvoker(request *model.ListImageTagsRequest) *ListImageTagsInvoker {
	requestDef := GenReqDefForListImageTags()
	return &ListImageTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ListImages(request *model.ListImagesRequest) (*model.ListImagesResponse, error) {
	requestDef := GenReqDefForListImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImagesResponse), nil
	}
}

func (c *ImsClient) ListImagesInvoker(request *model.ListImagesRequest) *ListImagesInvoker {
	requestDef := GenReqDefForListImages()
	return &ListImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ListImagesTags(request *model.ListImagesTagsRequest) (*model.ListImagesTagsResponse, error) {
	requestDef := GenReqDefForListImagesTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImagesTagsResponse), nil
	}
}

func (c *ImsClient) ListImagesTagsInvoker(request *model.ListImagesTagsRequest) *ListImagesTagsInvoker {
	requestDef := GenReqDefForListImagesTags()
	return &ListImagesTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ListOsVersions(request *model.ListOsVersionsRequest) (*model.ListOsVersionsResponse, error) {
	requestDef := GenReqDefForListOsVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOsVersionsResponse), nil
	}
}

func (c *ImsClient) ListOsVersionsInvoker(request *model.ListOsVersionsRequest) *ListOsVersionsInvoker {
	requestDef := GenReqDefForListOsVersions()
	return &ListOsVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

func (c *ImsClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) RegisterImage(request *model.RegisterImageRequest) (*model.RegisterImageResponse, error) {
	requestDef := GenReqDefForRegisterImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterImageResponse), nil
	}
}

func (c *ImsClient) RegisterImageInvoker(request *model.RegisterImageRequest) *RegisterImageInvoker {
	requestDef := GenReqDefForRegisterImage()
	return &RegisterImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ShowImageQuota(request *model.ShowImageQuotaRequest) (*model.ShowImageQuotaResponse, error) {
	requestDef := GenReqDefForShowImageQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageQuotaResponse), nil
	}
}

func (c *ImsClient) ShowImageQuotaInvoker(request *model.ShowImageQuotaRequest) *ShowImageQuotaInvoker {
	requestDef := GenReqDefForShowImageQuota()
	return &ShowImageQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

func (c *ImsClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ShowJobProgress(request *model.ShowJobProgressRequest) (*model.ShowJobProgressResponse, error) {
	requestDef := GenReqDefForShowJobProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobProgressResponse), nil
	}
}

func (c *ImsClient) ShowJobProgressInvoker(request *model.ShowJobProgressRequest) *ShowJobProgressInvoker {
	requestDef := GenReqDefForShowJobProgress()
	return &ShowJobProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) UpdateImage(request *model.UpdateImageRequest) (*model.UpdateImageResponse, error) {
	requestDef := GenReqDefForUpdateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateImageResponse), nil
	}
}

func (c *ImsClient) UpdateImageInvoker(request *model.UpdateImageRequest) *UpdateImageInvoker {
	requestDef := GenReqDefForUpdateImage()
	return &UpdateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ListVersions(request *model.ListVersionsRequest) (*model.ListVersionsResponse, error) {
	requestDef := GenReqDefForListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionsResponse), nil
	}
}

func (c *ImsClient) ListVersionsInvoker(request *model.ListVersionsRequest) *ListVersionsInvoker {
	requestDef := GenReqDefForListVersions()
	return &ListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

func (c *ImsClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceAddImageMember(request *model.GlanceAddImageMemberRequest) (*model.GlanceAddImageMemberResponse, error) {
	requestDef := GenReqDefForGlanceAddImageMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceAddImageMemberResponse), nil
	}
}

func (c *ImsClient) GlanceAddImageMemberInvoker(request *model.GlanceAddImageMemberRequest) *GlanceAddImageMemberInvoker {
	requestDef := GenReqDefForGlanceAddImageMember()
	return &GlanceAddImageMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceCreateImageMetadata(request *model.GlanceCreateImageMetadataRequest) (*model.GlanceCreateImageMetadataResponse, error) {
	requestDef := GenReqDefForGlanceCreateImageMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceCreateImageMetadataResponse), nil
	}
}

func (c *ImsClient) GlanceCreateImageMetadataInvoker(request *model.GlanceCreateImageMetadataRequest) *GlanceCreateImageMetadataInvoker {
	requestDef := GenReqDefForGlanceCreateImageMetadata()
	return &GlanceCreateImageMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceCreateTag(request *model.GlanceCreateTagRequest) (*model.GlanceCreateTagResponse, error) {
	requestDef := GenReqDefForGlanceCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceCreateTagResponse), nil
	}
}

func (c *ImsClient) GlanceCreateTagInvoker(request *model.GlanceCreateTagRequest) *GlanceCreateTagInvoker {
	requestDef := GenReqDefForGlanceCreateTag()
	return &GlanceCreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceDeleteImage(request *model.GlanceDeleteImageRequest) (*model.GlanceDeleteImageResponse, error) {
	requestDef := GenReqDefForGlanceDeleteImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceDeleteImageResponse), nil
	}
}

func (c *ImsClient) GlanceDeleteImageInvoker(request *model.GlanceDeleteImageRequest) *GlanceDeleteImageInvoker {
	requestDef := GenReqDefForGlanceDeleteImage()
	return &GlanceDeleteImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceDeleteImageMember(request *model.GlanceDeleteImageMemberRequest) (*model.GlanceDeleteImageMemberResponse, error) {
	requestDef := GenReqDefForGlanceDeleteImageMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceDeleteImageMemberResponse), nil
	}
}

func (c *ImsClient) GlanceDeleteImageMemberInvoker(request *model.GlanceDeleteImageMemberRequest) *GlanceDeleteImageMemberInvoker {
	requestDef := GenReqDefForGlanceDeleteImageMember()
	return &GlanceDeleteImageMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceDeleteTag(request *model.GlanceDeleteTagRequest) (*model.GlanceDeleteTagResponse, error) {
	requestDef := GenReqDefForGlanceDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceDeleteTagResponse), nil
	}
}

func (c *ImsClient) GlanceDeleteTagInvoker(request *model.GlanceDeleteTagRequest) *GlanceDeleteTagInvoker {
	requestDef := GenReqDefForGlanceDeleteTag()
	return &GlanceDeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceListImageMemberSchemas(request *model.GlanceListImageMemberSchemasRequest) (*model.GlanceListImageMemberSchemasResponse, error) {
	requestDef := GenReqDefForGlanceListImageMemberSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceListImageMemberSchemasResponse), nil
	}
}

func (c *ImsClient) GlanceListImageMemberSchemasInvoker(request *model.GlanceListImageMemberSchemasRequest) *GlanceListImageMemberSchemasInvoker {
	requestDef := GenReqDefForGlanceListImageMemberSchemas()
	return &GlanceListImageMemberSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceListImageMembers(request *model.GlanceListImageMembersRequest) (*model.GlanceListImageMembersResponse, error) {
	requestDef := GenReqDefForGlanceListImageMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceListImageMembersResponse), nil
	}
}

func (c *ImsClient) GlanceListImageMembersInvoker(request *model.GlanceListImageMembersRequest) *GlanceListImageMembersInvoker {
	requestDef := GenReqDefForGlanceListImageMembers()
	return &GlanceListImageMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceListImageSchemas(request *model.GlanceListImageSchemasRequest) (*model.GlanceListImageSchemasResponse, error) {
	requestDef := GenReqDefForGlanceListImageSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceListImageSchemasResponse), nil
	}
}

func (c *ImsClient) GlanceListImageSchemasInvoker(request *model.GlanceListImageSchemasRequest) *GlanceListImageSchemasInvoker {
	requestDef := GenReqDefForGlanceListImageSchemas()
	return &GlanceListImageSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceListImages(request *model.GlanceListImagesRequest) (*model.GlanceListImagesResponse, error) {
	requestDef := GenReqDefForGlanceListImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceListImagesResponse), nil
	}
}

func (c *ImsClient) GlanceListImagesInvoker(request *model.GlanceListImagesRequest) *GlanceListImagesInvoker {
	requestDef := GenReqDefForGlanceListImages()
	return &GlanceListImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceShowImage(request *model.GlanceShowImageRequest) (*model.GlanceShowImageResponse, error) {
	requestDef := GenReqDefForGlanceShowImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceShowImageResponse), nil
	}
}

func (c *ImsClient) GlanceShowImageInvoker(request *model.GlanceShowImageRequest) *GlanceShowImageInvoker {
	requestDef := GenReqDefForGlanceShowImage()
	return &GlanceShowImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceShowImageMember(request *model.GlanceShowImageMemberRequest) (*model.GlanceShowImageMemberResponse, error) {
	requestDef := GenReqDefForGlanceShowImageMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceShowImageMemberResponse), nil
	}
}

func (c *ImsClient) GlanceShowImageMemberInvoker(request *model.GlanceShowImageMemberRequest) *GlanceShowImageMemberInvoker {
	requestDef := GenReqDefForGlanceShowImageMember()
	return &GlanceShowImageMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceShowImageMemberSchemas(request *model.GlanceShowImageMemberSchemasRequest) (*model.GlanceShowImageMemberSchemasResponse, error) {
	requestDef := GenReqDefForGlanceShowImageMemberSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceShowImageMemberSchemasResponse), nil
	}
}

func (c *ImsClient) GlanceShowImageMemberSchemasInvoker(request *model.GlanceShowImageMemberSchemasRequest) *GlanceShowImageMemberSchemasInvoker {
	requestDef := GenReqDefForGlanceShowImageMemberSchemas()
	return &GlanceShowImageMemberSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceShowImageSchemas(request *model.GlanceShowImageSchemasRequest) (*model.GlanceShowImageSchemasResponse, error) {
	requestDef := GenReqDefForGlanceShowImageSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceShowImageSchemasResponse), nil
	}
}

func (c *ImsClient) GlanceShowImageSchemasInvoker(request *model.GlanceShowImageSchemasRequest) *GlanceShowImageSchemasInvoker {
	requestDef := GenReqDefForGlanceShowImageSchemas()
	return &GlanceShowImageSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceUpdateImage(request *model.GlanceUpdateImageRequest) (*model.GlanceUpdateImageResponse, error) {
	requestDef := GenReqDefForGlanceUpdateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceUpdateImageResponse), nil
	}
}

func (c *ImsClient) GlanceUpdateImageInvoker(request *model.GlanceUpdateImageRequest) *GlanceUpdateImageInvoker {
	requestDef := GenReqDefForGlanceUpdateImage()
	return &GlanceUpdateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ImsClient) GlanceUpdateImageMember(request *model.GlanceUpdateImageMemberRequest) (*model.GlanceUpdateImageMemberResponse, error) {
	requestDef := GenReqDefForGlanceUpdateImageMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GlanceUpdateImageMemberResponse), nil
	}
}

func (c *ImsClient) GlanceUpdateImageMemberInvoker(request *model.GlanceUpdateImageMemberRequest) *GlanceUpdateImageMemberInvoker {
	requestDef := GenReqDefForGlanceUpdateImageMember()
	return &GlanceUpdateImageMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
