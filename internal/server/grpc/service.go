package grpc

import (
	"context"
	"errors"
	"strings"

	"github.com/slem7451/anti_bruteforce/internal/entity/request"
	"github.com/slem7451/anti_bruteforce/internal/server"
	"github.com/slem7451/anti_bruteforce/internal/server/grpc/pb"
)

var errSubnetRequired = errors.New("subnet is required")

type service struct {
	pb.UnimplementedAuthServer
	app server.App
}

func requestToStruct(req *pb.Credits, isAuthMethod bool) (request.Credits, error) {
	if err := validateCredits(req, isAuthMethod); err != nil {
		return request.Credits{}, err
	}

	return request.Credits{
		IP:       req.GetIp(),
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	}, nil
}

func validateCredits(req *pb.Credits, isAuthMethod bool) error {
	msg := make([]string, 0, 3)

	if req.GetIp() == "" {
		msg = append(msg, "IP is required")
	}

	if req.GetLogin() == "" {
		msg = append(msg, "Login is required")
	}

	if isAuthMethod && req.GetPassword() == "" {
		msg = append(msg, "Password is required for this method")
	}

	if len(msg) != 0 {
		return errors.New(strings.Join(msg, "\n"))
	}

	return nil
}

func errResponse(err error) *pb.Response {
	errStr := err.Error()

	return &pb.Response{
		Ok:  false,
		Msg: &errStr,
	}
}

func (s *service) Auth(ctx context.Context, req *pb.Credits) (*pb.Response, error) {
	credits, err := requestToStruct(req, true)
	if err != nil {
		return errResponse(err), nil
	}

	res, err := s.app.ValidateAuth(ctx, credits)
	if err != nil {
		return errResponse(err), nil
	}

	return &pb.Response{
		Ok: res,
	}, nil
}

func (s *service) Reset(ctx context.Context, req *pb.Credits) (*pb.Response, error) {
	credits, err := requestToStruct(req, false)
	if err != nil {
		return errResponse(err), nil
	}

	err = s.app.RemoveLimit(ctx, credits)
	if err != nil {
		return errResponse(err), nil
	}

	return &pb.Response{
		Ok: true,
	}, nil
}

func (s *service) AddToBlacklist(ctx context.Context, req *pb.Subnet) (*pb.Response, error) {
	if req.GetSubnet() == "" {
		return errResponse(errSubnetRequired), nil
	}

	err := s.app.AddToBlacklist(ctx, req.GetSubnet())
	if err != nil {
		return errResponse(err), nil
	}

	return &pb.Response{
		Ok: true,
	}, nil
}

func (s *service) DeleteFromBlacklist(ctx context.Context, req *pb.Subnet) (*pb.Response, error) {
	if req.GetSubnet() == "" {
		return errResponse(errSubnetRequired), nil
	}

	err := s.app.DeleteFromBlacklist(ctx, req.GetSubnet())
	if err != nil {
		return errResponse(err), nil
	}

	return &pb.Response{
		Ok: true,
	}, nil
}

func (s *service) AddToWhitelist(ctx context.Context, req *pb.Subnet) (*pb.Response, error) {
	if req.GetSubnet() == "" {
		return errResponse(errSubnetRequired), nil
	}

	err := s.app.AddToWhitelist(ctx, req.GetSubnet())
	if err != nil {
		return errResponse(err), nil
	}

	return &pb.Response{
		Ok: true,
	}, nil
}

func (s *service) DeleteFromWhitelist(ctx context.Context, req *pb.Subnet) (*pb.Response, error) {
	if req.GetSubnet() == "" {
		return errResponse(errSubnetRequired), nil
	}

	err := s.app.DeleteFromWhitelist(ctx, req.GetSubnet())
	if err != nil {
		return errResponse(err), nil
	}

	return &pb.Response{
		Ok: true,
	}, nil
}
