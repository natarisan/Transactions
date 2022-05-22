package service

import(
	"codetest-docker/dto"
	"codetest-docker/domain"
	"codetest-docker/errs"
)

type TransactionService interface {
	Transaction(dto.TransactionRequest, string) *errs.AppError
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}

func(s DefaultTransactionService) Transaction(req dto.TransactionRequest, apiKey string) *errs.AppError {
	//dto変換
	transaction := domain.Transaction{
		UserID:      req.UserID,
		Amount:      req.Amount,
		Description: req.Description,
	}
	//ユーザ認証
	err := s.repo.Authorization(apiKey)
	if err != nil {
		return err
	}
	//トランザクション登録処理
	transactionErr := s.repo.RegisterTransaction(transaction)
	if transactionErr != nil {
		return transactionErr
	}
	
	return nil
}