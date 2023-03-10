package api

import (
	"database/sql"
	db "github.com/francivaldo-alves/gofinance-bankend/db/sqlc"
	"github.com/francivaldo-alves/gofinance-bankend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createCategoryRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req createCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	} else {

		arg := db.CreateCategoryParams{
			UserID:      req.UserID,
			Title:       req.Title,
			Type:        req.Type,
			Description: req.Description,
		}

		category, err := server.store.CreateCategory(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}

		ctx.JSON(http.StatusOK, category)
	}

}

// Funcação da PI para buscar uma category
type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req getCategoryRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	category, err := server.store.GetCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type getCategoriesRequest struct {
	UserID      int32  `json:"user_id" binding:"required" `
	Type        string `json:"type" binding:"required"`
	Title       string `json:"title" `
	Description string `json:"description"`
}

// Funcação da PI para buscar uma category
func (server *Server) getCategories(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req getCategoriesRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCategoriesParams{
		UserID:      req.UserID,
		Type:        req.Type,
		Title:       req.Title,
		Description: req.Description,
	}

	categories, err := server.store.GetCategories(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, categories)
}

type updateCategoryRequest struct {
	ID          int32  `json:"id" binding: "required`
	Title       string `json:"title" `
	Description string `json:"description" `
}

// Funcação da PI para cadastar uma categoria
func (server *Server) updateCategory(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req updateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}
	arg := db.UpdateCategoryParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	}
	category, err := server.store.UpdateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, category)
}

type deleteCategoryByRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

// Funcação da PI para deletar uma categoria
func (server *Server) deleteCategory(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req deleteCategoryByRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	err = server.store.DeleteCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, true)
}
