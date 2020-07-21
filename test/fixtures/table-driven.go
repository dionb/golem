package fixtures

import (
	"bytes"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gocraft/web"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	Name           string
	Method         string
	URL            string
	User           string //Authorization header value
	ReqBody        string
	ExpectedBody   string
	ExpectedStatus int
}

func RunTest(t *testing.T, test Case, subject *web.Router) {
	t.Run(test.Name, func(t *testing.T) {
		rw := &MockRW{}
		reqBody := bytes.NewBufferString(test.ReqBody)
		req, err := http.NewRequest(test.Method, test.URL, reqBody)
		if err != nil {
			log.Println(err.Error())
		}
		req.Header["Authorization"] = []string{test.User}
		assert.NoError(t, err)
		subject.ServeHTTP(rw, req)
		if test.ExpectedStatus == 0 {
			test.ExpectedStatus = 200
		}
		assert.Equal(t, test.ExpectedStatus, rw.status)
		assert.Regexp(t, test.ExpectedBody, rw.String())
		time.Sleep(time.Millisecond) // solves a race condition by making sure goroutines in the application get scheduled so that expected changes are applied before the next test that confirms they worked
	})
}
