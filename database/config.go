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
		log.Fatal("SUPABASE_DB_URL is not set in environment variables")
	}

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// DB.Migrator().DropTable(&models.Stream{}, &models.Message{})

	err = DB.AutoMigrate(
		&models.Category{}, &models.Stream{}, &models.Message{}, &models.StreamHistory{}, &models.ViewerHistory{}, &models.Highlight{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Dijalanin kalau diperlukan dan cek database dulu ya nak
	// err = SeedStreamingCategories() //dijalanin kalau perlu saja
	// err = SeedStreamInfos(100) //dijalanin kalau perlu saja
	// err = SeedStreams(100)
	// err = SeedStreamHistories(100)
	// err = SeedViewerHistories()
	// err = SeedMessages()

	if err != nil {
		log.Fatalln("failed to seed")
	}

	log.Println("Connected to Supabase PostgreSQL successfully!")
}
