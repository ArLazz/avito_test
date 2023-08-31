package endpoint

import (
	"context"
	"slugservice/service"
	"github.com/go-kit/kit/endpoint"
	"slugservice/transport"
)

func MakeCreateSegmentEndpoint(srv service.Service) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req, ok := request.(transport.CreateSegmentRequest)
		if !ok {
			return transport.CreateSegmentResponse{Err: errUnexpected}, errUnexpected
		}

		res, err := srv.CreateSegment(ctx, req.Slug)
		if err != nil {
			return transport.CreateSegmentResponse{Status: res, Err: err}, nil
		}

		return transport.CreateSegmentResponse{Status: res}, nil
	}
}

func (e Endpoints) CreateSegment(ctx context.Context, slug string) (string, error) {
	req := transport.CreateSegmentRequest{Slug: slug}

	resp, err := e.CreateSegmentEndpoint(ctx, req)
	if err != nil {
		return "error", err
	}

	createSegmentResp, ok := resp.(transport.CreateSegmentResponse)
	if !ok {
		return "error", errUnexpected
	}

	if createSegmentResp.Err != nil {
		return "error", createSegmentResp.Err
	}

	return createSegmentResp.Status, nil
}