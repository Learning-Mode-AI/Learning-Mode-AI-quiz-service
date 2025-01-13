package router

import (
	"Learning-Mode-AI-quiz-service/pkg/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Quiz Service Routes
	r.HandleFunc("/quiz/generate-quiz", handlers.GenerateQuiz).Methods("POST")       
	//r.HandleFunc("/quiz/{video_id}", handlers.getQuiz).Methods("GET")                    
	//r.HandleFunc("/quiz/{video_id}", handlers.deleteQuiz).Methods("DELETE")              

	return r
}
