
goautomock -template=testify -o "MockCloudFormation.go" -mock-pkg mockcloudformation -pkg=github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface CloudFormationAPI
mv mock_cloud_formation.go mockcloudformation/

goautomock -template=testify -o "MockEC2.go" -mock-pkg mockec2 -pkg=github.com/aws/aws-sdk-go/service/ec2/ec2iface EC2API
mv mock_e_c2.go mockec2/

# Hand editing required:
#
# type CloudFormationAPIMock struct {
#     mock.Mock
#     cloudformationiface.CloudFormationAPI  // temp workaround for https://github.com/ernesto-jimenez/goautomock/issues/3#issuecomment-276240746
# }
#
# type EC2APIMock struct {
#    mock.Mock
#    ec2iface.EC2API // temp workaround for https://github.com/ernesto-jimenez/goautomock/issues/3#issuecomment-276240746
# }