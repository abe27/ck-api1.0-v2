package configs

import (
	"github.com/abe27/api/models"
	"gorm.io/gorm"
)

var (
	Store           *gorm.DB
	API_TRIGGER_URL string
)

func SetDB() {
	if !Store.Migrator().HasTable(&models.User{}) {
		Store.AutoMigrate(&models.User{})
	}
	if !Store.Migrator().HasTable(&models.JwtToken{}) {
		Store.AutoMigrate(&models.JwtToken{})
	}
	if !Store.Migrator().HasTable(&models.Administrator{}) {
		Store.AutoMigrate(&models.Administrator{})
	}
	if !Store.Migrator().HasTable(&models.Area{}) {
		Store.AutoMigrate(&models.Area{})
	}
	if !Store.Migrator().HasTable(&models.Whs{}) {
		Store.AutoMigrate(&models.Whs{})
	}
	if !Store.Migrator().HasTable(&models.Factory{}) {
		Store.AutoMigrate(&models.Factory{})
	}
	if !Store.Migrator().HasTable(&models.ReceiveType{}) {
		Store.AutoMigrate(&models.ReceiveType{})
	}
	if !Store.Migrator().HasTable(&models.Unit{}) {
		Store.AutoMigrate(&models.Unit{})
	}
	if !Store.Migrator().HasTable(&models.PartType{}) {
		Store.AutoMigrate(&models.PartType{})
	}
	if !Store.Migrator().HasTable(&models.FileType{}) {
		Store.AutoMigrate(&models.FileType{})
	}
	if !Store.Migrator().HasTable(&models.Position{}) {
		Store.AutoMigrate(&models.Position{})
	}
	if !Store.Migrator().HasTable(&models.Department{}) {
		Store.AutoMigrate(&models.Department{})
	}
	if !Store.Migrator().HasTable(&models.PrefixName{}) {
		Store.AutoMigrate(&models.PrefixName{})
	}
	if !Store.Migrator().HasTable(&models.Profile{}) {
		Store.AutoMigrate(&models.Profile{})
	}
	if !Store.Migrator().HasTable(&models.Mailbox{}) {
		Store.AutoMigrate(&models.Mailbox{})
	}
	if !Store.Migrator().HasTable(&models.FileEdi{}) {
		Store.AutoMigrate(&models.FileEdi{})
	}
	if !Store.Migrator().HasTable(&models.CartonHistory{}) {
		Store.AutoMigrate(&models.CartonHistory{})
	}
	if !Store.Migrator().HasTable(&models.SyncLogger{}) {
		Store.AutoMigrate(&models.SyncLogger{})
	}
	if !Store.Migrator().HasTable(&models.Part{}) {
		Store.AutoMigrate(&models.Part{})
	}
	if !Store.Migrator().HasTable(&models.Ledger{}) {
		Store.AutoMigrate(&models.Ledger{})
	}
	if !Store.Migrator().HasTable(&models.Receive{}) {
		Store.AutoMigrate(&models.Receive{})
	}
	if !Store.Migrator().HasTable(&models.ReceiveDetail{}) {
		Store.AutoMigrate(&models.ReceiveDetail{})
	}
	if !Store.Migrator().HasTable(&models.Pc{}) {
		Store.AutoMigrate(&models.Pc{})
	}
	if !Store.Migrator().HasTable(&models.Commercial{}) {
		Store.AutoMigrate(&models.Commercial{})
	}
	if !Store.Migrator().HasTable(&models.SampleFlg{}) {
		Store.AutoMigrate(&models.SampleFlg{})
	}
	if !Store.Migrator().HasTable(&models.ReviseOrder{}) {
		Store.AutoMigrate(&models.ReviseOrder{})
	}
	if !Store.Migrator().HasTable(&models.Shipment{}) {
		Store.AutoMigrate(&models.Shipment{})
	}
	if !Store.Migrator().HasTable(&models.OrderZone{}) {
		Store.AutoMigrate(&models.OrderZone{})
	}
	if !Store.Migrator().HasTable(&models.OrderType{}) {
		Store.AutoMigrate(&models.OrderType{})
	}
	if !Store.Migrator().HasTable(&models.Affcode{}) {
		Store.AutoMigrate(&models.Affcode{})
	}
	if !Store.Migrator().HasTable(&models.Customer{}) {
		Store.AutoMigrate(&models.Customer{})
	}
	if !Store.Migrator().HasTable(&models.CustomerAddress{}) {
		Store.AutoMigrate(&models.CustomerAddress{})
	}
	if !Store.Migrator().HasTable(&models.Consignee{}) {
		Store.AutoMigrate(&models.Consignee{})
	}
	if !Store.Migrator().HasTable(&models.LastInvoice{}) {
		Store.AutoMigrate(&models.LastInvoice{})
	}
	if !Store.Migrator().HasTable(&models.OrderPlan{}) {
		Store.AutoMigrate(&models.OrderPlan{})
	}
	if !Store.Migrator().HasTable(&models.HistoryOrderPlan{}) {
		Store.AutoMigrate(&models.HistoryOrderPlan{})
	}
	if !Store.Migrator().HasTable(&models.OrderGroupType{}) {
		Store.AutoMigrate(&models.OrderGroupType{})
	}
	if !Store.Migrator().HasTable(&models.OrderGroup{}) {
		Store.AutoMigrate(&models.OrderGroup{})
	}
	if !Store.Migrator().HasTable(&models.OrderTitle{}) {
		Store.AutoMigrate(&models.OrderTitle{})
	}
	if !Store.Migrator().HasTable(&models.OrderLoadingArea{}) {
		Store.AutoMigrate(&models.OrderLoadingArea{})
	}
	if !Store.Migrator().HasTable(&models.Order{}) {
		Store.AutoMigrate(&models.Order{})
	}
	if !Store.Migrator().HasTable(&models.OrderDetail{}) {
		Store.AutoMigrate(&models.OrderDetail{})
	}
	if !Store.Migrator().HasTable(&models.Location{}) {
		Store.AutoMigrate(&models.Location{})
	}
	if !Store.Migrator().HasTable(&models.Carton{}) {
		Store.AutoMigrate(&models.Carton{})
	}
	if !Store.Migrator().HasTable(&models.AutoGenerateInvoice{}) {
		Store.AutoMigrate(&models.AutoGenerateInvoice{})
	}
	if !Store.Migrator().HasTable(&models.LineNotifyToken{}) {
		Store.AutoMigrate(&models.LineNotifyToken{})
	}
	if !Store.Migrator().HasTable(&models.CartonNotReceive{}) {
		Store.AutoMigrate(&models.CartonNotReceive{})
	}
	if !Store.Migrator().HasTable(&models.PalletType{}) {
		Store.AutoMigrate(&models.PalletType{})
	}
	if !Store.Migrator().HasTable(&models.LastFticket{}) {
		Store.AutoMigrate(&models.LastFticket{})
	}
	if !Store.Migrator().HasTable(&models.Pallet{}) {
		Store.AutoMigrate(&models.Pallet{})
	}
	if !Store.Migrator().HasTable(&models.PalletDetail{}) {
		Store.AutoMigrate(&models.PalletDetail{})
	}
	if !Store.Migrator().HasTable(&models.OrderPrepare{}) {
		Store.AutoMigrate(&models.OrderPrepare{})
	}
	if !Store.Migrator().HasTable(&models.ImportInvoiceTap{}) {
		Store.AutoMigrate(&models.ImportInvoiceTap{})
	}

	if !Store.Migrator().HasTable(&models.AssetType{}) {
		Store.AutoMigrate(&models.AssetType{})
	}

	if !Store.Migrator().HasTable(&models.PrintShippingLabel{}) {
		Store.AutoMigrate(&models.PrintShippingLabel{})
	}

	if !Store.Migrator().HasTable(&models.PlanningDay{}) {
		Store.AutoMigrate(&models.PlanningDay{})
	}

	if !Store.Migrator().HasTable(&models.SchedulePlan{}) {
		Store.AutoMigrate(&models.SchedulePlan{})
	}
}
