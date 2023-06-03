package model

import (
	"fmt"
	"ginweb/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type BaseModel struct {
	Id          int       `gorm:"primary_key" json:"id"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:datetime(0);autoUpdateTime" json:"updatedTime"`
	CreatedTime time.Time `gorm:"column:created_time;type:datetime(0);autoCreateTime" json:"createdTime"`
	IsDeleted   string    `gorm:"column:is_deleted;default:'0'" json:"isDeleted"`
}

var db *gorm.DB
var err error

/*
*
init db
*/
func SetUpDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.ConfSetting.Database.User,
		config.ConfSetting.Database.Password,
		config.ConfSetting.Database.Host,
		config.ConfSetting.Database.Name)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("sqldb error")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(1000)
}
