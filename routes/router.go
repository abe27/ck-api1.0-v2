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
	// Use Router Middleware
	app := r.Use(services.AuthorizationRequired)
	auth := app.Group("auth")
	auth.Get("/me", controllers.Profile)
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

	// Prefix Name Router
	prefixname := app.Group("/prefixname")
	prefixname.Get("", controllers.GetAllPrefixName)
	prefixname.Post("", controllers.CreatePrefixName)
	prefixname.Get("/:id", controllers.ShowPrefixNameByID)
	prefixname.Put("/:id", controllers.UpdatePrefixNameByID)
	prefixname.Delete("/:id", controllers.DeletePrefixNameByID)

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

	// Part Type Router
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
	receiveType.Put("/:id", controllers.UpdateReceiveEntByID)
	receiveType.Delete("/:id", controllers.DeleteReceiveEntByID)

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
	orderEnt.Get("", controllers.GetAllOrderGroup)
	orderEnt.Post("", controllers.CreateOrderGroup)
	orderEnt.Get("/:id", controllers.ShowOrderGroupByID)
	orderEnt.Put("/:id", controllers.UpdateOrderGroupByID)
	orderEnt.Delete("/:id", controllers.DeleteOrderGroupByID)
	orderEnt.Patch("/generate", controllers.GenerateOrder)

	orderDetail := orderGroup.Group("/detail")
	orderDetail.Get("", controllers.GetAllOrderGroup)
	orderDetail.Post("", controllers.CreateOrderGroup)
	orderDetail.Get("/:id", controllers.ShowOrderGroupByID)
	orderDetail.Put("/:id", controllers.UpdateOrderGroupByID)
	orderDetail.Delete("/:id", controllers.DeleteOrderGroupByID)

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
}
