package endpoint

import (
	"context"
	"slugservice/service"
	"github.com/go-kit/kit/endpoint"
	"slugservice/transport"
)

func MakeDeleteSegmentEndpoint(srv service.Service) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req, ok := request.(transport.DeleteSegmentRequest)
		if !ok {
			return transport.DeleteSegmentResponse{Err: errUnexpected}, errUnexpected
		}

		res, err := srv.DeleteSegment(ctx, req.Slug)
		if err != nil {
			return transport.DeleteSegmentResponse{Status: res, Err: err}, nil
		}

		return transport.DeleteSegmentResponse{Status: res}, nil
	}
}


func (e Endpoints) DeleteSegment(ctx context.Context, slug string) (string, error) {
	req := transport.DeleteSegmentRequest{Slug: slug}

	resp, err := e.DeleteSegmentEndpoint(ctx, req)
	if err != nil {
		return "error", err
	}

	deleteSegmentResp, ok := resp.(transport.DeleteSegmentResponse)
	if !ok {
		return "error", errUnexpected
	}

	if deleteSegmentResp.Err != nil {
		return "error", deleteSegmentResp.Err
	}

	return deleteSegmentResp.Status, nil
}
