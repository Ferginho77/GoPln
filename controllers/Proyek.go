package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"GoPln/config" // Ganti dengan path package db/config Anda
	"GoPln/models"
)

// CREATE
func CreateProyek(c *gin.Context) {
	var input models.Proyek

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Format data tidak valid",
			"error":   err.Error(),
		})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat Proyek",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Proyek berhasil dibuat",
		"data":    input,
	})
}

// READ ALL
func GetProyeks(c *gin.Context) {
	var Proyeks []models.Proyek

	if err := config.DB.Order("ProyekId DESC").Find(&Proyeks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data Proyek",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Proyek berhasil diambil",
		"data":    Proyeks,
	})
}

// READ BY ID
func GetProyekById(c *gin.Context) {
	id := c.Param("ProyekId")
	var Proyek models.Proyek

	if err := config.DB.First(&Proyek, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Proyek tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil Proyek",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Proyek berhasil diambil",
		"data":    Proyek,
	})
}

// UPDATE
func UpdateProyek(c *gin.Context) {
	id := c.Param("ProyekId")
	var Proyek models.Proyek

	// Cek apakah data ada
	if err := config.DB.First(&Proyek, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Proyek tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal memperbarui Proyek",
			"error":   err.Error(),
		})
		return
	}

	var input models.Proyek
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Format data tidak valid",
			"error":   err.Error(),
		})
		return
	}

	// Update kolom spesifik
	if err := config.DB.Model(&Proyek).Updates(models.Proyek{
		NamaProyek: input.NamaProyek,
		Jenis:      input.Jenis,
		Perizinan:  input.Perizinan,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal memperbarui Proyek",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Proyek berhasil diperbarui",
		"data":    Proyek,
	})
}

// DELETE
func DeleteProyek(c *gin.Context) {
	id := c.Param("ProyekId")
	var Proyek models.Proyek

	if err := config.DB.First(&Proyek, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Proyek tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus Proyek",
			"error":   err.Error(),
		})
		return
	}

	if err := config.DB.Delete(&Proyek).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus Proyek",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Proyek berhasil dihapus",
	})
}