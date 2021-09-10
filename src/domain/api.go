package domain

import (
	"context"

	"github.com/plagioriginal/api-gateway/protos/protos"
)

type LoginRequest struct {
	Username string `validate:"required,min=3"`
	Password string `validate:"required,min=8"`
}

type AddUserRequest struct {
	Username string `validate:"required,min=3"`
	Password string `validate:"required,min=3"`
	Role     string `validate:"required"`
	JwtToken string `validate:"required"`
}

type ListProjectsRequest struct {
	JwtToken string `validate:"required"`
	PageNum  uint
}

// Services to be used by the HTTP Handler.
// Should have all the gRPC clients into it's base, so it could do all the requests.
//
// @todo: define the `interfaces{}` with the Objects from the gRPC clients.
type APIService interface {
	Login(ctx context.Context, loginRequest LoginRequest) (*protos.TokenResponse, error)
	RefreshJWT(ctx context.Context, refreshToken string) (*protos.TokenResponse, error)
	AddUser(ctx context.Context, userRequest AddUserRequest) (*protos.UserResponse, error)
	ListProjects(ctx context.Context, listProjectReq ListProjectsRequest) (interface{}, error)
}
