package common

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func JSON(w http.ResponseWriter, msg interface{}) error {
	httpResponse := struct {
		Message interface{} `json:"message"`
	}{
		Message: msg,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(httpResponse)
}

func JSONErr(w http.ResponseWriter, err error, statusCode int) error {
	httpError := struct {
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Code:    statusCode,
		Status:  http.StatusText(statusCode),
		Message: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(httpError)
}

func GetParamsFromUrL(r *http.Request, params string) (int, error) {
	parsedParams := chi.URLParam(r, params)
	if len(parsedParams) == 0 {
		return 0, errors.New("nil URL params")
	}

	intParam, err := strconv.Atoi(parsedParams)
	if err != nil {
		return 0, err
	}
	return intParam, nil
}

func LoadConfig() map[string]string {
	if err := godotenv.Load("postgres.env"); err != nil {
		log.Fatal("error loading env file:", err)
	}

	var pgConfig = map[string]string{
		"host":     os.Getenv("POSTGRES_HOST"),
		"port":     os.Getenv("POSTGRES_PORT"),
		"username": os.Getenv("POSTGRES_USERNAME"),
		"password": os.Getenv("POSTGRES_PASSWORD"),
		"dbname":   os.Getenv("POSTGRES_DBNAME")}
	return pgConfig
}
