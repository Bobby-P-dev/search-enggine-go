package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Bobby-P-dev/search-enggine-go/src/configs"
	"github.com/Bobby-P-dev/search-enggine-go/src/controllers"
	"github.com/Bobby-P-dev/search-enggine-go/src/repositories"
	"github.com/Bobby-P-dev/search-enggine-go/src/routes"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var db *pgxpool.Pool

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: file .env tidak ditemukan, menggunakan environment variable sistem")
	}

	var err error
	db, err = configs.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v\n", err)
	}
	defer db.Close()

	log.Println("Successfully connected to database")

	userRepository := repositories.NewUserRepositoryImpl(db)
	userController := controllers.NewUserControllerImpl(db, userRepository)
	router := routes.SetupRoutes(userController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di http://localhost:%s\n", port)
	log.Printf("Endpoint user tersedia di http://localhost:%s/api/v1/users\n", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server gagal berjalan: %v\n", err)
	}
}
