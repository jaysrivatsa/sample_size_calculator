package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jaysrivatsa/sample_size_calculator/models"
	"github.com/jaysrivatsa/sample_size_calculator/utils"
)

func addtiveHandler(ctx *gin.Context) {
	var request models.AdditiveRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "errot": "unable to parse request"})
		return
	}

	sampleSizeRequired := utils.SampleSizeAdditive(request)

	ctx.JSON(http.StatusOK, gin.H{"message": "Sample size computed", "data": sampleSizeRequired})
}
