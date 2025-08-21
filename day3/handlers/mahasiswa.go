package handlers

import (
	"GoLang/day3/models"

	"github.com/gin-gonic/gin"
)

func ListMahasiswa(c *gin.Context) {
	c.String(200, "Showing Mahasiswas")
}

func CreateMahasiswa(c *gin.Context) {
	c.String(200, "Create data mahasiswa baru")
}

func GetMahasiswas(data []models.Mahasiswa) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, data)
	}
}