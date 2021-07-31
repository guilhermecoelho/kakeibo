package configurations

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/guilhermecoelho/kakeibo/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBPostgre *sql.DB
var DBgorm *gorm.DB

func InitDatabasePostgre() {
	db, err := sql.Open("postgres", "postgres://postgres:Postgres2018!@localhost/Kakeibo?sslmode=disable")
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected on postgre!\n")
	DBPostgre = db
}

func InitDatabaseGorm() {
	dsn := "host=localhost user=postgres password=Postgres2018! dbname=Kakeibo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	fmt.Printf("Connected on postgre with gorm!\n")

	DBgorm = db

}

func CreateDatabase() {

	DBgorm.Debug().Migrator().DropTable(&models.Group{})
	DBgorm.Debug().AutoMigrate(&models.Group{})
	DBgorm.Debug().Migrator().DropTable(&models.Category{})
	DBgorm.Debug().AutoMigrate(&models.Category{})
	DBgorm.Debug().Migrator().DropTable(&models.Expense{})
	DBgorm.Debug().AutoMigrate(&models.Expense{})
	DBgorm.Debug().Migrator().DropTable(&models.Income{})
	DBgorm.Debug().AutoMigrate(&models.Income{})
	DBgorm.Debug().Migrator().DropTable(&models.User{})
	DBgorm.Debug().AutoMigrate(&models.User{})
}
