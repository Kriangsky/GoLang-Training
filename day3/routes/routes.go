package routes

import (
	"GoLang/day3/handlers"
	"GoLang/day3/models"
	"github.com/gin-gonic/gin"
)



func SetUpRoute(r *gin.Engine, data []models.Mahasiswa){
	apiGroup := r.Group("/api")
	{
		mahasiswaGroup := apiGroup.Group("/mahasiswa")
		{
			mahasiswaGroup.GET("/", handlers.ListMahasiswa)
			mahasiswaGroup.POST("/", handlers.CreateMahasiswa)
		}
	}
}

