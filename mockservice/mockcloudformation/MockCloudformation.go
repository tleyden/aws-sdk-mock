package mockcloudformation

import "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"

//go:generate goautomock -template=testify CloudFormationAPI


// Wrap the official cloudformationiface.CloudFormationAPI so we can use goautomock to generate
// Mocks for it
type CloudFormationAPI interface {

	cloudformationiface.CloudFormationAPI


}



