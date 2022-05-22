package app

import(
	"codetest-docker/domain"
	"codetest-docker/service"
	"codetest-docker/logger"
	"log"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ServerAddr string
	ServerPort int
	DbUser     string
	DbAddr     string
	DbName     string
}

var Config ConfigList

func Init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	Config = ConfigList{
		ServerAddr: cfg.Section("server").Key("serverAddr").String(),
		ServerPort: cfg.Section("server").Key("serverPort").MustInt(),
		DbUser:     cfg.Section("db").Key("dbUser").String(),
		DbAddr:     cfg.Section("db").Key("dbAddr").String(),
		DbName:     cfg.Section("db").Key("dbName").String(),
	}
}

func Go() {
	Init()
	router := mux.NewRouter()
	transactionRepository := domain.NewTransactionRepositoryDb(getDbClient())
	th := TransactionHandler{service.NewTransactionService(transactionRepository)}
	router.HandleFunc("/transactions", th.Transaction).Methods(http.MethodPost)
	serverAddr := Config.ServerAddr
	serverPort := Config.ServerPort
	logger.Info(fmt.Sprintf("サーバアドレス%s,ポート%d", serverAddr, serverPort))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", serverAddr, serverPort), router); err != nil {
		log.Println(err)
	}
}

func getDbClient() *sqlx.DB {
	dbUser := Config.DbUser
	dbAddr := Config.DbAddr
	dbName := Config.DbName
	dataSource := fmt.Sprintf("%s@tcp(%s)/%s", dbUser, dbAddr, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	return client
}