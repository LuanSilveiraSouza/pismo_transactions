package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/account"
	server "github.com/LuanSilveiraSouza/pismo_transactions.git/src/server"
	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/transaction"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rs/cors"
)

func main() {

	errs := make(chan error)
	ctx := context.Background()

	port := 3000
	httpAddr := flag.String("http", fmt.Sprintf(":%d", port), "http listen address")

	db, err := ConnectDatabase()
	if err != nil {
		fmt.Println("database connection error", err)
		os.Exit(-1)
	}

	err = db.AutoMigrate(&account.Account{}, &transaction.TransactionType{}, &transaction.Transaction{})
	if err != nil {
		fmt.Println("database migration error", err)
		os.Exit(-1)
	}

	err = transaction.TransactionTypeSeed(db)
	if err != nil {
		fmt.Println("database seed error", err)
		os.Exit(-1)
	}

	accountSvc := account.NewService(db)
	transactionSvc := transaction.NewService(db)

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-signalChan)
	}()

	go func() {
		fmt.Println("start http server...")

		handler := server.NewHTTPServer(ctx, accountSvc, transactionSvc)
		corsConfig := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
		})
		handler = corsConfig.Handler(handler)

		fmt.Printf("app listening on port %s\n", *httpAddr)

		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	fmt.Errorf("exit", <-errs)
}

func ConnectDatabase() (*gorm.DB, error) {
	hostname := os.Getenv("DB_HOSTNAME")

	if hostname == "" {
		hostname = "localhost"
	}

	dsn := fmt.Sprintf("user:password@tcp(%s:3306)/db?charset=utf8mb4&parseTime=True&loc=Local", hostname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
