package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"

	"os"
	"pi-incident-paging-system/alert"
	"pi-incident-paging-system/webhook_data"
	"strconv"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS
)

func main() {
	// Load Environment variables from .env
	godotenv.Load()

	// Load up database
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	_, gormErr := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	if gormErr != nil {
		panic(gormErr)
	}

	// Run database migrations
	migrationsDirectory, _ := iofs.New(migrations, "migrations")

	migrations, migrationsErr := migrate.NewWithSourceInstance(
		"iofs", migrationsDirectory, databaseUrl)

	if migrationsErr == nil {
		migrations.Up()
	}

	// Setup routing
	r := gin.Default()

	// 1923929
	// 138075000

	r.POST("/opsgenie/alert/created", func(c *gin.Context) {
		var input webhook_data.OpsgenieWebhookCreate

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		alert.SendAlert(input.Alert.Priority, input.Alert.Message, 138075000, []string{
			"1923929",
		})
	})

	r.Run()
}