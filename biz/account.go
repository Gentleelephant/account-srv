package biz

import (
	"account-srv/handler"
	"context"
	"github.com/Gentleelephant/proto-center/pb/model"
)

type AccountServer struct {
	*pb.UnimplementedAccountServiceServer
}

func (a *AccountServer) GetAccountList(ctx context.Context, req *pb.AccountPagingRequest) (*pb.AccountListRes, error) {
	reply, err := handler.GetAccountList(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) GetAccountByMobile(ctx context.Context, req *pb.MobileRequest) (*pb.AccountRes, error) {
	reply, err := handler.GetAccountByMobile(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) GetAccountByID(ctx context.Context, req *pb.IDRequest) (*pb.AccountRes, error) {
	reply, err := handler.GetAccountByID(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) AddAccount(ctx context.Context, req *pb.AddAccountRequest) (*pb.AccountRes, error) {
	reply, err := handler.AddAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.AccountRes, error) {
	reply, err := handler.UpdateAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordRes, error) {
	reply, err := handler.CheckPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
