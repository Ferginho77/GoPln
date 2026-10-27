package models

type Bidang struct {
	BidangId uint   `json:"BidangId" gorm:"column:BidangId;primaryKey"`
	NamaBidang   string `json:"NamaBidang"   gorm:"column:NamaBidang"`
	Jenis        string `json:"Jenis"        gorm:"column:Jenis;type:enum('SUTT', 'GI')"`
	Perizinan         string   `json:"Perizinan"         gorm:"column:Perizinan;type:enum('KKPR', 'Penetapan Lokasi', 'PPKH', 'Lingkungan', 'PBG')"`
}

func (Bidang) TableName() string {
	return "Bidang"
}
