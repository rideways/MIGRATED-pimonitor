package sensu

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `[{"check": {"handle":true, "status": 1, "name": "check1"}}]`)
}

type MockBlinker struct{}

func (m *MockBlinker) blink(colour string) {
}

func TestGetSingleTest(t *testing.T) {
	handler := &MyHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	checker := Checker{Resty: resty.New().R(), SensuAPIURL: server.URL, Username: "dave", Password: "xxxx", UnknownColour: "#ffffff"}
	result, err := checker.GetChecks()

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, 1, len(*result), "should contain 1 result")

	for _, v := range *result {
		assert.Equal(t, "check1", v.Check.Name)
	}
}
