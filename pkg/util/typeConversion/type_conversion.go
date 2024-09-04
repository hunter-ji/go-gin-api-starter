// @Title type_conversion.go
// @Description
// @Author Hunter 2024/9/4 20:29

package typeConversion

import (
	"bytes"
	"encoding/json"
	"io"
)

func StructToReader(s interface{}) (io.Reader, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func StringToStruct(data string, s interface{}) error {
	return json.Unmarshal([]byte(data), s)
}
