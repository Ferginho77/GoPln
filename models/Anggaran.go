package models

import "time"

type Anggaran struct {
	AnggaranId      uint      `json:"AnggaranId" gorm:"column:AnggaranId;primaryKey;autoIncrement"`
	Bidang          string    `json:"Bidang" gorm:"column:Bidang;type:enum('Perizinan', 'Pengadaan Tanah', 'Kompensasi ROW', 'Sertifikasi');not null"`
	Tahapan         string    `json:"Tahapan" gorm:"column:Tahapan;type:varchar(150);not null"`
	RencanaAnggaran int       `json:"RencanaAnggaran" gorm:"column:RencanaAnggaran;not null"`
	Realisasi       int       `json:"Realisasi" gorm:"column:Realisasi;not null"`
	TanggalAnggaran time.Time `json:"TanggalAnggaran" gorm:"column:TanggalAnggaran;not null"`
	ProyekId        uint      `json:"ProyekId" gorm:"column:ProyekId;not null"`
}

// Menyesuaikan nama tabel agar sama dengan konfigurasi Sequelize (tableName: "anggaran")
func (Anggaran) TableName() string {
	return "anggaran"
}