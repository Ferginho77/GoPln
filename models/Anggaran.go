package models

import (
	"time"
	"github.com/shopspring/decimal"
)


type Anggaran struct {
	AnggaranId      uint      `json:"AnggaranId" gorm:"column:AnggaranId;primaryKey;autoIncrement"`
	Bidang          string    `json:"Bidang" gorm:"column:Bidang;type:enum('Perizinan', 'Pengadaan Tanah', 'Kompensasi ROW', 'Sertifikasi');not null"`
	Tahapan         string    `json:"Tahapan" gorm:"column:Tahapan;type:varchar(150);not null"`
	RencanaAnggaran decimal.Decimal `json:"RencanaAnggaran" gorm:"column:RencanaAnggaran;type:decimal(15,2);not null"` // Perbaikan di sini
	Realisasi       decimal.Decimal `json:"Realisasi" gorm:"column:Realisasi;type:decimal(15,2);not null"`
	TanggalAnggaran time.Time `json:"TanggalAnggaran" gorm:"column:TanggalAnggaran;not null"`
	ProyekId        uint      `json:"ProyekId" gorm:"column:ProyekId;not null"`
}

// Menyesuaikan nama tabel agar sama dengan konfigurasi Sequelize (tableName: "anggaran")
func (Anggaran) TableName() string {
	return "anggaran"
}