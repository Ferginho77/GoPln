package models

type Proyek struct {
	ProyekId uint   `json:"ProyekId" gorm:"column:ProyekId;primaryKey"`
	NamaProyek   string `json:"NamaProyek"   gorm:"column:NamaProyek"`
	Jenis        string `json:"Jenis"        gorm:"column:Jenis;type:enum('SUTT', 'GI')"`
	Perizinan         string   `json:"Perizinan"         gorm:"column:Perizinan;type:enum('KKPR', 'Penetapan Lokasi', 'PPKH', 'Lingkungan', 'PBG')"`
}

func (Proyek) TableName() string {
	return "proyek"
}
