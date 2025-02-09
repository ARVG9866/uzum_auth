package login_v1

import (
	"context"

	"github.com/Shemistan/uzum_auth/internal/models"
	"github.com/Shemistan/uzum_auth/internal/service/login_v1"
	pb "github.com/Shemistan/uzum_auth/pkg/login_v1"
)

type Login struct {
	pb.UnimplementedLoginV1Server
	Service login_v1.ILoginService
}

func (s *Login) Login(ctx context.Context, req *pb.Login_Request) (*pb.Login_Response, error) {
	res, err := s.Service.Login(ctx, &models.AuthUser{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.Login_Response{
		RefreshToken: res.Refresh,
		AccessToken:  res.Access,
	}, nil
}

func (s *Login) Check(ctx context.Context, _ *pb.Check_Request) (*pb.Check_Response, error) {
	user_id, err := s.Service.Check(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.Check_Response{
		UserId: user_id,
	}, nil
}
