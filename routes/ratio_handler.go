package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaysrivatsa/sample_size_calculator/models"
	"github.com/jaysrivatsa/sample_size_calculator/utils"
)

func ratioHandler(ctx *gin.Context) {
	var request models.RatioRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "errot": "unable to parse request"})
		return
	}

	sampleSizeRequired := utils.SampleSizeRatio(request)

	ctx.JSON(http.StatusOK, gin.H{"message": "Sample size computed", "data": sampleSizeRequired})
}
