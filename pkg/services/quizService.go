package services
import "strings"

import (
    "bytes"      
    "encoding/json" 
    "fmt"       
    "net/http"   
	"log"
    "io"

    "Learning-Mode-AI-quiz-service/pkg/config"
)


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
    payload := map[string]string{
        "video_id": videoID,
    }

    log.Printf("Payload sent to AI Service: %+v", payload)

    jsonData, _ := json.Marshal(payload)
    url := fmt.Sprintf("%s/ai/generate-quiz", config.AIHost)
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("failed to call AI service: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("AI service returned status code %d", resp.StatusCode)
    }

    rawBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read AI service response: %w", err)
    }

    log.Printf("Raw AI Response Body: %s", string(rawBody))

    // Parse the raw response
    var parsedResponse struct {
        Choices []struct {
            Message struct {
                Content string `json:"content"`
            } `json:"message"`
        } `json:"choices"`
    }

    if err := json.Unmarshal(rawBody, &parsedResponse); err != nil {
        return nil, fmt.Errorf("failed to parse AI response: %w", err)
    }

    if len(parsedResponse.Choices) == 0 {
        return nil, fmt.Errorf("no choices returned in AI response")
    }

    // Extract the quiz content
    content := parsedResponse.Choices[0].Message.Content
    var embeddedData struct {
        Questions []Question `json:"questions"`
    }

    if err := json.Unmarshal([]byte(content), &embeddedData); err != nil {
        return nil, fmt.Errorf("failed to parse embedded JSON content: %w", err)
    }

    aiResponse := &AIResponse{
        Questions: embeddedData.Questions,
    }

    // Add this cleanup function
    aiResponse.Questions = CleanValidQuestions(aiResponse.Questions)


    return aiResponse, nil
}



// CleanValidQuestions keeps questions with unique options, one correct answer, and no overlap with the question text.
func CleanValidQuestions(questions []Question) []Question {
    var validQuestions []Question

    for _, q := range questions {
        matchCount := 0
        seenOptions := map[string]bool{}
        uniqueOptions := []Option{}
        skipQuestion := false

        lowerQuestion := strings.ToLower(strings.TrimSpace(q.Text)) //to compare with options

        for _, opt := range q.Options {
            cleanOpt := strings.ToLower(strings.TrimSpace(opt.Option))

            // Skip if option are the same as the question
            if cleanOpt == lowerQuestion || strings.Contains(cleanOpt, lowerQuestion) || strings.Contains(lowerQuestion, cleanOpt) {
                log.Printf("Skipping question: %s", q.Text)
                skipQuestion = true
                break
            }

            // For deduplicate options
            if !seenOptions[opt.Option] {
                seenOptions[opt.Option] = true
                uniqueOptions = append(uniqueOptions, opt)
            }

            if opt.Option == q.Answer {
                matchCount++
            }
        }

        if skipQuestion {
            continue
        }

        // Keep only if the answer matches exactly one option
        if matchCount == 1 {
            q.Options = uniqueOptions
            validQuestions = append(validQuestions, q)
        } else {
            log.Printf("Not a valid question: %d — Q='%s'", matchCount, q.Text)
        }
    }

    return validQuestions
}