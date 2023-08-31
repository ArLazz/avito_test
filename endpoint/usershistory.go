package endpoint

import (
	"context"
	"slugservice/service"
	"github.com/go-kit/kit/endpoint"
	"slugservice/transport"
)

func MakeUsersHistoryEndpoint(srv service.Service) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req, ok := request.(transport.UsersHistoryRequest)
		if !ok {
			return transport.UsersHistoryResponse{Err: errUnexpected}, errUnexpected
		}

		res, err := srv.UsersHistory(ctx, req.Year, req.Month)
		if err != nil {
			return transport.UsersHistoryResponse{Link: res, Err: err}, nil
		}
		
		return transport.UsersHistoryResponse{Link: res}, nil
	}
}

func (e Endpoints) UsersHistory(ctx context.Context, year string, month string) (string, error) {
	req := transport.UsersHistoryRequest{Year: year, Month: month}

	resp, err := e.UsersHistoryEndpoint(ctx, req)
	if err != nil {
		return "error", err
	}

	usersHistoryResp, ok := resp.(transport.UsersHistoryResponse)
	if !ok {
		return "error", errUnexpected
	}

	if usersHistoryResp.Err != nil {
		return "error", usersHistoryResp.Err
	}

	return usersHistoryResp.Link, nil
}