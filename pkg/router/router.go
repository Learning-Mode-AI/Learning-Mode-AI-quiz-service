package router

import (
	"Youtube-Learning-Mode-Quiz-Service/pkg/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/generate-quiz", handlers.GenerateQuiz).Methods("POST")
	r.HandleFunc("/fetch-quiz", handlers.FetchQuiz).Methods("GET")

	return r
}
