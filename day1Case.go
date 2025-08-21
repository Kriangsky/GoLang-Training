package main

import "fmt"
import "strings"

func main(){
	type Alamat struct{
		jalan string
		kota string
		kodePos int
	}

	type Mahasiswa struct{
		nama string
		umur int
		hobi []string
		alamat Alamat
		nomor *string
	}
	
	nomorHp := "081299817566"

	data := []Mahasiswa{
		{
			nama : "Sky",
			umur : 12,
			hobi : []string{"Berenang", "Bermain"},
			alamat : Alamat{
				jalan : "Pluit timur",
				kota : "Jakarta",
				kodePos : 14450,
			},
		},{
			nama : "Peter",
			umur : 12,
			hobi : []string{"Hockey", "Bermain"},
			alamat : Alamat{
				jalan : "Pluit utara",
				kota : "Tangerang",
				kodePos : 14450,
			},
			nomor : &nomorHp,
		},
	}

	for index, value := range(data){
		fmt.Println("=============================")
		fmt.Println("Mahasiswa ke -", index+1)
		fmt.Printf("Nama : %s\n", value.nama)
		fmt.Printf("Umur : %d\n", value.umur)
		gabungHobi := strings.Join(value.hobi, ", ")
		fmt.Printf("Hobi :%s\n", gabungHobi)
		fmt.Printf("Alamat :%s, %s, %d\n", value.alamat.jalan, value.alamat.kota, value.alamat.kodePos)
		if value.nomor == nil{
			fmt.Printf("Nomor HP : Belum Diisi\n")
		} else {
			fmt.Printf("Nomor HP : %d\n", *value.nomor)
		}
	}
	
}