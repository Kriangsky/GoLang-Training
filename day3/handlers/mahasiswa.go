package handlers

import (
	"GoLang/day3/models"
	"GoLang/day3/schemas"
	"GoLang/day3/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var modelMahasiswa []models.Mahasiswa

func ListMahasiswa(c *gin.Context) {
	var response []schemas.ListMahasiswaResponse

	for _, mhs := range(modelMahasiswa){
		rata2 := utils.RataRataNilai(mhs.Nilai...)
		keteranganLulus, _ := utils.CekStatusKelulusan(rata2)

		response = append(response, schemas.ListMahasiswaResponse{
			ID: mhs.ID,
			Nama: mhs.Nama,
			RataRataNilai: rata2,
			KeteranganLulus: keteranganLulus,
		})
	}
	c.JSON(200, response)


}

func CreateMahasiswa(c *gin.Context) {
	var input schemas.CreateMahasiswaRequest
	err := c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = utils.CekNoHP(input.NoHP)
	if err != nil{
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	err = utils.ValidasiNilai(input.Nilai)
	if err != nil{
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	rataNilai := utils.RataRataNilai(input.Nilai...)
	keteranganLulus, tof := utils.CekStatusKelulusan(rataNilai)

	idMhs := len(modelMahasiswa) + 1

	newMahasiswa := models.Mahasiswa{
		ID : idMhs,
		Nama : input.Nama,
		Umur : input.Umur,
		Hobi : input.Hobi,
		Address : models.Alamat{
			Jalan : input.Alamat.Jalan,
			Kota : input.Alamat.Kota,
			KodePos: input.Alamat.KodePos,
		},
		Nomor : input.NoHP,
		Nilai : input.Nilai,
		Lulus : tof,
		KeteranganLulus: keteranganLulus,
	}

	modelMahasiswa = append(modelMahasiswa, newMahasiswa)

	c.JSON(http.StatusOK, newMahasiswa)
}

func GetMahasiswas(data []models.Mahasiswa) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, data)
	}
}