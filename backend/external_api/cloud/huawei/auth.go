package huawei

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
)

func NewClient(accessKey string, secretKey string, region string, projectId string) basic.Credentials {
	auth := basic.NewCredentialsBuilder().WithAk(accessKey).WithSk(secretKey).WithProjectId(projectId).Build()
	return auth
}
