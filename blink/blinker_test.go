package blink

import (
	"testing"

	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"
)

type MockResty struct {
	Req    *resty.Request
	Params map[string]string
}

func (m *MockResty) SetQueryParams(params map[string]string) *resty.Request {
	for p, v := range params {
		m.Req.QueryParam.Add(p, v)
	}
	return m.Req
}

func (m *MockResty) Get(url string) (*resty.Response, error) {
	return nil, nil
}

func (m *MockResty) SetBasicAuth(username, password string) *resty.Request {
	m.Req.UserInfo = &resty.User{Username: username, Password: password}
	return m.Req
}

func TestBlinkRed(t *testing.T) {
	testObj := MockResty{Req: resty.New().R()}
	blink := Blinker{Resty: &testObj, BlinkURL: "http://localhost"}

	blink.Blink("ff0000")

	assert.Equal(t, "http://localhost/blink1/fadeToRGB?rgb=%23ff0000", testObj.Req.URL, "url should be correct")
}

func TestBlinkGreen(t *testing.T) {
	testObj := MockResty{Req: resty.New().R()}
	blink := Blinker{Resty: &testObj, BlinkURL: "http://localhost"}
	blink.Blink("00ff00")

	assert.Equal(t, "http://localhost/blink1/fadeToRGB?rgb=%2300ff00", testObj.Req.URL, "url should be correct")
}
