package schemas

type AlamatRequest struct {
	Jalan   string `json:"jalan"`
	Kota    string `json:"kota"`
	KodePos string `json:"kode_pos"`
}

type CreateMahasiswaRequest struct {
	Nama   string        `json:"nama" binding:"required"`
	Umur   int           `json:"umur" binding:"required, min=11, max=30"`
	Hobi   []string      `json:"hobi"`
	NoHP   *string       `json:"no_hp"`
	Nilai  []int         `json:"nilai"`
	Alamat AlamatRequest `json:"alamat"`
}
