package dto

import(
	"codetest-docker/errs"
)

type TransactionRequest struct {
	UserID       int    `json:"user_id"`
	Amount       int    `json:"amount"`
	Description  string `json:"description"`
}

func(r TransactionRequest) Validate() *errs.AppError {
	if r.UserID <= 0 {
		return errs.NewValidationError("ユーザIDが0以下です。")
	}
	if r.Amount > 1000 {
		return errs.NewValidationError("値段が1000を超えています。")
	} else if r.Amount <= 0 {
		return errs.NewValidationError("値段を0以下にすることはできません。")
	}
	return nil
}
