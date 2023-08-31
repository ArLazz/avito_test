package pkg

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"slugservice/endpoint"
	"slugservice/transport"
)

func NewHTTPServer(endpoints endpoint.Endpoints) *http.ServeMux {
	router := http.NewServeMux()

	// создадим простой middleware
	// он будет устанавливать для всех запросов,
	// зарегистрированных через него, тип ответа "application/json"
	handle := func(pattern string, handler http.Handler) {
		router.Handle(pattern, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Type", "application/json; charset=utf-8")
			handler.ServeHTTP(writer, request)
		}))
	}
	handle(
		"/createsegment",
		httptransport.NewServer(
			endpoints.CreateSegmentEndpoint,
			transport.DecodeCreateSegmentRequest,
			transport.EncodeResponse,
		),
	)
	handle(
		"/deletesegment",
		httptransport.NewServer(
			endpoints.DeleteSegmentEndpoint,
			transport.DecodeDeleteSegmentRequest,
			transport.EncodeResponse,
		),
	)
	handle(
		"/adduser",
		httptransport.NewServer(
			endpoints.AddUserEndpoint,
			transport.DecodeAddUserRequest,
			transport.EncodeResponse,
		),
	)
	handle(
		"/usersegments",
		httptransport.NewServer(
			endpoints.UserSegmentsEndpoint,
			transport.DecodeUserSegmentsRequest,
			transport.EncodeResponse,
		),
	)
	handle(
		"/usershistory",
		httptransport.NewServer(
			endpoints.UsersHistoryEndpoint,
			transport.DecodeUsersHistoryRequest,
			transport.EncodeResponse,
		),
	)
	return router

}
