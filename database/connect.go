package database

import (
	"fmt"
	"log"
	"os"
	"ticketing-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal konek ke database! Cek Docker kamu.", err)
	}

	fmt.Println("Sukses konek ke Database!")

	err = database.AutoMigrate(
		&models.User{},
		&models.UserStats{},
		&models.Project{},
		&models.ProjectMember{},
		&models.ProjectInvite{},
		&models.Ticket{},
		&models.Comment{},
		&models.TicketHistory{},
		&models.Badge{},
		&models.UserBadge{},
		&models.XPEvent{},
		&models.Challenge{},
		&models.ChallengeParticipant{},
		&models.LeaderboardSnapshot{},
	)

	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}

	DB = database
}
