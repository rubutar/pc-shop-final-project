package repository

import (
	"context"
	"pc-shop-final-project/domain/entity/settlement"
)

type InterfaceSettlement interface {
	CreateSettle(ctx context.Context, idUser int, idCustomer int, idCoupon int, settle *settlement.Settlement) error
	ReadSettle(ctx context.Context) ([]*settlement.Settlement, error)
	UpdateSettle(ctx context.Context, code string) error
	DeleteSettle(ctx context.Context, code string) error
}

type InterfaceSettlementItem interface {
	CreatePurchase(ctx context.Context, purchase []*settlement.SettlePurchase) error
	ReadPurchase(ctx context.Context) ([]*settlement.SettlePurchase, error)
	DeletePurchase(ctx context.Context, codeItem string) error
}
