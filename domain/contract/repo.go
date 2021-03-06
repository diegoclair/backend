package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/safra-team-35/backend/domain/entity"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Ping() PingRepo
	QRCode() QRCodeRepo
	Company() CompanyRepo
	User() UserRepo
}

// PingRepo defines the data set for ping
type PingRepo interface{}

// QRCodeRepo defines the data set for qrcode
type QRCodeRepo interface {
	CreateQRCode(qrcode entity.QRCode) resterrors.RestErr
	GetByHash(hash string) (qrcode entity.QRCode, err resterrors.RestErr)
}

// CompanyRepo defines the data set for qrcode
type CompanyRepo interface {
	GetCompanyIDByUUID(uuid string) (int64, resterrors.RestErr)
}

// UserRepo defines the data set for user
type UserRepo interface {
	GetUserAddress(userID int64) (address []entity.Address, err resterrors.RestErr)
	CreateOrder(order entity.Order) (err resterrors.RestErr)
	GetOrderSummary(userID int64) (summary []entity.OrderSummary, restErr resterrors.RestErr)
}
