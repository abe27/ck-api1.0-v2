package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// initial database
	dns := "host=" + os.Getenv("DBHOST") +
		" user=" + os.Getenv("DBUSER") +
		" dbname=" + os.Getenv("DBNAME") +
		" port=" + os.Getenv("DBPORT") +
		" sslmode=" + os.Getenv("SSLMODE") +
		" TimeZone=" + os.Getenv("TZNAME") + ""
	if len(os.Getenv("DBPASSWORD")) > 0 {
		dns = "host=" + os.Getenv("DBHOST") +
			" user=" + os.Getenv("DBUSER") +
			" password=" + os.Getenv("DBPASSWORD") +
			" dbname=" + os.Getenv("DBNAME") +
			" port=" + os.Getenv("DBPORT") +
			" sslmode=" + os.Getenv("SSLMODE") +
			" TimeZone=" + os.Getenv("TZNAME") + ""
	}

	// fmt.Printf("DNS: %s\n", dns)
	configs.Store, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migration DB
	configs.Store.AutoMigrate(&models.User{})
	configs.Store.AutoMigrate(&models.JwtToken{})
	configs.Store.AutoMigrate(&models.Administrator{})
	configs.Store.AutoMigrate(&models.Area{})
	configs.Store.AutoMigrate(&models.Whs{})
	configs.Store.AutoMigrate(&models.Factory{})
	configs.Store.AutoMigrate(&models.ReceiveType{})
	configs.Store.AutoMigrate(&models.Unit{})
	configs.Store.AutoMigrate(&models.PartType{})
	configs.Store.AutoMigrate(&models.FileType{})
	configs.Store.AutoMigrate(&models.Position{})
	configs.Store.AutoMigrate(&models.Department{})
	configs.Store.AutoMigrate(&models.PrefixName{})
	configs.Store.AutoMigrate(&models.Profile{})
	configs.Store.AutoMigrate(&models.Mailbox{})
	configs.Store.AutoMigrate(&models.FileEdi{})
	configs.Store.AutoMigrate(&models.CartonHistory{})
	configs.Store.AutoMigrate(&models.SyncLogger{})
	configs.Store.AutoMigrate(&models.Part{})
	configs.Store.AutoMigrate(&models.Ledger{})
	configs.Store.AutoMigrate(&models.Receive{})
	configs.Store.AutoMigrate(&models.ReceiveDetail{})
	configs.Store.AutoMigrate(&models.Pc{})
	configs.Store.AutoMigrate(&models.Commercial{})
	configs.Store.AutoMigrate(&models.SampleFlg{})
	configs.Store.AutoMigrate(&models.ReviseOrder{})
	configs.Store.AutoMigrate(&models.Shipment{})
	configs.Store.AutoMigrate(&models.OrderZone{})
	configs.Store.AutoMigrate(&models.OrderType{})
	configs.Store.AutoMigrate(&models.Affcode{})
	configs.Store.AutoMigrate(&models.Customer{})
	configs.Store.AutoMigrate(&models.CustomerAddress{})
	configs.Store.AutoMigrate(&models.Consignee{})
	configs.Store.AutoMigrate(&models.LastInvoice{})
	configs.Store.AutoMigrate(&models.OrderPlan{})
	configs.Store.AutoMigrate(&models.OrderGroupType{})
	configs.Store.AutoMigrate(&models.OrderGroup{})
	configs.Store.AutoMigrate(&models.OrderTitle{})
	configs.Store.AutoMigrate(&models.OrderLoadingArea{})
	configs.Store.AutoMigrate(&models.Order{})
	configs.Store.AutoMigrate(&models.OrderDetail{})
	configs.Store.AutoMigrate(&models.Location{})
	configs.Store.AutoMigrate(&models.Carton{})
	// configs.Store.AutoMigrate(&models.LocationType{})
	// configs.Store.AutoMigrate(&models.LocationAddress{})
	// configs.Store.AutoMigrate(&models.LocationAreaType{})
}

func main() {
	// Create config variable
	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "SPL Server API Service", // add custom server header
		AppName:       "API Version 1.0",
		BodyLimit:     10 * 1024 * 1024, // this is the default limit of 10MB
	}

	app := fiber.New(config)
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Static("/", "./public")
	routes.SetUpRouter(app)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("ON_PORT")))
}
