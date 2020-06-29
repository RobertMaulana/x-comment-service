package comments_db

import (
	"database/sql"
	"fmt"
	commentSchema "github.com/RobertMaulana/x-comment-service/schema"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var (
	Client *sql.DB
	Database *gorm.DB
)

func init() {
	_ = godotenv.Load()

	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	schema := os.Getenv("PG_DBNAME")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, schema)
	var err error
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil{
		fmt.Printf("err %#v", err)
		panic(err)
	}
	Database = db
	Client = db.DB()
	autoCreate := os.Getenv("DB_AUTO_CREATE")
	if autoCreate == "true" {
		fmt.Println("Dropping and recreating all tables...")
		commentSchema.AutoMigrate(db)
		fmt.Println("All tables recreated successfully...")
	}
}