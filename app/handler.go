package app

import(
	"codetest-docker/dto"
	"codetest-docker/service"
	"codetest-docker/logger"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	service service.TransactionService
}

func(h TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	var transactionRequest dto.TransactionRequest
	var apiKey string

	//jsonデコード
	if err := json.NewDecoder(r.Body).Decode(&transactionRequest); err != nil {
		logger.Error("jsonリクエストをデコードする際にエラーが発生しました。" + err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	} 
	//リクエスト　バリデーションチェック
	validationErr := transactionRequest.Validate()
	if validationErr != nil {
		logger.Error("リクエストの値が不正です。")
		writeResponse(w, validationErr.Code, validationErr.Message)
		return
	}
	//トランザクション登録
	apiKey = r.Header.Get("apikey")
	transactionErr := h.service.Transaction(transactionRequest, apiKey)
	if transactionErr != nil {
		writeResponse(w, transactionErr.Code, transactionErr.Message)
		return
	} else {
		writeResponse(w, http.StatusCreated, "トランザクション登録完了")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}