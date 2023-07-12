package server

import (
	"context"
	"net/http"

	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/account"
	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/common"
	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/transaction"
	"github.com/gorilla/mux"
)

func ConfigureRoutes(ctx context.Context, r *mux.Router, accountSvc account.Service, transactionSvc transaction.Service) {
	subRoute := r.PathPrefix("/api/").Subrouter()

	healthRoutes(ctx, subRoute)
	accountRoutes(ctx, subRoute, accountSvc, transactionSvc)
	transactionRoutes(ctx, subRoute, transactionSvc)
}

func healthRoutes(ctx context.Context, r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		common.ResponseSuccess(w, "API Health")
	}).Methods("GET")
}

func accountRoutes(ctx context.Context, r *mux.Router, accountSvc account.Service, transactionSvc transaction.Service) {
	r.HandleFunc("/account/{id}", accountSvc.Get).Methods("GET")
	r.HandleFunc("/account/{id}/transactions", transactionSvc.List).Methods("GET")
	r.HandleFunc("/account", accountSvc.Add).Methods("POST")
}

func transactionRoutes(ctx context.Context, r *mux.Router, transactionSvc transaction.Service) {
	r.HandleFunc("/transaction", transactionSvc.Add).Methods("POST")
}
