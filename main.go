package main

import (
	"AbundancePath/app"
	"AbundancePath/core/model"
	"AbundancePath/core/repository"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type userRepository struct{}

func (r *userRepository) Save(user model.User) (model.User, error) {
    return user, nil
}

func (r *userRepository) FindByID(id int) (model.User, error) {
    return model.User{ID: id, Name: "John", Email: "john@example.com"}, nil
}

type userService struct {
    repo repository.UserRepository
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
    return s.repo.Save(user)
}

func (s *userService) GetUser(id int) (model.User, error) {
    return s.repo.FindByID(id)
}
type Config struct {
	DB *gorm.DB
}

// LoadConfig loads the application configuration
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connectToDatabase()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DB: db,
	}

	return config, nil
}

func connectToDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
func main() {

    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = LoadConfig()
	if err != nil {
		log.Fatal("Error loading application configuration:", err)
	}
    repo := &userRepository{}
    userService := &userService{repo: repo}
    app := app.New(userService)
    err = app.Run()
    if err != nil {
        fmt.Printf("Failed to run app: %v", err)
    }
}
