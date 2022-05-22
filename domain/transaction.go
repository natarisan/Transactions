package domain

import (
	"codetest-docker/errs"
)

type Transaction struct {
	UserID       int      `db:"user_id"`
	Amount       int      `db:"amount"`
	Description  string   `db:"description"`
}

type TransactionRepository interface {
	Authorization(string) *errs.AppError
	RegisterTransaction(Transaction) *errs.AppError
}