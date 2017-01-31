/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package mockcloudformation

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"

	request "github.com/aws/aws-sdk-go/aws/request"
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
        "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"

)

// CloudFormationAPIMock mock
type CloudFormationAPIMock struct {
	mock.Mock
	cloudformationiface.CloudFormationAPI  // temp workaround for https://github.com/ernesto-jimenez/goautomock/issues/3#issuecomment-276240746
}

func NewCloudFormationAPIMock() *CloudFormationAPIMock {
	return &CloudFormationAPIMock{}
}

// CancelUpdateStack mocked method
func (m *CloudFormationAPIMock) CancelUpdateStack(p0 *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.CancelUpdateStackOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.CancelUpdateStackOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CancelUpdateStackRequest mocked method
func (m *CloudFormationAPIMock) CancelUpdateStackRequest(p0 *cloudformation.CancelUpdateStackInput) (*request.Request, *cloudformation.CancelUpdateStackOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.CancelUpdateStackOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.CancelUpdateStackOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ContinueUpdateRollback mocked method
func (m *CloudFormationAPIMock) ContinueUpdateRollback(p0 *cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.ContinueUpdateRollbackOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.ContinueUpdateRollbackOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ContinueUpdateRollbackRequest mocked method
func (m *CloudFormationAPIMock) ContinueUpdateRollbackRequest(p0 *cloudformation.ContinueUpdateRollbackInput) (*request.Request, *cloudformation.ContinueUpdateRollbackOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.ContinueUpdateRollbackOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.ContinueUpdateRollbackOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateChangeSet mocked method
func (m *CloudFormationAPIMock) CreateChangeSet(p0 *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.CreateChangeSetOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.CreateChangeSetOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateChangeSetRequest mocked method
func (m *CloudFormationAPIMock) CreateChangeSetRequest(p0 *cloudformation.CreateChangeSetInput) (*request.Request, *cloudformation.CreateChangeSetOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.CreateChangeSetOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.CreateChangeSetOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateStack mocked method
func (m *CloudFormationAPIMock) CreateStack(p0 *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.CreateStackOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.CreateStackOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateStackRequest mocked method
func (m *CloudFormationAPIMock) CreateStackRequest(p0 *cloudformation.CreateStackInput) (*request.Request, *cloudformation.CreateStackOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.CreateStackOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.CreateStackOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteChangeSet mocked method
func (m *CloudFormationAPIMock) DeleteChangeSet(p0 *cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.DeleteChangeSetOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.DeleteChangeSetOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteChangeSetRequest mocked method
func (m *CloudFormationAPIMock) DeleteChangeSetRequest(p0 *cloudformation.DeleteChangeSetInput) (*request.Request, *cloudformation.DeleteChangeSetOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.DeleteChangeSetOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.DeleteChangeSetOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteStack mocked method
func (m *CloudFormationAPIMock) DeleteStack(p0 *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.DeleteStackOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.DeleteStackOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteStackRequest mocked method
func (m *CloudFormationAPIMock) DeleteStackRequest(p0 *cloudformation.DeleteStackInput) (*request.Request, *cloudformation.DeleteStackOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.DeleteStackOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.DeleteStackOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeAccountLimits mocked method
func (m *CloudFormationAPIMock) DescribeAccountLimits(p0 *cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.DescribeAccountLimitsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.DescribeAccountLimitsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeAccountLimitsRequest mocked method
func (m *CloudFormationAPIMock) DescribeAccountLimitsRequest(p0 *cloudformation.DescribeAccountLimitsInput) (*request.Request, *cloudformation.DescribeAccountLimitsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.DescribeAccountLimitsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.DescribeAccountLimitsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeChangeSet mocked method
func (m *CloudFormationAPIMock) DescribeChangeSet(p0 *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.DescribeChangeSetOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.DescribeChangeSetOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeChangeSetRequest mocked method
func (m *CloudFormationAPIMock) DescribeChangeSetRequest(p0 *cloudformation.DescribeChangeSetInput) (*request.Request, *cloudformation.DescribeChangeSetOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.DescribeChangeSetOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.DescribeChangeSetOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStackEvents mocked method
func (m *CloudFormationAPIMock) DescribeStackEvents(p0 *cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.DescribeStackEventsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.DescribeStackEventsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStackEventsPages mocked method
func (m *CloudFormationAPIMock) DescribeStackEventsPages(p0 *cloudformation.DescribeStackEventsInput, p1 func(*cloudformation.DescribeStackEventsOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// DescribeStackEventsRequest mocked method
func (m *CloudFormationAPIMock) DescribeStackEventsRequest(p0 *cloudformation.DescribeStackEventsInput) (*request.Request, *cloudformation.DescribeStackEventsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.DescribeStackEventsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.DescribeStackEventsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStackResource mocked method
func (m *CloudFormationAPIMock) DescribeStackResource(p0 *cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.DescribeStackResourceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.DescribeStackResourceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStackResourceRequest mocked method
func (m *CloudFormationAPIMock) DescribeStackResourceRequest(p0 *cloudformation.DescribeStackResourceInput) (*request.Request, *cloudformation.DescribeStackResourceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.DescribeStackResourceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.DescribeStackResourceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStackResources mocked method
func (m *CloudFormationAPIMock) DescribeStackResources(p0 *cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.DescribeStackResourcesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.DescribeStackResourcesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStackResourcesRequest mocked method
func (m *CloudFormationAPIMock) DescribeStackResourcesRequest(p0 *cloudformation.DescribeStackResourcesInput) (*request.Request, *cloudformation.DescribeStackResourcesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.DescribeStackResourcesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.DescribeStackResourcesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStacks mocked method
func (m *CloudFormationAPIMock) DescribeStacks(p0 *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.DescribeStacksOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.DescribeStacksOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStacksPages mocked method
func (m *CloudFormationAPIMock) DescribeStacksPages(p0 *cloudformation.DescribeStacksInput, p1 func(*cloudformation.DescribeStacksOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// DescribeStacksRequest mocked method
func (m *CloudFormationAPIMock) DescribeStacksRequest(p0 *cloudformation.DescribeStacksInput) (*request.Request, *cloudformation.DescribeStacksOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.DescribeStacksOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.DescribeStacksOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EstimateTemplateCost mocked method
func (m *CloudFormationAPIMock) EstimateTemplateCost(p0 *cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.EstimateTemplateCostOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.EstimateTemplateCostOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EstimateTemplateCostRequest mocked method
func (m *CloudFormationAPIMock) EstimateTemplateCostRequest(p0 *cloudformation.EstimateTemplateCostInput) (*request.Request, *cloudformation.EstimateTemplateCostOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.EstimateTemplateCostOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.EstimateTemplateCostOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ExecuteChangeSet mocked method
func (m *CloudFormationAPIMock) ExecuteChangeSet(p0 *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.ExecuteChangeSetOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.ExecuteChangeSetOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ExecuteChangeSetRequest mocked method
func (m *CloudFormationAPIMock) ExecuteChangeSetRequest(p0 *cloudformation.ExecuteChangeSetInput) (*request.Request, *cloudformation.ExecuteChangeSetOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.ExecuteChangeSetOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.ExecuteChangeSetOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetStackPolicy mocked method
func (m *CloudFormationAPIMock) GetStackPolicy(p0 *cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.GetStackPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.GetStackPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetStackPolicyRequest mocked method
func (m *CloudFormationAPIMock) GetStackPolicyRequest(p0 *cloudformation.GetStackPolicyInput) (*request.Request, *cloudformation.GetStackPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.GetStackPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.GetStackPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetTemplate mocked method
func (m *CloudFormationAPIMock) GetTemplate(p0 *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.GetTemplateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.GetTemplateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetTemplateRequest mocked method
func (m *CloudFormationAPIMock) GetTemplateRequest(p0 *cloudformation.GetTemplateInput) (*request.Request, *cloudformation.GetTemplateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.GetTemplateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.GetTemplateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetTemplateSummary mocked method
func (m *CloudFormationAPIMock) GetTemplateSummary(p0 *cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.GetTemplateSummaryOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.GetTemplateSummaryOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetTemplateSummaryRequest mocked method
func (m *CloudFormationAPIMock) GetTemplateSummaryRequest(p0 *cloudformation.GetTemplateSummaryInput) (*request.Request, *cloudformation.GetTemplateSummaryOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.GetTemplateSummaryOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.GetTemplateSummaryOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListChangeSets mocked method
func (m *CloudFormationAPIMock) ListChangeSets(p0 *cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.ListChangeSetsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.ListChangeSetsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListChangeSetsRequest mocked method
func (m *CloudFormationAPIMock) ListChangeSetsRequest(p0 *cloudformation.ListChangeSetsInput) (*request.Request, *cloudformation.ListChangeSetsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.ListChangeSetsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.ListChangeSetsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListStackResources mocked method
func (m *CloudFormationAPIMock) ListStackResources(p0 *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.ListStackResourcesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.ListStackResourcesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListStackResourcesPages mocked method
func (m *CloudFormationAPIMock) ListStackResourcesPages(p0 *cloudformation.ListStackResourcesInput, p1 func(*cloudformation.ListStackResourcesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListStackResourcesRequest mocked method
func (m *CloudFormationAPIMock) ListStackResourcesRequest(p0 *cloudformation.ListStackResourcesInput) (*request.Request, *cloudformation.ListStackResourcesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.ListStackResourcesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.ListStackResourcesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListStacks mocked method
func (m *CloudFormationAPIMock) ListStacks(p0 *cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.ListStacksOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.ListStacksOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListStacksPages mocked method
func (m *CloudFormationAPIMock) ListStacksPages(p0 *cloudformation.ListStacksInput, p1 func(*cloudformation.ListStacksOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListStacksRequest mocked method
func (m *CloudFormationAPIMock) ListStacksRequest(p0 *cloudformation.ListStacksInput) (*request.Request, *cloudformation.ListStacksOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.ListStacksOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.ListStacksOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SetStackPolicy mocked method
func (m *CloudFormationAPIMock) SetStackPolicy(p0 *cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.SetStackPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.SetStackPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SetStackPolicyRequest mocked method
func (m *CloudFormationAPIMock) SetStackPolicyRequest(p0 *cloudformation.SetStackPolicyInput) (*request.Request, *cloudformation.SetStackPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.SetStackPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.SetStackPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SignalResource mocked method
func (m *CloudFormationAPIMock) SignalResource(p0 *cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.SignalResourceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.SignalResourceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SignalResourceRequest mocked method
func (m *CloudFormationAPIMock) SignalResourceRequest(p0 *cloudformation.SignalResourceInput) (*request.Request, *cloudformation.SignalResourceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.SignalResourceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.SignalResourceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateStack mocked method
func (m *CloudFormationAPIMock) UpdateStack(p0 *cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.UpdateStackOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.UpdateStackOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateStackRequest mocked method
func (m *CloudFormationAPIMock) UpdateStackRequest(p0 *cloudformation.UpdateStackInput) (*request.Request, *cloudformation.UpdateStackOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.UpdateStackOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.UpdateStackOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ValidateTemplate mocked method
func (m *CloudFormationAPIMock) ValidateTemplate(p0 *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error) {

	ret := m.Called(p0)

	var r0 *cloudformation.ValidateTemplateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *cloudformation.ValidateTemplateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ValidateTemplateRequest mocked method
func (m *CloudFormationAPIMock) ValidateTemplateRequest(p0 *cloudformation.ValidateTemplateInput) (*request.Request, *cloudformation.ValidateTemplateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *cloudformation.ValidateTemplateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *cloudformation.ValidateTemplateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}
