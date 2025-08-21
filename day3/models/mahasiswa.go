package models

type Alamat struct {
	Jalan   string
	Kota    string
	KodePos int
}

type Mahasiswa struct {
	Nama            string
	Umur            int
	Hobi            []string
	Address         Alamat
	Nomor           *string
	Nilai           []int
	Lulus           bool
	KeteranganLulus string
}