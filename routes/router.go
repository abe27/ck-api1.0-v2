package routes

import (
	"github.com/abe27/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controllers.HandlerHello)

	// Group Prefix Router
	r := c.Group("/api/v1")
	// User
	user := r.Group("/auth")
	user.Post("/register", controllers.Register)
	user.Post("/login", controllers.Login)
	user.Get("/me", controllers.Profile)
	user.Get("/verify", controllers.Verify)
	user.Get("/logout", controllers.Logout)

	// Administrator Router
	administrator := r.Group("administrator")
	administrator.Get("", controllers.GetAllAdministrator)
	administrator.Post("", controllers.CreateAdministrator)
	administrator.Get("/:id", controllers.ShowAdministratorByID)
	administrator.Put("/:id", controllers.UpdateAdministratorByID)
	administrator.Delete("/:id", controllers.DeleteAdministratorByID)

	// Area Router
	area := r.Group("/area")
	area.Get("", controllers.GetAllArea)
	area.Post("", controllers.CreateArea)
	area.Get("/:id", controllers.ShowAreaByID)
	area.Put("/:id", controllers.UpdateAreaByID)
	area.Delete("/:id", controllers.DeleteAreaByID)

	// Whs Router
	whs := r.Group("/whs")
	whs.Get("", controllers.GetAllWhs)
	whs.Post("", controllers.CreateWhs)
	whs.Get("/:id", controllers.ShowWhsByID)
	whs.Put("/:id", controllers.UpdateWhsByID)
	whs.Delete("/:id", controllers.DeleteWhsByID)

	// Factory Router
	factory := r.Group("/factory")
	factory.Get("", controllers.GetAllFactory)
	factory.Post("", controllers.CreateFactory)
	factory.Get("/:id", controllers.ShowFactoryByID)
	factory.Put("/:id", controllers.UpdateFactoryByID)
	factory.Delete("/:id", controllers.DeleteFactoryByID)

	// Prefix Name Router
	prefixname := r.Group("/prefixname")
	prefixname.Get("", controllers.GetAllPrefixName)
	prefixname.Post("", controllers.CreatePrefixName)
	prefixname.Get("/:id", controllers.ShowPrefixNameByID)
	prefixname.Put("/:id", controllers.UpdatePrefixNameByID)
	prefixname.Delete("/:id", controllers.DeletePrefixNameByID)

	// Position Router
	position := r.Group("/position")
	position.Get("", controllers.GetAllPosition)
	position.Post("", controllers.CreatePosition)
	position.Get("/:id", controllers.ShowPositionByID)
	position.Put("/:id", controllers.UpdatePositionByID)
	position.Delete("/:id", controllers.DeletePositionByID)

	// Department Router
	department := r.Group("/department")
	department.Get("", controllers.GetAllDepartment)
	department.Post("", controllers.CreateDepartment)
	department.Get("/:id", controllers.ShowDepartmentByID)
	department.Put("/:id", controllers.UpdateDepartmentByID)
	department.Delete("/:id", controllers.DeleteDepartmentByID)

	// Unit Router
	unit := r.Group("/unit")
	unit.Get("", controllers.GetAllUnit)
	unit.Post("", controllers.CreateUnit)
	unit.Get("/:id", controllers.ShowUnitByID)
	unit.Put("/:id", controllers.UpdateUnitByID)
	unit.Delete("/:id", controllers.DeleteUnitByID)

	// Part Type Router
	part_type := r.Group("parttype")
	part_type.Get("", controllers.GetAllPartType)
	part_type.Post("", controllers.CreatePartType)
	part_type.Get("/:id", controllers.ShowPartTypeByID)
	part_type.Put("/:id", controllers.UpdatePartTypeByID)
	part_type.Delete("/:id", controllers.DeletePartTypeByID)
	// Part Type Router
	edi := r.Group("edi")
	edi_type := edi.Group("type")
	edi_type.Get("", controllers.GetAllFileType)
	edi_type.Post("", controllers.CreateFileType)
	edi_type.Get("/:id", controllers.ShowFileTypeByID)
	edi_type.Put("/:id", controllers.UpdateFileTypeByID)
	edi_type.Delete("/:id", controllers.DeleteFileTypeByID)

	mailbox := edi.Group("mailbox")
	mailbox.Get("", controllers.GetAllMailbox)
	mailbox.Post("", controllers.CreateMailbox)
	mailbox.Get("/:id", controllers.ShowMailboxByID)
	mailbox.Put("/:id", controllers.UpdateMailboxByID)
	mailbox.Delete("/:id", controllers.DeleteMailboxByID)

	part := r.Group("part")
	part.Get("", controllers.GetAllPart)
	part.Post("", controllers.CreatePart)
	part.Get("/:id", controllers.ShowPartByID)
	part.Put("/:id", controllers.UpdatePartByID)
	part.Delete("/:id", controllers.DeletePartByID)

	ledger := r.Group("ledger")
	ledger.Get("", controllers.GetAllLedger)
	ledger.Post("", controllers.CreateLedger)
	ledger.Get("/:id", controllers.ShowLedgerByID)
	ledger.Put("/:id", controllers.UpdateLedgerByID)
	ledger.Delete("/:id", controllers.DeleteLedgerByID)

	file_edi := edi.Group("file")
	file_edi.Get("", controllers.GetAllFileEdi)
	file_edi.Post("", controllers.CreateFileEdi)
	file_edi.Get("/:id", controllers.ShowFileEdiByID)
	file_edi.Put("/:id", controllers.UpdateFileEdiByID)
	file_edi.Delete("/:id", controllers.DeleteFileEdiByID)
	// Receive Type
	receive := r.Group("receive")
	receive_type := receive.Group("type")
	receive_type.Get("", controllers.GetAllReceiveType)
	receive_type.Post("", controllers.CreateReceiveType)
	receive_type.Get("/:id", controllers.ShowReceiveTypeByID)
	receive_type.Put("/:id", controllers.UpdateReceiveTypeByID)
	receive_type.Delete("/:id", controllers.DeleteReceiveTypeByID)
}
