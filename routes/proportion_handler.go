package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jaysrivatsa/sample_size_calculator/models"
	"github.com/jaysrivatsa/sample_size_calculator/utils"
	"log"
	"net/http"
)

func proportionHandler(ctx *gin.Context) {
	var request models.ProportionRequest
	err := ctx.ShouldBind(&request)
	log.Println(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "errot": "unable to parse request"})
		return
	}

	sampleSizeRequired := utils.SampleSizeProportion(request)

	ctx.JSON(http.StatusOK, gin.H{"message": "Sample size computed", "data": sampleSizeRequired})
}
