package v2

import (
	http_client "github.com/tmcloud-sdk/tmcloud-sdk-go/core"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/invoker"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/services/nat/v2/model"
)

type NatClient struct {
	HcClient *http_client.HcHttpClient
}

func NewNatClient(hcClient *http_client.HcHttpClient) *NatClient {
	return &NatClient{HcClient: hcClient}
}

func NatClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *NatClient) BatchCreateNatGatewayDnatRules(request *model.BatchCreateNatGatewayDnatRulesRequest) (*model.BatchCreateNatGatewayDnatRulesResponse, error) {
	requestDef := GenReqDefForBatchCreateNatGatewayDnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateNatGatewayDnatRulesResponse), nil
	}
}

func (c *NatClient) BatchCreateNatGatewayDnatRulesInvoker(request *model.BatchCreateNatGatewayDnatRulesRequest) *BatchCreateNatGatewayDnatRulesInvoker {
	requestDef := GenReqDefForBatchCreateNatGatewayDnatRules()
	return &BatchCreateNatGatewayDnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreateNatGatewayDnatRule(request *model.CreateNatGatewayDnatRuleRequest) (*model.CreateNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForCreateNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewayDnatRuleResponse), nil
	}
}

func (c *NatClient) CreateNatGatewayDnatRuleInvoker(request *model.CreateNatGatewayDnatRuleRequest) *CreateNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForCreateNatGatewayDnatRule()
	return &CreateNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeleteNatGatewayDnatRule(request *model.DeleteNatGatewayDnatRuleRequest) (*model.DeleteNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForDeleteNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewayDnatRuleResponse), nil
	}
}

func (c *NatClient) DeleteNatGatewayDnatRuleInvoker(request *model.DeleteNatGatewayDnatRuleRequest) *DeleteNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForDeleteNatGatewayDnatRule()
	return &DeleteNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListNatGatewayDnatRules(request *model.ListNatGatewayDnatRulesRequest) (*model.ListNatGatewayDnatRulesResponse, error) {
	requestDef := GenReqDefForListNatGatewayDnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewayDnatRulesResponse), nil
	}
}

func (c *NatClient) ListNatGatewayDnatRulesInvoker(request *model.ListNatGatewayDnatRulesRequest) *ListNatGatewayDnatRulesInvoker {
	requestDef := GenReqDefForListNatGatewayDnatRules()
	return &ListNatGatewayDnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowNatGatewayDnatRule(request *model.ShowNatGatewayDnatRuleRequest) (*model.ShowNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForShowNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewayDnatRuleResponse), nil
	}
}

func (c *NatClient) ShowNatGatewayDnatRuleInvoker(request *model.ShowNatGatewayDnatRuleRequest) *ShowNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForShowNatGatewayDnatRule()
	return &ShowNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdateNatGatewayDnatRule(request *model.UpdateNatGatewayDnatRuleRequest) (*model.UpdateNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForUpdateNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewayDnatRuleResponse), nil
	}
}

func (c *NatClient) UpdateNatGatewayDnatRuleInvoker(request *model.UpdateNatGatewayDnatRuleRequest) *UpdateNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForUpdateNatGatewayDnatRule()
	return &UpdateNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreateNatGateway(request *model.CreateNatGatewayRequest) (*model.CreateNatGatewayResponse, error) {
	requestDef := GenReqDefForCreateNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewayResponse), nil
	}
}

func (c *NatClient) CreateNatGatewayInvoker(request *model.CreateNatGatewayRequest) *CreateNatGatewayInvoker {
	requestDef := GenReqDefForCreateNatGateway()
	return &CreateNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeleteNatGateway(request *model.DeleteNatGatewayRequest) (*model.DeleteNatGatewayResponse, error) {
	requestDef := GenReqDefForDeleteNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewayResponse), nil
	}
}

func (c *NatClient) DeleteNatGatewayInvoker(request *model.DeleteNatGatewayRequest) *DeleteNatGatewayInvoker {
	requestDef := GenReqDefForDeleteNatGateway()
	return &DeleteNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListNatGateways(request *model.ListNatGatewaysRequest) (*model.ListNatGatewaysResponse, error) {
	requestDef := GenReqDefForListNatGateways()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewaysResponse), nil
	}
}

func (c *NatClient) ListNatGatewaysInvoker(request *model.ListNatGatewaysRequest) *ListNatGatewaysInvoker {
	requestDef := GenReqDefForListNatGateways()
	return &ListNatGatewaysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowNatGateway(request *model.ShowNatGatewayRequest) (*model.ShowNatGatewayResponse, error) {
	requestDef := GenReqDefForShowNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewayResponse), nil
	}
}

func (c *NatClient) ShowNatGatewayInvoker(request *model.ShowNatGatewayRequest) *ShowNatGatewayInvoker {
	requestDef := GenReqDefForShowNatGateway()
	return &ShowNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdateNatGateway(request *model.UpdateNatGatewayRequest) (*model.UpdateNatGatewayResponse, error) {
	requestDef := GenReqDefForUpdateNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewayResponse), nil
	}
}

func (c *NatClient) UpdateNatGatewayInvoker(request *model.UpdateNatGatewayRequest) *UpdateNatGatewayInvoker {
	requestDef := GenReqDefForUpdateNatGateway()
	return &UpdateNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreateNatGatewaySnatRule(request *model.CreateNatGatewaySnatRuleRequest) (*model.CreateNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForCreateNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewaySnatRuleResponse), nil
	}
}

func (c *NatClient) CreateNatGatewaySnatRuleInvoker(request *model.CreateNatGatewaySnatRuleRequest) *CreateNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForCreateNatGatewaySnatRule()
	return &CreateNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeleteNatGatewaySnatRule(request *model.DeleteNatGatewaySnatRuleRequest) (*model.DeleteNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForDeleteNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewaySnatRuleResponse), nil
	}
}

func (c *NatClient) DeleteNatGatewaySnatRuleInvoker(request *model.DeleteNatGatewaySnatRuleRequest) *DeleteNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForDeleteNatGatewaySnatRule()
	return &DeleteNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListNatGatewaySnatRules(request *model.ListNatGatewaySnatRulesRequest) (*model.ListNatGatewaySnatRulesResponse, error) {
	requestDef := GenReqDefForListNatGatewaySnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewaySnatRulesResponse), nil
	}
}

func (c *NatClient) ListNatGatewaySnatRulesInvoker(request *model.ListNatGatewaySnatRulesRequest) *ListNatGatewaySnatRulesInvoker {
	requestDef := GenReqDefForListNatGatewaySnatRules()
	return &ListNatGatewaySnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowNatGatewaySnatRule(request *model.ShowNatGatewaySnatRuleRequest) (*model.ShowNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForShowNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewaySnatRuleResponse), nil
	}
}

func (c *NatClient) ShowNatGatewaySnatRuleInvoker(request *model.ShowNatGatewaySnatRuleRequest) *ShowNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForShowNatGatewaySnatRule()
	return &ShowNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdateNatGatewaySnatRule(request *model.UpdateNatGatewaySnatRuleRequest) (*model.UpdateNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForUpdateNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewaySnatRuleResponse), nil
	}
}

func (c *NatClient) UpdateNatGatewaySnatRuleInvoker(request *model.UpdateNatGatewaySnatRuleRequest) *UpdateNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForUpdateNatGatewaySnatRule()
	return &UpdateNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
