package routes

import (
	"GoLang/day3/handlers"
	"GoLang/day3/models"
	"github.com/gin-gonic/gin"
)



func SetUpRoute(r *gin.Engine, data []models.Mahasiswa){
	// r.GET("/api/mahasiswa", handlers.ListMahasiswa)
	r.POST("/api/mahasiswa", handlers.CreateMahasiswa)
	r.GET("/api/mahasiswa", handlers.GetMahasiswas(data))
}

