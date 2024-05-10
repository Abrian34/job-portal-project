package migration

import (
	"job-portal-project/api/config"

	entities "job-portal-project/api/entities"

	//transactionworkshopentities "job-portal-project/api/entities/transaction/workshop"

	"time"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Migrate() {
	config.InitEnvConfigs(false, "")
	logEntry := "Auto Migrating..."

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.EnvConfigs.DBHost, config.EnvConfigs.DBUser, config.EnvConfigs.DBPass, config.EnvConfigs.DBName, config.EnvConfigs.DBPort)

	//init logger
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	//constraint foreign key tidak akan ke create jika DisableForeignKeyConstraintWhenMigrating: true
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // Set the logger for GORM
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "dbo.", // schema name
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	// db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		//TablePrefix:   "dbo.", // schema name
	// 		SingularTable: false,
	// 	}, DisableForeignKeyConstraintWhenMigrating: false})

	db.AutoMigrate( // sesuai urutan foreign key
		&entities.Job{},
	)

	if db != nil && db.Error != nil {
		log.Printf("%s Failed with error %s", logEntry, db.Error)
		panic(err)
	}

	log.Printf("%s Success", logEntry)
}
