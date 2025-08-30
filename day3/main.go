package main

import (


	"GoLang/day3/models"
	"GoLang/day3/routes"

	"github.com/gin-gonic/gin"
	"errors"
	"fmt"
	"strings"
)

func CekStatusKelulusan(num float64) (string, bool) {
	switch {
	case num > 60 && num < 70:
		return "Tidak lulus", false
	case num > 70 && num < 80:
		return "Remedial", false
	case num > 80 && num < 90:
		return "Lulus", true
	case num > 90 && num < 100:
		return "Lulus dengan pujian", true
	default:
		return ">//<", false
	}
}

func RataRataNilai(score ...int) float64{
	if len(score) == 0 {
		return 0
	}

	total := 0
	for _, value := range(score){
		total += value
	}

	average := float64(total) / float64(len(score))
	return average
	
}

func ValidasiNilai(score []int) error {
	for _, value := range score {
		if value <= 0 || value >= 100 {
			return errors.New("nilai ga valid")
		}
	}
	return nil
}

func CekNoHP(num string) error {

	length := len(num)

	if length > 10 && length < 13 {
		return errors.New("nomor hp ga valid")
	}
	return nil
}

func main() {
	// http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "masa seh")
	// })

	// fmt.Println("Tes")
	// http.ListenAndServe(":8080", nil)

	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Panic :", r)
		}
	}()
	nomor1 := "081299817566"
	nomor2 := "081511428222"

	data := []models.Mahasiswa{
		{
			Nama: "Sky",
			Umur: 18,
			Hobi: []string{"Berenang", "Basket"},
			Address: models.Alamat{
				Jalan:   "Kedondong",
				Kota:    "Gedong Panjang",
				KodePos: "1224",
			},
			Nomor: &nomor1,
			Nilai: []int{88, 94, 98},
		},{
			Nama: "John Doe",
			Umur: 18,
			Hobi: []string{"Bermain", "Clubbing"},
			Address: models.Alamat{
				Jalan:   "Kedondong",
				Kota:    "Gedong Panjang",
				KodePos: "1224",
			},
			Nomor: &nomor2,
			Nilai: []int{70, 60, 98},
		},
	}

	avg := 0.0

	for index, _ := range data {
		err := ValidasiNilai(data[index].Nilai)

		if err == nil {
			fmt.Println("Yey nilai sudah benar")
		} else {
			fmt.Println(err)
		}

		err = CekNoHP("081299817566")

		if err == nil {
			fmt.Println("Yey nomor HP sudah benar")
		} else {
			fmt.Println(err)
		}
		status, torf := CekStatusKelulusan(avg)
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
		fmt.Println("Average Score:", RataRataNilai(data[index].Nilai...))

	}
	fmt.Println("=======================")


	r := gin.Default()
	routes.SetUpRoute(r, data)
	r.Run()
}