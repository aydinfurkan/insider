package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"insider/src/config"
	"io"
	"net/http"
	"time"
)

type WebhookRequest struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

type WebhookResponse struct {
	Message   string `json:"message"`
	MessageId string `json:"messageId"`
}

type WebhookService struct {
	webhookURL   string
	authKey      string
	client       *http.Client
	redisService *RedisService
}

func NewWebhookService(cfg *config.ConfigType, redisService *RedisService) *WebhookService {
	return &WebhookService{
		webhookURL:   cfg.WEBHOOK_URL,
		redisService: redisService,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (ws *WebhookService) SendMessage(to, content string) (*WebhookResponse, error) {
	payload := WebhookRequest{
		To:      to,
		Content: content,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequest("POST", ws.webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := ws.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send webhook request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("webhook request failed with status: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var webhookResponse WebhookResponse
	if err := json.Unmarshal(responseBody, &webhookResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	err = ws.redisService.Set(webhookResponse.MessageId, time.Now().String(), 24*time.Hour)

	return &webhookResponse, nil
}
