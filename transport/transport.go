package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

// Сначала описываем "модели" запросов и ответов нашего сервиса

type CreateSegmentRequest struct {
	Slug string `json:"slug_name"`
}

type CreateSegmentResponse struct {
	Status string `json:"status"`
	Err    error  `json:"err,omitempty"`
}

type UserSegmentsRequest struct {
	UserId string `json:"user_id"`
}

type UserSegmentsResponse struct {
	Status string `json:"status"`
	Err    error  `json:"err,omitempty"`
}

type DeleteSegmentRequest struct {
	Slug string `json:"slug_name"`
}

type DeleteSegmentResponse struct {
	Status string `json:"status"`
	Err    error  `json:"err,omitempty"`
}

type AddUserRequest struct {
	AddSlug    string `json:"add_slug"`
	DeleteSlug string `json:"delete_slug"`
	UserId     string `json:"user_id"`
}

type AddUserResponse struct {
	Status string `json:"status"`
	Err    error  `json:"err,omitempty"`
}

type UsersHistoryRequest struct {
	Year    string `json:"year"`
	Month string `json:"month"`
}

type UsersHistoryResponse struct {
	Link string `json:"link"`
	Err    error  `json:"err,omitempty"`
}

// Затем описываем "декодеры" для входящих запросов

func DecodeCreateSegmentRequest(
	_ context.Context,
	r *http.Request,
) (interface{}, error) {
	var req CreateSegmentRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}
func DecodeDeleteSegmentRequest(
	_ context.Context,
	r *http.Request,
) (interface{}, error) {
	var req DeleteSegmentRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeAddUserRequest(
	_ context.Context,
	r *http.Request,
) (interface{}, error) {
	var req AddUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeUserSegmentsRequest(
	_ context.Context,
	r *http.Request,
) (interface{}, error) {
	var req UserSegmentsRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeUsersHistoryRequest(
	_ context.Context,
	r *http.Request,
) (interface{}, error) {
	var req UsersHistoryRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func EncodeResponse(
	_ context.Context,
	w http.ResponseWriter,
	response interface{},
) error {
	return json.NewEncoder(w).Encode(response)
}
