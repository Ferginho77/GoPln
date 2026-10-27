package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"GoPln/config" // ganti dengan path package db/config Anda
	"GoPln/models"
)

// READ ALL
func GetAnggaran(c *gin.Context) {
	var anggaran []models.Anggaran

	// Query dengan order AnggaranId DESC
	if err := config.DB.Order("AnggaranId DESC").Find(&anggaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, anggaran)
}

// CREATE
func CreateAnggaran(c *gin.Context) {
	var input models.Anggaran

	// Binding JSON request body ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validasi input field wajib
	if input.Bidang == "" || input.Tahapan == "" || input.RencanaAnggaran == 0 ||
		input.Realisasi == 0 || input.TanggalAnggaran.IsZero() || input.ProyekId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Semua field wajib diisi"})
		return
	}

	// Insert data ke database
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Anggaran berhasil dibuat",
		"data":    input,
	})
}

// DELETE
func DeleteAnggaran(c *gin.Context) {
	id := c.Param("AnggaranId")
	var anggaran models.Anggaran

	// Cek keberadaan data berdasarkan Primary Key
	if err := config.DB.First(&anggaran, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Anggaran tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Hapus data
	if err := config.DB.Delete(&anggaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Anggaran berhasil dihapus"})
}

// UPDATE
func UpdateAnggaran(c *gin.Context) {
	id := c.Param("AnggaranId")
	var anggaran models.Anggaran

	// Cek keberadaan data
	if err := config.DB.First(&anggaran, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Anggaran tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var input models.Anggaran
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Update data di database
	if err := config.DB.Model(&anggaran).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Anggaran berhasil diperbarui",
		"data":    anggaran,
	})
}