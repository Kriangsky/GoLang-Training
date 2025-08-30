package models

type Alamat struct {
	Jalan   string
	Kota    string
	KodePos string
}

type Mahasiswa struct {
	ID              int
	Nama            string
	Umur            int
	Hobi            []string
	Address         Alamat
	Nomor           *string
	Nilai           []int
	Lulus           bool
	KeteranganLulus string
}
