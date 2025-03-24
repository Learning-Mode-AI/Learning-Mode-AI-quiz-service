package handlers

import (
	"Learning-Mode-AI-quiz-service/pkg/services"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type GenerateQuizRequest struct {
	VideoID string `json:"video_id"`
}

type GenerateQuizResponse struct {
	QuizID    string             `json:"quiz_id"`
	Questions []services.Question `json:"questions"`
}

// GenerateQuiz handles the /quiz/generate-quiz POST requests.
func GenerateQuizHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	logger := logrus.WithFields(logrus.Fields{
		"handler": "GenerateQuizHandler",
		"service": "quiz_service",
		"method": r.Method,
		"path": r.URL.Path,
	})

	logger.Info("🎯 Received quiz generation request")

	// Only allow POST method
	if r.Method != http.MethodPost {
		logger.WithFields(logrus.Fields{
			"received_method": r.Method,
			"expected_method": http.MethodPost,
		}).Warn("⚠️ Invalid HTTP method")
		
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var request struct {
		VideoID string `json:"video_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"duration_ms": time.Since(startTime).Milliseconds(),
		}).Error("❌ Failed to decode request body")
		
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate video ID
	if request.VideoID == "" {
		logger.WithFields(logrus.Fields{
			"duration_ms": time.Since(startTime).Milliseconds(),
		}).Warn("⚠️ Missing video_id in request")
		
		http.Error(w, "video_id is required", http.StatusBadRequest)
		return
	}

	logger.WithFields(logrus.Fields{
		"video_id": request.VideoID,
	}).Info("🎬 Processing video for quiz generation")

	// Call service to generate quiz
	quiz, err := services.FetchQuizFromAI(request.VideoID)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"video_id": request.VideoID,
			"duration_ms": time.Since(startTime).Milliseconds(),
		}).Error("❌ Failed to generate quiz")
		
		http.Error(w, "Failed to generate quiz", http.StatusInternalServerError)
		return
	}

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(quiz); err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"video_id": request.VideoID,
			"duration_ms": time.Since(startTime).Milliseconds(),
		}).Error("❌ Failed to encode response")
		
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	logger.WithFields(logrus.Fields{
		"video_id": request.VideoID,
		"question_count": len(quiz.Questions),
		"duration_ms": time.Since(startTime).Milliseconds(),
	}).Info("✅ Successfully generated and returned quiz")
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
