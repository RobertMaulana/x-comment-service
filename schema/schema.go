package schema

import (
	"github.com/jinzhu/gorm"
	"time"
)

var (
	Database *gorm.DB
)

type TableInterface interface {
	Pk() string
	Ref() string
	AddForeignKeys()
	InsertDefaults()
}

type Base struct {
	Id        int        `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index"`
}

func AutoMigrate(database *gorm.DB) {

	Database = database

	// Drop all the tables
	//database.DropTableIfExists(
	//	Comment{},
	//)

	// Auto migrate
	database.AutoMigrate(
		Comments{},
		Organizations{},
	)

	// Relationship
	Comments{}.AddForeignKeys()

	// Relationship
	Organizations{}.InsertDefaults()
}
