// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"suxenia-finance/pkg/common/infrastructure/cache"
	"suxenia-finance/pkg/common/infrastructure/logs"
	"suxenia-finance/pkg/kyc/application"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/drivers"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/repos"
	"suxenia-finance/pkg/kyc/infrastructure/routes"
	application2 "suxenia-finance/pkg/wallet/application"
	drivers2 "suxenia-finance/pkg/wallet/infrastructure/persistence/drivers"
	repos2 "suxenia-finance/pkg/wallet/infrastructure/persistence/repos"
	routes2 "suxenia-finance/pkg/wallet/infrastructure/routes"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func InitalizeApplication() (*Application, error) {
	db := NewDBInstance()
	sugaredLogger := logs.NewLogger()
	bankKycDriver, err := drivers.NewBankycDriver(db, sugaredLogger)
	if err != nil {
		return nil, err
	}
	iBankingKycRepo, err := repos.NewBankycRepo(bankKycDriver)
	if err != nil {
		return nil, err
	}
	bankingKYCApplication, err := application.NewBankingKycApplication(iBankingKycRepo, sugaredLogger)
	if err != nil {
		return nil, err
	}
	kycRoutes, err := routes.NewKycRoute(bankingKYCApplication, sugaredLogger)
	if err != nil {
		return nil, err
	}
	paymentDriver, err := drivers2.NewPaymentDriver(db, sugaredLogger)
	if err != nil {
		return nil, err
	}
	cacheCache := cache.NewRedisCache()
	walletDriver, err := drivers2.NewWalletDriver(db, sugaredLogger)
	if err != nil {
		return nil, err
	}
	walletRepo, err := repos2.NewWalletRepo(walletDriver)
	if err != nil {
		return nil, err
	}
	walletTransactionDriver, err := drivers2.NewWalletTransactionDriver(db, sugaredLogger)
	if err != nil {
		return nil, err
	}
	paymentApplication, err := application2.NewPaymentApplication(paymentDriver, cacheCache, walletRepo, walletTransactionDriver, sugaredLogger)
	if err != nil {
		return nil, err
	}
	paymentApi, err := routes2.NewPaymentApi(paymentApplication, sugaredLogger)
	if err != nil {
		return nil, err
	}
	mainApplication, err := NewApplication(kycRoutes, paymentApi)
	if err != nil {
		return nil, err
	}
	return mainApplication, nil
}
