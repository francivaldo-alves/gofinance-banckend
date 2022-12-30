package api

import (
	"database/sql"
	"github.com/francivaldo-alves/gofinance-bankend/util"
	"net/http"
	"time"

	db "github.com/francivaldo-alves/gofinance-bankend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	UserID      int32     `json:"user_id" binding:"required"`
	CategoryID  int32     `json:"category_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Value       int32     `json:"value" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

// Funcação da PI para cadastar um account
func (server *Server) createAccount(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}
	var categoryId = req.CategoryID
	var accountType = req.Type
	category, err := server.store.GetCategory(ctx, categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	var categoryTypeIsDiffentOdAccountType = category.Type != accountType
	if categoryTypeIsDiffentOdAccountType {
		ctx.JSON(http.StatusBadRequest, "Account type is diferente od Category type")
	} else {

		arg := db.CreateAccountParams{
			UserID:      req.UserID,
			CategoryID:  categoryId,
			Title:       req.Title,
			Type:        accountType,
			Description: req.Description,
			Value:       req.Value,
			Date:        req.Date,
		}
		account, err := server.store.CreateAccount(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		ctx.JSON(http.StatusOK, account)
	}

}

// Funcação da PI para buscar uma account
type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}

	var req getAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountsRequest struct {
	UserID      int32     `json:"user_id" binding:"required" `
	CategoryID  int32     `json:"category_id"`
	Type        string    `json:"type" binding: required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" `
}

// Funcação da PI para buscar uma account
func (server *Server) getAccounts(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req getAccountsRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAccountsParams{
		UserID: req.UserID,
		Type:   req.Type,
		CategoryID: sql.NullInt32{
			Int32: req.CategoryID,
			Valid: req.CategoryID > 0,
		},
		Title:       req.Title,
		Description: req.Description,
		Date: sql.NullTime{
			Time:  req.Date,
			Valid: !req.Date.IsZero(),
		},
	}
	accounts, err := server.store.GetAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type updateAccountRequest struct {
	ID          int32  `json:"id" `
	Title       string `json:"title" `
	Description string `json:"description" `
	Value       int32  `json:"value"`
}

// Funcação da PI para cadastar uma categoria
func (server *Server) updateAccount(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req updateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}
	arg := db.UpdateAccountParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Value:       req.Value,
	}
	account, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}

type deleteAccountByRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

// Funcação da PI para deletar uma categoria
func (server *Server) deleteAccount(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req deleteAccountByRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	err = server.store.DeleteAccount(ctx, req.ID)
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

type GetAccountsGraphParams struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountGraph(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}

	var req GetAccountsGraphParams
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsGraphParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	countGraph, err := server.store.GetAccountsGraph(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, countGraph)
}

type GetAccountsReportsParams struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountRports(ctx *gin.Context) {
	erroOnValidetToken := util.GetTokenInHeaderAndVerify(ctx)
	if erroOnValidetToken != nil {
		return
	}
	var req GetAccountsReportsParams
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsReportsParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	sumReports, err := server.store.GetAccountsReports(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sumReports)
}
