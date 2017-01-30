package mockcloudformation

import "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"

// See README.md for instructions on how to generate mocks.  Requires tweaks
// above and beyond running "go generate"

//go:generate goautomock -template=testify CloudFormationAPI

// Wrap the official cloudformationiface.CloudFormationAPI so we can use goautomock to generate mocks
// 
// Initially I tried running go:generate goautomock -template=testify cloudformationiface.CloudFormationAPI 
// but ran into errors, hence this workaround
// 
type CloudFormationAPI interface {

	cloudformationiface.CloudFormationAPI

}



