package services

import (
	"fmt"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
)

func CreateCartonHistoryData(obj *models.CartonHistory) bool {
	db := configs.Store
	var sysLog models.SyncLogger
	sysLog.Title = "Create Carton history"
	sysLog.Description = fmt.Sprintf("%s created successfully", obj.SerialNo)
	sysLog.IsSuccess = true
	err := db.Create(&obj).Error
	if err != nil {
		sysLog.Title = fmt.Sprintf("Create %s is error", obj.SerialNo)
		sysLog.Description = err.Error()
		sysLog.IsSuccess = false
	}
	db.Create(&sysLog)
	return true
}
