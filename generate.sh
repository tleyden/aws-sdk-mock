
goautomock -template=testify -o "MockCloudFormation.go" -mock-pkg mockcloudformation -pkg=github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface CloudFormationAPI
