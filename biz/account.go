package biz

import (
	"context"
	"github.com/Gentleelephant/account-srv/config"
	"github.com/Gentleelephant/account-srv/handler"
	"github.com/Gentleelephant/proto-center/pb/model"
)

type AccountServer struct {
	*pb.UnimplementedAccountServiceServer
}

func (a *AccountServer) GetAccountList(ctx context.Context, req *pb.AccountPagingRequest) (*pb.AccountListRes, error) {
	reply, err := handler.GetAccountList(ctx, config.DB, req.GetPageNo(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) GetAccountByMobile(ctx context.Context, req *pb.MobileRequest) (*pb.AccountRes, error) {
	reply, err := handler.GetAccountByMobile(ctx, config.DB, req.Mobile)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) GetAccountByID(ctx context.Context, req *pb.IDRequest) (*pb.AccountRes, error) {
	reply, err := handler.GetAccountByID(ctx, config.DB, req.Id)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) AddAccount(ctx context.Context, req *pb.AddAccountRequest) (*pb.AccountRes, error) {
	reply, err := handler.AddAccount(ctx, config.DB, req.Mobile, req.Password, req.Nickname, req.Gender)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.AccountRes, error) {
	reply, err := handler.UpdateAccount(ctx, config.DB, req.Id, req.Mobile, req.Password, req.Nickname, req.Gender)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (a *AccountServer) CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordRes, error) {
	reply, err := handler.CheckPassword(ctx, config.DB, req.AccountId, req.Password)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
