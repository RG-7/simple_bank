package api

import (
	"database/sql"
	"net/http"

	db "github.com/RG-7/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR INR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	// if invalid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonsponse(err))
		return
	}

	// if valid
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResonsponse(err))
	}

	ctx.JSON(http.StatusOK, account)

}

// ------------------- Get Account By ID ------------------------
// struct for get request
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// func to get account by id!
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	// if invalid
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonsponse(err))
		return
	}

	// if valid
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResonsponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResonsponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// ------------------- Get List of Account ------------------------
// struct to get list of account request
type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// fucn to get lis tof account
func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonsponse(err))
		return
	}

	// if valid
	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResonsponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
