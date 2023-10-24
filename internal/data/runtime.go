package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// Implement the `Marshaler` interface
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
