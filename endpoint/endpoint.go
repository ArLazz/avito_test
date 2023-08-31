package endpoint

import (
	"errors"
	"github.com/go-kit/kit/endpoint"
)

var (
	errUnexpected = errors.New("unexpected error")
)


type Endpoints struct {
	CreateSegmentEndpoint endpoint.Endpoint
	DeleteSegmentEndpoint endpoint.Endpoint
	AddUserEndpoint       endpoint.Endpoint
	UserSegmentsEndpoint  endpoint.Endpoint
	UsersHistoryEndpoint endpoint.Endpoint
}




