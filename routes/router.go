package routes

import (
	"github.com/abe27/api/controllers"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controllers.HandlerHello)

	// Group Prefix Router
	r := c.Group("/api/v1")
	// User
	r.Post("/register", controllers.Register)
	r.Post("/login", controllers.Login)
	r.Group("/carton/history").Post("", controllers.CreateCartonHistory)
	r.Group("/notify").Get("", controllers.GetAllLineNotifyToken)
	r.Group("/receive/notscan").Get("", controllers.GetAllCartonNotReceive)
	r.Group("/receive/notscan").Put("/:id", controllers.UpdateCartonNotReceiveByID)

	// Start Group Router
	log := r.Group("/logs")
	log.Get("", controllers.GetAllSyncLogger)
	log.Post("", controllers.CreateSyncLogger)
	log.Get("/:id", controllers.ShowSyncLoggerByID)
	log.Put("/:id", controllers.UpdateSyncLoggerByID)
	log.Delete("/:id", controllers.DeleteSyncLoggerByID)
	// Use Router Middleware
	app := r.Use(services.AuthorizationRequired)

	fileUpload := app.Group("upload")
	fileUpload.Post("/receive", controllers.UploadReceiveExcel)
	fileUpload.Post("/invoice/tap", controllers.ImportInvoiceTap)
	fileUpload.Post("/invoice/client_tap", controllers.ClientImportInvoiceTap)
	// fileUpload.Get("/stock", controllers.Verify)
	// fileUpload.Get("/carton", controllers.Logout)

	auth := app.Group("auth")
	auth.Get("/me", controllers.Profile)
	auth.Put("/me", controllers.UpdateProfile)
	auth.Get("/verify", controllers.Verify)
	auth.Get("/logout", controllers.Logout)

	// Administrator Router
	administrator := app.Group("administrator")
	administrator.Get("", controllers.GetAllAdministrator)
	administrator.Post("", controllers.CreateAdministrator)
	administrator.Get("/:id", controllers.ShowAdministratorByID)
	administrator.Put("/:id", controllers.UpdateAdministratorByID)
	administrator.Delete("/:id", controllers.DeleteAdministratorByID)

	// Area Router
	area := app.Group("/area")
	area.Get("", controllers.GetAllArea)
	area.Post("", controllers.CreateArea)
	area.Get("/:id", controllers.ShowAreaByID)
	area.Put("/:id", controllers.UpdateAreaByID)
	area.Delete("/:id", controllers.DeleteAreaByID)

	// Whs Router
	stock := app.Group("/stock/data")
	stock.Get("", controllers.GetAllStock)
	stock.Post("", controllers.CreateStock)
	stock.Get("/:id", controllers.ShowStockByID)
	stock.Put("/:id", controllers.UpdateStockByID)
	stock.Delete("/:id", controllers.DeleteStockByID)

	shelve := app.Group("/stock/shelve")
	shelve.Get("/:shelve_no", controllers.GetAllStockByShelve)
	stockSerialNo := app.Group("/stock/serial_no")
	stockSerialNo.Get("/:serial_no", controllers.GetAllStockBySerialNo)
	stockSerialNo.Put("/:serial_no", controllers.UpdateStockBySerialNo)

	// Whs Router
	whs := app.Group("/whs")
	whs.Get("", controllers.GetAllWhs)
	whs.Post("", controllers.CreateWhs)
	whs.Get("/:id", controllers.ShowWhsByID)
	whs.Put("/:id", controllers.UpdateWhsByID)
	whs.Delete("/:id", controllers.DeleteWhsByID)

	// Factory Router
	factory := app.Group("/factory")
	factory.Get("", controllers.GetAllFactory)
	factory.Post("", controllers.CreateFactory)
	factory.Get("/:id", controllers.ShowFactoryByID)
	factory.Put("/:id", controllers.UpdateFactoryByID)
	factory.Delete("/:id", controllers.DeleteFactoryByID)

	// Factory Router
	factoryAutoGen := app.Group("/generate/invoice")
	factoryAutoGen.Get("", controllers.GetAllAutoGenerateInvoice)
	factoryAutoGen.Post("", controllers.CreateAutoGenerateInvoice)
	factoryAutoGen.Get("/:id", controllers.ShowAutoGenerateInvoiceByID)
	factoryAutoGen.Put("/:id", controllers.UpdateAutoGenerateInvoiceByID)
	factoryAutoGen.Delete("/:id", controllers.DeleteAutoGenerateInvoiceByID)

	lineToken := app.Group("/notify")
	lineToken.Post("", controllers.CreateLineNotifyToken)
	lineToken.Get("/:id", controllers.ShowLineNotifyTokenByID)
	lineToken.Put("/:id", controllers.UpdateLineNotifyTokenByID)
	lineToken.Delete("/:id", controllers.DeleteLineNotifyTokenByID)

	// Prefix Name Router
	prefixName := app.Group("/prefixname")
	prefixName.Get("", controllers.GetAllPrefixName)
	prefixName.Post("", controllers.CreatePrefixName)
	prefixName.Get("/:id", controllers.ShowPrefixNameByID)
	prefixName.Put("/:id", controllers.UpdatePrefixNameByID)
	prefixName.Delete("/:id", controllers.DeletePrefixNameByID)

	// Position Router
	position := app.Group("/position")
	position.Get("", controllers.GetAllPosition)
	position.Post("", controllers.CreatePosition)
	position.Get("/:id", controllers.ShowPositionByID)
	position.Put("/:id", controllers.UpdatePositionByID)
	position.Delete("/:id", controllers.DeletePositionByID)

	// Department Router
	department := app.Group("/department")
	department.Get("", controllers.GetAllDepartment)
	department.Post("", controllers.CreateDepartment)
	department.Get("/:id", controllers.ShowDepartmentByID)
	department.Put("/:id", controllers.UpdateDepartmentByID)
	department.Delete("/:id", controllers.DeleteDepartmentByID)

	// Unit Router
	unit := app.Group("/unit")
	unit.Get("", controllers.GetAllUnit)
	unit.Post("", controllers.CreateUnit)
	unit.Get("/:id", controllers.ShowUnitByID)
	unit.Put("/:id", controllers.UpdateUnitByID)
	unit.Delete("/:id", controllers.DeleteUnitByID)

	// Unit Router
	asset := app.Group("/asset")
	assetType := asset.Group("/type")
	assetType.Get("", controllers.GetAllAssetType)
	assetType.Post("", controllers.CreateAssetType)
	assetType.Get("/:id", controllers.ShowAssetTypeByID)
	assetType.Put("/:id", controllers.UpdateAssetTypeByID)
	assetType.Delete("/:id", controllers.DeleteAssetTypeByID)

	// Pc Router
	pc := app.Group("/pc")
	pc.Get("", controllers.GetAllPc)
	pc.Post("", controllers.CreatePc)
	pc.Get("/:id", controllers.ShowPcByID)
	pc.Put("/:id", controllers.UpdatePcByID)
	pc.Delete("/:id", controllers.DeletePcByID)

	// Commercial Router
	commercial := app.Group("/commercial")
	commercial.Get("", controllers.GetAllCommercial)
	commercial.Post("", controllers.CreateCommercial)
	commercial.Get("/:id", controllers.ShowCommercialByID)
	commercial.Put("/:id", controllers.UpdateCommercialByID)
	commercial.Delete("/:id", controllers.DeleteCommercialByID)

	// SampleFlg Router
	sampleflg := app.Group("/sampleflg")
	sampleflg.Get("", controllers.GetAllSampleFlg)
	sampleflg.Post("", controllers.CreateSampleFlg)
	sampleflg.Get("/:id", controllers.ShowSampleFlgByID)
	sampleflg.Put("/:id", controllers.UpdateSampleFlgByID)
	sampleflg.Delete("/:id", controllers.DeleteSampleFlgByID)

	// Shipment Type Router
	shipment := app.Group("shipment")
	shipment.Get("", controllers.GetAllShipment)
	shipment.Post("", controllers.CreateShipment)
	shipment.Get("/:id", controllers.ShowShipmentByID)
	shipment.Put("/:id", controllers.UpdateShipmentByID)
	shipment.Delete("/:id", controllers.DeleteShipmentByID)

	location := app.Group("location")
	location.Get("", controllers.GetAllLocation)
	location.Post("", controllers.CreateLocation)
	location.Get("/:id", controllers.ShowLocationByID)
	location.Put("/:id", controllers.UpdateLocationByID)
	location.Delete("/:id", controllers.DeleteLocationByID)

	// Part Type Router
	partType := app.Group("parttype")
	partType.Get("", controllers.GetAllPartType)
	partType.Post("", controllers.CreatePartType)
	partType.Get("/:id", controllers.ShowPartTypeByID)
	partType.Put("/:id", controllers.UpdatePartTypeByID)
	partType.Delete("/:id", controllers.DeletePartTypeByID)

	palletType := app.Group("pallettype")
	palletType.Get("", controllers.GetAllPalletType)
	palletType.Post("", controllers.CreatePalletType)
	palletType.Get("/:id", controllers.ShowPalletTypeByID)
	palletType.Put("/:id", controllers.UpdatePalletTypeByID)
	palletType.Delete("/:id", controllers.DeletePalletTypeByID)

	// Part Type Router
	part := app.Group("part")
	part.Get("", controllers.GetAllPart)
	part.Post("", controllers.CreatePart)
	part.Get("/:id", controllers.ShowPartByID)
	part.Put("/:id", controllers.UpdatePartByID)
	part.Delete("/:id", controllers.DeletePartByID)

	ledger := app.Group("ledger")
	ledger.Get("", controllers.GetAllLedger)
	ledger.Post("", controllers.CreateLedger)
	ledger.Get("/:id", controllers.ShowLedgerByID)
	ledger.Put("/:id", controllers.UpdateLedgerByID)
	ledger.Delete("/:id", controllers.DeleteLedgerByID)

	// Planning Day Router
	planning := app.Group("planning_day")
	planning.Get("", controllers.GetAllPlanningDay)
	planning.Post("", controllers.CreatePlanningDay)
	planning.Get("/:id", controllers.ShowPlanningDayByID)
	planning.Put("/:id", controllers.UpdatePlanningDayByID)
	planning.Delete("/:id", controllers.DeletePlanningDayByID)

	edi := app.Group("edi")
	ediType := edi.Group("type")
	ediType.Get("", controllers.GetAllFileType)
	ediType.Post("", controllers.CreateFileType)
	ediType.Get("/:id", controllers.ShowFileTypeByID)
	ediType.Put("/:id", controllers.UpdateFileTypeByID)
	ediType.Delete("/:id", controllers.DeleteFileTypeByID)

	mailbox := edi.Group("mailbox")
	mailbox.Get("", controllers.GetAllMailbox)
	mailbox.Post("", controllers.CreateMailbox)
	mailbox.Get("/:id", controllers.ShowMailboxByID)
	mailbox.Put("/:id", controllers.UpdateMailboxByID)
	mailbox.Delete("/:id", controllers.DeleteMailboxByID)

	lastFTicket := app.Group("last/fticket")
	lastFTicket.Get("", controllers.GetAllLastFTicket)
	lastFTicket.Post("", controllers.CreateLastFTicket)
	lastFTicket.Get("/:id", controllers.ShowLastFTicketByID)
	lastFTicket.Put("/:id", controllers.UpdateLastFTicketByID)
	lastFTicket.Delete("/:id", controllers.DeleteLastFTicketByID)

	fileEdi := edi.Group("file")
	fileEdi.Get("", controllers.GetAllFileEdi)
	fileEdi.Post("", controllers.CreateFileEdi)
	fileEdi.Get("/:id", controllers.ShowFileEdiByID)
	fileEdi.Put("/:id", controllers.UpdateFileEdiByID)
	fileEdi.Delete("/:id", controllers.DeleteFileEdiByID)
	fileEdi.Patch("", controllers.CheckFileEdiByID)
	// Receive Type
	receive := app.Group("receive")
	receiveType := receive.Group("type")
	receiveType.Get("", controllers.GetAllReceiveType)
	receiveType.Post("", controllers.CreateReceiveType)
	receiveType.Get("/:id", controllers.ShowReceiveTypeByID)
	receiveType.Put("/:id", controllers.UpdateReceiveTypeByID)
	receiveType.Delete("/:id", controllers.DeleteReceiveTypeByID)

	receiveEnt := receive.Group("/ent")
	receiveEnt.Get("", controllers.GetAllReceiveEnt)
	receiveEnt.Post("", controllers.CreateReceiveEnt)
	receiveEnt.Get("/:id", controllers.ShowReceiveEntByID)
	receiveEnt.Put("/:id", controllers.UpdateReceiveEntByID)
	receiveEnt.Delete("/:id", controllers.DeleteReceiveEntByID)

	//ReviseOrder Router
	orderGroup := app.Group("/order")
	revise := orderGroup.Group("/revise")
	revise.Get("", controllers.GetAllReviseOrder)
	revise.Post("", controllers.CreateReviseOrder)
	revise.Get("/:id", controllers.ShowReviseOrderByID)
	revise.Put("/:id", controllers.UpdateReviseOrderByID)
	revise.Delete("/:id", controllers.DeleteReviseOrderByID)

	orderZone := orderGroup.Group("/zone")
	orderZone.Get("", controllers.GetAllOrderZone)
	orderZone.Post("", controllers.CreateOrderZone)
	orderZone.Get("/:id", controllers.ShowOrderZoneByID)
	orderZone.Put("/:id", controllers.UpdateOrderZoneByID)
	orderZone.Delete("/:id", controllers.DeleteOrderZoneByID)

	orderType := orderGroup.Group("/type")
	orderType.Get("", controllers.GetAllOrderType)
	orderType.Post("", controllers.CreateOrderType)
	orderType.Get("/:id", controllers.ShowOrderTypeByID)
	orderType.Put("/:id", controllers.UpdateOrderTypeByID)
	orderType.Delete("/:id", controllers.DeleteOrderTypeByID)

	orderGroupType := orderGroup.Group("/grouptype")
	orderGroupType.Get("", controllers.GetAllOrderGroupType)
	orderGroupType.Post("", controllers.CreateOrderGroupType)
	orderGroupType.Get("/:id", controllers.ShowOrderGroupTypeByID)
	orderGroupType.Put("/:id", controllers.UpdateOrderGroupTypeByID)
	orderGroupType.Delete("/:id", controllers.DeleteOrderGroupTypeByID)

	orderGroupConsignee := orderGroup.Group("/consignee")
	orderGroupConsignee.Get("", controllers.GetAllOrderGroup)
	orderGroupConsignee.Post("", controllers.CreateOrderGroup)
	orderGroupConsignee.Get("/:id", controllers.ShowOrderGroupByID)
	orderGroupConsignee.Put("/:id", controllers.UpdateOrderGroupByID)
	orderGroupConsignee.Delete("/:id", controllers.DeleteOrderGroupByID)

	orderTitle := orderGroup.Group("/title")
	orderTitle.Get("", controllers.GetAllOrderTitle)
	orderTitle.Post("", controllers.CreateOrderTitle)
	orderTitle.Get("/:id", controllers.ShowOrderTitleByID)
	orderTitle.Put("/:id", controllers.UpdateOrderTitleByID)
	orderTitle.Delete("/:id", controllers.DeleteOrderTitleByID)

	orderLoadingArea := orderGroup.Group("/loading")
	orderLoadingArea.Get("", controllers.ShowAllOrderLoadingArea)
	orderLoadingArea.Post("", controllers.CreateOrderLoadingArea)
	orderLoadingArea.Get("/:id", controllers.ShowOrderLoadingAreaByID)
	orderLoadingArea.Put("/:id", controllers.UpdateOrderLoadingAreaByID)
	orderLoadingArea.Delete("/:id", controllers.DeleteOrderLoadingAreaByID)

	orderEnt := orderGroup.Group("/ent")
	orderEnt.Get("", controllers.GetAllOrder)
	// orderEnt.Post("", controllers.GetAllOrder)
	orderEnt.Get("/:id", controllers.ShowOrderByID)
	orderEnt.Put("/:id", controllers.UpdateOrderByID)
	orderEnt.Delete("/:id", controllers.DeleteOrderGroupByID)
	orderEnt.Patch("", controllers.GenerateOrder)

	orderDetail := orderGroup.Group("/detail")
	orderDetail.Get("", controllers.GetAllOrderDetail)
	orderDetail.Post("", controllers.CreateOrderDetail)
	orderDetail.Get("/:id", controllers.ShowOrderDetailByID)
	orderDetail.Put("/:id", controllers.UpdateOrderDetailByID)
	orderDetail.Patch("/:id", controllers.UpdateSyncOrderDetailByID)
	orderDetail.Delete("/:id", controllers.DeleteOrderDetailByID)

	orderPallet := orderGroup.Group("/pallet")
	orderPallet.Get("", controllers.GetAllOrderPallet)
	orderPallet.Post("", controllers.CreateOrderPallet)
	orderPallet.Get("/:id", controllers.ShowOrderPalletByID)
	orderPallet.Put("/:id", controllers.UpdateOrderPalletByID)
	orderPallet.Delete("/:id", controllers.DeleteOrderPalletByID)

	orderPlan := orderGroup.Group("/plan")
	orderPlan.Get("", controllers.GetAllOrderPlan)
	orderPlan.Post("", controllers.CreateOrderPlan)
	orderPlan.Get("/:id", controllers.ShowOrderPlanByID)
	orderPlan.Put("/:id", controllers.UpdateOrderPlanByID)
	orderPlan.Patch("", controllers.GetAllOrderPlanToSync)
	orderPlan.Delete("/:id", controllers.DeleteOrderPlanByID)

	printLabel := orderGroup.Group("/label")
	printLabel.Post("", controllers.CreatePrintLabel)

	orderShort := orderGroup.Group("/short")
	orderShort.Get("", controllers.ShowAllOrderShort)
	orderShort.Post("", controllers.CreateOrderShort)
	orderShort.Get("/:id", controllers.ShowOrderShortByID)
	orderShort.Put("/:id", controllers.UpdateOrderShortByID)
	orderShort.Delete("/:id", controllers.DeleteOrderShortByID)

	affcode := app.Group("affcode")
	affcode.Get("", controllers.GetAllAffcode)
	affcode.Post("", controllers.CreateAffcode)
	affcode.Get("/:id", controllers.ShowAffcodeByID)
	affcode.Put("/:id", controllers.UpdateAffcodeByID)
	affcode.Delete("/:id", controllers.DeleteAffcodeByID)

	customer := app.Group("customer")
	customer.Get("", controllers.GetAllCustomer)
	customer.Post("", controllers.CreateCustomer)
	customer.Get("/:id", controllers.ShowCustomerByID)
	customer.Put("/:id", controllers.UpdateCustomerByID)
	customer.Delete("/:id", controllers.DeleteCustomerByID)

	customerAddress := app.Group("customeraddress")
	customerAddress.Get("", controllers.GetAllCustomerAddress)
	customerAddress.Post("", controllers.CreateCustomerAddress)
	customerAddress.Get("/:id", controllers.ShowCustomerAddressByID)
	customerAddress.Put("/:id", controllers.UpdateCustomerAddressByID)
	customerAddress.Delete("/:id", controllers.DeleteCustomerAddressByID)

	consignee := app.Group("consignee")
	consignee.Get("", controllers.GetAllConsignee)
	consignee.Post("", controllers.CreateConsignee)
	consignee.Get("/:id", controllers.ShowConsigneeByID)
	consignee.Put("/:id", controllers.UpdateConsigneeByID)
	consignee.Delete("/:id", controllers.DeleteConsigneeByID)

	invoiceGroup := app.Group("invoice")
	fticketGroup := invoiceGroup.Group("/shipping_label")
	fticketGroup.Get("", controllers.GetAllShippingLabel)
	fticketGroup.Post("", controllers.CreateShippingLabel)
	fticketGroup.Get("/:id", controllers.ShowShippingLabelByID)
	fticketGroup.Put("/:id", controllers.UpdateShippingLabelByID)
	fticketGroup.Delete("/:id", controllers.DeleteShippingLabelByID)
}
