package routes

import (
	"encoding/json"
	"net/http"
	"time"
	"strings"
	"github.com/jjyamangg/go_transactions_postgresql/database"
	"github.com/jjyamangg/go_transactions_postgresql/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API for the Transactions!"))
}

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "JSON invalid: "+err.Error(), http.StatusBadRequest)
		return
	}
	if transaction.Amount <= 0 {
		http.Error(w, "The amount must be greater than 0.", http.StatusBadRequest)
		return
	}

	var currency models.Currency
	if err := database.DB.First(&currency, transaction.CurrencyID).Error; err != nil {
		http.Error(w, "currency_id invalid or not found", http.StatusBadRequest)
		return
	}

	var txType models.TransactionType
	if err := database.DB.First(&txType, transaction.TypeTransactionID).Error; err != nil {
		http.Error(w, "type_transaction_id invalid or not found", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		http.Error(w, "Error saving transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&transaction)
}

func GetTransactionsByDateRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Leer los parámetros from y to
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	if fromStr == "" || toStr == "" {
		http.Error(w, "You must provide ‘from’ and ‘to’ parameters in YYYYY-MM-DD format", http.StatusBadRequest)
		return
	}

	from, err1 := time.Parse("2006-01-02", fromStr)
	to, err2 := time.Parse("2006-01-02", toStr)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid date format. Use YYYYY-MM-DD", http.StatusBadRequest)
		return
	}

	// Buscar transacciones en el rango
	var transactions []models.Transaction
	if err := database.DB.
		Preload("TypeTransaction").
		Where("created_at BETWEEN ? AND ?", from, to).
		Order("created_at ASC").
		Find(&transactions).Error; err != nil {
		http.Error(w, "Error while querying transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Construir respuesta con saldo acumulado
	var response []models.TransactionResponse
	var runningBalance int

	for _, tx := range transactions {
		if strings.ToLower(tx.TypeTransaction.Name) == "deposit" {
			runningBalance += tx.Amount
		} else {
			runningBalance -= tx.Amount
		}

		response = append(response, models.TransactionResponse{
			ID:                tx.ID,
			AccountID:         tx.AccountID,
			Amount:            tx.Amount,
			CurrencyID:        tx.CurrencyID,
			TypeTransactionID: tx.TypeTransactionID,
			Description:       tx.Description,
			CreatedAt:         tx.CreatedAt,
			RunningBalance:    runningBalance,
		})
	}

	for i, j := 0, len(response)-1; i < j; i, j = i+1, j-1 {
		response[i], response[j] = response[j], response[i]
	}

	json.NewEncoder(w).Encode(response)
}