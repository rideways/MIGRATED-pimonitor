package rest

import "github.com/go-resty/resty"

// IResty resty rest
type IResty interface {
	SetBasicAuth(username, password string) *resty.Request
	SetQueryParams(map[string]string) *resty.Request
	Get(string) (*resty.Response, error)
}
