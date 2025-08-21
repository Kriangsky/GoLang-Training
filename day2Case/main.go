package main

import (
	"GoLang/day2Case/models"
	"GoLang/day2Case/utils"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(models.Tambah(1, 2))

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Jangan takut: ", r)
		}
	}()

	// defer fmt.Println("Tes1")
	// defer fmt.Println("Tes2")
	// panic("OH NOOOO")
	nomorSky := "0812992772"
	// nomorJohn := "1293712937192"

	var data []models.Mahasiswa = []models.Mahasiswa{
		{
			Nama: "Sky",
			Umur: 18,
			Hobi: []string{"Berenang", "Basket"},
			Address: models.Alamat{
				Jalan:   "Kedondong",
				Kota:    "Gedong Panjang",
				KodePos: 1224,
			},
			Nomor: &nomorSky,
			Nilai: []int{88, 94, 98},
		},
		{
			Nama: "John Doe",
			Umur: 18,
			Hobi: []string{"Bermain", "Clubbing"},
			Address: models.Alamat{
				Jalan:   "Kedondong",
				Kota:    "Gedong Panjang",
				KodePos: 1224,
			},
			Nilai: []int{80, 60, 98},
		},
	}

	avg := 0.0

	for index, _ := range data {
		err := utils.ValidasiNilai(data[index].Nilai)

		if err == nil {
			fmt.Println("Yey nilai sudah benar")
		} else {
			fmt.Println(err)
		}

		err = utils.CekNoHP("081299817566")

		if err == nil {
			fmt.Println("Yey nomor HP sudah benar")
		} else {
			fmt.Println(err)
		}

		avg = utils.RataRataNilai(data[index].Nilai...)
		status, torf := utils.CekStatusKelulusan(avg)
		data[index].KeteranganLulus = status
		data[index].Lulus = torf
	}

	for index, value := range data {
		fmt.Println("=======================")
		fmt.Println("Mahasiswa ke - ", index+1)
		fmt.Println("Nama:", value.Nama)
		fmt.Println("Umur:", value.Umur)
		hobbies := strings.Join(value.Hobi, ", ")
		fmt.Println("Hobi:", hobbies)
		fmt.Printf("Kota: %s, %s, %d\n", value.Address.Jalan, value.Address.Kota, value.Address.KodePos)
		if value.Nomor != nil {
			fmt.Println("Nomor:", *value.Nomor)
		} else {
			fmt.Println("Nomor: Belum diisi")
		}
		fmt.Println("Keterangan Lulus:", value.KeteranganLulus)
		fmt.Println("Lulus:", value.Lulus)
		fmt.Println("Score:", value.Nilai)
		fmt.Println("Average Score:", avg)

	}
	fmt.Println("=======================")



}
