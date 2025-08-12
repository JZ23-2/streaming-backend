package database

import (
	"log"
	"main/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("SUPABASE_DB_URL")
	if dsn == "" {
		log.Fatal("❌ SUPABASE_DB_URL is not set in environment variables")
	}

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// DB.Migrator().DropTable(&models.User{}, &models.Stream{})

	err = DB.AutoMigrate(
		&models.User{}, &models.Stream{}, &models.Message{},
	)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Connected to Supabase PostgreSQL successfully!")
}
