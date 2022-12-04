package database

import (
	"database/sql"
	"fiber-crud/pkg/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           string
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	TT           string
}

type Customer struct {
	Id          string     `gorm:"primaryKey;column:Id"`
	Name        string     `gorm:"column:Name"`
	Email       string     `gorm:"column:Email"`
	Age         int        `gorm:"column:Age"`
	CreatedDate *time.Time `gorm:"column:CreatedDate"`
	IsActive    bool       `gorm:"column:IsActive"`
}

func (Customer) TableName() string {
	return "CustomerX"
}

var _db *gorm.DB

func InitMsql() (gormDB *gorm.DB, db *sql.DB) {
	cf := getConfigDb()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	gormDB, err := gorm.Open(sqlserver.Open(cf+"&charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		CreateBatchSize: 300,
		Logger:          newLogger})
	if err != nil {
		panic(err)
	}

	db, err = gormDB.DB()
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return
}

func InitialDB() {
	configDb := getConfigDb()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(sqlserver.Open(configDb+"&charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		// PrepareStmt:            true,
		SkipDefaultTransaction: true,
		// CreateBatchSize:        300,
		Logger: newLogger})
	// Logger: newLogger})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	_db = db
	//Auto Migrate
	// db.AutoMigrate(&Customer{})
}

func getConfigDb() string {
	conf := config.Database()
	if conf.Port > 0 {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s:%d?database=%s",
			conf.User,
			conf.Pass,
			conf.Server,
			conf.Port,
			conf.DatabaseName,
		)
	} else {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=%s",
			conf.User,
			conf.Pass,
			conf.Server,
			conf.DatabaseName,
		)
	}
}

func Db() *gorm.DB {
	return _db
}
