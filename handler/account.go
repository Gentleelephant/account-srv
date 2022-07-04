package handler

import (
	"account-srv/biz"
	"account-srv/custom_error"
	"account-srv/internal"
	"account-srv/model"
	"context"
	"errors"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func Paginate(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo <= 0 {
			pageNo = 1
		}
		switch {
		case pageSize <= 0:
			pageSize = 1
		case pageSize > 100:
			pageSize = 100
		}
		return db.Offset((pageNo - 1) * pageSize).Limit(pageSize)
	}
}

func GetAccountList(ctx context.Context, req *pb.AccountPagingRequest) (*pb.AccountListRes, error) {
	var accountList []*model.Account
	result := internal.DB.Model(&model.Account{}).Scopes(Paginate(int(req.PageNo), int(req.PageSize))).Find(&accountList)
	if result.Error != nil {
		return nil, result.Error
	}
	accountListRes := &pb.AccountListRes{}
	accountListRes.Total = int32(result.RowsAffected)
	for _, v := range accountList {
		accountListRes.AccountList = append(accountListRes.AccountList, Model2Pb(v))
	}
	return accountListRes, nil
}

func GetAccountByMobile(ctx context.Context, req *pb.MobileRequest) (*pb.AccountRes, error) {
	account := &model.Account{}
	res := internal.DB.Model(&model.Account{}).Where("mobile = ?", req.Mobile).Find(account)
	if res.Error != nil {
		return nil, res.Error
	}
	return Model2Pb(account), nil
}

func GetAccountByID(ctx context.Context, req *pb.IDRequest) (*pb.AccountRes, error) {
	account := &model.Account{}
	res := internal.DB.Model(&model.Account{}).Where("id = ?", req.Id).Find(account)
	if res.Error != nil {
		return nil, res.Error
	}
	return Model2Pb(account), nil
}

func AddAccount(ctx context.Context, req *pb.AddAccountRequest) (*pb.AccountRes, error) {
	account := &model.Account{}

	find := internal.DB.Model(&model.Account{}).Where("mobile = ?", req.Mobile).Find(account)
	if find.RowsAffected == 1 {
		return nil, errors.New(custom_error.AccountExists)
	}
	err := copier.Copy(&account, &req)
	if err != nil {
		return nil, err
	}
	slat, encodeSlat := biz.PasswordEncode(req.Password)
	account.Salt = slat
	account.Password = encodeSlat
	result := internal.DB.Model(&model.Account{}).Create(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return Model2Pb(account), nil
}

func UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.AccountRes, error) {
	account := &model.Account{}
	err := copier.Copy(&account, &req)
	if err != nil {
		return nil, err
	}
	find := internal.DB.Model(&model.Account{}).Where("id = ?", req.Id).Find(account)
	if find.RowsAffected != 1 {
		return nil, errors.New(custom_error.AccountNotFound)
	}
	slat, encodePassword := biz.PasswordEncode(req.Password)
	account.Salt = slat
	account.Password = encodePassword
	result := internal.DB.Model(&model.Account{}).Where("id = ?", req.Id).Updates(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return Model2Pb(account), nil
}

func CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordRes, error) {
	account := &model.Account{}
	res := internal.DB.Model(&model.Account{}).Where("id = ?", req.AccountId).Find(account)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, errors.New(custom_error.AccountNotFound)
	}
	if account.Salt == "" {
		return nil, errors.New(custom_error.AccountSlatIsEmpty)
	}
	verify := biz.PasswordVerify(req.Password, account.Salt, account.Password)
	return &pb.CheckPasswordRes{Result: verify}, nil
}

func Model2Pb(account *model.Account) *pb.AccountRes {
	res := &pb.AccountRes{}
	err := copier.Copy(&res, &account)
	if err != nil {
		return nil
	}
	return res
}

func Pb2Model(accountRes *pb.AccountRes) *model.Account {
	res := &model.Account{}
	err := copier.Copy(&res, &accountRes)
	if err != nil {
		return nil
	}
	return res
}
