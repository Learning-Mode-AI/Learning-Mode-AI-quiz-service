package services

import (
    "bytes"      
    "encoding/json" 
    "fmt"       
    "net/http"   
    "io"

    "github.com/sirupsen/logrus"
    "Learning-Mode-AI-quiz-service/pkg/config"
)

func init() {
    // Configure logrus to use JSON formatter
    logrus.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
        PrettyPrint:     false, // Set to true for development if needed
    })
    logrus.SetLevel(logrus.InfoLevel)
}

type Quiz struct {
	QuizID    string     `json:"quiz_id"`
	Questions []Question `json:"questions"`
}

type AIResponse struct {
    Questions []Question `json:"questions"`
}

type Option struct {
    Option      string `json:"option"`
    Explanation string `json:"explanation"`
}

type Question struct {
    Text    string   `json:"text"`
    Options   []Option `json:"options"`
    Answer  string   `json:"answer"`
    Timestamp string   `json:"timestamp"`
}



type RawAIResponse struct {
    Summary string `json:"summary"`
}


func FetchQuizFromAI(videoID string) (*AIResponse, error) {
    // Create quiz ID from video ID for now
    quizID := fmt.Sprintf("quiz_%s", videoID)
    
    payload := map[string]string{
        "video_id": videoID,
    }

    logrus.WithFields(logrus.Fields{
        "quiz_id": quizID,
        "payload": payload,
    }).Info("📤 Sending request to AI Service")

    jsonData, _ := json.Marshal(payload)
    url := fmt.Sprintf("%s/ai/generate-quiz", config.AIHost)
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "quiz_id": quizID,
            "error": err,
        }).Error("❌ Failed to call AI service")
        return nil, fmt.Errorf("failed to call AI service: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        logrus.WithFields(logrus.Fields{
            "quiz_id": quizID,
            "status_code": resp.StatusCode,
        }).Warn("⚠️ AI service returned non-200 status code")
        return nil, fmt.Errorf("AI service returned status code %d", resp.StatusCode)
    }

    rawBody, err := io.ReadAll(resp.Body)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "quiz_id": quizID,
            "error": err,
        }).Error("❌ Failed to read AI service response")
        return nil, fmt.Errorf("failed to read AI service response: %w", err)
    }

    logrus.WithFields(logrus.Fields{
        "quiz_id": quizID,
        "response": string(rawBody),
    }).Info("📥 Received raw response from AI service")

    // Parse the raw response
    var parsedResponse struct {
        Choices []struct {
            Message struct {
                Content string `json:"content"`
            } `json:"message"`
        } `json:"choices"`
    }

    if err := json.Unmarshal(rawBody, &parsedResponse); err != nil {
        logrus.WithFields(logrus.Fields{
            "quiz_id": quizID,
            "error": err,
        }).Error("❌ Failed to parse AI response")
        return nil, fmt.Errorf("failed to parse AI response: %w", err)
    }

    if len(parsedResponse.Choices) == 0 {
        logrus.WithFields(logrus.Fields{
            "quiz_id": quizID,
        }).Warn("⚠️ No choices returned in AI response")
        return nil, fmt.Errorf("no choices returned in AI response")
    }

    // Extract the quiz content
    content := parsedResponse.Choices[0].Message.Content
    var embeddedData struct {
        Questions []Question `json:"questions"`
    }

    if err := json.Unmarshal([]byte(content), &embeddedData); err != nil {
        logrus.WithFields(logrus.Fields{
            "quiz_id": quizID,
            "error": err,
        }).Error("❌ Failed to parse embedded JSON content")
        return nil, fmt.Errorf("failed to parse embedded JSON content: %w", err)
    }

    aiResponse := &AIResponse{
        Questions: embeddedData.Questions,
    }

    logrus.WithFields(logrus.Fields{
        "quiz_id": quizID,
        "question_count": len(aiResponse.Questions),
    }).Info("✅ Successfully generated quiz")

    return aiResponse, nil
}
