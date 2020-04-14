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

func NewPhraseHandler(repository repositories.IRepository) *PhraseHandler {
	return &PhraseHandler{Repository: repository}
}

func (self *PhraseHandler) Handle(context *gin.Context) {
	paging := self.getPaging(context)
	sort := self.getSort(context)
	phrases := self.Repository.GetPhraseStatsList(paging, sort)

	context.JSON(http.StatusOK, phrases)
}

func (self *PhraseHandler) getPaging(context *gin.Context) *models.Paging {
	return &models.Paging {
		Offset: self.getPathInt(context, "offset", 0),
		Limit:  self.getPathInt(context, "limit", 10),
	}
}

func (self *PhraseHandler) getSort(context *gin.Context) *models.Sort {
	fieldName := context.Query("sortBy")
	order := context.Query("order")

	var sort *models.Sort

	if fieldName != "" && order != "" {
		sort = &models.Sort {
			FieldName: fieldName,
			Order:  order,
		}
	} else {
		return nil
	}

	return sort
}

func (self PhraseHandler) getPathInt(context *gin.Context, name string, defaultValue int64) int64 {
	var result int64

	value := context.Query(name)
	if value == "" {
		result = defaultValue
	} else {
		result, _ = strconv.ParseInt(value, 10, 64)
	}

	if result < 0 {
		result = defaultValue
	}

	return result
}
