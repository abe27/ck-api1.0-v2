package services

import (
	"fmt"
	"time"
)

var HelloWorld = fmt.Sprintf("Hello World, %s", (time.Now()).Format("2006-01-02 15:04:05"))

func MessageShowAll(txt string) string {
	return fmt.Sprintf("Show All, %s on %s", txt, (time.Now()).Format("2006-01-02 15:04:05"))
}

func MessageNotFound(txt string) string {
	return fmt.Sprintf("Not found, %s", txt)
}

var MessageRegister = func(msg string) string { return "Register " + msg + " is completed!" }
var MessageInputValidationError = "Invalid value!"
var MessagePasswordNotMatched = "Password is valid!"
var MessageSystemError = "System is error!"
var MessageAuthentication = "Welcome to API Server!"
var MessageNotFoundUser = "User is not found!"
var MessagePasswordNotMatch = "Password is not matched!"
var MessageUserNotAuthenticated = "User is not authenticated!"
var MessageNotFoundTokenKey = "Token key is not found!"
var MessageTokenIsExpired = "Token is expired!"
var MessageUserLeave = "User is logout!"
var MessageUserNotActive = "User is not active!"

var MessageShowAllData = func(title string) string { return fmt.Sprintf("Show All `%s`!", title) }
var MessageCreatedData = func(title *string) string { return fmt.Sprintf("Create Data `%s(%d)` is completed", *title, title) }
var MessageShowDataByID = func(title *string) string { return fmt.Sprintf("Show Data by ID: `%s(%d)`", *title, title) }
var MessageUpdateDataByID = func(title *string) string {
	return fmt.Sprintf("Update Data by ID: `%s(%d)` is completed!", *title, title)
}
var MessageNotFoundData = func(title *string) string {
	return fmt.Sprintf("Not found `%s(%d)`!", *title, title)
}
var MessageDuplicateData = func(title *string) string {
	return fmt.Sprintf("`%v(%d)` is Duplicate!", *title, title)
}
var MessageDeleteData = func(title *string) string {
	return fmt.Sprintf("Delete Data by ID: `%s(%d)` is completed.", *title, title)
}
var MessageUploadFileError = func(title *string) string {
	return fmt.Sprintf("Upload File Error: `%s(%d)` is completed.", *title, title)
}
