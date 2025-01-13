package handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"Learning-Mode-AI-quiz-service/pkg/services"
)

type GenerateQuizRequest struct {
	VideoID string `json:"video_id"`
}

type GenerateQuizResponse struct {
	QuizID    string             `json:"quiz_id"`
	Questions []services.Question `json:"questions"`
}

// GenerateQuiz handles the /quiz/generate-quiz POST requests.
func GenerateQuiz(w http.ResponseWriter, r *http.Request) {
	var req GenerateQuizRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request:", err)
		return
	}

	// Check if the quiz exists in Redis
	quiz, err := services.GetQuizFromRedis(req.VideoID)
if err != nil {
    http.Error(w, "Failed to fetch quiz from Redis", http.StatusInternalServerError)
    log.Println("Error fetching quiz from Redis:", err)
    return
}

// If quiz exists in Redis
if quiz != nil {
    log.Printf("Quiz retrieved from Redis for VideoID: %s", req.VideoID)
    sendQuizResponse(w, req.VideoID, quiz.Questions)
    return
}

// If quiz does not exist in Redis, fetch from AI service
quiz, err = services.FetchQuizFromAI(req.VideoID)
if err != nil {
    http.Error(w, "Failed to generate quiz", http.StatusInternalServerError)
    log.Println("Error fetching quiz from AI service:", err)
    return
}

// Store the fetched quiz in Redis
if err := services.StoreQuizInRedis(req.VideoID, quiz); err != nil {
    log.Println("Error storing quiz in Redis:", err)
    // Continue responding even if Redis storage fails
}


	// Respond with the fetched quiz
	sendQuizResponse(w, req.VideoID, quiz.Questions)
}

func sendQuizResponse(w http.ResponseWriter, quizID string, questions []services.Question) {
	response := GenerateQuizResponse{
		QuizID:    quizID,
		Questions: questions,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
	}
}
