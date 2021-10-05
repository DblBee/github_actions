package main

import (
	"fmt"
	"os"

	"github.com/dblbee/github_actions/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var repo models.AccountRepo

func SetupEnvFile() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	godotenv.Load(".env")
}

func init() {
	SetupEnvFile()

	repo = models.NewAccountRepo()
}

func main() {
	app := fiber.New(fiber.Config{})

	app.Get("/", HealthCheck)

	app.Listen(":" + os.Getenv("PORT"))
}

func HealthCheck(c *fiber.Ctx) error {
	GetAccount()
	c.SendString("Alive - V1.0.1")
	return c.SendStatus(200)
}

func GetAccount() {
	account, err := repo.Get(uuid.UUID{})
	fmt.Println("Error: ", err)
	fmt.Println("Account: ", account)
}
