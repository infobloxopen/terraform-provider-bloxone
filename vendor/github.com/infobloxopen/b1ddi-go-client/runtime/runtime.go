// Package runtime contains BloxOne DDI API helper
// runtime abstractions for internal client use.
package runtime

import (
	"fmt"
	"github.com/go-openapi/runtime"
	"io"
	"io/ioutil"
	"path"
	"strings"
)

// TrimIDPrefix removes app ID and resource type prefixes from the ID property.
//
// If no prefix is found, TrimIDPrefix will return the original id value.
//
// More about BloxOne DDI resource IDs:
// https://github.com/infobloxopen/atlas-app-toolkit/tree/v1.1.2/rpc/resource#resource
func TrimIDPrefix(pathPattern string, id string) string {
	if !path.IsAbs(id) {
		id = "/" + id
	}
	prefix := pathPattern
	for !strings.HasPrefix(id, prefix) {
		prefix = prefix[0 : len(prefix)-1]
		if len(prefix) == 0 {
			return strings.TrimPrefix(id, "/")
		}
	}
	if strings.HasPrefix(pathPattern, prefix+"{id}") {
		return strings.TrimPrefix(id, prefix)
	} else {
		return strings.TrimPrefix(id, "/")
	}

}

// NewAPIHTTPError creates a new API HTTP error.
func NewAPIHTTPError(opName string, payload io.Reader, code int) *APIHTTPError {
	body, err := ioutil.ReadAll(payload)
	if err != nil {
		body = []byte("Failed to read response")
	}
	return &APIHTTPError{
		runtime.APIError{
			OperationName: opName,
			Response:      string(body),
			Code:          code,
		},
	}
}

// APIHTTPError wraps runtime.APIError, and modifies response processing logic.
type APIHTTPError struct {
	runtime.APIError
}

// Error method prints APIHTTPError error message.
func (a *APIHTTPError) Error() string {
	return fmt.Sprintf("%s (status %d): \n%s", a.OperationName, a.Code, a.Response)
}
