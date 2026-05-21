package booking // veya config / database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Uyarı: .env dosyası bulunamadı, sistem ortam değişkenlerine bakılıyor...")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("veritabanı açılırken hata oluştu: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("veritabanına ulaşılamıyor: %w", err)
	}

	return db, nil
}
