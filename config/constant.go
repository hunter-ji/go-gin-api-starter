// @Title constant.go
// @Description
// @Author Hunter 2024/9/4 10:47

package config

// CommonSplicePrefix Common prefixes for splicing, such as token, redis key, etc., to distinguish different projects
const CommonSplicePrefix = "go-gin-api-starter"

// NodeEnv current running environment
const (
	Development = "development"
	Production  = "production"
	Test        = "test"
)
