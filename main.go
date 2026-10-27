package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"GoPln/config"
	"GoPln/controllers"
)

func main() {
	// 1. Load file .env terlebih dahulu sebelum inisialisasi modul lain
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: File .env tidak ditemukan, menggunakan environment variable sistem")
	}

	// 2. Inisialisasi koneksi database
	config.Conn()

	// 3. Inisialisasi router Gin
	r := gin.Default()

	// 4. Konfigurasi Middleware CORS (menggunakan AllowAllOrigins agar tidak konfliktual)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	// 5. Routing API
	// Endpoint Anggaran
	r.GET("/anggaran", controllers.GetAnggaran)
	r.POST("/anggaran", controllers.CreateAnggaran)
	r.PUT("/anggaran/:AnggaranId", controllers.UpdateAnggaran)
	r.DELETE("/anggaran/:AnggaranId", controllers.DeleteAnggaran)

	// Endpoint Proyek
	r.GET("/proyek", controllers.GetProyeks)
	r.GET("/proyek/:ProyekId", controllers.GetProyekById)
	r.POST("/proyek", controllers.CreateProyek)
	r.PUT("/proyek/:ProyekId", controllers.UpdateProyek)
	r.DELETE("/proyek/:id", controllers.DeleteProyek)

	// 6. Jalankan Server HTTP (default port 8080 jika PORT di .env tidak diisi)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan pada port :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}