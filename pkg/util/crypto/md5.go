// @Title md5.go
// @Description
// @Author Hunter 2024/9/3 17:34

package crypto

import (
	"crypto/md5"
	"fmt"
)

// Md5
// @Description: MD5
// @param originMessage origin message
// @return result encrypted message
func Md5(originMessage string) (result string) {
	data := []byte(originMessage)
	result = fmt.Sprintf("%x", md5.Sum(data))
	return
}
