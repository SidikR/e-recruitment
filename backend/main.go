package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func main() {
	// Konfigurasi koneksi ke database PostgreSQL
	databaseURL := "user=postgres password=postgres dbname=e_rekruitment sslmode=disable"
	var err error
	db, err = gorm.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Migrate model-model Anda ke dalam database (opsional)
	// db.AutoMigrate(&Model{})

	// Inisialisasi router Gin
	r := gin.Default()

	// Definisikan rute-rute Anda dan gunakan koneksi database
	// ...

	// Jalankan server Gin
	r.Run(":8080")
}
