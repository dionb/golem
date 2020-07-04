package fixtures

import (
	"bytes"
	"net/http"
)

type MockRW struct {
	bytes.Buffer
	status int
}

func (rw *MockRW) Header() http.Header {
	return http.Header{}
}

func (rw *MockRW) WriteHeader(status int) {
	rw.status = status
}
