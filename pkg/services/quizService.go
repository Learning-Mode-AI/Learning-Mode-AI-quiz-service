package services

import (
	"Youtube-Learning-Mode-Quiz-Service/pkg/services/redisService"
	"fmt"

	"github.com/openai/openai-go-sdk"
)

func GenerateQuiz(sessionID, videoID string, timestamps []string) (map[string]interface{}, error) {
	client := openai.NewClient("<OPENAI_API_KEY>")

	// Placeholder for future OpenAI Assistant integration
	prompt := fmt.Sprintf("Generate a quiz for video ID %s using timestamps %v. Include MC and TF questions.", videoID, timestamps)

	req := openai.CompletionRequest{
		Model:     "text-davinci-003",
		Prompt:    prompt,
		MaxTokens: 1000,
	}

	resp, err := client.Completions.Create(req)
	if err != nil {
		return nil, err
	}

	quiz := map[string]interface{}{
		"quiz": resp.Choices[0].Text,
	}

	// Store the quiz in Redis
	err = redisService.StoreQuiz(sessionID, quiz)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func FetchQuiz(sessionID string) (map[string]interface{}, error) {
	// Fetch quiz from Redis
	quiz, err := redisService.FetchQuiz(sessionID)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}
