package handlers

import (
	"Learning-Mode-AI-quiz-service/pkg/services"
	"encoding/json"
	"fmt"
	"net/http"

	logrus "github.com/sirupsen/logrus"
)

type GenerateQuizRequest struct {
	VideoID string `json:"video_id"`
	UserID  string `json:"user_id"`
}

type GenerateQuizResponse struct {
	QuizID    string              `json:"quiz_id"`
	Questions []services.Question `json:"questions"`
}

// Initialize a logger with basic configuration
var logger = logrus.New()

// GenerateQuiz handles the /quiz/generate-quiz POST requests.
func GenerateQuiz(w http.ResponseWriter, r *http.Request) {
	var req GenerateQuizRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		logger.WithFields(logrus.Fields{
			"error":   err.Error(),
			"video_id": req.VideoID,
			"user_id": req.UserID,
		}).Error("Error decoding request")
		return
	}

	// Generate a QuizID using VideoID and UserID
	quizID := fmt.Sprintf("quiz_%s_%s", req.UserID, req.VideoID)

	// Check if the quiz exists in Redis
	quiz, err := services.GetQuizFromRedis(req.UserID, req.VideoID)
	if err != nil {
		http.Error(w, "Failed to fetch quiz from Redis", http.StatusInternalServerError)
		logger.WithFields(logrus.Fields{
			"error":    err.Error(),
			"video_id": req.VideoID,
			"user_id":  req.UserID,
		}).Error("Error fetching quiz from Redis")
		return
	}

	// If quiz exists in Redis
	if quiz != nil {
		logger.WithFields(logrus.Fields{
			"video_id": req.VideoID,
			"user_id":  req.UserID,
			"quiz_id":  quizID,
		}).Info("Quiz retrieved from Redis successfully")
		sendQuizResponse(w, quizID, quiz.Questions)
		return
	}

	// If quiz does not exist in Redis, fetch from AI service
	quiz, err = services.FetchQuizFromAI(req.VideoID, req.UserID)
	if err != nil {
		http.Error(w, "Failed to generate quiz", http.StatusInternalServerError)
		logger.WithFields(logrus.Fields{
			"error":    err.Error(),
			"video_id": req.VideoID,
			"user_id":  req.UserID,
		}).Error("Error fetching quiz from AI service")
		return
	}

	// Store the fetched quiz in Redis
	if err := services.StoreQuizInRedis(req.UserID, req.VideoID, quiz); err != nil {
		logger.WithFields(logrus.Fields{
			"error":    err.Error(),
			"video_id": req.VideoID,
			"user_id":  req.UserID,
			"quiz_id":  quizID,
		}).Warn("Error storing quiz in Redis. Continuing with response.")
		// Continue responding even if Redis storage fails
	}

	// Respond with the fetched quiz
	sendQuizResponse(w, quizID, quiz.Questions)
}

func sendQuizResponse(w http.ResponseWriter, quizID string, questions []services.Question) {
	response := GenerateQuizResponse{
		QuizID:    quizID,
		Questions: questions,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error encoding response")
	}
}
