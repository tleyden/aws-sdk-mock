
## Instructions to regenerate mocks

- Run `go generate`
- Rename generated file to `AutoMockCloudformation.go` (so it's not in a `_test` file)
- Hand edit `CloudFormationAPIMock` struct to embed `cloudformationiface.CloudFormationAPI`, since it doesn't generate methods for _all_ methods in interface

```
type CloudFormationAPIMock struct {
	mock.Mock
	cloudformationiface.CloudFormationAPI
}
```