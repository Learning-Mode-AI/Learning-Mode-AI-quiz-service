package handlers

import (
	"Youtube-Learning-Mode-Quiz-Service/pkg/services"
	"encoding/json"
	"net/http"
)

func GenerateQuiz(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		SessionID  string   `json:"session_id"`
		VideoID    string   `json:"video_id"`
		Timestamps []string `json:"timestamps"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Call the quiz service to generate the quiz
	quiz, err := services.GenerateQuiz(requestBody.SessionID, requestBody.VideoID, requestBody.Timestamps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(quiz) // Send back the generated quiz
}

func FetchQuiz(w http.ResponseWriter, r *http.Request) {
	quizID := r.URL.Query().Get("quiz_id")
	if quizID == "" {
		http.Error(w, "Missing quiz_id", http.StatusBadRequest)
		return
	}

	quiz, err := services.FetchQuiz(quizID)
	if err != nil {
		http.Error(w, "Quiz not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(quiz) // Send back the fetched quiz
}
