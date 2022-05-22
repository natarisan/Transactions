package domain

import(
	"codetest-docker/errs"
	"github.com/jmoiron/sqlx"
	"database/sql"
	"sync"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

var mutex sync.Mutex

//ユーザ認証
func(d TransactionRepositoryDb) Authorization(apiKey string) *errs.AppError {
	sqlSelect := "SELECT * FROM users WHERE api_key = ?"
	row := d.client.QueryRow(sqlSelect, apiKey)
	var id      int
	var name    string
	var key     string
	err := row.Scan(&id, &name, &key)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewAuthenticationError("無効なユーザです。")
		}
		return errs.NewUnexpectedError("予期せぬDBエラー(Authorization)" + err.Error())
	}
	return nil
}

//トランザクション登録
func(d TransactionRepositoryDb) RegisterTransaction(transaction Transaction) *errs.AppError {
	mutex.Lock()
	defer mutex.Unlock()
	sumAmount := 0
	sqlSelect := "SELECT SUM(amount) FROM transactions WHERE user_id = ?"
	sqlInsert := "INSERT INTO transactions(user_id, amount, description) VALUES(?, ?, ?)"
	//そのユーザの現在の合計値を取得
	row := d.client.QueryRow(sqlSelect, transaction.UserID)
	row.Scan(&sumAmount)
	//その合計値に本トランザクションの値を足す
	sumAmount = transaction.Amount + sumAmount
	if sumAmount > 1000 {
		return errs.NewPaymentRequiredError("取引金額の合計が1000を超えてしまいます。")
	}
	
	_, err2 := d.client.Exec(sqlInsert, transaction.UserID, transaction.Amount, transaction.Description)
	if err2 != nil {
		return errs.NewUnexpectedError("予期せぬDBエラー(RegisterTransaction3)" + err2.Error())
	}
	return nil	
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}