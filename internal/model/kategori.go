package model

type RequestKategori struct {
	Tahun string `json:"tahun" gorm:"primaryKey"`
	Nama  string `json:"nama"`
}

type Kategori struct {
}
