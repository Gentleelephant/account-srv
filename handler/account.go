package handler

import (
	"context"
	"errors"
	"github.com/Gentleelephant/account-srv/custom_error"
	"github.com/Gentleelephant/account-srv/internal"
	"github.com/Gentleelephant/account-srv/model"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
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

func GetAccountList(ctx context.Context, db *gorm.DB, pageNo, pageSize uint32) (*pb.AccountListRes, error) {
	zap.L().Info("GetAccountList", zap.Uint32("pageNo", pageNo), zap.Uint32("pageSize", pageSize))
	var accountList []*model.Account
	result := db.Model(&model.Account{}).Scopes(Paginate(int(pageNo), int(pageSize))).Find(&accountList)
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

func GetAccountByMobile(ctx context.Context, db *gorm.DB, mobile string) (*pb.AccountRes, error) {
	account := &model.Account{}
	res := db.Model(&model.Account{}).Where("mobile = ?", mobile).Find(account)
	if res.Error != nil {
		return nil, res.Error
	}
	return Model2Pb(account), nil
}

func GetAccountByID(ctx context.Context, db *gorm.DB, id int32) (*pb.AccountRes, error) {
	account := &model.Account{}
	res := db.Model(&model.Account{}).Where("id = ?", id).Find(account)
	if res.Error != nil {
		return nil, res.Error
	}
	return Model2Pb(account), nil
}

func AddAccount(ctx context.Context, db *gorm.DB, mobile, password, nickname, gender string) (*pb.AccountRes, error) {
	account := &model.Account{}

	find := db.Model(&model.Account{}).Where("mobile = ?", mobile).Find(account)
	if find.RowsAffected == 1 {
		return nil, errors.New(custom_error.AccountExists)
	}
	slat, encodeSlat := internal.PasswordEncode(password)
	account.Salt = slat
	account.Password = encodeSlat
	account.Mobile = mobile
	account.Nickname = nickname
	account.Gender = gender
	result := db.Model(&model.Account{}).Create(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return Model2Pb(account), nil
}

func UpdateAccount(ctx context.Context, db *gorm.DB, id uint32, mobile, password, nickname, gender string) (*pb.AccountRes, error) {
	account := &model.Account{}
	find := db.Model(&model.Account{}).Where("id = ?", id).Find(account)
	if find.RowsAffected != 1 {
		return nil, errors.New(custom_error.AccountNotFound)
	}
	slat, encodePassword := internal.PasswordEncode(password)
	account.Salt = slat
	account.Password = encodePassword
	account.Mobile = mobile
	account.Nickname = nickname
	account.Gender = gender
	result := db.Model(&model.Account{}).Where("id = ?", id).Updates(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return Model2Pb(account), nil
}

func CheckPassword(ctx context.Context, db *gorm.DB, id, password string) (*pb.CheckPasswordRes, error) {
	account := &model.Account{}
	res := db.Model(&model.Account{}).Where("id = ?", id).Find(account)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, errors.New(custom_error.AccountNotFound)
	}
	if account.Salt == "" {
		return nil, errors.New(custom_error.AccountSlatIsEmpty)
	}
	verify := internal.PasswordVerify(password, account.Salt, account.Password)
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
