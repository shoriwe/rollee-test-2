package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/rollee-test-2/models"
)

func (h *Handler) AddWord(ctx *gin.Context) {
	var word models.Word
	bErr := ctx.Bind(&word)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: bErr.Error()})
	}
	qErr := h.c.AddWord(&word)
	if qErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: qErr.Error()})
	}
	ctx.JSON(http.StatusOK, Status{Succeed: true})
}

func (h *Handler) QueryWord(ctx *gin.Context) {
	word, qErr := h.c.QueryWord(ctx.Param(PatternParam))
	if qErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: qErr.Error()})
	}
	if word == nil {
		ctx.JSON(http.StatusOK, models.Word{IsNull: true})
		ctx.Done()
		return
	}
	ctx.JSON(http.StatusOK, word)
	ctx.Done()
}
