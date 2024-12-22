package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/DRAGONBORN1999/Golang_project_site/pkg/calculation"
)

var (
	ErrServer              = "Internal server error"
	ErrBadRequest          = "Bad request"
	ErrUnproccesableEntity = "Expression is not valid"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
}

type Error struct {
	Error string `json:"error"`
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":8080", nil)
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errorResponse := Error{ErrBadRequest}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)

		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		errorResponse := Error{ErrUnproccesableEntity}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)

		return

	}

	json.NewEncoder(w).Encode(Response{Result: fmt.Sprintf("%f", result)})
}
