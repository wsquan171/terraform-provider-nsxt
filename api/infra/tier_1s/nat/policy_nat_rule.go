//nolint:revive
package nat

// The following file has been autogenerated. Please avoid any changes!
import (
	"errors"

	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	client1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/global_infra/tier_1s/nat"
	model1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	client0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/tier_1s/nat"
	model0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	client2 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/orgs/projects/infra/tier_1s/nat"

	utl "github.com/vmware/terraform-provider-nsxt/api/utl"
)

type PolicyNatRuleClientContext utl.ClientContext

func NewNatRulesClient(sessionContext utl.SessionContext, connector vapiProtocolClient_.Connector) *PolicyNatRuleClientContext {
	var client interface{}

	switch sessionContext.ClientType {

	case utl.Local:
		client = client0.NewNatRulesClient(connector)

	case utl.Global:
		client = client1.NewNatRulesClient(connector)

	case utl.Multitenancy:
		client = client2.NewNatRulesClient(connector)

	default:
		return nil
	}
	return &PolicyNatRuleClientContext{Client: client, ClientType: sessionContext.ClientType, ProjectID: sessionContext.ProjectID}
}

func (c PolicyNatRuleClientContext) Get(tier1IdParam string, natIdParam string, natRuleIdParam string) (model0.PolicyNatRule, error) {
	var obj model0.PolicyNatRule
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.NatRulesClient)
		obj, err = client.Get(tier1IdParam, natIdParam, natRuleIdParam)
		if err != nil {
			return obj, err
		}

	case utl.Global:
		client := c.Client.(client1.NatRulesClient)
		gmObj, err1 := client.Get(tier1IdParam, natIdParam, natRuleIdParam)
		if err1 != nil {
			return obj, err1
		}
		var rawObj interface{}
		rawObj, err = utl.ConvertModelBindingType(gmObj, model1.PolicyNatRuleBindingType(), model0.PolicyNatRuleBindingType())
		obj = rawObj.(model0.PolicyNatRule)

	case utl.Multitenancy:
		client := c.Client.(client2.NatRulesClient)
		obj, err = client.Get(utl.DefaultOrgID, c.ProjectID, tier1IdParam, natIdParam, natRuleIdParam)
		if err != nil {
			return obj, err
		}

	default:
		return obj, errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c PolicyNatRuleClientContext) Patch(tier1IdParam string, natIdParam string, natRuleIdParam string, policyNatRuleParam model0.PolicyNatRule) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.NatRulesClient)
		err = client.Patch(tier1IdParam, natIdParam, natRuleIdParam, policyNatRuleParam)

	case utl.Global:
		client := c.Client.(client1.NatRulesClient)
		gmObj, err1 := utl.ConvertModelBindingType(policyNatRuleParam, model0.PolicyNatRuleBindingType(), model1.PolicyNatRuleBindingType())
		if err1 != nil {
			return err1
		}
		err = client.Patch(tier1IdParam, natIdParam, natRuleIdParam, gmObj.(model1.PolicyNatRule))

	case utl.Multitenancy:
		client := c.Client.(client2.NatRulesClient)
		err = client.Patch(utl.DefaultOrgID, c.ProjectID, tier1IdParam, natIdParam, natRuleIdParam, policyNatRuleParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c PolicyNatRuleClientContext) Update(tier1IdParam string, natIdParam string, natRuleIdParam string, policyNatRuleParam model0.PolicyNatRule) (model0.PolicyNatRule, error) {
	var err error
	var obj model0.PolicyNatRule

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.NatRulesClient)
		obj, err = client.Update(tier1IdParam, natIdParam, natRuleIdParam, policyNatRuleParam)

	case utl.Global:
		client := c.Client.(client1.NatRulesClient)
		gmObj, err := utl.ConvertModelBindingType(policyNatRuleParam, model0.PolicyNatRuleBindingType(), model1.PolicyNatRuleBindingType())
		if err != nil {
			return obj, err
		}
		gmObj, err = client.Update(tier1IdParam, natIdParam, natRuleIdParam, gmObj.(model1.PolicyNatRule))
		if err != nil {
			return obj, err
		}
		obj1, err1 := utl.ConvertModelBindingType(gmObj, model1.PolicyNatRuleBindingType(), model0.PolicyNatRuleBindingType())
		if err1 != nil {
			return obj, err1
		}
		obj = obj1.(model0.PolicyNatRule)

	case utl.Multitenancy:
		client := c.Client.(client2.NatRulesClient)
		obj, err = client.Update(utl.DefaultOrgID, c.ProjectID, tier1IdParam, natIdParam, natRuleIdParam, policyNatRuleParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c PolicyNatRuleClientContext) Delete(tier1IdParam string, natIdParam string, natRuleIdParam string) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.NatRulesClient)
		err = client.Delete(tier1IdParam, natIdParam, natRuleIdParam)

	case utl.Global:
		client := c.Client.(client1.NatRulesClient)
		err = client.Delete(tier1IdParam, natIdParam, natRuleIdParam)

	case utl.Multitenancy:
		client := c.Client.(client2.NatRulesClient)
		err = client.Delete(utl.DefaultOrgID, c.ProjectID, tier1IdParam, natIdParam, natRuleIdParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c PolicyNatRuleClientContext) List(tier1IdParam string, natIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model0.PolicyNatRuleListResult, error) {
	var err error
	var obj model0.PolicyNatRuleListResult

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.NatRulesClient)
		obj, err = client.List(tier1IdParam, natIdParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)

	case utl.Global:
		client := c.Client.(client1.NatRulesClient)
		gmObj, err := client.List(tier1IdParam, natIdParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)
		if err != nil {
			return obj, err
		}
		obj1, err1 := utl.ConvertModelBindingType(gmObj, model1.PolicyNatRuleListResultBindingType(), model0.PolicyNatRuleListResultBindingType())
		if err1 != nil {
			return obj, err1
		}
		obj = obj1.(model0.PolicyNatRuleListResult)

	case utl.Multitenancy:
		client := c.Client.(client2.NatRulesClient)
		obj, err = client.List(utl.DefaultOrgID, c.ProjectID, tier1IdParam, natIdParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}