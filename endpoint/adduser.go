package endpoint

import (
	"context"
	"slugservice/service"
	"github.com/go-kit/kit/endpoint"
	"slugservice/transport"
)


func MakeAdduserEndpoint(srv service.Service) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req, ok := request.(transport.AddUserRequest)
		if !ok {
			return transport.AddUserResponse{Err: errUnexpected}, errUnexpected
		}

		res, err := srv.AddUser(ctx, req.AddSlug, req.DeleteSlug, req.UserId)
		if err != nil {
			return transport.AddUserResponse{Status: res, Err: err}, nil
		}
		return transport.AddUserResponse{Status: res}, nil
	}
}




func (e Endpoints) AddUser(ctx context.Context, addSlug string, deleteSlug string, userId string) (string, error) {
	req := transport.AddUserRequest{AddSlug: addSlug, DeleteSlug: deleteSlug, UserId: userId}

	resp, err := e.AddUserEndpoint(ctx, req)
	if err != nil {
		return "error", err
	}

	addUserResp, ok := resp.(transport.AddUserResponse)
	if !ok {
		return "error", errUnexpected
	}

	if addUserResp.Err != nil {
		return "error", addUserResp.Err
	}

	return addUserResp.Status, nil
}