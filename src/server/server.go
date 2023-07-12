package server

import (
	"context"
	"net/http"
	"time"

	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/account"
	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/transaction"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, accountSvc account.Service, transactionSvc transaction.Service) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	ConfigureRoutes(ctx, r, accountSvc, transactionSvc)
	return http.TimeoutHandler(r, time.Second*30, "Timeout!")
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		rw.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token, x-fetch")
		rw.Header().Add("Access-Control-Expose-Headers", "Authorization")
		rw.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(rw, r)
	})
}
