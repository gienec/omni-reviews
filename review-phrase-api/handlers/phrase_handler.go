package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phrase-api/models"
	"phrase-api/repositories"
	"strconv"
)

type PhraseHandler struct {
	Repository repositories.IRepository
}

func (phraseHandler PhraseHandler) Handle(context *gin.Context) {
	paging := phraseHandler.getPaging(context)
	sort := phraseHandler.getSort(context)
	phrases := phraseHandler.Repository.GetPhraseStatsList(paging, sort)

	context.JSON(http.StatusOK, phrases)
}

func (phraseHandler PhraseHandler) getPaging(context *gin.Context) models.Paging {
	return models.Paging {
		Offset: phraseHandler.getPathInt(context, "offset", 0),
		Limit:  phraseHandler.getPathInt(context, "limit", 0),
	}
}

func (phraseHandler PhraseHandler) getSort(context *gin.Context) models.Sort {
	fieldName := context.Params.ByName("sortBy")
	order := context.Params.ByName("order")

	return models.Sort {
		FieldName: fieldName,
		Order:  order,
	}
}

func (phraseHandler PhraseHandler) getPathInt(context *gin.Context, name string, defaultValue int) int {
	var result int

	value := context.Params.ByName(name)
	if value == "" {
		result = defaultValue
	} else {
		result, _ = strconv.Atoi(value)
	}

	return result
}
