package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type Sample struct{
	Paragraph string `json:"paragraph"`
}

func getSample(ctx *gin.Context) {
	var sample Sample
	if err:=ctx.ShouldBindJSON(&sample);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"paragraph":sample.Paragraph})
}

func main(){
	engine := gin.Default()
	engine.POST("/", getSample)
	engine.Run(":8080")
}