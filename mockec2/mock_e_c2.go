/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package mockec2

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"

	request "github.com/aws/aws-sdk-go/aws/request"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// EC2APIMock mock
type EC2APIMock struct {
	mock.Mock
	ec2iface.EC2API
}

func NewEC2APIMock() *EC2APIMock {
	return &EC2APIMock{}
}

// AcceptReservedInstancesExchangeQuote mocked method
func (m *EC2APIMock) AcceptReservedInstancesExchangeQuote(p0 *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AcceptReservedInstancesExchangeQuoteOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AcceptReservedInstancesExchangeQuoteOutput:
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

// AcceptReservedInstancesExchangeQuoteRequest mocked method
func (m *EC2APIMock) AcceptReservedInstancesExchangeQuoteRequest(p0 *ec2.AcceptReservedInstancesExchangeQuoteInput) (*request.Request, *ec2.AcceptReservedInstancesExchangeQuoteOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AcceptReservedInstancesExchangeQuoteOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AcceptReservedInstancesExchangeQuoteOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AcceptVpcPeeringConnection mocked method
func (m *EC2APIMock) AcceptVpcPeeringConnection(p0 *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AcceptVpcPeeringConnectionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AcceptVpcPeeringConnectionOutput:
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

// AcceptVpcPeeringConnectionRequest mocked method
func (m *EC2APIMock) AcceptVpcPeeringConnectionRequest(p0 *ec2.AcceptVpcPeeringConnectionInput) (*request.Request, *ec2.AcceptVpcPeeringConnectionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AcceptVpcPeeringConnectionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AcceptVpcPeeringConnectionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AllocateAddress mocked method
func (m *EC2APIMock) AllocateAddress(p0 *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AllocateAddressOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AllocateAddressOutput:
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

// AllocateAddressRequest mocked method
func (m *EC2APIMock) AllocateAddressRequest(p0 *ec2.AllocateAddressInput) (*request.Request, *ec2.AllocateAddressOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AllocateAddressOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AllocateAddressOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AllocateHosts mocked method
func (m *EC2APIMock) AllocateHosts(p0 *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AllocateHostsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AllocateHostsOutput:
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

// AllocateHostsRequest mocked method
func (m *EC2APIMock) AllocateHostsRequest(p0 *ec2.AllocateHostsInput) (*request.Request, *ec2.AllocateHostsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AllocateHostsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AllocateHostsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AssignPrivateIpAddresses mocked method
func (m *EC2APIMock) AssignPrivateIpAddresses(p0 *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AssignPrivateIpAddressesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AssignPrivateIpAddressesOutput:
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

// AssignPrivateIpAddressesRequest mocked method
func (m *EC2APIMock) AssignPrivateIpAddressesRequest(p0 *ec2.AssignPrivateIpAddressesInput) (*request.Request, *ec2.AssignPrivateIpAddressesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AssignPrivateIpAddressesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AssignPrivateIpAddressesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AssociateAddress mocked method
func (m *EC2APIMock) AssociateAddress(p0 *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AssociateAddressOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AssociateAddressOutput:
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

// AssociateAddressRequest mocked method
func (m *EC2APIMock) AssociateAddressRequest(p0 *ec2.AssociateAddressInput) (*request.Request, *ec2.AssociateAddressOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AssociateAddressOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AssociateAddressOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AssociateDhcpOptions mocked method
func (m *EC2APIMock) AssociateDhcpOptions(p0 *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AssociateDhcpOptionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AssociateDhcpOptionsOutput:
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

// AssociateDhcpOptionsRequest mocked method
func (m *EC2APIMock) AssociateDhcpOptionsRequest(p0 *ec2.AssociateDhcpOptionsInput) (*request.Request, *ec2.AssociateDhcpOptionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AssociateDhcpOptionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AssociateDhcpOptionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AssociateRouteTable mocked method
func (m *EC2APIMock) AssociateRouteTable(p0 *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AssociateRouteTableOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AssociateRouteTableOutput:
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

// AssociateRouteTableRequest mocked method
func (m *EC2APIMock) AssociateRouteTableRequest(p0 *ec2.AssociateRouteTableInput) (*request.Request, *ec2.AssociateRouteTableOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AssociateRouteTableOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AssociateRouteTableOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachClassicLinkVpc mocked method
func (m *EC2APIMock) AttachClassicLinkVpc(p0 *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AttachClassicLinkVpcOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AttachClassicLinkVpcOutput:
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

// AttachClassicLinkVpcRequest mocked method
func (m *EC2APIMock) AttachClassicLinkVpcRequest(p0 *ec2.AttachClassicLinkVpcInput) (*request.Request, *ec2.AttachClassicLinkVpcOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AttachClassicLinkVpcOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AttachClassicLinkVpcOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachInternetGateway mocked method
func (m *EC2APIMock) AttachInternetGateway(p0 *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AttachInternetGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AttachInternetGatewayOutput:
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

// AttachInternetGatewayRequest mocked method
func (m *EC2APIMock) AttachInternetGatewayRequest(p0 *ec2.AttachInternetGatewayInput) (*request.Request, *ec2.AttachInternetGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AttachInternetGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AttachInternetGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachNetworkInterface mocked method
func (m *EC2APIMock) AttachNetworkInterface(p0 *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AttachNetworkInterfaceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AttachNetworkInterfaceOutput:
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

// AttachNetworkInterfaceRequest mocked method
func (m *EC2APIMock) AttachNetworkInterfaceRequest(p0 *ec2.AttachNetworkInterfaceInput) (*request.Request, *ec2.AttachNetworkInterfaceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AttachNetworkInterfaceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AttachNetworkInterfaceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachVolume mocked method
func (m *EC2APIMock) AttachVolume(p0 *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {

	ret := m.Called(p0)

	var r0 *ec2.VolumeAttachment
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.VolumeAttachment:
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

// AttachVolumeRequest mocked method
func (m *EC2APIMock) AttachVolumeRequest(p0 *ec2.AttachVolumeInput) (*request.Request, *ec2.VolumeAttachment) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.VolumeAttachment
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.VolumeAttachment:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachVpnGateway mocked method
func (m *EC2APIMock) AttachVpnGateway(p0 *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AttachVpnGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AttachVpnGatewayOutput:
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

// AttachVpnGatewayRequest mocked method
func (m *EC2APIMock) AttachVpnGatewayRequest(p0 *ec2.AttachVpnGatewayInput) (*request.Request, *ec2.AttachVpnGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AttachVpnGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AttachVpnGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AuthorizeSecurityGroupEgress mocked method
func (m *EC2APIMock) AuthorizeSecurityGroupEgress(p0 *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AuthorizeSecurityGroupEgressOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AuthorizeSecurityGroupEgressOutput:
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

// AuthorizeSecurityGroupEgressRequest mocked method
func (m *EC2APIMock) AuthorizeSecurityGroupEgressRequest(p0 *ec2.AuthorizeSecurityGroupEgressInput) (*request.Request, *ec2.AuthorizeSecurityGroupEgressOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AuthorizeSecurityGroupEgressOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AuthorizeSecurityGroupEgressOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AuthorizeSecurityGroupIngress mocked method
func (m *EC2APIMock) AuthorizeSecurityGroupIngress(p0 *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.AuthorizeSecurityGroupIngressOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.AuthorizeSecurityGroupIngressOutput:
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

// AuthorizeSecurityGroupIngressRequest mocked method
func (m *EC2APIMock) AuthorizeSecurityGroupIngressRequest(p0 *ec2.AuthorizeSecurityGroupIngressInput) (*request.Request, *ec2.AuthorizeSecurityGroupIngressOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.AuthorizeSecurityGroupIngressOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.AuthorizeSecurityGroupIngressOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// BundleInstance mocked method
func (m *EC2APIMock) BundleInstance(p0 *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.BundleInstanceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.BundleInstanceOutput:
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

// BundleInstanceRequest mocked method
func (m *EC2APIMock) BundleInstanceRequest(p0 *ec2.BundleInstanceInput) (*request.Request, *ec2.BundleInstanceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.BundleInstanceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.BundleInstanceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CancelBundleTask mocked method
func (m *EC2APIMock) CancelBundleTask(p0 *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CancelBundleTaskOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CancelBundleTaskOutput:
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

// CancelBundleTaskRequest mocked method
func (m *EC2APIMock) CancelBundleTaskRequest(p0 *ec2.CancelBundleTaskInput) (*request.Request, *ec2.CancelBundleTaskOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CancelBundleTaskOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CancelBundleTaskOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CancelConversionTask mocked method
func (m *EC2APIMock) CancelConversionTask(p0 *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CancelConversionTaskOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CancelConversionTaskOutput:
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

// CancelConversionTaskRequest mocked method
func (m *EC2APIMock) CancelConversionTaskRequest(p0 *ec2.CancelConversionTaskInput) (*request.Request, *ec2.CancelConversionTaskOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CancelConversionTaskOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CancelConversionTaskOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CancelExportTask mocked method
func (m *EC2APIMock) CancelExportTask(p0 *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CancelExportTaskOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CancelExportTaskOutput:
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

// CancelExportTaskRequest mocked method
func (m *EC2APIMock) CancelExportTaskRequest(p0 *ec2.CancelExportTaskInput) (*request.Request, *ec2.CancelExportTaskOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CancelExportTaskOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CancelExportTaskOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CancelImportTask mocked method
func (m *EC2APIMock) CancelImportTask(p0 *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CancelImportTaskOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CancelImportTaskOutput:
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

// CancelImportTaskRequest mocked method
func (m *EC2APIMock) CancelImportTaskRequest(p0 *ec2.CancelImportTaskInput) (*request.Request, *ec2.CancelImportTaskOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CancelImportTaskOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CancelImportTaskOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CancelReservedInstancesListing mocked method
func (m *EC2APIMock) CancelReservedInstancesListing(p0 *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CancelReservedInstancesListingOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CancelReservedInstancesListingOutput:
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

// CancelReservedInstancesListingRequest mocked method
func (m *EC2APIMock) CancelReservedInstancesListingRequest(p0 *ec2.CancelReservedInstancesListingInput) (*request.Request, *ec2.CancelReservedInstancesListingOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CancelReservedInstancesListingOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CancelReservedInstancesListingOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CancelSpotFleetRequests mocked method
func (m *EC2APIMock) CancelSpotFleetRequests(p0 *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CancelSpotFleetRequestsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CancelSpotFleetRequestsOutput:
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

// CancelSpotFleetRequestsRequest mocked method
func (m *EC2APIMock) CancelSpotFleetRequestsRequest(p0 *ec2.CancelSpotFleetRequestsInput) (*request.Request, *ec2.CancelSpotFleetRequestsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CancelSpotFleetRequestsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CancelSpotFleetRequestsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CancelSpotInstanceRequests mocked method
func (m *EC2APIMock) CancelSpotInstanceRequests(p0 *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CancelSpotInstanceRequestsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CancelSpotInstanceRequestsOutput:
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

// CancelSpotInstanceRequestsRequest mocked method
func (m *EC2APIMock) CancelSpotInstanceRequestsRequest(p0 *ec2.CancelSpotInstanceRequestsInput) (*request.Request, *ec2.CancelSpotInstanceRequestsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CancelSpotInstanceRequestsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CancelSpotInstanceRequestsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ConfirmProductInstance mocked method
func (m *EC2APIMock) ConfirmProductInstance(p0 *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ConfirmProductInstanceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ConfirmProductInstanceOutput:
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

// ConfirmProductInstanceRequest mocked method
func (m *EC2APIMock) ConfirmProductInstanceRequest(p0 *ec2.ConfirmProductInstanceInput) (*request.Request, *ec2.ConfirmProductInstanceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ConfirmProductInstanceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ConfirmProductInstanceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CopyImage mocked method
func (m *EC2APIMock) CopyImage(p0 *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CopyImageOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CopyImageOutput:
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

// CopyImageRequest mocked method
func (m *EC2APIMock) CopyImageRequest(p0 *ec2.CopyImageInput) (*request.Request, *ec2.CopyImageOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CopyImageOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CopyImageOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CopySnapshot mocked method
func (m *EC2APIMock) CopySnapshot(p0 *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CopySnapshotOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CopySnapshotOutput:
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

// CopySnapshotRequest mocked method
func (m *EC2APIMock) CopySnapshotRequest(p0 *ec2.CopySnapshotInput) (*request.Request, *ec2.CopySnapshotOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CopySnapshotOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CopySnapshotOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateCustomerGateway mocked method
func (m *EC2APIMock) CreateCustomerGateway(p0 *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateCustomerGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateCustomerGatewayOutput:
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

// CreateCustomerGatewayRequest mocked method
func (m *EC2APIMock) CreateCustomerGatewayRequest(p0 *ec2.CreateCustomerGatewayInput) (*request.Request, *ec2.CreateCustomerGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateCustomerGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateCustomerGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateDhcpOptions mocked method
func (m *EC2APIMock) CreateDhcpOptions(p0 *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateDhcpOptionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateDhcpOptionsOutput:
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

// CreateDhcpOptionsRequest mocked method
func (m *EC2APIMock) CreateDhcpOptionsRequest(p0 *ec2.CreateDhcpOptionsInput) (*request.Request, *ec2.CreateDhcpOptionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateDhcpOptionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateDhcpOptionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateFlowLogs mocked method
func (m *EC2APIMock) CreateFlowLogs(p0 *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateFlowLogsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateFlowLogsOutput:
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

// CreateFlowLogsRequest mocked method
func (m *EC2APIMock) CreateFlowLogsRequest(p0 *ec2.CreateFlowLogsInput) (*request.Request, *ec2.CreateFlowLogsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateFlowLogsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateFlowLogsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateImage mocked method
func (m *EC2APIMock) CreateImage(p0 *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateImageOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateImageOutput:
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

// CreateImageRequest mocked method
func (m *EC2APIMock) CreateImageRequest(p0 *ec2.CreateImageInput) (*request.Request, *ec2.CreateImageOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateImageOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateImageOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateInstanceExportTask mocked method
func (m *EC2APIMock) CreateInstanceExportTask(p0 *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateInstanceExportTaskOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateInstanceExportTaskOutput:
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

// CreateInstanceExportTaskRequest mocked method
func (m *EC2APIMock) CreateInstanceExportTaskRequest(p0 *ec2.CreateInstanceExportTaskInput) (*request.Request, *ec2.CreateInstanceExportTaskOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateInstanceExportTaskOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateInstanceExportTaskOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateInternetGateway mocked method
func (m *EC2APIMock) CreateInternetGateway(p0 *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateInternetGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateInternetGatewayOutput:
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

// CreateInternetGatewayRequest mocked method
func (m *EC2APIMock) CreateInternetGatewayRequest(p0 *ec2.CreateInternetGatewayInput) (*request.Request, *ec2.CreateInternetGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateInternetGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateInternetGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateKeyPair mocked method
func (m *EC2APIMock) CreateKeyPair(p0 *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateKeyPairOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateKeyPairOutput:
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

// CreateKeyPairRequest mocked method
func (m *EC2APIMock) CreateKeyPairRequest(p0 *ec2.CreateKeyPairInput) (*request.Request, *ec2.CreateKeyPairOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateKeyPairOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateKeyPairOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateNatGateway mocked method
func (m *EC2APIMock) CreateNatGateway(p0 *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateNatGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateNatGatewayOutput:
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

// CreateNatGatewayRequest mocked method
func (m *EC2APIMock) CreateNatGatewayRequest(p0 *ec2.CreateNatGatewayInput) (*request.Request, *ec2.CreateNatGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateNatGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateNatGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateNetworkAcl mocked method
func (m *EC2APIMock) CreateNetworkAcl(p0 *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateNetworkAclOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateNetworkAclOutput:
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

// CreateNetworkAclEntry mocked method
func (m *EC2APIMock) CreateNetworkAclEntry(p0 *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateNetworkAclEntryOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateNetworkAclEntryOutput:
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

// CreateNetworkAclEntryRequest mocked method
func (m *EC2APIMock) CreateNetworkAclEntryRequest(p0 *ec2.CreateNetworkAclEntryInput) (*request.Request, *ec2.CreateNetworkAclEntryOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateNetworkAclEntryOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateNetworkAclEntryOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateNetworkAclRequest mocked method
func (m *EC2APIMock) CreateNetworkAclRequest(p0 *ec2.CreateNetworkAclInput) (*request.Request, *ec2.CreateNetworkAclOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateNetworkAclOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateNetworkAclOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateNetworkInterface mocked method
func (m *EC2APIMock) CreateNetworkInterface(p0 *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateNetworkInterfaceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateNetworkInterfaceOutput:
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

// CreateNetworkInterfaceRequest mocked method
func (m *EC2APIMock) CreateNetworkInterfaceRequest(p0 *ec2.CreateNetworkInterfaceInput) (*request.Request, *ec2.CreateNetworkInterfaceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateNetworkInterfaceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateNetworkInterfaceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreatePlacementGroup mocked method
func (m *EC2APIMock) CreatePlacementGroup(p0 *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreatePlacementGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreatePlacementGroupOutput:
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

// CreatePlacementGroupRequest mocked method
func (m *EC2APIMock) CreatePlacementGroupRequest(p0 *ec2.CreatePlacementGroupInput) (*request.Request, *ec2.CreatePlacementGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreatePlacementGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreatePlacementGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateReservedInstancesListing mocked method
func (m *EC2APIMock) CreateReservedInstancesListing(p0 *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateReservedInstancesListingOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateReservedInstancesListingOutput:
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

// CreateReservedInstancesListingRequest mocked method
func (m *EC2APIMock) CreateReservedInstancesListingRequest(p0 *ec2.CreateReservedInstancesListingInput) (*request.Request, *ec2.CreateReservedInstancesListingOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateReservedInstancesListingOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateReservedInstancesListingOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateRoute mocked method
func (m *EC2APIMock) CreateRoute(p0 *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateRouteOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateRouteOutput:
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

// CreateRouteRequest mocked method
func (m *EC2APIMock) CreateRouteRequest(p0 *ec2.CreateRouteInput) (*request.Request, *ec2.CreateRouteOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateRouteOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateRouteOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateRouteTable mocked method
func (m *EC2APIMock) CreateRouteTable(p0 *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateRouteTableOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateRouteTableOutput:
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

// CreateRouteTableRequest mocked method
func (m *EC2APIMock) CreateRouteTableRequest(p0 *ec2.CreateRouteTableInput) (*request.Request, *ec2.CreateRouteTableOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateRouteTableOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateRouteTableOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateSecurityGroup mocked method
func (m *EC2APIMock) CreateSecurityGroup(p0 *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateSecurityGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateSecurityGroupOutput:
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

// CreateSecurityGroupRequest mocked method
func (m *EC2APIMock) CreateSecurityGroupRequest(p0 *ec2.CreateSecurityGroupInput) (*request.Request, *ec2.CreateSecurityGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateSecurityGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateSecurityGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateSnapshot mocked method
func (m *EC2APIMock) CreateSnapshot(p0 *ec2.CreateSnapshotInput) (*ec2.Snapshot, error) {

	ret := m.Called(p0)

	var r0 *ec2.Snapshot
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.Snapshot:
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

// CreateSnapshotRequest mocked method
func (m *EC2APIMock) CreateSnapshotRequest(p0 *ec2.CreateSnapshotInput) (*request.Request, *ec2.Snapshot) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.Snapshot
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.Snapshot:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateSpotDatafeedSubscription mocked method
func (m *EC2APIMock) CreateSpotDatafeedSubscription(p0 *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateSpotDatafeedSubscriptionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateSpotDatafeedSubscriptionOutput:
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

// CreateSpotDatafeedSubscriptionRequest mocked method
func (m *EC2APIMock) CreateSpotDatafeedSubscriptionRequest(p0 *ec2.CreateSpotDatafeedSubscriptionInput) (*request.Request, *ec2.CreateSpotDatafeedSubscriptionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateSpotDatafeedSubscriptionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateSpotDatafeedSubscriptionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateSubnet mocked method
func (m *EC2APIMock) CreateSubnet(p0 *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateSubnetOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateSubnetOutput:
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

// CreateSubnetRequest mocked method
func (m *EC2APIMock) CreateSubnetRequest(p0 *ec2.CreateSubnetInput) (*request.Request, *ec2.CreateSubnetOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateSubnetOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateSubnetOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateTags mocked method
func (m *EC2APIMock) CreateTags(p0 *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateTagsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateTagsOutput:
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

// CreateTagsRequest mocked method
func (m *EC2APIMock) CreateTagsRequest(p0 *ec2.CreateTagsInput) (*request.Request, *ec2.CreateTagsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateTagsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateTagsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVolume mocked method
func (m *EC2APIMock) CreateVolume(p0 *ec2.CreateVolumeInput) (*ec2.Volume, error) {

	ret := m.Called(p0)

	var r0 *ec2.Volume
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.Volume:
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

// CreateVolumeRequest mocked method
func (m *EC2APIMock) CreateVolumeRequest(p0 *ec2.CreateVolumeInput) (*request.Request, *ec2.Volume) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.Volume
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.Volume:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVpc mocked method
func (m *EC2APIMock) CreateVpc(p0 *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateVpcOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateVpcOutput:
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

// CreateVpcEndpoint mocked method
func (m *EC2APIMock) CreateVpcEndpoint(p0 *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateVpcEndpointOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateVpcEndpointOutput:
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

// CreateVpcEndpointRequest mocked method
func (m *EC2APIMock) CreateVpcEndpointRequest(p0 *ec2.CreateVpcEndpointInput) (*request.Request, *ec2.CreateVpcEndpointOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateVpcEndpointOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateVpcEndpointOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVpcPeeringConnection mocked method
func (m *EC2APIMock) CreateVpcPeeringConnection(p0 *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateVpcPeeringConnectionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateVpcPeeringConnectionOutput:
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

// CreateVpcPeeringConnectionRequest mocked method
func (m *EC2APIMock) CreateVpcPeeringConnectionRequest(p0 *ec2.CreateVpcPeeringConnectionInput) (*request.Request, *ec2.CreateVpcPeeringConnectionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateVpcPeeringConnectionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateVpcPeeringConnectionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVpcRequest mocked method
func (m *EC2APIMock) CreateVpcRequest(p0 *ec2.CreateVpcInput) (*request.Request, *ec2.CreateVpcOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateVpcOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateVpcOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVpnConnection mocked method
func (m *EC2APIMock) CreateVpnConnection(p0 *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateVpnConnectionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateVpnConnectionOutput:
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

// CreateVpnConnectionRequest mocked method
func (m *EC2APIMock) CreateVpnConnectionRequest(p0 *ec2.CreateVpnConnectionInput) (*request.Request, *ec2.CreateVpnConnectionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateVpnConnectionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateVpnConnectionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVpnConnectionRoute mocked method
func (m *EC2APIMock) CreateVpnConnectionRoute(p0 *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateVpnConnectionRouteOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateVpnConnectionRouteOutput:
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

// CreateVpnConnectionRouteRequest mocked method
func (m *EC2APIMock) CreateVpnConnectionRouteRequest(p0 *ec2.CreateVpnConnectionRouteInput) (*request.Request, *ec2.CreateVpnConnectionRouteOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateVpnConnectionRouteOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateVpnConnectionRouteOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVpnGateway mocked method
func (m *EC2APIMock) CreateVpnGateway(p0 *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.CreateVpnGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.CreateVpnGatewayOutput:
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

// CreateVpnGatewayRequest mocked method
func (m *EC2APIMock) CreateVpnGatewayRequest(p0 *ec2.CreateVpnGatewayInput) (*request.Request, *ec2.CreateVpnGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.CreateVpnGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.CreateVpnGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteCustomerGateway mocked method
func (m *EC2APIMock) DeleteCustomerGateway(p0 *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteCustomerGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteCustomerGatewayOutput:
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

// DeleteCustomerGatewayRequest mocked method
func (m *EC2APIMock) DeleteCustomerGatewayRequest(p0 *ec2.DeleteCustomerGatewayInput) (*request.Request, *ec2.DeleteCustomerGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteCustomerGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteCustomerGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteDhcpOptions mocked method
func (m *EC2APIMock) DeleteDhcpOptions(p0 *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteDhcpOptionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteDhcpOptionsOutput:
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

// DeleteDhcpOptionsRequest mocked method
func (m *EC2APIMock) DeleteDhcpOptionsRequest(p0 *ec2.DeleteDhcpOptionsInput) (*request.Request, *ec2.DeleteDhcpOptionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteDhcpOptionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteDhcpOptionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteFlowLogs mocked method
func (m *EC2APIMock) DeleteFlowLogs(p0 *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteFlowLogsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteFlowLogsOutput:
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

// DeleteFlowLogsRequest mocked method
func (m *EC2APIMock) DeleteFlowLogsRequest(p0 *ec2.DeleteFlowLogsInput) (*request.Request, *ec2.DeleteFlowLogsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteFlowLogsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteFlowLogsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteInternetGateway mocked method
func (m *EC2APIMock) DeleteInternetGateway(p0 *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteInternetGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteInternetGatewayOutput:
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

// DeleteInternetGatewayRequest mocked method
func (m *EC2APIMock) DeleteInternetGatewayRequest(p0 *ec2.DeleteInternetGatewayInput) (*request.Request, *ec2.DeleteInternetGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteInternetGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteInternetGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteKeyPair mocked method
func (m *EC2APIMock) DeleteKeyPair(p0 *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteKeyPairOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteKeyPairOutput:
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

// DeleteKeyPairRequest mocked method
func (m *EC2APIMock) DeleteKeyPairRequest(p0 *ec2.DeleteKeyPairInput) (*request.Request, *ec2.DeleteKeyPairOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteKeyPairOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteKeyPairOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteNatGateway mocked method
func (m *EC2APIMock) DeleteNatGateway(p0 *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteNatGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteNatGatewayOutput:
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

// DeleteNatGatewayRequest mocked method
func (m *EC2APIMock) DeleteNatGatewayRequest(p0 *ec2.DeleteNatGatewayInput) (*request.Request, *ec2.DeleteNatGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteNatGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteNatGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteNetworkAcl mocked method
func (m *EC2APIMock) DeleteNetworkAcl(p0 *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteNetworkAclOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteNetworkAclOutput:
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

// DeleteNetworkAclEntry mocked method
func (m *EC2APIMock) DeleteNetworkAclEntry(p0 *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteNetworkAclEntryOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteNetworkAclEntryOutput:
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

// DeleteNetworkAclEntryRequest mocked method
func (m *EC2APIMock) DeleteNetworkAclEntryRequest(p0 *ec2.DeleteNetworkAclEntryInput) (*request.Request, *ec2.DeleteNetworkAclEntryOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteNetworkAclEntryOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteNetworkAclEntryOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteNetworkAclRequest mocked method
func (m *EC2APIMock) DeleteNetworkAclRequest(p0 *ec2.DeleteNetworkAclInput) (*request.Request, *ec2.DeleteNetworkAclOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteNetworkAclOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteNetworkAclOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteNetworkInterface mocked method
func (m *EC2APIMock) DeleteNetworkInterface(p0 *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteNetworkInterfaceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteNetworkInterfaceOutput:
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

// DeleteNetworkInterfaceRequest mocked method
func (m *EC2APIMock) DeleteNetworkInterfaceRequest(p0 *ec2.DeleteNetworkInterfaceInput) (*request.Request, *ec2.DeleteNetworkInterfaceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteNetworkInterfaceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteNetworkInterfaceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeletePlacementGroup mocked method
func (m *EC2APIMock) DeletePlacementGroup(p0 *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeletePlacementGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeletePlacementGroupOutput:
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

// DeletePlacementGroupRequest mocked method
func (m *EC2APIMock) DeletePlacementGroupRequest(p0 *ec2.DeletePlacementGroupInput) (*request.Request, *ec2.DeletePlacementGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeletePlacementGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeletePlacementGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteRoute mocked method
func (m *EC2APIMock) DeleteRoute(p0 *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteRouteOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteRouteOutput:
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

// DeleteRouteRequest mocked method
func (m *EC2APIMock) DeleteRouteRequest(p0 *ec2.DeleteRouteInput) (*request.Request, *ec2.DeleteRouteOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteRouteOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteRouteOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteRouteTable mocked method
func (m *EC2APIMock) DeleteRouteTable(p0 *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteRouteTableOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteRouteTableOutput:
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

// DeleteRouteTableRequest mocked method
func (m *EC2APIMock) DeleteRouteTableRequest(p0 *ec2.DeleteRouteTableInput) (*request.Request, *ec2.DeleteRouteTableOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteRouteTableOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteRouteTableOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSecurityGroup mocked method
func (m *EC2APIMock) DeleteSecurityGroup(p0 *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteSecurityGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteSecurityGroupOutput:
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

// DeleteSecurityGroupRequest mocked method
func (m *EC2APIMock) DeleteSecurityGroupRequest(p0 *ec2.DeleteSecurityGroupInput) (*request.Request, *ec2.DeleteSecurityGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteSecurityGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteSecurityGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSnapshot mocked method
func (m *EC2APIMock) DeleteSnapshot(p0 *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteSnapshotOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteSnapshotOutput:
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

// DeleteSnapshotRequest mocked method
func (m *EC2APIMock) DeleteSnapshotRequest(p0 *ec2.DeleteSnapshotInput) (*request.Request, *ec2.DeleteSnapshotOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteSnapshotOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteSnapshotOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSpotDatafeedSubscription mocked method
func (m *EC2APIMock) DeleteSpotDatafeedSubscription(p0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteSpotDatafeedSubscriptionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteSpotDatafeedSubscriptionOutput:
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

// DeleteSpotDatafeedSubscriptionRequest mocked method
func (m *EC2APIMock) DeleteSpotDatafeedSubscriptionRequest(p0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DeleteSpotDatafeedSubscriptionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteSpotDatafeedSubscriptionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteSpotDatafeedSubscriptionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSubnet mocked method
func (m *EC2APIMock) DeleteSubnet(p0 *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteSubnetOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteSubnetOutput:
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

// DeleteSubnetRequest mocked method
func (m *EC2APIMock) DeleteSubnetRequest(p0 *ec2.DeleteSubnetInput) (*request.Request, *ec2.DeleteSubnetOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteSubnetOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteSubnetOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteTags mocked method
func (m *EC2APIMock) DeleteTags(p0 *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteTagsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteTagsOutput:
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

// DeleteTagsRequest mocked method
func (m *EC2APIMock) DeleteTagsRequest(p0 *ec2.DeleteTagsInput) (*request.Request, *ec2.DeleteTagsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteTagsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteTagsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVolume mocked method
func (m *EC2APIMock) DeleteVolume(p0 *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteVolumeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteVolumeOutput:
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

// DeleteVolumeRequest mocked method
func (m *EC2APIMock) DeleteVolumeRequest(p0 *ec2.DeleteVolumeInput) (*request.Request, *ec2.DeleteVolumeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteVolumeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteVolumeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVpc mocked method
func (m *EC2APIMock) DeleteVpc(p0 *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteVpcOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteVpcOutput:
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

// DeleteVpcEndpoints mocked method
func (m *EC2APIMock) DeleteVpcEndpoints(p0 *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteVpcEndpointsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteVpcEndpointsOutput:
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

// DeleteVpcEndpointsRequest mocked method
func (m *EC2APIMock) DeleteVpcEndpointsRequest(p0 *ec2.DeleteVpcEndpointsInput) (*request.Request, *ec2.DeleteVpcEndpointsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteVpcEndpointsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteVpcEndpointsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVpcPeeringConnection mocked method
func (m *EC2APIMock) DeleteVpcPeeringConnection(p0 *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteVpcPeeringConnectionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteVpcPeeringConnectionOutput:
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

// DeleteVpcPeeringConnectionRequest mocked method
func (m *EC2APIMock) DeleteVpcPeeringConnectionRequest(p0 *ec2.DeleteVpcPeeringConnectionInput) (*request.Request, *ec2.DeleteVpcPeeringConnectionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteVpcPeeringConnectionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteVpcPeeringConnectionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVpcRequest mocked method
func (m *EC2APIMock) DeleteVpcRequest(p0 *ec2.DeleteVpcInput) (*request.Request, *ec2.DeleteVpcOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteVpcOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteVpcOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVpnConnection mocked method
func (m *EC2APIMock) DeleteVpnConnection(p0 *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteVpnConnectionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteVpnConnectionOutput:
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

// DeleteVpnConnectionRequest mocked method
func (m *EC2APIMock) DeleteVpnConnectionRequest(p0 *ec2.DeleteVpnConnectionInput) (*request.Request, *ec2.DeleteVpnConnectionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteVpnConnectionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteVpnConnectionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVpnConnectionRoute mocked method
func (m *EC2APIMock) DeleteVpnConnectionRoute(p0 *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteVpnConnectionRouteOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteVpnConnectionRouteOutput:
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

// DeleteVpnConnectionRouteRequest mocked method
func (m *EC2APIMock) DeleteVpnConnectionRouteRequest(p0 *ec2.DeleteVpnConnectionRouteInput) (*request.Request, *ec2.DeleteVpnConnectionRouteOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteVpnConnectionRouteOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteVpnConnectionRouteOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVpnGateway mocked method
func (m *EC2APIMock) DeleteVpnGateway(p0 *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeleteVpnGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeleteVpnGatewayOutput:
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

// DeleteVpnGatewayRequest mocked method
func (m *EC2APIMock) DeleteVpnGatewayRequest(p0 *ec2.DeleteVpnGatewayInput) (*request.Request, *ec2.DeleteVpnGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeleteVpnGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeleteVpnGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeregisterImage mocked method
func (m *EC2APIMock) DeregisterImage(p0 *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DeregisterImageOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DeregisterImageOutput:
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

// DeregisterImageRequest mocked method
func (m *EC2APIMock) DeregisterImageRequest(p0 *ec2.DeregisterImageInput) (*request.Request, *ec2.DeregisterImageOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DeregisterImageOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DeregisterImageOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeAccountAttributes mocked method
func (m *EC2APIMock) DescribeAccountAttributes(p0 *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeAccountAttributesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeAccountAttributesOutput:
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

// DescribeAccountAttributesRequest mocked method
func (m *EC2APIMock) DescribeAccountAttributesRequest(p0 *ec2.DescribeAccountAttributesInput) (*request.Request, *ec2.DescribeAccountAttributesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeAccountAttributesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeAccountAttributesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeAddresses mocked method
func (m *EC2APIMock) DescribeAddresses(p0 *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeAddressesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeAddressesOutput:
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

// DescribeAddressesRequest mocked method
func (m *EC2APIMock) DescribeAddressesRequest(p0 *ec2.DescribeAddressesInput) (*request.Request, *ec2.DescribeAddressesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeAddressesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeAddressesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeAvailabilityZones mocked method
func (m *EC2APIMock) DescribeAvailabilityZones(p0 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeAvailabilityZonesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeAvailabilityZonesOutput:
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

// DescribeAvailabilityZonesRequest mocked method
func (m *EC2APIMock) DescribeAvailabilityZonesRequest(p0 *ec2.DescribeAvailabilityZonesInput) (*request.Request, *ec2.DescribeAvailabilityZonesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeAvailabilityZonesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeAvailabilityZonesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeBundleTasks mocked method
func (m *EC2APIMock) DescribeBundleTasks(p0 *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeBundleTasksOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeBundleTasksOutput:
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

// DescribeBundleTasksRequest mocked method
func (m *EC2APIMock) DescribeBundleTasksRequest(p0 *ec2.DescribeBundleTasksInput) (*request.Request, *ec2.DescribeBundleTasksOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeBundleTasksOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeBundleTasksOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeClassicLinkInstances mocked method
func (m *EC2APIMock) DescribeClassicLinkInstances(p0 *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeClassicLinkInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeClassicLinkInstancesOutput:
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

// DescribeClassicLinkInstancesRequest mocked method
func (m *EC2APIMock) DescribeClassicLinkInstancesRequest(p0 *ec2.DescribeClassicLinkInstancesInput) (*request.Request, *ec2.DescribeClassicLinkInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeClassicLinkInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeClassicLinkInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeConversionTasks mocked method
func (m *EC2APIMock) DescribeConversionTasks(p0 *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeConversionTasksOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeConversionTasksOutput:
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

// DescribeConversionTasksRequest mocked method
func (m *EC2APIMock) DescribeConversionTasksRequest(p0 *ec2.DescribeConversionTasksInput) (*request.Request, *ec2.DescribeConversionTasksOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeConversionTasksOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeConversionTasksOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeCustomerGateways mocked method
func (m *EC2APIMock) DescribeCustomerGateways(p0 *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeCustomerGatewaysOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeCustomerGatewaysOutput:
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

// DescribeCustomerGatewaysRequest mocked method
func (m *EC2APIMock) DescribeCustomerGatewaysRequest(p0 *ec2.DescribeCustomerGatewaysInput) (*request.Request, *ec2.DescribeCustomerGatewaysOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeCustomerGatewaysOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeCustomerGatewaysOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeDhcpOptions mocked method
func (m *EC2APIMock) DescribeDhcpOptions(p0 *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeDhcpOptionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeDhcpOptionsOutput:
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

// DescribeDhcpOptionsRequest mocked method
func (m *EC2APIMock) DescribeDhcpOptionsRequest(p0 *ec2.DescribeDhcpOptionsInput) (*request.Request, *ec2.DescribeDhcpOptionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeDhcpOptionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeDhcpOptionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeExportTasks mocked method
func (m *EC2APIMock) DescribeExportTasks(p0 *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeExportTasksOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeExportTasksOutput:
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

// DescribeExportTasksRequest mocked method
func (m *EC2APIMock) DescribeExportTasksRequest(p0 *ec2.DescribeExportTasksInput) (*request.Request, *ec2.DescribeExportTasksOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeExportTasksOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeExportTasksOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeFlowLogs mocked method
func (m *EC2APIMock) DescribeFlowLogs(p0 *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeFlowLogsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeFlowLogsOutput:
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

// DescribeFlowLogsRequest mocked method
func (m *EC2APIMock) DescribeFlowLogsRequest(p0 *ec2.DescribeFlowLogsInput) (*request.Request, *ec2.DescribeFlowLogsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeFlowLogsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeFlowLogsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeHostReservationOfferings mocked method
func (m *EC2APIMock) DescribeHostReservationOfferings(p0 *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeHostReservationOfferingsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeHostReservationOfferingsOutput:
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

// DescribeHostReservationOfferingsRequest mocked method
func (m *EC2APIMock) DescribeHostReservationOfferingsRequest(p0 *ec2.DescribeHostReservationOfferingsInput) (*request.Request, *ec2.DescribeHostReservationOfferingsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeHostReservationOfferingsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeHostReservationOfferingsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeHostReservations mocked method
func (m *EC2APIMock) DescribeHostReservations(p0 *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeHostReservationsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeHostReservationsOutput:
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

// DescribeHostReservationsRequest mocked method
func (m *EC2APIMock) DescribeHostReservationsRequest(p0 *ec2.DescribeHostReservationsInput) (*request.Request, *ec2.DescribeHostReservationsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeHostReservationsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeHostReservationsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeHosts mocked method
func (m *EC2APIMock) DescribeHosts(p0 *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeHostsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeHostsOutput:
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

// DescribeHostsRequest mocked method
func (m *EC2APIMock) DescribeHostsRequest(p0 *ec2.DescribeHostsInput) (*request.Request, *ec2.DescribeHostsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeHostsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeHostsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeIdFormat mocked method
func (m *EC2APIMock) DescribeIdFormat(p0 *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeIdFormatOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeIdFormatOutput:
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

// DescribeIdFormatRequest mocked method
func (m *EC2APIMock) DescribeIdFormatRequest(p0 *ec2.DescribeIdFormatInput) (*request.Request, *ec2.DescribeIdFormatOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeIdFormatOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeIdFormatOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeIdentityIdFormat mocked method
func (m *EC2APIMock) DescribeIdentityIdFormat(p0 *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeIdentityIdFormatOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeIdentityIdFormatOutput:
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

// DescribeIdentityIdFormatRequest mocked method
func (m *EC2APIMock) DescribeIdentityIdFormatRequest(p0 *ec2.DescribeIdentityIdFormatInput) (*request.Request, *ec2.DescribeIdentityIdFormatOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeIdentityIdFormatOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeIdentityIdFormatOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeImageAttribute mocked method
func (m *EC2APIMock) DescribeImageAttribute(p0 *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeImageAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeImageAttributeOutput:
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

// DescribeImageAttributeRequest mocked method
func (m *EC2APIMock) DescribeImageAttributeRequest(p0 *ec2.DescribeImageAttributeInput) (*request.Request, *ec2.DescribeImageAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeImageAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeImageAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeImages mocked method
func (m *EC2APIMock) DescribeImages(p0 *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeImagesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeImagesOutput:
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

// DescribeImagesRequest mocked method
func (m *EC2APIMock) DescribeImagesRequest(p0 *ec2.DescribeImagesInput) (*request.Request, *ec2.DescribeImagesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeImagesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeImagesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeImportImageTasks mocked method
func (m *EC2APIMock) DescribeImportImageTasks(p0 *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeImportImageTasksOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeImportImageTasksOutput:
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

// DescribeImportImageTasksRequest mocked method
func (m *EC2APIMock) DescribeImportImageTasksRequest(p0 *ec2.DescribeImportImageTasksInput) (*request.Request, *ec2.DescribeImportImageTasksOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeImportImageTasksOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeImportImageTasksOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeImportSnapshotTasks mocked method
func (m *EC2APIMock) DescribeImportSnapshotTasks(p0 *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeImportSnapshotTasksOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeImportSnapshotTasksOutput:
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

// DescribeImportSnapshotTasksRequest mocked method
func (m *EC2APIMock) DescribeImportSnapshotTasksRequest(p0 *ec2.DescribeImportSnapshotTasksInput) (*request.Request, *ec2.DescribeImportSnapshotTasksOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeImportSnapshotTasksOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeImportSnapshotTasksOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeInstanceAttribute mocked method
func (m *EC2APIMock) DescribeInstanceAttribute(p0 *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeInstanceAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeInstanceAttributeOutput:
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

// DescribeInstanceAttributeRequest mocked method
func (m *EC2APIMock) DescribeInstanceAttributeRequest(p0 *ec2.DescribeInstanceAttributeInput) (*request.Request, *ec2.DescribeInstanceAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeInstanceAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeInstanceAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeInstanceStatus mocked method
func (m *EC2APIMock) DescribeInstanceStatus(p0 *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeInstanceStatusOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeInstanceStatusOutput:
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

// DescribeInstanceStatusPages mocked method
func (m *EC2APIMock) DescribeInstanceStatusPages(p0 *ec2.DescribeInstanceStatusInput, p1 func(*ec2.DescribeInstanceStatusOutput, bool) bool) error {

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

// DescribeInstanceStatusRequest mocked method
func (m *EC2APIMock) DescribeInstanceStatusRequest(p0 *ec2.DescribeInstanceStatusInput) (*request.Request, *ec2.DescribeInstanceStatusOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeInstanceStatusOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeInstanceStatusOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeInstances mocked method
func (m *EC2APIMock) DescribeInstances(p0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeInstancesOutput:
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

// DescribeInstancesPages mocked method
func (m *EC2APIMock) DescribeInstancesPages(p0 *ec2.DescribeInstancesInput, p1 func(*ec2.DescribeInstancesOutput, bool) bool) error {

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

// DescribeInstancesRequest mocked method
func (m *EC2APIMock) DescribeInstancesRequest(p0 *ec2.DescribeInstancesInput) (*request.Request, *ec2.DescribeInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeInternetGateways mocked method
func (m *EC2APIMock) DescribeInternetGateways(p0 *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeInternetGatewaysOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeInternetGatewaysOutput:
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

// DescribeInternetGatewaysRequest mocked method
func (m *EC2APIMock) DescribeInternetGatewaysRequest(p0 *ec2.DescribeInternetGatewaysInput) (*request.Request, *ec2.DescribeInternetGatewaysOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeInternetGatewaysOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeInternetGatewaysOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeKeyPairs mocked method
func (m *EC2APIMock) DescribeKeyPairs(p0 *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeKeyPairsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeKeyPairsOutput:
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

// DescribeKeyPairsRequest mocked method
func (m *EC2APIMock) DescribeKeyPairsRequest(p0 *ec2.DescribeKeyPairsInput) (*request.Request, *ec2.DescribeKeyPairsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeKeyPairsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeKeyPairsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeMovingAddresses mocked method
func (m *EC2APIMock) DescribeMovingAddresses(p0 *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeMovingAddressesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeMovingAddressesOutput:
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

// DescribeMovingAddressesRequest mocked method
func (m *EC2APIMock) DescribeMovingAddressesRequest(p0 *ec2.DescribeMovingAddressesInput) (*request.Request, *ec2.DescribeMovingAddressesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeMovingAddressesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeMovingAddressesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeNatGateways mocked method
func (m *EC2APIMock) DescribeNatGateways(p0 *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeNatGatewaysOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeNatGatewaysOutput:
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

// DescribeNatGatewaysRequest mocked method
func (m *EC2APIMock) DescribeNatGatewaysRequest(p0 *ec2.DescribeNatGatewaysInput) (*request.Request, *ec2.DescribeNatGatewaysOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeNatGatewaysOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeNatGatewaysOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeNetworkAcls mocked method
func (m *EC2APIMock) DescribeNetworkAcls(p0 *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeNetworkAclsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeNetworkAclsOutput:
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

// DescribeNetworkAclsRequest mocked method
func (m *EC2APIMock) DescribeNetworkAclsRequest(p0 *ec2.DescribeNetworkAclsInput) (*request.Request, *ec2.DescribeNetworkAclsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeNetworkAclsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeNetworkAclsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeNetworkInterfaceAttribute mocked method
func (m *EC2APIMock) DescribeNetworkInterfaceAttribute(p0 *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeNetworkInterfaceAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeNetworkInterfaceAttributeOutput:
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

// DescribeNetworkInterfaceAttributeRequest mocked method
func (m *EC2APIMock) DescribeNetworkInterfaceAttributeRequest(p0 *ec2.DescribeNetworkInterfaceAttributeInput) (*request.Request, *ec2.DescribeNetworkInterfaceAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeNetworkInterfaceAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeNetworkInterfaceAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeNetworkInterfaces mocked method
func (m *EC2APIMock) DescribeNetworkInterfaces(p0 *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeNetworkInterfacesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeNetworkInterfacesOutput:
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

// DescribeNetworkInterfacesRequest mocked method
func (m *EC2APIMock) DescribeNetworkInterfacesRequest(p0 *ec2.DescribeNetworkInterfacesInput) (*request.Request, *ec2.DescribeNetworkInterfacesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeNetworkInterfacesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeNetworkInterfacesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribePlacementGroups mocked method
func (m *EC2APIMock) DescribePlacementGroups(p0 *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribePlacementGroupsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribePlacementGroupsOutput:
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

// DescribePlacementGroupsRequest mocked method
func (m *EC2APIMock) DescribePlacementGroupsRequest(p0 *ec2.DescribePlacementGroupsInput) (*request.Request, *ec2.DescribePlacementGroupsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribePlacementGroupsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribePlacementGroupsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribePrefixLists mocked method
func (m *EC2APIMock) DescribePrefixLists(p0 *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribePrefixListsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribePrefixListsOutput:
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

// DescribePrefixListsRequest mocked method
func (m *EC2APIMock) DescribePrefixListsRequest(p0 *ec2.DescribePrefixListsInput) (*request.Request, *ec2.DescribePrefixListsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribePrefixListsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribePrefixListsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeRegions mocked method
func (m *EC2APIMock) DescribeRegions(p0 *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeRegionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeRegionsOutput:
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

// DescribeRegionsRequest mocked method
func (m *EC2APIMock) DescribeRegionsRequest(p0 *ec2.DescribeRegionsInput) (*request.Request, *ec2.DescribeRegionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeRegionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeRegionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeReservedInstances mocked method
func (m *EC2APIMock) DescribeReservedInstances(p0 *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeReservedInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeReservedInstancesOutput:
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

// DescribeReservedInstancesListings mocked method
func (m *EC2APIMock) DescribeReservedInstancesListings(p0 *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeReservedInstancesListingsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeReservedInstancesListingsOutput:
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

// DescribeReservedInstancesListingsRequest mocked method
func (m *EC2APIMock) DescribeReservedInstancesListingsRequest(p0 *ec2.DescribeReservedInstancesListingsInput) (*request.Request, *ec2.DescribeReservedInstancesListingsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeReservedInstancesListingsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeReservedInstancesListingsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeReservedInstancesModifications mocked method
func (m *EC2APIMock) DescribeReservedInstancesModifications(p0 *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeReservedInstancesModificationsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeReservedInstancesModificationsOutput:
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

// DescribeReservedInstancesModificationsPages mocked method
func (m *EC2APIMock) DescribeReservedInstancesModificationsPages(p0 *ec2.DescribeReservedInstancesModificationsInput, p1 func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error {

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

// DescribeReservedInstancesModificationsRequest mocked method
func (m *EC2APIMock) DescribeReservedInstancesModificationsRequest(p0 *ec2.DescribeReservedInstancesModificationsInput) (*request.Request, *ec2.DescribeReservedInstancesModificationsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeReservedInstancesModificationsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeReservedInstancesModificationsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeReservedInstancesOfferings mocked method
func (m *EC2APIMock) DescribeReservedInstancesOfferings(p0 *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeReservedInstancesOfferingsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeReservedInstancesOfferingsOutput:
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

// DescribeReservedInstancesOfferingsPages mocked method
func (m *EC2APIMock) DescribeReservedInstancesOfferingsPages(p0 *ec2.DescribeReservedInstancesOfferingsInput, p1 func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error {

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

// DescribeReservedInstancesOfferingsRequest mocked method
func (m *EC2APIMock) DescribeReservedInstancesOfferingsRequest(p0 *ec2.DescribeReservedInstancesOfferingsInput) (*request.Request, *ec2.DescribeReservedInstancesOfferingsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeReservedInstancesOfferingsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeReservedInstancesOfferingsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeReservedInstancesRequest mocked method
func (m *EC2APIMock) DescribeReservedInstancesRequest(p0 *ec2.DescribeReservedInstancesInput) (*request.Request, *ec2.DescribeReservedInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeReservedInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeReservedInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeRouteTables mocked method
func (m *EC2APIMock) DescribeRouteTables(p0 *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeRouteTablesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeRouteTablesOutput:
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

// DescribeRouteTablesRequest mocked method
func (m *EC2APIMock) DescribeRouteTablesRequest(p0 *ec2.DescribeRouteTablesInput) (*request.Request, *ec2.DescribeRouteTablesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeRouteTablesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeRouteTablesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeScheduledInstanceAvailability mocked method
func (m *EC2APIMock) DescribeScheduledInstanceAvailability(p0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeScheduledInstanceAvailabilityOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeScheduledInstanceAvailabilityOutput:
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

// DescribeScheduledInstanceAvailabilityRequest mocked method
func (m *EC2APIMock) DescribeScheduledInstanceAvailabilityRequest(p0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*request.Request, *ec2.DescribeScheduledInstanceAvailabilityOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeScheduledInstanceAvailabilityOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeScheduledInstanceAvailabilityOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeScheduledInstances mocked method
func (m *EC2APIMock) DescribeScheduledInstances(p0 *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeScheduledInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeScheduledInstancesOutput:
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

// DescribeScheduledInstancesRequest mocked method
func (m *EC2APIMock) DescribeScheduledInstancesRequest(p0 *ec2.DescribeScheduledInstancesInput) (*request.Request, *ec2.DescribeScheduledInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeScheduledInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeScheduledInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSecurityGroupReferences mocked method
func (m *EC2APIMock) DescribeSecurityGroupReferences(p0 *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSecurityGroupReferencesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSecurityGroupReferencesOutput:
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

// DescribeSecurityGroupReferencesRequest mocked method
func (m *EC2APIMock) DescribeSecurityGroupReferencesRequest(p0 *ec2.DescribeSecurityGroupReferencesInput) (*request.Request, *ec2.DescribeSecurityGroupReferencesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSecurityGroupReferencesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSecurityGroupReferencesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSecurityGroups mocked method
func (m *EC2APIMock) DescribeSecurityGroups(p0 *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSecurityGroupsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSecurityGroupsOutput:
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

// DescribeSecurityGroupsRequest mocked method
func (m *EC2APIMock) DescribeSecurityGroupsRequest(p0 *ec2.DescribeSecurityGroupsInput) (*request.Request, *ec2.DescribeSecurityGroupsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSecurityGroupsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSecurityGroupsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSnapshotAttribute mocked method
func (m *EC2APIMock) DescribeSnapshotAttribute(p0 *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSnapshotAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSnapshotAttributeOutput:
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

// DescribeSnapshotAttributeRequest mocked method
func (m *EC2APIMock) DescribeSnapshotAttributeRequest(p0 *ec2.DescribeSnapshotAttributeInput) (*request.Request, *ec2.DescribeSnapshotAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSnapshotAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSnapshotAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSnapshots mocked method
func (m *EC2APIMock) DescribeSnapshots(p0 *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSnapshotsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSnapshotsOutput:
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

// DescribeSnapshotsPages mocked method
func (m *EC2APIMock) DescribeSnapshotsPages(p0 *ec2.DescribeSnapshotsInput, p1 func(*ec2.DescribeSnapshotsOutput, bool) bool) error {

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

// DescribeSnapshotsRequest mocked method
func (m *EC2APIMock) DescribeSnapshotsRequest(p0 *ec2.DescribeSnapshotsInput) (*request.Request, *ec2.DescribeSnapshotsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSnapshotsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSnapshotsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSpotDatafeedSubscription mocked method
func (m *EC2APIMock) DescribeSpotDatafeedSubscription(p0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSpotDatafeedSubscriptionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSpotDatafeedSubscriptionOutput:
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

// DescribeSpotDatafeedSubscriptionRequest mocked method
func (m *EC2APIMock) DescribeSpotDatafeedSubscriptionRequest(p0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DescribeSpotDatafeedSubscriptionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSpotDatafeedSubscriptionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSpotDatafeedSubscriptionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSpotFleetInstances mocked method
func (m *EC2APIMock) DescribeSpotFleetInstances(p0 *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSpotFleetInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSpotFleetInstancesOutput:
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

// DescribeSpotFleetInstancesRequest mocked method
func (m *EC2APIMock) DescribeSpotFleetInstancesRequest(p0 *ec2.DescribeSpotFleetInstancesInput) (*request.Request, *ec2.DescribeSpotFleetInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSpotFleetInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSpotFleetInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSpotFleetRequestHistory mocked method
func (m *EC2APIMock) DescribeSpotFleetRequestHistory(p0 *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSpotFleetRequestHistoryOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSpotFleetRequestHistoryOutput:
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

// DescribeSpotFleetRequestHistoryRequest mocked method
func (m *EC2APIMock) DescribeSpotFleetRequestHistoryRequest(p0 *ec2.DescribeSpotFleetRequestHistoryInput) (*request.Request, *ec2.DescribeSpotFleetRequestHistoryOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSpotFleetRequestHistoryOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSpotFleetRequestHistoryOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSpotFleetRequests mocked method
func (m *EC2APIMock) DescribeSpotFleetRequests(p0 *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSpotFleetRequestsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSpotFleetRequestsOutput:
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

// DescribeSpotFleetRequestsPages mocked method
func (m *EC2APIMock) DescribeSpotFleetRequestsPages(p0 *ec2.DescribeSpotFleetRequestsInput, p1 func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool) error {

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

// DescribeSpotFleetRequestsRequest mocked method
func (m *EC2APIMock) DescribeSpotFleetRequestsRequest(p0 *ec2.DescribeSpotFleetRequestsInput) (*request.Request, *ec2.DescribeSpotFleetRequestsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSpotFleetRequestsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSpotFleetRequestsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSpotInstanceRequests mocked method
func (m *EC2APIMock) DescribeSpotInstanceRequests(p0 *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSpotInstanceRequestsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSpotInstanceRequestsOutput:
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

// DescribeSpotInstanceRequestsRequest mocked method
func (m *EC2APIMock) DescribeSpotInstanceRequestsRequest(p0 *ec2.DescribeSpotInstanceRequestsInput) (*request.Request, *ec2.DescribeSpotInstanceRequestsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSpotInstanceRequestsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSpotInstanceRequestsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSpotPriceHistory mocked method
func (m *EC2APIMock) DescribeSpotPriceHistory(p0 *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSpotPriceHistoryOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSpotPriceHistoryOutput:
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

// DescribeSpotPriceHistoryPages mocked method
func (m *EC2APIMock) DescribeSpotPriceHistoryPages(p0 *ec2.DescribeSpotPriceHistoryInput, p1 func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error {

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

// DescribeSpotPriceHistoryRequest mocked method
func (m *EC2APIMock) DescribeSpotPriceHistoryRequest(p0 *ec2.DescribeSpotPriceHistoryInput) (*request.Request, *ec2.DescribeSpotPriceHistoryOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSpotPriceHistoryOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSpotPriceHistoryOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeStaleSecurityGroups mocked method
func (m *EC2APIMock) DescribeStaleSecurityGroups(p0 *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeStaleSecurityGroupsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeStaleSecurityGroupsOutput:
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

// DescribeStaleSecurityGroupsRequest mocked method
func (m *EC2APIMock) DescribeStaleSecurityGroupsRequest(p0 *ec2.DescribeStaleSecurityGroupsInput) (*request.Request, *ec2.DescribeStaleSecurityGroupsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeStaleSecurityGroupsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeStaleSecurityGroupsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeSubnets mocked method
func (m *EC2APIMock) DescribeSubnets(p0 *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeSubnetsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeSubnetsOutput:
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

// DescribeSubnetsRequest mocked method
func (m *EC2APIMock) DescribeSubnetsRequest(p0 *ec2.DescribeSubnetsInput) (*request.Request, *ec2.DescribeSubnetsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeSubnetsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeSubnetsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeTags mocked method
func (m *EC2APIMock) DescribeTags(p0 *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeTagsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeTagsOutput:
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

// DescribeTagsPages mocked method
func (m *EC2APIMock) DescribeTagsPages(p0 *ec2.DescribeTagsInput, p1 func(*ec2.DescribeTagsOutput, bool) bool) error {

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

// DescribeTagsRequest mocked method
func (m *EC2APIMock) DescribeTagsRequest(p0 *ec2.DescribeTagsInput) (*request.Request, *ec2.DescribeTagsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeTagsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeTagsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVolumeAttribute mocked method
func (m *EC2APIMock) DescribeVolumeAttribute(p0 *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVolumeAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVolumeAttributeOutput:
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

// DescribeVolumeAttributeRequest mocked method
func (m *EC2APIMock) DescribeVolumeAttributeRequest(p0 *ec2.DescribeVolumeAttributeInput) (*request.Request, *ec2.DescribeVolumeAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVolumeAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVolumeAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVolumeStatus mocked method
func (m *EC2APIMock) DescribeVolumeStatus(p0 *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVolumeStatusOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVolumeStatusOutput:
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

// DescribeVolumeStatusPages mocked method
func (m *EC2APIMock) DescribeVolumeStatusPages(p0 *ec2.DescribeVolumeStatusInput, p1 func(*ec2.DescribeVolumeStatusOutput, bool) bool) error {

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

// DescribeVolumeStatusRequest mocked method
func (m *EC2APIMock) DescribeVolumeStatusRequest(p0 *ec2.DescribeVolumeStatusInput) (*request.Request, *ec2.DescribeVolumeStatusOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVolumeStatusOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVolumeStatusOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVolumes mocked method
func (m *EC2APIMock) DescribeVolumes(p0 *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVolumesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVolumesOutput:
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

// DescribeVolumesPages mocked method
func (m *EC2APIMock) DescribeVolumesPages(p0 *ec2.DescribeVolumesInput, p1 func(*ec2.DescribeVolumesOutput, bool) bool) error {

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

// DescribeVolumesRequest mocked method
func (m *EC2APIMock) DescribeVolumesRequest(p0 *ec2.DescribeVolumesInput) (*request.Request, *ec2.DescribeVolumesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVolumesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVolumesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpcAttribute mocked method
func (m *EC2APIMock) DescribeVpcAttribute(p0 *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpcAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpcAttributeOutput:
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

// DescribeVpcAttributeRequest mocked method
func (m *EC2APIMock) DescribeVpcAttributeRequest(p0 *ec2.DescribeVpcAttributeInput) (*request.Request, *ec2.DescribeVpcAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpcAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpcAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpcClassicLink mocked method
func (m *EC2APIMock) DescribeVpcClassicLink(p0 *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpcClassicLinkOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpcClassicLinkOutput:
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

// DescribeVpcClassicLinkDnsSupport mocked method
func (m *EC2APIMock) DescribeVpcClassicLinkDnsSupport(p0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpcClassicLinkDnsSupportOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpcClassicLinkDnsSupportOutput:
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

// DescribeVpcClassicLinkDnsSupportRequest mocked method
func (m *EC2APIMock) DescribeVpcClassicLinkDnsSupportRequest(p0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DescribeVpcClassicLinkDnsSupportOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpcClassicLinkDnsSupportOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpcClassicLinkDnsSupportOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpcClassicLinkRequest mocked method
func (m *EC2APIMock) DescribeVpcClassicLinkRequest(p0 *ec2.DescribeVpcClassicLinkInput) (*request.Request, *ec2.DescribeVpcClassicLinkOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpcClassicLinkOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpcClassicLinkOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpcEndpointServices mocked method
func (m *EC2APIMock) DescribeVpcEndpointServices(p0 *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpcEndpointServicesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpcEndpointServicesOutput:
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

// DescribeVpcEndpointServicesRequest mocked method
func (m *EC2APIMock) DescribeVpcEndpointServicesRequest(p0 *ec2.DescribeVpcEndpointServicesInput) (*request.Request, *ec2.DescribeVpcEndpointServicesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpcEndpointServicesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpcEndpointServicesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpcEndpoints mocked method
func (m *EC2APIMock) DescribeVpcEndpoints(p0 *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpcEndpointsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpcEndpointsOutput:
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

// DescribeVpcEndpointsRequest mocked method
func (m *EC2APIMock) DescribeVpcEndpointsRequest(p0 *ec2.DescribeVpcEndpointsInput) (*request.Request, *ec2.DescribeVpcEndpointsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpcEndpointsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpcEndpointsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpcPeeringConnections mocked method
func (m *EC2APIMock) DescribeVpcPeeringConnections(p0 *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpcPeeringConnectionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpcPeeringConnectionsOutput:
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

// DescribeVpcPeeringConnectionsRequest mocked method
func (m *EC2APIMock) DescribeVpcPeeringConnectionsRequest(p0 *ec2.DescribeVpcPeeringConnectionsInput) (*request.Request, *ec2.DescribeVpcPeeringConnectionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpcPeeringConnectionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpcPeeringConnectionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpcs mocked method
func (m *EC2APIMock) DescribeVpcs(p0 *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpcsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpcsOutput:
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

// DescribeVpcsRequest mocked method
func (m *EC2APIMock) DescribeVpcsRequest(p0 *ec2.DescribeVpcsInput) (*request.Request, *ec2.DescribeVpcsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpcsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpcsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpnConnections mocked method
func (m *EC2APIMock) DescribeVpnConnections(p0 *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpnConnectionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpnConnectionsOutput:
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

// DescribeVpnConnectionsRequest mocked method
func (m *EC2APIMock) DescribeVpnConnectionsRequest(p0 *ec2.DescribeVpnConnectionsInput) (*request.Request, *ec2.DescribeVpnConnectionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpnConnectionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpnConnectionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DescribeVpnGateways mocked method
func (m *EC2APIMock) DescribeVpnGateways(p0 *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DescribeVpnGatewaysOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DescribeVpnGatewaysOutput:
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

// DescribeVpnGatewaysRequest mocked method
func (m *EC2APIMock) DescribeVpnGatewaysRequest(p0 *ec2.DescribeVpnGatewaysInput) (*request.Request, *ec2.DescribeVpnGatewaysOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DescribeVpnGatewaysOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DescribeVpnGatewaysOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachClassicLinkVpc mocked method
func (m *EC2APIMock) DetachClassicLinkVpc(p0 *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DetachClassicLinkVpcOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DetachClassicLinkVpcOutput:
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

// DetachClassicLinkVpcRequest mocked method
func (m *EC2APIMock) DetachClassicLinkVpcRequest(p0 *ec2.DetachClassicLinkVpcInput) (*request.Request, *ec2.DetachClassicLinkVpcOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DetachClassicLinkVpcOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DetachClassicLinkVpcOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachInternetGateway mocked method
func (m *EC2APIMock) DetachInternetGateway(p0 *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DetachInternetGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DetachInternetGatewayOutput:
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

// DetachInternetGatewayRequest mocked method
func (m *EC2APIMock) DetachInternetGatewayRequest(p0 *ec2.DetachInternetGatewayInput) (*request.Request, *ec2.DetachInternetGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DetachInternetGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DetachInternetGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachNetworkInterface mocked method
func (m *EC2APIMock) DetachNetworkInterface(p0 *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DetachNetworkInterfaceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DetachNetworkInterfaceOutput:
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

// DetachNetworkInterfaceRequest mocked method
func (m *EC2APIMock) DetachNetworkInterfaceRequest(p0 *ec2.DetachNetworkInterfaceInput) (*request.Request, *ec2.DetachNetworkInterfaceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DetachNetworkInterfaceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DetachNetworkInterfaceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachVolume mocked method
func (m *EC2APIMock) DetachVolume(p0 *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {

	ret := m.Called(p0)

	var r0 *ec2.VolumeAttachment
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.VolumeAttachment:
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

// DetachVolumeRequest mocked method
func (m *EC2APIMock) DetachVolumeRequest(p0 *ec2.DetachVolumeInput) (*request.Request, *ec2.VolumeAttachment) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.VolumeAttachment
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.VolumeAttachment:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachVpnGateway mocked method
func (m *EC2APIMock) DetachVpnGateway(p0 *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DetachVpnGatewayOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DetachVpnGatewayOutput:
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

// DetachVpnGatewayRequest mocked method
func (m *EC2APIMock) DetachVpnGatewayRequest(p0 *ec2.DetachVpnGatewayInput) (*request.Request, *ec2.DetachVpnGatewayOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DetachVpnGatewayOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DetachVpnGatewayOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DisableVgwRoutePropagation mocked method
func (m *EC2APIMock) DisableVgwRoutePropagation(p0 *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DisableVgwRoutePropagationOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DisableVgwRoutePropagationOutput:
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

// DisableVgwRoutePropagationRequest mocked method
func (m *EC2APIMock) DisableVgwRoutePropagationRequest(p0 *ec2.DisableVgwRoutePropagationInput) (*request.Request, *ec2.DisableVgwRoutePropagationOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DisableVgwRoutePropagationOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DisableVgwRoutePropagationOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DisableVpcClassicLink mocked method
func (m *EC2APIMock) DisableVpcClassicLink(p0 *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DisableVpcClassicLinkOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DisableVpcClassicLinkOutput:
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

// DisableVpcClassicLinkDnsSupport mocked method
func (m *EC2APIMock) DisableVpcClassicLinkDnsSupport(p0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DisableVpcClassicLinkDnsSupportOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DisableVpcClassicLinkDnsSupportOutput:
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

// DisableVpcClassicLinkDnsSupportRequest mocked method
func (m *EC2APIMock) DisableVpcClassicLinkDnsSupportRequest(p0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DisableVpcClassicLinkDnsSupportOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DisableVpcClassicLinkDnsSupportOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DisableVpcClassicLinkDnsSupportOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DisableVpcClassicLinkRequest mocked method
func (m *EC2APIMock) DisableVpcClassicLinkRequest(p0 *ec2.DisableVpcClassicLinkInput) (*request.Request, *ec2.DisableVpcClassicLinkOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DisableVpcClassicLinkOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DisableVpcClassicLinkOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DisassociateAddress mocked method
func (m *EC2APIMock) DisassociateAddress(p0 *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DisassociateAddressOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DisassociateAddressOutput:
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

// DisassociateAddressRequest mocked method
func (m *EC2APIMock) DisassociateAddressRequest(p0 *ec2.DisassociateAddressInput) (*request.Request, *ec2.DisassociateAddressOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DisassociateAddressOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DisassociateAddressOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DisassociateRouteTable mocked method
func (m *EC2APIMock) DisassociateRouteTable(p0 *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.DisassociateRouteTableOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.DisassociateRouteTableOutput:
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

// DisassociateRouteTableRequest mocked method
func (m *EC2APIMock) DisassociateRouteTableRequest(p0 *ec2.DisassociateRouteTableInput) (*request.Request, *ec2.DisassociateRouteTableOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.DisassociateRouteTableOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.DisassociateRouteTableOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EnableVgwRoutePropagation mocked method
func (m *EC2APIMock) EnableVgwRoutePropagation(p0 *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.EnableVgwRoutePropagationOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.EnableVgwRoutePropagationOutput:
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

// EnableVgwRoutePropagationRequest mocked method
func (m *EC2APIMock) EnableVgwRoutePropagationRequest(p0 *ec2.EnableVgwRoutePropagationInput) (*request.Request, *ec2.EnableVgwRoutePropagationOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.EnableVgwRoutePropagationOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.EnableVgwRoutePropagationOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EnableVolumeIO mocked method
func (m *EC2APIMock) EnableVolumeIO(p0 *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.EnableVolumeIOOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.EnableVolumeIOOutput:
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

// EnableVolumeIORequest mocked method
func (m *EC2APIMock) EnableVolumeIORequest(p0 *ec2.EnableVolumeIOInput) (*request.Request, *ec2.EnableVolumeIOOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.EnableVolumeIOOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.EnableVolumeIOOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EnableVpcClassicLink mocked method
func (m *EC2APIMock) EnableVpcClassicLink(p0 *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.EnableVpcClassicLinkOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.EnableVpcClassicLinkOutput:
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

// EnableVpcClassicLinkDnsSupport mocked method
func (m *EC2APIMock) EnableVpcClassicLinkDnsSupport(p0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.EnableVpcClassicLinkDnsSupportOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.EnableVpcClassicLinkDnsSupportOutput:
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

// EnableVpcClassicLinkDnsSupportRequest mocked method
func (m *EC2APIMock) EnableVpcClassicLinkDnsSupportRequest(p0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.EnableVpcClassicLinkDnsSupportOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.EnableVpcClassicLinkDnsSupportOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.EnableVpcClassicLinkDnsSupportOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EnableVpcClassicLinkRequest mocked method
func (m *EC2APIMock) EnableVpcClassicLinkRequest(p0 *ec2.EnableVpcClassicLinkInput) (*request.Request, *ec2.EnableVpcClassicLinkOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.EnableVpcClassicLinkOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.EnableVpcClassicLinkOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetConsoleOutput mocked method
func (m *EC2APIMock) GetConsoleOutput(p0 *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.GetConsoleOutputOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.GetConsoleOutputOutput:
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

// GetConsoleOutputRequest mocked method
func (m *EC2APIMock) GetConsoleOutputRequest(p0 *ec2.GetConsoleOutputInput) (*request.Request, *ec2.GetConsoleOutputOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.GetConsoleOutputOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.GetConsoleOutputOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetConsoleScreenshot mocked method
func (m *EC2APIMock) GetConsoleScreenshot(p0 *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.GetConsoleScreenshotOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.GetConsoleScreenshotOutput:
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

// GetConsoleScreenshotRequest mocked method
func (m *EC2APIMock) GetConsoleScreenshotRequest(p0 *ec2.GetConsoleScreenshotInput) (*request.Request, *ec2.GetConsoleScreenshotOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.GetConsoleScreenshotOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.GetConsoleScreenshotOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetHostReservationPurchasePreview mocked method
func (m *EC2APIMock) GetHostReservationPurchasePreview(p0 *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.GetHostReservationPurchasePreviewOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.GetHostReservationPurchasePreviewOutput:
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

// GetHostReservationPurchasePreviewRequest mocked method
func (m *EC2APIMock) GetHostReservationPurchasePreviewRequest(p0 *ec2.GetHostReservationPurchasePreviewInput) (*request.Request, *ec2.GetHostReservationPurchasePreviewOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.GetHostReservationPurchasePreviewOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.GetHostReservationPurchasePreviewOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetPasswordData mocked method
func (m *EC2APIMock) GetPasswordData(p0 *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.GetPasswordDataOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.GetPasswordDataOutput:
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

// GetPasswordDataRequest mocked method
func (m *EC2APIMock) GetPasswordDataRequest(p0 *ec2.GetPasswordDataInput) (*request.Request, *ec2.GetPasswordDataOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.GetPasswordDataOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.GetPasswordDataOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetReservedInstancesExchangeQuote mocked method
func (m *EC2APIMock) GetReservedInstancesExchangeQuote(p0 *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.GetReservedInstancesExchangeQuoteOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.GetReservedInstancesExchangeQuoteOutput:
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

// GetReservedInstancesExchangeQuoteRequest mocked method
func (m *EC2APIMock) GetReservedInstancesExchangeQuoteRequest(p0 *ec2.GetReservedInstancesExchangeQuoteInput) (*request.Request, *ec2.GetReservedInstancesExchangeQuoteOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.GetReservedInstancesExchangeQuoteOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.GetReservedInstancesExchangeQuoteOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ImportImage mocked method
func (m *EC2APIMock) ImportImage(p0 *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ImportImageOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ImportImageOutput:
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

// ImportImageRequest mocked method
func (m *EC2APIMock) ImportImageRequest(p0 *ec2.ImportImageInput) (*request.Request, *ec2.ImportImageOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ImportImageOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ImportImageOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ImportInstance mocked method
func (m *EC2APIMock) ImportInstance(p0 *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ImportInstanceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ImportInstanceOutput:
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

// ImportInstanceRequest mocked method
func (m *EC2APIMock) ImportInstanceRequest(p0 *ec2.ImportInstanceInput) (*request.Request, *ec2.ImportInstanceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ImportInstanceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ImportInstanceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ImportKeyPair mocked method
func (m *EC2APIMock) ImportKeyPair(p0 *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ImportKeyPairOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ImportKeyPairOutput:
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

// ImportKeyPairRequest mocked method
func (m *EC2APIMock) ImportKeyPairRequest(p0 *ec2.ImportKeyPairInput) (*request.Request, *ec2.ImportKeyPairOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ImportKeyPairOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ImportKeyPairOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ImportSnapshot mocked method
func (m *EC2APIMock) ImportSnapshot(p0 *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ImportSnapshotOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ImportSnapshotOutput:
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

// ImportSnapshotRequest mocked method
func (m *EC2APIMock) ImportSnapshotRequest(p0 *ec2.ImportSnapshotInput) (*request.Request, *ec2.ImportSnapshotOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ImportSnapshotOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ImportSnapshotOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ImportVolume mocked method
func (m *EC2APIMock) ImportVolume(p0 *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ImportVolumeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ImportVolumeOutput:
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

// ImportVolumeRequest mocked method
func (m *EC2APIMock) ImportVolumeRequest(p0 *ec2.ImportVolumeInput) (*request.Request, *ec2.ImportVolumeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ImportVolumeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ImportVolumeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyHosts mocked method
func (m *EC2APIMock) ModifyHosts(p0 *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyHostsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyHostsOutput:
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

// ModifyHostsRequest mocked method
func (m *EC2APIMock) ModifyHostsRequest(p0 *ec2.ModifyHostsInput) (*request.Request, *ec2.ModifyHostsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyHostsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyHostsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyIdFormat mocked method
func (m *EC2APIMock) ModifyIdFormat(p0 *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyIdFormatOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyIdFormatOutput:
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

// ModifyIdFormatRequest mocked method
func (m *EC2APIMock) ModifyIdFormatRequest(p0 *ec2.ModifyIdFormatInput) (*request.Request, *ec2.ModifyIdFormatOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyIdFormatOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyIdFormatOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyIdentityIdFormat mocked method
func (m *EC2APIMock) ModifyIdentityIdFormat(p0 *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyIdentityIdFormatOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyIdentityIdFormatOutput:
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

// ModifyIdentityIdFormatRequest mocked method
func (m *EC2APIMock) ModifyIdentityIdFormatRequest(p0 *ec2.ModifyIdentityIdFormatInput) (*request.Request, *ec2.ModifyIdentityIdFormatOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyIdentityIdFormatOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyIdentityIdFormatOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyImageAttribute mocked method
func (m *EC2APIMock) ModifyImageAttribute(p0 *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyImageAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyImageAttributeOutput:
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

// ModifyImageAttributeRequest mocked method
func (m *EC2APIMock) ModifyImageAttributeRequest(p0 *ec2.ModifyImageAttributeInput) (*request.Request, *ec2.ModifyImageAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyImageAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyImageAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyInstanceAttribute mocked method
func (m *EC2APIMock) ModifyInstanceAttribute(p0 *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyInstanceAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyInstanceAttributeOutput:
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

// ModifyInstanceAttributeRequest mocked method
func (m *EC2APIMock) ModifyInstanceAttributeRequest(p0 *ec2.ModifyInstanceAttributeInput) (*request.Request, *ec2.ModifyInstanceAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyInstanceAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyInstanceAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyInstancePlacement mocked method
func (m *EC2APIMock) ModifyInstancePlacement(p0 *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyInstancePlacementOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyInstancePlacementOutput:
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

// ModifyInstancePlacementRequest mocked method
func (m *EC2APIMock) ModifyInstancePlacementRequest(p0 *ec2.ModifyInstancePlacementInput) (*request.Request, *ec2.ModifyInstancePlacementOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyInstancePlacementOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyInstancePlacementOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyNetworkInterfaceAttribute mocked method
func (m *EC2APIMock) ModifyNetworkInterfaceAttribute(p0 *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyNetworkInterfaceAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyNetworkInterfaceAttributeOutput:
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

// ModifyNetworkInterfaceAttributeRequest mocked method
func (m *EC2APIMock) ModifyNetworkInterfaceAttributeRequest(p0 *ec2.ModifyNetworkInterfaceAttributeInput) (*request.Request, *ec2.ModifyNetworkInterfaceAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyNetworkInterfaceAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyNetworkInterfaceAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyReservedInstances mocked method
func (m *EC2APIMock) ModifyReservedInstances(p0 *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyReservedInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyReservedInstancesOutput:
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

// ModifyReservedInstancesRequest mocked method
func (m *EC2APIMock) ModifyReservedInstancesRequest(p0 *ec2.ModifyReservedInstancesInput) (*request.Request, *ec2.ModifyReservedInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyReservedInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyReservedInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifySnapshotAttribute mocked method
func (m *EC2APIMock) ModifySnapshotAttribute(p0 *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifySnapshotAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifySnapshotAttributeOutput:
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

// ModifySnapshotAttributeRequest mocked method
func (m *EC2APIMock) ModifySnapshotAttributeRequest(p0 *ec2.ModifySnapshotAttributeInput) (*request.Request, *ec2.ModifySnapshotAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifySnapshotAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifySnapshotAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifySpotFleetRequest mocked method
func (m *EC2APIMock) ModifySpotFleetRequest(p0 *ec2.ModifySpotFleetRequestInput) (*ec2.ModifySpotFleetRequestOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifySpotFleetRequestOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifySpotFleetRequestOutput:
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

// ModifySpotFleetRequestRequest mocked method
func (m *EC2APIMock) ModifySpotFleetRequestRequest(p0 *ec2.ModifySpotFleetRequestInput) (*request.Request, *ec2.ModifySpotFleetRequestOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifySpotFleetRequestOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifySpotFleetRequestOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifySubnetAttribute mocked method
func (m *EC2APIMock) ModifySubnetAttribute(p0 *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifySubnetAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifySubnetAttributeOutput:
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

// ModifySubnetAttributeRequest mocked method
func (m *EC2APIMock) ModifySubnetAttributeRequest(p0 *ec2.ModifySubnetAttributeInput) (*request.Request, *ec2.ModifySubnetAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifySubnetAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifySubnetAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyVolumeAttribute mocked method
func (m *EC2APIMock) ModifyVolumeAttribute(p0 *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyVolumeAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyVolumeAttributeOutput:
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

// ModifyVolumeAttributeRequest mocked method
func (m *EC2APIMock) ModifyVolumeAttributeRequest(p0 *ec2.ModifyVolumeAttributeInput) (*request.Request, *ec2.ModifyVolumeAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyVolumeAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyVolumeAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyVpcAttribute mocked method
func (m *EC2APIMock) ModifyVpcAttribute(p0 *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyVpcAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyVpcAttributeOutput:
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

// ModifyVpcAttributeRequest mocked method
func (m *EC2APIMock) ModifyVpcAttributeRequest(p0 *ec2.ModifyVpcAttributeInput) (*request.Request, *ec2.ModifyVpcAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyVpcAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyVpcAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyVpcEndpoint mocked method
func (m *EC2APIMock) ModifyVpcEndpoint(p0 *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyVpcEndpointOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyVpcEndpointOutput:
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

// ModifyVpcEndpointRequest mocked method
func (m *EC2APIMock) ModifyVpcEndpointRequest(p0 *ec2.ModifyVpcEndpointInput) (*request.Request, *ec2.ModifyVpcEndpointOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyVpcEndpointOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyVpcEndpointOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ModifyVpcPeeringConnectionOptions mocked method
func (m *EC2APIMock) ModifyVpcPeeringConnectionOptions(p0 *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ModifyVpcPeeringConnectionOptionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ModifyVpcPeeringConnectionOptionsOutput:
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

// ModifyVpcPeeringConnectionOptionsRequest mocked method
func (m *EC2APIMock) ModifyVpcPeeringConnectionOptionsRequest(p0 *ec2.ModifyVpcPeeringConnectionOptionsInput) (*request.Request, *ec2.ModifyVpcPeeringConnectionOptionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ModifyVpcPeeringConnectionOptionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ModifyVpcPeeringConnectionOptionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// MonitorInstances mocked method
func (m *EC2APIMock) MonitorInstances(p0 *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.MonitorInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.MonitorInstancesOutput:
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

// MonitorInstancesRequest mocked method
func (m *EC2APIMock) MonitorInstancesRequest(p0 *ec2.MonitorInstancesInput) (*request.Request, *ec2.MonitorInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.MonitorInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.MonitorInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// MoveAddressToVpc mocked method
func (m *EC2APIMock) MoveAddressToVpc(p0 *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.MoveAddressToVpcOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.MoveAddressToVpcOutput:
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

// MoveAddressToVpcRequest mocked method
func (m *EC2APIMock) MoveAddressToVpcRequest(p0 *ec2.MoveAddressToVpcInput) (*request.Request, *ec2.MoveAddressToVpcOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.MoveAddressToVpcOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.MoveAddressToVpcOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PurchaseHostReservation mocked method
func (m *EC2APIMock) PurchaseHostReservation(p0 *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.PurchaseHostReservationOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.PurchaseHostReservationOutput:
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

// PurchaseHostReservationRequest mocked method
func (m *EC2APIMock) PurchaseHostReservationRequest(p0 *ec2.PurchaseHostReservationInput) (*request.Request, *ec2.PurchaseHostReservationOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.PurchaseHostReservationOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.PurchaseHostReservationOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PurchaseReservedInstancesOffering mocked method
func (m *EC2APIMock) PurchaseReservedInstancesOffering(p0 *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.PurchaseReservedInstancesOfferingOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.PurchaseReservedInstancesOfferingOutput:
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

// PurchaseReservedInstancesOfferingRequest mocked method
func (m *EC2APIMock) PurchaseReservedInstancesOfferingRequest(p0 *ec2.PurchaseReservedInstancesOfferingInput) (*request.Request, *ec2.PurchaseReservedInstancesOfferingOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.PurchaseReservedInstancesOfferingOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.PurchaseReservedInstancesOfferingOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PurchaseScheduledInstances mocked method
func (m *EC2APIMock) PurchaseScheduledInstances(p0 *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.PurchaseScheduledInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.PurchaseScheduledInstancesOutput:
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

// PurchaseScheduledInstancesRequest mocked method
func (m *EC2APIMock) PurchaseScheduledInstancesRequest(p0 *ec2.PurchaseScheduledInstancesInput) (*request.Request, *ec2.PurchaseScheduledInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.PurchaseScheduledInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.PurchaseScheduledInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RebootInstances mocked method
func (m *EC2APIMock) RebootInstances(p0 *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RebootInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RebootInstancesOutput:
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

// RebootInstancesRequest mocked method
func (m *EC2APIMock) RebootInstancesRequest(p0 *ec2.RebootInstancesInput) (*request.Request, *ec2.RebootInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RebootInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RebootInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RegisterImage mocked method
func (m *EC2APIMock) RegisterImage(p0 *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RegisterImageOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RegisterImageOutput:
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

// RegisterImageRequest mocked method
func (m *EC2APIMock) RegisterImageRequest(p0 *ec2.RegisterImageInput) (*request.Request, *ec2.RegisterImageOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RegisterImageOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RegisterImageOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RejectVpcPeeringConnection mocked method
func (m *EC2APIMock) RejectVpcPeeringConnection(p0 *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RejectVpcPeeringConnectionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RejectVpcPeeringConnectionOutput:
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

// RejectVpcPeeringConnectionRequest mocked method
func (m *EC2APIMock) RejectVpcPeeringConnectionRequest(p0 *ec2.RejectVpcPeeringConnectionInput) (*request.Request, *ec2.RejectVpcPeeringConnectionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RejectVpcPeeringConnectionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RejectVpcPeeringConnectionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ReleaseAddress mocked method
func (m *EC2APIMock) ReleaseAddress(p0 *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ReleaseAddressOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ReleaseAddressOutput:
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

// ReleaseAddressRequest mocked method
func (m *EC2APIMock) ReleaseAddressRequest(p0 *ec2.ReleaseAddressInput) (*request.Request, *ec2.ReleaseAddressOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ReleaseAddressOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ReleaseAddressOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ReleaseHosts mocked method
func (m *EC2APIMock) ReleaseHosts(p0 *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ReleaseHostsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ReleaseHostsOutput:
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

// ReleaseHostsRequest mocked method
func (m *EC2APIMock) ReleaseHostsRequest(p0 *ec2.ReleaseHostsInput) (*request.Request, *ec2.ReleaseHostsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ReleaseHostsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ReleaseHostsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ReplaceNetworkAclAssociation mocked method
func (m *EC2APIMock) ReplaceNetworkAclAssociation(p0 *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ReplaceNetworkAclAssociationOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ReplaceNetworkAclAssociationOutput:
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

// ReplaceNetworkAclAssociationRequest mocked method
func (m *EC2APIMock) ReplaceNetworkAclAssociationRequest(p0 *ec2.ReplaceNetworkAclAssociationInput) (*request.Request, *ec2.ReplaceNetworkAclAssociationOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ReplaceNetworkAclAssociationOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ReplaceNetworkAclAssociationOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ReplaceNetworkAclEntry mocked method
func (m *EC2APIMock) ReplaceNetworkAclEntry(p0 *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ReplaceNetworkAclEntryOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ReplaceNetworkAclEntryOutput:
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

// ReplaceNetworkAclEntryRequest mocked method
func (m *EC2APIMock) ReplaceNetworkAclEntryRequest(p0 *ec2.ReplaceNetworkAclEntryInput) (*request.Request, *ec2.ReplaceNetworkAclEntryOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ReplaceNetworkAclEntryOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ReplaceNetworkAclEntryOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ReplaceRoute mocked method
func (m *EC2APIMock) ReplaceRoute(p0 *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ReplaceRouteOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ReplaceRouteOutput:
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

// ReplaceRouteRequest mocked method
func (m *EC2APIMock) ReplaceRouteRequest(p0 *ec2.ReplaceRouteInput) (*request.Request, *ec2.ReplaceRouteOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ReplaceRouteOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ReplaceRouteOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ReplaceRouteTableAssociation mocked method
func (m *EC2APIMock) ReplaceRouteTableAssociation(p0 *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ReplaceRouteTableAssociationOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ReplaceRouteTableAssociationOutput:
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

// ReplaceRouteTableAssociationRequest mocked method
func (m *EC2APIMock) ReplaceRouteTableAssociationRequest(p0 *ec2.ReplaceRouteTableAssociationInput) (*request.Request, *ec2.ReplaceRouteTableAssociationOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ReplaceRouteTableAssociationOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ReplaceRouteTableAssociationOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ReportInstanceStatus mocked method
func (m *EC2APIMock) ReportInstanceStatus(p0 *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ReportInstanceStatusOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ReportInstanceStatusOutput:
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

// ReportInstanceStatusRequest mocked method
func (m *EC2APIMock) ReportInstanceStatusRequest(p0 *ec2.ReportInstanceStatusInput) (*request.Request, *ec2.ReportInstanceStatusOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ReportInstanceStatusOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ReportInstanceStatusOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RequestSpotFleet mocked method
func (m *EC2APIMock) RequestSpotFleet(p0 *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RequestSpotFleetOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RequestSpotFleetOutput:
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

// RequestSpotFleetRequest mocked method
func (m *EC2APIMock) RequestSpotFleetRequest(p0 *ec2.RequestSpotFleetInput) (*request.Request, *ec2.RequestSpotFleetOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RequestSpotFleetOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RequestSpotFleetOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RequestSpotInstances mocked method
func (m *EC2APIMock) RequestSpotInstances(p0 *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RequestSpotInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RequestSpotInstancesOutput:
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

// RequestSpotInstancesRequest mocked method
func (m *EC2APIMock) RequestSpotInstancesRequest(p0 *ec2.RequestSpotInstancesInput) (*request.Request, *ec2.RequestSpotInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RequestSpotInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RequestSpotInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResetImageAttribute mocked method
func (m *EC2APIMock) ResetImageAttribute(p0 *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ResetImageAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ResetImageAttributeOutput:
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

// ResetImageAttributeRequest mocked method
func (m *EC2APIMock) ResetImageAttributeRequest(p0 *ec2.ResetImageAttributeInput) (*request.Request, *ec2.ResetImageAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ResetImageAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ResetImageAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResetInstanceAttribute mocked method
func (m *EC2APIMock) ResetInstanceAttribute(p0 *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ResetInstanceAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ResetInstanceAttributeOutput:
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

// ResetInstanceAttributeRequest mocked method
func (m *EC2APIMock) ResetInstanceAttributeRequest(p0 *ec2.ResetInstanceAttributeInput) (*request.Request, *ec2.ResetInstanceAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ResetInstanceAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ResetInstanceAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResetNetworkInterfaceAttribute mocked method
func (m *EC2APIMock) ResetNetworkInterfaceAttribute(p0 *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ResetNetworkInterfaceAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ResetNetworkInterfaceAttributeOutput:
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

// ResetNetworkInterfaceAttributeRequest mocked method
func (m *EC2APIMock) ResetNetworkInterfaceAttributeRequest(p0 *ec2.ResetNetworkInterfaceAttributeInput) (*request.Request, *ec2.ResetNetworkInterfaceAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ResetNetworkInterfaceAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ResetNetworkInterfaceAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResetSnapshotAttribute mocked method
func (m *EC2APIMock) ResetSnapshotAttribute(p0 *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.ResetSnapshotAttributeOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.ResetSnapshotAttributeOutput:
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

// ResetSnapshotAttributeRequest mocked method
func (m *EC2APIMock) ResetSnapshotAttributeRequest(p0 *ec2.ResetSnapshotAttributeInput) (*request.Request, *ec2.ResetSnapshotAttributeOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.ResetSnapshotAttributeOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.ResetSnapshotAttributeOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RestoreAddressToClassic mocked method
func (m *EC2APIMock) RestoreAddressToClassic(p0 *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RestoreAddressToClassicOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RestoreAddressToClassicOutput:
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

// RestoreAddressToClassicRequest mocked method
func (m *EC2APIMock) RestoreAddressToClassicRequest(p0 *ec2.RestoreAddressToClassicInput) (*request.Request, *ec2.RestoreAddressToClassicOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RestoreAddressToClassicOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RestoreAddressToClassicOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RevokeSecurityGroupEgress mocked method
func (m *EC2APIMock) RevokeSecurityGroupEgress(p0 *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RevokeSecurityGroupEgressOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RevokeSecurityGroupEgressOutput:
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

// RevokeSecurityGroupEgressRequest mocked method
func (m *EC2APIMock) RevokeSecurityGroupEgressRequest(p0 *ec2.RevokeSecurityGroupEgressInput) (*request.Request, *ec2.RevokeSecurityGroupEgressOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RevokeSecurityGroupEgressOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RevokeSecurityGroupEgressOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RevokeSecurityGroupIngress mocked method
func (m *EC2APIMock) RevokeSecurityGroupIngress(p0 *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RevokeSecurityGroupIngressOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RevokeSecurityGroupIngressOutput:
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

// RevokeSecurityGroupIngressRequest mocked method
func (m *EC2APIMock) RevokeSecurityGroupIngressRequest(p0 *ec2.RevokeSecurityGroupIngressInput) (*request.Request, *ec2.RevokeSecurityGroupIngressOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RevokeSecurityGroupIngressOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RevokeSecurityGroupIngressOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RunInstances mocked method
func (m *EC2APIMock) RunInstances(p0 *ec2.RunInstancesInput) (*ec2.Reservation, error) {

	ret := m.Called(p0)

	var r0 *ec2.Reservation
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.Reservation:
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

// RunInstancesRequest mocked method
func (m *EC2APIMock) RunInstancesRequest(p0 *ec2.RunInstancesInput) (*request.Request, *ec2.Reservation) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.Reservation
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.Reservation:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RunScheduledInstances mocked method
func (m *EC2APIMock) RunScheduledInstances(p0 *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.RunScheduledInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.RunScheduledInstancesOutput:
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

// RunScheduledInstancesRequest mocked method
func (m *EC2APIMock) RunScheduledInstancesRequest(p0 *ec2.RunScheduledInstancesInput) (*request.Request, *ec2.RunScheduledInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.RunScheduledInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.RunScheduledInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// StartInstances mocked method
func (m *EC2APIMock) StartInstances(p0 *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.StartInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.StartInstancesOutput:
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

// StartInstancesRequest mocked method
func (m *EC2APIMock) StartInstancesRequest(p0 *ec2.StartInstancesInput) (*request.Request, *ec2.StartInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.StartInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.StartInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// StopInstances mocked method
func (m *EC2APIMock) StopInstances(p0 *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.StopInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.StopInstancesOutput:
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

// StopInstancesRequest mocked method
func (m *EC2APIMock) StopInstancesRequest(p0 *ec2.StopInstancesInput) (*request.Request, *ec2.StopInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.StopInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.StopInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// TerminateInstances mocked method
func (m *EC2APIMock) TerminateInstances(p0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.TerminateInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.TerminateInstancesOutput:
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

// TerminateInstancesRequest mocked method
func (m *EC2APIMock) TerminateInstancesRequest(p0 *ec2.TerminateInstancesInput) (*request.Request, *ec2.TerminateInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.TerminateInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.TerminateInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UnassignPrivateIpAddresses mocked method
func (m *EC2APIMock) UnassignPrivateIpAddresses(p0 *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.UnassignPrivateIpAddressesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.UnassignPrivateIpAddressesOutput:
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

// UnassignPrivateIpAddressesRequest mocked method
func (m *EC2APIMock) UnassignPrivateIpAddressesRequest(p0 *ec2.UnassignPrivateIpAddressesInput) (*request.Request, *ec2.UnassignPrivateIpAddressesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.UnassignPrivateIpAddressesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.UnassignPrivateIpAddressesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UnmonitorInstances mocked method
func (m *EC2APIMock) UnmonitorInstances(p0 *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {

	ret := m.Called(p0)

	var r0 *ec2.UnmonitorInstancesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *ec2.UnmonitorInstancesOutput:
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

// UnmonitorInstancesRequest mocked method
func (m *EC2APIMock) UnmonitorInstancesRequest(p0 *ec2.UnmonitorInstancesInput) (*request.Request, *ec2.UnmonitorInstancesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *ec2.UnmonitorInstancesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *ec2.UnmonitorInstancesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// WaitUntilBundleTaskComplete mocked method
func (m *EC2APIMock) WaitUntilBundleTaskComplete(p0 *ec2.DescribeBundleTasksInput) error {

	ret := m.Called(p0)

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

// WaitUntilConversionTaskCancelled mocked method
func (m *EC2APIMock) WaitUntilConversionTaskCancelled(p0 *ec2.DescribeConversionTasksInput) error {

	ret := m.Called(p0)

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

// WaitUntilConversionTaskCompleted mocked method
func (m *EC2APIMock) WaitUntilConversionTaskCompleted(p0 *ec2.DescribeConversionTasksInput) error {

	ret := m.Called(p0)

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

// WaitUntilConversionTaskDeleted mocked method
func (m *EC2APIMock) WaitUntilConversionTaskDeleted(p0 *ec2.DescribeConversionTasksInput) error {

	ret := m.Called(p0)

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

// WaitUntilCustomerGatewayAvailable mocked method
func (m *EC2APIMock) WaitUntilCustomerGatewayAvailable(p0 *ec2.DescribeCustomerGatewaysInput) error {

	ret := m.Called(p0)

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

// WaitUntilExportTaskCancelled mocked method
func (m *EC2APIMock) WaitUntilExportTaskCancelled(p0 *ec2.DescribeExportTasksInput) error {

	ret := m.Called(p0)

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

// WaitUntilExportTaskCompleted mocked method
func (m *EC2APIMock) WaitUntilExportTaskCompleted(p0 *ec2.DescribeExportTasksInput) error {

	ret := m.Called(p0)

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

// WaitUntilImageAvailable mocked method
func (m *EC2APIMock) WaitUntilImageAvailable(p0 *ec2.DescribeImagesInput) error {

	ret := m.Called(p0)

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

// WaitUntilImageExists mocked method
func (m *EC2APIMock) WaitUntilImageExists(p0 *ec2.DescribeImagesInput) error {

	ret := m.Called(p0)

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

// WaitUntilInstanceExists mocked method
func (m *EC2APIMock) WaitUntilInstanceExists(p0 *ec2.DescribeInstancesInput) error {

	ret := m.Called(p0)

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

// WaitUntilInstanceRunning mocked method
func (m *EC2APIMock) WaitUntilInstanceRunning(p0 *ec2.DescribeInstancesInput) error {

	ret := m.Called(p0)

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

// WaitUntilInstanceStatusOk mocked method
func (m *EC2APIMock) WaitUntilInstanceStatusOk(p0 *ec2.DescribeInstanceStatusInput) error {

	ret := m.Called(p0)

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

// WaitUntilInstanceStopped mocked method
func (m *EC2APIMock) WaitUntilInstanceStopped(p0 *ec2.DescribeInstancesInput) error {

	ret := m.Called(p0)

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

// WaitUntilInstanceTerminated mocked method
func (m *EC2APIMock) WaitUntilInstanceTerminated(p0 *ec2.DescribeInstancesInput) error {

	ret := m.Called(p0)

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

// WaitUntilKeyPairExists mocked method
func (m *EC2APIMock) WaitUntilKeyPairExists(p0 *ec2.DescribeKeyPairsInput) error {

	ret := m.Called(p0)

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

// WaitUntilNatGatewayAvailable mocked method
func (m *EC2APIMock) WaitUntilNatGatewayAvailable(p0 *ec2.DescribeNatGatewaysInput) error {

	ret := m.Called(p0)

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

// WaitUntilNetworkInterfaceAvailable mocked method
func (m *EC2APIMock) WaitUntilNetworkInterfaceAvailable(p0 *ec2.DescribeNetworkInterfacesInput) error {

	ret := m.Called(p0)

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

// WaitUntilPasswordDataAvailable mocked method
func (m *EC2APIMock) WaitUntilPasswordDataAvailable(p0 *ec2.GetPasswordDataInput) error {

	ret := m.Called(p0)

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

// WaitUntilSnapshotCompleted mocked method
func (m *EC2APIMock) WaitUntilSnapshotCompleted(p0 *ec2.DescribeSnapshotsInput) error {

	ret := m.Called(p0)

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

// WaitUntilSpotInstanceRequestFulfilled mocked method
func (m *EC2APIMock) WaitUntilSpotInstanceRequestFulfilled(p0 *ec2.DescribeSpotInstanceRequestsInput) error {

	ret := m.Called(p0)

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

// WaitUntilSubnetAvailable mocked method
func (m *EC2APIMock) WaitUntilSubnetAvailable(p0 *ec2.DescribeSubnetsInput) error {

	ret := m.Called(p0)

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

// WaitUntilSystemStatusOk mocked method
func (m *EC2APIMock) WaitUntilSystemStatusOk(p0 *ec2.DescribeInstanceStatusInput) error {

	ret := m.Called(p0)

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

// WaitUntilVolumeAvailable mocked method
func (m *EC2APIMock) WaitUntilVolumeAvailable(p0 *ec2.DescribeVolumesInput) error {

	ret := m.Called(p0)

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

// WaitUntilVolumeDeleted mocked method
func (m *EC2APIMock) WaitUntilVolumeDeleted(p0 *ec2.DescribeVolumesInput) error {

	ret := m.Called(p0)

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

// WaitUntilVolumeInUse mocked method
func (m *EC2APIMock) WaitUntilVolumeInUse(p0 *ec2.DescribeVolumesInput) error {

	ret := m.Called(p0)

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

// WaitUntilVpcAvailable mocked method
func (m *EC2APIMock) WaitUntilVpcAvailable(p0 *ec2.DescribeVpcsInput) error {

	ret := m.Called(p0)

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

// WaitUntilVpcExists mocked method
func (m *EC2APIMock) WaitUntilVpcExists(p0 *ec2.DescribeVpcsInput) error {

	ret := m.Called(p0)

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

// WaitUntilVpcPeeringConnectionExists mocked method
func (m *EC2APIMock) WaitUntilVpcPeeringConnectionExists(p0 *ec2.DescribeVpcPeeringConnectionsInput) error {

	ret := m.Called(p0)

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

// WaitUntilVpnConnectionAvailable mocked method
func (m *EC2APIMock) WaitUntilVpnConnectionAvailable(p0 *ec2.DescribeVpnConnectionsInput) error {

	ret := m.Called(p0)

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

// WaitUntilVpnConnectionDeleted mocked method
func (m *EC2APIMock) WaitUntilVpnConnectionDeleted(p0 *ec2.DescribeVpnConnectionsInput) error {

	ret := m.Called(p0)

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
