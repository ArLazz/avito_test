package endpoint

import (
	"context"
	"slugservice/service"
	"github.com/go-kit/kit/endpoint"
	"slugservice/transport"
)


func MakeUserSegmentsEndpoint(srv service.Service) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req, ok := request.(transport.UserSegmentsRequest)
		if !ok {
			return transport.UserSegmentsResponse{Err: errUnexpected}, errUnexpected
		}

		res, err := srv.UserSegments(ctx, req.UserId)
		if err != nil {
			return transport.UserSegmentsResponse{Status: res, Err: err}, nil
		}
		
		return transport.UserSegmentsResponse{Status: res}, nil
	}
}

func (e Endpoints) UserSegments(ctx context.Context, userId string) (string, error) {
	req := transport.UserSegmentsRequest{UserId: userId}

	resp, err := e.UserSegmentsEndpoint(ctx, req)
	if err != nil {
		return "error", err
	}

	userSegmentsResp, ok := resp.(transport.UserSegmentsResponse)
	if !ok {
		return "error", errUnexpected
	}

	if userSegmentsResp.Err != nil {
		return "error", userSegmentsResp.Err
	}

	return userSegmentsResp.Status, nil
}