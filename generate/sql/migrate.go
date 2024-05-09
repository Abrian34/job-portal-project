package migration

import (
	"job-portal-project/api/config"

	// entities "job-portal-project/api/entities"

	//transactionworkshopentities "job-portal-project/api/entities/transaction/workshop"

	"time"

	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Migrate() {
	config.InitEnvConfigs(false, "")
	logEntry := "Auto Migrating..."

	dsn := fmt.Sprintf(
		`%s://%s:%s@%s:%v?database=%s`,
		config.EnvConfigs.DBDriver,
		config.EnvConfigs.DBUser,
		config.EnvConfigs.DBPass,
		config.EnvConfigs.DBHost,
		config.EnvConfigs.DBPort,
		config.EnvConfigs.DBName,
	)

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
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
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
	// &masteroperationentities.OperationModelMapping{},
	// &masteroperationentities.OperationFrt{},
	// &masteroperationentities.OperationGroup{},
	// &masteroperationentities.OperationSection{},
	// &masteroperationentities.OperationKey{},
	// &masteroperationentities.OperationEntries{},
	// &masteroperationentities.OperationCode{},

	// &masterwarehouseentities.WarehouseGroup{},
	// &masterwarehouseentities.WarehouseMaster{},
	// &masterwarehouseentities.WarehouseLocation{},
	// &masterwarehouseentities.WarehouseLocationDefinition{},
	// &masterwarehouseentities.WarehouseLocationDefinitionLevel{},

	// &masteritementities.MarkupMaster{},
	// &masteritementities.PrincipleBrandParent{},
	// &masteritementities.UomType{},
	// &masteritementities.PriceList{},
	// &masteritementities.ItemLocationSource{},
	// &masteritementities.ItemLocationDetail{},
	// &masteritementities.ItemLocation{},
	//&masteritementities.PurchasePrice{},
	//&masteritementities.PurchasePriceDetail{},
	// &masteritementities.ItemDetail{},
	// &masteritementities.ItemImport{},
	// &masteritementities.DiscountPercent{},
	// &masteritementities.ItemSubstitute{},
	// &masteritementities.ItemSubstituteDetail{},
	// &masteritementities.ItemPackage{},
	// &masteritementities.ItemPackageDetail{},
	// &masteritementities.PrincipleBrandParent{},
	// &masteritementities.UomType{},
	// &masteritementities.Uom{},
	//&masteritementities.Bom{},
	//&masteritementities.BomDetail{},
	// &masteritementities.ItemClass{},
	// &masteritementities.Item{},
	// &masteritementities.MarkupRate{},
	// &masteritementities.ItemLevel{},

	// &entities.IncentiveGroup{},
	// &entities.ForecastMaster{},
	// &entities.MovingCode{},
	// &entities.ShiftSchedule{},
	// &entities.IncentiveGroup{},
	// &entities.ForecastMaster{},
	// &entities.MovingCode{},
	// &entities.ShiftSchedule{},
	// &entities.IncentiveMaster{},
	// &entities.IncentiveGroupDetail{},
	// &entities.SkillLevel{},
	// &entities.WarrantyFreeService{},
	// &entities.DeductionList{},
	// &entities.DeductionDetail{},
	// &entities.FieldActionEligibleVehicleItem{},
	// &entities.FieldActionEligibleVehicle{},
	// &entities.FieldAction{},
	// &entities.Discount{},

	// &transactionentities.SupplySlip{},
	// &transactionentities.SupplySlipDetail{},
	// &transactionworkshopentities.WorkOrder{},
	// &transactionentities.ServiceLog{},
	// &transactionworkshopentities.BookingEstimation{},
	)

	if db != nil && db.Error != nil {
		log.Printf("%s Failed with error %s", logEntry, db.Error)
		panic(err)
	}

	log.Printf("%s Success", logEntry)
}