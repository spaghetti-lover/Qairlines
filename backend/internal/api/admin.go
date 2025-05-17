package api

// import (
// 	"database/sql"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	db "github.com/spaghetti-lover/qairlines/db/sqlc"
// )

// type createAdminRequest struct {
// 	Owner    string `json:"owner" binding:"required"`
// 	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
// }

// // // Simple health check handler
// // func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
// // 	resp := HealthResponse{
// // 		Status:  "OK",
// // 		Version: "1.0.0",
// // 	}
// // 	w.Header().Set("Content-Type", "application/json")
// // 	w.WriteHeader(http.StatusOK)
// // 	if err := json.NewEncoder(w).Encode(resp); err != nil {
// // 		errorResponse(w, err, http.StatusInternalServerError)
// // 		return
// // 	}
// // }

// func (server *Server) handleCreateAdmin(ctx *gin.Context) {
// 	var req createAdminRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateAccountParams{
// 		Owner:    req.Owner,
// 		Currency: req.Currency,
// 		Balance:  0,
// 	}

// 	account, err := server.store.CreateAccount(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)
// }

// type getAccountRequest struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) getAccount(ctx *gin.Context) {
// 	var req getAccountRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	account, err := server.store.GetAccount(ctx, req.ID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}

// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)
// }

// type listAccountRequest struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

// func (server *Server) listAccount(ctx *gin.Context) {
// 	var req listAccountRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.ListAccountsParams{
// 		Limit:  req.PageSize,
// 		Offset: (req.PageID - 1) * req.PageSize,
// 	}

// 	accounts, err := server.store.ListAccounts(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, accounts)
// }
